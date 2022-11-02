package service

import (
	"github.com/devstackq/smtp-mailer/internal/models"
	"github.com/devstackq/smtp-mailer/internal/repository"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUser(userRepo repository.UserRepository) *UserService {
	return &UserService{
		repo: userRepo,
	}
}
func (u UserService) Create(user models.User) error {
	return u.repo.Create(user)
}
func (u UserService) GetListUser() ([]models.User, error) {
	return u.repo.GetListUser()
}

func (u UserService) GetUserByID(id string) (*models.User, error) {
	return u.repo.GetUserById(id)
}
