package redditmongo

import (
	"testing"
)

func TestFromEnv(t *testing.T) {
	_, err := MongoStorage{}.FromEnv()

	if err != nil {
		t.Error(err)
	}
}