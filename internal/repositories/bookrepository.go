package repositories

import (
	"github.com/nyae44/GoLibrary/internal/models"
	"gorm.io/gorm"
)

type BookRepository interface {
	Create(book *models.Book) error
	FindById(id uint) (*models.Book, error)
	Update(book *models.Book) error
	DeleteById(id uint) error
}

type bookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) BookRepository {
	return &bookRepository{db: db}
}

func (r *bookRepository) Create(book *models.Book) error {
	return r.db.Create(book).Error
}
func (r *bookRepository) FindById(id uint) (*models.Book, error) {
	var book models.Book
	err := r.db.First(&book, id).Error
	return &book, err
}

func (r *bookRepository) Update(book *models.Book) error {
	return r.db.Save(book).Error
}

func (r *bookRepository) DeleteById(id uint) error {
	return r.db.Delete(&models.Book{}, id).Error
}
