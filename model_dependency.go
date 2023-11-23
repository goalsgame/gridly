/*
Gridly API

Gridly API documentation

API version: 4.21.5
Contact: support@gridly.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gridly

import (
	"encoding/json"
)

// Dependency struct for Dependency
type Dependency struct {
	Id *string `json:"id,omitempty"`
	SourceColumnId string `json:"sourceColumnId"`
	TargetColumnId string `json:"targetColumnId"`
}

// NewDependency instantiates a new Dependency object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDependency(sourceColumnId string, targetColumnId string) *Dependency {
	this := Dependency{}
	this.SourceColumnId = sourceColumnId
	this.TargetColumnId = targetColumnId
	return &this
}

// NewDependencyWithDefaults instantiates a new Dependency object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDependencyWithDefaults() *Dependency {
	this := Dependency{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Dependency) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dependency) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Dependency) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Dependency) SetId(v string) {
	o.Id = &v
}

// GetSourceColumnId returns the SourceColumnId field value
func (o *Dependency) GetSourceColumnId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SourceColumnId
}

// GetSourceColumnIdOk returns a tuple with the SourceColumnId field value
// and a boolean to check if the value has been set.
func (o *Dependency) GetSourceColumnIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceColumnId, true
}

// SetSourceColumnId sets field value
func (o *Dependency) SetSourceColumnId(v string) {
	o.SourceColumnId = v
}

// GetTargetColumnId returns the TargetColumnId field value
func (o *Dependency) GetTargetColumnId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TargetColumnId
}

// GetTargetColumnIdOk returns a tuple with the TargetColumnId field value
// and a boolean to check if the value has been set.
func (o *Dependency) GetTargetColumnIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetColumnId, true
}

// SetTargetColumnId sets field value
func (o *Dependency) SetTargetColumnId(v string) {
	o.TargetColumnId = v
}

func (o Dependency) MarshalJSON() ([]byte, error) {
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

type NullableDependency struct {
	value *Dependency
	isSet bool
}

func (v NullableDependency) Get() *Dependency {
	return v.value
}

func (v *NullableDependency) Set(val *Dependency) {
	v.value = val
	v.isSet = true
}

func (v NullableDependency) IsSet() bool {
	return v.isSet
}

func (v *NullableDependency) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDependency(val *Dependency) *NullableDependency {
	return &NullableDependency{value: val, isSet: true}
}

func (v NullableDependency) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDependency) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


