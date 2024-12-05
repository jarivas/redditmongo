package redditmongo

import "github.com/jarivas/redditscraper"

func main() {
	m, err := MongoStorage{}.FromEnv()
	
	if err != nil {
		
	}

}