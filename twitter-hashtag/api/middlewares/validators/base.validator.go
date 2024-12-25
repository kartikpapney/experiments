package validators

import (
	"net/http"
	"strconv"
	requestModel "twitterApi/models/request"
	responseModel "twitterApi/models/response"

	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

func PathIdValidator() gin.HandlerFunc {
	return func(c *gin.Context) {

		id := c.Param("id")
		err := validation.Validate(id, is.MongoID)
		if err != nil {
			responseModel.SendErrorResponse(c, http.StatusBadRequest, "invalid id: "+id)
			return
		}

		idHex, err := requestModel.NewPathIdRequest(id)
		if err != nil {
			responseModel.SendErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}

		c.Set("id", idHex)

		c.Next()
	}
}

func PaginationRequestValidator() gin.HandlerFunc {
	return func(c *gin.Context) {

		limit := c.DefaultQuery("limit", "10")
		limitInt, err := strconv.Atoi(limit)
		if err != nil || limitInt > 10 || limitInt <= 0 {
			responseModel.SendErrorResponse(c, http.StatusBadRequest, "invalid limit, should be a positive integer not greater than 10")
			return
		}

		page := c.DefaultQuery("page", "0")
		pageInt, err := strconv.Atoi(page)
		if err != nil || pageInt < 0 {
			responseModel.SendErrorResponse(c, http.StatusBadRequest, "invalid page, should be a non-negative integer")
			return
		}

		paginatedRequest, err := requestModel.NewPaginatedRequest(int64(limitInt), int64(pageInt))
		if err != nil {
			responseModel.SendErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}

		c.Set("pagination", paginatedRequest)
		c.Next()
	}
}
