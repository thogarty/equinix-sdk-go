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

// checks if the OrderSignatureRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderSignatureRequest{}

// OrderSignatureRequest struct for OrderSignatureRequest
type OrderSignatureRequest struct {
	Signatory            OrderSignatureRequestSignatory `json:"signatory"`
	Delegate             *OrderSignatureDelegateRequest `json:"delegate,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OrderSignatureRequest OrderSignatureRequest

// NewOrderSignatureRequest instantiates a new OrderSignatureRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderSignatureRequest(signatory OrderSignatureRequestSignatory) *OrderSignatureRequest {
	this := OrderSignatureRequest{}
	this.Signatory = signatory
	return &this
}

// NewOrderSignatureRequestWithDefaults instantiates a new OrderSignatureRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderSignatureRequestWithDefaults() *OrderSignatureRequest {
	this := OrderSignatureRequest{}
	return &this
}

// GetSignatory returns the Signatory field value
func (o *OrderSignatureRequest) GetSignatory() OrderSignatureRequestSignatory {
	if o == nil {
		var ret OrderSignatureRequestSignatory
		return ret
	}

	return o.Signatory
}

// GetSignatoryOk returns a tuple with the Signatory field value
// and a boolean to check if the value has been set.
func (o *OrderSignatureRequest) GetSignatoryOk() (*OrderSignatureRequestSignatory, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Signatory, true
}

// SetSignatory sets field value
func (o *OrderSignatureRequest) SetSignatory(v OrderSignatureRequestSignatory) {
	o.Signatory = v
}

// GetDelegate returns the Delegate field value if set, zero value otherwise.
func (o *OrderSignatureRequest) GetDelegate() OrderSignatureDelegateRequest {
	if o == nil || IsNil(o.Delegate) {
		var ret OrderSignatureDelegateRequest
		return ret
	}
	return *o.Delegate
}

// GetDelegateOk returns a tuple with the Delegate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderSignatureRequest) GetDelegateOk() (*OrderSignatureDelegateRequest, bool) {
	if o == nil || IsNil(o.Delegate) {
		return nil, false
	}
	return o.Delegate, true
}

// HasDelegate returns a boolean if a field has been set.
func (o *OrderSignatureRequest) HasDelegate() bool {
	if o != nil && !IsNil(o.Delegate) {
		return true
	}

	return false
}

// SetDelegate gets a reference to the given OrderSignatureDelegateRequest and assigns it to the Delegate field.
func (o *OrderSignatureRequest) SetDelegate(v OrderSignatureDelegateRequest) {
	o.Delegate = &v
}

func (o OrderSignatureRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderSignatureRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["signatory"] = o.Signatory
	if !IsNil(o.Delegate) {
		toSerialize["delegate"] = o.Delegate
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OrderSignatureRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"signatory",
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

	varOrderSignatureRequest := _OrderSignatureRequest{}

	err = json.Unmarshal(data, &varOrderSignatureRequest)

	if err != nil {
		return err
	}

	*o = OrderSignatureRequest(varOrderSignatureRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "signatory")
		delete(additionalProperties, "delegate")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOrderSignatureRequest struct {
	value *OrderSignatureRequest
	isSet bool
}

func (v NullableOrderSignatureRequest) Get() *OrderSignatureRequest {
	return v.value
}

func (v *NullableOrderSignatureRequest) Set(val *OrderSignatureRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderSignatureRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderSignatureRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderSignatureRequest(val *OrderSignatureRequest) *NullableOrderSignatureRequest {
	return &NullableOrderSignatureRequest{value: val, isSet: true}
}

func (v NullableOrderSignatureRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderSignatureRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
