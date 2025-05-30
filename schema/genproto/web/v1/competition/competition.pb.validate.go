// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: competition.proto

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

// Validate checks the field values on GetCompetitionListRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetCompetitionListRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetCompetitionListRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetCompetitionListRequestMultiError, or nil if none found.
func (m *GetCompetitionListRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetCompetitionListRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Page

	// no validation rules for Number

	// no validation rules for Search

	if len(errors) > 0 {
		return GetCompetitionListRequestMultiError(errors)
	}

	return nil
}

// GetCompetitionListRequestMultiError is an error wrapping multiple validation
// errors returned by GetCompetitionListRequest.ValidateAll() if the
// designated constraints aren't met.
type GetCompetitionListRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetCompetitionListRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetCompetitionListRequestMultiError) AllErrors() []error { return m }

// GetCompetitionListRequestValidationError is the validation error returned by
// GetCompetitionListRequest.Validate if the designated constraints aren't met.
type GetCompetitionListRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetCompetitionListRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetCompetitionListRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetCompetitionListRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetCompetitionListRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetCompetitionListRequestValidationError) ErrorName() string {
	return "GetCompetitionListRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetCompetitionListRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetCompetitionListRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetCompetitionListRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetCompetitionListRequestValidationError{}

// Validate checks the field values on GetCompetitionListResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetCompetitionListResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetCompetitionListResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetCompetitionListResponseMultiError, or nil if none found.
func (m *GetCompetitionListResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetCompetitionListResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetCompetitionList() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, GetCompetitionListResponseValidationError{
						field:  fmt.Sprintf("CompetitionList[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, GetCompetitionListResponseValidationError{
						field:  fmt.Sprintf("CompetitionList[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return GetCompetitionListResponseValidationError{
					field:  fmt.Sprintf("CompetitionList[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for Total

	if len(errors) > 0 {
		return GetCompetitionListResponseMultiError(errors)
	}

	return nil
}

// GetCompetitionListResponseMultiError is an error wrapping multiple
// validation errors returned by GetCompetitionListResponse.ValidateAll() if
// the designated constraints aren't met.
type GetCompetitionListResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetCompetitionListResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetCompetitionListResponseMultiError) AllErrors() []error { return m }

// GetCompetitionListResponseValidationError is the validation error returned
// by GetCompetitionListResponse.Validate if the designated constraints aren't met.
type GetCompetitionListResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetCompetitionListResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetCompetitionListResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetCompetitionListResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetCompetitionListResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetCompetitionListResponseValidationError) ErrorName() string {
	return "GetCompetitionListResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetCompetitionListResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetCompetitionListResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetCompetitionListResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetCompetitionListResponseValidationError{}

// Validate checks the field values on
// GetCompetitionListResponse_CompetitionItem with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *GetCompetitionListResponse_CompetitionItem) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on
// GetCompetitionListResponse_CompetitionItem with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in
// GetCompetitionListResponse_CompetitionItemMultiError, or nil if none found.
func (m *GetCompetitionListResponse_CompetitionItem) ValidateAll() error {
	return m.validate(true)
}

func (m *GetCompetitionListResponse_CompetitionItem) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Name

	// no validation rules for OwnerId

	// no validation rules for StartTime

	// no validation rules for Deadline

	// no validation rules for Status

	// no validation rules for Category

	// no validation rules for Permission

	// no validation rules for Quantity

	if len(errors) > 0 {
		return GetCompetitionListResponse_CompetitionItemMultiError(errors)
	}

	return nil
}

// GetCompetitionListResponse_CompetitionItemMultiError is an error wrapping
// multiple validation errors returned by
// GetCompetitionListResponse_CompetitionItem.ValidateAll() if the designated
// constraints aren't met.
type GetCompetitionListResponse_CompetitionItemMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetCompetitionListResponse_CompetitionItemMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetCompetitionListResponse_CompetitionItemMultiError) AllErrors() []error { return m }

// GetCompetitionListResponse_CompetitionItemValidationError is the validation
// error returned by GetCompetitionListResponse_CompetitionItem.Validate if
// the designated constraints aren't met.
type GetCompetitionListResponse_CompetitionItemValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetCompetitionListResponse_CompetitionItemValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetCompetitionListResponse_CompetitionItemValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetCompetitionListResponse_CompetitionItemValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetCompetitionListResponse_CompetitionItemValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetCompetitionListResponse_CompetitionItemValidationError) ErrorName() string {
	return "GetCompetitionListResponse_CompetitionItemValidationError"
}

// Error satisfies the builtin error interface
func (e GetCompetitionListResponse_CompetitionItemValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetCompetitionListResponse_CompetitionItem.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetCompetitionListResponse_CompetitionItemValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetCompetitionListResponse_CompetitionItemValidationError{}
