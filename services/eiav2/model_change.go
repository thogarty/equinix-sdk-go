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

// checks if the Change type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Change{}

// Change struct for Change
type Change struct {
	Href                 string     `json:"href"`
	Uuid                 string     `json:"uuid"`
	Type                 ChangeType `json:"type"`
	AdditionalProperties map[string]interface{}
}

type _Change Change

// NewChange instantiates a new Change object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChange(href string, uuid string, type_ ChangeType) *Change {
	this := Change{}
	this.Href = href
	this.Uuid = uuid
	this.Type = type_
	return &this
}

// NewChangeWithDefaults instantiates a new Change object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChangeWithDefaults() *Change {
	this := Change{}
	return &this
}

// GetHref returns the Href field value
func (o *Change) GetHref() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Href
}

// GetHrefOk returns a tuple with the Href field value
// and a boolean to check if the value has been set.
func (o *Change) GetHrefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Href, true
}

// SetHref sets field value
func (o *Change) SetHref(v string) {
	o.Href = v
}

// GetUuid returns the Uuid field value
func (o *Change) GetUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value
// and a boolean to check if the value has been set.
func (o *Change) GetUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uuid, true
}

// SetUuid sets field value
func (o *Change) SetUuid(v string) {
	o.Uuid = v
}

// GetType returns the Type field value
func (o *Change) GetType() ChangeType {
	if o == nil {
		var ret ChangeType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Change) GetTypeOk() (*ChangeType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Change) SetType(v ChangeType) {
	o.Type = v
}

func (o Change) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Change) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["href"] = o.Href
	toSerialize["uuid"] = o.Uuid
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Change) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"href",
		"uuid",
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

	varChange := _Change{}

	err = json.Unmarshal(data, &varChange)

	if err != nil {
		return err
	}

	*o = Change(varChange)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "href")
		delete(additionalProperties, "uuid")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableChange struct {
	value *Change
	isSet bool
}

func (v NullableChange) Get() *Change {
	return v.value
}

func (v *NullableChange) Set(val *Change) {
	v.value = val
	v.isSet = true
}

func (v NullableChange) IsSet() bool {
	return v.isSet
}

func (v *NullableChange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChange(val *Change) *NullableChange {
	return &NullableChange{value: val, isSet: true}
}

func (v NullableChange) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
