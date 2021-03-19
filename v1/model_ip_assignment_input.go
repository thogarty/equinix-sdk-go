/*
 * Metal API
 *
 * This is the API for Equinix Metal Product. Interact with your devices, user account, and projects.
 *
 * API version: 1.0.0
 * Contact: support@equinixmetal.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1

import (
	"encoding/json"
)

// IPAssignmentInput struct for IPAssignmentInput
type IPAssignmentInput struct {
	Address string `json:"address"`
	Manageable *bool `json:"manageable,omitempty"`
	Customdata *map[string]interface{} `json:"customdata,omitempty"`
}

// NewIPAssignmentInput instantiates a new IPAssignmentInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIPAssignmentInput(address string) *IPAssignmentInput {
	this := IPAssignmentInput{}
	this.Address = address
	return &this
}

// NewIPAssignmentInputWithDefaults instantiates a new IPAssignmentInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIPAssignmentInputWithDefaults() *IPAssignmentInput {
	this := IPAssignmentInput{}
	return &this
}

// GetAddress returns the Address field value
func (o *IPAssignmentInput) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *IPAssignmentInput) GetAddressOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *IPAssignmentInput) SetAddress(v string) {
	o.Address = v
}

// GetManageable returns the Manageable field value if set, zero value otherwise.
func (o *IPAssignmentInput) GetManageable() bool {
	if o == nil || o.Manageable == nil {
		var ret bool
		return ret
	}
	return *o.Manageable
}

// GetManageableOk returns a tuple with the Manageable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPAssignmentInput) GetManageableOk() (*bool, bool) {
	if o == nil || o.Manageable == nil {
		return nil, false
	}
	return o.Manageable, true
}

// HasManageable returns a boolean if a field has been set.
func (o *IPAssignmentInput) HasManageable() bool {
	if o != nil && o.Manageable != nil {
		return true
	}

	return false
}

// SetManageable gets a reference to the given bool and assigns it to the Manageable field.
func (o *IPAssignmentInput) SetManageable(v bool) {
	o.Manageable = &v
}

// GetCustomdata returns the Customdata field value if set, zero value otherwise.
func (o *IPAssignmentInput) GetCustomdata() map[string]interface{} {
	if o == nil || o.Customdata == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Customdata
}

// GetCustomdataOk returns a tuple with the Customdata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPAssignmentInput) GetCustomdataOk() (*map[string]interface{}, bool) {
	if o == nil || o.Customdata == nil {
		return nil, false
	}
	return o.Customdata, true
}

// HasCustomdata returns a boolean if a field has been set.
func (o *IPAssignmentInput) HasCustomdata() bool {
	if o != nil && o.Customdata != nil {
		return true
	}

	return false
}

// SetCustomdata gets a reference to the given map[string]interface{} and assigns it to the Customdata field.
func (o *IPAssignmentInput) SetCustomdata(v map[string]interface{}) {
	o.Customdata = &v
}

func (o IPAssignmentInput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["address"] = o.Address
	}
	if o.Manageable != nil {
		toSerialize["manageable"] = o.Manageable
	}
	if o.Customdata != nil {
		toSerialize["customdata"] = o.Customdata
	}
	return json.Marshal(toSerialize)
}

type NullableIPAssignmentInput struct {
	value *IPAssignmentInput
	isSet bool
}

func (v NullableIPAssignmentInput) Get() *IPAssignmentInput {
	return v.value
}

func (v *NullableIPAssignmentInput) Set(val *IPAssignmentInput) {
	v.value = val
	v.isSet = true
}

func (v NullableIPAssignmentInput) IsSet() bool {
	return v.isSet
}

func (v *NullableIPAssignmentInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIPAssignmentInput(val *IPAssignmentInput) *NullableIPAssignmentInput {
	return &NullableIPAssignmentInput{value: val, isSet: true}
}

func (v NullableIPAssignmentInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIPAssignmentInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


