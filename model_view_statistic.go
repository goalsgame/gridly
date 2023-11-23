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

// ViewStatistic struct for ViewStatistic
type ViewStatistic struct {
	RecordCount *int64 `json:"recordCount,omitempty"`
	Translations *map[string]ColumnStatistic `json:"translations,omitempty"`
}

// NewViewStatistic instantiates a new ViewStatistic object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewViewStatistic() *ViewStatistic {
	this := ViewStatistic{}
	return &this
}

// NewViewStatisticWithDefaults instantiates a new ViewStatistic object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewViewStatisticWithDefaults() *ViewStatistic {
	this := ViewStatistic{}
	return &this
}

// GetRecordCount returns the RecordCount field value if set, zero value otherwise.
func (o *ViewStatistic) GetRecordCount() int64 {
	if o == nil || o.RecordCount == nil {
		var ret int64
		return ret
	}
	return *o.RecordCount
}

// GetRecordCountOk returns a tuple with the RecordCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewStatistic) GetRecordCountOk() (*int64, bool) {
	if o == nil || o.RecordCount == nil {
		return nil, false
	}
	return o.RecordCount, true
}

// HasRecordCount returns a boolean if a field has been set.
func (o *ViewStatistic) HasRecordCount() bool {
	if o != nil && o.RecordCount != nil {
		return true
	}

	return false
}

// SetRecordCount gets a reference to the given int64 and assigns it to the RecordCount field.
func (o *ViewStatistic) SetRecordCount(v int64) {
	o.RecordCount = &v
}

// GetTranslations returns the Translations field value if set, zero value otherwise.
func (o *ViewStatistic) GetTranslations() map[string]ColumnStatistic {
	if o == nil || o.Translations == nil {
		var ret map[string]ColumnStatistic
		return ret
	}
	return *o.Translations
}

// GetTranslationsOk returns a tuple with the Translations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewStatistic) GetTranslationsOk() (*map[string]ColumnStatistic, bool) {
	if o == nil || o.Translations == nil {
		return nil, false
	}
	return o.Translations, true
}

// HasTranslations returns a boolean if a field has been set.
func (o *ViewStatistic) HasTranslations() bool {
	if o != nil && o.Translations != nil {
		return true
	}

	return false
}

// SetTranslations gets a reference to the given map[string]ColumnStatistic and assigns it to the Translations field.
func (o *ViewStatistic) SetTranslations(v map[string]ColumnStatistic) {
	o.Translations = &v
}

func (o ViewStatistic) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.RecordCount != nil {
		toSerialize["recordCount"] = o.RecordCount
	}
	if o.Translations != nil {
		toSerialize["translations"] = o.Translations
	}
	return json.Marshal(toSerialize)
}

type NullableViewStatistic struct {
	value *ViewStatistic
	isSet bool
}

func (v NullableViewStatistic) Get() *ViewStatistic {
	return v.value
}

func (v *NullableViewStatistic) Set(val *ViewStatistic) {
	v.value = val
	v.isSet = true
}

func (v NullableViewStatistic) IsSet() bool {
	return v.isSet
}

func (v *NullableViewStatistic) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableViewStatistic(val *ViewStatistic) *NullableViewStatistic {
	return &NullableViewStatistic{value: val, isSet: true}
}

func (v NullableViewStatistic) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableViewStatistic) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


