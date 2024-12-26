package services

import (
	"database/sql"
	"fmt"
	models "twitterApi/models/db"
)

func GetTopTweets(db *sql.DB) ([]models.Tweet, int, error) {

	rows, _ := db.Query("SELECT hashtag, tweet_count FROM tweets_count ORDER BY tweet_count DESC LIMIT 10")
	var totalTweetCount int
	err := db.QueryRow("SELECT SUM(tweet_count) FROM tweets_count").Scan(&totalTweetCount)
	if err != nil {
		return nil, -1, fmt.Errorf("could not execute query: %v", err)
	}
	defer rows.Close()

	var topTweets []models.Tweet

	for rows.Next() {
		var tweet models.Tweet
		tweet.HashTag = "#" + tweet.HashTag
		if err := rows.Scan(&tweet.HashTag, &tweet.TweetCount); err != nil {
			return nil, -1, fmt.Errorf("could not scan row: %v", err)
		}
		topTweets = append(topTweets, tweet)
	}

	if err := rows.Err(); err != nil {
		return nil, -1, fmt.Errorf("error iterating over rows: %v", err)
	}

	return topTweets, totalTweetCount, nil
}
