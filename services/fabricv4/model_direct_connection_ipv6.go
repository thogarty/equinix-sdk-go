/*
Equinix Fabric API v4

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fabricv4

import (
	"encoding/json"
	"fmt"
)

// checks if the DirectConnectionIpv6 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DirectConnectionIpv6{}

// DirectConnectionIpv6 struct for DirectConnectionIpv6
type DirectConnectionIpv6 struct {
	// Equinix side Interface IP address
	EquinixIfaceIp       string `json:"equinixIfaceIp"`
	AdditionalProperties map[string]interface{}
}

type _DirectConnectionIpv6 DirectConnectionIpv6

// NewDirectConnectionIpv6 instantiates a new DirectConnectionIpv6 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDirectConnectionIpv6(equinixIfaceIp string) *DirectConnectionIpv6 {
	this := DirectConnectionIpv6{}
	this.EquinixIfaceIp = equinixIfaceIp
	return &this
}

// NewDirectConnectionIpv6WithDefaults instantiates a new DirectConnectionIpv6 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDirectConnectionIpv6WithDefaults() *DirectConnectionIpv6 {
	this := DirectConnectionIpv6{}
	return &this
}

// GetEquinixIfaceIp returns the EquinixIfaceIp field value
func (o *DirectConnectionIpv6) GetEquinixIfaceIp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EquinixIfaceIp
}

// GetEquinixIfaceIpOk returns a tuple with the EquinixIfaceIp field value
// and a boolean to check if the value has been set.
func (o *DirectConnectionIpv6) GetEquinixIfaceIpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EquinixIfaceIp, true
}

// SetEquinixIfaceIp sets field value
func (o *DirectConnectionIpv6) SetEquinixIfaceIp(v string) {
	o.EquinixIfaceIp = v
}

func (o DirectConnectionIpv6) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DirectConnectionIpv6) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["equinixIfaceIp"] = o.EquinixIfaceIp

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DirectConnectionIpv6) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"equinixIfaceIp",
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

	varDirectConnectionIpv6 := _DirectConnectionIpv6{}

	err = json.Unmarshal(data, &varDirectConnectionIpv6)

	if err != nil {
		return err
	}

	*o = DirectConnectionIpv6(varDirectConnectionIpv6)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "equinixIfaceIp")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDirectConnectionIpv6 struct {
	value *DirectConnectionIpv6
	isSet bool
}

func (v NullableDirectConnectionIpv6) Get() *DirectConnectionIpv6 {
	return v.value
}

func (v *NullableDirectConnectionIpv6) Set(val *DirectConnectionIpv6) {
	v.value = val
	v.isSet = true
}

func (v NullableDirectConnectionIpv6) IsSet() bool {
	return v.isSet
}

func (v *NullableDirectConnectionIpv6) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDirectConnectionIpv6(val *DirectConnectionIpv6) *NullableDirectConnectionIpv6 {
	return &NullableDirectConnectionIpv6{value: val, isSet: true}
}

func (v NullableDirectConnectionIpv6) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDirectConnectionIpv6) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
