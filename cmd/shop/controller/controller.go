package controller

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/goagile/mongoshop/cmd/shop/db"
)

type Controller struct{}

func New() *Controller {
	return new(Controller)
}

// GetBooks godoc
// @Summary Show all books
// @Tags books
// @Produce json
// @Success 200 {array} db.Book
// @Router /books [get]
func (ctrl *Controller) GetBooks(c *gin.Context) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	books, err := db.FindBooks(ctx)
	if err != nil {
		log.Println("FindBooks", err)
		c.String(http.StatusNotFound, "books are not found")
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": books})
}

func (ctrl *Controller) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func (ctrl *Controller) Books(c *gin.Context) {
	c.HTML(http.StatusOK, "books.html", nil)
}

func (ctrl *Controller) NoRoute(c *gin.Context) {
	c.HTML(http.StatusOK, "notfound.html", nil)
}
