package main

import (
	"ccs/cmd/ccs/app"
	"ccs/pkg/core/services"
	"ccs/token"
	"context"
	"flag"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"net/http"

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

	router.GlobalOPTIONS = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Header.Get("Access-Control-Request-Method") != "" {
			// Set CORS headers
			w.Header().Set("Content-Type", "application/json, text/html")
			//			w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, OPTIONS")
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, RefreshToken")
			w.Header().Set("Accept", "*/*")
			w.WriteHeader(http.StatusOK)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	})
	password, err := HashPassword("shah")
	fmt.Println("Im pass = ", password)
	server := app.NewMainServer(router, pool, svc, tokenSvc)

	server.Start()
}
