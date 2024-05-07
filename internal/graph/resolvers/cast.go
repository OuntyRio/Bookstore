package resolvers

import (
	"github.com/ountyrio/gql-bookstore/internal/common"
	gqlmodel "github.com/ountyrio/gql-bookstore/internal/graph/model"
	"github.com/ountyrio/gql-bookstore/internal/types"
)

func CastAuthorDtoToGql(d *types.Author) *gqlmodel.Author {
	return &gqlmodel.Author{
		ID:        &d.ID,
		Firstname: &d.Firstname,
		Lastname:  &d.Lastname,
		Books:     common.SliceMapPtr(CastBookDtoToGql, d.Books),
		CreatedAt: &d.CreatedAt,
		UpdatedAt: &d.UpdatedAt,
	}
}

func CastBookDtoToGql(d *types.Book) *gqlmodel.Book {
	return &gqlmodel.Book{
		ID:      &d.ID,
		Title:   &d.Title,
		Authors: common.SliceMapPtr(CastAuthorDtoToGql, d.Authors),
		Genre: CastGenreDtoToGql(&types.Genre{
			ID: d.Genre,
		}),
		CreatedAt: &d.CreatedAt,
		UpdatedAt: &d.UpdatedAt,
	}
}

func CastGenreDtoToGql(d *types.Genre) *gqlmodel.Genre {
	return &gqlmodel.Genre{
		ID:        &d.ID,
		Name:      &d.Name,
		CreatedAt: &d.CreatedAt,
		UpdatedAt: &d.UpdatedAt,
	}
}
