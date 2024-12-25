package requestModel

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PaginatedRequest struct {
	Limit int64 `form:"limit" json:"limit"`
	Page  int64 `form:"page" json:"page"`
}

func NewPaginatedRequest(limit int64, page int64) (*PaginatedRequest, error) {
	paginatedRequest := PaginatedRequest{
		Limit: int64(limit),
		Page:  int64(page),
	}

	err := validation.ValidateStruct(&paginatedRequest,
		validation.Field(&paginatedRequest.Limit, validation.Required, validation.Min(1), validation.Max(10)),
		validation.Field(&paginatedRequest.Page, validation.Required, validation.Min(1)),
	)

	if err != nil {
		return nil, err
	}

	return &paginatedRequest, nil
}

type PathIdRequest struct {
	Id primitive.ObjectID `form:"id"`
}

func NewPathIdRequest(id string) (*PathIdRequest, error) {
	result, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	pathIdRequest := &PathIdRequest{
		Id: result,
	}
	return pathIdRequest, nil
}

type RequestMetadata struct {
	UserId primitive.ObjectID `json:"id" bson:"_id"`
}

func NewRequestMetadata(userId primitive.ObjectID) (*RequestMetadata, error) {
	return &RequestMetadata{
		UserId: userId,
	}, nil
}
