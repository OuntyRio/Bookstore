package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.45

import (
	"context"

	"github.com/ountyrio/gql-bookstore/internal/graph/generated"
	gqlmodel "github.com/ountyrio/gql-bookstore/internal/graph/model"
	"github.com/ountyrio/gql-bookstore/internal/types"
)

// Create is the resolver for the create field.
func (r *authorsMutationNSResolver) Create(ctx context.Context, obj *gqlmodel.AuthorsMutationNs, input *gqlmodel.AuthorsMutationNsInput) (*gqlmodel.AuthorsMutationNsMutationResult, error) {
	body := &types.AuthorCreateRequestDto{
		Lastname:  *input.Lastname,
		Firstname: *input.Firstname,
	}

	author, err := r.authorService.Create(ctx, body)

	if err != nil {
		return nil, err
	}

	res := &gqlmodel.AuthorsMutationNsMutationResult{
		Changed: CastAuthorDtoToGql(&author.Created),
	}

	return res, nil
}

// Update is the resolver for the update field.
func (r *authorsMutationNSResolver) Update(ctx context.Context, obj *gqlmodel.AuthorsMutationNs, input *gqlmodel.AuthorsMutationNsInput) (*gqlmodel.AuthorsMutationNsMutationResult, error) {
	body := &types.AuthorUpdateRequestDto{
		ID:        *input.ID,
		Lastname:  *input.Lastname,
		Firstname: *input.Firstname,
	}

	author, err := r.authorService.Update(ctx, body)

	if err != nil {
		return nil, err
	}

	res := &gqlmodel.AuthorsMutationNsMutationResult{
		Changed: CastAuthorDtoToGql(&author.Updated),
	}

	return res, nil
}

// Delete is the resolver for the delete field.
func (r *authorsMutationNSResolver) Delete(ctx context.Context, obj *gqlmodel.AuthorsMutationNs, input *gqlmodel.AuthorsMutationNsInput) (*gqlmodel.AuthorsMutationNsMutationResult, error) {
	body := &types.AuthorIDRequestDto{
		ID: *input.ID,
	}

	err := r.authorService.Delete(ctx, body)

	if err != nil {
		return nil, err
	}

	res := &gqlmodel.AuthorsMutationNsMutationResult{
		Changed: &gqlmodel.Author{
			ID: input.ID,
		},
	}

	return res, nil
}

// AuthorsMutationNS returns generated.AuthorsMutationNSResolver implementation.
func (r *Resolver) AuthorsMutationNS() generated.AuthorsMutationNSResolver {
	return &authorsMutationNSResolver{r}
}

type authorsMutationNSResolver struct{ *Resolver }
