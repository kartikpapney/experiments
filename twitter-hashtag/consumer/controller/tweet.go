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

func IncreamentHashtagCount(db *sql.DB, hashtag string, count int) error {
	txn, err := db.Begin()
	// defer txn.Rollback()
	if err != nil {

		return fmt.Errorf("could not begin transaction: %v", err)
	}

	_, err = txn.Exec(`
		INSERT INTO tweets_count (hashtag, tweet_count)
		VALUES ($1, 1)
		ON CONFLICT (hashtag)  -- Conflict on the 'hashtag' column (PRIMARY KEY)
		DO UPDATE SET tweet_count = tweets_count.tweet_count + $2;
	`, hashtag, count)

	if err != nil {
		return fmt.Errorf("could not exec transaction: %v", err)
	}

	err = txn.Commit()

	if err != nil {
		return fmt.Errorf("could not commit transaction: %v", err)
	}
	fmt.Printf("#%s is increased by %d\n", hashtag, count)
	return nil
}
