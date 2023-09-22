package models

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Url struct {
	ID             primitive.ObjectID  `bson:"_id,omitempty"`
	ShortURL       string              `bson:"shortURL"`
	OriginalURL    string              `bson:"originalURL"`
	CreationDate   time.Time           `bson:"creationDate"`
	ExpirationDate time.Time           `bson:"expirationDate,omitempty"`
	HitCount       int                 `bson:"hitCount"`
	UserID         *primitive.ObjectID `bson:"userId,omitempty"`
}

type UrlModel struct {
	DB *mongo.Collection
}

func (m *UrlModel) FindOriginalUrl(shortURL string) (string, error) {
	var result Url
	fmt.Printf("Looking for matching with short URL: %s\n", shortURL)
	err := m.DB.FindOne(context.TODO(), bson.D{{Key: "shortURL", Value: shortURL}}).Decode(&result)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			log.Fatal(err)
			return "", err
		}
		log.Fatal(err)
		return "", err
	}

	return result.OriginalURL, nil
}

func (m *UrlModel) CreateUrl(shortUrl string, originalUrl string) error {
	url := Url{
		ShortURL:       shortUrl,
		OriginalURL:    originalUrl,
		CreationDate:   time.Now(),
		ExpirationDate: time.Now().Add(time.Hour),
		HitCount:       0,
	}
	_, err := m.DB.InsertOne(context.TODO(), url)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("Successfully created url: ", url)
	return nil
}
