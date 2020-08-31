package services

import (
	"ccs/models"
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/pgxpool"
	"golang.org/x/crypto/bcrypt"
	"log"
	"strconv"
	"time"
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

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (receiver *UserSvc) GetUserById(id string) (User models.UserDTO, err error) {
	conn, err := receiver.pool.Acquire(context.Background())
	if err != nil {
		log.Printf("can't get connection %e", err)
		return
	}
	defer conn.Release()

	err = receiver.pool.QueryRow(context.Background(), getUserByIdDML, id).Scan(
		&User.Id,
		&User.Name,
		&User.Surname,
		&User.LastName,
		&User.Login,
		//&ignore,
		&User.Phone,
		&User.Role,
		&User.Status)
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
			&User.Status)
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

func (receiver *UserSvc) AddNewUser(User models.SaveUser) (err error){
	User.Password, err = HashPassword(User.Password)
	if err != nil {
		fmt.Println("can't do your pass to hash")
		return err
	}
	conn, err := receiver.pool.Acquire(context.Background())
	if err != nil {
		log.Printf("can't get connection %e", err)
		return err
	}
	defer conn.Release()
	fmt.Println("User = ", User)
	_, err = conn.Exec(context.Background(), userSaveDML, User.Name, User.Surname, User.LastName,
		User.Login, User.Password, User.Phone)
	if err != nil {
		log.Print("can't add to db err is = ", err)
		return err
	}
	return nil
}

func (receiver *UserSvc) EditUser(User models.SaveUser, id string) (err error){
	User.Password, err = HashPassword(User.Password)
	if err != nil {
		fmt.Println("can't do your pass to hash")
		return err
	}
	conn, err := receiver.pool.Acquire(context.Background())
	if err != nil {
		log.Printf("can't get connection %e", err)
		return err
	}
	defer conn.Release()
	fmt.Println("User = ", User)
	_, err = conn.Exec(context.Background(), editUserDML, User.Name, User.Surname, User.LastName,
		User.Login, User.Password, User.Phone, id)
	if err != nil {
		log.Print("can't edit to db err is = ", err)
		return err
	}
	return nil
}

func (receiver *UserSvc) SetStateAndDate(State models.StatesDTO, id string) (err error) {
	conn, err := receiver.pool.Acquire(context.Background())
	if err != nil {
		log.Fatalf("can't get connection %e", err)
		return err
	}
	defer conn.Release()
	fmt.Println("ID = ", id)
	atoi, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("can't conver to Int")
		return
	}
	//	fmt.Println("Unix time = ", time.Now().Unix())
	_, err = conn.Exec(context.Background(), setStateAndTimeDML, int64(atoi), State.Time, State.Status, time.Now().Unix(), time.Now())
	if err != nil {
		log.Print("can't add to db err is = ", err)
		return err
	}
	_, err = conn.Exec(context.Background(), editUserStateDML, State.Status, int64(atoi))
	if err != nil {
		log.Print("can't add edit User StateDML = ", err)
		return err
	}
	return nil
}

func (receiver *UserSvc) GetUserStats(id string, from int64) (States []models.State, err error) {
	conn, err := receiver.pool.Acquire(context.Background())
	if err != nil {
		log.Printf("can't get connection %e", err)
		return
	}
	defer conn.Release()
	rows, err := conn.Query(context.Background(), getUserStatsDML, id, from)
	if err != nil {
		fmt.Printf("can't read user rows %e", err)
		return
	}
	defer rows.Close()

	for rows.Next(){
		State := models.State{}
		err := rows.Scan(
			&State.ID,
			&State.UserId,
			&State.WorkTime,
			&State.Status,
			&State.UnixDate,
			&State.TimeDate)
		if err != nil {
			fmt.Println("can't scan err is = ", err)
		}
		States = append(States, State)
	}
	if rows.Err() != nil {
		log.Printf("rows err %s", err)
		return nil, rows.Err()
	}
	return
}