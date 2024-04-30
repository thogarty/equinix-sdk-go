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

// checks if the OrderSignatureDelegateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderSignatureDelegateRequest{}

// OrderSignatureDelegateRequest struct for OrderSignatureDelegateRequest
type OrderSignatureDelegateRequest struct {
	FirstName *string `json:"firstName,omitempty"`
	LastName  *string `json:"lastName,omitempty"`
	// Email address that the signature request should be sent to in case of DELEGATE signature
	Email                string `json:"email"`
	AdditionalProperties map[string]interface{}
}

type _OrderSignatureDelegateRequest OrderSignatureDelegateRequest

// NewOrderSignatureDelegateRequest instantiates a new OrderSignatureDelegateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderSignatureDelegateRequest(email string) *OrderSignatureDelegateRequest {
	this := OrderSignatureDelegateRequest{}
	this.Email = email
	return &this
}

// NewOrderSignatureDelegateRequestWithDefaults instantiates a new OrderSignatureDelegateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderSignatureDelegateRequestWithDefaults() *OrderSignatureDelegateRequest {
	this := OrderSignatureDelegateRequest{}
	return &this
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *OrderSignatureDelegateRequest) GetFirstName() string {
	if o == nil || IsNil(o.FirstName) {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderSignatureDelegateRequest) GetFirstNameOk() (*string, bool) {
	if o == nil || IsNil(o.FirstName) {
		return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *OrderSignatureDelegateRequest) HasFirstName() bool {
	if o != nil && !IsNil(o.FirstName) {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *OrderSignatureDelegateRequest) SetFirstName(v string) {
	o.FirstName = &v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *OrderSignatureDelegateRequest) GetLastName() string {
	if o == nil || IsNil(o.LastName) {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderSignatureDelegateRequest) GetLastNameOk() (*string, bool) {
	if o == nil || IsNil(o.LastName) {
		return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *OrderSignatureDelegateRequest) HasLastName() bool {
	if o != nil && !IsNil(o.LastName) {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *OrderSignatureDelegateRequest) SetLastName(v string) {
	o.LastName = &v
}

// GetEmail returns the Email field value
func (o *OrderSignatureDelegateRequest) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *OrderSignatureDelegateRequest) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *OrderSignatureDelegateRequest) SetEmail(v string) {
	o.Email = v
}

func (o OrderSignatureDelegateRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderSignatureDelegateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FirstName) {
		toSerialize["firstName"] = o.FirstName
	}
	if !IsNil(o.LastName) {
		toSerialize["lastName"] = o.LastName
	}
	toSerialize["email"] = o.Email

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OrderSignatureDelegateRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"email",
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

	varOrderSignatureDelegateRequest := _OrderSignatureDelegateRequest{}

	err = json.Unmarshal(data, &varOrderSignatureDelegateRequest)

	if err != nil {
		return err
	}

	*o = OrderSignatureDelegateRequest(varOrderSignatureDelegateRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "firstName")
		delete(additionalProperties, "lastName")
		delete(additionalProperties, "email")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOrderSignatureDelegateRequest struct {
	value *OrderSignatureDelegateRequest
	isSet bool
}

func (v NullableOrderSignatureDelegateRequest) Get() *OrderSignatureDelegateRequest {
	return v.value
}

func (v *NullableOrderSignatureDelegateRequest) Set(val *OrderSignatureDelegateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderSignatureDelegateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderSignatureDelegateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderSignatureDelegateRequest(val *OrderSignatureDelegateRequest) *NullableOrderSignatureDelegateRequest {
	return &NullableOrderSignatureDelegateRequest{value: val, isSet: true}
}

func (v NullableOrderSignatureDelegateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderSignatureDelegateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
