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

// Userdata struct for Userdata
type Userdata struct {
	Userdata *string `json:"userdata,omitempty"`
}

// NewUserdata instantiates a new Userdata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserdata() *Userdata {
	this := Userdata{}
	return &this
}

// NewUserdataWithDefaults instantiates a new Userdata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserdataWithDefaults() *Userdata {
	this := Userdata{}
	return &this
}

// GetUserdata returns the Userdata field value if set, zero value otherwise.
func (o *Userdata) GetUserdata() string {
	if o == nil || o.Userdata == nil {
		var ret string
		return ret
	}
	return *o.Userdata
}

// GetUserdataOk returns a tuple with the Userdata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Userdata) GetUserdataOk() (*string, bool) {
	if o == nil || o.Userdata == nil {
		return nil, false
	}
	return o.Userdata, true
}

// HasUserdata returns a boolean if a field has been set.
func (o *Userdata) HasUserdata() bool {
	if o != nil && o.Userdata != nil {
		return true
	}

	return false
}

// SetUserdata gets a reference to the given string and assigns it to the Userdata field.
func (o *Userdata) SetUserdata(v string) {
	o.Userdata = &v
}

func (o Userdata) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Userdata != nil {
		toSerialize["userdata"] = o.Userdata
	}
	return json.Marshal(toSerialize)
}

type NullableUserdata struct {
	value *Userdata
	isSet bool
}

func (v NullableUserdata) Get() *Userdata {
	return v.value
}

func (v *NullableUserdata) Set(val *Userdata) {
	v.value = val
	v.isSet = true
}

func (v NullableUserdata) IsSet() bool {
	return v.isSet
}

func (v *NullableUserdata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserdata(val *Userdata) *NullableUserdata {
	return &NullableUserdata{value: val, isSet: true}
}

func (v NullableUserdata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserdata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


