// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: tests/plugin/validate.proto

package plugin

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

// Validate checks the field values on Interface with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Interface) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Interface with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in InterfaceMultiError, or nil
// if none found.
func (m *Interface) ValidateAll() error {
	return m.validate(true)
}

func (m *Interface) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Name

	if _, ok := _Interface_Status_NotInLookup[m.GetStatus()]; ok {
		err := InterfaceValidationError{
			field:  "Status",
			reason: "value must not be in list [0]",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if _, ok := InterfaceStatus_name[int32(m.GetStatus())]; !ok {
		err := InterfaceValidationError{
			field:  "Status",
			reason: "value must be one of the defined enum values",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	for idx, item := range m.GetAddresses() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, InterfaceValidationError{
						field:  fmt.Sprintf("Addresses[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, InterfaceValidationError{
						field:  fmt.Sprintf("Addresses[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return InterfaceValidationError{
					field:  fmt.Sprintf("Addresses[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return InterfaceMultiError(errors)
	}
	return nil
}

// InterfaceMultiError is an error wrapping multiple validation errors returned
// by Interface.ValidateAll() if the designated constraints aren't met.
type InterfaceMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m InterfaceMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m InterfaceMultiError) AllErrors() []error { return m }

// InterfaceValidationError is the validation error returned by
// Interface.Validate if the designated constraints aren't met.
type InterfaceValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e InterfaceValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e InterfaceValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e InterfaceValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e InterfaceValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e InterfaceValidationError) ErrorName() string { return "InterfaceValidationError" }

// Error satisfies the builtin error interface
func (e InterfaceValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sInterface.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = InterfaceValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = InterfaceValidationError{}

var _Interface_Status_NotInLookup = map[InterfaceStatus]struct{}{
	0: {},
}

// Validate checks the field values on IPAddress with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *IPAddress) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on IPAddress with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in IPAddressMultiError, or nil
// if none found.
func (m *IPAddress) ValidateAll() error {
	return m.validate(true)
}

func (m *IPAddress) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	switch m.Address.(type) {

	case *IPAddress_IPV4:

		if ip := net.ParseIP(m.GetIPV4()); ip == nil || ip.To4() == nil {
			err := IPAddressValidationError{
				field:  "Ipv4",
				reason: "value must be a valid IPv4 address",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	case *IPAddress_IPV6:

		if ip := net.ParseIP(m.GetIPV6()); ip == nil || ip.To4() != nil {
			err := IPAddressValidationError{
				field:  "Ipv6",
				reason: "value must be a valid IPv6 address",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if len(errors) > 0 {
		return IPAddressMultiError(errors)
	}
	return nil
}

// IPAddressMultiError is an error wrapping multiple validation errors returned
// by IPAddress.ValidateAll() if the designated constraints aren't met.
type IPAddressMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m IPAddressMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m IPAddressMultiError) AllErrors() []error { return m }

// IPAddressValidationError is the validation error returned by
// IPAddress.Validate if the designated constraints aren't met.
type IPAddressValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e IPAddressValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e IPAddressValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e IPAddressValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e IPAddressValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e IPAddressValidationError) ErrorName() string { return "IPAddressValidationError" }

// Error satisfies the builtin error interface
func (e IPAddressValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sIPAddress.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = IPAddressValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = IPAddressValidationError{}
