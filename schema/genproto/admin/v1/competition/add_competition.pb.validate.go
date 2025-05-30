// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: add_competition.proto

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

// Validate checks the field values on AddCompetitionRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *AddCompetitionRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AddCompetitionRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// AddCompetitionRequestMultiError, or nil if none found.
func (m *AddCompetitionRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *AddCompetitionRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Name

	// no validation rules for StartTime

	// no validation rules for Deadline

	// no validation rules for Category

	// no validation rules for Permission

	if len(errors) > 0 {
		return AddCompetitionRequestMultiError(errors)
	}

	return nil
}

// AddCompetitionRequestMultiError is an error wrapping multiple validation
// errors returned by AddCompetitionRequest.ValidateAll() if the designated
// constraints aren't met.
type AddCompetitionRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AddCompetitionRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AddCompetitionRequestMultiError) AllErrors() []error { return m }

// AddCompetitionRequestValidationError is the validation error returned by
// AddCompetitionRequest.Validate if the designated constraints aren't met.
type AddCompetitionRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AddCompetitionRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AddCompetitionRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AddCompetitionRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AddCompetitionRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AddCompetitionRequestValidationError) ErrorName() string {
	return "AddCompetitionRequestValidationError"
}

// Error satisfies the builtin error interface
func (e AddCompetitionRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAddCompetitionRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AddCompetitionRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AddCompetitionRequestValidationError{}

// Validate checks the field values on AddCompetitionResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *AddCompetitionResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AddCompetitionResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// AddCompetitionResponseMultiError, or nil if none found.
func (m *AddCompetitionResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *AddCompetitionResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return AddCompetitionResponseMultiError(errors)
	}

	return nil
}

// AddCompetitionResponseMultiError is an error wrapping multiple validation
// errors returned by AddCompetitionResponse.ValidateAll() if the designated
// constraints aren't met.
type AddCompetitionResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AddCompetitionResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AddCompetitionResponseMultiError) AllErrors() []error { return m }

// AddCompetitionResponseValidationError is the validation error returned by
// AddCompetitionResponse.Validate if the designated constraints aren't met.
type AddCompetitionResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AddCompetitionResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AddCompetitionResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AddCompetitionResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AddCompetitionResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AddCompetitionResponseValidationError) ErrorName() string {
	return "AddCompetitionResponseValidationError"
}

// Error satisfies the builtin error interface
func (e AddCompetitionResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAddCompetitionResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AddCompetitionResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AddCompetitionResponseValidationError{}
