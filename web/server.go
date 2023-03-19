package web

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/mtricolici/file-hub-go/config"
	"github.com/mtricolici/file-hub-go/web/handlers"
)

var (
	server *gin.Engine
)

func setWebRoutes() {
	server.GET("/", handlers.Homepage)
}

func startWebServer() {
	ip := config.Get().ListenInterface()
	port := config.Get().ListenPort()

	listenOn := fmt.Sprintf("%s:%d", ip, port)
	fmt.Printf("Starting web server at %s ...\n", listenOn)

	if err := server.Run(listenOn); err != nil {
		panic(err)
	}
}

func InitAndStart() {
	gin.SetMode(gin.ReleaseMode)
	server = gin.New()
	server.LoadHTMLGlob(config.Get().TemplatesDir())

	setWebRoutes()
	startWebServer()
}
