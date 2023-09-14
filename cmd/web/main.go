package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/frankie-mur/gorter/internal/models"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type application struct {
	urls *models.UrlModel
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environment variable")
	}

	client, err := openDB(uri)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	srv := http.NewServeMux()

	coll := client.Database("GorterDB").Collection("gorter")

	app := &application{
		urls: &models.UrlModel{DB: coll},
	}
	//TODO: WHY ISNT THIS WORKING
	srv.HandleFunc("/shorten", app.urlCreate)

	http.ListenAndServe(":4000", srv)

}

func openDB(uri string) (*mongo.Client, error) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	return client, nil
}
