package main

import (
	"context"
	"database/sql"
	"log"
	"strings"
	"sync"
	"time"

	db "github.com/mide7/go/rss_aggregator/db/sqlc"
)

func startScraping(database *db.Queries, concurrency int, timeBetweenRequest time.Duration) {
	log.Printf("Starting scraping on %v concurrency and %v seconds between requests", concurrency, timeBetweenRequest)

	ticker := time.NewTicker(timeBetweenRequest)

	for ; ; <-ticker.C {
		feeds, err := database.GetNextFeedsToFetch(context.Background(), int32(concurrency))
		if err != nil {
			log.Printf("Error fetching feeds: %v", err)
			continue
		}

		wg := &sync.WaitGroup{}

		for _, feed := range feeds {
			wg.Add(1)

			go scrapeFeed(database, feed, wg)
		}

		wg.Wait()
	}
}

func scrapeFeed(database *db.Queries, feed db.Feed, wg *sync.WaitGroup) {
	defer wg.Done()

	_, err := database.MarkFeedAsFetched(context.Background(), feed.ID)

	if err != nil {
		log.Printf("Error marking feed %v as fetched: %v", feed.ID, err)
		return
	}

	rssFeed, err := urlToFeed(feed.Url)
	if err != nil {
		log.Printf("Error fetching feed %v: %v", feed.ID, err)
		return
	}

	for _, item := range rssFeed.Channel.Items {
		pubDate, err := time.Parse(time.RFC1123Z, item.PubDate)
		if err != nil {
			log.Printf("Error parsing date %v: %v", item.PubDate, err)
			continue
		}

		log.Println("Found post", item.Title, "for feed", feed.Title)
		_, err = database.CreatePost(context.Background(), db.CreatePostParams{
			Title:       item.Title,
			Url:         item.Link,
			Description: sql.NullString{String: item.Description, Valid: item.Description != ""},
			FeedID:      feed.ID,
			PublishedAt: pubDate,
		})

		if err != nil {
			if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
				continue
			}
			log.Printf("Error creating post: %v", err)
		}
	}

	log.Printf("Feed %s collected, %v posts found", feed.Title, len(rssFeed.Channel.Items))
}
