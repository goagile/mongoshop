package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/goagile/mongoshop/db"
)

func main() {
	ctx := context.Background()

	db.Init(ctx, "mongodb://127.0.0.1:27017")

	r := gin.Default()
	r.GET("/books", GetBooks)
	r.Run(":8081")

	db.Client.Disconnect(ctx)
}

// GetBooks - GET /books
func GetBooks(c *gin.Context) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	books, err := db.FindBooks(ctx)
	if err != nil {
		log.Println("FindBooks", err)
		c.String(http.StatusNotFound, "books are not found")
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": books})
}
