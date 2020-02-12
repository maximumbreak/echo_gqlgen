package resolver

import (
	"context"

	test_gqlgen "github.com/beforesecond/gqlgen-todos"
	"github.com/beforesecond/gqlgen-todos/models"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{}

func (r *Resolver) Mutation() test_gqlgen.MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() test_gqlgen.QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateTodo(ctx context.Context, input models.NewTodo) (*models.Todo, error) {
	panic("not implemented")
}
func (r *mutationResolver) CreateUser(ctx context.Context, input models.InputUser) (*models.User, error) {
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Todos(ctx context.Context) ([]*models.Todo, error) {
	panic("not implemented")
}
