package types

import (
	"context"
)

type Book struct {
	ID        int
	Title     string
	Authors   []Author
	Genre     int
	CreatedAt string
	UpdatedAt string
}

type BookService interface {
	GetById(ctx context.Context, request *BookIDRequestDto) (*BookGetByIdResponseDto, error)
	Query(ctx context.Context, request *BookQueryRequestDto) (*BookQueryResponseDto, error)
	Create(ctx context.Context, request *BookCreateRequestDto) (*BookCreateResponseDto, error)
	Update(ctx context.Context, request *BookUpdateRequestDto) (*BookUpdateResponseDto, error)
	Delete(ctx context.Context, request *BookIDRequestDto) error
}

type BookIDRequestDto struct {
	ID int
}

type BookQueryRequestDto struct {
	Limit int
	Page  int
	Field string
	Order string
}

type BookCreateRequestDto struct {
	Title   string
	Authors []int
	Genre   int
}

type BookUpdateRequestDto struct {
	ID      int
	Title   string
	Authors []int
	Genre   int
}

type BookGetByIdResponseDto struct {
	Book Book
}

type BookQueryResponseDto struct {
	Books []Book
	Count int64
}

type BookCreateResponseDto struct {
	Created Book
}

type BookUpdateResponseDto struct {
	Updated Book
}
