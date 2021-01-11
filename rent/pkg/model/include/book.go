package include

import "time"

type Book struct {
	ID          int64
	Status      int8
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
