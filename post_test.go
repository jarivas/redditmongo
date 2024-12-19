package redditmongo

import (
	"testing"

	"github.com/jarivas/redditscraper"
)

const testCollection string = "redditdev"

var m *MongoStorage

func TestFromScraped(t *testing.T) {
	rp := &redditscraper.Post{
		Id:    "asd",
		Title: "asd",
		Body:  "asd",
	}

	p := Post{}.FromScraped(rp, testCollection)

	if p.Id != "asd" {
		t.Errorf("Wrong id: %v", p.Id)
	}

	if p.Title != "asd" {
		t.Errorf("Wrong title: %v", p.Title)
	}

	if p.Body != "asd" {
		t.Errorf("Wrong body: %v", p.Body)
	}

	if p.subreddit != testCollection {
		t.Errorf("Wrong subreddit: %v", p.Body)
	}
}

func TestValidate(t *testing.T) {
	rp := &redditscraper.Post{
		Id:    "asd",
		Title: "asd",
		Body:  "asd",
	}

	p := Post{}.FromScraped(rp, testCollection)

	if !p.Validate() {
		t.Error("Problem with validation")
	}
}

func TestGetLast(t *testing.T) {
	TestSave(t)

	post, err := Post{}.GetLast(m, testCollection)

	if err != nil {
		t.Error(err)
	}

	if post == nil {
		t.Error("post is nil")
	}
}

func TestCheckExists(t *testing.T) {
	setMongoStorageTest(t)
	err := m.ResetColection(testCollection)

	if err != nil {
		t.Error(err)
	}

	rp := &redditscraper.Post{
		Id:    "asd123213",
		Title: "asd",
		Body:  "asd",
	}

	p := Post{}.FromScraped(rp, testCollection)

	exists, err := p.CheckExists(m)

	if err != nil {
		t.Error(err)
	}

	if exists {
		t.Error("exists when it should not")
	}
}

func TestSave(t *testing.T) {
	setMongoStorageTest(t)

	err := m.ResetColection(testCollection)

	if err != nil {
		t.Error(err)
	}

	rp := &redditscraper.Post{
		Id:    "asd",
		Title: "asd",
		Body:  "asd",
	}

	p := Post{}.FromScraped(rp, testCollection)

	if err != nil {
		t.Error(err)
	}

	err = p.Save(m)

	if err != nil {
		t.Error(err)
	}

	exists, err := p.CheckExists(m)

	if err != nil {
		t.Error(err)
	}

	if !exists {
		t.Error("post not created")
	}
}

func setMongoStorageTest(t *testing.T) {
	if m != nil {
		return
	}

	var err error

	m, err = MongoStorage{}.FromEnv()

	if err != nil {
		t.Error(err)
	}
}
