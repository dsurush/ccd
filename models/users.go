package models

type User struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	LastName string `json:"last_name"`
	Login    string `json:"login"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
	Role     string `json:"role"`
	Status   bool   `json:"status"`
}

type UserDTO struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	LastName string `json:"last_name"`
	Login    string `json:"login"`
	//	Password string `json:"password"`
	Phone  string `json:"phone"`
	Role   string `json:"role"`
	Status bool   `json:"status"`
}

type SaveUser struct {
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	LastName string `json:"last_name"`
	Login    string `json:"login"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
}
