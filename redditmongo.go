package redditmongo

import (
	"errors"

	"github.com/jarivas/redditscraper"
)

type RedditMongo struct {
	ms *MongoStorage
	rs *redditscraper.RedditScraper
}

func (rm RedditMongo) New(ms *MongoStorage, rs *redditscraper.RedditScraper) (*RedditMongo, error) {
	return &RedditMongo{
		ms: ms,
		rs: rs,
	}, nil
}

func (rm *RedditMongo) Scrape(nextId string, s chan<- string, e chan<- error) {
	c := make(chan *redditscraper.CachedPosts)

	go rm.rs.Scrape(c, e, nextId)

	for posts := range c {
		rm.receivePosts(posts, s, e)
	}
}

func (rm *RedditMongo) receivePosts(posts *redditscraper.CachedPosts, s chan<- string, e chan<- error) {
	p := posts.GetPosts()

	if len(p) == 0 {
		e <- errors.New("empty posts from scraper")
	}

	for _, post := range p {
		p := Post{}.FromScraped(post, rm.rs.GetSubreddit())

		err := p.Save(rm.ms)

		if err != nil {
			e <- err
		}

		s <- p.Id
	}
}
