package main

import (
	"context"

	"github.com/gin-gonic/gin"
	_ "github.com/goagile/mongoshop/api/docs"
	"github.com/goagile/mongoshop/cmd/shop/controller"
	"github.com/goagile/mongoshop/pkg/db"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

const (
	DBURI  = "mongodb://127.0.0.1:27017"
	SRVURI = ":8080"
	TMPL   = "./web/templates/*"
	CSS    = "./web/static/css"
	JS     = "./web/static/js"
)

// @title Hello
// @BasePath /api/v1
func main() {
	ctx := context.Background()
	db.Connect(ctx, DBURI)
	defer db.Disconnect(ctx)

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

	r.Run(SRVURI)
}
