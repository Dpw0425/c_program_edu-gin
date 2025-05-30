// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: question_list.proto

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

// Validate checks the field values on QuestionListRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *QuestionListRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on QuestionListRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// QuestionListRequestMultiError, or nil if none found.
func (m *QuestionListRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *QuestionListRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Search

	// no validation rules for Page

	// no validation rules for Number

	if len(errors) > 0 {
		return QuestionListRequestMultiError(errors)
	}

	return nil
}

// QuestionListRequestMultiError is an error wrapping multiple validation
// errors returned by QuestionListRequest.ValidateAll() if the designated
// constraints aren't met.
type QuestionListRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m QuestionListRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m QuestionListRequestMultiError) AllErrors() []error { return m }

// QuestionListRequestValidationError is the validation error returned by
// QuestionListRequest.Validate if the designated constraints aren't met.
type QuestionListRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e QuestionListRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e QuestionListRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e QuestionListRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e QuestionListRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e QuestionListRequestValidationError) ErrorName() string {
	return "QuestionListRequestValidationError"
}

// Error satisfies the builtin error interface
func (e QuestionListRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sQuestionListRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = QuestionListRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = QuestionListRequestValidationError{}

// Validate checks the field values on QuestionListResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *QuestionListResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on QuestionListResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// QuestionListResponseMultiError, or nil if none found.
func (m *QuestionListResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *QuestionListResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetQuestionList() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, QuestionListResponseValidationError{
						field:  fmt.Sprintf("QuestionList[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, QuestionListResponseValidationError{
						field:  fmt.Sprintf("QuestionList[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return QuestionListResponseValidationError{
					field:  fmt.Sprintf("QuestionList[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for Total

	if len(errors) > 0 {
		return QuestionListResponseMultiError(errors)
	}

	return nil
}

// QuestionListResponseMultiError is an error wrapping multiple validation
// errors returned by QuestionListResponse.ValidateAll() if the designated
// constraints aren't met.
type QuestionListResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m QuestionListResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m QuestionListResponseMultiError) AllErrors() []error { return m }

// QuestionListResponseValidationError is the validation error returned by
// QuestionListResponse.Validate if the designated constraints aren't met.
type QuestionListResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e QuestionListResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e QuestionListResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e QuestionListResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e QuestionListResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e QuestionListResponseValidationError) ErrorName() string {
	return "QuestionListResponseValidationError"
}

// Error satisfies the builtin error interface
func (e QuestionListResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sQuestionListResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = QuestionListResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = QuestionListResponseValidationError{}

// Validate checks the field values on QuestionListResponse_QuestionItem with
// the rules defined in the proto definition for this message. If any rules
// are violated, the first error encountered is returned, or nil if there are
// no violations.
func (m *QuestionListResponse_QuestionItem) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on QuestionListResponse_QuestionItem
// with the rules defined in the proto definition for this message. If any
// rules are violated, the result is a list of violation errors wrapped in
// QuestionListResponse_QuestionItemMultiError, or nil if none found.
func (m *QuestionListResponse_QuestionItem) ValidateAll() error {
	return m.validate(true)
}

func (m *QuestionListResponse_QuestionItem) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Title

	// no validation rules for Degree

	// no validation rules for PassingRate

	// no validation rules for OwnerId

	// no validation rules for Content

	// no validation rules for Answer

	// no validation rules for Status

	if len(errors) > 0 {
		return QuestionListResponse_QuestionItemMultiError(errors)
	}

	return nil
}

// QuestionListResponse_QuestionItemMultiError is an error wrapping multiple
// validation errors returned by
// QuestionListResponse_QuestionItem.ValidateAll() if the designated
// constraints aren't met.
type QuestionListResponse_QuestionItemMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m QuestionListResponse_QuestionItemMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m QuestionListResponse_QuestionItemMultiError) AllErrors() []error { return m }

// QuestionListResponse_QuestionItemValidationError is the validation error
// returned by QuestionListResponse_QuestionItem.Validate if the designated
// constraints aren't met.
type QuestionListResponse_QuestionItemValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e QuestionListResponse_QuestionItemValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e QuestionListResponse_QuestionItemValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e QuestionListResponse_QuestionItemValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e QuestionListResponse_QuestionItemValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e QuestionListResponse_QuestionItemValidationError) ErrorName() string {
	return "QuestionListResponse_QuestionItemValidationError"
}

// Error satisfies the builtin error interface
func (e QuestionListResponse_QuestionItemValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sQuestionListResponse_QuestionItem.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = QuestionListResponse_QuestionItemValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = QuestionListResponse_QuestionItemValidationError{}
