package resolver

import (
	"context"

	"github.com/beforesecond/gqlgen-todos/models"
)

func (r *queryResolver) Users(ctx context.Context) ([]*models.User, error) {

	var userModel = []*models.User{0: {
		Username: "usertest",
		Password: "123456",
	}}
	return userModel, nil
}

func (r *mutationResolver) Login(ctx context.Context, input models.InputLogin) (*models.UserLogin, error) {
	var userModel = models.UserLogin{
		Username: "usertest",
		Token:    "123456",
	}
	return &userModel, nil
}
