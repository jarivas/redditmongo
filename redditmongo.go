package redditmongo

import (
	"errors"

	"github.com/jarivas/redditscraper"
)

type RedditMongo struct {
	ms *mongoStorage
}

func (rm RedditMongo) New(mp *MongoParams) (*RedditMongo, error) {
	if !mp.validate() {
		return nil, errors.New("invalid mongo params")
	}

	ms, err := mongoStorage{}.New(mp)

	if err != nil {
		return nil, err
	}

	return &RedditMongo{
		ms: ms,
	}, nil
}

func (rm RedditMongo) FromEnv() (*RedditMongo, error) {
	mp, err := MongoParams{}.FromEnv()

	if err != nil {
		return nil, err
	}

	return rm.New(mp)
}

func (rm *RedditMongo) Scrape(rp *RedditParams, s chan<- string, e chan<- error) {
	if !rp.validate() {
		e <- errors.New("invalid reddit params")
		return
	}

	scraper, err := rp.getScraper()

	if err != nil {
		e <- err
		return
	}

	subreddit := rp.subreddit
	c := make(chan *redditscraper.CachedPosts)

	go scraper.Scrape(c, e, rp.nextId)

	for posts := range c {
		rm.receivePosts(posts, subreddit, s, e)
	}
}

func (rm RedditMongo) receivePosts(posts *redditscraper.CachedPosts, subreddit string, s chan<- string, e chan<- error) {
	p := posts.GetPosts()

	if len(p) == 0 {
		e <- errors.New("empty posts from scraper")
	}

	for _, post := range p {
		p := Post{}.FromScraped(post, subreddit)

		err := p.Save(rm.ms)

		if err != nil {
			e <- err
		}

		s <- p.Id
	}
}
