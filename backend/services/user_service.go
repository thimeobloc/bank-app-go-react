package services

import (
	"banque-app/backend/models"
	"banque-app/backend/repositories"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	Repo *repositories.UserRepository
}

func (s *UserService) RegisterUser(name, email, password string) error {
	existingUser, _ := s.Repo.FindByEmail(email)
	if existingUser != nil {
		return fmt.Errorf("email already in use")
	}
	hashed, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	user := models.User{Name: name, Email: email, PasswordHash: string(hashed)}
	return s.Repo.Create(user)
}

func (s *UserService) LoginUser(email, password string) (*models.User, error) {
	existingUser, err := s.Repo.FindByEmail(email)
	if err != nil {
		return nil, fmt.Errorf("erreur serveur")
	}

	if existingUser == nil {
		return nil, fmt.Errorf("user not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(existingUser.PasswordHash), []byte(password))
	if err != nil {
		return nil, fmt.Errorf("password incorrect")
	}

	return existingUser, nil
}

func (s *UserService) DeleteUser(email string) error {
	user, err := s.Repo.FindByEmail(email)
	if err != nil || user == nil {
		return fmt.Errorf("user not found")
	}
	return s.Repo.Delete(email)
}
