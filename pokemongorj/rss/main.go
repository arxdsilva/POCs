package main

import (
	"net/url"

	"github.com/mmcdole/gofeed"
)

func main() {
	f, err := GetNewsFromRSS("https://google.com/somefeed")
	if err != nil {
		return
	}
	for _, item := range f.Items {
		// check github.com/arxdsilva/POCs/pokemongorj/tweeting
		if err = tweet(item.Title); err != nil {
			continue
		}
	}
}

// GetNewsFromRSS fetches a pokemon Go RSS feed by google
// and transforms it in data to be tweeted
func GetNewsFromRSS(feedURL string) (feed *gofeed.Feed, err error) {
	fp := gofeed.NewParser()
	feed, err = fp.ParseURL(feedURL)
	if err != nil {
		return
	}
	for _, item := range feed.Items {
		u, errP := url.Parse(item.Link)
		if errP != nil {
			return nil, errP
		}
		rq, errP := url.ParseQuery(u.RawQuery)
		if errP != nil {
			return nil, errP
		}
		item.Link = rq["url"][0]
	}
	return
}
