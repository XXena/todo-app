package service

import (
	"crypto/sha1"
	"fmt"
	"github.com/XXena/todo-app"
	"github.com/XXena/todo-app/package/repository"
)

const salt = "sdlfghdfjhgpsfdhgpadifbvpia"

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

// https://golangs.org/oop

func (s *AuthService) CreateUser(user todo.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)

	return s.repo.CreateUser(user)
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
