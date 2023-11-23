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
	"os"
)

// UploadZipRequest struct for UploadZipRequest
type UploadZipRequest struct {
	ColumnId string `json:"columnId"`
	FileMappings string `json:"fileMappings"`
	File *os.File `json:"file"`
}

// NewUploadZipRequest instantiates a new UploadZipRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUploadZipRequest(columnId string, fileMappings string, file *os.File) *UploadZipRequest {
	this := UploadZipRequest{}
	this.ColumnId = columnId
	this.FileMappings = fileMappings
	this.File = file
	return &this
}

// NewUploadZipRequestWithDefaults instantiates a new UploadZipRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUploadZipRequestWithDefaults() *UploadZipRequest {
	this := UploadZipRequest{}
	return &this
}

// GetColumnId returns the ColumnId field value
func (o *UploadZipRequest) GetColumnId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ColumnId
}

// GetColumnIdOk returns a tuple with the ColumnId field value
// and a boolean to check if the value has been set.
func (o *UploadZipRequest) GetColumnIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ColumnId, true
}

// SetColumnId sets field value
func (o *UploadZipRequest) SetColumnId(v string) {
	o.ColumnId = v
}

// GetFileMappings returns the FileMappings field value
func (o *UploadZipRequest) GetFileMappings() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FileMappings
}

// GetFileMappingsOk returns a tuple with the FileMappings field value
// and a boolean to check if the value has been set.
func (o *UploadZipRequest) GetFileMappingsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FileMappings, true
}

// SetFileMappings sets field value
func (o *UploadZipRequest) SetFileMappings(v string) {
	o.FileMappings = v
}

// GetFile returns the File field value
func (o *UploadZipRequest) GetFile() *os.File {
	if o == nil {
		var ret *os.File
		return ret
	}

	return o.File
}

// GetFileOk returns a tuple with the File field value
// and a boolean to check if the value has been set.
func (o *UploadZipRequest) GetFileOk() (**os.File, bool) {
	if o == nil {
		return nil, false
	}
	return &o.File, true
}

// SetFile sets field value
func (o *UploadZipRequest) SetFile(v *os.File) {
	o.File = v
}

func (o UploadZipRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["columnId"] = o.ColumnId
	}
	if true {
		toSerialize["fileMappings"] = o.FileMappings
	}
	if true {
		toSerialize["file"] = o.File
	}
	return json.Marshal(toSerialize)
}

type NullableUploadZipRequest struct {
	value *UploadZipRequest
	isSet bool
}

func (v NullableUploadZipRequest) Get() *UploadZipRequest {
	return v.value
}

func (v *NullableUploadZipRequest) Set(val *UploadZipRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUploadZipRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUploadZipRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUploadZipRequest(val *UploadZipRequest) *NullableUploadZipRequest {
	return &NullableUploadZipRequest{value: val, isSet: true}
}

func (v NullableUploadZipRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUploadZipRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


