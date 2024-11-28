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

// checks if the RoutingProtocolRequestBase type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RoutingProtocolRequestBase{}

// RoutingProtocolRequestBase struct for RoutingProtocolRequestBase
type RoutingProtocolRequestBase struct {
	Tags []string `json:"tags,omitempty"`
	// Allowed values: - BGP - DIRECT - STATIC
	Type string `json:"type"`
	// Name of the routing protocol instance.
	Name *string `json:"name,omitempty"`
	// Description of the routing protocol instance
	Description          *string `json:"description,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RoutingProtocolRequestBase RoutingProtocolRequestBase

// NewRoutingProtocolRequestBase instantiates a new RoutingProtocolRequestBase object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoutingProtocolRequestBase(type_ string) *RoutingProtocolRequestBase {
	this := RoutingProtocolRequestBase{}
	this.Type = type_
	return &this
}

// NewRoutingProtocolRequestBaseWithDefaults instantiates a new RoutingProtocolRequestBase object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoutingProtocolRequestBaseWithDefaults() *RoutingProtocolRequestBase {
	this := RoutingProtocolRequestBase{}
	return &this
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *RoutingProtocolRequestBase) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoutingProtocolRequestBase) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *RoutingProtocolRequestBase) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *RoutingProtocolRequestBase) SetTags(v []string) {
	o.Tags = v
}

// GetType returns the Type field value
func (o *RoutingProtocolRequestBase) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *RoutingProtocolRequestBase) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *RoutingProtocolRequestBase) SetType(v string) {
	o.Type = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *RoutingProtocolRequestBase) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoutingProtocolRequestBase) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *RoutingProtocolRequestBase) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *RoutingProtocolRequestBase) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *RoutingProtocolRequestBase) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoutingProtocolRequestBase) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *RoutingProtocolRequestBase) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *RoutingProtocolRequestBase) SetDescription(v string) {
	o.Description = &v
}

func (o RoutingProtocolRequestBase) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RoutingProtocolRequestBase) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	toSerialize["type"] = o.Type
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RoutingProtocolRequestBase) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
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

	varRoutingProtocolRequestBase := _RoutingProtocolRequestBase{}

	err = json.Unmarshal(data, &varRoutingProtocolRequestBase)

	if err != nil {
		return err
	}

	*o = RoutingProtocolRequestBase(varRoutingProtocolRequestBase)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "tags")
		delete(additionalProperties, "type")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRoutingProtocolRequestBase struct {
	value *RoutingProtocolRequestBase
	isSet bool
}

func (v NullableRoutingProtocolRequestBase) Get() *RoutingProtocolRequestBase {
	return v.value
}

func (v *NullableRoutingProtocolRequestBase) Set(val *RoutingProtocolRequestBase) {
	v.value = val
	v.isSet = true
}

func (v NullableRoutingProtocolRequestBase) IsSet() bool {
	return v.isSet
}

func (v *NullableRoutingProtocolRequestBase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoutingProtocolRequestBase(val *RoutingProtocolRequestBase) *NullableRoutingProtocolRequestBase {
	return &NullableRoutingProtocolRequestBase{value: val, isSet: true}
}

func (v NullableRoutingProtocolRequestBase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoutingProtocolRequestBase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}