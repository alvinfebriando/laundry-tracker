package entity

import "time"

type LaundryDetail struct {
	Id            uint
	LaundryId     uint
	Laundry       Laundry
	LaundryItemId uint
	LaundryItem   LaundryItem
	Status        string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     time.Time
}
