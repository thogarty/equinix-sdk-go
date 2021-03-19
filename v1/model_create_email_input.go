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

// CreateEmailInput struct for CreateEmailInput
type CreateEmailInput struct {
	Address string `json:"address"`
}

// NewCreateEmailInput instantiates a new CreateEmailInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateEmailInput(address string) *CreateEmailInput {
	this := CreateEmailInput{}
	this.Address = address
	return &this
}

// NewCreateEmailInputWithDefaults instantiates a new CreateEmailInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateEmailInputWithDefaults() *CreateEmailInput {
	this := CreateEmailInput{}
	return &this
}

// GetAddress returns the Address field value
func (o *CreateEmailInput) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *CreateEmailInput) GetAddressOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *CreateEmailInput) SetAddress(v string) {
	o.Address = v
}

func (o CreateEmailInput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["address"] = o.Address
	}
	return json.Marshal(toSerialize)
}

type NullableCreateEmailInput struct {
	value *CreateEmailInput
	isSet bool
}

func (v NullableCreateEmailInput) Get() *CreateEmailInput {
	return v.value
}

func (v *NullableCreateEmailInput) Set(val *CreateEmailInput) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateEmailInput) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateEmailInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateEmailInput(val *CreateEmailInput) *NullableCreateEmailInput {
	return &NullableCreateEmailInput{value: val, isSet: true}
}

func (v NullableCreateEmailInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateEmailInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


