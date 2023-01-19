package service

import (
	"context"
	"my-crud-app/internal/domain"
)

type BooksRepository interface {
	GetBook(ctx context.Context, id int) (domain.Book, error)
	GetBooks(ctx context.Context) ([]domain.Book, error)
	CreateBook(ctx context.Context, book domain.CreateBookInput) error
	UpdateBook(ctx context.Context, id int, book domain.CreateBookInput) error
	DeleteBook(ctx context.Context, id int) error
}

type Books struct {
	repo BooksRepository
}

func NewBooks(repo BooksRepository) *Books {
	return &Books{
		repo: repo,
	}
}

func (b *Books) CreateBook(ctx context.Context, book domain.CreateBookInput) error {
	return b.repo.CreateBook(ctx, book)
}

func (b *Books) GetBook(ctx context.Context, id int) (domain.Book, error) {
	return b.repo.GetBook(ctx, id)
}

func (b *Books) GetBooks(ctx context.Context) ([]domain.Book, error) {
	return b.repo.GetBooks(ctx)
}

func (b *Books) DeleteBook(ctx context.Context, id int) error {
	return b.repo.DeleteBook(ctx, id)
}

func (b *Books) UpdateBook(ctx context.Context, id int, inp domain.CreateBookInput) error {
	return b.repo.UpdateBook(ctx, id, inp)
}
