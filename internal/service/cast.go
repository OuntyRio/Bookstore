package service

import (
	"github.com/ountyrio/gql-bookstore/internal/common"
	"github.com/ountyrio/gql-bookstore/internal/service/model"
	"github.com/ountyrio/gql-bookstore/internal/types"
)

const TIME_FORMAT = "2006-01-02 15:04:05"

func CastAuthorModelToDto(m *model.Author) types.Author {
	return types.Author{
		ID:        m.ID,
		Firstname: m.Firstname,
		Lastname:  m.Lastname,
		Books:     common.SliceMap(CastBookModelToDto, m.Books),
		CreatedAt: m.CreatedAt.Format(TIME_FORMAT),
	}
}

func CastBookModelToDto(m *model.Book) types.Book {
	return types.Book{
		ID:        m.ID,
		Title:     m.Title,
		Genre:     m.GenreID,
		Authors:   common.SliceMap(CastAuthorModelToDto, m.Authors),
		CreatedAt: m.CreatedAt.Format(TIME_FORMAT),
		UpdatedAt: m.UpdatedAt.Format(TIME_FORMAT),
	}
}

func CastGenreModelToDto(m *model.Genre) types.Genre {
	return types.Genre{
		ID:        m.ID,
		Name:      m.Name,
		CreatedAt: m.CreatedAt.Format(TIME_FORMAT),
		UpdatedAt: m.UpdatedAt.Format(TIME_FORMAT),
	}
}
