package redditmongo

import (
	"errors"
	"os"
)

type MongoParams struct {
	url    string
	dbName string
}

func (m MongoParams) New(url, dbName string) (*MongoParams, error) {
	if url == "" || dbName == "" {
		return nil, errors.New("empty mongo params")
	}

	return &MongoParams{
		url:    url,
		dbName: dbName,
	}, nil
}

func (m MongoParams) FromEnv() (*MongoParams, error) {
	url := os.Getenv("MONGO_CONNECTION_STRING")
	dbName := os.Getenv("MONGO_DB_NAME")

	return m.New(url, dbName)
}

func (m MongoParams) validate() bool {
	return m.url != "" && m.dbName != ""
}
