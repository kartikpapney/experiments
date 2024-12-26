-- -- Drop tables if they exist to start fresh
-- DROP TABLE IF EXISTS tweets_count;

-- Table: tweets_count
-- Stores information about tweets_count.
CREATE TABLE IF NOT EXISTS tweets_count (
    tweet_count INT,  -- Represents the count of tweets
    hashtag VARCHAR(255) PRIMARY KEY  -- Unique hashtag as the primary key
);
