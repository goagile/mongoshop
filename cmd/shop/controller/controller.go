package controller

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/goagile/mongoshop/pkg/book"
)

type Controller struct{}

func New() *Controller {
	return new(Controller)
}

// GetBooks godoc
// @Summary Show all books
// @Tags books
// @Produce json
// @Success 200 {array} map[string]interface{}
// @Router /books [get]
func (ctrl *Controller) GetBooks(c *gin.Context) {
	ctx := context.Background()
	ctx, _ = context.WithTimeout(ctx, 10*time.Second)
	bs, err := book.FindAll(ctx)
	if err != nil {
		log.Println("FindAll", err)
		c.String(http.StatusNotFound, "books are not found")
		return
	}
	c.JSON(http.StatusOK, gin.H{"items": bs})
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
