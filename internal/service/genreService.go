package service

import (
	"context"
	"fmt"

	"github.com/ountyrio/gql-bookstore/internal/common"
	"github.com/ountyrio/gql-bookstore/internal/service/model"
	"github.com/ountyrio/gql-bookstore/internal/service/repository"
	"github.com/ountyrio/gql-bookstore/internal/types"
)

type GenreServiceImpl struct {
	genreRepository repository.GenreRepository
}

func (gs *GenreServiceImpl) GetById(ctx context.Context, request *types.GenreIDRequestDto) (*types.GenreGetByIdResponseDto, error) {
	genre, err := gs.genreRepository.GetById(ctx, request.ID)
	if err != nil {
		return nil, err
	}

	if genre.ID == 0 {
		return nil, fmt.Errorf("genre with id=%d hasn`t found", request.ID)
	}

	res := &types.GenreGetByIdResponseDto{
		Genre: CastGenreModelToDto(genre),
	}
	return res, nil
}

func (gs *GenreServiceImpl) Query(ctx context.Context, request *types.GenreQueryRequestDto) (*types.GenreQueryResponseDto, error) {
	genres, count, err := gs.genreRepository.Query(ctx, &repository.Query{
		Limit: request.Limit,
		Page:  request.Page,
		Field: request.Field,
		Order: request.Order,
	})

	if err != nil {
		return nil, err
	}

	res := &types.GenreQueryResponseDto{
		Genres: common.SliceMap(CastGenreModelToDto, genres),
		Count:  count,
	}

	return res, nil
}

func (bs *GenreServiceImpl) Create(ctx context.Context, request *types.GenreCreateRequestDto) (*types.GenreCreateResponseDto, error) {
	genre, err := bs.genreRepository.Create(ctx, &model.Genre{
		Name: request.Name,
	})
	if err != nil {
		return nil, err
	}
	res := &types.GenreCreateResponseDto{
		Created: CastGenreModelToDto(genre),
	}
	return res, nil
}

func (bs *GenreServiceImpl) Update(ctx context.Context, request *types.GenreUpdateRequestDto) (*types.GenreUpdateResponseDto, error) {
	genre, err := bs.genreRepository.Update(ctx, &model.Genre{
		ID:   request.ID,
		Name: request.Name,
	})
	if err != nil {
		return nil, err
	}
	res := &types.GenreUpdateResponseDto{
		Updated: CastGenreModelToDto(genre),
	}
	return res, nil
}

func (bs *GenreServiceImpl) Delete(ctx context.Context, request *types.GenreIDRequestDto) error {
	err := bs.genreRepository.Delete(ctx, request.ID)
	if err != nil {
		return err
	}
	return nil
}

func NewGenreService(r repository.GenreRepository) *GenreServiceImpl {
	return &GenreServiceImpl{
		genreRepository: r,
	}
}
