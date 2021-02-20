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
	block, err := mongodb.DBConnection.FindBlockByHash(hash)
	if err != nil {
		return nil, err
	}

	return block, nil
}

func (r *queryResolver) Blocks(ctx context.Context) ([]*model.Block, error) {
	blocks, err := mongodb.DBConnection.AllBlocks()
	if err != nil {
		return nil, err
	}

	return blocks, nil
}

func (r *queryResolver) Transaction(ctx context.Context, hash string) (*model.Transaction, error) {
	transaction, err := mongodb.DBConnection.FindTransactionByHash(hash)
	if err != nil {
		return nil, err
	}

	return transaction, err
}

func (r *queryResolver) Transactions(ctx context.Context) ([]*model.Transaction, error) {
	transactions, err := mongodb.DBConnection.AllTransactions()
	if err != nil {
		return nil, err
	}

	return transactions, err
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
