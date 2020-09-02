// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/extensions/filters/udp/udp_proxy/v3/udp_proxy.proto

package envoy_extensions_filters_udp_udp_proxy_v3

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

	"github.com/golang/protobuf/ptypes"
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
	_ = ptypes.DynamicAny{}
)

// define the regex for a UUID once up-front
var _udp_proxy_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on UdpProxyConfig with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *UdpProxyConfig) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetStatPrefix()) < 1 {
		return UdpProxyConfigValidationError{
			field:  "StatPrefix",
			reason: "value length must be at least 1 bytes",
		}
	}

	if v, ok := interface{}(m.GetIdleTimeout()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UdpProxyConfigValidationError{
				field:  "IdleTimeout",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for UseOriginalSrcIp

	if len(m.GetHashPolicies()) > 1 {
		return UdpProxyConfigValidationError{
			field:  "HashPolicies",
			reason: "value must contain no more than 1 item(s)",
		}
	}

	for idx, item := range m.GetHashPolicies() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return UdpProxyConfigValidationError{
					field:  fmt.Sprintf("HashPolicies[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	switch m.RouteSpecifier.(type) {

	case *UdpProxyConfig_Cluster:

		if len(m.GetCluster()) < 1 {
			return UdpProxyConfigValidationError{
				field:  "Cluster",
				reason: "value length must be at least 1 bytes",
			}
		}

	default:
		return UdpProxyConfigValidationError{
			field:  "RouteSpecifier",
			reason: "value is required",
		}

	}

	return nil
}

// UdpProxyConfigValidationError is the validation error returned by
// UdpProxyConfig.Validate if the designated constraints aren't met.
type UdpProxyConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UdpProxyConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UdpProxyConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UdpProxyConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UdpProxyConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UdpProxyConfigValidationError) ErrorName() string { return "UdpProxyConfigValidationError" }

// Error satisfies the builtin error interface
func (e UdpProxyConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUdpProxyConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UdpProxyConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UdpProxyConfigValidationError{}

// Validate checks the field values on UdpProxyConfig_HashPolicy with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *UdpProxyConfig_HashPolicy) Validate() error {
	if m == nil {
		return nil
	}

	switch m.PolicySpecifier.(type) {

	case *UdpProxyConfig_HashPolicy_SourceIp:

		if m.GetSourceIp() != true {
			return UdpProxyConfig_HashPolicyValidationError{
				field:  "SourceIp",
				reason: "value must equal true",
			}
		}

	default:
		return UdpProxyConfig_HashPolicyValidationError{
			field:  "PolicySpecifier",
			reason: "value is required",
		}

	}

	return nil
}

// UdpProxyConfig_HashPolicyValidationError is the validation error returned by
// UdpProxyConfig_HashPolicy.Validate if the designated constraints aren't met.
type UdpProxyConfig_HashPolicyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UdpProxyConfig_HashPolicyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UdpProxyConfig_HashPolicyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UdpProxyConfig_HashPolicyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UdpProxyConfig_HashPolicyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UdpProxyConfig_HashPolicyValidationError) ErrorName() string {
	return "UdpProxyConfig_HashPolicyValidationError"
}

// Error satisfies the builtin error interface
func (e UdpProxyConfig_HashPolicyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUdpProxyConfig_HashPolicy.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UdpProxyConfig_HashPolicyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UdpProxyConfig_HashPolicyValidationError{}
