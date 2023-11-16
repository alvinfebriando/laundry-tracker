package entity

import "time"

type LaundryItem struct {
	Id          uint
	UserId      uint
	User        User
	Slug        string
	Photo       string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
}
