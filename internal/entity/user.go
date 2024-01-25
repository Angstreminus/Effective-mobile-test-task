package entity

import (
	uuid "github.com/satori/go.uuid"
)

type User struct {
	ID          uuid.UUID
	Name        string
	Surname     string
	Patronymic  string
	Gender      string
	Nationality string
	CreatedAt   string
	UpdatedAt   string
	DeletedAt   string
	Age         int
	IsDeleted   bool
}
