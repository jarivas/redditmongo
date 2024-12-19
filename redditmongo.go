package redditmongo

import (
	"github.com/jarivas/redditscraper"
)

type RedditMongo struct {
	ms *MongoStorage
	rs *redditscraper.RedditScraper
	Subreddit string
}

func (rm RedditMongo) New(ms *MongoStorage, rs *redditscraper.RedditScraper, subreddit string) *RedditMongo {
	return &RedditMongo{
		ms: ms,
		rs: rs,
		Subreddit: subreddit,
	}
}

func (rm RedditMongo) FromEnv(subreddit string) (*RedditMongo, error) {
	ms, err := MongoStorage{}.FromEnv()

	if err != nil {
		return nil, err
	}

	rs, err := redditscraper.RedditScraper{}.FromEnv(subreddit)

	if err != nil {
		return nil, err
	}
	
	return RedditMongo{}.New(ms, rs, subreddit), nil
}

func (rm *RedditMongo) Scrape(e chan<- error) {
	var err error

	p := make(chan *redditscraper.Post)

	listing := redditscraper.PostListing{
		Limit: redditscraper.MaxPosts,
		Id: rm.getLastId(),
	}

	go rm.rs.Listen(redditscraper.SubredditTop, listing, p, e)

	for post := range p {
		p := Post{}.FromScraped(post, rm.Subreddit)

		if p.Validate() {
			err = p.Save(rm.ms)
		}

		if err != nil {
			e <- err
			err = nil
		}
	}
}

func (rm *RedditMongo) getLastId() string {
	p, err := Post{}.GetLast(rm.ms, rm.Subreddit)

	if err != nil || p == nil {
		return ""
	}

	return p.Id
}