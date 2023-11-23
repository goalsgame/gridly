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

// DateFormat struct for DateFormat
type DateFormat struct {
	Name *string `json:"name,omitempty"`
	Format *string `json:"format,omitempty"`
}

// NewDateFormat instantiates a new DateFormat object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDateFormat() *DateFormat {
	this := DateFormat{}
	return &this
}

// NewDateFormatWithDefaults instantiates a new DateFormat object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDateFormatWithDefaults() *DateFormat {
	this := DateFormat{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DateFormat) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DateFormat) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DateFormat) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DateFormat) SetName(v string) {
	o.Name = &v
}

// GetFormat returns the Format field value if set, zero value otherwise.
func (o *DateFormat) GetFormat() string {
	if o == nil || o.Format == nil {
		var ret string
		return ret
	}
	return *o.Format
}

// GetFormatOk returns a tuple with the Format field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DateFormat) GetFormatOk() (*string, bool) {
	if o == nil || o.Format == nil {
		return nil, false
	}
	return o.Format, true
}

// HasFormat returns a boolean if a field has been set.
func (o *DateFormat) HasFormat() bool {
	if o != nil && o.Format != nil {
		return true
	}

	return false
}

// SetFormat gets a reference to the given string and assigns it to the Format field.
func (o *DateFormat) SetFormat(v string) {
	o.Format = &v
}

func (o DateFormat) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Format != nil {
		toSerialize["format"] = o.Format
	}
	return json.Marshal(toSerialize)
}

type NullableDateFormat struct {
	value *DateFormat
	isSet bool
}

func (v NullableDateFormat) Get() *DateFormat {
	return v.value
}

func (v *NullableDateFormat) Set(val *DateFormat) {
	v.value = val
	v.isSet = true
}

func (v NullableDateFormat) IsSet() bool {
	return v.isSet
}

func (v *NullableDateFormat) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDateFormat(val *DateFormat) *NullableDateFormat {
	return &NullableDateFormat{value: val, isSet: true}
}

func (v NullableDateFormat) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDateFormat) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


