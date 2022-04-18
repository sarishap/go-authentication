package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/sarishap/go-authentication/graph/generated"
	"github.com/sarishap/go-authentication/graph/model"
)

func (r *mutationResolver) CreateUserDetail(ctx context.Context, input *model.NewUserDetail) (*model.UserDetail, error) {
	userdetails := &model.UserDetail{
		ID:      fmt.Sprintf("%d", rand.Int()),
		Name:    input.Name,
		Phone:   input.Phone,
		Address: input.Address,
		User:    &model.User{ID: input.UserID, Username: input.Name},
	}
	r.userdetails = append(r.userdetails, userdetails)
	return userdetails, nil
}

func (r *mutationResolver) Register(ctx context.Context, input *model.NewUser) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) Login(ctx context.Context, input model.Login) (*model.AuthResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) ForgotPassword(ctx context.Context, newPassword string, key string) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) RefreshToken(ctx context.Context, token string, refreshToken string) (*model.AuthResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) UserDetails(ctx context.Context) ([]*model.UserDetail, error) {
	return r.userdetails, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
