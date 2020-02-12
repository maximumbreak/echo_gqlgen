package resolver

import (
	"context"

	"github.com/beforesecond/gqlgen-todos/models"
	"github.com/beforesecond/gqlgen-todos/service"
)

func (r *queryResolver) Users(ctx context.Context) ([]*models.User, error) {

	var userModel = []*models.User{0: {
		Username: "usertest",
		//	Password: "123456",
	}}
	return userModel, nil
}

func (r *mutationResolver) Login(ctx context.Context, input models.InputLogin) (*models.AuthResponse, error) {
	// var userModel = models.UserLogin{
	// 	Username: "usertest",
	// 	Token:    "123456",
	// }
	token, err := service.AuthTokenHandlerByGraphQL(input)
	//log.Print(token)
	if err != nil {
		return nil, nil
	} else {
		return token, nil
	}
	return nil, nil
}

func (r *mutationResolver) CreateUser(ctx context.Context, input models.InputUser) (*models.User, error) {
	user, err := service.AuthRegisterGraphQLHandler(input)

	if err != nil {
		return nil, nil
	} else {
		result := models.User{
			ID:       user.ID,
			Username: user.Username,
			Token:    user.Token,
		}
		return &result, nil
	}
}
