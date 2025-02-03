package handlers

import (
	"encoding/json"
	"github.com/nyae44/GoLibrary/internal/models"
	"github.com/nyae44/GoLibrary/internal/services"
	"net/http"
	"strconv"
)

type TransactionHandler struct {
	transactionService services.TransactionService
}

func NewTransactionHandler(transactionService services.TransactionService) *TransactionHandler {
	return &TransactionHandler{transactionService: transactionService}
}

// HandleBorrowBook borrow book
func (h *TransactionHandler) HandleBorrowBook(w http.ResponseWriter, r *http.Request) {
	var transaction models.Transaction
	if err := json.NewDecoder(r.Body).Decode(&transaction); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	createdTransaction, err := h.transactionService.BorrowBook(&transaction)
	if err != nil {
		if err.Error() == "insufficient balance to borrow book" {
			http.Error(w, err.Error(), http.StatusPaymentRequired)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(createdTransaction)
	if err != nil {
		return
	}
}

// HandleReturnBook handles the request to return a book
func (h *TransactionHandler) HandleReturnBook(w http.ResponseWriter, r *http.Request) {
	var request struct {
		TransactionID uint `json:"transaction_id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	transaction, err := h.transactionService.GetTransactionByID(request.TransactionID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	if err := json.NewEncoder(w).Encode(transaction); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

//HandleTransactionByID handles request to get a transaction by ID

func (h *TransactionHandler) HandleTransactionByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(w, "Invalid transaction ID", http.StatusBadRequest)
		return
	}

	transaction, err := h.transactionService.GetTransactionByID(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(transaction)
}

//HandleListTransactions handles request to list all Transactions

func (h *TransactionHandler) HandleGetTransactions(w http.ResponseWriter, r *http.Request) {
	transactions, err := h.transactionService.ListTransactions()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(transactions)
}

// HandleGetTransactionByMember handles request to get transactions by member ID
func (h *TransactionHandler) HandleGetTransactionByMember(w http.ResponseWriter, r *http.Request) {
	memberIDStr := r.URL.Query().Get("member_id")
	memberID, err := strconv.Atoi(memberIDStr)
	if err != nil {
		http.Error(w, "Invalid member ID", http.StatusBadRequest)
		return
	}
	transactions, err := h.transactionService.GetTransactionsByMember(uint(memberID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(transactions)
}

//HandleGetTransactionByBook handles request to get transactions by book borrowed

func (h *TransactionHandler) HandleGetTransactionByBook(w http.ResponseWriter, r *http.Request) {
	bookIDStr := r.URL.Query().Get("book_id")
	bookID, err := strconv.Atoi(bookIDStr)
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}
	transactions, err := h.transactionService.GetTransactionsByBook(uint(bookID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(transactions)
}
