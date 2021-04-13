package mongodb

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"pranavt61/icon-graphql-api/graph/model"
)

type DB struct {
	client *mongo.Client
}

var DBConnection *DB

func ConnectClient(url string, user string, pass string) error {
	// Check if connection already exists
	if DBConnection != nil {
		return nil
	}

	connection_uri := fmt.Sprintf("mongodb://%s:%s@%s", user, pass, url)
	client, err := mongo.NewClient(options.Client().ApplyURI(connection_uri))
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		return err
	}

	// Set global client
	DBConnection = &DB{
		client: client,
	}

	return nil
}

func (db *DB) FindBlockByHash(hash string) (*model.Block, error) {
	collection := db.client.Database("icon").Collection("blocks")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res := collection.FindOne(ctx, bson.M{"hash": hash})
	block := model.Block{}
	res.Decode(&block)
	return &block, nil
}

func (db *DB) Blocks(skip int, limit int) ([]*model.Block, error) {
	collection := db.client.Database("icon").Collection("blocks")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	findOptions := options.FindOptions{}
	findOptions.SetSkip(int64(skip))
	findOptions.SetLimit(int64(limit))
	findOptions.SetSort(bson.D{{"number", -1}})
	cur, err := collection.Find(ctx, bson.D{}, &findOptions)
	if err != nil {
		return nil, err
	}
	var blocks []*model.Block
	for cur.Next(ctx) {
		var block *model.Block
		err := cur.Decode(&block)
		if err != nil {
			return nil, err
		}
		blocks = append(blocks, block)
	}
	return blocks, nil
}

func (db *DB) FindTransactionByHash(hash string) (*model.Transaction, error) {
	collection := db.client.Database("icon").Collection("transactions")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res := collection.FindOne(ctx, bson.M{"hash": hash})
	transaction := model.Transaction{}
	res.Decode(&transaction)
	return &transaction, nil
}

func (db *DB) AllTransactions() ([]*model.Transaction, error) {
	collection := db.client.Database("icon").Collection("transactions")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	var transactions []*model.Transaction
	for cur.Next(ctx) {
		var transaction *model.Transaction
		err := cur.Decode(&transaction)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, transaction)
	}
	return transactions, nil
}

func (db *DB) FindLogsByTransactionHash(transactionHash string) ([]*model.Log, error) {
	collection := db.client.Database("icon").Collection("logs")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	cur, err := collection.Find(ctx, bson.M{"transaction_hash": transactionHash})
	if err != nil {
		return nil, err
	}
	var logs []*model.Log
	for cur.Next(ctx) {
		var log *model.Log
		err := cur.Decode(&log)
		if err != nil {
			return nil, err
		}
		logs = append(logs, log)
	}
	return logs, nil
}
