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

// ListObjectsRequest struct for ListObjectsRequest
type ListObjectsRequest struct {
	AuthorizationModelId *string              `json:"authorization_model_id,omitempty"yaml:"authorization_model_id,omitempty"`
	Type                 string               `json:"type"yaml:"type"`
	Relation             string               `json:"relation"yaml:"relation"`
	User                 string               `json:"user"yaml:"user"`
	ContextualTuples     *ContextualTupleKeys `json:"contextual_tuples,omitempty"yaml:"contextual_tuples,omitempty"`
	// Additional request context that will be used to evaluate any ABAC conditions encountered in the query evaluation.
	Context     *map[string]interface{} `json:"context,omitempty"yaml:"context,omitempty"`
	Consistency *ConsistencyPreference  `json:"consistency,omitempty"yaml:"consistency,omitempty"`
}

// NewListObjectsRequest instantiates a new ListObjectsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListObjectsRequest(type_ string, relation string, user string) *ListObjectsRequest {
	this := ListObjectsRequest{}
	this.Type = type_
	this.Relation = relation
	this.User = user
	var consistency ConsistencyPreference = CONSISTENCYPREFERENCE_UNSPECIFIED
	this.Consistency = &consistency
	return &this
}

// NewListObjectsRequestWithDefaults instantiates a new ListObjectsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListObjectsRequestWithDefaults() *ListObjectsRequest {
	this := ListObjectsRequest{}
	var consistency ConsistencyPreference = CONSISTENCYPREFERENCE_UNSPECIFIED
	this.Consistency = &consistency
	return &this
}

// GetAuthorizationModelId returns the AuthorizationModelId field value if set, zero value otherwise.
func (o *ListObjectsRequest) GetAuthorizationModelId() string {
	if o == nil || o.AuthorizationModelId == nil {
		var ret string
		return ret
	}
	return *o.AuthorizationModelId
}

// GetAuthorizationModelIdOk returns a tuple with the AuthorizationModelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListObjectsRequest) GetAuthorizationModelIdOk() (*string, bool) {
	if o == nil || o.AuthorizationModelId == nil {
		return nil, false
	}
	return o.AuthorizationModelId, true
}

// HasAuthorizationModelId returns a boolean if a field has been set.
func (o *ListObjectsRequest) HasAuthorizationModelId() bool {
	if o != nil && o.AuthorizationModelId != nil {
		return true
	}

	return false
}

// SetAuthorizationModelId gets a reference to the given string and assigns it to the AuthorizationModelId field.
func (o *ListObjectsRequest) SetAuthorizationModelId(v string) {
	o.AuthorizationModelId = &v
}

// GetType returns the Type field value
func (o *ListObjectsRequest) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ListObjectsRequest) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ListObjectsRequest) SetType(v string) {
	o.Type = v
}

// GetRelation returns the Relation field value
func (o *ListObjectsRequest) GetRelation() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Relation
}

// GetRelationOk returns a tuple with the Relation field value
// and a boolean to check if the value has been set.
func (o *ListObjectsRequest) GetRelationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Relation, true
}

// SetRelation sets field value
func (o *ListObjectsRequest) SetRelation(v string) {
	o.Relation = v
}

// GetUser returns the User field value
func (o *ListObjectsRequest) GetUser() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *ListObjectsRequest) GetUserOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *ListObjectsRequest) SetUser(v string) {
	o.User = v
}

// GetContextualTuples returns the ContextualTuples field value if set, zero value otherwise.
func (o *ListObjectsRequest) GetContextualTuples() ContextualTupleKeys {
	if o == nil || o.ContextualTuples == nil {
		var ret ContextualTupleKeys
		return ret
	}
	return *o.ContextualTuples
}

// GetContextualTuplesOk returns a tuple with the ContextualTuples field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListObjectsRequest) GetContextualTuplesOk() (*ContextualTupleKeys, bool) {
	if o == nil || o.ContextualTuples == nil {
		return nil, false
	}
	return o.ContextualTuples, true
}

// HasContextualTuples returns a boolean if a field has been set.
func (o *ListObjectsRequest) HasContextualTuples() bool {
	if o != nil && o.ContextualTuples != nil {
		return true
	}

	return false
}

// SetContextualTuples gets a reference to the given ContextualTupleKeys and assigns it to the ContextualTuples field.
func (o *ListObjectsRequest) SetContextualTuples(v ContextualTupleKeys) {
	o.ContextualTuples = &v
}

// GetContext returns the Context field value if set, zero value otherwise.
func (o *ListObjectsRequest) GetContext() map[string]interface{} {
	if o == nil || o.Context == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListObjectsRequest) GetContextOk() (*map[string]interface{}, bool) {
	if o == nil || o.Context == nil {
		return nil, false
	}
	return o.Context, true
}

// HasContext returns a boolean if a field has been set.
func (o *ListObjectsRequest) HasContext() bool {
	if o != nil && o.Context != nil {
		return true
	}

	return false
}

// SetContext gets a reference to the given map[string]interface{} and assigns it to the Context field.
func (o *ListObjectsRequest) SetContext(v map[string]interface{}) {
	o.Context = &v
}

// GetConsistency returns the Consistency field value if set, zero value otherwise.
func (o *ListObjectsRequest) GetConsistency() ConsistencyPreference {
	if o == nil || o.Consistency == nil {
		var ret ConsistencyPreference
		return ret
	}
	return *o.Consistency
}

// GetConsistencyOk returns a tuple with the Consistency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListObjectsRequest) GetConsistencyOk() (*ConsistencyPreference, bool) {
	if o == nil || o.Consistency == nil {
		return nil, false
	}
	return o.Consistency, true
}

// HasConsistency returns a boolean if a field has been set.
func (o *ListObjectsRequest) HasConsistency() bool {
	if o != nil && o.Consistency != nil {
		return true
	}

	return false
}

// SetConsistency gets a reference to the given ConsistencyPreference and assigns it to the Consistency field.
func (o *ListObjectsRequest) SetConsistency(v ConsistencyPreference) {
	o.Consistency = &v
}

func (o ListObjectsRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AuthorizationModelId != nil {
		toSerialize["authorization_model_id"] = o.AuthorizationModelId
	}
	toSerialize["type"] = o.Type
	toSerialize["relation"] = o.Relation
	toSerialize["user"] = o.User
	if o.ContextualTuples != nil {
		toSerialize["contextual_tuples"] = o.ContextualTuples
	}
	if o.Context != nil {
		toSerialize["context"] = o.Context
	}
	if o.Consistency != nil {
		toSerialize["consistency"] = o.Consistency
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

type NullableListObjectsRequest struct {
	value *ListObjectsRequest
	isSet bool
}

func (v NullableListObjectsRequest) Get() *ListObjectsRequest {
	return v.value
}

func (v *NullableListObjectsRequest) Set(val *ListObjectsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableListObjectsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableListObjectsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListObjectsRequest(val *ListObjectsRequest) *NullableListObjectsRequest {
	return &NullableListObjectsRequest{value: val, isSet: true}
}

func (v NullableListObjectsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListObjectsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}