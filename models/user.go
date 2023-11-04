package models

import (
	"time"
)

type User struct {
	UserId    int
	FirstName string
	LastName  string
	DateBirth time.Time
	Email     string
	Password  string
	Avatar    string
}
