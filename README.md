# Reddit Mongo
## Description
Scraps a particular and saves its posts to a mongodb

## Install
```go get https://github.com/jarivas/redditmongo```

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
	"github.com/jarivas/redditsmongo"
    "log"
)

func main() {
	err := RedditMongo("AmItheasshole", 100, 3000)

	if err != nil {
		log.Fatal(err)
	}
}
```

For more flexibility please check: