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

// UploadedFile struct for UploadedFile
type UploadedFile struct {
	Id *string `json:"id,omitempty"`
	OriginalName *string `json:"originalName,omitempty"`
	ContentType *string `json:"contentType,omitempty"`
	Size *int64 `json:"size,omitempty"`
}

// NewUploadedFile instantiates a new UploadedFile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUploadedFile() *UploadedFile {
	this := UploadedFile{}
	return &this
}

// NewUploadedFileWithDefaults instantiates a new UploadedFile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUploadedFileWithDefaults() *UploadedFile {
	this := UploadedFile{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UploadedFile) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UploadedFile) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UploadedFile) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *UploadedFile) SetId(v string) {
	o.Id = &v
}

// GetOriginalName returns the OriginalName field value if set, zero value otherwise.
func (o *UploadedFile) GetOriginalName() string {
	if o == nil || o.OriginalName == nil {
		var ret string
		return ret
	}
	return *o.OriginalName
}

// GetOriginalNameOk returns a tuple with the OriginalName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UploadedFile) GetOriginalNameOk() (*string, bool) {
	if o == nil || o.OriginalName == nil {
		return nil, false
	}
	return o.OriginalName, true
}

// HasOriginalName returns a boolean if a field has been set.
func (o *UploadedFile) HasOriginalName() bool {
	if o != nil && o.OriginalName != nil {
		return true
	}

	return false
}

// SetOriginalName gets a reference to the given string and assigns it to the OriginalName field.
func (o *UploadedFile) SetOriginalName(v string) {
	o.OriginalName = &v
}

// GetContentType returns the ContentType field value if set, zero value otherwise.
func (o *UploadedFile) GetContentType() string {
	if o == nil || o.ContentType == nil {
		var ret string
		return ret
	}
	return *o.ContentType
}

// GetContentTypeOk returns a tuple with the ContentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UploadedFile) GetContentTypeOk() (*string, bool) {
	if o == nil || o.ContentType == nil {
		return nil, false
	}
	return o.ContentType, true
}

// HasContentType returns a boolean if a field has been set.
func (o *UploadedFile) HasContentType() bool {
	if o != nil && o.ContentType != nil {
		return true
	}

	return false
}

// SetContentType gets a reference to the given string and assigns it to the ContentType field.
func (o *UploadedFile) SetContentType(v string) {
	o.ContentType = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *UploadedFile) GetSize() int64 {
	if o == nil || o.Size == nil {
		var ret int64
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UploadedFile) GetSizeOk() (*int64, bool) {
	if o == nil || o.Size == nil {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *UploadedFile) HasSize() bool {
	if o != nil && o.Size != nil {
		return true
	}

	return false
}

// SetSize gets a reference to the given int64 and assigns it to the Size field.
func (o *UploadedFile) SetSize(v int64) {
	o.Size = &v
}

func (o UploadedFile) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.OriginalName != nil {
		toSerialize["originalName"] = o.OriginalName
	}
	if o.ContentType != nil {
		toSerialize["contentType"] = o.ContentType
	}
	if o.Size != nil {
		toSerialize["size"] = o.Size
	}
	return json.Marshal(toSerialize)
}

type NullableUploadedFile struct {
	value *UploadedFile
	isSet bool
}

func (v NullableUploadedFile) Get() *UploadedFile {
	return v.value
}

func (v *NullableUploadedFile) Set(val *UploadedFile) {
	v.value = val
	v.isSet = true
}

func (v NullableUploadedFile) IsSet() bool {
	return v.isSet
}

func (v *NullableUploadedFile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUploadedFile(val *UploadedFile) *NullableUploadedFile {
	return &NullableUploadedFile{value: val, isSet: true}
}

func (v NullableUploadedFile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUploadedFile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


