package validators

import (
	"net/http"
	responseModel "twitterApi/models/response"

	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

func GetAllProductValidator() gin.HandlerFunc {
	return func(c *gin.Context) {

		id := c.Param("id")
		err := validation.Validate(id, is.MongoID)
		if err != nil {
			responseModel.SendErrorResponse(c, http.StatusBadRequest, "invalid id: "+id)
			return
		}

		c.Next()
	}
}
