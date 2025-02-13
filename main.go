package main

import (
	"embed"

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
