package book

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	DB    *mongo.Database
	Books *mongo.Collection
)

// Book - Model Entity
type Book struct {
	Title  string `json:"title", bson:"title"`
	Author string `json:"author", bson:"author"`
}

// FindAll ...
func FindAll(ctx context.Context) ([]*Book, error) {
	var bs []*Book
	c, err := Books.Find(ctx, bson.M{})
	if err != nil {
		return bs, err
	}
	if err := c.All(ctx, &bs); err != nil {
		return bs, err
	}
	defer c.Close(ctx)
	return bs, nil
}
