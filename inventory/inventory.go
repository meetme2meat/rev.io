package inventory

import (
	"context"

	"github.com/meetme2meat/rev.io/http"
)

type Inventory struct {
	http.HTTPClient
}

func (i Inventory) Create(ctx context.Context, params interface{}) InventoryResponse {
	response, err := i.POST(ctx, "https://restapi.rev.io/v1/InventoryItem")
	if err != nil {
		return InventoryResponse{
			Code:    500,
			Error:   err,
			Message: "BAD_REQUEST",
		}
	}

	defer response.Body.Close()
	return InventoryResponse{
		Code:    200,
		Error:   nil,
		Message: "",
	}
}

func (i Inventory) Get(ctx context.Context, params interface{}) InventoryResponse {
	response, err := i.GET(ctx, "https://restapi.rev.io/v1/InventoryItem")
	if err != nil {
		return InventoryResponse{
			Code:    500,
			Error:   err,
			Message: "BAD_REQUEST",
		}
	}

	defer response.Body.Close()
	return InventoryResponse{
		Code:    200,
		Error:   nil,
		Message: "",
	}
}

type InventoryResponse struct {
	Code    uint // status code
	Error   error
	Message string
	Data    []map[string]interface{}
}
