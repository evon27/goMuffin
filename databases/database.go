package databases

import (
	"log"

	"git.wh64.net/muffin/goMuffin/configs"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func connect() *mongo.Client {
	client, err := mongo.Connect(options.Client().ApplyURI(configs.Config.DatabaseURL))
	if err != nil {
		log.Fatalln(err)
	}

	return client
}

var Client *mongo.Client = connect()
