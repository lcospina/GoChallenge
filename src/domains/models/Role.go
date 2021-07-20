package models

import (
	"time"
)

type Role struct {
	ID          int
	Description string
	CreatedAt   time.Time
	DeleteAt    time.Time
	UpdateAt    time.Time
}
