package services

import (
	"database/sql"
	"fmt"
	models "twitterApi/models/db"
)

func GetTopTweets(db *sql.DB) ([]models.Tweet, error) {

	rows, err := db.Query("SELECT hashtag, tweet_count FROM tweets_count ORDER BY tweet_count DESC LIMIT 10")
	if err != nil {
		return nil, fmt.Errorf("could not execute query: %v", err)
	}
	defer rows.Close()

	var topTweets []models.Tweet

	for rows.Next() {
		var tweet models.Tweet
		if err := rows.Scan(&tweet.HashTag, &tweet.TweetCount); err != nil {
			return nil, fmt.Errorf("could not scan row: %v", err)
		}
		topTweets = append(topTweets, tweet)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over rows: %v", err)
	}

	return topTweets, nil
}
