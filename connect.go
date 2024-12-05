package redditmongo

import (
	"context"
	"errors"
	"os"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect(url, dbName string) error {
	if url == "" || dbName != "" {
		return errors.New("impossible to connect, no data")
	}

	return connectAndCreateDb(dbName, url)
}

func ConnectUsingEnv() error {
	dbName := os.Getenv("MONGO_DB_NAME")
	url := os.Getenv("MONGO_CONNECTION_STRING")

	return connectAndCreateDb(dbName, url)
}

func getCollection() *mgm.Collection {
	return mgm.CollectionByName(mongoCollectionName)
}

func connectAndCreateDb(dbName, url string) error {
	err := mgm.SetDefaultConfig(nil, dbName, options.Client().ApplyURI(url))

	if err != nil {
		return err
	}

	_, client, _, err := mgm.DefaultConfigs()

	if err != nil {
		return err
	}

	c := client.Database(dbName).Collection(mongoCollectionName)
	i := mongo.IndexModel{
		Keys: bson.D{{Key: "Id", Value: 1}},
	}
	_, err = c.Indexes().CreateOne(context.TODO(), i)

	return err
}
