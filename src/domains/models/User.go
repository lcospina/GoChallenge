package models

import (
	"time"
)

type User struct {
	ID          int
	FirstName   string
	LastName    string
	Email       string
	NumberPhone string
	RoleID      int
	Role        Role
	CreatedAt   time.Time
	DeleteAt    time.Time
	UpdateAt    time.Time
}
