package app

import (
	"ccs/pkg/core/services"
	"ccs/token"
	"fmt"
	"github.com/jackc/pgx/pgxpool"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type MainServer struct {
	router *httprouter.Router
	pool *pgxpool.Pool
	svc *services.UserSvc
	tokenSvc *token.TokenSvc
}

func NewMainServer(router *httprouter.Router, pool *pgxpool.Pool, svc *services.UserSvc, tokenSvc *token.TokenSvc) *MainServer {
	return &MainServer{router: router, pool: pool, svc: svc, tokenSvc: tokenSvc}
}

func (server *MainServer) Start() {
	err := server.svc.DbInit()
	if err != nil {
		fmt.Println("no build")
		//TODO: CLOSE ME
	}

	server.InitRoutes()

}

func (server *MainServer) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	server.router.ServeHTTP(writer, request)
}
