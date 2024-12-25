package middlewares

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

// DBMiddleware is the middleware that injects the db connection into the request context
func DBMiddleware(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Add the db connection to the request context
		c.Set("db", db)
		// Proceed with the next handler
		c.Next()
	}
}
