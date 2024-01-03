// Code generated by protoc-gen-jsonshim. DO NOT EDIT.
package v1beta1

import (
	bytes "bytes"
	jsonpb "github.com/golang/protobuf/jsonpb"
)

// MarshalJSON is a custom marshaler for VirtualService
func (this *VirtualService) MarshalJSON() ([]byte, error) {
	str, err := VirtualServiceMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for VirtualService
func (this *VirtualService) UnmarshalJSON(b []byte) error {
	return VirtualServiceUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for Destination
func (this *Destination) MarshalJSON() ([]byte, error) {
	str, err := VirtualServiceMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for Destination
func (this *Destination) UnmarshalJSON(b []byte) error {
	return VirtualServiceUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for HTTPRoute
func (this *HTTPRoute) MarshalJSON() ([]byte, error) {
	str, err := VirtualServiceMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for HTTPRoute
func (this *HTTPRoute) UnmarshalJSON(b []byte) error {
	return VirtualServiceUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for Delegate
func (this *Delegate) MarshalJSON() ([]byte, error) {
	str, err := VirtualServiceMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for Delegate
func (this *Delegate) UnmarshalJSON(b []byte) error {
	return VirtualServiceUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for Headers
func (this *Headers) MarshalJSON() ([]byte, error) {
	str, err := VirtualServiceMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for Headers
func (this *Headers) UnmarshalJSON(b []byte) error {
	return VirtualServiceUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for Headers_HeaderOperations
func (this *Headers_HeaderOperations) MarshalJSON() ([]byte, error) {
	str, err := VirtualServiceMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for Headers_HeaderOperations
func (this *Headers_HeaderOperations) UnmarshalJSON(b []byte) error {
	return VirtualServiceUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for TLSRoute
func (this *TLSRoute) MarshalJSON() ([]byte, error) {
	str, err := VirtualServiceMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for TLSRoute
func (this *TLSRoute) UnmarshalJSON(b []byte) error {
	return VirtualServiceUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for TCPRoute
func (this *TCPRoute) MarshalJSON() ([]byte, error) {
	str, err := VirtualServiceMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for TCPRoute
func (this *TCPRoute) UnmarshalJSON(b []byte) error {
	return VirtualServiceUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for HTTPMatchRequest
func (this *HTTPMatchRequest) MarshalJSON() ([]byte, error) {
	str, err := VirtualServiceMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for HTTPMatchRequest
func (this *HTTPMatchRequest) UnmarshalJSON(b []byte) error {
	return VirtualServiceUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for HTTPRouteDestination
func (this *HTTPRouteDestination) MarshalJSON() ([]byte, error) {
	str, err := VirtualServiceMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for HTTPRouteDestination
func (this *HTTPRouteDestination) UnmarshalJSON(b []byte) error {
	return VirtualServiceUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for RouteDestination
func (this *RouteDestination) MarshalJSON() ([]byte, error) {
	str, err := VirtualServiceMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for RouteDestination
func (this *RouteDestination) UnmarshalJSON(b []byte) error {
	return VirtualServiceUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for L4MatchAttributes
func (this *L4MatchAttributes) MarshalJSON() ([]byte, error) {
	str, err := VirtualServiceMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for L4MatchAttributes
func (this *L4MatchAttributes) UnmarshalJSON(b []byte) error {
	return VirtualServiceUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for TLSMatchAttributes
func (this *TLSMatchAttributes) MarshalJSON() ([]byte, error) {
	str, err := VirtualServiceMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for TLSMatchAttributes
func (this *TLSMatchAttributes) UnmarshalJSON(b []byte) error {
	return VirtualServiceUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for HTTPRedirect
func (this *HTTPRedirect) MarshalJSON() ([]byte, error) {
	str, err := VirtualServiceMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for HTTPRedirect
func (this *HTTPRedirect) UnmarshalJSON(b []byte) error {
	return VirtualServiceUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for HTTPDirectResponse
func (this *HTTPDirectResponse) MarshalJSON() ([]byte, error) {
	str, err := VirtualServiceMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for HTTPDirectResponse
func (this *HTTPDirectResponse) UnmarshalJSON(b []byte) error {
	return VirtualServiceUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for HTTPBody
func (this *HTTPBody) MarshalJSON() ([]byte, error) {
	str, err := VirtualServiceMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for HTTPBody
func (this *HTTPBody) UnmarshalJSON(b []byte) error {
	return VirtualServiceUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for HTTPRewrite
func (this *HTTPRewrite) MarshalJSON() ([]byte, error) {
	str, err := VirtualServiceMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for HTTPRewrite
func (this *HTTPRewrite) UnmarshalJSON(b []byte) error {
	return VirtualServiceUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for RegexRewrite
func (this *RegexRewrite) MarshalJSON() ([]byte, error) {
	str, err := VirtualServiceMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for RegexRewrite
func (this *RegexRewrite) UnmarshalJSON(b []byte) error {
	return VirtualServiceUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for StringMatch
func (this *StringMatch) MarshalJSON() ([]byte, error) {
	str, err := VirtualServiceMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for StringMatch
func (this *StringMatch) UnmarshalJSON(b []byte) error {
	return VirtualServiceUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for HTTPRetry
func (this *HTTPRetry) MarshalJSON() ([]byte, error) {
	str, err := VirtualServiceMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for HTTPRetry
func (this *HTTPRetry) UnmarshalJSON(b []byte) error {
	return VirtualServiceUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for CorsPolicy
func (this *CorsPolicy) MarshalJSON() ([]byte, error) {
	str, err := VirtualServiceMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for CorsPolicy
func (this *CorsPolicy) UnmarshalJSON(b []byte) error {
	return VirtualServiceUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for HTTPFaultInjection
func (this *HTTPFaultInjection) MarshalJSON() ([]byte, error) {
	str, err := VirtualServiceMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for HTTPFaultInjection
func (this *HTTPFaultInjection) UnmarshalJSON(b []byte) error {
	return VirtualServiceUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for HTTPFaultInjection_Delay
func (this *HTTPFaultInjection_Delay) MarshalJSON() ([]byte, error) {
	str, err := VirtualServiceMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for HTTPFaultInjection_Delay
func (this *HTTPFaultInjection_Delay) UnmarshalJSON(b []byte) error {
	return VirtualServiceUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for HTTPFaultInjection_Abort
func (this *HTTPFaultInjection_Abort) MarshalJSON() ([]byte, error) {
	str, err := VirtualServiceMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for HTTPFaultInjection_Abort
func (this *HTTPFaultInjection_Abort) UnmarshalJSON(b []byte) error {
	return VirtualServiceUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for HTTPMirrorPolicy
func (this *HTTPMirrorPolicy) MarshalJSON() ([]byte, error) {
	str, err := VirtualServiceMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for HTTPMirrorPolicy
func (this *HTTPMirrorPolicy) UnmarshalJSON(b []byte) error {
	return VirtualServiceUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for PortSelector
func (this *PortSelector) MarshalJSON() ([]byte, error) {
	str, err := VirtualServiceMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for PortSelector
func (this *PortSelector) UnmarshalJSON(b []byte) error {
	return VirtualServiceUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for Percent
func (this *Percent) MarshalJSON() ([]byte, error) {
	str, err := VirtualServiceMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for Percent
func (this *Percent) UnmarshalJSON(b []byte) error {
	return VirtualServiceUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for HTTPInternalActiveRedirect
func (this *HTTPInternalActiveRedirect) MarshalJSON() ([]byte, error) {
	str, err := VirtualServiceMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for HTTPInternalActiveRedirect
func (this *HTTPInternalActiveRedirect) UnmarshalJSON(b []byte) error {
	return VirtualServiceUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for HTTPInternalActiveRedirect_RedirectPolicy
func (this *HTTPInternalActiveRedirect_RedirectPolicy) MarshalJSON() ([]byte, error) {
	str, err := VirtualServiceMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for HTTPInternalActiveRedirect_RedirectPolicy
func (this *HTTPInternalActiveRedirect_RedirectPolicy) UnmarshalJSON(b []byte) error {
	return VirtualServiceUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for RegexMatchAndSubstitute
func (this *RegexMatchAndSubstitute) MarshalJSON() ([]byte, error) {
	str, err := VirtualServiceMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for RegexMatchAndSubstitute
func (this *RegexMatchAndSubstitute) UnmarshalJSON(b []byte) error {
	return VirtualServiceUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

var (
	VirtualServiceMarshaler   = &jsonpb.Marshaler{}
	VirtualServiceUnmarshaler = &jsonpb.Unmarshaler{AllowUnknownFields: true}
)
