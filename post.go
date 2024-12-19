package redditmongo

import (
	"errors"

	"github.com/jarivas/redditscraper"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Post struct {
	mgm.DefaultModel `bson:",inline"`
	Id               string `json:"id" bson:"id"`
	Title            string `json:"title" bson:"title"`
	Body             string `json:"body" bson:"body"`
	subreddit        string
}

func (p Post) FromScraped(post *redditscraper.Post, subreddit string) *Post {
	body := post.Body

	if len(body) == 0 {
		body = post.Title
	}
	return &Post{
		Id:        post.Id,
		Title:     post.Title,
		Body:      body,
		subreddit: subreddit,
	}
}

func (p Post) GetLast(m *MongoStorage, subreddit string) (*Post, error) {
	result := []Post{}
	var limit int64 = 1

	err := m.GetCollection(subreddit).SimpleFind(&result, bson.M{}, &options.FindOptions{
		Limit: &limit,
		Sort: bson.M{
			"_id": -1,
		},
	})

	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			return nil, nil
		}

		return nil, err
	}

	if len(result) == 0 {
		return nil, nil
	}

	return &result[0], nil
}

func (p *Post) Validate() bool {
	return p.Id != "" && p.Title != "" && p.Body != ""
}

func (p *Post) CheckExists(m *MongoStorage) (bool, error) {
	if p.Id == "" {
		return false, errors.New("empty model on check exists")
	}

	post := &Post{}

	err := m.GetCollection(p.subreddit).First(
		bson.M{
			"id": p.Id,
		}, post)

	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			return false, nil
		}

		return false, err
	}

	if post.Id == p.Id {
		return true, nil
	}

	return false, nil
}

func (p *Post) Save(m *MongoStorage) error {
	exists, err := p.CheckExists(m)

	if err != nil {
		return err
	}

	if exists {
		return nil
	}

	return m.GetCollection(p.subreddit).Create(p)
}
