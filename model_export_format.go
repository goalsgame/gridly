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

// ExportFormat the model 'ExportFormat'
type ExportFormat string

// List of ExportFormat
const (
	TMX ExportFormat = "tmx"
	CSV ExportFormat = "csv"
	XLSX ExportFormat = "xlsx"
)

// All allowed values of ExportFormat enum
var AllowedExportFormatEnumValues = []ExportFormat{
	"tmx",
	"csv",
	"xlsx",
}

func (v *ExportFormat) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ExportFormat(value)
	for _, existing := range AllowedExportFormatEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ExportFormat", value)
}

// NewExportFormatFromValue returns a pointer to a valid ExportFormat
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewExportFormatFromValue(v string) (*ExportFormat, error) {
	ev := ExportFormat(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ExportFormat: valid values are %v", v, AllowedExportFormatEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ExportFormat) IsValid() bool {
	for _, existing := range AllowedExportFormatEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ExportFormat value
func (v ExportFormat) Ptr() *ExportFormat {
	return &v
}

type NullableExportFormat struct {
	value *ExportFormat
	isSet bool
}

func (v NullableExportFormat) Get() *ExportFormat {
	return v.value
}

func (v *NullableExportFormat) Set(val *ExportFormat) {
	v.value = val
	v.isSet = true
}

func (v NullableExportFormat) IsSet() bool {
	return v.isSet
}

func (v *NullableExportFormat) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExportFormat(val *ExportFormat) *NullableExportFormat {
	return &NullableExportFormat{value: val, isSet: true}
}

func (v NullableExportFormat) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExportFormat) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

