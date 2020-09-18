package app

import (
	"ccs/middleware/authorized"
	"ccs/middleware/corss"
	"ccs/middleware/jwt"
	"ccs/middleware/logger"
	"ccs/settings"
	"ccs/token"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"reflect"
	"time"

	//"time"
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
	server.router.GET(`/api/users/:id/info`, logger.Logger(`Get user state by id: `)(corss.Middleware(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))(authorized.Authorized([]string{`admin`, `user`}, jwt.FromContext)(server.GetUserStatsForAdminHandler)))))

	server.router.POST(`/api/settings/change-password`, logger.Logger(`Change pass: `)(corss.Middleware(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))(authorized.Authorized([]string{`admin`, `user`}, jwt.FromContext)(server.SetNewPassHandler)))))
	settings.AppSettings = settings.ReadSettings("./settings.json")
	port := fmt.Sprintf(":%d", settings.AppSettings.AppParams.PortRun)
	log.Println(http.ListenAndServe(port, server))
}

func test()  {

	err := bcrypt.CompareHashAndPassword([]byte(`$2a$14$iFltmkBEzTcuNVRWAPTJ2.gu7Y3O77FgADPrmCWkmtnnaa1MMkyta`), []byte(`surush`))
	if err != nil {
	//	err = ErrInvalidPasswordOrLogin
		fmt.Println(err)
	} else {
		fmt.Println("I am fine")
	}

//	fmt.Println()

//	time := fmt.Sprintf("%d:%d", time.Now().Hour(), time.Now().Minute())
//	fmt.Println(time)
	//
	t := time.Now()
	rounded := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	fmt.Println(t)
	fmt.Println(rounded)

}