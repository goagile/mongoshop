package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// NewClient - create new mongo client
func NewClient(ctx context.Context, uri string) *mongo.Client {
	opts := options.Client().ApplyURI(uri)

	c, err := mongo.NewClient(opts)
	if err != nil {
		log.Fatalf("DB %v err: %v", uri, err)
	}

	if err := c.Connect(ctx); err != nil {
		log.Fatalf("DB %v Connect:%v", uri, err)
	}

	if err := c.Ping(ctx, nil); err != nil {
		log.Fatalf("DB %v Ping:%v", uri, err)
	}

	return c
}
