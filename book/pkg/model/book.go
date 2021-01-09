package model

import "time"

type Book struct {
	ID          int64 `gorm:"primary_key"`
	Status      int8  `gorm:"default:1"`
	RentStatus  int8  `gorm:"default:1"`
	ImgUrl      string
	ISBN        string
	BookName    string
	Publisher   string
	Category    string
	Author      string
	Description string
	CreateTime  time.Time
	UpdateTime  time.Time
}
