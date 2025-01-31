package repositories

import (
	"github.com/nyae44/GoLibrary/internal/models"
	"gorm.io/gorm"
)

type TransactionRepository interface {
	Create(transaction *models.Transaction) (*models.Transaction, error)
	Delete(id int) error
	FindById(id uint) (*models.Transaction, error)
	FindAll() ([]*models.Transaction, error)
	Update(transaction *models.Transaction) (*models.Transaction, error)
	FindByMember(memberId int) ([]*models.Transaction, error)
	FindByBookId(bookId int) ([]*models.Transaction, error)
}

type transactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &transactionRepository{db: db}
}

func (r *transactionRepository) Create(transaction *models.Transaction) (*models.Transaction, error) {
	result := r.db.Create(transaction)
	if result.Error != nil {
		return nil, result.Error
	}
	return transaction, nil
}

func (r *transactionRepository) Delete(id int) error {
	transaction := &models.Transaction{}
	result := r.db.Delete(transaction, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *transactionRepository) FindById(id uint) (*models.Transaction, error) {
	transaction := &models.Transaction{}
	result := r.db.First(transaction, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return transaction, nil
}
func (r *transactionRepository) FindAll() ([]*models.Transaction, error) {
	transactions := make([]*models.Transaction, 0)
	result := r.db.Find(&transactions)
	if result.Error != nil {
		return nil, result.Error
	}
	return transactions, nil
}

func (r *transactionRepository) Update(transaction *models.Transaction) (*models.Transaction, error) {
	result := r.db.Save(transaction)
	if result.Error != nil {
		return nil, result.Error
	}
	return transaction, nil
}
func (r *transactionRepository) FindByBookId(bookId int) ([]*models.Transaction, error) {
	transactions := make([]*models.Transaction, 0)
	result := r.db.Where("book_id = ?", bookId).Find(&transactions)
	if result.Error != nil {
		return nil, result.Error
	}
	return transactions, nil
}
func (r *transactionRepository) FindByMember(memberId int) ([]*models.Transaction, error) {
	transactions := make([]*models.Transaction, 0)
	result := r.db.Where("member_id = ?", memberId).Find(&transactions)
	if result.Error != nil {
		return nil, result.Error
	}
	return transactions, nil
}
