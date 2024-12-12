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

func (rm RedditMongo) FromEnv()(*RedditMongo, error) {
	mp, err := MongoParams{}.FromEnv()

	if err != nil {
		return nil, err
	}

	return rm.New(mp)
}

func (rm *RedditMongo) Scrape(rp *RedditParams, s chan <-string) error {
	if !rp.validate() {
		return errors.New("invalid reddit params")
	}

	scraper, err := rp.getScraper()

	if err != nil {
		return err
	}

	subreddit := rp.subreddit
	c := make(chan *redditscraper.CachedPosts)
	e := make(chan error)

	go scraper.Scrape(c, e, rp.nextId)

	for {
		select {
		case posts := <-c:
			err = rm.receivePosts(posts, subreddit, s)

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

func (rm RedditMongo) receivePosts(posts *redditscraper.CachedPosts, subreddit string, s chan <-string) error {
	p := posts.GetPosts()

	if len(p) == 0 {
		return errors.New("empty posts from scraper")
	}

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