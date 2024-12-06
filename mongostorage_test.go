package redditmongo

import (
	"testing"
)

func TestFromEnv(t *testing.T) {
	mp, err := MongoParams{}.FromEnv()

	if err != nil {
		t.Error(err)
	}

	_, err = mongoStorage{}.New(mp)

	if err != nil {
		t.Error(err)
	}
}
