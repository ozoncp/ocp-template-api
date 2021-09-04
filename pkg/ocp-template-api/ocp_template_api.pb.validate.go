// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: ozoncp/ocp_template_api/v1/ocp_template_api.proto

package ocp_template_api

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
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
)

// Validate checks the field values on Template with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Template) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	// no validation rules for Foo

	return nil
}

// TemplateValidationError is the validation error returned by
// Template.Validate if the designated constraints aren't met.
type TemplateValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TemplateValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TemplateValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TemplateValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TemplateValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TemplateValidationError) ErrorName() string { return "TemplateValidationError" }

// Error satisfies the builtin error interface
func (e TemplateValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTemplate.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TemplateValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TemplateValidationError{}

// Validate checks the field values on CreateTemplateV1Request with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CreateTemplateV1Request) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetFoo() <= 0 {
		return CreateTemplateV1RequestValidationError{
			field:  "Foo",
			reason: "value must be greater than 0",
		}
	}

	return nil
}

// CreateTemplateV1RequestValidationError is the validation error returned by
// CreateTemplateV1Request.Validate if the designated constraints aren't met.
type CreateTemplateV1RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateTemplateV1RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateTemplateV1RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateTemplateV1RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateTemplateV1RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateTemplateV1RequestValidationError) ErrorName() string {
	return "CreateTemplateV1RequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateTemplateV1RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateTemplateV1Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateTemplateV1RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateTemplateV1RequestValidationError{}

// Validate checks the field values on CreateTemplateV1Response with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CreateTemplateV1Response) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	return nil
}

// CreateTemplateV1ResponseValidationError is the validation error returned by
// CreateTemplateV1Response.Validate if the designated constraints aren't met.
type CreateTemplateV1ResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateTemplateV1ResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateTemplateV1ResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateTemplateV1ResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateTemplateV1ResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateTemplateV1ResponseValidationError) ErrorName() string {
	return "CreateTemplateV1ResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CreateTemplateV1ResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateTemplateV1Response.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateTemplateV1ResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateTemplateV1ResponseValidationError{}
