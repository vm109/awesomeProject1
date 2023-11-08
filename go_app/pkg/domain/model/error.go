package model

import (
	"fmt"
	"github.com/pkg/errors"
)

const (
	ErrNotFound                ErrorCode = "ErrNotFound"
	ErrInvalidData             ErrorCode = "ErrInvalidData"
	ErrInvalidParam            ErrorCode = "ErrInvalidParam"
	ErrAlreadyExists           ErrorCode = "ErrAlreadyExists"
	ErrUrlAlreadyExists        ErrorCode = "ErrUrlAlreadyExists"
	ErrGatewayComm             ErrorCode = "ErrGatewayComm"
	ErrNotImplemented          ErrorCode = "ErrNotImplemented"
	ErrInternalServerException ErrorCode = "InternalServerException"
	ErrRateLimitExceeded       ErrorCode = "ErrRateLimitExceeded"
	ErrPreconditionFailed      ErrorCode = "ErrPreconditionFailed"
	ErrRestrictedContent       ErrorCode = "ErrRestrictedContent"
	ErrStatusLocked            ErrorCode = "ErrStatusLocked"
	ErrInvalidBSON             ErrorCode = "ErrInvalidBSON"
	ErrSizeExceeded            ErrorCode = "ErrorSizeExceeded"
	ErrInvalidAns              ErrorCode = "ErrInvalidAns"
)

type ErrorCode string

type Error struct {
	Code      ErrorCode `json:"error_code"`
	Msg       string    `json:"error_message"`
	WithTrace error     `json:"-"`
}

func NewErrorf(code ErrorCode, msgFormat string, args ...interface{}) error {
	e := Error{code, fmt.Sprintf(msgFormat, args...), nil}
	e.WithTrace = errors.WithStack(&e)
	return &e
}
func (e *Error) Error() string {
	return fmt.Sprintf("%s - %s", e.Code, e.Msg)
}

func EndpointNotFoundError() error {
	return NewErrorf(ErrNotFound, "not found")
}

func InvalidData() error {
	return NewErrorf(ErrInvalidData, "Invalid Data")
}
