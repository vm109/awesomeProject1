package model

import (
	"context"
)

const (
	ErrorResponseErrorCodeJsonKey    = "error_code"
	ErrorResponseErrorMessageJsonKey = "error_message"
	ErrorResponseRequestIdJsonKey    = "request_id"
)

type ErrorResponse struct {
	RequestId string    `json:"request_id,omitempty"`
	Message   string    `json:"error_message,omitempty"`
	Code      ErrorCode `json:"error_code,omitempty"`
}

func ExtractRequestId(ctx context.Context) string {
	if fields, ok := ctx.Value(&struct{}{}).(map[string]string); ok {
		if value, ok := fields[ErrorResponseRequestIdJsonKey]; ok {
			return value
		}
	}
	return ""
}
