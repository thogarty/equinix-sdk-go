/*
Equinix Fabric API v4

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fabricv4

import (
	"encoding/json"
)

// checks if the PortInterface type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PortInterface{}

// PortInterface Port interface
type PortInterface struct {
	// Port interface type
	Type                 *string `json:"type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PortInterface PortInterface

// NewPortInterface instantiates a new PortInterface object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPortInterface() *PortInterface {
	this := PortInterface{}
	return &this
}

// NewPortInterfaceWithDefaults instantiates a new PortInterface object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPortInterfaceWithDefaults() *PortInterface {
	this := PortInterface{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PortInterface) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortInterface) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PortInterface) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *PortInterface) SetType(v string) {
	o.Type = &v
}

func (o PortInterface) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PortInterface) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PortInterface) UnmarshalJSON(data []byte) (err error) {
	varPortInterface := _PortInterface{}

	err = json.Unmarshal(data, &varPortInterface)

	if err != nil {
		return err
	}

	*o = PortInterface(varPortInterface)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePortInterface struct {
	value *PortInterface
	isSet bool
}

func (v NullablePortInterface) Get() *PortInterface {
	return v.value
}

func (v *NullablePortInterface) Set(val *PortInterface) {
	v.value = val
	v.isSet = true
}

func (v NullablePortInterface) IsSet() bool {
	return v.isSet
}

func (v *NullablePortInterface) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePortInterface(val *PortInterface) *NullablePortInterface {
	return &NullablePortInterface{value: val, isSet: true}
}

func (v NullablePortInterface) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePortInterface) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
