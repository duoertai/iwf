/*
Workflow APIs

This APIs for iwf SDKs to operate workflows

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package iwfidl

import (
	"encoding/json"
)

// WorkflowSearchResponse struct for WorkflowSearchResponse
type WorkflowSearchResponse struct {
	QueryAttributes []WorkflowSearchResponseEntry `json:"queryAttributes,omitempty"`
}

// NewWorkflowSearchResponse instantiates a new WorkflowSearchResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowSearchResponse() *WorkflowSearchResponse {
	this := WorkflowSearchResponse{}
	return &this
}

// NewWorkflowSearchResponseWithDefaults instantiates a new WorkflowSearchResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowSearchResponseWithDefaults() *WorkflowSearchResponse {
	this := WorkflowSearchResponse{}
	return &this
}

// GetQueryAttributes returns the QueryAttributes field value if set, zero value otherwise.
func (o *WorkflowSearchResponse) GetQueryAttributes() []WorkflowSearchResponseEntry {
	if o == nil || o.QueryAttributes == nil {
		var ret []WorkflowSearchResponseEntry
		return ret
	}
	return o.QueryAttributes
}

// GetQueryAttributesOk returns a tuple with the QueryAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowSearchResponse) GetQueryAttributesOk() ([]WorkflowSearchResponseEntry, bool) {
	if o == nil || o.QueryAttributes == nil {
		return nil, false
	}
	return o.QueryAttributes, true
}

// HasQueryAttributes returns a boolean if a field has been set.
func (o *WorkflowSearchResponse) HasQueryAttributes() bool {
	if o != nil && o.QueryAttributes != nil {
		return true
	}

	return false
}

// SetQueryAttributes gets a reference to the given []WorkflowSearchResponseEntry and assigns it to the QueryAttributes field.
func (o *WorkflowSearchResponse) SetQueryAttributes(v []WorkflowSearchResponseEntry) {
	o.QueryAttributes = v
}

func (o WorkflowSearchResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.QueryAttributes != nil {
		toSerialize["queryAttributes"] = o.QueryAttributes
	}
	return json.Marshal(toSerialize)
}

type NullableWorkflowSearchResponse struct {
	value *WorkflowSearchResponse
	isSet bool
}

func (v NullableWorkflowSearchResponse) Get() *WorkflowSearchResponse {
	return v.value
}

func (v *NullableWorkflowSearchResponse) Set(val *WorkflowSearchResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowSearchResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowSearchResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowSearchResponse(val *WorkflowSearchResponse) *NullableWorkflowSearchResponse {
	return &NullableWorkflowSearchResponse{value: val, isSet: true}
}

func (v NullableWorkflowSearchResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowSearchResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


