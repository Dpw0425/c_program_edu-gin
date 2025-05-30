// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: test_data.proto

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

// Validate checks the field values on GetTestDataRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetTestDataRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetTestDataRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetTestDataRequestMultiError, or nil if none found.
func (m *GetTestDataRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetTestDataRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for QuestionId

	if len(errors) > 0 {
		return GetTestDataRequestMultiError(errors)
	}

	return nil
}

// GetTestDataRequestMultiError is an error wrapping multiple validation errors
// returned by GetTestDataRequest.ValidateAll() if the designated constraints
// aren't met.
type GetTestDataRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetTestDataRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetTestDataRequestMultiError) AllErrors() []error { return m }

// GetTestDataRequestValidationError is the validation error returned by
// GetTestDataRequest.Validate if the designated constraints aren't met.
type GetTestDataRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetTestDataRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetTestDataRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetTestDataRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetTestDataRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetTestDataRequestValidationError) ErrorName() string {
	return "GetTestDataRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetTestDataRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetTestDataRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetTestDataRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetTestDataRequestValidationError{}

// Validate checks the field values on GetTestDataResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetTestDataResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetTestDataResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetTestDataResponseMultiError, or nil if none found.
func (m *GetTestDataResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetTestDataResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetTestDataList() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, GetTestDataResponseValidationError{
						field:  fmt.Sprintf("TestDataList[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, GetTestDataResponseValidationError{
						field:  fmt.Sprintf("TestDataList[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return GetTestDataResponseValidationError{
					field:  fmt.Sprintf("TestDataList[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return GetTestDataResponseMultiError(errors)
	}

	return nil
}

// GetTestDataResponseMultiError is an error wrapping multiple validation
// errors returned by GetTestDataResponse.ValidateAll() if the designated
// constraints aren't met.
type GetTestDataResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetTestDataResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetTestDataResponseMultiError) AllErrors() []error { return m }

// GetTestDataResponseValidationError is the validation error returned by
// GetTestDataResponse.Validate if the designated constraints aren't met.
type GetTestDataResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetTestDataResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetTestDataResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetTestDataResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetTestDataResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetTestDataResponseValidationError) ErrorName() string {
	return "GetTestDataResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetTestDataResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetTestDataResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetTestDataResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetTestDataResponseValidationError{}

// Validate checks the field values on UpdateTestDataRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UpdateTestDataRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateTestDataRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateTestDataRequestMultiError, or nil if none found.
func (m *UpdateTestDataRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateTestDataRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Input

	// no validation rules for Output

	if len(errors) > 0 {
		return UpdateTestDataRequestMultiError(errors)
	}

	return nil
}

// UpdateTestDataRequestMultiError is an error wrapping multiple validation
// errors returned by UpdateTestDataRequest.ValidateAll() if the designated
// constraints aren't met.
type UpdateTestDataRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateTestDataRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateTestDataRequestMultiError) AllErrors() []error { return m }

// UpdateTestDataRequestValidationError is the validation error returned by
// UpdateTestDataRequest.Validate if the designated constraints aren't met.
type UpdateTestDataRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateTestDataRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateTestDataRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateTestDataRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateTestDataRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateTestDataRequestValidationError) ErrorName() string {
	return "UpdateTestDataRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateTestDataRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateTestDataRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateTestDataRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateTestDataRequestValidationError{}

// Validate checks the field values on UpdateTestDataResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UpdateTestDataResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateTestDataResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateTestDataResponseMultiError, or nil if none found.
func (m *UpdateTestDataResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateTestDataResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return UpdateTestDataResponseMultiError(errors)
	}

	return nil
}

// UpdateTestDataResponseMultiError is an error wrapping multiple validation
// errors returned by UpdateTestDataResponse.ValidateAll() if the designated
// constraints aren't met.
type UpdateTestDataResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateTestDataResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateTestDataResponseMultiError) AllErrors() []error { return m }

