package app

import (
	"ccs/middleware/authorized"
	"ccs/middleware/corss"
	"ccs/middleware/jwt"
	"ccs/middleware/logger"
	"ccs/pkg/core/services"
	"ccs/settings"
	"ccs/token"
	"fmt"
	"log"
	"net/http"
	"reflect"
	//"time"
)

func (server *MainServer) InitRoutes() {
	fmt.Println("Init routes")
	test(server)
	server.router.POST("/api/login", logger.Logger(`Create Token for user: `)(corss.Middleware(server.LoginHandler)))
	server.router.GET(`/api/users`, server.GetUsersHandler)
	//server.router.GET(`/api/users`, logger.Logger(`Get all users: `)(corss.Middleware(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))(authorized.Authorized([]string{`admin`}, jwt.FromContext)(server.GetUsersHandler)))))
	//  Список user - ов с рабочим временем
	server.router.GET(`/api/user-worktime`, logger.Logger(`Get user by id: `)(corss.Middleware(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))(authorized.Authorized([]string{`admin`}, jwt.FromContext)(server.GetUsersWithWorkTimeHandler)))))

	server.router.GET(`/api/users/:id`, logger.Logger(`Get user by id: `)(corss.Middleware(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))(authorized.Authorized([]string{`admin`}, jwt.FromContext)(server.GetUserByIdHandler)))))

	server.router.POST(`/api/users/add`, logger.Logger(`Add all user: `)(corss.Middleware(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))(authorized.Authorized([]string{`admin`}, jwt.FromContext)(server.AddNewUserHandler)))))
	server.router.POST(`/api/users/edit/:id`, logger.Logger(`Edit user by id: `)(corss.Middleware(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))(authorized.Authorized([]string{`admin`}, jwt.FromContext)(server.EditUserHandler)))))

	server.router.POST(`/api/click-state`, logger.Logger(`Set click: `)(corss.Middleware(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))(authorized.Authorized([]string{`admin`, `user`}, jwt.FromContext)(server.SetStateAndDateHandler)))))
	server.router.POST(`/api/userstate`, logger.Logger(`Get user state: `)(corss.Middleware(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))(authorized.Authorized([]string{`admin`, `user`}, jwt.FromContext)(server.GetUserStatsHandler)))))

	server.router.GET(`/api/users-states`, logger.Logger(`Get users states: `)(corss.Middleware(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))(authorized.Authorized([]string{`admin`, `user`}, jwt.FromContext)(server.GetUsersStatsHandler)))))
	server.router.GET(`/api/users/:id/info`, logger.Logger(`Get user state by id: `)(corss.Middleware(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))(authorized.Authorized([]string{`admin`, `user`}, jwt.FromContext)(server.GetUserStatsForAdminHandler)))))
	server.router.GET(`/api/report`, logger.Logger(`Get report: `)(corss.Middleware(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))(authorized.Authorized([]string{`admin`, `user`}, jwt.FromContext)(server.ReportHandler)))))
	server.router.POST(`/api/settings/change-password`, logger.Logger(`Change pass: `)(corss.Middleware(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))(authorized.Authorized([]string{`admin`, `user`}, jwt.FromContext)(server.SetNewPassHandler)))))

	server.router.POST(`/api/status-confirm`, logger.Logger(`Status confirm: `)(corss.Middleware(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))(authorized.Authorized([]string{`admin`, `user`}, jwt.FromContext)(server.StatusConfirmHandler)))))

	server.router.POST(`/api/exit`, logger.Logger(`Exit click: `)(corss.Middleware(jwt.JWT(reflect.TypeOf((*token.Payload)(nil)).Elem(), []byte(`surush`))(authorized.Authorized([]string{`admin`, `user`}, jwt.FromContext)(server.ExitClickHandler)))))

	settings.AppSettings = settings.ReadSettings("./settings.json")
	port := fmt.Sprintf(":%d", settings.AppSettings.AppParams.PortRun)
	log.Println(http.ListenAndServe(port, server))
}

func test(server *MainServer)  {

	//err := server.svc.FixTimeLogin("1")
	//if err != nil {
	//	fmt.Println("pizda")
	//} else {
	//	fmt.Println("xuynya")
	//}
	//day := models.GetUnixTimeStartOfDay(time.Now())
	//fmt.Println(day)
	//err := server.svc.SetVisitTime(`2`)
	//if err != nil {
	//	fmt.Println("pizda")
	//} else {
	//	fmt.Println("Uspeshno")
	//}
	//
	//sprintf := fmt.Sprintf("%s", time.Now())
	//fmt.Println(sprintf[0:10], "|", "+")
	//me, err := server.svc.TestMe(sprintf[0:10])
	//if err != nil {
	//	fmt.Println("pizda")
	//} else {
	//	fmt.Println(me)
	//}
	//
	//_, err := server.svc.CheckHasFixForToday(1)
	//if err != nil {
	//	fmt.Println("Pizdec naxoy blyat")
	//	return
	//} else {
	//	fmt.Println("YESSS")
	//}
	//err := bcrypt.CompareHashAndPassword([]byte(`$2a$14$iFltmkBEzTcuNVRWAPTJ2.gu7Y3O77FgADPrmCWkmtnnaa1MMkyta`), []byte(`surush`))
	//if err != nil {
	////	err = ErrInvalidPasswordOrLogin
	//	fmt.Println(err)
	//} else {
	//	fmt.Println("I am fine")
	//}

//	fmt.Println()

//	time := fmt.Sprintf("%d:%d", time.Now().Hour(), time.Now().Minute())
//	fmt.Println(time)
	//
	//t := time.Now()
	//rounded := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	//fmt.Println(rounded.Unix())
//	fmt.Println(time.Now().Unix())
//	server.svc.UpdateToFixLoginTime(`1`)
//	report, err := server.svc.GetReport("2020-08-01", `2020-12-30`)
//	if err != nil {
//		fmt.Println("XXXX")
//	} else {
//		fmt.Println(report)
//	}
	password, err := services.HashPassword(`0089`)
	if err != nil {
		fmt.Println(`xaxa`)
	} else {
		fmt.Println(password)
	}
}