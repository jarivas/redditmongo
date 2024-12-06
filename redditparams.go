package redditmongo

import (
	"errors"
	"github.com/jarivas/redditscraper"
)

const redditMaxPosts int = 100
const redditWaitMilliseconds int = 3000

type RedditParams struct {
	subreddit        string
	maxPosts         int
	waitMilliseconds int
}

func (r RedditParams) New(subreddit string, maxPosts int, waitMilliseconds int) (*RedditParams, error) {
	if subreddit == "" {
		return nil, errors.New("empty subreddit")
	}

	if r.waitMilliseconds == 0 {
		waitMilliseconds = redditWaitMilliseconds
	}

	return &RedditParams{
		subreddit:        subreddit,
		maxPosts:         maxPosts,
		waitMilliseconds: waitMilliseconds,
	}, nil
}

func (r RedditParams) Default(subreddit string) (*RedditParams, error) {
	return r.New(subreddit, redditMaxPosts, redditWaitMilliseconds)
}

func (r RedditParams) validate() bool {
	return r.subreddit != "" && r.maxPosts != 0 && r.waitMilliseconds != 0
}

func (r RedditParams) getScraper() (*redditscraper.RedditScraper, error) {
	if !r.validate() {
		return nil, errors.New("invalid reddit params")
	}

	return redditscraper.RedditScraper{}.New(r.subreddit, r.maxPosts, r.waitMilliseconds)
}
