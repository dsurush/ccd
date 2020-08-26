package app

import (
	"ccs/middleware/authorized"
	"ccs/middleware/corss"
	"ccs/middleware/jwt"
	"ccs/middleware/logger"
	"ccs/settings"
	"ccs/token"
	"fmt"
	"log"
	"net/http"
	"reflect"
)

func (server *MainServer) InitRoutes() {
	fmt.Println("Init routes")

	server.router.POST("/api/login", logger.Logger(`Create Token for user:`)(corss.Middleware(server.LoginHandler)))

	server.router.GET(`/api/users`, logger.Logger(`Get all users:`)(corss.Middleware(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))(authorized.Authorized([]string{`admin`}, jwt.FromContext)(server.GetUsersHandler)))))
	server.router.GET(`/api/users/:id`, logger.Logger(`Get all user by id:`)(corss.Middleware(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))(authorized.Authorized([]string{`admin`}, jwt.FromContext)(server.GetUserByIdHandler)))))

	server.router.POST(`/api/users/add`, server.AddNewUser)
	server.router.POST(`/api/users/edit/:id`, server.EditUser)

	settings.AppSettings = settings.ReadSettings("./settings.json")
	port := fmt.Sprintf(":%d", settings.AppSettings.AppParams.PortRun)
	log.Println(http.ListenAndServe(port, server))
}
