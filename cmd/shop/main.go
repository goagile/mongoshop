package main

import (
	"context"
	"fmt"

	"github.com/goagile/mongoshop/internal/db"
	"github.com/goagile/mongoshop/internal/server"
)

const (
	DBURI  = "mongodb://127.0.0.1:27017"
	SRVURI = ":8080"
)

func main() {
	ctx := context.Background()
	db.Connect(ctx, DBURI)
	defer db.Disconnect(ctx)

	server.Setup()
	fmt.Println("Server start at:", SRVURI)
	server.Run(SRVURI)
}
