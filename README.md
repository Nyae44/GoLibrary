# GoLibrary API

GoLibrary is a simple library management system built using **Golang**, **GORM**, and **PostgreSQL**. It provides an API for managing books, members, and transactions (borrowing and returning books).

## Features
- Add, update, and delete books
- Register and manage library members
- Borrow and return books with transaction tracking
- Ensure book availability and member balance validation

## Project Structure
```plaintext
GoLibrary/
├── cmd/
│   └── library/
│       └── main.go
├── internal/
│   ├── handlers/          # HTTP request handlers
│   ├── models/            # Database models
│   ├── repositories/      # Database access layer
│   ├── services/          # Business logic layer
│   ├── utils/             # Utility functions
├── go.mod
├── go.sum
├── LICENSE
├── README.md
```

## Installation

### Prerequisites
- **Go 1.22+** installed
- **PostgreSQL** installed and running

### Clone the Repository
```sh
git clone https://github.com/nyae44/GoLibrary.git
cd GoLibrary
```

### Setup Database
Update your **PostgreSQL** connection string in `cmd/library/main.go`:
```go
db, err := gorm.Open(postgres.Open("your_postgres_dsn"), &gorm.Config{})
```

Run database migrations:
```sh
go run cmd/library/main.go
```

### Run the API Server
```sh
go run cmd/library/main.go
```
Server runs on **http://localhost:8080**

## API Endpoints

### Books
- `POST /books` - Add a new book
- `GET /books` - Get all books
- `GET /books/{id}` - Get book by ID
- `PUT /books/{id}` - Update book details
- `DELETE /books/{id}` - Delete a book

### Members
- `POST /members` - Register a new member
- `GET /members` - Get all members
- `GET /members/{id}` - Get member by ID

### Transactions
- `POST /transactions/borrow` - Borrow a book
- `POST /transactions/return` - Return a book
- `GET /transactions` - Get all transactions

## Contributing
Feel free to fork and submit pull requests! 🚀

## License
This project is licensed under the **MIT License**.
