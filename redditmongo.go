package redditmongo

import (
	"errors"

	"github.com/jarivas/redditscraper"
)

func RedditMongo(rp *RedditParams, mp *MongoParams) error {
	if !rp.validate() {
		return errors.New("invalid reddit params")
	}

	if !mp.validate(){
		return errors.New("invalid mongo params")
	}

	return redditMongoHelper(rp, mp)
}

func RedditMongoFromEnv(subreddit string) error {
	rp, err := RedditParams{}.Default(subreddit)

	if err != nil {
		return err
	}

	mp, err := MongoParams{}.FromEnv()

	if err != nil {
		return err
	}

	return redditMongoHelper(rp, mp)
}

func redditMongoHelper(rp *RedditParams, mp *MongoParams) error {
	scraper, err := rp.getScraper()

	if err != nil {
		return err
	}

	storage, err := mongoStorage{}.New(mp)

	if err != nil {
		return err
	}

	c := make(chan *redditscraper.CachedPosts)
	e := make(chan error)

	go scraper.Scrape(c, e)

	for {
		select {
			case posts := <- c: 
			for _, post := range(posts.GetPosts()) {
				p := Post{}.FromScraped(post, rp.subreddit)
	
				p.Save(storage)
			}
			case err := <- e:
				return err
		}
	}
}