package redditmongo

import (
	"errors"

	"github.com/jarivas/redditscraper"
)

type RedditMongo struct {
	rp *RedditParams
	ms *mongoStorage
}

func (rm RedditMongo) New(rp *RedditParams, mp *MongoParams) (*RedditMongo, error) {
	if !rp.validate() {
		return nil, errors.New("invalid reddit params")
	}

	if !mp.validate() {
		return nil, errors.New("invalid mongo params")
	}

	return getRedditMongoHelper(rp, mp)
}

func (rm RedditMongo) FromEnv(rp *RedditParams) (*RedditMongo, error) {
	mp, err := MongoParams{}.FromEnv()

	if err != nil {
		return nil, err
	}

	return getRedditMongoHelper(rp, mp)
}

func (rm RedditMongo) Scrape(s chan <-string) error {
	scraper, err := rm.rp.getScraper()

	if err != nil {
		return err
	}

	c := make(chan *redditscraper.CachedPosts)
	e := make(chan error)

	go scraper.Scrape(c, e, rm.rp.nextId)

	for {
		select {
		case posts := <-c:
			err = rm.receivePosts(posts, s)

			if err != nil {
				close(c)
				close(e)
				return err
			}
		case err = <-e:
			close(c)
			close(e)
			return err
		}
	}
}

func (rm RedditMongo) receivePosts(posts *redditscraper.CachedPosts, s chan <-string) error {
	p := posts.GetPosts()

	if len(p) == 0 {
		return errors.New("empty posts from scraper")
	}
	subreddit := rm.rp.subreddit

	for _, post := range p {
		p := Post{}.FromScraped(post, subreddit)

		err := p.Save(rm.ms)

		if err != nil {
			return err
		}

		s <- p.Id
	}

	return nil
}

func getRedditMongoHelper(rp *RedditParams, mp *MongoParams) (*RedditMongo, error) {
	ms, err := mongoStorage{}.New(mp)

	if err != nil {
		return nil, err
	}

	return &RedditMongo{
		rp: rp,
		ms: ms,
	}, nil
}
