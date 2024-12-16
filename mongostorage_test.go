package redditmongo

import (
	"testing"
)

func TestNewMongoStorage(t *testing.T) {
	rm, err := MongoStorage{}.FromEnv()

	if err != nil {
		t.Error(err)
	}

	if rm == nil {
		t.Error("rm is nil")
	}
}

func TestGetCollection(t *testing.T) {
	rm, err := MongoStorage{}.FromEnv()

	if err != nil {
		t.Error(err)
	}

	c := "test"
	err = rm.CreateCollection(c)

	if err != nil {
		t.Error(err)
	}

	col := rm.GetCollection(c)

	if col == nil {
		t.Error("Collection is nil")
	}
}