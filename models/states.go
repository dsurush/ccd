package models

import "time"

type StatesDTO struct {
	Time int64 `json:"time"`
	Status bool `json:"status"`
}

type State struct {
	ID int64 `json:"id"`
	UserId int64 `json:"user_id"`
	WorkTime int64 `json:"work_time"`
	Status bool `json:"status"`
	UnixDate int64 `json:"unix_date"`
	TimeDate time.Time `json:"time_date"`
}