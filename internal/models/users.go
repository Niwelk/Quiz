package models

import (
	"time"
)

type User struct {
	Id        uint32
	Login     string
	Password  string
	Email     string
	CreatedAt time.Time
}
