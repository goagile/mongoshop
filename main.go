package main

import (
	"context"

	"github.com/goagile/mongoshop/db"
	"github.com/goagile/mongoshop/server"
)

const (
	DBURI  = "mongodb://127.0.0.1:27017"
	SRVURI = ":8081"
)

func main() {
	ctx := context.Background()
	db.Connect(ctx, DBURI)
	defer db.Disconnect(ctx)

	server.Setup()
	server.Run(SRVURI)
}
