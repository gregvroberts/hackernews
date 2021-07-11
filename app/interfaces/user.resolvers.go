package interfaces

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/gregvroberts/hackernews/app/models"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input models.NewUser) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) Login(ctx context.Context, input models.Login) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) RefreshToken(ctx context.Context, input models.RefreshTokenInput) (string, error) {
	panic(fmt.Errorf("not implemented"))
}
