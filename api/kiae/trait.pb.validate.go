// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: kiae/trait.proto

package kiae

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

// Validate checks the field values on Trait with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Trait) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Trait with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in TraitMultiError, or nil if none found.
func (m *Trait) ValidateAll() error {
	return m.validate(true)
}

func (m *Trait) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Name

	// no validation rules for Type

	{
		sorted_keys := make([]string, len(m.GetProperties()))
		i := 0
		for key := range m.GetProperties() {
			sorted_keys[i] = key
			i++
		}
		sort.Slice(sorted_keys, func(i, j int) bool { return sorted_keys[i] < sorted_keys[j] })
		for _, key := range sorted_keys {
			val := m.GetProperties()[key]
			_ = val

			// no validation rules for Properties[key]

			if all {
				switch v := interface{}(val).(type) {
				case interface{ ValidateAll() error }:
					if err := v.ValidateAll(); err != nil {
						errors = append(errors, TraitValidationError{
							field:  fmt.Sprintf("Properties[%v]", key),
							reason: "embedded message failed validation",
							cause:  err,
						})
					}
				case interface{ Validate() error }:
					if err := v.Validate(); err != nil {
						errors = append(errors, TraitValidationError{
							field:  fmt.Sprintf("Properties[%v]", key),
							reason: "embedded message failed validation",
							cause:  err,
						})
					}
				}
			} else if v, ok := interface{}(val).(interface{ Validate() error }); ok {
				if err := v.Validate(); err != nil {
					return TraitValidationError{
						field:  fmt.Sprintf("Properties[%v]", key),
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		}
	}

	if len(errors) > 0 {
		return TraitMultiError(errors)
	}

	return nil
}

// TraitMultiError is an error wrapping multiple validation errors returned by
// Trait.ValidateAll() if the designated constraints aren't met.
type TraitMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m TraitMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m TraitMultiError) AllErrors() []error { return m }

// TraitValidationError is the validation error returned by Trait.Validate if
// the designated constraints aren't met.
type TraitValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TraitValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TraitValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TraitValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TraitValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TraitValidationError) ErrorName() string { return "TraitValidationError" }

// Error satisfies the builtin error interface
func (e TraitValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTrait.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TraitValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TraitValidationError{}
