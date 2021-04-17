package repository

import (
	"github.com/XXena/todo-app"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

// https://golangs.org/oop

func (r *AuthPostgres) CreateUser(user todo.User) (int, error) {
	return 0, nil
}
