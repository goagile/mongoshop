package main

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/gin-gonic/gin"
	_ "github.com/goagile/mongoshop/api/docs"
	"github.com/goagile/mongoshop/cmd/shop/controller"
	"github.com/goagile/mongoshop/pkg/book"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

const (
	DBAddr  = "mongodb://127.0.0.1:27017"
	SrvAddr = "127.0.0.1:8080"
	TMPL    = "./web/templates/*"
	CSS     = "./web/static/css"
	JS      = "./web/static/js"
)

// @title Hello
// @BasePath /api/v1
func main() {
	ctx := context.Background()
	c := setupDBClient(ctx, DBAddr)
	defer c.Disconnect(ctx)

	s := setupWebServer()
	s.Run(SrvAddr)
}

func setupDBClient(ctx context.Context, uri string) *mongo.Client {
	opts := options.Client().ApplyURI(uri)
	c, err := mongo.NewClient(opts)
	if err != nil {
		log.Fatalf("DB %v err: %v", DBAddr, err)
	}
	if err := c.Connect(ctx); err != nil {
		log.Fatalf("DB Connect:%v", err)
	}
	book.DB = c.Database("bookstore")
	book.Books = book.DB.Collection("books")
	return c
}

func setupWebServer() *gin.Engine {
	r := gin.New()
	gin.SetMode(gin.DebugMode)
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.LoadHTMLGlob(TMPL)
	r.Static("/css", CSS)
	r.Static("/js", JS)

	c := controller.New()
	r.NoRoute(c.NoRoute)
	r.GET("/", c.Index)
	r.GET("/books", c.Books)

	api := r.Group("api")
	v1 := api.Group("v1")
	v1.GET("/books", c.GetBooks)

	r.GET("/swagger/*any",
		ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
