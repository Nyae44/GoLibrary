package main

import (
	"github.com/nyae44/GoLibrary/internal/handlers"
	"github.com/nyae44/GoLibrary/internal/models"
	"github.com/nyae44/GoLibrary/internal/repositories"
	"github.com/nyae44/GoLibrary/internal/services"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func main() {
	dsn := "host=localhost user=postgres password=example dbname=postgres port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database")
	}
	// Auto migrate the models
	err = db.AutoMigrate(&models.User{}, models.Book{}, models.Member{}, models.Transaction{})
	if err != nil {
		log.Fatal("Failed to migrate database")
	}
	//Initialize repositories
	memberRepo := repositories.NewMemberRepository(db)
	bookRepo := repositories.NewBookRepository(db)
	transactionRepo := repositories.NewTransactionRepository(db)

	// Initialize services
	memberService := services.NewMemberService(memberRepo)
	bookService := services.NewBookService(bookRepo)
	transactionService := services.NewTransactionService(transactionRepo, bookRepo, memberRepo)

	// Initialize handlers
	bookHandler := handlers.NewBookHandler(bookService)
	memberHandler := handlers.NewMemberHandler(memberService)
	transactionHandler := handlers.NewTransactionHandler(transactionService)

	// Routes
	// Books
	http.HandleFunc("/books/create-book", bookHandler.HandleCreateBook)
	http.HandleFunc("/books/{id}", bookHandler.HandleGetBookByID)
	http.HandleFunc("books/update-book", bookHandler.HandleUpdateBook)
	http.HandleFunc("/books/delete-book", bookHandler.HandleDeleteBook)

	// Members
	http.HandleFunc("/members/create-member", memberHandler.HandleCreateMember)
	http.HandleFunc("/members/{id}", memberHandler.HandleGetMemberByID)
	http.HandleFunc("/members/update-member", memberHandler.HandleUpdateMember)
	http.HandleFunc("/members", memberHandler.HandleListMembers)
	//http.HandleFunc("members/delete-member", memberHandler.HandleDeleteMember)

	//Transactions
	http.HandleFunc("/transactions", transactionHandler.HandleGetTransactions)
	http.HandleFunc("/transactions/{id}", transactionHandler.HandleTransactionByID)
	http.HandleFunc("/transaction/borrow", transactionHandler.HandleBorrowBook)
	http.HandleFunc("/transaction/return", transactionHandler.HandleReturnBook)

	log.Println("Starting server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
