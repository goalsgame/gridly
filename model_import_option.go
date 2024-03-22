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
	"fmt"
)

// ImportOption the model 'ImportOption'
type ImportOption string

// List of ImportOption
const (
	ADD ImportOption = "ADD"
	UPDATE ImportOption = "UPDATE"
	UPDATE_ONLY ImportOption = "UPDATE_ONLY"
)

// All allowed values of ImportOption enum
var AllowedImportOptionEnumValues = []ImportOption{
	"ADD",
	"UPDATE",
	"UPDATE_ONLY",
}

func (v *ImportOption) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ImportOption(value)
	for _, existing := range AllowedImportOptionEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ImportOption", value)
}

// NewImportOptionFromValue returns a pointer to a valid ImportOption
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewImportOptionFromValue(v string) (*ImportOption, error) {
	ev := ImportOption(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ImportOption: valid values are %v", v, AllowedImportOptionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ImportOption) IsValid() bool {
	for _, existing := range AllowedImportOptionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ImportOption value
func (v ImportOption) Ptr() *ImportOption {
	return &v
}

type NullableImportOption struct {
	value *ImportOption
	isSet bool
}

func (v NullableImportOption) Get() *ImportOption {
	return v.value
}

func (v *NullableImportOption) Set(val *ImportOption) {
	v.value = val
	v.isSet = true
}

func (v NullableImportOption) IsSet() bool {
	return v.isSet
}

func (v *NullableImportOption) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImportOption(val *ImportOption) *NullableImportOption {
	return &NullableImportOption{value: val, isSet: true}
}

func (v NullableImportOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImportOption) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

