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

// View struct for View
type View struct {
	Id *string `json:"id,omitempty"`
	Columns []ViewColumn `json:"columns,omitempty"`
	GridId *string `json:"gridId,omitempty"`
	GridStatus *string `json:"gridStatus,omitempty"`
	Name *string `json:"name,omitempty"`
	Records []Record `json:"records,omitempty"`
}

// NewView instantiates a new View object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewView() *View {
	this := View{}
	return &this
}

// NewViewWithDefaults instantiates a new View object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewViewWithDefaults() *View {
	this := View{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *View) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *View) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *View) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *View) SetId(v string) {
	o.Id = &v
}

// GetColumns returns the Columns field value if set, zero value otherwise.
func (o *View) GetColumns() []ViewColumn {
	if o == nil || o.Columns == nil {
		var ret []ViewColumn
		return ret
	}
	return o.Columns
}

// GetColumnsOk returns a tuple with the Columns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *View) GetColumnsOk() ([]ViewColumn, bool) {
	if o == nil || o.Columns == nil {
		return nil, false
	}
	return o.Columns, true
}

// HasColumns returns a boolean if a field has been set.
func (o *View) HasColumns() bool {
	if o != nil && o.Columns != nil {
		return true
	}

	return false
}

// SetColumns gets a reference to the given []ViewColumn and assigns it to the Columns field.
func (o *View) SetColumns(v []ViewColumn) {
	o.Columns = v
}

// GetGridId returns the GridId field value if set, zero value otherwise.
func (o *View) GetGridId() string {
	if o == nil || o.GridId == nil {
		var ret string
		return ret
	}
	return *o.GridId
}

// GetGridIdOk returns a tuple with the GridId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *View) GetGridIdOk() (*string, bool) {
	if o == nil || o.GridId == nil {
		return nil, false
	}
	return o.GridId, true
}

// HasGridId returns a boolean if a field has been set.
func (o *View) HasGridId() bool {
	if o != nil && o.GridId != nil {
		return true
	}

	return false
}

// SetGridId gets a reference to the given string and assigns it to the GridId field.
func (o *View) SetGridId(v string) {
	o.GridId = &v
}

// GetGridStatus returns the GridStatus field value if set, zero value otherwise.
func (o *View) GetGridStatus() string {
	if o == nil || o.GridStatus == nil {
		var ret string
		return ret
	}
	return *o.GridStatus
}

// GetGridStatusOk returns a tuple with the GridStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *View) GetGridStatusOk() (*string, bool) {
	if o == nil || o.GridStatus == nil {
		return nil, false
	}
	return o.GridStatus, true
}

// HasGridStatus returns a boolean if a field has been set.
func (o *View) HasGridStatus() bool {
	if o != nil && o.GridStatus != nil {
		return true
	}

	return false
}

// SetGridStatus gets a reference to the given string and assigns it to the GridStatus field.
func (o *View) SetGridStatus(v string) {
	o.GridStatus = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *View) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *View) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *View) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *View) SetName(v string) {
	o.Name = &v
}

// GetRecords returns the Records field value if set, zero value otherwise.
func (o *View) GetRecords() []Record {
	if o == nil || o.Records == nil {
		var ret []Record
		return ret
	}
	return o.Records
}

// GetRecordsOk returns a tuple with the Records field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *View) GetRecordsOk() ([]Record, bool) {
	if o == nil || o.Records == nil {
		return nil, false
	}
	return o.Records, true
}

// HasRecords returns a boolean if a field has been set.
func (o *View) HasRecords() bool {
	if o != nil && o.Records != nil {
		return true
	}

	return false
}

// SetRecords gets a reference to the given []Record and assigns it to the Records field.
func (o *View) SetRecords(v []Record) {
	o.Records = v
}

func (o View) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Columns != nil {
		toSerialize["columns"] = o.Columns
	}
	if o.GridId != nil {
		toSerialize["gridId"] = o.GridId
	}
	if o.GridStatus != nil {
		toSerialize["gridStatus"] = o.GridStatus
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Records != nil {
		toSerialize["records"] = o.Records
	}
	return json.Marshal(toSerialize)
}

type NullableView struct {
	value *View
	isSet bool
}

func (v NullableView) Get() *View {
	return v.value
}

func (v *NullableView) Set(val *View) {
	v.value = val
	v.isSet = true
}

func (v NullableView) IsSet() bool {
	return v.isSet
}

func (v *NullableView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableView(val *View) *NullableView {
	return &NullableView{value: val, isSet: true}
}

func (v NullableView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


