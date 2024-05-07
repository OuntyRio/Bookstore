package service

import (
	"context"
	"fmt"

	"github.com/ountyrio/gql-bookstore/internal/common"
	"github.com/ountyrio/gql-bookstore/internal/service/model"
	"github.com/ountyrio/gql-bookstore/internal/service/repository"
	"github.com/ountyrio/gql-bookstore/internal/types"
)

type BookServiceImpl struct {
	authorRepository repository.AuthorRepository
	bookRepository   repository.BookRepository
	genreRepository  repository.GenreRepository
}

func (bs *BookServiceImpl) GetById(ctx context.Context, request *types.BookIDRequestDto) (*types.BookGetByIdResponseDto, error) {
	book, err := bs.bookRepository.GetById(ctx, request.ID)
	if err != nil {
		return nil, err
	}

	if book.ID == 0 {
		return nil, fmt.Errorf("book with id=%d hasn`t found", request.ID)
	}

	res := &types.BookGetByIdResponseDto{
		Book: CastBookModelToDto(book),
	}
	return res, nil
}

func (bs *BookServiceImpl) Query(ctx context.Context, request *types.BookQueryRequestDto) (*types.BookQueryResponseDto, error) {
	books, count, err := bs.bookRepository.Query(ctx, &repository.Query{
		Limit: request.Limit,
		Page:  request.Page,
		Field: request.Field,
		Order: request.Order,
	})

	if err != nil {
		return nil, err
	}

	res := &types.BookQueryResponseDto{
		Books: common.SliceMap(CastBookModelToDto, books),
		Count: count,
	}

	return res, nil
}

func (bs *BookServiceImpl) Create(ctx context.Context, request *types.BookCreateRequestDto) (*types.BookCreateResponseDto, error) {
	authors, err := bs.authorRepository.GetMany(ctx, request.Authors)
	if err != nil {
		return nil, err
	}

	if len(authors) == 0 {
		return nil, fmt.Errorf("Authors are not found")
	}

	genre, err := bs.genreRepository.GetById(ctx, request.Genre)
	if err != nil {
		return nil, err
	}

	if genre.ID == 0 {
		return nil, fmt.Errorf("Genre is not found")
	}

	book, err := bs.bookRepository.Create(ctx, &model.Book{
		Title:   request.Title,
		GenreID: genre.ID,
		Authors: authors,
	})

	if err != nil {
		return nil, err
	}

	res := &types.BookCreateResponseDto{
		Created: CastBookModelToDto(book),
	}
	return res, nil
}

func (bs *BookServiceImpl) Update(ctx context.Context, request *types.BookUpdateRequestDto) (*types.BookUpdateResponseDto, error) {
	authors, err := bs.authorRepository.GetMany(ctx, request.Authors)
	if err != nil {
		return nil, err
	}

	if len(authors) == 0 {
		return nil, fmt.Errorf("Authors are not found")
	}

	genre, err := bs.genreRepository.GetById(ctx, request.Genre)
	if err != nil {
		return nil, err
	}

	if genre.ID == 0 {
		return nil, fmt.Errorf("Genre is not found")
	}

	book, err := bs.bookRepository.Update(ctx, &model.Book{
		ID:      request.ID,
		Title:   request.Title,
		GenreID: genre.ID,
		Authors: authors,
	})
	if err != nil {
		return nil, err
	}

	res := &types.BookUpdateResponseDto{
		Updated: CastBookModelToDto(book),
	}
	return res, nil
}

func (bs *BookServiceImpl) Delete(ctx context.Context, request *types.BookIDRequestDto) error {
	err := bs.bookRepository.Delete(ctx, request.ID)
	if err != nil {
		return err
	}
	return nil
}

func NewBookService(ar repository.AuthorRepository, br repository.BookRepository, gr repository.GenreRepository) *BookServiceImpl {
	return &BookServiceImpl{
		authorRepository: ar,
		bookRepository:   br,
		genreRepository:  gr,
	}
}
