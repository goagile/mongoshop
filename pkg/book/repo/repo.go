package repo

import (
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	DB    *mongo.Database
	Books *mongo.Collection
)
