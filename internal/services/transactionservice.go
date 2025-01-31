package services

import (
	"errors"
	"github.com/nyae44/GoLibrary/internal/models"
	"github.com/nyae44/GoLibrary/internal/repositories"
)

type transactionService struct {
	transactionRepo repositories.TransactionRepository
	bookRepo        repositories.BookRepository
	memberRepo      repositories.MemberRepository
}

func NewTransactionService(transactionRepo repositories.TransactionRepository, bookRepo repositories.BookRepository, memberRepo repositories.MemberRepository) *transactionService {
	return &transactionService{transactionRepo: transactionRepo, bookRepo: bookRepo, memberRepo: memberRepo}
}

func (s *transactionService) BorrowBook(transaction *models.Transaction) (*models.Transaction, error) {
	// Check if book exists and is available
	book, err := s.bookRepo.FindById(transaction.BookID)
	if err != nil {
		return nil, errors.New("book not found")
	}
	if book.Status != "Available" {
		return nil, errors.New("book is not available")
	}

	// Check if member exists
	member, err := s.memberRepo.GetById(transaction.BookID)
	if err != nil {
		return nil, errors.New("member not found")
	}

	// Check if member has sufficient balance to pay the fee
	const rentalFee = 100
	if member.Balance < rentalFee {
		return nil, errors.New("not enough balance to rent a book")
	}
	// Deduct fee from member balance
	member.Balance -= rentalFee
	if _, err := s.memberRepo.Update(member); err != nil {
		return nil, err
	}
	// Update book status to not available
	book.Status = "Not available"
	if _, err := s.bookRepo.Update(book); err != nil {
		return nil, err
	}
	// Set the fee and create the transaction
	transaction.FeesCharged = rentalFee
	transaction.Type = "Borrow"
	return s.transactionRepo.Create(transaction)
}

func (s *transactionService) ReturnBook(transaction *models.Transaction) error {
	// Find the transaction
	transaction, err := s.transactionRepo.FindById(transaction.ID)
	if err != nil {
		return errors.New("transaction not found")
	}

	//Check if a transaction type is a borrow transaction
	if transaction.Type != "borrow" {
		return errors.New("invalid transaction type")
	}

	// Find the book
	book, err := s.bookRepo.FindById(transaction.BookID)
	if err != nil {
		return errors.New("book not found")
	}
	// Update book status to available
	book.Status = "Available"

	if _, err := s.bookRepo.Update(book); err != nil {
		return err
	}

	// Create a return transaction no fee for returning
	returnTransaction := &models.Transaction{
		MemberID:    transaction.MemberID,
		BookID:      transaction.BookID,
		Type:        "Return",
		FeesCharged: 0,
	}
	_, err = s.transactionRepo.Create(returnTransaction)
	return err
}
