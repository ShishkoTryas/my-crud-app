package db

import (
	"context"
	"database/sql"
	log "github.com/sirupsen/logrus"
	"my-crud-app/internal/domain"
)

type Book struct {
	db *sql.DB
}

func NewBooks(db *sql.DB) *Book {
	return &Book{
		db: db,
	}
}

func (b *Book) GetBook(ctx context.Context, id int) (domain.Book, error) {
	var book domain.Book

	err := b.db.QueryRow("SELECT * FROM books where id = $1", id).Scan(&book.Id, &book.Name, &book.Price, &book.Time)

	return book, err
}

func (b *Book) GetBooks(ctx context.Context) ([]domain.Book, error) {
	var books []domain.Book

	rows, err := b.db.Query("SELECT * FROM books")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		b := domain.Book{}
		err := rows.Scan(&b.Id, &b.Name, &b.Price, &b.Time)
		if err != nil {
			log.Fatal(err)
		}

		books = append(books, b)
	}

	return books, err
}

func (b *Book) CreateBook(ctx context.Context, book domain.CreateBookInput) error {
	_, err := b.db.Exec("INSERT into books (name, price) VALUES ($1, $2)", book.Name, book.Price)

	return err
}

func (b *Book) UpdateBook(ctx context.Context, id int, book domain.CreateBookInput) error {
	_, err := b.db.Exec("update books set name=$1, price=$2 where id=$3", book.Name, book.Price, id)
	return err
}

func (b *Book) DeleteBook(ctx context.Context, id int) error {
	_, err := b.db.Exec("delete from books where id=$1", id)

	return err
}
