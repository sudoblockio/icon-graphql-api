package mongodb

import (
	"context"
	"fmt"
	"log"
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

func ConnectClient(url string, user string, pass string) {
	connection_uri := fmt.Sprintf("mongodb://%s:%s@%s", user, pass, url)
	client, err := mongo.NewClient(options.Client().ApplyURI(connection_uri))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	DBConnection = &DB{
		client: client,
	}
}

func (db *DB) FindByID(hash string) *model.Block {
	collection := db.client.Database("icon").Collection("blocks")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res := collection.FindOne(ctx, bson.M{"hash": hash})
	block := model.Block{}
	res.Decode(&block)
	return &block
}

func (db *DB) All() []*model.Block {
	collection := db.client.Database("icon").Collection("blocks")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	var blocks []*model.Block
	for cur.Next(ctx) {
		var block *model.Block
		err := cur.Decode(&block)
		if err != nil {
			log.Fatal(err)
		}
		blocks = append(blocks, block)
	}
	return blocks
}
