package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	SrvAddr                string
	DBURI                  string
	DBName                 string
	DBTimeout              = 10 * time.Second
	ProductsCollectionName = "products"
	DBClient               *mongo.Client
	DB                     *mongo.Database
	Products               *mongo.Collection
)

func main() {
	// Params
	flag.StringVar(&SrvAddr, "a", "127.0.0.1:8081", "server addr string, example: '127.0.0.1:8081'")
	flag.StringVar(&DBURI, "d", "127.0.0.1:27017", "mongo db server addr string, example: '127.0.0.1:27017'")
	flag.StringVar(&DBName, "n", "shop", "mongo db name, example: 'shop'")
	flag.Parse()

	// Database
	uri := fmt.Sprintf("mongodb://%v", DBURI)
	DBClient, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatalf("DB %v err: %v", DBURI, err)
	}

	ctx, fin := context.WithTimeout(context.Background(), DBTimeout)
	if err := DBClient.Connect(ctx); err != nil {
		log.Fatalf("DB Connect:%v", err)
	}
	defer func() {
		if err := DBClient.Disconnect(ctx); err != nil {
			log.Fatal(err)
		}
		fin()
	}()

	DB = DBClient.Database(DBName)
	Products = DB.Collection(ProductsCollectionName)

	// HTTP Server
	http.Handle("/", http.FileServer(http.Dir("static")))
	http.HandleFunc("/product", product)

	log.Println("Server starts at:", SrvAddr)
	if err := http.ListenAndServe(SrvAddr, nil); err != nil {
		log.Fatalf("ListenAndServe:%v", err)
	}
}

// GET http://host:port/product?slug=wheel-barrow-9092
func product(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "panic")
			log.Println("product handler", err)
			return
		}
	}()

	if http.MethodGet != r.Method {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "%v is not allowed.", r.Method)
		return
	}

	q := r.URL.Query()
	if len(q["slug"]) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%v slug parameter must be specified", r.URL)
		return
	}

	slug := strings.TrimSpace(q["slug"][0])
	if "" == slug {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%v slug parameter must be not empty", r.URL)
		return
	}

	product, err := getProduct(slug)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		log.Printf("%v getProduct %v", r.URL, err)
		fmt.Fprintf(w, "%v", err)
		return
	}

	log.Println("Found product", product)

	b, err := json.Marshal(product)
	if err != nil {
		log.Printf("Marshal err: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "err with product, slug %v", slug)
		return
	}

	w.WriteHeader(http.StatusOK)
	io.WriteString(w, string(b))
}

// getProduct - return one product
func getProduct(slug string) (map[string]interface{}, error) {
	ctx, fin := context.WithTimeout(context.Background(), DBTimeout)
	defer fin()
	r := Products.FindOne(ctx, bson.M{
		"slug": slug,
	})
	if err := r.Err(); err != nil {
		return nil, fmt.Errorf("FindOne: %v", err)
	}

	var p map[string]interface{}
	if err := r.Decode(&p); err != nil {
		return nil, fmt.Errorf("Decode:%v", err)
	}

	return p, nil
}
