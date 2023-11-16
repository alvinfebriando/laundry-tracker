package entity

import "time"

type User struct {
	Id           uint
	Name         string
	LaundryItems []LaundryItem
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    time.Time
}
