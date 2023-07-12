// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/api/v2/core/event_service_config.proto

package core

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

// Validate checks the field values on EventServiceConfig with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *EventServiceConfig) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on EventServiceConfig with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// EventServiceConfigMultiError, or nil if none found.
func (m *EventServiceConfig) ValidateAll() error {
	return m.validate(true)
}

func (m *EventServiceConfig) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	oneofConfigSourceSpecifierPresent := false
	switch v := m.ConfigSourceSpecifier.(type) {
	case *EventServiceConfig_GrpcService:
		if v == nil {
			err := EventServiceConfigValidationError{
				field:  "ConfigSourceSpecifier",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		oneofConfigSourceSpecifierPresent = true

		if all {
			switch v := interface{}(m.GetGrpcService()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, EventServiceConfigValidationError{
						field:  "GrpcService",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, EventServiceConfigValidationError{
						field:  "GrpcService",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetGrpcService()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return EventServiceConfigValidationError{
					field:  "GrpcService",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		_ = v // ensures v is used
	}
	if !oneofConfigSourceSpecifierPresent {
		err := EventServiceConfigValidationError{
			field:  "ConfigSourceSpecifier",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return EventServiceConfigMultiError(errors)
	}

	return nil
}

// EventServiceConfigMultiError is an error wrapping multiple validation errors
// returned by EventServiceConfig.ValidateAll() if the designated constraints
// aren't met.
type EventServiceConfigMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m EventServiceConfigMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m EventServiceConfigMultiError) AllErrors() []error { return m }

// EventServiceConfigValidationError is the validation error returned by
// EventServiceConfig.Validate if the designated constraints aren't met.
type EventServiceConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e EventServiceConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e EventServiceConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e EventServiceConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e EventServiceConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e EventServiceConfigValidationError) ErrorName() string {
	return "EventServiceConfigValidationError"
}

// Error satisfies the builtin error interface
func (e EventServiceConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sEventServiceConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = EventServiceConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = EventServiceConfigValidationError{}
