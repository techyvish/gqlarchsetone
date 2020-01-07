package graph

import (
	"context"

	"github.com/techyvish/gqlarchsetone/graph/models"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{}

func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Person(ctx context.Context) ([]*models.Person, error) {
	panic("not implemented")
}
func (r *queryResolver) Pet(ctx context.Context) ([]*models.Pet, error) {
	panic("not implemented")
}
