package entity

import "time"

type Laundry struct {
	Id         uint
	UserId     uint
	User       *User
	StoreId    uint
	Store      *Store
	Date       time.Time
	PickedDate time.Time
	Price      int
	Weight     float32
	Type       int
	Details    []LaundryDetail
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  time.Time
}
