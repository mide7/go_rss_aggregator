package main

import (
	"time"

	"github.com/google/uuid"
	db "github.com/mide7/go_rss_aggregator/db/sqlc"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	ApiKey    string    `json:"api_key"`
}

func databaseUserToUser(user db.User) User {
	return User{
		ID:        user.ID,
		Name:      user.Name,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		ApiKey:    user.ApiKey,
	}
}

func databaseUsersToUsers(users []db.User) []User {
	var userArray []User
	for _, user := range users {
		userArray = append(userArray, databaseUserToUser(user))
	}
	return userArray
}

type Feed struct {
	ID        uuid.UUID `json:"id"`
	Title     string    `json:"title"`
	Url       string    `json:"url"`
	UserID    uuid.UUID `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func databaseFeedToFeed(feed db.Feed) Feed {
	return Feed{
		ID:        feed.ID,
		Title:     feed.Title,
		Url:       feed.Url,
		UserID:    feed.UserID,
		CreatedAt: feed.CreatedAt,
		UpdatedAt: feed.UpdatedAt,
	}
}

func databaseFeedsToFeeds(feeds []db.Feed) []Feed {
	var feedArray []Feed
	for _, feed := range feeds {
		feedArray = append(feedArray, databaseFeedToFeed(feed))
	}
	return feedArray
}

type FeedFollow struct {
	ID        uuid.UUID `json:"id"`
	UserID    uuid.UUID `json:"user_id"`
	FeedID    uuid.UUID `json:"feed_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ListFeedFollowsRow struct {
	ID        uuid.UUID `json:"id"`
	UserID    uuid.UUID `json:"user_id"`
	FeedID    uuid.UUID `json:"feed_id"`
	User      any       `json:"user"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func databaseFeedFollowToFeedFollow(feedFollow db.FeedFollow) FeedFollow {
	return FeedFollow{
		ID:        feedFollow.ID,
		UserID:    feedFollow.UserID,
		FeedID:    feedFollow.FeedID,
		CreatedAt: feedFollow.CreatedAt,
		UpdatedAt: feedFollow.UpdatedAt,
	}
}

func databaseFeedFollowsToFeedFollows(feedFollows []db.FeedFollow) []FeedFollow {
	var feedFollowArray []FeedFollow
	for _, feedFollow := range feedFollows {
		feedFollowArray = append(feedFollowArray, databaseFeedFollowToFeedFollow(feedFollow))
	}
	return feedFollowArray
}

func databaseListFeedFollowsToFeedFollows(feedFollows []db.ListFeedFollowsRow) []ListFeedFollowsRow {
	var feedFollowArray []ListFeedFollowsRow
	for _, feedFollow := range feedFollows {
		feedFollowArray = append(feedFollowArray, ListFeedFollowsRow{
			ID:        feedFollow.ID,
			UserID:    feedFollow.UserID,
			FeedID:    feedFollow.FeedID,
			User:      feedFollow.User,
			CreatedAt: feedFollow.CreatedAt,
			UpdatedAt: feedFollow.UpdatedAt,
		})
	}
	return feedFollowArray
}

type Post struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Url         string    `json:"url"`
	Description string    `json:"description"`
	FeedID      uuid.UUID `json:"feed_id"`
	PublishedAt time.Time `json:"published_at"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func databasePostToPost(post db.Post) Post {
	var description string
	if post.Description.Valid {
		description = post.Description.String
	}
	return Post{
		ID:          post.ID,
		Title:       post.Title,
		Url:         post.Url,
		Description: description,
		FeedID:      post.FeedID,
		PublishedAt: post.PublishedAt,
		CreatedAt:   post.CreatedAt,
		UpdatedAt:   post.UpdatedAt,
	}
}

func databasePostsToPosts(posts []db.Post) []Post {
	var postArray []Post
	for _, post := range posts {
		postArray = append(postArray, databasePostToPost(post))
	}
	return postArray
}
