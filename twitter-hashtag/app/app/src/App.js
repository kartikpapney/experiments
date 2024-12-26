import React, { useEffect, useState } from "react";

const TrendingTweets = () => {
  const [data, setData] = useState(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  useEffect(() => {
    let intervalId;

    const fetchData = async () => {
      try {
        const response = await fetch("http://localhost:3000/api/v1/trending");
        if (!response.ok) {
          throw new Error("Failed to fetch data");
        }
        const result = await response.json();
        setData(result.data);
        setLoading(false);
      } catch (err) {
        setError(err.message);
        setLoading(false);
      }
    };

    fetchData();

    intervalId = setInterval(() => {
      fetchData();
    }, 1000);

    return () => clearInterval(intervalId);
  }, []);

  if (loading) {
    return <div>Loading...</div>;
  }

  if (error) {
    return <div>Error: {error}</div>;
  }

  return (
    <div style={{margin: "10px"}}>
      <h1>Trending Tweets</h1>
      <p>Total Tweets: {data.totalTweets}</p>
      <ul>
        {data.trending.map((trend, index) => (
          <li key={index}>
            <strong>#{trend.HashTag}</strong>: {trend.TweetCount} tweets
          </li>
        ))}
      </ul>
    </div>
  );
};

export default TrendingTweets;
