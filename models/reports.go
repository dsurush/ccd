package models

import "time"

type Report struct {
	Name       string    `json:"name"`
	Surname    string    `json:"surname"`
	Position   string    `json:"position"`
	LoginDate  []string  `json:"login_date"`
	LogoutDate []string  `json:"logout_date"`
	Work       int64     `json:"work"`
	Rest       int64     `json:"rest"`
	Time       time.Time `json:"time"`
}
