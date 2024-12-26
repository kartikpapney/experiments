package controllers

import (
	"database/sql"
	"net/http"
	responseModel "twitterApi/models/response"
	"twitterApi/services"

	"github.com/gin-gonic/gin"
)

func GetTopTweets(c *gin.Context) {
	response := &responseModel.Response{
		StatusCode: http.StatusBadRequest,
		Success:    false,
	}

	db, exists := c.Get("db")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "DB connection not found"})
		return
	}

	postgres, ok := db.(*sql.DB)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid DB connection"})
		return
	}

	tweets, totalTweets, err := services.GetTopTweets(postgres)

	if err != nil {
		response.Message = err.Error()
		response.SendResponse(c)
		return
	}

	response.StatusCode = http.StatusOK
	response.Success = true
	response.Data = gin.H{
		"totalTweets": totalTweets,
		"trending":    tweets,
	}

	response.SendResponse(c)
}
