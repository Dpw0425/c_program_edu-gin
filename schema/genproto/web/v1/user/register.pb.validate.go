// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: register.proto

package web

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

// Validate checks the field values on UserRegisterRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UserRegisterRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UserRegisterRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UserRegisterRequestMultiError, or nil if none found.
func (m *UserRegisterRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *UserRegisterRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for UserName

	// no validation rules for Password

	// no validation rules for StudentId

	// no validation rules for Avatar

	// no validation rules for Email

	// no validation rules for Grade

	// no validation rules for VerifyCode

	if len(errors) > 0 {
		return UserRegisterRequestMultiError(errors)
	}

	return nil
}

// UserRegisterRequestMultiError is an error wrapping multiple validation
// errors returned by UserRegisterRequest.ValidateAll() if the designated
// constraints aren't met.
type UserRegisterRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UserRegisterRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UserRegisterRequestMultiError) AllErrors() []error { return m }

// UserRegisterRequestValidationError is the validation error returned by
// UserRegisterRequest.Validate if the designated constraints aren't met.
type UserRegisterRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UserRegisterRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UserRegisterRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UserRegisterRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UserRegisterRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UserRegisterRequestValidationError) ErrorName() string {
	return "UserRegisterRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UserRegisterRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUserRegisterRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UserRegisterRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UserRegisterRequestValidationError{}

// Validate checks the field values on UserRegisterResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UserRegisterResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UserRegisterResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UserRegisterResponseMultiError, or nil if none found.
func (m *UserRegisterResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *UserRegisterResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return UserRegisterResponseMultiError(errors)
	}

	return nil
}

// UserRegisterResponseMultiError is an error wrapping multiple validation
// errors returned by UserRegisterResponse.ValidateAll() if the designated
// constraints aren't met.
type UserRegisterResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UserRegisterResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UserRegisterResponseMultiError) AllErrors() []error { return m }

// UserRegisterResponseValidationError is the validation error returned by
// UserRegisterResponse.Validate if the designated constraints aren't met.
type UserRegisterResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UserRegisterResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UserRegisterResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UserRegisterResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UserRegisterResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UserRegisterResponseValidationError) ErrorName() string {
	return "UserRegisterResponseValidationError"
}

// Error satisfies the builtin error interface
func (e UserRegisterResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUserRegisterResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UserRegisterResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UserRegisterResponseValidationError{}
