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
	dsn := "host=localhost user=daley password=postgres dbname=postgres port=5432"
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
	http.HandleFunc("/books", bookHandler.HandleCreateBook)
	http.HandleFunc("/members", memberHandler.HandleCreateMember)
	http.HandleFunc("/transactions", transactionHandler.HandleGetTransactions)

	log.Println("Starting server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
