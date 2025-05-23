package user

import (
	"time"
)

type Role string

const (
	RoleAdmin Role = "Admin"
	RoleUser  Role = "User"
)

type EmailVerify struct {
	IsVerified bool
	Token      string
}

type Entity struct {
	Id   string
	Name string

	Email       string
	EmailVerify EmailVerify
	Password    string

	Role Role

	CreatedAt  time.Time
	UpdatedAt  time.Time
	ArchivedAt *time.Time
}
