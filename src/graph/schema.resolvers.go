package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"pranavt61/icon-graphql-api/graph/generated"
	"pranavt61/icon-graphql-api/graph/model"
	"pranavt61/icon-graphql-api/mongodb"
)

func (r *queryResolver) Block(ctx context.Context, hash string) (*model.Block, error) {
	return mongodb.DBConnection.FindByID(hash), nil
}

func (r *queryResolver) Blocks(ctx context.Context) ([]*model.Block, error) {
	return mongodb.DBConnection.All(), nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
