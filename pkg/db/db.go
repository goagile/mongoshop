package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	Client *mongo.Client
	DB     *mongo.Database
	Books  *mongo.Collection
)

// Connect - open Client connection
func Connect(ctx context.Context, uri string) {
	Client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatalf("DB %v err: %v", uri, err)
	}
	ctx, _ = context.WithTimeout(ctx, 10*time.Second)
	if err := Client.Connect(ctx); err != nil {
		log.Fatalf("DB Connect:%v", err)
	}
	DB = Client.Database("bookstore")
	Books = DB.Collection("books")
}

// Disconnect - close Client connection
func Disconnect(ctx context.Context) {
	Client.Disconnect(ctx)
}

// Book - Model Entity
type Book struct {
	ID     primitive.ObjectID `json:"-", bson:"_id"`
	Title  string             `json:"title", bson:"title"`
	Author string             `json:"author", bson:"author"`
}

// FindBooks ...
func FindBooks(ctx context.Context) ([]*Book, error) {
	var books []*Book
	cursor, err := Books.Find(ctx, bson.M{})
	if err != nil {
		return books, err
	}
	if err := cursor.All(ctx, &books); err != nil {
		return books, err
	}
	defer cursor.Close(ctx)
	return books, nil
}
