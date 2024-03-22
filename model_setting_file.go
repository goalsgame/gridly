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

// SettingFile struct for SettingFile
type SettingFile struct {
	Category *FileCategory `json:"category,omitempty"`
	Files []UploadedFile `json:"files,omitempty"`
}

// NewSettingFile instantiates a new SettingFile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSettingFile() *SettingFile {
	this := SettingFile{}
	return &this
}

// NewSettingFileWithDefaults instantiates a new SettingFile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSettingFileWithDefaults() *SettingFile {
	this := SettingFile{}
	return &this
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *SettingFile) GetCategory() FileCategory {
	if o == nil || o.Category == nil {
		var ret FileCategory
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingFile) GetCategoryOk() (*FileCategory, bool) {
	if o == nil || o.Category == nil {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *SettingFile) HasCategory() bool {
	if o != nil && o.Category != nil {
		return true
	}

	return false
}

// SetCategory gets a reference to the given FileCategory and assigns it to the Category field.
func (o *SettingFile) SetCategory(v FileCategory) {
	o.Category = &v
}

// GetFiles returns the Files field value if set, zero value otherwise.
func (o *SettingFile) GetFiles() []UploadedFile {
	if o == nil || o.Files == nil {
		var ret []UploadedFile
		return ret
	}
	return o.Files
}

// GetFilesOk returns a tuple with the Files field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingFile) GetFilesOk() ([]UploadedFile, bool) {
	if o == nil || o.Files == nil {
		return nil, false
	}
	return o.Files, true
}

// HasFiles returns a boolean if a field has been set.
func (o *SettingFile) HasFiles() bool {
	if o != nil && o.Files != nil {
		return true
	}

	return false
}

// SetFiles gets a reference to the given []UploadedFile and assigns it to the Files field.
func (o *SettingFile) SetFiles(v []UploadedFile) {
	o.Files = v
}

func (o SettingFile) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Category != nil {
		toSerialize["category"] = o.Category
	}
	if o.Files != nil {
		toSerialize["files"] = o.Files
	}
	return json.Marshal(toSerialize)
}

type NullableSettingFile struct {
	value *SettingFile
	isSet bool
}

func (v NullableSettingFile) Get() *SettingFile {
	return v.value
}

func (v *NullableSettingFile) Set(val *SettingFile) {
	v.value = val
	v.isSet = true
}

func (v NullableSettingFile) IsSet() bool {
	return v.isSet
}

func (v *NullableSettingFile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSettingFile(val *SettingFile) *NullableSettingFile {
	return &NullableSettingFile{value: val, isSet: true}
}

func (v NullableSettingFile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSettingFile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


