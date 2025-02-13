package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(c *gin.Context) {
	var loginRequest LoginRequest
	err := json.NewDecoder(c.Request.Body).Decode(&loginRequest)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "无效的参数",
			"success": false,
		})
		return
	}
	username := loginRequest.Username
	password := loginRequest.Password
	if username == "" || password == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "无效的参数",
			"success": false,
		})
		return
	}

	// 此处直接登录成功
	c.JSON(http.StatusOK, gin.H{
		"message":          "登录成功",
		"success":          true,
		"status":           "ok",
		"currentAuthority": "admin",
		"type":             "account",
		"token":            "xxxxxx",
	})
}

func CurrentUser(c *gin.Context) {
	jsonStr := `
{
    "success": true,
    "data": {
        "name": "Serati Ma",
        "avatar": "https://gw.alipayobjects.com/zos/antfincdn/XAosXuNZyF/BiazfanxmamNRoxxVxka.png",
        "userid": "00000001",
        "email": "antdesign@alipay.com",
        "signature": "海纳百川，有容乃大",
        "title": "交互专家",
        "group": "蚂蚁金服－某某某事业群－某某平台部－某某技术部－UED",
        "tags": [
            {
                "key": "0",
                "label": "很有想法的"
            },
            {
                "key": "1",
                "label": "专注设计"
            },
            {
                "key": "2",
                "label": "辣~"
            },
            {
                "key": "3",
                "label": "大长腿"
            },
            {
                "key": "4",
                "label": "川妹子"
            },
            {
                "key": "5",
                "label": "海纳百川"
            }
        ],
        "notifyCount": 12,
        "unreadCount": 11,
        "country": "China",
        "access": "admin",
        "geographic": {
            "province": {
                "label": "浙江省",
                "key": "330000"
            },
            "city": {
                "label": "杭州市",
                "key": "330100"
            }
        },
        "address": "西湖区工专路 77 号",
        "phone": "0752-268888888"
    }
}
`
	c.Header("Content-Type", "application/json")
	c.String(http.StatusOK, jsonStr)
}
