/*
Equinix Internet Access API

Equinix Internet Access provides direct access to the Internet with scalable bandwidth options in IBX data centers.

API version: 2.0
Contact: api-support@equinix.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eiav2

import (
	"encoding/json"
	"fmt"
)

// checks if the ProjectReference type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectReference{}

// ProjectReference struct for ProjectReference
type ProjectReference struct {
	Href string `json:"href"`
	// Customer Resource Hierarchy Project identifier
	ProjectId            string `json:"projectId"`
	AdditionalProperties map[string]interface{}
}

type _ProjectReference ProjectReference

// NewProjectReference instantiates a new ProjectReference object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectReference(href string, projectId string) *ProjectReference {
	this := ProjectReference{}
	this.Href = href
	this.ProjectId = projectId
	return &this
}

// NewProjectReferenceWithDefaults instantiates a new ProjectReference object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectReferenceWithDefaults() *ProjectReference {
	this := ProjectReference{}
	return &this
}

// GetHref returns the Href field value
func (o *ProjectReference) GetHref() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Href
}

// GetHrefOk returns a tuple with the Href field value
// and a boolean to check if the value has been set.
func (o *ProjectReference) GetHrefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Href, true
}

// SetHref sets field value
func (o *ProjectReference) SetHref(v string) {
	o.Href = v
}

// GetProjectId returns the ProjectId field value
func (o *ProjectReference) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *ProjectReference) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *ProjectReference) SetProjectId(v string) {
	o.ProjectId = v
}

func (o ProjectReference) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectReference) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["href"] = o.Href
	toSerialize["projectId"] = o.ProjectId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ProjectReference) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"href",
		"projectId",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varProjectReference := _ProjectReference{}

	err = json.Unmarshal(data, &varProjectReference)

	if err != nil {
		return err
	}

	*o = ProjectReference(varProjectReference)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "href")
		delete(additionalProperties, "projectId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableProjectReference struct {
	value *ProjectReference
	isSet bool
}

func (v NullableProjectReference) Get() *ProjectReference {
	return v.value
}

func (v *NullableProjectReference) Set(val *ProjectReference) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectReference) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectReference) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectReference(val *ProjectReference) *NullableProjectReference {
	return &NullableProjectReference{value: val, isSet: true}
}

func (v NullableProjectReference) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectReference) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
