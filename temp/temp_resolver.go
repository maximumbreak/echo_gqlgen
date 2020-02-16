package temp

import (
	"context"

	"github.com/beforesecond/gqlgen-todos/generated"
	"github.com/beforesecond/gqlgen-todos/models"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{}

func (r *Resolver) Mutation() generated.MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() generated.QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateTodo(ctx context.Context, input models.NewTodo) (*models.Todo, error) {
	panic("not implemented")
}
func (r *mutationResolver) CreateUser(ctx context.Context, input models.InputUser) (*models.User, error) {
	panic("not implemented")
}
func (r *mutationResolver) Login(ctx context.Context, input models.InputLogin) (*models.AuthResponse, error) {
	panic("not implemented")
}
func (r *mutationResolver) RevokeToken(ctx context.Context, token string) (*models.AuthResponse, error) {
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Todos(ctx context.Context) ([]*models.Todo, error) {
	panic("not implemented")
}
func (r *queryResolver) Users(ctx context.Context) ([]*models.User, error) {
	panic("not implemented")
}
