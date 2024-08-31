package model

import "time"

type Task struct {
	ID          int `gorm:"primary_key"`
	Name        string
	Description string
	Priority    int
	CreatedAt   time.Time
	Status      bool
}
