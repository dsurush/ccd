package app

import (
	"ccs/middleware/logger"
	"ccs/settings"
	"fmt"
	"log"
	"net/http"
)

func (server *MainServer) InitRoutes() {
	fmt.Println("Init routes")

	server.router.POST("/api/login", logger.Logger(`Create Token for user:`)(server.LoginHandler))
//	server.router.POST(`/api/login`, logger.Logger())
	settings.AppSettings = settings.ReadSettings("./settings.json")
	port := fmt.Sprintf(":%d", settings.AppSettings.AppParams.PortRun)
	fmt.Printf("Server is listening port %s...\n", port)
	log.Println(http.ListenAndServe(port, server))
}
