/*
Gridly API

Gridly API documentation

API version: 4.15.1
Contact: support@gridly.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gridly

import (
	"encoding/json"
)

// CreateDependency struct for CreateDependency
type CreateDependency struct {
	Id *string `json:"id,omitempty"`
	SourceColumnId string `json:"sourceColumnId"`
	TargetColumnId string `json:"targetColumnId"`
}

// NewCreateDependency instantiates a new CreateDependency object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateDependency(sourceColumnId string, targetColumnId string) *CreateDependency {
	this := CreateDependency{}
	this.SourceColumnId = sourceColumnId
	this.TargetColumnId = targetColumnId
	return &this
}

// NewCreateDependencyWithDefaults instantiates a new CreateDependency object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateDependencyWithDefaults() *CreateDependency {
	this := CreateDependency{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CreateDependency) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDependency) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CreateDependency) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CreateDependency) SetId(v string) {
	o.Id = &v
}

// GetSourceColumnId returns the SourceColumnId field value
func (o *CreateDependency) GetSourceColumnId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SourceColumnId
}

// GetSourceColumnIdOk returns a tuple with the SourceColumnId field value
// and a boolean to check if the value has been set.
func (o *CreateDependency) GetSourceColumnIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceColumnId, true
}

// SetSourceColumnId sets field value
func (o *CreateDependency) SetSourceColumnId(v string) {
	o.SourceColumnId = v
}

// GetTargetColumnId returns the TargetColumnId field value
func (o *CreateDependency) GetTargetColumnId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TargetColumnId
}

// GetTargetColumnIdOk returns a tuple with the TargetColumnId field value
// and a boolean to check if the value has been set.
func (o *CreateDependency) GetTargetColumnIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetColumnId, true
}

// SetTargetColumnId sets field value
func (o *CreateDependency) SetTargetColumnId(v string) {
	o.TargetColumnId = v
}

func (o CreateDependency) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["sourceColumnId"] = o.SourceColumnId
	}
	if true {
		toSerialize["targetColumnId"] = o.TargetColumnId
	}
	return json.Marshal(toSerialize)
}

type NullableCreateDependency struct {
	value *CreateDependency
	isSet bool
}

func (v NullableCreateDependency) Get() *CreateDependency {
	return v.value
}

func (v *NullableCreateDependency) Set(val *CreateDependency) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateDependency) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateDependency) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateDependency(val *CreateDependency) *NullableCreateDependency {
	return &NullableCreateDependency{value: val, isSet: true}
}

func (v NullableCreateDependency) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateDependency) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


