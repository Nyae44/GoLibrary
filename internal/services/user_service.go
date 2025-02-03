package services

import (
	"github.com/nyae44/GoLibrary/internal/models"
	"github.com/nyae44/GoLibrary/internal/repositories"
)

type userService struct {
	userRepo repositories.UserRepository
}

func NewUserService(userRepo repositories.UserRepository) *userService {
	return &userService{userRepo: userRepo}
}

func (s *userService) CreateUser(user *models.User) error {
	return s.userRepo.Create(user)
}
func (s *userService) UpdateUser(user *models.User) error {
	return s.userRepo.Update(user)
}
func (s *userService) DeleteUser(user *models.User) error {
	return s.userRepo.Delete(user.ID)
}

func (s *userService) GetUser(id uint) (*models.User, error) {
	return s.userRepo.FindById(id)
}
func (s *userService) GetUserByName(name string) (*models.User, error) {
	return s.userRepo.FindByName(name)
}
func (s *userService) GetUserByEmail(email string) (*models.User, error) {
	return s.userRepo.FindByEmail(email)
}
