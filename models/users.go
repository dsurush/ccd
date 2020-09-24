package models

type User struct {
	Id         int64  `json:"id"`
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	LastName   string `json:"last_name"`
	Login      string `json:"login"`
	Password   string `json:"password"`
	Phone      string `json:"phone"`
	Role       string `json:"role"`
	Status     bool   `json:"status"`
	Position   string `json:"position"`
	StatusLine bool   `json:"status_line"`
}

type UserDTO struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	LastName string `json:"last_name"`
	Login    string `json:"login"`
	//	Password string `json:"password"`
	Phone      string `json:"phone"`
	Role       string `json:"role"`
	Status     bool   `json:"status"`
	Position   string `json:"position"`
	StatusLine bool   `json:"status_line"`
}

type UserWithWorkTimeDTO struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	LastName string `json:"last_name"`
	Login    string `json:"login"`
	//	Password string `json:"password"`
	Phone      string `json:"phone"`
	Role       string `json:"role"`
	Status     bool   `json:"status"`
	Position   string `json:"position"`
	StatusLine bool   `json:"status_line"`
	Worked     int64  `json:"worked"`
	Rest       int64  `json:"rest"`
}

type SaveUser struct {
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	LastName string `json:"last_name"`
	Login    string `json:"login"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
	Position string `json:"position"`
}

type ChangePassword struct {
	Password    string `json:"password"`
	NewPassword string `json:"new_password"`
}
