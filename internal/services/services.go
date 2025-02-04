package services

import "github.com/nyae44/GoLibrary/internal/models"

type BookService interface {
	CreateBook(book *models.Book) (*models.Book, error)
	GetBookByID(id string) (*models.Book, error)
	UpdateBook(book *models.Book) (*models.Book, error)
	DeleteBook(id uint) error
	ListBooks() ([]models.Book, error)
}
type MemberService interface {
	CreateMember(member *models.Member) error
	GetMemberByID(id string) (*models.Member, error)
	UpdateMember(member *models.Member) error
	DeleteMember(id uint) error
	ListMembers() ([]models.Member, error)
}

type TransactionService interface {
	BorrowBook(transaction *models.Transaction) (*models.Transaction, error)
	ReturnBook(transaction *models.Transaction) error
	GetTransactionByID(id uint) (*models.Transaction, error)
	ListTransactions() ([]models.Transaction, error)
	GetTransactionsByMember(memberID uint) ([]models.Transaction, error)
	GetTransactionsByBook(bookID uint) ([]models.Transaction, error)
}

type UserService interface {
	CreateUser(user *models.User) error
	GetUserByID(id uint) (*models.User, error)
	UpdateUser(user *models.User) error
	DeleteUser(id uint) error
	ListUsers() ([]models.User, error)
	GetUserByName(name string) (*models.User, error)
}
