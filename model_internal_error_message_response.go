/**
 * Go SDK for OpenFGA
 *
 * API version: 1.x
 * Website: https://openfga.dev
 * Documentation: https://openfga.dev/docs
 * Support: https://openfga.dev/community
 * License: [Apache-2.0](https://github.com/openfga/go-sdk/blob/main/LICENSE)
 *
 * NOTE: This file was auto generated by OpenAPI Generator (https://openapi-generator.tech). DO NOT EDIT.
 */

package openfga

import (
	"bytes"

	"encoding/json"
)

// InternalErrorMessageResponse struct for InternalErrorMessageResponse
type InternalErrorMessageResponse struct {
	Code    *InternalErrorCode `json:"code,omitempty"yaml:"code,omitempty"`
	Message *string            `json:"message,omitempty"yaml:"message,omitempty"`
}

// NewInternalErrorMessageResponse instantiates a new InternalErrorMessageResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInternalErrorMessageResponse() *InternalErrorMessageResponse {
	this := InternalErrorMessageResponse{}
	var code InternalErrorCode = INTERNALERRORCODE_NO_INTERNAL_ERROR
	this.Code = &code
	return &this
}

// NewInternalErrorMessageResponseWithDefaults instantiates a new InternalErrorMessageResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInternalErrorMessageResponseWithDefaults() *InternalErrorMessageResponse {
	this := InternalErrorMessageResponse{}
	var code InternalErrorCode = INTERNALERRORCODE_NO_INTERNAL_ERROR
	this.Code = &code
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *InternalErrorMessageResponse) GetCode() InternalErrorCode {
	if o == nil || o.Code == nil {
		var ret InternalErrorCode
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalErrorMessageResponse) GetCodeOk() (*InternalErrorCode, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *InternalErrorMessageResponse) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given InternalErrorCode and assigns it to the Code field.
func (o *InternalErrorMessageResponse) SetCode(v InternalErrorCode) {
	o.Code = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *InternalErrorMessageResponse) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalErrorMessageResponse) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *InternalErrorMessageResponse) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *InternalErrorMessageResponse) SetMessage(v string) {
	o.Message = &v
}

func (o InternalErrorMessageResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Code != nil {
		toSerialize["code"] = o.Code
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	var b bytes.Buffer
	enc := json.NewEncoder(&b)
	enc.SetEscapeHTML(false)
	err := enc.Encode(toSerialize)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

type NullableInternalErrorMessageResponse struct {
	value *InternalErrorMessageResponse
	isSet bool
}

func (v NullableInternalErrorMessageResponse) Get() *InternalErrorMessageResponse {
	return v.value
}

func (v *NullableInternalErrorMessageResponse) Set(val *InternalErrorMessageResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableInternalErrorMessageResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableInternalErrorMessageResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInternalErrorMessageResponse(val *InternalErrorMessageResponse) *NullableInternalErrorMessageResponse {
	return &NullableInternalErrorMessageResponse{value: val, isSet: true}
}

func (v NullableInternalErrorMessageResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInternalErrorMessageResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}