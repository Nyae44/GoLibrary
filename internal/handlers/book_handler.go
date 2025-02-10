package handlers

import (
	"encoding/json"
	"github.com/nyae44/GoLibrary/internal/models"
	"github.com/nyae44/GoLibrary/internal/services"
	"net/http"
	"strconv"
)

type BookHandler struct {
	booksService services.BookService
}

func NewBookHandler(booksService services.BookService) *BookHandler {
	return &BookHandler{booksService: booksService}
}

//HandleCreateBook a request to handle the creation of a book

func (h *BookHandler) HandleCreateBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	createdBook, err := h.booksService.CreateBook(&book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(createdBook)
	if err != nil {
		return
	}
}

//HandleGetBookByID handles request to fetch a book by ID

func (h *BookHandler) HandleGetBookByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
	}

	book, err := h.booksService.GetBookByID(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(book)
	if err != nil {
		return
	}

}

// HandleUpdateBook handles requests to update book details
func (h *BookHandler) HandleUpdateBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	updatedBook, err := h.booksService.UpdateBook(&book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(updatedBook)
	if err != nil {
		return
	}
}

//HandleDeleteBook handles requests to delete a book

func (h *BookHandler) HandleDeleteBook(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
	}

	if err := h.booksService.DeleteBook(uint(id)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte("Book deleted successfully!"))
	if err != nil {
		return
	}

}
