// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: delete_bank.proto

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

// Validate checks the field values on DeleteBankRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *DeleteBankRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteBankRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteBankRequestMultiError, or nil if none found.
func (m *DeleteBankRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteBankRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return DeleteBankRequestMultiError(errors)
	}

	return nil
}

// DeleteBankRequestMultiError is an error wrapping multiple validation errors
// returned by DeleteBankRequest.ValidateAll() if the designated constraints
// aren't met.
type DeleteBankRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteBankRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteBankRequestMultiError) AllErrors() []error { return m }

// DeleteBankRequestValidationError is the validation error returned by
// DeleteBankRequest.Validate if the designated constraints aren't met.
type DeleteBankRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteBankRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteBankRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteBankRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteBankRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteBankRequestValidationError) ErrorName() string {
	return "DeleteBankRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteBankRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteBankRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteBankRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteBankRequestValidationError{}

// Validate checks the field values on DeleteBankResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DeleteBankResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteBankResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteBankResponseMultiError, or nil if none found.
func (m *DeleteBankResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteBankResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return DeleteBankResponseMultiError(errors)
	}

	return nil
}

// DeleteBankResponseMultiError is an error wrapping multiple validation errors
// returned by DeleteBankResponse.ValidateAll() if the designated constraints
// aren't met.
type DeleteBankResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteBankResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteBankResponseMultiError) AllErrors() []error { return m }

// DeleteBankResponseValidationError is the validation error returned by
// DeleteBankResponse.Validate if the designated constraints aren't met.
type DeleteBankResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteBankResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteBankResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteBankResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteBankResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteBankResponseValidationError) ErrorName() string {
	return "DeleteBankResponseValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteBankResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteBankResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteBankResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteBankResponseValidationError{}

// Validate checks the field values on ExcludeQuestionRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ExcludeQuestionRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ExcludeQuestionRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ExcludeQuestionRequestMultiError, or nil if none found.
func (m *ExcludeQuestionRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ExcludeQuestionRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for BankId

	// no validation rules for QuestionId

	if len(errors) > 0 {
		return ExcludeQuestionRequestMultiError(errors)
	}

	return nil
}

// ExcludeQuestionRequestMultiError is an error wrapping multiple validation
// errors returned by ExcludeQuestionRequest.ValidateAll() if the designated
// constraints aren't met.
type ExcludeQuestionRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ExcludeQuestionRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ExcludeQuestionRequestMultiError) AllErrors() []error { return m }

// ExcludeQuestionRequestValidationError is the validation error returned by
// ExcludeQuestionRequest.Validate if the designated constraints aren't met.
type ExcludeQuestionRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ExcludeQuestionRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ExcludeQuestionRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ExcludeQuestionRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ExcludeQuestionRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ExcludeQuestionRequestValidationError) ErrorName() string {
	return "ExcludeQuestionRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ExcludeQuestionRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sExcludeQuestionRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ExcludeQuestionRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ExcludeQuestionRequestValidationError{}

// Validate checks the field values on ExcludeQuestionResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ExcludeQuestionResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ExcludeQuestionResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ExcludeQuestionResponseMultiError, or nil if none found.
func (m *ExcludeQuestionResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ExcludeQuestionResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return ExcludeQuestionResponseMultiError(errors)
	}

	return nil
}

// ExcludeQuestionResponseMultiError is an error wrapping multiple validation
// errors returned by ExcludeQuestionResponse.ValidateAll() if the designated
// constraints aren't met.
type ExcludeQuestionResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ExcludeQuestionResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ExcludeQuestionResponseMultiError) AllErrors() []error { return m }

// ExcludeQuestionResponseValidationError is the validation error returned by
// ExcludeQuestionResponse.Validate if the designated constraints aren't met.
type ExcludeQuestionResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ExcludeQuestionResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ExcludeQuestionResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ExcludeQuestionResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ExcludeQuestionResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ExcludeQuestionResponseValidationError) ErrorName() string {
	return "ExcludeQuestionResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ExcludeQuestionResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sExcludeQuestionResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ExcludeQuestionResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ExcludeQuestionResponseValidationError{}
