package types

import (
	"context"
)

type Author struct {
	ID        int
	Firstname string
	Lastname  string
	Books     []Book
	CreatedAt string
	UpdatedAt string
}

type AuthorService interface {
	GetById(ctx context.Context, request *AuthorIDRequestDto) (*AuthorGetByIdResponseDto, error)
	Query(ctx context.Context, request *AuthorQueryRequestDto) (*AuthorQueryResponseDto, error)
	Create(ctx context.Context, request *AuthorCreateRequestDto) (*AuthorCreateResponseDto, error)
	Update(ctx context.Context, request *AuthorUpdateRequestDto) (*AuthorUpdateResponseDto, error)
	Delete(ctx context.Context, request *AuthorIDRequestDto) error
}

type AuthorIDRequestDto struct {
	ID int
}

type AuthorQueryRequestDto struct {
	Limit int
	Page  int
	Field string
	Order string
}

type AuthorQueryResponseDto struct {
	Authors []Author
	Count   int64
}

type AuthorCreateRequestDto struct {
	Firstname string
	Lastname  string
}

type AuthorUpdateRequestDto struct {
	ID        int
	Firstname string
	Lastname  string
}

type AuthorGetByIdResponseDto struct {
	Author Author
}

type AuthorCreateResponseDto struct {
	Created Author
}

type AuthorUpdateResponseDto struct {
	Updated Author
}
