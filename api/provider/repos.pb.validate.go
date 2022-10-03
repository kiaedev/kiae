// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: provider/repos.proto

package provider

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

// Validate checks the field values on ListReposRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ListReposRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListReposRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListReposRequestMultiError, or nil if none found.
func (m *ListReposRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ListReposRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Provider

	if len(errors) > 0 {
		return ListReposRequestMultiError(errors)
	}

	return nil
}

// ListReposRequestMultiError is an error wrapping multiple validation errors
// returned by ListReposRequest.ValidateAll() if the designated constraints
// aren't met.
type ListReposRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListReposRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListReposRequestMultiError) AllErrors() []error { return m }

// ListReposRequestValidationError is the validation error returned by
// ListReposRequest.Validate if the designated constraints aren't met.
type ListReposRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListReposRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListReposRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListReposRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListReposRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListReposRequestValidationError) ErrorName() string { return "ListReposRequestValidationError" }

// Error satisfies the builtin error interface
func (e ListReposRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListReposRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListReposRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListReposRequestValidationError{}

// Validate checks the field values on ListReposResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ListReposResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListReposResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListReposResponseMultiError, or nil if none found.
func (m *ListReposResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ListReposResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetItems() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ListReposResponseValidationError{
						field:  fmt.Sprintf("Items[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ListReposResponseValidationError{
						field:  fmt.Sprintf("Items[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListReposResponseValidationError{
					field:  fmt.Sprintf("Items[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for Total

	if len(errors) > 0 {
		return ListReposResponseMultiError(errors)
	}

	return nil
}

// ListReposResponseMultiError is an error wrapping multiple validation errors
// returned by ListReposResponse.ValidateAll() if the designated constraints
// aren't met.
type ListReposResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListReposResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListReposResponseMultiError) AllErrors() []error { return m }

// ListReposResponseValidationError is the validation error returned by
// ListReposResponse.Validate if the designated constraints aren't met.
type ListReposResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListReposResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListReposResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListReposResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListReposResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListReposResponseValidationError) ErrorName() string {
	return "ListReposResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ListReposResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListReposResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListReposResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListReposResponseValidationError{}

// Validate checks the field values on Repo with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *Repo) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Repo with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in RepoMultiError, or nil if none found.
func (m *Repo) ValidateAll() error {
	return m.validate(true)
}

func (m *Repo) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Name

	// no validation rules for FullName

	// no validation rules for Intro

	// no validation rules for GitUrl

	// no validation rules for HttpUrl

	if all {
		switch v := interface{}(m.GetCreatedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, RepoValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, RepoValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCreatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RepoValidationError{
				field:  "CreatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetUpdatedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, RepoValidationError{
					field:  "UpdatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, RepoValidationError{
					field:  "UpdatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUpdatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RepoValidationError{
				field:  "UpdatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return RepoMultiError(errors)
	}

	return nil
}

// RepoMultiError is an error wrapping multiple validation errors returned by
// Repo.ValidateAll() if the designated constraints aren't met.
type RepoMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RepoMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RepoMultiError) AllErrors() []error { return m }

// RepoValidationError is the validation error returned by Repo.Validate if the
// designated constraints aren't met.
type RepoValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RepoValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RepoValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RepoValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RepoValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RepoValidationError) ErrorName() string { return "RepoValidationError" }

// Error satisfies the builtin error interface
func (e RepoValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRepo.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RepoValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RepoValidationError{}

// Validate checks the field values on ListBranchesRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ListBranchesRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListBranchesRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListBranchesRequestMultiError, or nil if none found.
func (m *ListBranchesRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ListBranchesRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Provider

	// no validation rules for RepoName

	if len(errors) > 0 {
		return ListBranchesRequestMultiError(errors)
	}

	return nil
}

// ListBranchesRequestMultiError is an error wrapping multiple validation
// errors returned by ListBranchesRequest.ValidateAll() if the designated
// constraints aren't met.
type ListBranchesRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListBranchesRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListBranchesRequestMultiError) AllErrors() []error { return m }

// ListBranchesRequestValidationError is the validation error returned by
// ListBranchesRequest.Validate if the designated constraints aren't met.
type ListBranchesRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListBranchesRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListBranchesRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListBranchesRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListBranchesRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListBranchesRequestValidationError) ErrorName() string {
	return "ListBranchesRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ListBranchesRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListBranchesRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListBranchesRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListBranchesRequestValidationError{}

// Validate checks the field values on ListBranchesResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ListBranchesResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListBranchesResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListBranchesResponseMultiError, or nil if none found.
func (m *ListBranchesResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ListBranchesResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetItems() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ListBranchesResponseValidationError{
						field:  fmt.Sprintf("Items[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ListBranchesResponseValidationError{
						field:  fmt.Sprintf("Items[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListBranchesResponseValidationError{
					field:  fmt.Sprintf("Items[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for Total

	if len(errors) > 0 {
		return ListBranchesResponseMultiError(errors)
	}

	return nil
}

// ListBranchesResponseMultiError is an error wrapping multiple validation
// errors returned by ListBranchesResponse.ValidateAll() if the designated
// constraints aren't met.
type ListBranchesResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListBranchesResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListBranchesResponseMultiError) AllErrors() []error { return m }

// ListBranchesResponseValidationError is the validation error returned by
// ListBranchesResponse.Validate if the designated constraints aren't met.
type ListBranchesResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListBranchesResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListBranchesResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListBranchesResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListBranchesResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListBranchesResponseValidationError) ErrorName() string {
	return "ListBranchesResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ListBranchesResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListBranchesResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListBranchesResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListBranchesResponseValidationError{}

// Validate checks the field values on Branch with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Branch) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Branch with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in BranchMultiError, or nil if none found.
func (m *Branch) ValidateAll() error {
	return m.validate(true)
}

func (m *Branch) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Name

	if all {
		switch v := interface{}(m.GetCommit()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, BranchValidationError{
					field:  "Commit",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, BranchValidationError{
					field:  "Commit",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCommit()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return BranchValidationError{
				field:  "Commit",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return BranchMultiError(errors)
	}

	return nil
}

// BranchMultiError is an error wrapping multiple validation errors returned by
// Branch.ValidateAll() if the designated constraints aren't met.
type BranchMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m BranchMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m BranchMultiError) AllErrors() []error { return m }

// BranchValidationError is the validation error returned by Branch.Validate if
// the designated constraints aren't met.
type BranchValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e BranchValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e BranchValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e BranchValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e BranchValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e BranchValidationError) ErrorName() string { return "BranchValidationError" }

// Error satisfies the builtin error interface
func (e BranchValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sBranch.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = BranchValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = BranchValidationError{}

// Validate checks the field values on ListTagsRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ListTagsRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListTagsRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListTagsRequestMultiError, or nil if none found.
func (m *ListTagsRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ListTagsRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Provider

	// no validation rules for RepoName

	if len(errors) > 0 {
		return ListTagsRequestMultiError(errors)
	}

	return nil
}

// ListTagsRequestMultiError is an error wrapping multiple validation errors
// returned by ListTagsRequest.ValidateAll() if the designated constraints
// aren't met.
type ListTagsRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListTagsRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListTagsRequestMultiError) AllErrors() []error { return m }

// ListTagsRequestValidationError is the validation error returned by
// ListTagsRequest.Validate if the designated constraints aren't met.
type ListTagsRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListTagsRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListTagsRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListTagsRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListTagsRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListTagsRequestValidationError) ErrorName() string { return "ListTagsRequestValidationError" }

// Error satisfies the builtin error interface
func (e ListTagsRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListTagsRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListTagsRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListTagsRequestValidationError{}

// Validate checks the field values on ListTagsResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ListTagsResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListTagsResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListTagsResponseMultiError, or nil if none found.
func (m *ListTagsResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ListTagsResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetItems() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ListTagsResponseValidationError{
						field:  fmt.Sprintf("Items[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ListTagsResponseValidationError{
						field:  fmt.Sprintf("Items[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListTagsResponseValidationError{
					field:  fmt.Sprintf("Items[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for Total

	if len(errors) > 0 {
		return ListTagsResponseMultiError(errors)
	}

	return nil
}

// ListTagsResponseMultiError is an error wrapping multiple validation errors
// returned by ListTagsResponse.ValidateAll() if the designated constraints
// aren't met.
type ListTagsResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListTagsResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListTagsResponseMultiError) AllErrors() []error { return m }

// ListTagsResponseValidationError is the validation error returned by
// ListTagsResponse.Validate if the designated constraints aren't met.
type ListTagsResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListTagsResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListTagsResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListTagsResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListTagsResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListTagsResponseValidationError) ErrorName() string { return "ListTagsResponseValidationError" }

// Error satisfies the builtin error interface
func (e ListTagsResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListTagsResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListTagsResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListTagsResponseValidationError{}

// Validate checks the field values on Tag with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *Tag) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Tag with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in TagMultiError, or nil if none found.
func (m *Tag) ValidateAll() error {
	return m.validate(true)
}

func (m *Tag) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Name

	if all {
		switch v := interface{}(m.GetCommit()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, TagValidationError{
					field:  "Commit",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, TagValidationError{
					field:  "Commit",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCommit()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return TagValidationError{
				field:  "Commit",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return TagMultiError(errors)
	}

	return nil
}

// TagMultiError is an error wrapping multiple validation errors returned by
// Tag.ValidateAll() if the designated constraints aren't met.
type TagMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m TagMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m TagMultiError) AllErrors() []error { return m }

// TagValidationError is the validation error returned by Tag.Validate if the
// designated constraints aren't met.
type TagValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TagValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TagValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TagValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TagValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TagValidationError) ErrorName() string { return "TagValidationError" }

// Error satisfies the builtin error interface
func (e TagValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTag.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TagValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TagValidationError{}

// Validate checks the field values on Commit with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Commit) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Commit with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in CommitMultiError, or nil if none found.
func (m *Commit) ValidateAll() error {
	return m.validate(true)
}

func (m *Commit) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Sha1

	// no validation rules for ShortId

	// no validation rules for Message

	// no validation rules for CommitterName

	// no validation rules for CommitterEmail

	if all {
		switch v := interface{}(m.GetCreatedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CommitValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CommitValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCreatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CommitValidationError{
				field:  "CreatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetUpdatedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CommitValidationError{
					field:  "UpdatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CommitValidationError{
					field:  "UpdatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUpdatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CommitValidationError{
				field:  "UpdatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return CommitMultiError(errors)
	}

	return nil
}

// CommitMultiError is an error wrapping multiple validation errors returned by
// Commit.ValidateAll() if the designated constraints aren't met.
type CommitMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CommitMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CommitMultiError) AllErrors() []error { return m }

// CommitValidationError is the validation error returned by Commit.Validate if
// the designated constraints aren't met.
type CommitValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CommitValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CommitValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CommitValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CommitValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CommitValidationError) ErrorName() string { return "CommitValidationError" }

// Error satisfies the builtin error interface
func (e CommitValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCommit.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CommitValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CommitValidationError{}