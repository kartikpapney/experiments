package routes

import (
	"twitterApi/controllers"

	"github.com/gin-gonic/gin"
)

func TrendingRoute(router *gin.RouterGroup) {
	auth := router.Group("/trending")
	{
		auth.GET(
			"",
			controllers.GetTopTweets,
		)
	}
}
