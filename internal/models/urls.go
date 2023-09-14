package models

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Url struct {
	ID             primitive.ObjectID  `bson:"_id,omitempty"`
	ShortUrl       string              `bson:"shortURL"`
	OriginalURL    string              `bson:"originalURL"`
	CreationDate   time.Time           `bson:"creationDate"`
	ExpirationDate *time.Time          `bson:"expirationDate,omitempty"`
	HitCount       int                 `bson:"hitCount"`
	UserID         *primitive.ObjectID `bson:"userId,omitempty"`
}

type UrlModel struct {
	DB *mongo.Collection
}

func (m *UrlModel) FindById(id string) (*Url, error) {
	var result Url
	err := m.DB.FindOne(context.TODO(), bson.D{{Key: "_id", Value: id}}).Decode(&result)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, err
		}
		return nil, err
	}

	return &result, nil
}

// //title := "Back to the Future"
// 	var result bson.M
// 	err = coll.FindOne(context.TODO(), bson.D{{"title", title}}).Decode(&result)
// 	if err == mongo.ErrNoDocuments {
// 		fmt.Printf("No document was found with the title %s\n", title)
// 		return
// 	}
// 	if err != nil {
// 		panic(err)
// 	}
// 	jsonData, err := json.MarshalIndent(result, "", "    ")
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Printf("%s\n", jsonData)
