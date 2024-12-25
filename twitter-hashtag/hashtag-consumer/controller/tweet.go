package controller

import (
	"database/sql"
	"fmt"
)

type Tweet struct {
	PostId           int    `json:"postId"`
	HashTag          string `json:"hashTag"`
	HashtagPartition int    `json:"hashtagPartition"`
}

func (tweet *Tweet) IncreamentHashtagCount(db *sql.DB) error {
	txn, err := db.Begin()
	if err != nil {
		txn.Rollback()
		return fmt.Errorf("could not begin transaction: %v", err)
	}

	_, err = db.Exec(`
		INSERT INTO tweets_count (hashtag, tweet_count)
		VALUES ($1, 1)
		ON CONFLICT (hashtag)  -- Conflict on the 'hashtag' column (PRIMARY KEY)
		DO UPDATE SET tweet_count = tweets_count.tweet_count + 1;
	`, tweet.HashTag)

	if err != nil {
		txn.Rollback()
		return fmt.Errorf("could not exec transaction: %v", err)
	}

	err = txn.Commit()
	if err != nil {
		return fmt.Errorf("could not commit transaction: %v", err)
	}
	return nil
}
