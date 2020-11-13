package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	addr                   = "127.0.0.1:8080"
	dburi                  = "mongodb://127.0.0.1:27017"
	dbtimeout              = 1 * time.Second
	DBClient               *mongo.Client
	DB                     *mongo.Database
	DBName                 = "shop"
	ProductsCollectionName = "products"
	Products               *mongo.Collection
)

func main() {
	// Database
	DBClient, err := mongo.NewClient(
		options.Client().ApplyURI(dburi))
	if err != nil {
		log.Fatalf("DB %v err: %v", dburi, err)
	}
	ctx, _ := context.WithTimeout(context.Background(), dbtimeout)
	if err := DBClient.Connect(ctx); err != nil {
		log.Fatalf("DB Connect:%v", err)
	}
	defer DBClient.Disconnect(ctx)
	DB = DBClient.Database(DBName)
	Products = DB.Collection(ProductsCollectionName)

	// HTTP Server
	http.Handle("/", http.FileServer(http.Dir("static")))
	http.HandleFunc("/product", product)
	log.Println("Server starts at:", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("ListenAndServe:%v", err)
	}
}

// GET http://host:port/product?slug=wheel-barrow-9092
func product(res http.ResponseWriter, req *http.Request) {
	if http.MethodGet != req.Method {
		res.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(res, "%v is not allowed.", req.Method)
		return
	}

	q := req.URL.Query()
	if len(q["slug"]) == 0 {
		res.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(res, "%v slug parameter must be specified", req.URL)
		return
	}

	slug := strings.TrimSpace(q["slug"][0])
	if "" == slug {
		res.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(res, "%v slug parameter must be not empty", req.URL)
		return
	}

	product, err := getProduct(slug)
	if err != nil {
		res.WriteHeader(http.StatusNotFound)
		log.Printf("%v getProduct %v", req.URL, err)
		fmt.Fprintf(res, "%v", err)
		return
	}

	log.Println("Found product", product)

	bytproduct, err := json.Marshal(product)
	if err != nil {
		log.Printf("Marshal err: %v", err)
		res.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(res, "err with product, slug %v", slug)
		return
	}

	res.WriteHeader(http.StatusOK)
	fmt.Fprint(res, bytproduct)
}

func getProduct(slug string) (map[string]interface{}, error) {
	r := Products.FindOne(context.TODO(), bson.M{
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
