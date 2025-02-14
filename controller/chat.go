package controller

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/sashabaranov/go-openai"
	"net/http"
)

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

func ChatSSE(c *gin.Context) {
	// 设置SSE headers
	c.Writer.Header().Set("Content-Type", "text/event-stream")
	c.Writer.Header().Set("Cache-Control", "no-cache")
	c.Writer.Header().Set("Connection", "keep-alive")

	// 读取请求体
	var message Message
	if err := c.BindJSON(&message); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config := openai.DefaultConfig("sk-GLLjoRQHbrOf9mniE38212791569423eAa1e6c7e111cB6D6")
	config.BaseURL = "http://oneapi.chenfeng.info/v1/"
	// 创建OpenAI客户端
	client := openai.NewClientWithConfig(config)

	// 准备OpenAI请求
	req := openai.ChatCompletionRequest{
		Model: "Qwen2.5-32B-Instruct",
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    message.Role,
				Content: message.Content,
			},
		},
		Stream: true,
	}

	// 创建流式响应
	stream, err := client.CreateChatCompletionStream(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer stream.Close()

	// 通知客户端不要缓存响应
	c.Writer.Flush()

	// 读取流式响应并发送到SSE
	for {
		response, err := stream.Recv()
		if err != nil {
			c.Writer.Write([]byte("data: [DONE]\n\n"))
			c.Writer.Flush()
			return
		}

		if len(response.Choices) > 0 {
			content := response.Choices[0].Delta.Content
			if content != "" {
				// 发送SSE消息
				c.Writer.Write([]byte("data: " + content + "\n\n"))
				c.Writer.Flush()
			}
		}
	}
}
