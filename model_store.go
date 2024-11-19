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

// Store struct for Store
type Store struct {
	Id        string     `json:"id"yaml:"id"`
	Name      string     `json:"name"yaml:"name"`
	CreatedAt time.Time  `json:"created_at"yaml:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"yaml:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"yaml:"deleted_at,omitempty"`
}

// NewStore instantiates a new Store object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStore(id string, name string, createdAt time.Time, updatedAt time.Time) *Store {
	this := Store{}
	this.Id = id
	this.Name = name
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	return &this
}

// NewStoreWithDefaults instantiates a new Store object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStoreWithDefaults() *Store {
	this := Store{}
	return &this
}

// GetId returns the Id field value
func (o *Store) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Store) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Store) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *Store) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Store) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Store) SetName(v string) {
	o.Name = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *Store) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *Store) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *Store) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *Store) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *Store) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *Store) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise.
func (o *Store) GetDeletedAt() time.Time {
	if o == nil || o.DeletedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.DeletedAt
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Store) GetDeletedAtOk() (*time.Time, bool) {
	if o == nil || o.DeletedAt == nil {
		return nil, false
	}
	return o.DeletedAt, true
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *Store) HasDeletedAt() bool {
	if o != nil && o.DeletedAt != nil {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given time.Time and assigns it to the DeletedAt field.
func (o *Store) SetDeletedAt(v time.Time) {
	o.DeletedAt = &v
}

func (o Store) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["updated_at"] = o.UpdatedAt
	if o.DeletedAt != nil {
		toSerialize["deleted_at"] = o.DeletedAt
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

type NullableStore struct {
	value *Store
	isSet bool
}

func (v NullableStore) Get() *Store {
	return v.value
}

func (v *NullableStore) Set(val *Store) {
	v.value = val
	v.isSet = true
}

func (v NullableStore) IsSet() bool {
	return v.isSet
}

func (v *NullableStore) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStore(val *Store) *NullableStore {
	return &NullableStore{value: val, isSet: true}
}

func (v NullableStore) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStore) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}