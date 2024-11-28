/*
Equinix Internet Access API

Equinix Internet Access provides direct access to the Internet with scalable bandwidth options in IBX data centers.

Contact: api-support@equinix.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eiav2

import (
	"encoding/json"
	"fmt"
)

// checks if the ProjectReadModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectReadModel{}

// ProjectReadModel struct for ProjectReadModel
type ProjectReadModel struct {
	// Project URI
	Href string `json:"href"`
	// identifier of the project
	ProjectId string `json:"projectId"`
	// name of the project
	ProjectName          string `json:"projectName"`
	ParentOrganizationId string `json:"parentOrganizationId"`
	// Name of the parent organization
	ParentOrganizationName string `json:"parentOrganizationName"`
	AdditionalProperties   map[string]interface{}
}

type _ProjectReadModel ProjectReadModel

// NewProjectReadModel instantiates a new ProjectReadModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectReadModel(href string, projectId string, projectName string, parentOrganizationId string, parentOrganizationName string) *ProjectReadModel {
	this := ProjectReadModel{}
	this.Href = href
	this.ProjectId = projectId
	this.ProjectName = projectName
	this.ParentOrganizationId = parentOrganizationId
	this.ParentOrganizationName = parentOrganizationName
	return &this
}

// NewProjectReadModelWithDefaults instantiates a new ProjectReadModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectReadModelWithDefaults() *ProjectReadModel {
	this := ProjectReadModel{}
	return &this
}

// GetHref returns the Href field value
func (o *ProjectReadModel) GetHref() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Href
}

// GetHrefOk returns a tuple with the Href field value
// and a boolean to check if the value has been set.
func (o *ProjectReadModel) GetHrefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Href, true
}

// SetHref sets field value
func (o *ProjectReadModel) SetHref(v string) {
	o.Href = v
}

// GetProjectId returns the ProjectId field value
func (o *ProjectReadModel) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *ProjectReadModel) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *ProjectReadModel) SetProjectId(v string) {
	o.ProjectId = v
}

// GetProjectName returns the ProjectName field value
func (o *ProjectReadModel) GetProjectName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectName
}

// GetProjectNameOk returns a tuple with the ProjectName field value
// and a boolean to check if the value has been set.
func (o *ProjectReadModel) GetProjectNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectName, true
}

// SetProjectName sets field value
func (o *ProjectReadModel) SetProjectName(v string) {
	o.ProjectName = v
}

// GetParentOrganizationId returns the ParentOrganizationId field value
func (o *ProjectReadModel) GetParentOrganizationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ParentOrganizationId
}

// GetParentOrganizationIdOk returns a tuple with the ParentOrganizationId field value
// and a boolean to check if the value has been set.
func (o *ProjectReadModel) GetParentOrganizationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ParentOrganizationId, true
}

// SetParentOrganizationId sets field value
func (o *ProjectReadModel) SetParentOrganizationId(v string) {
	o.ParentOrganizationId = v
}

// GetParentOrganizationName returns the ParentOrganizationName field value
func (o *ProjectReadModel) GetParentOrganizationName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ParentOrganizationName
}

// GetParentOrganizationNameOk returns a tuple with the ParentOrganizationName field value
// and a boolean to check if the value has been set.
func (o *ProjectReadModel) GetParentOrganizationNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ParentOrganizationName, true
}

// SetParentOrganizationName sets field value
func (o *ProjectReadModel) SetParentOrganizationName(v string) {
	o.ParentOrganizationName = v
}

func (o ProjectReadModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectReadModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["href"] = o.Href
	toSerialize["projectId"] = o.ProjectId
	toSerialize["projectName"] = o.ProjectName
	toSerialize["parentOrganizationId"] = o.ParentOrganizationId
	toSerialize["parentOrganizationName"] = o.ParentOrganizationName

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ProjectReadModel) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"href",
		"projectId",
		"projectName",
		"parentOrganizationId",
		"parentOrganizationName",
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

	varProjectReadModel := _ProjectReadModel{}

	err = json.Unmarshal(data, &varProjectReadModel)

	if err != nil {
		return err
	}

	*o = ProjectReadModel(varProjectReadModel)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "href")
		delete(additionalProperties, "projectId")
		delete(additionalProperties, "projectName")
		delete(additionalProperties, "parentOrganizationId")
		delete(additionalProperties, "parentOrganizationName")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableProjectReadModel struct {
	value *ProjectReadModel
	isSet bool
}

func (v NullableProjectReadModel) Get() *ProjectReadModel {
	return v.value
}

func (v *NullableProjectReadModel) Set(val *ProjectReadModel) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectReadModel) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectReadModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectReadModel(val *ProjectReadModel) *NullableProjectReadModel {
	return &NullableProjectReadModel{value: val, isSet: true}
}

func (v NullableProjectReadModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectReadModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
