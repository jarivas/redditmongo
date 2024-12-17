# Reddit Mongo
## Description
Scraps a particular subreddit and saves its posts to a mongodb

## Install
```go get github.com/jarivas/redditmongo```

## Usage
**.env**
```bash
REDDIT_USERNAME=reddit_bot
REDDIT_PASSWORD=snoo
REDDIT_CLIENT_ID=p-jcoLKBynTLew
REDDIT_APP_SECRET=gko_LXELoV07ZBNUXrvWZfzE3aI
MONGO_PORT=27017
MONGO_PORT_UI=8081
MONGO_DB_NAME=reddit
MONGO_INITDB_ROOT_USERNAME=root
MONGO_INITDB_ROOT_PASSWORD=example
```
**demo.go**
```golang
package demo

import (
	"github.com/jarivas/redditmongo"
    "log"
)

func main() {
	rm, err := RedditMongo{}.FromEnv("redditdev")

	if err != nil {
		log.Fatal(err)
	}

	e := make(chan error)

	go rm.Scrape(e)

	for err = range(e) {
		log.Fatal(err)
	}
}
```
