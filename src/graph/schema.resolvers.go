package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"pranavt61/icon-graphql-api/graph/generated"
	"pranavt61/icon-graphql-api/graph/model"
)

func (r *queryResolver) Blocks(ctx context.Context) ([]*model.Block, error) {
	blocks := make([]*model.Block, 0, 0)

	block := model.Block{"hash", 100}

	blocks = append(blocks, &block)

	return blocks, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
