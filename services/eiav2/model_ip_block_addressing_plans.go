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

// checks if the IpBlockAddressingPlans type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IpBlockAddressingPlans{}

// IpBlockAddressingPlans struct for IpBlockAddressingPlans
type IpBlockAddressingPlans struct {
	Size int32 `json:"size"`
	// The purpose of IP Subnet
	Purpose string `json:"purpose"`
	// Number of ip addresses to be used immediatelly
	Immediate int32 `json:"immediate"`
	// Number of ip addresses to be used after 3 months
	AfterThreeMonths     int32 `json:"afterThreeMonths"`
	AdditionalProperties map[string]interface{}
}

type _IpBlockAddressingPlans IpBlockAddressingPlans

// NewIpBlockAddressingPlans instantiates a new IpBlockAddressingPlans object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIpBlockAddressingPlans(size int32, purpose string, immediate int32, afterThreeMonths int32) *IpBlockAddressingPlans {
	this := IpBlockAddressingPlans{}
	this.Size = size
	this.Purpose = purpose
	this.Immediate = immediate
	this.AfterThreeMonths = afterThreeMonths
	return &this
}

// NewIpBlockAddressingPlansWithDefaults instantiates a new IpBlockAddressingPlans object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIpBlockAddressingPlansWithDefaults() *IpBlockAddressingPlans {
	this := IpBlockAddressingPlans{}
	return &this
}

// GetSize returns the Size field value
func (o *IpBlockAddressingPlans) GetSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Size
}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
func (o *IpBlockAddressingPlans) GetSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Size, true
}

// SetSize sets field value
func (o *IpBlockAddressingPlans) SetSize(v int32) {
	o.Size = v
}

// GetPurpose returns the Purpose field value
func (o *IpBlockAddressingPlans) GetPurpose() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Purpose
}

// GetPurposeOk returns a tuple with the Purpose field value
// and a boolean to check if the value has been set.
func (o *IpBlockAddressingPlans) GetPurposeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Purpose, true
}

// SetPurpose sets field value
func (o *IpBlockAddressingPlans) SetPurpose(v string) {
	o.Purpose = v
}

// GetImmediate returns the Immediate field value
func (o *IpBlockAddressingPlans) GetImmediate() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Immediate
}

// GetImmediateOk returns a tuple with the Immediate field value
// and a boolean to check if the value has been set.
func (o *IpBlockAddressingPlans) GetImmediateOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Immediate, true
}

// SetImmediate sets field value
func (o *IpBlockAddressingPlans) SetImmediate(v int32) {
	o.Immediate = v
}

// GetAfterThreeMonths returns the AfterThreeMonths field value
func (o *IpBlockAddressingPlans) GetAfterThreeMonths() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AfterThreeMonths
}

// GetAfterThreeMonthsOk returns a tuple with the AfterThreeMonths field value
// and a boolean to check if the value has been set.
func (o *IpBlockAddressingPlans) GetAfterThreeMonthsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AfterThreeMonths, true
}

// SetAfterThreeMonths sets field value
func (o *IpBlockAddressingPlans) SetAfterThreeMonths(v int32) {
	o.AfterThreeMonths = v
}

func (o IpBlockAddressingPlans) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IpBlockAddressingPlans) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["size"] = o.Size
	toSerialize["purpose"] = o.Purpose
	toSerialize["immediate"] = o.Immediate
	toSerialize["afterThreeMonths"] = o.AfterThreeMonths

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *IpBlockAddressingPlans) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"size",
		"purpose",
		"immediate",
		"afterThreeMonths",
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

	varIpBlockAddressingPlans := _IpBlockAddressingPlans{}

	err = json.Unmarshal(data, &varIpBlockAddressingPlans)

	if err != nil {
		return err
	}

	*o = IpBlockAddressingPlans(varIpBlockAddressingPlans)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "size")
		delete(additionalProperties, "purpose")
		delete(additionalProperties, "immediate")
		delete(additionalProperties, "afterThreeMonths")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIpBlockAddressingPlans struct {
	value *IpBlockAddressingPlans
	isSet bool
}

func (v NullableIpBlockAddressingPlans) Get() *IpBlockAddressingPlans {
	return v.value
}

func (v *NullableIpBlockAddressingPlans) Set(val *IpBlockAddressingPlans) {
	v.value = val
	v.isSet = true
}

func (v NullableIpBlockAddressingPlans) IsSet() bool {
	return v.isSet
}

func (v *NullableIpBlockAddressingPlans) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIpBlockAddressingPlans(val *IpBlockAddressingPlans) *NullableIpBlockAddressingPlans {
	return &NullableIpBlockAddressingPlans{value: val, isSet: true}
}

func (v NullableIpBlockAddressingPlans) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIpBlockAddressingPlans) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
