package models

import (
	"time"
)

type User struct {
	UserId    int
	FirstName string
	LastName  string
	BirthDate time.Time
	Email     string
	Password  string
	Avatar    string
	IsDeleted bool
	IsLocked  bool
}
