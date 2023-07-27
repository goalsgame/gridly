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

// Formula struct for Formula
type Formula struct {
	FormulaText string `json:"formulaText"`
	AlwaysFormatResultValueAsList *bool `json:"alwaysFormatResultValueAsList,omitempty"`
	DetectResultValueType *string `json:"detectResultValueType,omitempty"`
}

// NewFormula instantiates a new Formula object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFormula(formulaText string) *Formula {
	this := Formula{}
	this.FormulaText = formulaText
	return &this
}

// NewFormulaWithDefaults instantiates a new Formula object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFormulaWithDefaults() *Formula {
	this := Formula{}
	return &this
}

// GetFormulaText returns the FormulaText field value
func (o *Formula) GetFormulaText() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FormulaText
}

// GetFormulaTextOk returns a tuple with the FormulaText field value
// and a boolean to check if the value has been set.
func (o *Formula) GetFormulaTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FormulaText, true
}

// SetFormulaText sets field value
func (o *Formula) SetFormulaText(v string) {
	o.FormulaText = v
}

// GetAlwaysFormatResultValueAsList returns the AlwaysFormatResultValueAsList field value if set, zero value otherwise.
func (o *Formula) GetAlwaysFormatResultValueAsList() bool {
	if o == nil || o.AlwaysFormatResultValueAsList == nil {
		var ret bool
		return ret
	}
	return *o.AlwaysFormatResultValueAsList
}

// GetAlwaysFormatResultValueAsListOk returns a tuple with the AlwaysFormatResultValueAsList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Formula) GetAlwaysFormatResultValueAsListOk() (*bool, bool) {
	if o == nil || o.AlwaysFormatResultValueAsList == nil {
		return nil, false
	}
	return o.AlwaysFormatResultValueAsList, true
}

// HasAlwaysFormatResultValueAsList returns a boolean if a field has been set.
func (o *Formula) HasAlwaysFormatResultValueAsList() bool {
	if o != nil && o.AlwaysFormatResultValueAsList != nil {
		return true
	}

	return false
}

// SetAlwaysFormatResultValueAsList gets a reference to the given bool and assigns it to the AlwaysFormatResultValueAsList field.
func (o *Formula) SetAlwaysFormatResultValueAsList(v bool) {
	o.AlwaysFormatResultValueAsList = &v
}

// GetDetectResultValueType returns the DetectResultValueType field value if set, zero value otherwise.
func (o *Formula) GetDetectResultValueType() string {
	if o == nil || o.DetectResultValueType == nil {
		var ret string
		return ret
	}
	return *o.DetectResultValueType
}

// GetDetectResultValueTypeOk returns a tuple with the DetectResultValueType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Formula) GetDetectResultValueTypeOk() (*string, bool) {
	if o == nil || o.DetectResultValueType == nil {
		return nil, false
	}
	return o.DetectResultValueType, true
}

// HasDetectResultValueType returns a boolean if a field has been set.
func (o *Formula) HasDetectResultValueType() bool {
	if o != nil && o.DetectResultValueType != nil {
		return true
	}

	return false
}

// SetDetectResultValueType gets a reference to the given string and assigns it to the DetectResultValueType field.
func (o *Formula) SetDetectResultValueType(v string) {
	o.DetectResultValueType = &v
}

func (o Formula) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["formulaText"] = o.FormulaText
	}
	if o.AlwaysFormatResultValueAsList != nil {
		toSerialize["alwaysFormatResultValueAsList"] = o.AlwaysFormatResultValueAsList
	}
	if o.DetectResultValueType != nil {
		toSerialize["detectResultValueType"] = o.DetectResultValueType
	}
	return json.Marshal(toSerialize)
}

type NullableFormula struct {
	value *Formula
	isSet bool
}

func (v NullableFormula) Get() *Formula {
	return v.value
}

func (v *NullableFormula) Set(val *Formula) {
	v.value = val
	v.isSet = true
}

func (v NullableFormula) IsSet() bool {
	return v.isSet
}

func (v *NullableFormula) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFormula(val *Formula) *NullableFormula {
	return &NullableFormula{value: val, isSet: true}
}

func (v NullableFormula) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFormula) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


