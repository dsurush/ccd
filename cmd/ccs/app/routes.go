package app

import (
	"ccs/middleware/corss"
	"ccs/middleware/logger"
	"ccs/settings"
	"fmt"
	"log"
	"net/http"
)

func (server *MainServer) InitRoutes() {
	fmt.Println("Init routes")

	server.router.POST("/api/login", logger.Logger(`Create Token for user:`)(corss.Middleware(server.LoginHandler)))
	server.router.GET(`/api/users`, logger.Logger(`Get all users:`)(server.GetUsersHandler))
	server.router.GET(`/api/users/:id`, logger.Logger(`Get all user by id:`)(server.GetUserByIdHandler))

	settings.AppSettings = settings.ReadSettings("./settings.json")
	port := fmt.Sprintf(":%d", settings.AppSettings.AppParams.PortRun)
	log.Println(http.ListenAndServe(port, server))
}
