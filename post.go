package redditmongo

import (
	"errors"

	"github.com/jarivas/redditscraper"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
)

type Post struct {
	mgm.DefaultModel `bson:",inline"`
	Id               string `json:"id" bson:"id"`
	Title            string `json:"title" bson:"title"`
	Body             string `json:"body" bson:"body"`
}

func (p Post) FromScraped(post *redditscraper.Post) Post {
	return Post{
		Id:    post.Id,
		Title: post.Title,
		Body:  post.Body,
	}
}

func (p Post) Validate() bool {
	return p.Id != "" && p.Title != "" && p.Body != ""
}

func (p *Post) CheckExists() (bool, error) {
	if p.Id == "" {
		return false, errors.New("empty model")
	}

	post := &Post{}

	err := getCollection().First(
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

func (p Post) Save() error {

	if !p.Validate() {
		return errors.New("empty model")
	}

	exists, err := p.CheckExists()

	if err != nil {
		return err
	}

	if exists {
		return nil
	}

	return getCollection().Create(&p)
}
