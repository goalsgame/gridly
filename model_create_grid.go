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

// CreateGrid struct for CreateGrid
type CreateGrid struct {
	Id *string `json:"id,omitempty"`
	Name string `json:"name"`
	TemplateGridId *string `json:"templateGridId,omitempty"`
	RecordIdentifierType *string `json:"recordIdentifierType,omitempty"`
	Columns []CreateColumn `json:"columns,omitempty"`
	Metadata *map[string]string `json:"metadata,omitempty"`
}

// NewCreateGrid instantiates a new CreateGrid object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateGrid(name string) *CreateGrid {
	this := CreateGrid{}
	this.Name = name
	return &this
}

// NewCreateGridWithDefaults instantiates a new CreateGrid object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateGridWithDefaults() *CreateGrid {
	this := CreateGrid{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CreateGrid) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateGrid) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CreateGrid) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CreateGrid) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value
func (o *CreateGrid) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateGrid) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateGrid) SetName(v string) {
	o.Name = v
}

// GetTemplateGridId returns the TemplateGridId field value if set, zero value otherwise.
func (o *CreateGrid) GetTemplateGridId() string {
	if o == nil || o.TemplateGridId == nil {
		var ret string
		return ret
	}
	return *o.TemplateGridId
}

// GetTemplateGridIdOk returns a tuple with the TemplateGridId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateGrid) GetTemplateGridIdOk() (*string, bool) {
	if o == nil || o.TemplateGridId == nil {
		return nil, false
	}
	return o.TemplateGridId, true
}

// HasTemplateGridId returns a boolean if a field has been set.
func (o *CreateGrid) HasTemplateGridId() bool {
	if o != nil && o.TemplateGridId != nil {
		return true
	}

	return false
}

// SetTemplateGridId gets a reference to the given string and assigns it to the TemplateGridId field.
func (o *CreateGrid) SetTemplateGridId(v string) {
	o.TemplateGridId = &v
}

// GetRecordIdentifierType returns the RecordIdentifierType field value if set, zero value otherwise.
func (o *CreateGrid) GetRecordIdentifierType() string {
	if o == nil || o.RecordIdentifierType == nil {
		var ret string
		return ret
	}
	return *o.RecordIdentifierType
}

// GetRecordIdentifierTypeOk returns a tuple with the RecordIdentifierType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateGrid) GetRecordIdentifierTypeOk() (*string, bool) {
	if o == nil || o.RecordIdentifierType == nil {
		return nil, false
	}
	return o.RecordIdentifierType, true
}

// HasRecordIdentifierType returns a boolean if a field has been set.
func (o *CreateGrid) HasRecordIdentifierType() bool {
	if o != nil && o.RecordIdentifierType != nil {
		return true
	}

	return false
}

// SetRecordIdentifierType gets a reference to the given string and assigns it to the RecordIdentifierType field.
func (o *CreateGrid) SetRecordIdentifierType(v string) {
	o.RecordIdentifierType = &v
}

// GetColumns returns the Columns field value if set, zero value otherwise.
func (o *CreateGrid) GetColumns() []CreateColumn {
	if o == nil || o.Columns == nil {
		var ret []CreateColumn
		return ret
	}
	return o.Columns
}

// GetColumnsOk returns a tuple with the Columns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateGrid) GetColumnsOk() ([]CreateColumn, bool) {
	if o == nil || o.Columns == nil {
		return nil, false
	}
	return o.Columns, true
}

// HasColumns returns a boolean if a field has been set.
func (o *CreateGrid) HasColumns() bool {
	if o != nil && o.Columns != nil {
		return true
	}

	return false
}

// SetColumns gets a reference to the given []CreateColumn and assigns it to the Columns field.
func (o *CreateGrid) SetColumns(v []CreateColumn) {
	o.Columns = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *CreateGrid) GetMetadata() map[string]string {
	if o == nil || o.Metadata == nil {
		var ret map[string]string
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateGrid) GetMetadataOk() (*map[string]string, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *CreateGrid) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]string and assigns it to the Metadata field.
func (o *CreateGrid) SetMetadata(v map[string]string) {
	o.Metadata = &v
}

func (o CreateGrid) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.TemplateGridId != nil {
		toSerialize["templateGridId"] = o.TemplateGridId
	}
	if o.RecordIdentifierType != nil {
		toSerialize["recordIdentifierType"] = o.RecordIdentifierType
	}
	if o.Columns != nil {
		toSerialize["columns"] = o.Columns
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	return json.Marshal(toSerialize)
}

type NullableCreateGrid struct {
	value *CreateGrid
	isSet bool
}

func (v NullableCreateGrid) Get() *CreateGrid {
	return v.value
}

func (v *NullableCreateGrid) Set(val *CreateGrid) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateGrid) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateGrid) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateGrid(val *CreateGrid) *NullableCreateGrid {
	return &NullableCreateGrid{value: val, isSet: true}
}

func (v NullableCreateGrid) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateGrid) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


