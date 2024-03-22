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

// ReferencedColumn struct for ReferencedColumn
type ReferencedColumn struct {
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	SelectionOptions []string `json:"selectionOptions,omitempty"`
}

// NewReferencedColumn instantiates a new ReferencedColumn object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReferencedColumn() *ReferencedColumn {
	this := ReferencedColumn{}
	return &this
}

// NewReferencedColumnWithDefaults instantiates a new ReferencedColumn object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReferencedColumnWithDefaults() *ReferencedColumn {
	this := ReferencedColumn{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ReferencedColumn) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReferencedColumn) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ReferencedColumn) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ReferencedColumn) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ReferencedColumn) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReferencedColumn) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ReferencedColumn) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ReferencedColumn) SetName(v string) {
	o.Name = &v
}

// GetSelectionOptions returns the SelectionOptions field value if set, zero value otherwise.
func (o *ReferencedColumn) GetSelectionOptions() []string {
	if o == nil || o.SelectionOptions == nil {
		var ret []string
		return ret
	}
	return o.SelectionOptions
}

// GetSelectionOptionsOk returns a tuple with the SelectionOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReferencedColumn) GetSelectionOptionsOk() ([]string, bool) {
	if o == nil || o.SelectionOptions == nil {
		return nil, false
	}
	return o.SelectionOptions, true
}

// HasSelectionOptions returns a boolean if a field has been set.
func (o *ReferencedColumn) HasSelectionOptions() bool {
	if o != nil && o.SelectionOptions != nil {
		return true
	}

	return false
}

// SetSelectionOptions gets a reference to the given []string and assigns it to the SelectionOptions field.
func (o *ReferencedColumn) SetSelectionOptions(v []string) {
	o.SelectionOptions = v
}

func (o ReferencedColumn) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.SelectionOptions != nil {
		toSerialize["selectionOptions"] = o.SelectionOptions
	}
	return json.Marshal(toSerialize)
}

type NullableReferencedColumn struct {
	value *ReferencedColumn
	isSet bool
}

func (v NullableReferencedColumn) Get() *ReferencedColumn {
	return v.value
}

func (v *NullableReferencedColumn) Set(val *ReferencedColumn) {
	v.value = val
	v.isSet = true
}

func (v NullableReferencedColumn) IsSet() bool {
	return v.isSet
}

func (v *NullableReferencedColumn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReferencedColumn(val *ReferencedColumn) *NullableReferencedColumn {
	return &NullableReferencedColumn{value: val, isSet: true}
}

func (v NullableReferencedColumn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReferencedColumn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


