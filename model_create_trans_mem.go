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

// CreateTransMem struct for CreateTransMem
type CreateTransMem struct {
	Name string `json:"name"`
	Description *string `json:"description,omitempty"`
	ProjectIds []int64 `json:"projectIds,omitempty"`
	FuzzyMatch *bool `json:"fuzzyMatch,omitempty"`
	IsDisabled *bool `json:"isDisabled,omitempty"`
	IsPausedConsuming *bool `json:"isPausedConsuming,omitempty"`
	PopulateTranslationStatus *TranslationStatus `json:"populateTranslationStatus,omitempty"`
}

// NewCreateTransMem instantiates a new CreateTransMem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateTransMem(name string) *CreateTransMem {
	this := CreateTransMem{}
	this.Name = name
	return &this
}

// NewCreateTransMemWithDefaults instantiates a new CreateTransMem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateTransMemWithDefaults() *CreateTransMem {
	this := CreateTransMem{}
	return &this
}

// GetName returns the Name field value
func (o *CreateTransMem) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateTransMem) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateTransMem) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreateTransMem) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTransMem) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateTransMem) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreateTransMem) SetDescription(v string) {
	o.Description = &v
}

// GetProjectIds returns the ProjectIds field value if set, zero value otherwise.
func (o *CreateTransMem) GetProjectIds() []int64 {
	if o == nil || o.ProjectIds == nil {
		var ret []int64
		return ret
	}
	return o.ProjectIds
}

// GetProjectIdsOk returns a tuple with the ProjectIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTransMem) GetProjectIdsOk() ([]int64, bool) {
	if o == nil || o.ProjectIds == nil {
		return nil, false
	}
	return o.ProjectIds, true
}

// HasProjectIds returns a boolean if a field has been set.
func (o *CreateTransMem) HasProjectIds() bool {
	if o != nil && o.ProjectIds != nil {
		return true
	}

	return false
}

// SetProjectIds gets a reference to the given []int64 and assigns it to the ProjectIds field.
func (o *CreateTransMem) SetProjectIds(v []int64) {
	o.ProjectIds = v
}

// GetFuzzyMatch returns the FuzzyMatch field value if set, zero value otherwise.
func (o *CreateTransMem) GetFuzzyMatch() bool {
	if o == nil || o.FuzzyMatch == nil {
		var ret bool
		return ret
	}
	return *o.FuzzyMatch
}

// GetFuzzyMatchOk returns a tuple with the FuzzyMatch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTransMem) GetFuzzyMatchOk() (*bool, bool) {
	if o == nil || o.FuzzyMatch == nil {
		return nil, false
	}
	return o.FuzzyMatch, true
}

// HasFuzzyMatch returns a boolean if a field has been set.
func (o *CreateTransMem) HasFuzzyMatch() bool {
	if o != nil && o.FuzzyMatch != nil {
		return true
	}

	return false
}

// SetFuzzyMatch gets a reference to the given bool and assigns it to the FuzzyMatch field.
func (o *CreateTransMem) SetFuzzyMatch(v bool) {
	o.FuzzyMatch = &v
}

// GetIsDisabled returns the IsDisabled field value if set, zero value otherwise.
func (o *CreateTransMem) GetIsDisabled() bool {
	if o == nil || o.IsDisabled == nil {
		var ret bool
		return ret
	}
	return *o.IsDisabled
}

// GetIsDisabledOk returns a tuple with the IsDisabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTransMem) GetIsDisabledOk() (*bool, bool) {
	if o == nil || o.IsDisabled == nil {
		return nil, false
	}
	return o.IsDisabled, true
}

// HasIsDisabled returns a boolean if a field has been set.
func (o *CreateTransMem) HasIsDisabled() bool {
	if o != nil && o.IsDisabled != nil {
		return true
	}

	return false
}

// SetIsDisabled gets a reference to the given bool and assigns it to the IsDisabled field.
func (o *CreateTransMem) SetIsDisabled(v bool) {
	o.IsDisabled = &v
}

// GetIsPausedConsuming returns the IsPausedConsuming field value if set, zero value otherwise.
func (o *CreateTransMem) GetIsPausedConsuming() bool {
	if o == nil || o.IsPausedConsuming == nil {
		var ret bool
		return ret
	}
	return *o.IsPausedConsuming
}

// GetIsPausedConsumingOk returns a tuple with the IsPausedConsuming field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTransMem) GetIsPausedConsumingOk() (*bool, bool) {
	if o == nil || o.IsPausedConsuming == nil {
		return nil, false
	}
	return o.IsPausedConsuming, true
}

// HasIsPausedConsuming returns a boolean if a field has been set.
func (o *CreateTransMem) HasIsPausedConsuming() bool {
	if o != nil && o.IsPausedConsuming != nil {
		return true
	}

	return false
}

// SetIsPausedConsuming gets a reference to the given bool and assigns it to the IsPausedConsuming field.
func (o *CreateTransMem) SetIsPausedConsuming(v bool) {
	o.IsPausedConsuming = &v
}

// GetPopulateTranslationStatus returns the PopulateTranslationStatus field value if set, zero value otherwise.
func (o *CreateTransMem) GetPopulateTranslationStatus() TranslationStatus {
	if o == nil || o.PopulateTranslationStatus == nil {
		var ret TranslationStatus
		return ret
	}
	return *o.PopulateTranslationStatus
}

// GetPopulateTranslationStatusOk returns a tuple with the PopulateTranslationStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTransMem) GetPopulateTranslationStatusOk() (*TranslationStatus, bool) {
	if o == nil || o.PopulateTranslationStatus == nil {
		return nil, false
	}
	return o.PopulateTranslationStatus, true
}

// HasPopulateTranslationStatus returns a boolean if a field has been set.
func (o *CreateTransMem) HasPopulateTranslationStatus() bool {
	if o != nil && o.PopulateTranslationStatus != nil {
		return true
	}

	return false
}

// SetPopulateTranslationStatus gets a reference to the given TranslationStatus and assigns it to the PopulateTranslationStatus field.
func (o *CreateTransMem) SetPopulateTranslationStatus(v TranslationStatus) {
	o.PopulateTranslationStatus = &v
}

func (o CreateTransMem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.ProjectIds != nil {
		toSerialize["projectIds"] = o.ProjectIds
	}
	if o.FuzzyMatch != nil {
		toSerialize["fuzzyMatch"] = o.FuzzyMatch
	}
	if o.IsDisabled != nil {
		toSerialize["isDisabled"] = o.IsDisabled
	}
	if o.IsPausedConsuming != nil {
		toSerialize["isPausedConsuming"] = o.IsPausedConsuming
	}
	if o.PopulateTranslationStatus != nil {
		toSerialize["populateTranslationStatus"] = o.PopulateTranslationStatus
	}
	return json.Marshal(toSerialize)
}

type NullableCreateTransMem struct {
	value *CreateTransMem
	isSet bool
}

func (v NullableCreateTransMem) Get() *CreateTransMem {
	return v.value
}

func (v *NullableCreateTransMem) Set(val *CreateTransMem) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateTransMem) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateTransMem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateTransMem(val *CreateTransMem) *NullableCreateTransMem {
	return &NullableCreateTransMem{value: val, isSet: true}
}

func (v NullableCreateTransMem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateTransMem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


