package models

import "time"

type Report struct {
	Name       string    `json:"name"`
	Surname    string    `json:"surname"`
	LoginDate  []string  `json:"login_date"`
	LogoutDate []string  `json:"logout_date"`
	Sum        int64     `json:"sum"`
	Time       time.Time `json:"time"`
}