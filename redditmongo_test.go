package redditmongo

import (
	"log"
	"testing"

	"github.com/jarivas/redditscraper"
)

func TestNewRedditMongo(t *testing.T) {
	getNewRedditMongo(t)
}

func TestScrape(t *testing.T) {
	rm := getNewRedditMongo(t)

	s := make(chan string)
	e := make(chan error)

	go rm.Scrape(s, e)

	for {
		select {
		case lastId := <-s:
			log.Println(lastId)
			return
		case err := <-e:
			log.Fatal(err)
		}
	}
}

func getNewRedditMongo(t *testing.T) *RedditMongo{
	ms, err := MongoStorage{}.FromEnv()

	if err != nil {
		t.Error(err)
	}

	rs, err := redditscraper.RedditScraper{}.FromEnv("AmItheasshole")

	if err != nil {
		t.Error(err)
	}
	
	rm, err := RedditMongo{}.New(ms, rs)

	if err != nil {
		t.Error(err)
	}

	if rm == nil {
		t.Error("rm is nil")
	}

	return rm
}