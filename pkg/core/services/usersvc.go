package services

import (
	"ccs/models"
	"ccs/token"
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
	ddls := []string{createUsersDDL, createStatesDDL}
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
		&User.Status,
		&User.Position,
		&User.StatusLine)
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
			&User.Status,
			&User.Position,
			&User.StatusLine)
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
// User With WorkTime DTO
func (receiver *UserSvc) GetUsersWithWorkTime() (Users []models.UserWithWorkTimeDTO, err error) {
	conn, err := receiver.pool.Acquire(context.Background())
	if err != nil {
		log.Printf("can't get connection %e", err)
		return
	}
	defer conn.Release()
	StartOfDay := models.GetUnixTimeStartOfDay(time.Now())
	rows, err := conn.Query(context.Background(), getUsersWithWorkTimeDML, StartOfDay, StartOfDay)
	if err != nil {
		fmt.Printf("can't read user rows %e", err)
		return
	}
	defer rows.Close()

	for rows.Next(){
		User := models.UserWithWorkTimeDTO{}
		//string := ""
		err := rows.Scan(
			&User.Id,
			&User.Name,
			&User.Surname,
			&User.LastName,
			&User.Login,
			&User.Phone,
			&User.Role,
			&User.Status,
			&User.Position,
			&User.StatusLine,
			&User.Worked,
			&User.Rest)
		if err != nil {
			fmt.Println("can't scan err is = ", err)
//			return
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
		User.Login, User.Password, User.Phone, User.Position)
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
	if User.Password == ``{
		_, err = conn.Exec(context.Background(), editUserWithoutPassDML, User.Name, User.Surname, User.LastName,
			User.Login, User.Phone, User.Position, id)
		if err != nil {
			log.Print("can't edit to db err is = ", err)
			return err
		}
	} else {
		_, err = conn.Exec(context.Background(), editUserDML, User.Name, User.Surname, User.LastName,
			User.Login, User.Password, User.Phone, User.Position, id)
		if err != nil {
			log.Print("can't edit to db err is = ", err)
			return err
		}
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

func (receiver *UserSvc) GetUsersStats(interval models.TimeInterval) (States []models.TotalState, err error) {
	conn, err := receiver.pool.Acquire(context.Background())
	if err != nil {
		log.Printf("can't get connection %e", err)
		return
	}
	defer conn.Release()
	rows, err := conn.Query(context.Background(), geUsersStatsDML, interval.From, interval.To)
	if err != nil {
		fmt.Printf("can't read user rows %e", err)
		return
	}
	defer rows.Close()

	for rows.Next(){
		State := models.TotalState{}
		err := rows.Scan(
			&State.Name,
			&State.Surname,
			&State.WorkTime,
			&State.CountInterval)
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

func (receiver *UserSvc) GetUserStatsForAdmin(id string, interval models.TimeInterval) (States []models.State, err error) {
	conn, err := receiver.pool.Acquire(context.Background())
	if err != nil {
		log.Printf("can't get connection %e", err)
		return
	}
	defer conn.Release()
	rows, err := conn.Query(context.Background(), getUserStatsForAdminDML, id, interval.From, interval.To)
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

func (receiver *UserSvc) ChangePassword(id string, pass string, newPass string) (err error){
	conn, err := receiver.pool.Acquire(context.Background())
	if err != nil {
		log.Printf("can't get connection %e", err)
		return
	}
	defer conn.Release()
	var password string
	err = receiver.pool.QueryRow(context.Background(), getUserPassByIdDML, id).Scan(&password)
	if err != nil {
		fmt.Printf("Can't scan %e", err)
		return
	}
//TODO: HERE
	err = bcrypt.CompareHashAndPassword([]byte(password), []byte(pass))
	if err != nil {
		err = token.ErrInvalidPasswordOrLogin
		fmt.Println(err)
	}
	HashPass, err := HashPassword(newPass)
	_, err = receiver.pool.Exec(context.Background(), setUserPassByIdDML, HashPass, id)

	if err != nil {
		fmt.Printf("Can't set new pass %e", err)
		return
	}
	
	return
}

func (receiver *UserSvc) SetStatusLine(login string, statusLine bool) (err error) {
	conn, err := receiver.pool.Acquire(context.Background())
	if err != nil {
		log.Fatalf("can't get connection %e", err)
		return err
	}
	defer conn.Release()

	_, err = conn.Exec(context.Background(), editUserStatusLineDML, statusLine, login)
	if err != nil {
		log.Print("can't add to db status_line true, err is  = ", err)
		return err
	}
	return nil
}

//
func (receiver *UserSvc) SetStatusLineById(id string, statusLine bool) (err error) {
	conn, err := receiver.pool.Acquire(context.Background())
	if err != nil {
		log.Fatalf("can't get connection %e", err)
		return err
	}
	defer conn.Release()

	_, err = conn.Exec(context.Background(), editUserStatusLineByIdDML, statusLine, id)
	if err != nil {
		log.Print("can't add to db status_line true, err is  = ", err)
		return err
	}
	return nil
}

///
func (receiver *UserSvc) SetStatusById(id string, status bool) (err error) {
	conn, err := receiver.pool.Acquire(context.Background())
	if err != nil {
		log.Fatalf("can't get connection %e", err)
		return err
	}
	defer conn.Release()

	_, err = conn.Exec(context.Background(), editUserStatusByIdDML, status, id)
	if err != nil {
		log.Print("can't add to db status true, err is  = ", err)
		return err
	}
	return nil
}



func (receiver *UserSvc) ExitClick(id string, State models.StatesDTO) (err error){
	const StatusFalse = false
	err = receiver.SetStatusLineById(id, StatusFalse)
	if err != nil {
		log.Print("can't add to db status_line false, err is  = ", err)
		return err
	}
	err = receiver.SetStateAndDate(State, id)
	if err != nil {
		log.Print("can't set to db state and date, err is = ", err)
		return err
	}
	//
	conn, err := receiver.pool.Acquire(context.Background())
	if err != nil {
		log.Fatalf("can't get connection %e", err)
		return err
	}
	defer conn.Release()
	_, err = conn.Exec(context.Background(), editUserStateDML, false, id)
	if err != nil {
		log.Print("can't add edit User StateDML = ", err)
		return err
	}
	//
	return
}

func (receiver *UserSvc) FixTimeLogin(id string) (err error) {
	conn, err := receiver.pool.Acquire(context.Background())
	if err != nil {
		log.Printf("can't get connection %e", err)
		return
	}
	defer conn.Release()
	hour := fmt.Sprintf("%s:%s", strconv.Itoa(time.Now().Hour()), strconv.Itoa(time.Now().Minute()))
	var hours []string
	hours = append(hours, hour)
	_, err = conn.Exec(context.Background(), FixLoginTime, id, time.Now().Unix(), hours, time.Now())
	if err != nil {
		fmt.Printf(" Cant Get %e", err)
		return
	}
	return
}

func (receiver *UserSvc) TestMe(time string) (Reports []models.Report, err error) {
	conn, err := receiver.pool.Acquire(context.Background())
	if err != nil {
		log.Printf("can't get connection %e", err)
		return
	}
	defer conn.Release()
	rows, err := conn.Query(context.Background(), `SELECT us.name, us.surname, lt.login_date, lt.logout_date, sum(swd.work_time), swd.time_date FROM public.states_with_date as swd,
		login_times as lt, users as us
		where lt.user_id = swd.user_id and lt.time_date = swd.time_date and us.id = lt.user_id and lt.time_date < ($1)
		group by swd.user_id, swd.time_date, lt.login_date, lt.logout_date, us.name, us.surname`, time)
	if err != nil {
		fmt.Printf("can't read user rows %e", err)
		return
	}
	defer rows.Close()

	for rows.Next(){
		Report := models.Report{}
		err := rows.Scan(
			&Report.Name,
			&Report.Surname,
			&Report.LoginDate,
			&Report.LogoutDate,
			&Report.Sum,
			&Report.Time)
		if err != nil {
			fmt.Println("can't scan err is = ", err)
		}
		Reports = append(Reports, Report)
	}
	if rows.Err() != nil {
		log.Printf("rows err %s", err)
		return nil, rows.Err()
	}
	return
}
func (receiver *UserSvc) CheckHasFixForToday(id int64) (newId int64, err error){
	conn, err := receiver.pool.Acquire(context.Background())
	if err != nil {
		log.Printf("can't get connection %e", err)
		return
	}
	defer conn.Release()
	sprintf := fmt.Sprintf("%s", time.Now())
	TimeDate := sprintf[0:10]
	fmt.Println(sprintf[0:10])
	var idNew int64
	_ = conn.QueryRow(context.Background(), `Select id from login_times where 
user_id = ($1) and time_date = ($2)`, id, TimeDate).Scan(&idNew)
	fmt.Println("I am newID = (%d)", idNew)
	return
}

func (receiver *UserSvc) SetLoginTime() {

}