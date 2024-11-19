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
	"time"
)

// CreateStoreResponse struct for CreateStoreResponse
type CreateStoreResponse struct {
	Id        string    `json:"id"yaml:"id"`
	Name      string    `json:"name"yaml:"name"`
	CreatedAt time.Time `json:"created_at"yaml:"created_at"`
	UpdatedAt time.Time `json:"updated_at"yaml:"updated_at"`
}

// NewCreateStoreResponse instantiates a new CreateStoreResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateStoreResponse(id string, name string, createdAt time.Time, updatedAt time.Time) *CreateStoreResponse {
	this := CreateStoreResponse{}
	this.Id = id
	this.Name = name
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	return &this
}

// NewCreateStoreResponseWithDefaults instantiates a new CreateStoreResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateStoreResponseWithDefaults() *CreateStoreResponse {
	this := CreateStoreResponse{}
	return &this
}

// GetId returns the Id field value
func (o *CreateStoreResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CreateStoreResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CreateStoreResponse) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *CreateStoreResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateStoreResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateStoreResponse) SetName(v string) {
	o.Name = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *CreateStoreResponse) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *CreateStoreResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *CreateStoreResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *CreateStoreResponse) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *CreateStoreResponse) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *CreateStoreResponse) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

func (o CreateStoreResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["updated_at"] = o.UpdatedAt
	var b bytes.Buffer
	enc := json.NewEncoder(&b)
	enc.SetEscapeHTML(false)
	err := enc.Encode(toSerialize)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

type NullableCreateStoreResponse struct {
	value *CreateStoreResponse
	isSet bool
}

func (v NullableCreateStoreResponse) Get() *CreateStoreResponse {
	return v.value
}

func (v *NullableCreateStoreResponse) Set(val *CreateStoreResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateStoreResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateStoreResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateStoreResponse(val *CreateStoreResponse) *NullableCreateStoreResponse {
	return &NullableCreateStoreResponse{value: val, isSet: true}
}

func (v NullableCreateStoreResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateStoreResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}