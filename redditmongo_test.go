package redditmongo

import (
	"testing"
)

func TestRedditMongo(t *testing.T) {
	err := RedditMongoFromEnv("AmItheasshole", "")

	if err != nil {
		t.Error(err)
	}
}