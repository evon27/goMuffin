package databases

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type InsertText struct {
	Text      string
	Persona   string
	CreatedAt time.Time `bson:"created_at"`
}

type Text struct {
	Id        bson.ObjectID `bson:"_id"`
	Text      string
	Persona   string
	CreatedAt time.Time `bson:"created_at"`
}

var Texts *mongo.Collection = Client.Database("muffin_ai_test").Collection("text")
