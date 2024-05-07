package resolvers

import "github.com/ountyrio/gql-bookstore/internal/types"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	authorService types.AuthorService
	bookService   types.BookService
	genreService  types.GenreService
}

func NewResolver(
	as types.AuthorService,
	bs types.BookService,
	gs types.GenreService,
) *Resolver {
	return &Resolver{
		authorService: as,
		bookService:   bs,
		genreService:  gs,
	}
}
