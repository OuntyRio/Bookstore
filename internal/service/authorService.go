package service

import (
	"context"
	"fmt"

	"github.com/ountyrio/gql-bookstore/internal/common"
	"github.com/ountyrio/gql-bookstore/internal/service/model"
	"github.com/ountyrio/gql-bookstore/internal/service/repository"
	"github.com/ountyrio/gql-bookstore/internal/types"
)

type AuthorServiceImpl struct {
	authorRepository repository.AuthorRepository
}

func (as *AuthorServiceImpl) GetById(ctx context.Context, request *types.AuthorIDRequestDto) (*types.AuthorGetByIdResponseDto, error) {
	author, err := as.authorRepository.GetById(ctx, request.ID)
	if err != nil {
		return nil, err
	}

	if author.ID == 0 {
		return nil, fmt.Errorf("author with id=%d hasn`t found", request.ID)
	}

	res := &types.AuthorGetByIdResponseDto{
		Author: CastAuthorModelToDto(author),
	}

	return res, nil
}

func (as *AuthorServiceImpl) Query(ctx context.Context, request *types.AuthorQueryRequestDto) (*types.AuthorQueryResponseDto, error) {
	authors, count, err := as.authorRepository.Query(ctx, &repository.Query{
		Limit: request.Limit,
		Page:  request.Page,
		Field: request.Field,
		Order: request.Order,
	})

	if err != nil {
		return nil, err
	}

	res := &types.AuthorQueryResponseDto{
		Authors: common.SliceMap(CastAuthorModelToDto, authors),
		Count:   count,
	}

	return res, err
}

func (as *AuthorServiceImpl) Create(ctx context.Context, request *types.AuthorCreateRequestDto) (*types.AuthorCreateResponseDto, error) {
	author, err := as.authorRepository.Create(ctx, &model.Author{
		Firstname: request.Firstname,
		Lastname:  request.Lastname,
	})
	if err != nil {
		return nil, err
	}

	res := &types.AuthorCreateResponseDto{
		Created: CastAuthorModelToDto(author),
	}
	return res, nil
}

func (as *AuthorServiceImpl) Update(ctx context.Context, request *types.AuthorUpdateRequestDto) (*types.AuthorUpdateResponseDto, error) {
	author, err := as.authorRepository.Update(ctx, &model.Author{
		ID:        request.ID,
		Firstname: request.Firstname,
		Lastname:  request.Lastname,
	})
	if err != nil {
		return nil, err
	}
	result := &types.AuthorUpdateResponseDto{
		Updated: CastAuthorModelToDto(author),
	}
	return result, nil
}

func (as *AuthorServiceImpl) Delete(ctx context.Context, request *types.AuthorIDRequestDto) error {
	err := as.authorRepository.Delete(ctx, request.ID)
	if err != nil {
		return err
	}
	return nil
}

func NewAuthorService(r repository.AuthorRepository) *AuthorServiceImpl {
	return &AuthorServiceImpl{
		authorRepository: r,
	}
}
