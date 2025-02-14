import React from 'react';
import { useXChat, useXAgent,Bubble,Sender,XRequest,XStream } from '@ant-design/x';
import axios from 'axios';

async function fetchSSEStream(message: string) {
  const response = await fetch('/api/chat', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({ message }), // 假设需要发送message作为请求体
  });

  if (!response.ok) {
    throw new Error(`HTTP error! status: ${response.status}`);
  }

  if (!response.body) {
    throw new Error('Response body is null');
  }

  return new ReadableStream({
    async start(controller) {
      // @ts-ignore
      const reader = response.body.getReader();
      const decoder = new TextDecoder();
      let buffer = '';

      try {
        while (true) {
          const { done, value } = await reader.read();

          if (done) {
            if (buffer) {
              // 处理最后可能残留的数据
              const events = buffer.split('\n\n').filter(Boolean);
              for (const event of events) {
                controller.enqueue(event + '\n\n');
              }
            }
            break;
          }

          // 将新接收的数据添加到缓冲区
          buffer += decoder.decode(value, { stream: true });

          // 按SSE格式分割事件
          const events = buffer.split('\n\n');

          // 保留最后一个可能不完整的事件
          buffer = events.pop() || '';

          // 发送完整的事件
          for (const event of events) {
            if (event.trim()) {
              controller.enqueue(event + '\n\n');
            }
          }
        }
      } catch (error) {
        controller.error(error);
      } finally {
        controller.close();
        reader.releaseLock();
      }
    }
  });
}

const ChatBot = () => {
  const [lines, setLines] = React.useState<Record<string, string>[]>([]);
  const [agent] = useXAgent({
    // 模型推理服务地址
    baseURL: 'http://',
    // 模型名称
    model: 'gpt-3.5',
  });

  const {
    // 发起聊天请求
    onRequest,
    // 消息列表
    messages,
  } = useXChat({ agent });



  async function readStream() {
    // 🌟 Read the stream
    // @ts-ignore
    for await (const chunk of XStream({
      readableStream: fetchSSEStream('hello'),
    })) {
      console.log(chunk);
      setLines((pre) => [...pre, chunk]);
    }
  }

  return (
    <div>
      <Bubble.List items={messages} />
      <Sender onSubmit={readStream} />
    </div>
  );
};

export default ChatBot;
