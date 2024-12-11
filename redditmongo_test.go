package redditmongo

import (
	"errors"
	"testing"
)

func TestRedditMongoScrape(t *testing.T) {
	m := getMongoStorageTest(t)

	err := m.ResetColection(testCollection)

	if err != nil {
		t.Error(err)
	}

	rp, err := RedditParams{}.Default(testCollection)

	if err != nil {
		t.Error(err)
	}

	rm, err := RedditMongo{}.FromEnv(rp)

	if err != nil {
		t.Error(err)
	}

	s := make(chan string)

	go func() {
		err = rm.Scrape(s)

		if err != nil {
			t.Error(err)
		}
	}()

	nextId := <- s

	if nextId == "" {
		t.Error(errors.New("invalid nextId"))
	}
}