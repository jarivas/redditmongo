package redditmongo

import (
	"errors"
	"log"
	"testing"
)

func TestRedditMongoScrape(t *testing.T) {
	m := getMongoStorageTest(t)

	err := m.ResetColection(testCollection)

	if err != nil {
		t.Error(err)
	}

	rm, err := RedditMongo{}.FromEnv()

	if err != nil {
		t.Error(err)
	}

	rp, err := RedditParams{}.Default(testCollection)

	if err != nil {
		t.Error(err)
	}

	s := make(chan string)
	e := make(chan error)

	go rm.Scrape(rp, s, e)

	for {
		select {
		case lastId := <-s:
			log.Println(lastId)

			if lastId == "" {
				t.Error(errors.New("invalid nextId"))
			}

			return
		case err = <-e:
			t.Error(err)
		}

	}
}
