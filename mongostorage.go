package redditmongo

import (
	"context"
	"errors"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoStorage struct {
	url    string
	dbName string
}

func (m mongoStorage) New(mp *MongoParams) (*mongoStorage, error) {
	new := mongoStorage{
		url:    mp.url,
		dbName: mp.dbName,
	}

	err := new.connect()

	if err != nil {
		return nil, err
	}

	return &new, nil
}

func (m mongoStorage) GetCollection(name string) *mgm.Collection {
	return mgm.CollectionByName(name)
}

func (m mongoStorage) CreateCollection(name string) error {
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

func (m mongoStorage) ResetColection(name string) error {
	err := m.GetCollection(name).Drop(context.TODO())

	if err != nil {
		return err
	}

	return m.CreateCollection(name)
}

func (m mongoStorage) connect() error {
	if m.url == "" || m.dbName == "" {
		return errors.New("impossible to connect, no data")
	}

	return mgm.SetDefaultConfig(nil, m.dbName, options.Client().ApplyURI(m.url))
}
