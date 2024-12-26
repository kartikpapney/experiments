package controllers

import (
	"net/http"
	responseModel "twitterApi/models/response"

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
