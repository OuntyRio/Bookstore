package types

import "context"

type Genre struct {
	ID        int
	Name      string
	CreatedAt string
	UpdatedAt string
}

type GenreService interface {
	GetById(ctx context.Context, request *GenreIDRequestDto) (*GenreGetByIdResponseDto, error)
	Query(ctx context.Context, request *GenreQueryRequestDto) (*GenreQueryResponseDto, error)
	Create(ctx context.Context, request *GenreCreateRequestDto) (*GenreCreateResponseDto, error)
	Update(ctx context.Context, request *GenreUpdateRequestDto) (*GenreUpdateResponseDto, error)
	Delete(ctx context.Context, request *GenreIDRequestDto) error
}

type GenreIDRequestDto struct {
	ID int
}

type GenreQueryRequestDto struct {
	Limit int
	Page  int
	Field string
	Order string
}

type GenreQueryResponseDto struct {
	Genres []Genre
	Count  int64
}

type GenreCreateRequestDto struct {
	Name string
}

type GenreUpdateRequestDto struct {
	ID   int
	Name string
}

type GenreGetByIdResponseDto struct {
	Genre Genre
}

type GenreCreateResponseDto struct {
	Created Genre
}

type GenreUpdateResponseDto struct {
	Updated Genre
}
