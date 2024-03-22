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

// CreateGlossary struct for CreateGlossary
type CreateGlossary struct {
	Name string `json:"name"`
	Description *string `json:"description,omitempty"`
	Langs []string `json:"langs,omitempty"`
	Projects []GlossaryProject `json:"projects,omitempty"`
}

// NewCreateGlossary instantiates a new CreateGlossary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateGlossary(name string) *CreateGlossary {
	this := CreateGlossary{}
	this.Name = name
	return &this
}

// NewCreateGlossaryWithDefaults instantiates a new CreateGlossary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateGlossaryWithDefaults() *CreateGlossary {
	this := CreateGlossary{}
	return &this
}

// GetName returns the Name field value
func (o *CreateGlossary) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateGlossary) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateGlossary) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreateGlossary) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateGlossary) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateGlossary) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreateGlossary) SetDescription(v string) {
	o.Description = &v
}

// GetLangs returns the Langs field value if set, zero value otherwise.
func (o *CreateGlossary) GetLangs() []string {
	if o == nil || o.Langs == nil {
		var ret []string
		return ret
	}
	return o.Langs
}

// GetLangsOk returns a tuple with the Langs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateGlossary) GetLangsOk() ([]string, bool) {
	if o == nil || o.Langs == nil {
		return nil, false
	}
	return o.Langs, true
}

// HasLangs returns a boolean if a field has been set.
func (o *CreateGlossary) HasLangs() bool {
	if o != nil && o.Langs != nil {
		return true
	}

	return false
}

// SetLangs gets a reference to the given []string and assigns it to the Langs field.
func (o *CreateGlossary) SetLangs(v []string) {
	o.Langs = v
}

// GetProjects returns the Projects field value if set, zero value otherwise.
func (o *CreateGlossary) GetProjects() []GlossaryProject {
	if o == nil || o.Projects == nil {
		var ret []GlossaryProject
		return ret
	}
	return o.Projects
}

// GetProjectsOk returns a tuple with the Projects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateGlossary) GetProjectsOk() ([]GlossaryProject, bool) {
	if o == nil || o.Projects == nil {
		return nil, false
	}
	return o.Projects, true
}

// HasProjects returns a boolean if a field has been set.
func (o *CreateGlossary) HasProjects() bool {
	if o != nil && o.Projects != nil {
		return true
	}

	return false
}

// SetProjects gets a reference to the given []GlossaryProject and assigns it to the Projects field.
func (o *CreateGlossary) SetProjects(v []GlossaryProject) {
	o.Projects = v
}

func (o CreateGlossary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Langs != nil {
		toSerialize["langs"] = o.Langs
	}
	if o.Projects != nil {
		toSerialize["projects"] = o.Projects
	}
	return json.Marshal(toSerialize)
}

type NullableCreateGlossary struct {
	value *CreateGlossary
	isSet bool
}

func (v NullableCreateGlossary) Get() *CreateGlossary {
	return v.value
}

func (v *NullableCreateGlossary) Set(val *CreateGlossary) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateGlossary) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateGlossary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateGlossary(val *CreateGlossary) *NullableCreateGlossary {
	return &NullableCreateGlossary{value: val, isSet: true}
}

func (v NullableCreateGlossary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateGlossary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


