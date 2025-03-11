package postgres

import (
	"books/app/domain/entity"
	"database/sql"

	"github.com/google/uuid"
)

type BookRepository struct {
	db *sql.DB
}

// NewBookRepository returns new BookRepository.
func NewBookRepository(db * sql.DB) *BookRepository {
	return &BookRepository{
		db: db,
	}
}

func (r *BookRepository) CreateBook(book *entity.Book) (string, error) {
	id := uuid.New().String()
	err := r.db.QueryRow(
		"INSERT INTO books (id, title, author, isbn) VALUES ($1, $2, $3, $4) RETURNING id",
		id, book.Title, book.Author, book.ISBN,
	).Scan(&book.ID)
	
	if err != nil {
		return "", err
	}

	return id, nil
}