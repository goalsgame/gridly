/*
Gridly API

Gridly API documentation

API version: 4.29.1
Contact: support@gridly.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gridly

import (
	"encoding/json"
)

// PathSingle struct for PathSingle
type PathSingle struct {
	Path *string `json:"path,omitempty"`
}

// NewPathSingle instantiates a new PathSingle object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPathSingle() *PathSingle {
	this := PathSingle{}
	return &this
}

// NewPathSingleWithDefaults instantiates a new PathSingle object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPathSingleWithDefaults() *PathSingle {
	this := PathSingle{}
	return &this
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *PathSingle) GetPath() string {
	if o == nil || o.Path == nil {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PathSingle) GetPathOk() (*string, bool) {
	if o == nil || o.Path == nil {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *PathSingle) HasPath() bool {
	if o != nil && o.Path != nil {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *PathSingle) SetPath(v string) {
	o.Path = &v
}

func (o PathSingle) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Path != nil {
		toSerialize["path"] = o.Path
	}
	return json.Marshal(toSerialize)
}

type NullablePathSingle struct {
	value *PathSingle
	isSet bool
}

func (v NullablePathSingle) Get() *PathSingle {
	return v.value
}

func (v *NullablePathSingle) Set(val *PathSingle) {
	v.value = val
	v.isSet = true
}

func (v NullablePathSingle) IsSet() bool {
	return v.isSet
}

func (v *NullablePathSingle) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePathSingle(val *PathSingle) *NullablePathSingle {
	return &NullablePathSingle{value: val, isSet: true}
}

func (v NullablePathSingle) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePathSingle) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

