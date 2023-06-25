// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: auth_room.proto

package go_protocol

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

// Validate checks the field values on ActiveRoomInfo with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ActiveRoomInfo) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ActiveRoomInfo with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ActiveRoomInfoMultiError,
// or nil if none found.
func (m *ActiveRoomInfo) ValidateAll() error {
	return m.validate(true)
}

func (m *ActiveRoomInfo) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for RoomTitle

	// no validation rules for RoomId

	// no validation rules for Sid

	// no validation rules for NumOfJoinedParticipants

	// no validation rules for WebhookUrl

	// no validation rules for Metadata

	// no validation rules for CreatedAt

	// no validation rules for ParentRoomId

	if len(errors) > 0 {
		return ActiveRoomInfoMultiError(errors)
	}

	return nil
}

// ActiveRoomInfoMultiError is an error wrapping multiple validation errors
// returned by ActiveRoomInfo.ValidateAll() if the designated constraints
// aren't met.
type ActiveRoomInfoMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ActiveRoomInfoMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ActiveRoomInfoMultiError) AllErrors() []error { return m }

// ActiveRoomInfoValidationError is the validation error returned by
// ActiveRoomInfo.Validate if the designated constraints aren't met.
type ActiveRoomInfoValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ActiveRoomInfoValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ActiveRoomInfoValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ActiveRoomInfoValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ActiveRoomInfoValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ActiveRoomInfoValidationError) ErrorName() string { return "ActiveRoomInfoValidationError" }

// Error satisfies the builtin error interface
func (e ActiveRoomInfoValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sActiveRoomInfo.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ActiveRoomInfoValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ActiveRoomInfoValidationError{}

// Validate checks the field values on ActiveRoomWithParticipants with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ActiveRoomWithParticipants) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ActiveRoomWithParticipants with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ActiveRoomWithParticipantsMultiError, or nil if none found.
func (m *ActiveRoomWithParticipants) ValidateAll() error {
	return m.validate(true)
}

func (m *ActiveRoomWithParticipants) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return ActiveRoomWithParticipantsMultiError(errors)
	}

	return nil
}

// ActiveRoomWithParticipantsMultiError is an error wrapping multiple
// validation errors returned by ActiveRoomWithParticipants.ValidateAll() if
// the designated constraints aren't met.
type ActiveRoomWithParticipantsMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ActiveRoomWithParticipantsMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ActiveRoomWithParticipantsMultiError) AllErrors() []error { return m }

// ActiveRoomWithParticipantsValidationError is the validation error returned
// by ActiveRoomWithParticipants.Validate if the designated constraints aren't met.
type ActiveRoomWithParticipantsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ActiveRoomWithParticipantsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ActiveRoomWithParticipantsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ActiveRoomWithParticipantsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ActiveRoomWithParticipantsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ActiveRoomWithParticipantsValidationError) ErrorName() string {
	return "ActiveRoomWithParticipantsValidationError"
}

// Error satisfies the builtin error interface
func (e ActiveRoomWithParticipantsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sActiveRoomWithParticipants.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ActiveRoomWithParticipantsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ActiveRoomWithParticipantsValidationError{}

// Validate checks the field values on ActiveRoomInfoRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ActiveRoomInfoRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ActiveRoomInfoRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ActiveRoomInfoRequestMultiError, or nil if none found.
func (m *ActiveRoomInfoRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ActiveRoomInfoRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for RoomId

	if len(errors) > 0 {
		return ActiveRoomInfoRequestMultiError(errors)
	}

	return nil
}

// ActiveRoomInfoRequestMultiError is an error wrapping multiple validation
// errors returned by ActiveRoomInfoRequest.ValidateAll() if the designated
// constraints aren't met.
type ActiveRoomInfoRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ActiveRoomInfoRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ActiveRoomInfoRequestMultiError) AllErrors() []error { return m }

