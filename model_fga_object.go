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

// FgaObject Object represents an OpenFGA Object.  An Object is composed of a type and identifier (e.g. 'document:1')  See https://openfga.dev/docs/concepts#what-is-an-object
type FgaObject struct {
	Type string `json:"type"yaml:"type"`
	Id   string `json:"id"yaml:"id"`
}

// NewFgaObject instantiates a new FgaObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFgaObject(type_ string, id string) *FgaObject {
	this := FgaObject{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewFgaObjectWithDefaults instantiates a new FgaObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFgaObjectWithDefaults() *FgaObject {
	this := FgaObject{}
	return &this
}

// GetType returns the Type field value
func (o *FgaObject) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *FgaObject) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *FgaObject) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *FgaObject) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *FgaObject) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *FgaObject) SetId(v string) {
	o.Id = v
}

func (o FgaObject) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	var b bytes.Buffer
	enc := json.NewEncoder(&b)
	enc.SetEscapeHTML(false)
	err := enc.Encode(toSerialize)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

type NullableFgaObject struct {
	value *FgaObject
	isSet bool
}

func (v NullableFgaObject) Get() *FgaObject {
	return v.value
}

func (v *NullableFgaObject) Set(val *FgaObject) {
	v.value = val
	v.isSet = true
}

func (v NullableFgaObject) IsSet() bool {
	return v.isSet
}

func (v *NullableFgaObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFgaObject(val *FgaObject) *NullableFgaObject {
	return &NullableFgaObject{value: val, isSet: true}
}

func (v NullableFgaObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFgaObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
