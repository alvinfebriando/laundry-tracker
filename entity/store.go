package entity

import "time"

type Store struct {
	Id        uint
	Name      string
	Address   string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
