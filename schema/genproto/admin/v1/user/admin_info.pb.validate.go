// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: admin_info.proto

package admin

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on AdminInfoRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *AdminInfoRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AdminInfoRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// AdminInfoRequestMultiError, or nil if none found.
func (m *AdminInfoRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *AdminInfoRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return AdminInfoRequestMultiError(errors)
	}

	return nil
}

// AdminInfoRequestMultiError is an error wrapping multiple validation errors
// returned by AdminInfoRequest.ValidateAll() if the designated constraints
// aren't met.
type AdminInfoRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AdminInfoRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AdminInfoRequestMultiError) AllErrors() []error { return m }

// AdminInfoRequestValidationError is the validation error returned by
// AdminInfoRequest.Validate if the designated constraints aren't met.
type AdminInfoRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AdminInfoRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AdminInfoRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AdminInfoRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AdminInfoRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AdminInfoRequestValidationError) ErrorName() string { return "AdminInfoRequestValidationError" }

// Error satisfies the builtin error interface
func (e AdminInfoRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAdminInfoRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AdminInfoRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AdminInfoRequestValidationError{}

// Validate checks the field values on AdminInfoResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *AdminInfoResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AdminInfoResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// AdminInfoResponseMultiError, or nil if none found.
func (m *AdminInfoResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *AdminInfoResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for UserName

	// no validation rules for TeacherId

	// no validation rules for Status

	if len(errors) > 0 {
		return AdminInfoResponseMultiError(errors)
	}

	return nil
}

// AdminInfoResponseMultiError is an error wrapping multiple validation errors
// returned by AdminInfoResponse.ValidateAll() if the designated constraints
// aren't met.
type AdminInfoResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AdminInfoResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AdminInfoResponseMultiError) AllErrors() []error { return m }

// AdminInfoResponseValidationError is the validation error returned by
// AdminInfoResponse.Validate if the designated constraints aren't met.
type AdminInfoResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AdminInfoResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AdminInfoResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AdminInfoResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AdminInfoResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AdminInfoResponseValidationError) ErrorName() string {
	return "AdminInfoResponseValidationError"
}

// Error satisfies the builtin error interface
func (e AdminInfoResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAdminInfoResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AdminInfoResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AdminInfoResponseValidationError{}
