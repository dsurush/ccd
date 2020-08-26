package services

import (
	"ccs/models"
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/pgxpool"
	"log"
)

type UserSvc struct {
	pool *pgxpool.Pool
}

func NewUserSvc(pool *pgxpool.Pool) *UserSvc {
	if pool == nil {
		panic(errors.New("pool can't be nil")) // <- be accurate
	}
	return &UserSvc{pool: pool}
}

func (receiver *UserSvc) DbInit() error {
	fmt.Println("services initial")
	ddls := []string{createUsersDDL}
	for _, ddl := range ddls {
		_, err := receiver.pool.Exec(context.Background(), ddl)
		if err != nil {
			log.Printf("err, %e\n", err)
			return err
		}
	}
	return nil
}

type ResponseUsers struct {
	Page      int64         `json:"page"`
	TotalPage int64         `json:"totalPage"`
	URL       string        `json:"url"`
	Users     []models.User `json:"data"`
}

func (receiver *UserSvc) GetUserById(id string) (User models.UserDTO, err error) {
	conn, err := receiver.pool.Acquire(context.Background())
	if err != nil {
		log.Printf("can't get connection %e", err)
		return
	}
	defer conn.Release()
	// for ignoring password
	//var ignore string
	//
	err = receiver.pool.QueryRow(context.Background(), getUserByIdDML, id).Scan(
		&User.Id,
		&User.Name,
		&User.Surname,
		&User.LastName,
		&User.Login,
		//&ignore,
		&User.Phone,
		&User.Role)
	if err != nil {
		fmt.Printf("Can't scan %e", err)
	}
	return
}

func (receiver *UserSvc) GetUsers() (Users []models.UserDTO, err error) {
	conn, err := receiver.pool.Acquire(context.Background())
	if err != nil {
		log.Printf("can't get connection %e", err)
		return
	}
	defer conn.Release()
	rows, err := conn.Query(context.Background(), getUsersDML)
	if err != nil {
		fmt.Printf("can't read user rows %e", err)
		return
	}
	defer rows.Close()

	for rows.Next(){
		User := models.UserDTO{}
		err := rows.Scan(
			&User.Id,
			&User.Name,
			&User.Surname,
			&User.LastName,
			&User.Login,
			&User.Phone,
			&User.Role,
		)
		if err != nil {
			fmt.Println("can't scan err is = ", err)
		}
		Users = append(Users, User)
	}
	if rows.Err() != nil {
		log.Printf("rows err %s", err)
		return nil, rows.Err()
	}
	return
}