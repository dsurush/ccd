package main

import (
	"ccs/cmd/ccs/app"
	"ccs/pkg/core/services"
	"ccs/token"
	"context"
	"flag"
	"fmt"
	"golang.org/x/crypto/bcrypt"

	//	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/jackc/pgx/pgxpool"
	"github.com/julienschmidt/httprouter"
	"log"
)

var (
	dsn = flag.String("dsn", "postgres://localadmin:123456789@localhost:5432/ccd", "Postgres DSN")
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func main() {
	flag.Parse()
	router := httprouter.New()

	pool, err := pgxpool.Connect(context.Background(), *dsn)
	if err != nil {
		log.Printf("Owibka - %e", err)
		log.Fatal("BAD")
	} else {
		fmt.Println("Chix and Pux")
	}

	svc := services.NewUserSvc(pool)
	tokenSvc := token.NewTokenSvc([]byte(`surush`), pool)
	server := app.NewMainServer(router, pool, svc, tokenSvc)

	server.Start()
}
