package routes

import (
	"twitterApi/controllers"

	"github.com/gin-gonic/gin"
)

func PingRoute(router *gin.RouterGroup) {
	auth := router.Group("/ping")
	{
		auth.GET(
			"",
			controllers.GetTopTweets,
		)
	}
}
