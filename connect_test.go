package redditmongo

import (
	"testing"
)

func TestConnectUsingEnv(t *testing.T) {
	err := ConnectUsingEnv()

	if err != nil {
		t.Error(err)
	}
}

func TestGetCollection(t *testing.T) {
	err := ConnectUsingEnv()

	if err != nil {
		t.Error(err)
	}

	collection := getCollection()

	if collection.Name() != mongoCollectionName {
		t.Error("Problems with collectons")
	}
}