package repositories

import (
	"github.com/nyae44/GoLibrary/internal/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *models.User) error
	Delete(id uint) error
	FindById(id uint) (*models.User, error)
	FindByName(name string) (*models.User, error)
	FindByEmail(email string) (*models.User, error)
	Update(user *models.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}
func (r *userRepository) Create(user *models.User) error {
	return r.db.Create(user).Error
}
func (r *userRepository) Delete(id uint) error {
	return r.db.Delete(id).Error
}
func (r *userRepository) FindById(id uint) (*models.User, error) {
	user := &models.User{}
	if err := r.db.First(user, id).Error; err != nil {
		return nil, err
	}
	return user, nil
}
func (r *userRepository) FindByName(name string) (*models.User, error) {
	user := &models.User{}
	if err := r.db.Where("name = ?", name).First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
func (r *userRepository) FindByEmail(email string) (*models.User, error) {
	user := &models.User{}
	if err := r.db.Where("email = ?", email).First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
func (r *userRepository) Update(user *models.User) error {
	return r.db.Save(user).Error
}
