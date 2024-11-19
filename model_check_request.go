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

// CheckRequest struct for CheckRequest
type CheckRequest struct {
	TupleKey             CheckRequestTupleKey `json:"tuple_key"yaml:"tuple_key"`
	ContextualTuples     *ContextualTupleKeys `json:"contextual_tuples,omitempty"yaml:"contextual_tuples,omitempty"`
	AuthorizationModelId *string              `json:"authorization_model_id,omitempty"yaml:"authorization_model_id,omitempty"`
	// Defaults to false. Making it true has performance implications.
	Trace *bool `json:"trace,omitempty"yaml:"trace,omitempty"`
	// Additional request context that will be used to evaluate any ABAC conditions encountered in the query evaluation.
	Context     *map[string]interface{} `json:"context,omitempty"yaml:"context,omitempty"`
	Consistency *ConsistencyPreference  `json:"consistency,omitempty"yaml:"consistency,omitempty"`
}

// NewCheckRequest instantiates a new CheckRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCheckRequest(tupleKey CheckRequestTupleKey) *CheckRequest {
	this := CheckRequest{}
	this.TupleKey = tupleKey
	var consistency ConsistencyPreference = CONSISTENCYPREFERENCE_UNSPECIFIED
	this.Consistency = &consistency
	return &this
}

// NewCheckRequestWithDefaults instantiates a new CheckRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCheckRequestWithDefaults() *CheckRequest {
	this := CheckRequest{}
	var consistency ConsistencyPreference = CONSISTENCYPREFERENCE_UNSPECIFIED
	this.Consistency = &consistency
	return &this
}

// GetTupleKey returns the TupleKey field value
func (o *CheckRequest) GetTupleKey() CheckRequestTupleKey {
	if o == nil {
		var ret CheckRequestTupleKey
		return ret
	}

	return o.TupleKey
}

// GetTupleKeyOk returns a tuple with the TupleKey field value
// and a boolean to check if the value has been set.
func (o *CheckRequest) GetTupleKeyOk() (*CheckRequestTupleKey, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TupleKey, true
}

// SetTupleKey sets field value
func (o *CheckRequest) SetTupleKey(v CheckRequestTupleKey) {
	o.TupleKey = v
}

// GetContextualTuples returns the ContextualTuples field value if set, zero value otherwise.
func (o *CheckRequest) GetContextualTuples() ContextualTupleKeys {
	if o == nil || o.ContextualTuples == nil {
		var ret ContextualTupleKeys
		return ret
	}
	return *o.ContextualTuples
}

// GetContextualTuplesOk returns a tuple with the ContextualTuples field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckRequest) GetContextualTuplesOk() (*ContextualTupleKeys, bool) {
	if o == nil || o.ContextualTuples == nil {
		return nil, false
	}
	return o.ContextualTuples, true
}

// HasContextualTuples returns a boolean if a field has been set.
func (o *CheckRequest) HasContextualTuples() bool {
	if o != nil && o.ContextualTuples != nil {
		return true
	}

	return false
}

// SetContextualTuples gets a reference to the given ContextualTupleKeys and assigns it to the ContextualTuples field.
func (o *CheckRequest) SetContextualTuples(v ContextualTupleKeys) {
	o.ContextualTuples = &v
}

// GetAuthorizationModelId returns the AuthorizationModelId field value if set, zero value otherwise.
func (o *CheckRequest) GetAuthorizationModelId() string {
	if o == nil || o.AuthorizationModelId == nil {
		var ret string
		return ret
	}
	return *o.AuthorizationModelId
}

// GetAuthorizationModelIdOk returns a tuple with the AuthorizationModelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckRequest) GetAuthorizationModelIdOk() (*string, bool) {
	if o == nil || o.AuthorizationModelId == nil {
		return nil, false
	}
	return o.AuthorizationModelId, true
}

// HasAuthorizationModelId returns a boolean if a field has been set.
func (o *CheckRequest) HasAuthorizationModelId() bool {
	if o != nil && o.AuthorizationModelId != nil {
		return true
	}

	return false
}

// SetAuthorizationModelId gets a reference to the given string and assigns it to the AuthorizationModelId field.
func (o *CheckRequest) SetAuthorizationModelId(v string) {
	o.AuthorizationModelId = &v
}

// GetTrace returns the Trace field value if set, zero value otherwise.
func (o *CheckRequest) GetTrace() bool {
	if o == nil || o.Trace == nil {
		var ret bool
		return ret
	}
	return *o.Trace
}

// GetTraceOk returns a tuple with the Trace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckRequest) GetTraceOk() (*bool, bool) {
	if o == nil || o.Trace == nil {
		return nil, false
	}
	return o.Trace, true
}

// HasTrace returns a boolean if a field has been set.
func (o *CheckRequest) HasTrace() bool {
	if o != nil && o.Trace != nil {
		return true
	}

	return false
}

// SetTrace gets a reference to the given bool and assigns it to the Trace field.
func (o *CheckRequest) SetTrace(v bool) {
	o.Trace = &v
}

// GetContext returns the Context field value if set, zero value otherwise.
func (o *CheckRequest) GetContext() map[string]interface{} {
	if o == nil || o.Context == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckRequest) GetContextOk() (*map[string]interface{}, bool) {
	if o == nil || o.Context == nil {
		return nil, false
	}
	return o.Context, true
}

// HasContext returns a boolean if a field has been set.
func (o *CheckRequest) HasContext() bool {
	if o != nil && o.Context != nil {
		return true
	}

	return false
}

// SetContext gets a reference to the given map[string]interface{} and assigns it to the Context field.
func (o *CheckRequest) SetContext(v map[string]interface{}) {
	o.Context = &v
}

// GetConsistency returns the Consistency field value if set, zero value otherwise.
func (o *CheckRequest) GetConsistency() ConsistencyPreference {
	if o == nil || o.Consistency == nil {
		var ret ConsistencyPreference
		return ret
	}
	return *o.Consistency
}

// GetConsistencyOk returns a tuple with the Consistency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckRequest) GetConsistencyOk() (*ConsistencyPreference, bool) {
	if o == nil || o.Consistency == nil {
		return nil, false
	}
	return o.Consistency, true
}

// HasConsistency returns a boolean if a field has been set.
func (o *CheckRequest) HasConsistency() bool {
	if o != nil && o.Consistency != nil {
		return true
	}

	return false
}

// SetConsistency gets a reference to the given ConsistencyPreference and assigns it to the Consistency field.
func (o *CheckRequest) SetConsistency(v ConsistencyPreference) {
	o.Consistency = &v
}

func (o CheckRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["tuple_key"] = o.TupleKey
	if o.ContextualTuples != nil {
		toSerialize["contextual_tuples"] = o.ContextualTuples
	}
	if o.AuthorizationModelId != nil {
		toSerialize["authorization_model_id"] = o.AuthorizationModelId
	}
	if o.Trace != nil {
		toSerialize["trace"] = o.Trace
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

type NullableCheckRequest struct {
	value *CheckRequest
	isSet bool
}

func (v NullableCheckRequest) Get() *CheckRequest {
	return v.value
}

func (v *NullableCheckRequest) Set(val *CheckRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckRequest(val *CheckRequest) *NullableCheckRequest {
	return &NullableCheckRequest{value: val, isSet: true}
}

func (v NullableCheckRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}