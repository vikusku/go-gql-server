package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/vikusku/go-gql-server/internal/gql"
	"github.com/vikusku/go-gql-server/internal/gql/models"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input models.UserInput) (*models.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateUser(ctx context.Context, input models.UserInput) (*models.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteUser(ctx context.Context, userID string) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

var id , email = "010", "test@sok.fi"

func (r *queryResolver) Users(ctx context.Context, userID *string) ([]*models.User, error) {
	return []*models.User{
		{
			ID:     &id,
			Email:  &email,
			UserID: userID,
		},
	}, nil
}

// Mutation returns gql.MutationResolver implementation.
func (r *Resolver) Mutation() gql.MutationResolver { return &mutationResolver{r} }

// Query returns gql.QueryResolver implementation.
func (r *Resolver) Query() gql.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
