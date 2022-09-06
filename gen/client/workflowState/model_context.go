/*
WorkflowState APIs

This APIs for iwf-server to invoke user workflow code defined in WorkflowState using any iwf SDKs

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// Context struct for Context
type Context struct {
	WorkflowId *string `json:"workflowId,omitempty"`
	WorkflowRunId *string `json:"workflowRunId,omitempty"`
	WorkflowStartedTimestamp *int32 `json:"workflowStartedTimestamp,omitempty"`
	StateExecutionId *string `json:"stateExecutionId,omitempty"`
}

// NewContext instantiates a new Context object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContext() *Context {
	this := Context{}
	return &this
}

// NewContextWithDefaults instantiates a new Context object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContextWithDefaults() *Context {
	this := Context{}
	return &this
}

// GetWorkflowId returns the WorkflowId field value if set, zero value otherwise.
func (o *Context) GetWorkflowId() string {
	if o == nil || o.WorkflowId == nil {
		var ret string
		return ret
	}
	return *o.WorkflowId
}

// GetWorkflowIdOk returns a tuple with the WorkflowId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Context) GetWorkflowIdOk() (*string, bool) {
	if o == nil || o.WorkflowId == nil {
		return nil, false
	}
	return o.WorkflowId, true
}

// HasWorkflowId returns a boolean if a field has been set.
func (o *Context) HasWorkflowId() bool {
	if o != nil && o.WorkflowId != nil {
		return true
	}

	return false
}

// SetWorkflowId gets a reference to the given string and assigns it to the WorkflowId field.
func (o *Context) SetWorkflowId(v string) {
	o.WorkflowId = &v
}

// GetWorkflowRunId returns the WorkflowRunId field value if set, zero value otherwise.
func (o *Context) GetWorkflowRunId() string {
	if o == nil || o.WorkflowRunId == nil {
		var ret string
		return ret
	}
	return *o.WorkflowRunId
}

// GetWorkflowRunIdOk returns a tuple with the WorkflowRunId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Context) GetWorkflowRunIdOk() (*string, bool) {
	if o == nil || o.WorkflowRunId == nil {
		return nil, false
	}
	return o.WorkflowRunId, true
}

// HasWorkflowRunId returns a boolean if a field has been set.
func (o *Context) HasWorkflowRunId() bool {
	if o != nil && o.WorkflowRunId != nil {
		return true
	}

	return false
}

// SetWorkflowRunId gets a reference to the given string and assigns it to the WorkflowRunId field.
func (o *Context) SetWorkflowRunId(v string) {
	o.WorkflowRunId = &v
}

// GetWorkflowStartedTimestamp returns the WorkflowStartedTimestamp field value if set, zero value otherwise.
func (o *Context) GetWorkflowStartedTimestamp() int32 {
	if o == nil || o.WorkflowStartedTimestamp == nil {
		var ret int32
		return ret
	}
	return *o.WorkflowStartedTimestamp
}

// GetWorkflowStartedTimestampOk returns a tuple with the WorkflowStartedTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Context) GetWorkflowStartedTimestampOk() (*int32, bool) {
	if o == nil || o.WorkflowStartedTimestamp == nil {
		return nil, false
	}
	return o.WorkflowStartedTimestamp, true
}

// HasWorkflowStartedTimestamp returns a boolean if a field has been set.
func (o *Context) HasWorkflowStartedTimestamp() bool {
	if o != nil && o.WorkflowStartedTimestamp != nil {
		return true
	}

	return false
}

// SetWorkflowStartedTimestamp gets a reference to the given int32 and assigns it to the WorkflowStartedTimestamp field.
func (o *Context) SetWorkflowStartedTimestamp(v int32) {
	o.WorkflowStartedTimestamp = &v
}

// GetStateExecutionId returns the StateExecutionId field value if set, zero value otherwise.
func (o *Context) GetStateExecutionId() string {
	if o == nil || o.StateExecutionId == nil {
		var ret string
		return ret
	}
	return *o.StateExecutionId
}

// GetStateExecutionIdOk returns a tuple with the StateExecutionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Context) GetStateExecutionIdOk() (*string, bool) {
	if o == nil || o.StateExecutionId == nil {
		return nil, false
	}
	return o.StateExecutionId, true
}

// HasStateExecutionId returns a boolean if a field has been set.
func (o *Context) HasStateExecutionId() bool {
	if o != nil && o.StateExecutionId != nil {
		return true
	}

	return false
}

// SetStateExecutionId gets a reference to the given string and assigns it to the StateExecutionId field.
func (o *Context) SetStateExecutionId(v string) {
	o.StateExecutionId = &v
}

func (o Context) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.WorkflowId != nil {
		toSerialize["workflowId"] = o.WorkflowId
	}
	if o.WorkflowRunId != nil {
		toSerialize["workflowRunId"] = o.WorkflowRunId
	}
	if o.WorkflowStartedTimestamp != nil {
		toSerialize["workflowStartedTimestamp"] = o.WorkflowStartedTimestamp
	}
	if o.StateExecutionId != nil {
		toSerialize["stateExecutionId"] = o.StateExecutionId
	}
	return json.Marshal(toSerialize)
}

type NullableContext struct {
	value *Context
	isSet bool
}

func (v NullableContext) Get() *Context {
	return v.value
}

func (v *NullableContext) Set(val *Context) {
	v.value = val
	v.isSet = true
}

func (v NullableContext) IsSet() bool {
	return v.isSet
}

func (v *NullableContext) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContext(val *Context) *NullableContext {
	return &NullableContext{value: val, isSet: true}
}

func (v NullableContext) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContext) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


