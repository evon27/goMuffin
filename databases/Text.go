package databases

import (
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Text struct {
	Id        bson.ObjectID `bson:"_id"`
	Text      string
	Persona   string
	CreatedAt bson.DateTime `bson:"created_at"`
}

var Texts *mongo.Collection = Client.Database("muffin_ai_test").Collection("text")
