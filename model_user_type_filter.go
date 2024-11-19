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

// UserTypeFilter struct for UserTypeFilter
type UserTypeFilter struct {
	Type     string  `json:"type"yaml:"type"`
	Relation *string `json:"relation,omitempty"yaml:"relation,omitempty"`
}

// NewUserTypeFilter instantiates a new UserTypeFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserTypeFilter(type_ string) *UserTypeFilter {
	this := UserTypeFilter{}
	this.Type = type_
	return &this
}

// NewUserTypeFilterWithDefaults instantiates a new UserTypeFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserTypeFilterWithDefaults() *UserTypeFilter {
	this := UserTypeFilter{}
	return &this
}

// GetType returns the Type field value
func (o *UserTypeFilter) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *UserTypeFilter) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *UserTypeFilter) SetType(v string) {
	o.Type = v
}

// GetRelation returns the Relation field value if set, zero value otherwise.
func (o *UserTypeFilter) GetRelation() string {
	if o == nil || o.Relation == nil {
		var ret string
		return ret
	}
	return *o.Relation
}

// GetRelationOk returns a tuple with the Relation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserTypeFilter) GetRelationOk() (*string, bool) {
	if o == nil || o.Relation == nil {
		return nil, false
	}
	return o.Relation, true
}

// HasRelation returns a boolean if a field has been set.
func (o *UserTypeFilter) HasRelation() bool {
	if o != nil && o.Relation != nil {
		return true
	}

	return false
}

// SetRelation gets a reference to the given string and assigns it to the Relation field.
func (o *UserTypeFilter) SetRelation(v string) {
	o.Relation = &v
}

func (o UserTypeFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	if o.Relation != nil {
		toSerialize["relation"] = o.Relation
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

type NullableUserTypeFilter struct {
	value *UserTypeFilter
	isSet bool
}

func (v NullableUserTypeFilter) Get() *UserTypeFilter {
	return v.value
}

func (v *NullableUserTypeFilter) Set(val *UserTypeFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableUserTypeFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableUserTypeFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserTypeFilter(val *UserTypeFilter) *NullableUserTypeFilter {
	return &NullableUserTypeFilter{value: val, isSet: true}
}

func (v NullableUserTypeFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserTypeFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}