package repo

import (
	"context"
	"testing"

	"github.com/goagile/mongoshop/pkg/db"
	"go.mongodb.org/mongo-driver/mongo"
)

type Author struct {
	HRef string
	Name string
}

type PubOffice struct {
	HRef string
	Name string
}

type Cover struct {
	HRef  string
	Title string
}

type Book struct {
	HRef      string
	Price     int
	Discount  float64
	Title     string
	Cover     *Cover
	Authors   []*Author
	PubOffice *PubOffice
}

var (
	Waice        = &Author{Name: "Вайс", HRef: "."}
	Horton       = &Author{Name: "Хортон", HRef: "."}
	DMKPress     = &PubOffice{Name: "ДМК-Пресс", HRef: "."}
	ReactJSCover = &Cover{
		HRef:  "./img/razrabotka-web-react.jpg",
		Title: "Вайс, Хортон - Разработка веб-приложений в ReactJS",
	}
	ReactJSBook = &Book{
		HRef:      ".",
		Cover:     ReactJSCover,
		Price:     1800,
		Discount:  15,
		Title:     "Разработка веб-приложений в ReactJS",
		Authors:   []*Author{Waice, Horton},
		PubOffice: DMKPress,
	}
)

const TestDBAddr = "mongodb://127.0.0.1:27017"

var (
	Client *mongo.Client
)

func init() {
	ctx := context.Background()
	Client = db.NewClient(ctx, TestDBAddr)
}

func Test_X(t *testing.T) {
	// b := ReactJSBook
}
