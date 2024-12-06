package error

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"net/http"
	"sync"
)

var (
	Wrap         = errors.Wrap
	Wrapf        = errors.Wrapf
	Unwrap       = errors.Unwrap
	WithStack    = errors.WithStack
	WithMessage  = errors.WithMessage
	WithMessagef = errors.WithMessagef
	Cause        = errors.Cause
	Errorf       = errors.Errorf
	Is           = errors.Is
	As           = errors.As
)

const (
	DefaultBadRequestID            = "bad_request"
	DefaultUnauthorizedID          = "unauthorized"
	DefaultForbiddenID             = "forbidden"
	DefaultNotFoundID              = "not_found"
	DefaultMethodNotAllowedID      = "method_not_allowed"
	DefaultTooManyRequestsID       = "too_many_requests"
	DefaultRequestEntityTooLargeID = "request_entity_too_large"
	DefaultInternalServerErrorID   = "internal_server_error"
	DefaultConflictID              = "conflict"
	DefaultRequestTimeoutID        = "request_timeout"
)

// Error customize the error structure for implementation errors.Error interface
type Error struct {
	ID     string `json_utils:"id,omitempty"`     // 错误类型
	Code   int    `json_utils:"code,omitempty"`   // 错误码
	Detail string `json_utils:"detail,omitempty"` // 错误信息
	Status string `json_utils:"status,omitempty"` // 响应状态
}

func (e *Error) Error() string {
	errByte, _ := json.Marshal(e)
	return string(errByte)
}

func (e *Error) ErrorText() string {
	return e.Detail
}

func (e *Error) ErrorCode() int {
	return e.Code
}

// New custom error
func New(id, detail string, code int) error {
	return &Error{
		ID:     id,
		Code:   code,
		Detail: detail,
		Status: http.StatusText(int(code)),
	}
}

// Parse used to resolve string type to Error type.
// if failed it will set the string as Error.Detail
func Parse(err string) *Error {
	newErr := new(Error)
	result := json.Unmarshal([]byte(err), newErr)
	if result != nil {
		newErr.Detail = err
	}

	return newErr
}

// BadRequest generates a 400 error.
func BadRequest(id, format string, a ...interface{}) error {
	if id == "" {
		id = DefaultBadRequestID
	}

	return &Error{
		ID:     id,
		Code:   http.StatusBadRequest,
		Detail: fmt.Sprintf(format, a...),
		Status: http.StatusText(http.StatusBadRequest),
	}
}

// Unauthorized generates a 401 error.
func Unauthorized(id, format string, a ...interface{}) error {
	if id == "" {
		id = DefaultUnauthorizedID
	}

	return &Error{
		ID:     id,
		Code:   http.StatusUnauthorized,
		Detail: fmt.Sprintf(format, a...),
		Status: http.StatusText(http.StatusUnauthorized),
	}
}

// Forbidden generates a 403 error.
func Forbidden(id, format string, a ...interface{}) error {
	if id == "" {
		id = DefaultForbiddenID
	}

	return &Error{
		ID:     id,
		Code:   http.StatusForbidden,
		Detail: fmt.Sprintf(format, a...),
		Status: http.StatusText(http.StatusForbidden),
	}
}

// NotFound generates a 404 error.
func NotFound(id, format string, a ...interface{}) error {
	if id == "" {
		id = DefaultNotFoundID
	}

	return &Error{
		ID:     id,
		Code:   http.StatusNotFound,
		Detail: fmt.Sprintf(format, a...),
		Status: http.StatusText(http.StatusNotFound),
	}
}

// MethodNotAllowed generates a 405 error.
func MethodNotAllowed(id, format string, a ...interface{}) error {
	if id == "" {
		id = DefaultMethodNotAllowedID
	}

	return &Error{
		ID:     id,
		Code:   http.StatusMethodNotAllowed,
		Detail: fmt.Sprintf(format, a...),
		Status: http.StatusText(http.StatusMethodNotAllowed),
	}
}

// TooManyRequests generates a 429 error.
func TooManyRequests(id, format string, a ...interface{}) error {
	if id == "" {
		id = DefaultTooManyRequestsID
	}

	return &Error{
		ID:     id,
		Code:   http.StatusTooManyRequests,
		Detail: fmt.Sprintf(format, a...),
		Status: http.StatusText(http.StatusTooManyRequests),
	}
}

// Timeout generates a 408 error.
func Timeout(id, format string, a ...interface{}) error {
	if id == "" {
		id = DefaultRequestTimeoutID
	}

	return &Error{
		ID:     id,
		Code:   http.StatusRequestTimeout,
		Detail: fmt.Sprintf(format, a...),
		Status: http.StatusText(http.StatusRequestTimeout),
	}
}

// Conflict generates a 409 error.
func Conflict(id, format string, a ...interface{}) error {
	if id == "" {
		id = DefaultConflictID
	}

	return &Error{
		ID:     id,
		Code:   http.StatusConflict,
		Detail: fmt.Sprintf(format, a...),
		Status: http.StatusText(http.StatusConflict),
	}
}

// RequestEntityTooLarge generates a 413 error.
func RequestEntityTooLarge(id, format string, a ...interface{}) error {
	if id == "" {
		id = DefaultRequestEntityTooLargeID
	}

	return &Error{
		ID:     id,
		Code:   http.StatusRequestEntityTooLarge,
		Detail: fmt.Sprintf(format, a...),
		Status: http.StatusText(http.StatusRequestEntityTooLarge),
	}
}

// InternalServerError generates a 500 error.
func InternalServerError(id, format string, a ...interface{}) error {
	if id == "" {
		id = DefaultInternalServerErrorID
	}

	return &Error{
		ID:     id,
		Code:   http.StatusInternalServerError,
		Detail: fmt.Sprintf(format, a...),
		Status: http.StatusText(http.StatusInternalServerError),
	}
}

// Equal used to compare errors
func Equal(err1, err2 error) bool {
	var result1 *Error
	ok1 := As(err1, &result1)
	var result2 *Error
	ok2 := As(err2, &result2)

	if ok1 != ok2 {
		return false
	}

	if !ok1 {
		return Is(err1, err2)
	}

	if result1.Code != result2.Code {
		return false
	}

	return true
}

// FromError used to convert go error to *Error
func FromError(err error) *Error {
	if err == nil {
		return nil
	}

	var result *Error
	if ok := As(err, &result); ok && result != nil {
		return result
	}

	return Parse(err.Error())
}

type MultiError struct {
	lock   *sync.Mutex
	Errors []error
}

func (e *MultiError) Error() string {
	errByte, _ := json.Marshal(e)
	return string(errByte)
}

func NewMultiError() *MultiError {
	return &MultiError{
		lock:   &sync.Mutex{},
		Errors: make([]error, 0),
	}
}

func (e *MultiError) Append(err error) {
	e.Errors = append(e.Errors, err)
}

func (e *MultiError) AppendWithLock(err error) {
	e.lock.Lock()
	defer e.lock.Unlock()
	e.Append(err)
}

func (e *MultiError) HasErrors() bool {
	return len(e.Errors) > 0
}
