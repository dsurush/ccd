package models

import (
	"time"
)

type TimeInterval struct {
	From int64 `json:"from"`
	To int64 `json:"to"`
}

func GetUnixTimeStartOfDay() int64 {
	t := time.Now()
	rounded := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
//	fmt.Println(t.Unix())
	return rounded.Unix()
}