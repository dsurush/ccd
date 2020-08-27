package token

import (
	"ccs/models"
	"context"
	"errors"
	"fmt"
	"github.com/dsurush/jwt/pkg/jwt"
	"github.com/jackc/pgx/pgxpool"
	"golang.org/x/crypto/bcrypt"
	"log"
	"time"
)

type TokenSvc struct {
	secret []byte
	pool   *pgxpool.Pool
}

func NewTokenSvc(secret []byte, pool *pgxpool.Pool) *TokenSvc {
	return &TokenSvc{secret: secret, pool: pool}
}

type Payload struct {
	Id    int64  `json:"id"`
	Exp   int64  `json:"exp"`
	Login string `json:"login"`
	Role  string `json:"role"`
}

type RequestDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type ResponseDTO struct {
	Token string `json:"token"`
	Role string `json:"role"`
	State bool `json:"state"`
	Name string `json:"name"`
	Surname string `json:"surname"`
}

//var ErrInvalidLogin = errors.New("invalid login or password")
var ErrInvalidPasswordOrLogin = errors.New("invalid password")

func (receiver *TokenSvc) FindUserForPassCheck(login string) (User models.User, err error) {
	conn, err := receiver.pool.Acquire(context.Background())
	if err != nil {
		log.Printf("can't get connection %e", err)
		return
	}
	defer conn.Release()

	err = receiver.pool.QueryRow(context.Background(), `Select *from users where login = ($1)`, login).Scan(
		&User.Id,
		&User.Name,
		&User.Surname,
		&User.LastName,
		&User.Login,
		&User.Password,
		&User.Phone,
		&User.Role,
		&User.Status)
	if err != nil {
		fmt.Printf("Can't scan %e", err)
	}
	return
}
func (receiver *TokenSvc) Generate(context context.Context, request *RequestDTO) (response ResponseDTO, err error) {
	//user, err := models.FindUserByLogin(request.Username)
	user, err := receiver.FindUserForPassCheck(request.Username)
	if err != nil {
		err = ErrInvalidPasswordOrLogin
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		err = ErrInvalidPasswordOrLogin
		return
	}

	response.Token, err = jwt.Encode(Payload{
		Id:  user.Id,
		Exp: time.Now().Add(time.Hour * 10).Unix(),
		///Exp:   time.Now().Add(time.Second * 10).Unix(),
		Login: user.Login,
		Role:  user.Role,
	}, receiver.secret)
	return
}