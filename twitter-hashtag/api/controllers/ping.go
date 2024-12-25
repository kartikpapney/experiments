package controllers

import (
	"database/sql"
	"net/http"
	responseModel "twitterApi/models/response"
	"twitterApi/services"

	"github.com/gin-gonic/gin"
)

// Ping godoc
// @Summary      Ping
// @Description  check server
// @Tags         ping
// @Accept       json
// @Produce      json
// @Success      200  {object}  responseModel.Response
// @Router       /ping [get]
func Ping(c *gin.Context) {
	response := &responseModel.Response{
		StatusCode: http.StatusOK,
		Success:    true,
		Message:    "Working!",
	}

	response.SendResponse(c)
}

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

	tweets, err := services.GetTopTweets(postgres)

	if err != nil {
		response.Message = err.Error()
		response.SendResponse(c)
		return
	}

	response.StatusCode = http.StatusOK
	response.Success = true
	response.Data = gin.H{
		"trending": tweets,
	}
	response.SendResponse(c)
}
