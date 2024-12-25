package responseModel

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Response Base response
type PaginatedResponse struct {
	Limit      int64          `json:"limit"`
	Page       int64          `json:"page"`
	Total      int64          `json:"total"`
	StatusCode int            `json:"-"`
	Success    bool           `json:"success"`
	Message    string         `json:"message,omitempty"`
	Data       map[string]any `json:"data,omitempty"`
}

func (response *PaginatedResponse) SendResponse(c *gin.Context) {
	c.AbortWithStatusJSON(response.StatusCode, response)
}

func SendPaginatedResponseData(c *gin.Context, data gin.H, limit int64, page int64, total int64) {
	response := &PaginatedResponse{
		Limit:      limit,
		Total:      total,
		Page:       page,
		StatusCode: http.StatusOK,
		Success:    true,
		Data:       data,
	}
	response.SendResponse(c)
}

func SendPaginatedErrorResponse(c *gin.Context, status int, message string) {
	response := &PaginatedResponse{
		StatusCode: status,
		Success:    false,
		Message:    message,
	}
	response.SendResponse(c)
}
