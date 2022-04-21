package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strconv"

	"github.com/sarishap/go-authentication/graph/generated"
	"github.com/sarishap/go-authentication/graph/model"
	"github.com/sarishap/go-authentication/internal/userdetails"
	"github.com/sarishap/go-authentication/internal/users"
	"github.com/sarishap/go-authentication/jwt"
	"github.com/sarishap/go-authentication/middleware/auth"
)

func (r *mutationResolver) CreateUserDetail(ctx context.Context, input *model.NewUserDetail) (*model.UserDetail, error) {
	user := auth.ForContext(ctx)
	if user == nil {
		return &model.UserDetail{}, errors.New("access denied")
	}
	var userdetail userdetails.UserDetails
	userdetail.User = user
	UserID := userdetail.Save()
	graphqlUser := &model.User{
		ID:       user.ID,
		Username: user.Username,
	}
	return &model.UserDetail{ID: strconv.FormatInt(UserID, 10), Name: userdetail.Name, Address: userdetail.Address, Phone: userdetail.Phone, User: graphqlUser}, nil

}

func (r *mutationResolver) Register(ctx context.Context, input *model.NewUser) (string, error) {
	var user users.User
	user.Username = input.Username
	user.Password = input.Password
	user.Create()
	token, err := jwt.GenerateToken(user.Username)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (r *mutationResolver) Login(ctx context.Context, input model.Login) (string, error) {
	var user users.User
	user.Username = input.Username
	user.Password = input.Password
	correct := user.Authenticate()
	if !correct {
		return "", errors.New("invalid username or password")
	} else {
		log.Println("User authenticated")
	}

	token, err := jwt.GenerateToken(user.Username)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (r *mutationResolver) ForgotPassword(ctx context.Context, input *model.UpdatePassword) (string, error) {
	var user users.User
	user.Password = input.NewPassword
	user.Username = input.Username
	user.UpdatePassword()
	token, err := jwt.GenerateToken(user.Username)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (r *mutationResolver) RefreshToken(ctx context.Context, input model.RefreshTokenInput) (string, error) {
	username, err := jwt.ParseToken(input.Token)
	if err != nil {
		return "", fmt.Errorf("access denied")
	}
	token, err := jwt.GenerateToken(username)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (r *queryResolver) UserDetails(ctx context.Context) ([]*model.UserDetail, error) {
	var result []*model.UserDetail
	dbUserDetails := userdetails.FetchData()
	for _, userdetail := range dbUserDetails {
		graphqlUser := &model.User{
			//ID:       userdetail.User.ID,
			Username: userdetail.User.Username,
		}
		result = append(result, &model.UserDetail{ID: userdetail.ID, Name: userdetail.Name, Address: userdetail.Address, User: graphqlUser})
	}
	return result, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
