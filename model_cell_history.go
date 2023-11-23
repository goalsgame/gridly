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

// CellHistory struct for CellHistory
type CellHistory struct {
	DependencyStatus *string `json:"dependencyStatus,omitempty"`
	SourceStatus *string `json:"sourceStatus,omitempty"`
	ColumnId *string `json:"columnId,omitempty"`
	Value map[string]interface{} `json:"value,omitempty"`
}

// NewCellHistory instantiates a new CellHistory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCellHistory() *CellHistory {
	this := CellHistory{}
	return &this
}

// NewCellHistoryWithDefaults instantiates a new CellHistory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCellHistoryWithDefaults() *CellHistory {
	this := CellHistory{}
	return &this
}

// GetDependencyStatus returns the DependencyStatus field value if set, zero value otherwise.
func (o *CellHistory) GetDependencyStatus() string {
	if o == nil || o.DependencyStatus == nil {
		var ret string
		return ret
	}
	return *o.DependencyStatus
}

// GetDependencyStatusOk returns a tuple with the DependencyStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CellHistory) GetDependencyStatusOk() (*string, bool) {
	if o == nil || o.DependencyStatus == nil {
		return nil, false
	}
	return o.DependencyStatus, true
}

// HasDependencyStatus returns a boolean if a field has been set.
func (o *CellHistory) HasDependencyStatus() bool {
	if o != nil && o.DependencyStatus != nil {
		return true
	}

	return false
}

// SetDependencyStatus gets a reference to the given string and assigns it to the DependencyStatus field.
func (o *CellHistory) SetDependencyStatus(v string) {
	o.DependencyStatus = &v
}

// GetSourceStatus returns the SourceStatus field value if set, zero value otherwise.
func (o *CellHistory) GetSourceStatus() string {
	if o == nil || o.SourceStatus == nil {
		var ret string
		return ret
	}
	return *o.SourceStatus
}

// GetSourceStatusOk returns a tuple with the SourceStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CellHistory) GetSourceStatusOk() (*string, bool) {
	if o == nil || o.SourceStatus == nil {
		return nil, false
	}
	return o.SourceStatus, true
}

// HasSourceStatus returns a boolean if a field has been set.
func (o *CellHistory) HasSourceStatus() bool {
	if o != nil && o.SourceStatus != nil {
		return true
	}

	return false
}

// SetSourceStatus gets a reference to the given string and assigns it to the SourceStatus field.
func (o *CellHistory) SetSourceStatus(v string) {
	o.SourceStatus = &v
}

// GetColumnId returns the ColumnId field value if set, zero value otherwise.
func (o *CellHistory) GetColumnId() string {
	if o == nil || o.ColumnId == nil {
		var ret string
		return ret
	}
	return *o.ColumnId
}

// GetColumnIdOk returns a tuple with the ColumnId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CellHistory) GetColumnIdOk() (*string, bool) {
	if o == nil || o.ColumnId == nil {
		return nil, false
	}
	return o.ColumnId, true
}

// HasColumnId returns a boolean if a field has been set.
func (o *CellHistory) HasColumnId() bool {
	if o != nil && o.ColumnId != nil {
		return true
	}

	return false
}

// SetColumnId gets a reference to the given string and assigns it to the ColumnId field.
func (o *CellHistory) SetColumnId(v string) {
	o.ColumnId = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CellHistory) GetValue() map[string]interface{} {
	if o == nil || o.Value == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CellHistory) GetValueOk() (map[string]interface{}, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CellHistory) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given map[string]interface{} and assigns it to the Value field.
func (o *CellHistory) SetValue(v map[string]interface{}) {
	o.Value = v
}

func (o CellHistory) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DependencyStatus != nil {
		toSerialize["dependencyStatus"] = o.DependencyStatus
	}
	if o.SourceStatus != nil {
		toSerialize["sourceStatus"] = o.SourceStatus
	}
	if o.ColumnId != nil {
		toSerialize["columnId"] = o.ColumnId
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableCellHistory struct {
	value *CellHistory
	isSet bool
}

func (v NullableCellHistory) Get() *CellHistory {
	return v.value
}

func (v *NullableCellHistory) Set(val *CellHistory) {
	v.value = val
	v.isSet = true
}

func (v NullableCellHistory) IsSet() bool {
	return v.isSet
}

func (v *NullableCellHistory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCellHistory(val *CellHistory) *NullableCellHistory {
	return &NullableCellHistory{value: val, isSet: true}
}

func (v NullableCellHistory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCellHistory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


