package services

import (
	"github.com/nyae44/GoLibrary/internal/models"
	"github.com/nyae44/GoLibrary/internal/repositories"
)

type bookService struct {
	bookRepo repositories.BookRepository
}

func NewBookService(bookRepo repositories.BookRepository) *bookService {
	return &bookService{bookRepo: bookRepo}
}
func (s *bookService) CreateBook(book *models.Book) (error, error) {
	return s.bookRepo.Create(book), nil
}

func (s *bookService) GetBookByID(id uint) (*models.Book, error) {
	return s.bookRepo.FindById(id)
}

func (s *bookService) UpdateBook(book *models.Book) (*models.Book, error) {
	err, _ := s.bookRepo.Update(book)
	if err != nil {
		return nil, err
	}
	return book, nil
}

func (s *bookService) DeleteBook(id uint) error {
	return s.bookRepo.DeleteById(id)
}

func (s *bookService) ListBooks() ([]*models.Book, error) {
	return s.bookRepo.FindAll()
}
