package api

import (
	"context"
	"encoding/json"
	"time"

	"github.com/think-free/sportsnews/lib/logging"
)

// Response of the api, call NewResponse to create a new response with the given status and data, it will return a json byte array
type Response struct {
	Status   string      `json:"status"`
	Data     interface{} `json:"data,omitempty"`
	Metadata Metadata    `json:"metadata"`
}

type Metadata struct {
	CreatedAt  string `json:"createdAt,omitempty"`
	TotalItems int    `json:"totalItems,omitempty"`
}

func NewResponse(ctx context.Context, status string, data interface{}) []byte {
	// Creating response
	r := &Response{
		Status: status,
		Data:   data,
		Metadata: Metadata{
			CreatedAt: time.Now().Format(time.RFC3339),
		},
	}

	// Marshalling response
	js, err := json.Marshal(r)
	if err != nil {
		logging.L(ctx).Errorf("Error marshalling json: %s", err.Error())
		return []byte(JSONStatusError)
	}
	return js
}
