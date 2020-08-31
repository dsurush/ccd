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
	"time"
)

func (server *MainServer) InitRoutes() {
	fmt.Println("Init routes")
	test()
	server.router.POST("/api/login", logger.Logger(`Create Token for user: `)(corss.Middleware(server.LoginHandler)))

	server.router.GET(`/api/users`, logger.Logger(`Get all users: `)(corss.Middleware(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))(authorized.Authorized([]string{`admin`}, jwt.FromContext)(server.GetUsersHandler)))))
	server.router.GET(`/api/users/:id`, logger.Logger(`Get all user by id: `)(corss.Middleware(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))(authorized.Authorized([]string{`admin`}, jwt.FromContext)(server.GetUserByIdHandler)))))

	server.router.POST(`/api/users/add`, logger.Logger(`Add all user: `)(corss.Middleware(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))(authorized.Authorized([]string{`admin`}, jwt.FromContext)(server.AddNewUserHandler)))))
	server.router.POST(`/api/users/edit/:id`, logger.Logger(`Edit user by id: `)(corss.Middleware(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))(authorized.Authorized([]string{`admin`}, jwt.FromContext)(server.EditUserHandler)))))

	server.router.POST(`/api/click-state`, logger.Logger(`Set click: `)(corss.Middleware(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))(authorized.Authorized([]string{`admin`, `user`}, jwt.FromContext)(server.SetStateAndDateHandler)))))
	server.router.POST(`/api/userstate`, logger.Logger(`Get user state: `)(corss.Middleware(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))(authorized.Authorized([]string{`admin`, `user`}, jwt.FromContext)(server.GetUserStatsHandler)))))

	server.router.GET(`/api/users-states`, logger.Logger(`Get users states: `)(corss.Middleware(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))(authorized.Authorized([]string{`admin`, `user`}, jwt.FromContext)(server.GetUsersStatsHandler)))))
	settings.AppSettings = settings.ReadSettings("./settings.json")
	port := fmt.Sprintf(":%d", settings.AppSettings.AppParams.PortRun)
	log.Println(http.ListenAndServe(port, server))
}

func test()  {
	fmt.Println(time.Now().Format(time.RFC3339))
}