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

type MongoStorage struct {
	url    string
	dbName string
}

func (m MongoStorage) New(url, dbName string) (*MongoStorage, error) {
	err := m.connect(url, dbName)

	if err != nil {
		return nil, err
	}

	new := MongoStorage{
		url:    url,
		dbName: dbName,
	}

	return &new, nil
}

func (m MongoStorage) FromEnv() (*MongoStorage, error) {
	url := os.Getenv("MONGO_CONNECTION_STRING")
	dbName := os.Getenv("MONGO_DB_NAME")

	return m.New(url, dbName)
}

func (m MongoStorage) GetCollection(name string) *mgm.Collection {
	return mgm.CollectionByName(name)
}

func (m MongoStorage) CreateCollection(name string) error {
	_, client, _, err := mgm.DefaultConfigs()

	if err != nil {
		return err
	}

	c := client.Database(m.dbName).Collection(name)
	i := mongo.IndexModel{
		Keys: bson.D{
			{Key: "Id", Value: 1},
		},
	}
	_, err = c.Indexes().CreateOne(context.TODO(), i)

	return err
}

func (m MongoStorage) connect(url, dbName string) error {
	if url == "" || dbName != "" {
		return errors.New("impossible to connect, no data")
	}

	return mgm.SetDefaultConfig(nil, dbName, options.Client().ApplyURI(url))
}
