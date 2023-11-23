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

// SetCell struct for SetCell
type SetCell struct {
	ColumnId *string `json:"columnId,omitempty"`
	DependencyStatus *string `json:"dependencyStatus,omitempty"`
	LengthLimit *int32 `json:"lengthLimit,omitempty"`
	ReferencedIds []string `json:"referencedIds,omitempty"`
	SourceStatus *string `json:"sourceStatus,omitempty"`
	Value *interface{} `json:"value,omitempty"`
}

// NewSetCell instantiates a new SetCell object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSetCell() *SetCell {
	this := SetCell{}
	return &this
}

// NewSetCellWithDefaults instantiates a new SetCell object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSetCellWithDefaults() *SetCell {
	this := SetCell{}
	return &this
}

// GetColumnId returns the ColumnId field value if set, zero value otherwise.
func (o *SetCell) GetColumnId() string {
	if o == nil || o.ColumnId == nil {
		var ret string
		return ret
	}
	return *o.ColumnId
}

// GetColumnIdOk returns a tuple with the ColumnId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetCell) GetColumnIdOk() (*string, bool) {
	if o == nil || o.ColumnId == nil {
		return nil, false
	}
	return o.ColumnId, true
}

// HasColumnId returns a boolean if a field has been set.
func (o *SetCell) HasColumnId() bool {
	if o != nil && o.ColumnId != nil {
		return true
	}

	return false
}

// SetColumnId gets a reference to the given string and assigns it to the ColumnId field.
func (o *SetCell) SetColumnId(v string) {
	o.ColumnId = &v
}

// GetDependencyStatus returns the DependencyStatus field value if set, zero value otherwise.
func (o *SetCell) GetDependencyStatus() string {
	if o == nil || o.DependencyStatus == nil {
		var ret string
		return ret
	}
	return *o.DependencyStatus
}

// GetDependencyStatusOk returns a tuple with the DependencyStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetCell) GetDependencyStatusOk() (*string, bool) {
	if o == nil || o.DependencyStatus == nil {
		return nil, false
	}
	return o.DependencyStatus, true
}

// HasDependencyStatus returns a boolean if a field has been set.
func (o *SetCell) HasDependencyStatus() bool {
	if o != nil && o.DependencyStatus != nil {
		return true
	}

	return false
}

// SetDependencyStatus gets a reference to the given string and assigns it to the DependencyStatus field.
func (o *SetCell) SetDependencyStatus(v string) {
	o.DependencyStatus = &v
}

// GetLengthLimit returns the LengthLimit field value if set, zero value otherwise.
func (o *SetCell) GetLengthLimit() int32 {
	if o == nil || o.LengthLimit == nil {
		var ret int32
		return ret
	}
	return *o.LengthLimit
}

// GetLengthLimitOk returns a tuple with the LengthLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetCell) GetLengthLimitOk() (*int32, bool) {
	if o == nil || o.LengthLimit == nil {
		return nil, false
	}
	return o.LengthLimit, true
}

// HasLengthLimit returns a boolean if a field has been set.
func (o *SetCell) HasLengthLimit() bool {
	if o != nil && o.LengthLimit != nil {
		return true
	}

	return false
}

// SetLengthLimit gets a reference to the given int32 and assigns it to the LengthLimit field.
func (o *SetCell) SetLengthLimit(v int32) {
	o.LengthLimit = &v
}

// GetReferencedIds returns the ReferencedIds field value if set, zero value otherwise.
func (o *SetCell) GetReferencedIds() []string {
	if o == nil || o.ReferencedIds == nil {
		var ret []string
		return ret
	}
	return o.ReferencedIds
}

// GetReferencedIdsOk returns a tuple with the ReferencedIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetCell) GetReferencedIdsOk() ([]string, bool) {
	if o == nil || o.ReferencedIds == nil {
		return nil, false
	}
	return o.ReferencedIds, true
}

// HasReferencedIds returns a boolean if a field has been set.
func (o *SetCell) HasReferencedIds() bool {
	if o != nil && o.ReferencedIds != nil {
		return true
	}

	return false
}

// SetReferencedIds gets a reference to the given []string and assigns it to the ReferencedIds field.
func (o *SetCell) SetReferencedIds(v []string) {
	o.ReferencedIds = v
}

// GetSourceStatus returns the SourceStatus field value if set, zero value otherwise.
func (o *SetCell) GetSourceStatus() string {
	if o == nil || o.SourceStatus == nil {
		var ret string
		return ret
	}
	return *o.SourceStatus
}

// GetSourceStatusOk returns a tuple with the SourceStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetCell) GetSourceStatusOk() (*string, bool) {
	if o == nil || o.SourceStatus == nil {
		return nil, false
	}
	return o.SourceStatus, true
}

// HasSourceStatus returns a boolean if a field has been set.
func (o *SetCell) HasSourceStatus() bool {
	if o != nil && o.SourceStatus != nil {
		return true
	}

	return false
}

// SetSourceStatus gets a reference to the given string and assigns it to the SourceStatus field.
func (o *SetCell) SetSourceStatus(v string) {
	o.SourceStatus = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *SetCell) GetValue() interface{} {
	if o == nil || o.Value == nil {
		var ret interface{}
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetCell) GetValueOk() (*interface{}, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *SetCell) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given interface{} and assigns it to the Value field.
func (o *SetCell) SetValue(v interface{}) {
	o.Value = &v
}

func (o SetCell) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ColumnId != nil {
		toSerialize["columnId"] = o.ColumnId
	}
	if o.DependencyStatus != nil {
		toSerialize["dependencyStatus"] = o.DependencyStatus
	}
	if o.LengthLimit != nil {
		toSerialize["lengthLimit"] = o.LengthLimit
	}
	if o.ReferencedIds != nil {
		toSerialize["referencedIds"] = o.ReferencedIds
	}
	if o.SourceStatus != nil {
		toSerialize["sourceStatus"] = o.SourceStatus
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableSetCell struct {
	value *SetCell
	isSet bool
}

func (v NullableSetCell) Get() *SetCell {
	return v.value
}

func (v *NullableSetCell) Set(val *SetCell) {
	v.value = val
	v.isSet = true
}

func (v NullableSetCell) IsSet() bool {
	return v.isSet
}

func (v *NullableSetCell) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSetCell(val *SetCell) *NullableSetCell {
	return &NullableSetCell{value: val, isSet: true}
}

func (v NullableSetCell) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSetCell) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


