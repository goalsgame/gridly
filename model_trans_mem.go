/*
Gridly API

Gridly API documentation

API version: 4.33.0
Contact: support@gridly.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gridly

import (
	"encoding/json"
)

// TransMem struct for TransMem
type TransMem struct {
	Id *string `json:"id,omitempty"`
	ProjectIds []int64 `json:"projectIds,omitempty"`
	IsDisabled *bool `json:"isDisabled,omitempty"`
	IsPausedConsuming *bool `json:"isPausedConsuming,omitempty"`
	PopulateTranslationStatus *TranslationStatus `json:"populateTranslationStatus,omitempty"`
	ContextLookup *bool `json:"contextLookup,omitempty"`
	Name string `json:"name"`
	Description *string `json:"description,omitempty"`
	FuzzyMatch *bool `json:"fuzzyMatch,omitempty"`
	AllowAlternative *bool `json:"allowAlternative,omitempty"`
	AllowAlternativeHasSameRecordId *bool `json:"allowAlternativeHasSameRecordId,omitempty"`
}

// NewTransMem instantiates a new TransMem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransMem(name string) *TransMem {
	this := TransMem{}
	this.Name = name
	return &this
}

// NewTransMemWithDefaults instantiates a new TransMem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransMemWithDefaults() *TransMem {
	this := TransMem{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TransMem) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransMem) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TransMem) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *TransMem) SetId(v string) {
	o.Id = &v
}

// GetProjectIds returns the ProjectIds field value if set, zero value otherwise.
func (o *TransMem) GetProjectIds() []int64 {
	if o == nil || isNil(o.ProjectIds) {
		var ret []int64
		return ret
	}
	return o.ProjectIds
}

// GetProjectIdsOk returns a tuple with the ProjectIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransMem) GetProjectIdsOk() ([]int64, bool) {
	if o == nil || isNil(o.ProjectIds) {
    return nil, false
	}
	return o.ProjectIds, true
}

// HasProjectIds returns a boolean if a field has been set.
func (o *TransMem) HasProjectIds() bool {
	if o != nil && !isNil(o.ProjectIds) {
		return true
	}

	return false
}

// SetProjectIds gets a reference to the given []int64 and assigns it to the ProjectIds field.
func (o *TransMem) SetProjectIds(v []int64) {
	o.ProjectIds = v
}

// GetIsDisabled returns the IsDisabled field value if set, zero value otherwise.
func (o *TransMem) GetIsDisabled() bool {
	if o == nil || isNil(o.IsDisabled) {
		var ret bool
		return ret
	}
	return *o.IsDisabled
}

// GetIsDisabledOk returns a tuple with the IsDisabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransMem) GetIsDisabledOk() (*bool, bool) {
	if o == nil || isNil(o.IsDisabled) {
    return nil, false
	}
	return o.IsDisabled, true
}

// HasIsDisabled returns a boolean if a field has been set.
func (o *TransMem) HasIsDisabled() bool {
	if o != nil && !isNil(o.IsDisabled) {
		return true
	}

	return false
}

// SetIsDisabled gets a reference to the given bool and assigns it to the IsDisabled field.
func (o *TransMem) SetIsDisabled(v bool) {
	o.IsDisabled = &v
}

// GetIsPausedConsuming returns the IsPausedConsuming field value if set, zero value otherwise.
func (o *TransMem) GetIsPausedConsuming() bool {
	if o == nil || isNil(o.IsPausedConsuming) {
		var ret bool
		return ret
	}
	return *o.IsPausedConsuming
}

// GetIsPausedConsumingOk returns a tuple with the IsPausedConsuming field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransMem) GetIsPausedConsumingOk() (*bool, bool) {
	if o == nil || isNil(o.IsPausedConsuming) {
    return nil, false
	}
	return o.IsPausedConsuming, true
}

// HasIsPausedConsuming returns a boolean if a field has been set.
func (o *TransMem) HasIsPausedConsuming() bool {
	if o != nil && !isNil(o.IsPausedConsuming) {
		return true
	}

	return false
}

// SetIsPausedConsuming gets a reference to the given bool and assigns it to the IsPausedConsuming field.
func (o *TransMem) SetIsPausedConsuming(v bool) {
	o.IsPausedConsuming = &v
}

// GetPopulateTranslationStatus returns the PopulateTranslationStatus field value if set, zero value otherwise.
func (o *TransMem) GetPopulateTranslationStatus() TranslationStatus {
	if o == nil || isNil(o.PopulateTranslationStatus) {
		var ret TranslationStatus
		return ret
	}
	return *o.PopulateTranslationStatus
}

// GetPopulateTranslationStatusOk returns a tuple with the PopulateTranslationStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransMem) GetPopulateTranslationStatusOk() (*TranslationStatus, bool) {
	if o == nil || isNil(o.PopulateTranslationStatus) {
    return nil, false
	}
	return o.PopulateTranslationStatus, true
}

// HasPopulateTranslationStatus returns a boolean if a field has been set.
func (o *TransMem) HasPopulateTranslationStatus() bool {
	if o != nil && !isNil(o.PopulateTranslationStatus) {
		return true
	}

	return false
}

// SetPopulateTranslationStatus gets a reference to the given TranslationStatus and assigns it to the PopulateTranslationStatus field.
func (o *TransMem) SetPopulateTranslationStatus(v TranslationStatus) {
	o.PopulateTranslationStatus = &v
}

// GetContextLookup returns the ContextLookup field value if set, zero value otherwise.
func (o *TransMem) GetContextLookup() bool {
	if o == nil || isNil(o.ContextLookup) {
		var ret bool
		return ret
	}
	return *o.ContextLookup
}

// GetContextLookupOk returns a tuple with the ContextLookup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransMem) GetContextLookupOk() (*bool, bool) {
	if o == nil || isNil(o.ContextLookup) {
    return nil, false
	}
	return o.ContextLookup, true
}

// HasContextLookup returns a boolean if a field has been set.
func (o *TransMem) HasContextLookup() bool {
	if o != nil && !isNil(o.ContextLookup) {
		return true
	}

	return false
}

// SetContextLookup gets a reference to the given bool and assigns it to the ContextLookup field.
func (o *TransMem) SetContextLookup(v bool) {
	o.ContextLookup = &v
}

// GetName returns the Name field value
func (o *TransMem) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TransMem) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TransMem) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *TransMem) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransMem) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *TransMem) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *TransMem) SetDescription(v string) {
	o.Description = &v
}

// GetFuzzyMatch returns the FuzzyMatch field value if set, zero value otherwise.
func (o *TransMem) GetFuzzyMatch() bool {
	if o == nil || isNil(o.FuzzyMatch) {
		var ret bool
		return ret
	}
	return *o.FuzzyMatch
}

// GetFuzzyMatchOk returns a tuple with the FuzzyMatch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransMem) GetFuzzyMatchOk() (*bool, bool) {
	if o == nil || isNil(o.FuzzyMatch) {
    return nil, false
	}
	return o.FuzzyMatch, true
}

// HasFuzzyMatch returns a boolean if a field has been set.
func (o *TransMem) HasFuzzyMatch() bool {
	if o != nil && !isNil(o.FuzzyMatch) {
		return true
	}

	return false
}

// SetFuzzyMatch gets a reference to the given bool and assigns it to the FuzzyMatch field.
func (o *TransMem) SetFuzzyMatch(v bool) {
	o.FuzzyMatch = &v
}

// GetAllowAlternative returns the AllowAlternative field value if set, zero value otherwise.
func (o *TransMem) GetAllowAlternative() bool {
	if o == nil || isNil(o.AllowAlternative) {
		var ret bool
		return ret
	}
	return *o.AllowAlternative
}

// GetAllowAlternativeOk returns a tuple with the AllowAlternative field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransMem) GetAllowAlternativeOk() (*bool, bool) {
	if o == nil || isNil(o.AllowAlternative) {
    return nil, false
	}
	return o.AllowAlternative, true
}

// HasAllowAlternative returns a boolean if a field has been set.
func (o *TransMem) HasAllowAlternative() bool {
	if o != nil && !isNil(o.AllowAlternative) {
		return true
	}

	return false
}

// SetAllowAlternative gets a reference to the given bool and assigns it to the AllowAlternative field.
func (o *TransMem) SetAllowAlternative(v bool) {
	o.AllowAlternative = &v
}

// GetAllowAlternativeHasSameRecordId returns the AllowAlternativeHasSameRecordId field value if set, zero value otherwise.
func (o *TransMem) GetAllowAlternativeHasSameRecordId() bool {
	if o == nil || isNil(o.AllowAlternativeHasSameRecordId) {
		var ret bool
		return ret
	}
	return *o.AllowAlternativeHasSameRecordId
}

// GetAllowAlternativeHasSameRecordIdOk returns a tuple with the AllowAlternativeHasSameRecordId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransMem) GetAllowAlternativeHasSameRecordIdOk() (*bool, bool) {
	if o == nil || isNil(o.AllowAlternativeHasSameRecordId) {
    return nil, false
	}
	return o.AllowAlternativeHasSameRecordId, true
}

// HasAllowAlternativeHasSameRecordId returns a boolean if a field has been set.
func (o *TransMem) HasAllowAlternativeHasSameRecordId() bool {
	if o != nil && !isNil(o.AllowAlternativeHasSameRecordId) {
		return true
	}

	return false
}

// SetAllowAlternativeHasSameRecordId gets a reference to the given bool and assigns it to the AllowAlternativeHasSameRecordId field.
func (o *TransMem) SetAllowAlternativeHasSameRecordId(v bool) {
	o.AllowAlternativeHasSameRecordId = &v
}

func (o TransMem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.ProjectIds) {
		toSerialize["projectIds"] = o.ProjectIds
	}
	if !isNil(o.IsDisabled) {
		toSerialize["isDisabled"] = o.IsDisabled
	}
	if !isNil(o.IsPausedConsuming) {
		toSerialize["isPausedConsuming"] = o.IsPausedConsuming
	}
	if !isNil(o.PopulateTranslationStatus) {
		toSerialize["populateTranslationStatus"] = o.PopulateTranslationStatus
	}
	if !isNil(o.ContextLookup) {
		toSerialize["contextLookup"] = o.ContextLookup
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !isNil(o.FuzzyMatch) {
		toSerialize["fuzzyMatch"] = o.FuzzyMatch
	}
	if !isNil(o.AllowAlternative) {
		toSerialize["allowAlternative"] = o.AllowAlternative
	}
	if !isNil(o.AllowAlternativeHasSameRecordId) {
		toSerialize["allowAlternativeHasSameRecordId"] = o.AllowAlternativeHasSameRecordId
	}
	return json.Marshal(toSerialize)
}

type NullableTransMem struct {
	value *TransMem
	isSet bool
}

func (v NullableTransMem) Get() *TransMem {
	return v.value
}

func (v *NullableTransMem) Set(val *TransMem) {
	v.value = val
	v.isSet = true
}

func (v NullableTransMem) IsSet() bool {
	return v.isSet
}

func (v *NullableTransMem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransMem(val *TransMem) *NullableTransMem {
	return &NullableTransMem{value: val, isSet: true}
}

func (v NullableTransMem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransMem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


