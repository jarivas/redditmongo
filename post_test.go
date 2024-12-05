package redditmongo

import (
	"testing"

	"github.com/jarivas/redditscraper"
)

func TestFromScraped(t *testing.T) {
	rp := &redditscraper.Post{
		Id: "asd",
		Title: "asd",
		Body: "asd",
	}

	p := Post{}.FromScraped(rp)

	if p.Id != "asd" {
		t.Errorf("Wrong id: %v", p.Id)
	}

	if p.Title != "asd" {
		t.Errorf("Wrong title: %v", p.Title)
	}

	if p.Body != "asd" {
		t.Errorf("Wrong body: %v", p.Body)
	}
}

func TestValidate(t *testing.T) {
	rp := &redditscraper.Post{
		Id: "asd",
		Title: "asd",
		Body: "asd",
	}

	p := Post{}.FromScraped(rp)

	if !p.Validate() {
		t.Error("Problem with validation")
	}
}

func TestCheckExists(t *testing.T) {
	err := ConnectUsingEnv()

	if err != nil {
		t.Error(err)
	}

	rp := &redditscraper.Post{
		Id: "asd123213",
		Title: "asd",
		Body: "asd",
	}

	p := Post{}.FromScraped(rp)

	err = getCollection().Delete(&p)

	if err != nil {
		t.Error(err)
	}

	exists, err := p.CheckExists()

	if err != nil {
		t.Error(err)
	}

	if exists {
		t.Error("exists when it should not")
	}
}

func TestSave(t *testing.T) {
	err := ConnectUsingEnv()

	if err != nil {
		t.Error(err)
	}

	rp := &redditscraper.Post{
		Id: "asd",
		Title: "asd",
		Body: "asd",
	}

	p := Post{}.FromScraped(rp)

	err = getCollection().Delete(&p)

	if err != nil {
		t.Error(err)
	}

	err = p.Save()

	if err != nil {
		t.Error(err)
	}
	
	exists, err := p.CheckExists()

	if err != nil {
		t.Error(err)
	}

	if !exists {
		t.Error("post not created")
	}
}