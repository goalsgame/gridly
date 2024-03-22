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

// CreateShareView struct for CreateShareView
type CreateShareView struct {
	IncludeGridHistory *bool `json:"includeGridHistory,omitempty"`
}

// NewCreateShareView instantiates a new CreateShareView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateShareView() *CreateShareView {
	this := CreateShareView{}
	return &this
}

// NewCreateShareViewWithDefaults instantiates a new CreateShareView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateShareViewWithDefaults() *CreateShareView {
	this := CreateShareView{}
	return &this
}

// GetIncludeGridHistory returns the IncludeGridHistory field value if set, zero value otherwise.
func (o *CreateShareView) GetIncludeGridHistory() bool {
	if o == nil || o.IncludeGridHistory == nil {
		var ret bool
		return ret
	}
	return *o.IncludeGridHistory
}

// GetIncludeGridHistoryOk returns a tuple with the IncludeGridHistory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateShareView) GetIncludeGridHistoryOk() (*bool, bool) {
	if o == nil || o.IncludeGridHistory == nil {
		return nil, false
	}
	return o.IncludeGridHistory, true
}

// HasIncludeGridHistory returns a boolean if a field has been set.
func (o *CreateShareView) HasIncludeGridHistory() bool {
	if o != nil && o.IncludeGridHistory != nil {
		return true
	}

	return false
}

// SetIncludeGridHistory gets a reference to the given bool and assigns it to the IncludeGridHistory field.
func (o *CreateShareView) SetIncludeGridHistory(v bool) {
	o.IncludeGridHistory = &v
}

func (o CreateShareView) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.IncludeGridHistory != nil {
		toSerialize["includeGridHistory"] = o.IncludeGridHistory
	}
	return json.Marshal(toSerialize)
}

type NullableCreateShareView struct {
	value *CreateShareView
	isSet bool
}

func (v NullableCreateShareView) Get() *CreateShareView {
	return v.value
}

func (v *NullableCreateShareView) Set(val *CreateShareView) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateShareView) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateShareView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateShareView(val *CreateShareView) *NullableCreateShareView {
	return &NullableCreateShareView{value: val, isSet: true}
}

func (v NullableCreateShareView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateShareView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