// UpdateTestDataResponseValidationError is the validation error returned by
// UpdateTestDataResponse.Validate if the designated constraints aren't met.
type UpdateTestDataResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateTestDataResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateTestDataResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateTestDataResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateTestDataResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateTestDataResponseValidationError) ErrorName() string {
	return "UpdateTestDataResponseValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateTestDataResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateTestDataResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateTestDataResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateTestDataResponseValidationError{}

// Validate checks the field values on DeleteTestDataRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DeleteTestDataRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteTestDataRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteTestDataRequestMultiError, or nil if none found.
func (m *DeleteTestDataRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteTestDataRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return DeleteTestDataRequestMultiError(errors)
	}

	return nil
}

// DeleteTestDataRequestMultiError is an error wrapping multiple validation
// errors returned by DeleteTestDataRequest.ValidateAll() if the designated
// constraints aren't met.
type DeleteTestDataRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteTestDataRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteTestDataRequestMultiError) AllErrors() []error { return m }

// DeleteTestDataRequestValidationError is the validation error returned by
// DeleteTestDataRequest.Validate if the designated constraints aren't met.
type DeleteTestDataRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteTestDataRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteTestDataRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteTestDataRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteTestDataRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteTestDataRequestValidationError) ErrorName() string {
	return "DeleteTestDataRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteTestDataRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteTestDataRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteTestDataRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteTestDataRequestValidationError{}

// Validate checks the field values on DeleteTestDataResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DeleteTestDataResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteTestDataResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteTestDataResponseMultiError, or nil if none found.
func (m *DeleteTestDataResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteTestDataResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return DeleteTestDataResponseMultiError(errors)
	}

	return nil
}

// DeleteTestDataResponseMultiError is an error wrapping multiple validation
// errors returned by DeleteTestDataResponse.ValidateAll() if the designated
// constraints aren't met.
type DeleteTestDataResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteTestDataResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteTestDataResponseMultiError) AllErrors() []error { return m }

// DeleteTestDataResponseValidationError is the validation error returned by
// DeleteTestDataResponse.Validate if the designated constraints aren't met.
type DeleteTestDataResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteTestDataResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteTestDataResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteTestDataResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteTestDataResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteTestDataResponseValidationError) ErrorName() string {
	return "DeleteTestDataResponseValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteTestDataResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteTestDataResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteTestDataResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteTestDataResponseValidationError{}

// Validate checks the field values on GetTestDataResponse_TestDataItem with
// the rules defined in the proto definition for this message. If any rules
// are violated, the first error encountered is returned, or nil if there are
// no violations.
func (m *GetTestDataResponse_TestDataItem) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetTestDataResponse_TestDataItem with
// the rules defined in the proto definition for this message. If any rules
// are violated, the result is a list of violation errors wrapped in
// GetTestDataResponse_TestDataItemMultiError, or nil if none found.
func (m *GetTestDataResponse_TestDataItem) ValidateAll() error {
	return m.validate(true)
}

func (m *GetTestDataResponse_TestDataItem) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Input

	// no validation rules for Output

	// no validation rules for QuestionId

	if len(errors) > 0 {
		return GetTestDataResponse_TestDataItemMultiError(errors)
	}

	return nil
}

// GetTestDataResponse_TestDataItemMultiError is an error wrapping multiple
// validation errors returned by
// GetTestDataResponse_TestDataItem.ValidateAll() if the designated
// constraints aren't met.
type GetTestDataResponse_TestDataItemMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetTestDataResponse_TestDataItemMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetTestDataResponse_TestDataItemMultiError) AllErrors() []error { return m }

// GetTestDataResponse_TestDataItemValidationError is the validation error
// returned by GetTestDataResponse_TestDataItem.Validate if the designated
// constraints aren't met.
type GetTestDataResponse_TestDataItemValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetTestDataResponse_TestDataItemValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetTestDataResponse_TestDataItemValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetTestDataResponse_TestDataItemValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetTestDataResponse_TestDataItemValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetTestDataResponse_TestDataItemValidationError) ErrorName() string {
	return "GetTestDataResponse_TestDataItemValidationError"
}

// Error satisfies the builtin error interface
func (e GetTestDataResponse_TestDataItemValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetTestDataResponse_TestDataItem.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetTestDataResponse_TestDataItemValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetTestDataResponse_TestDataItemValidationError{}