// ActiveRoomInfoRequestValidationError is the validation error returned by
// ActiveRoomInfoRequest.Validate if the designated constraints aren't met.
type ActiveRoomInfoRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ActiveRoomInfoRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ActiveRoomInfoRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ActiveRoomInfoRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ActiveRoomInfoRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ActiveRoomInfoRequestValidationError) ErrorName() string {
	return "ActiveRoomInfoRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ActiveRoomInfoRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sActiveRoomInfoRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ActiveRoomInfoRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ActiveRoomInfoRequestValidationError{}

// Validate checks the field values on ActiveRoomInfoResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ActiveRoomInfoResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ActiveRoomInfoResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ActiveRoomInfoResponseMultiError, or nil if none found.
func (m *ActiveRoomInfoResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ActiveRoomInfoResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Status

	// no validation rules for Msg

	for idx, item := range m.GetParticipantInfo() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ActiveRoomInfoResponseValidationError{
						field:  fmt.Sprintf("ParticipantInfo[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ActiveRoomInfoResponseValidationError{
						field:  fmt.Sprintf("ParticipantInfo[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ActiveRoomInfoResponseValidationError{
					field:  fmt.Sprintf("ParticipantInfo[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if m.RoomInfo != nil {

		if all {
			switch v := interface{}(m.GetRoomInfo()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ActiveRoomInfoResponseValidationError{
						field:  "RoomInfo",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ActiveRoomInfoResponseValidationError{
						field:  "RoomInfo",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetRoomInfo()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ActiveRoomInfoResponseValidationError{
					field:  "RoomInfo",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return ActiveRoomInfoResponseMultiError(errors)
	}

	return nil
}

// ActiveRoomInfoResponseMultiError is an error wrapping multiple validation
// errors returned by ActiveRoomInfoResponse.ValidateAll() if the designated
// constraints aren't met.
type ActiveRoomInfoResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ActiveRoomInfoResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ActiveRoomInfoResponseMultiError) AllErrors() []error { return m }

// ActiveRoomInfoResponseValidationError is the validation error returned by
// ActiveRoomInfoResponse.Validate if the designated constraints aren't met.
type ActiveRoomInfoResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ActiveRoomInfoResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ActiveRoomInfoResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ActiveRoomInfoResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ActiveRoomInfoResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ActiveRoomInfoResponseValidationError) ErrorName() string {
	return "ActiveRoomInfoResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ActiveRoomInfoResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sActiveRoomInfoResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ActiveRoomInfoResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ActiveRoomInfoResponseValidationError{}

// Validate checks the field values on IsRoomActiveRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *IsRoomActiveRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on IsRoomActiveRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// IsRoomActiveRequestMultiError, or nil if none found.
func (m *IsRoomActiveRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *IsRoomActiveRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for RoomId

	if len(errors) > 0 {
		return IsRoomActiveRequestMultiError(errors)
	}

	return nil
}

// IsRoomActiveRequestMultiError is an error wrapping multiple validation
// errors returned by IsRoomActiveRequest.ValidateAll() if the designated
// constraints aren't met.
type IsRoomActiveRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m IsRoomActiveRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m IsRoomActiveRequestMultiError) AllErrors() []error { return m }

// IsRoomActiveRequestValidationError is the validation error returned by
// IsRoomActiveRequest.Validate if the designated constraints aren't met.
type IsRoomActiveRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e IsRoomActiveRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e IsRoomActiveRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e IsRoomActiveRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e IsRoomActiveRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e IsRoomActiveRequestValidationError) ErrorName() string {
	return "IsRoomActiveRequestValidationError"
}

// Error satisfies the builtin error interface
func (e IsRoomActiveRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sIsRoomActiveRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = IsRoomActiveRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = IsRoomActiveRequestValidationError{}

// Validate checks the field values on IsRoomActiveResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *IsRoomActiveResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on IsRoomActiveResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// IsRoomActiveResponseMultiError, or nil if none found.
func (m *IsRoomActiveResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *IsRoomActiveResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return IsRoomActiveResponseMultiError(errors)
	}

	return nil
}

// IsRoomActiveResponseMultiError is an error wrapping multiple validation
// errors returned by IsRoomActiveResponse.ValidateAll() if the designated
// constraints aren't met.
type IsRoomActiveResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m IsRoomActiveResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m IsRoomActiveResponseMultiError) AllErrors() []error { return m }

// IsRoomActiveResponseValidationError is the validation error returned by
// IsRoomActiveResponse.Validate if the designated constraints aren't met.
type IsRoomActiveResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e IsRoomActiveResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e IsRoomActiveResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e IsRoomActiveResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e IsRoomActiveResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e IsRoomActiveResponseValidationError) ErrorName() string {
	return "IsRoomActiveResponseValidationError"
}

// Error satisfies the builtin error interface
func (e IsRoomActiveResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sIsRoomActiveResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = IsRoomActiveResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = IsRoomActiveResponseValidationError{}

// Validate checks the field values on EndRoomRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *EndRoomRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on EndRoomRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in EndRoomRequestMultiError,
// or nil if none found.
func (m *EndRoomRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *EndRoomRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for RoomId

	if len(errors) > 0 {
		return EndRoomRequestMultiError(errors)
	}

	return nil
}

// EndRoomRequestMultiError is an error wrapping multiple validation errors
// returned by EndRoomRequest.ValidateAll() if the designated constraints
// aren't met.
type EndRoomRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m EndRoomRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m EndRoomRequestMultiError) AllErrors() []error { return m }

// EndRoomRequestValidationError is the validation error returned by
// EndRoomRequest.Validate if the designated constraints aren't met.
type EndRoomRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e EndRoomRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e EndRoomRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e EndRoomRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e EndRoomRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e EndRoomRequestValidationError) ErrorName() string { return "EndRoomRequestValidationError" }

// Error satisfies the builtin error interface
func (e EndRoomRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sEndRoomRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = EndRoomRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = EndRoomRequestValidationError{}
