/*
Gridly API

Gridly API documentation

API version: 3.20.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gridly

import (
	"encoding/json"
)

// UpdateDependency struct for UpdateDependency
type UpdateDependency struct {
	NewId *string `json:"newId,omitempty"`
	SourceColumnId *string `json:"sourceColumnId,omitempty"`
	TargetColumnId *string `json:"targetColumnId,omitempty"`
}

// NewUpdateDependency instantiates a new UpdateDependency object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateDependency() *UpdateDependency {
	this := UpdateDependency{}
	return &this
}

// NewUpdateDependencyWithDefaults instantiates a new UpdateDependency object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateDependencyWithDefaults() *UpdateDependency {
	this := UpdateDependency{}
	return &this
}

// GetNewId returns the NewId field value if set, zero value otherwise.
func (o *UpdateDependency) GetNewId() string {
	if o == nil || o.NewId == nil {
		var ret string
		return ret
	}
	return *o.NewId
}

// GetNewIdOk returns a tuple with the NewId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDependency) GetNewIdOk() (*string, bool) {
	if o == nil || o.NewId == nil {
		return nil, false
	}
	return o.NewId, true
}

// HasNewId returns a boolean if a field has been set.
func (o *UpdateDependency) HasNewId() bool {
	if o != nil && o.NewId != nil {
		return true
	}

	return false
}

// SetNewId gets a reference to the given string and assigns it to the NewId field.
func (o *UpdateDependency) SetNewId(v string) {
	o.NewId = &v
}

// GetSourceColumnId returns the SourceColumnId field value if set, zero value otherwise.
func (o *UpdateDependency) GetSourceColumnId() string {
	if o == nil || o.SourceColumnId == nil {
		var ret string
		return ret
	}
	return *o.SourceColumnId
}

// GetSourceColumnIdOk returns a tuple with the SourceColumnId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDependency) GetSourceColumnIdOk() (*string, bool) {
	if o == nil || o.SourceColumnId == nil {
		return nil, false
	}
	return o.SourceColumnId, true
}

// HasSourceColumnId returns a boolean if a field has been set.
func (o *UpdateDependency) HasSourceColumnId() bool {
	if o != nil && o.SourceColumnId != nil {
		return true
	}

	return false
}

// SetSourceColumnId gets a reference to the given string and assigns it to the SourceColumnId field.
func (o *UpdateDependency) SetSourceColumnId(v string) {
	o.SourceColumnId = &v
}

// GetTargetColumnId returns the TargetColumnId field value if set, zero value otherwise.
func (o *UpdateDependency) GetTargetColumnId() string {
	if o == nil || o.TargetColumnId == nil {
		var ret string
		return ret
	}
	return *o.TargetColumnId
}

// GetTargetColumnIdOk returns a tuple with the TargetColumnId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDependency) GetTargetColumnIdOk() (*string, bool) {
	if o == nil || o.TargetColumnId == nil {
		return nil, false
	}
	return o.TargetColumnId, true
}

// HasTargetColumnId returns a boolean if a field has been set.
func (o *UpdateDependency) HasTargetColumnId() bool {
	if o != nil && o.TargetColumnId != nil {
		return true
	}

	return false
}

// SetTargetColumnId gets a reference to the given string and assigns it to the TargetColumnId field.
func (o *UpdateDependency) SetTargetColumnId(v string) {
	o.TargetColumnId = &v
}

func (o UpdateDependency) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.NewId != nil {
		toSerialize["newId"] = o.NewId
	}
	if o.SourceColumnId != nil {
		toSerialize["sourceColumnId"] = o.SourceColumnId
	}
	if o.TargetColumnId != nil {
		toSerialize["targetColumnId"] = o.TargetColumnId
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateDependency struct {
	value *UpdateDependency
	isSet bool
}

func (v NullableUpdateDependency) Get() *UpdateDependency {
	return v.value
}

func (v *NullableUpdateDependency) Set(val *UpdateDependency) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateDependency) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateDependency) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateDependency(val *UpdateDependency) *NullableUpdateDependency {
	return &NullableUpdateDependency{value: val, isSet: true}
}

func (v NullableUpdateDependency) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateDependency) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


