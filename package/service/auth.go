package service

import (
	"github.com/XXena/todo-app"
	"github.com/XXena/todo-app/package/repository"
)

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user todo.User) (int, error) {
	return s.repo.CreateUser(user)
}
