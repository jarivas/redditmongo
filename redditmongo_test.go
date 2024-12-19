package redditmongo

import (
	"testing"
)

func TestScrape(t *testing.T) {
	rm, err := RedditMongo{}.FromEnv(testCollection)

	if err != nil {
		t.Error(err)
	}

	e := make(chan error)

	go rm.Scrape(e)

	for err = range(e) {
		t.Error(err)
	}
}