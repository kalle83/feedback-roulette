package model

type Question struct {
	Id   int64  `json:"id" gorm:"primary_key"`
	Text string `json:"text"`
}
