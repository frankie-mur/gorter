package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/a-h/templ"
	"github.com/frankie-mur/gorter/internal/models"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type application struct {
	urls  *models.UrlModel
	templ templ.Component
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

	r := chi.NewRouter()

	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))

	coll := client.Database("GorterDB").Collection("gorter")

	component := hello("Frankie")

	app := &application{
		urls:  &models.UrlModel{DB: coll},
		templ: component,
	}

	//render html template

	r.Get("/home", app.HomePage)

	r.Get("/shorten/*", app.urlFind)
	r.Post("/url/create", app.urlCreate)

	http.ListenAndServe(":4000", r)

}

func openDB(uri string) (*mongo.Client, error) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	return client, nil
}
