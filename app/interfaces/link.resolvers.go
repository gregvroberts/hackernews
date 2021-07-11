package interfaces

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/gregvroberts/hackernews/app/generated"
	"github.com/gregvroberts/hackernews/app/models"
)

func (r *mutationResolver) CreateLink(ctx context.Context, input models.NewLink) (*models.Link, error) {
	var link models.Link
	var user models.User
	link.Address = input.Address
	link.Title = input.Title
	user.Name = "test name"
	link.User = &user
	return &link, nil
}

func (r *queryResolver) Links(ctx context.Context) ([]*models.Link, error) {
	var links []*models.Link
	dummyLink := models.Link{
		Title:   "Our dummy link",
		Address: "https://ddress.com",
		User:    &models.User{Name: "admin"},
	}

	links = append(links, &dummyLink)
	return links, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
