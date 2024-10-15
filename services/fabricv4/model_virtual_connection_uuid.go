/*
Equinix Fabric API v4

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fabricv4

import (
	"encoding/json"
	"fmt"
)

// checks if the VirtualConnectionUuid type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VirtualConnectionUuid{}

// VirtualConnectionUuid UUID of the Fabric Connection Instance
type VirtualConnectionUuid struct {
	// Connection URI
	Href *string `json:"href,omitempty"`
	// Connection Type
	Type *string `json:"type,omitempty"`
	// Connection UUID.
	Uuid                 string `json:"uuid"`
	AdditionalProperties map[string]interface{}
}

type _VirtualConnectionUuid VirtualConnectionUuid

// NewVirtualConnectionUuid instantiates a new VirtualConnectionUuid object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualConnectionUuid(uuid string) *VirtualConnectionUuid {
	this := VirtualConnectionUuid{}
	this.Uuid = uuid
	return &this
}

// NewVirtualConnectionUuidWithDefaults instantiates a new VirtualConnectionUuid object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualConnectionUuidWithDefaults() *VirtualConnectionUuid {
	this := VirtualConnectionUuid{}
	return &this
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *VirtualConnectionUuid) GetHref() string {
	if o == nil || IsNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualConnectionUuid) GetHrefOk() (*string, bool) {
	if o == nil || IsNil(o.Href) {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *VirtualConnectionUuid) HasHref() bool {
	if o != nil && !IsNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *VirtualConnectionUuid) SetHref(v string) {
	o.Href = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *VirtualConnectionUuid) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualConnectionUuid) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *VirtualConnectionUuid) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *VirtualConnectionUuid) SetType(v string) {
	o.Type = &v
}

// GetUuid returns the Uuid field value
func (o *VirtualConnectionUuid) GetUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value
// and a boolean to check if the value has been set.
func (o *VirtualConnectionUuid) GetUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uuid, true
}

// SetUuid sets field value
func (o *VirtualConnectionUuid) SetUuid(v string) {
	o.Uuid = v
}

func (o VirtualConnectionUuid) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VirtualConnectionUuid) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Href) {
		toSerialize["href"] = o.Href
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	toSerialize["uuid"] = o.Uuid

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *VirtualConnectionUuid) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"uuid",
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

	varVirtualConnectionUuid := _VirtualConnectionUuid{}

	err = json.Unmarshal(data, &varVirtualConnectionUuid)

	if err != nil {
		return err
	}

	*o = VirtualConnectionUuid(varVirtualConnectionUuid)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "href")
		delete(additionalProperties, "type")
		delete(additionalProperties, "uuid")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVirtualConnectionUuid struct {
	value *VirtualConnectionUuid
	isSet bool
}

func (v NullableVirtualConnectionUuid) Get() *VirtualConnectionUuid {
	return v.value
}

func (v *NullableVirtualConnectionUuid) Set(val *VirtualConnectionUuid) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualConnectionUuid) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualConnectionUuid) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualConnectionUuid(val *VirtualConnectionUuid) *NullableVirtualConnectionUuid {
	return &NullableVirtualConnectionUuid{value: val, isSet: true}
}

func (v NullableVirtualConnectionUuid) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualConnectionUuid) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}