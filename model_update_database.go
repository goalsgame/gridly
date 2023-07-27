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

// UpdateDatabase struct for UpdateDatabase
type UpdateDatabase struct {
	Name string `json:"name"`
	Description *string `json:"description,omitempty"`
}

// NewUpdateDatabase instantiates a new UpdateDatabase object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateDatabase(name string) *UpdateDatabase {
	this := UpdateDatabase{}
	this.Name = name
	return &this
}

// NewUpdateDatabaseWithDefaults instantiates a new UpdateDatabase object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateDatabaseWithDefaults() *UpdateDatabase {
	this := UpdateDatabase{}
	return &this
}

// GetName returns the Name field value
func (o *UpdateDatabase) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UpdateDatabase) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *UpdateDatabase) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UpdateDatabase) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDatabase) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UpdateDatabase) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UpdateDatabase) SetDescription(v string) {
	o.Description = &v
}

func (o UpdateDatabase) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateDatabase struct {
	value *UpdateDatabase
	isSet bool
}

func (v NullableUpdateDatabase) Get() *UpdateDatabase {
	return v.value
}

func (v *NullableUpdateDatabase) Set(val *UpdateDatabase) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateDatabase) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateDatabase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateDatabase(val *UpdateDatabase) *NullableUpdateDatabase {
	return &NullableUpdateDatabase{value: val, isSet: true}
}

func (v NullableUpdateDatabase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateDatabase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


