package server

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/goagile/mongoshop/db"
)

var (
	router *gin.Engine
)

// Setup - configure http server
func Setup() {
	gin.SetMode(gin.DebugMode)

	router = gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.LoadHTMLGlob("./server/templates/*")
	router.Static("/css", "./server/static/css")
	router.Static("/js", "./server/static/js")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	router.GET("/books", func(c *gin.Context) {
		c.HTML(http.StatusOK, "books.html", nil)
	})

	// 404 page
	router.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusOK, "notfound.html", nil)
	})

	// API
	api := router.Group("api")
	v1 := api.Group("v1")
	{
		v1.GET("/books", GetBooks)
	}
}

// Run - run setups server
func Run(uri string) {
	router.Run(uri)
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
