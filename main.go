package main

import (
	"embed"
	"github.com/gin-contrib/cors"
	"time"

	"gin-template/router"

	"github.com/gin-gonic/gin"
	"log"
	"os"
)

//go:embed web/dist
var buildFS embed.FS

//go:embed web/dist/index.html
var indexPage []byte

var targetPath = "web/dist"

func main() {
	//if os.Getenv("GIN_MODE") != "debug" {
	//	gin.SetMode(gin.ReleaseMode)
	//}

	// Initialize HTTP server
	server := gin.Default()

	// 使用 cors.New() 配置 CORS
	config := cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},                            // 允许的来源
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}, // 允许的方法
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},          // 允许的头部
		ExposeHeaders:    []string{"Content-Length"},                                   // 允许客户端访问的响应头
		AllowCredentials: true,                                                         // 是否允许发送 Cookie
		MaxAge:           12 * time.Hour,                                               // 预检请求的有效期
	}
	server.Use(cors.New(config)) // 应用 CORS 配置

	router.SetRouter(server, buildFS, indexPage, targetPath)
	var port = os.Getenv("PORT")
	if port == "" {
		port = "2486"
	}
	err := server.Run(":" + port)
	if err != nil {
		log.Println(err)
	}
}
