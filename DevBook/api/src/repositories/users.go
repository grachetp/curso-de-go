package repositories

import (
	"api/src/models"
	"database/sql"
)

type Users struct {
	db *sql.DB
}

func NewUsersRepository(db *sql.DB) *Users {
	return &Users{db}
}

// func (inteface struct) nameFunction(params) (return params)
func (u Users) Create(user models.User) (uint64, error) {
	return 0, nil
}
