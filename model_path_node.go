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

// PathNode struct for PathNode
type PathNode struct {
	Name *string `json:"name,omitempty"`
	ParentPath *string `json:"parentPath,omitempty"`
}

// NewPathNode instantiates a new PathNode object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPathNode() *PathNode {
	this := PathNode{}
	return &this
}

// NewPathNodeWithDefaults instantiates a new PathNode object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPathNodeWithDefaults() *PathNode {
	this := PathNode{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PathNode) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PathNode) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PathNode) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PathNode) SetName(v string) {
	o.Name = &v
}

// GetParentPath returns the ParentPath field value if set, zero value otherwise.
func (o *PathNode) GetParentPath() string {
	if o == nil || o.ParentPath == nil {
		var ret string
		return ret
	}
	return *o.ParentPath
}

// GetParentPathOk returns a tuple with the ParentPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PathNode) GetParentPathOk() (*string, bool) {
	if o == nil || o.ParentPath == nil {
		return nil, false
	}
	return o.ParentPath, true
}

// HasParentPath returns a boolean if a field has been set.
func (o *PathNode) HasParentPath() bool {
	if o != nil && o.ParentPath != nil {
		return true
	}

	return false
}

// SetParentPath gets a reference to the given string and assigns it to the ParentPath field.
func (o *PathNode) SetParentPath(v string) {
	o.ParentPath = &v
}

func (o PathNode) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.ParentPath != nil {
		toSerialize["parentPath"] = o.ParentPath
	}
	return json.Marshal(toSerialize)
}

type NullablePathNode struct {
	value *PathNode
	isSet bool
}

func (v NullablePathNode) Get() *PathNode {
	return v.value
}

func (v *NullablePathNode) Set(val *PathNode) {
	v.value = val
	v.isSet = true
}

func (v NullablePathNode) IsSet() bool {
	return v.isSet
}

func (v *NullablePathNode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePathNode(val *PathNode) *NullablePathNode {
	return &NullablePathNode{value: val, isSet: true}
}

func (v NullablePathNode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePathNode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


