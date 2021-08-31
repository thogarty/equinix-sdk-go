/*
Metal API

This is the API for Equinix Metal. The API allows you to programmatically interact with all of your Equinix Metal resources, including devices, networks, addresses, organizations, projects, and your user account.  The official API docs are hosted at <https://metal.equinix.com/developers/api>. 

API version: 1.0.0
Contact: support@equinixmetal.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1

import (
	"encoding/json"
)

// SelfServiceReservationItemRequest struct for SelfServiceReservationItemRequest
type SelfServiceReservationItemRequest struct {
	MetroId *string `json:"metro_id,omitempty"`
	PlanId *string `json:"plan_id,omitempty"`
	Quantity *int32 `json:"quantity,omitempty"`
	Amount *float32 `json:"amount,omitempty"`
	Term *string `json:"term,omitempty"`
}

// NewSelfServiceReservationItemRequest instantiates a new SelfServiceReservationItemRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSelfServiceReservationItemRequest() *SelfServiceReservationItemRequest {
	this := SelfServiceReservationItemRequest{}
	return &this
}

// NewSelfServiceReservationItemRequestWithDefaults instantiates a new SelfServiceReservationItemRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSelfServiceReservationItemRequestWithDefaults() *SelfServiceReservationItemRequest {
	this := SelfServiceReservationItemRequest{}
	return &this
}

// GetMetroId returns the MetroId field value if set, zero value otherwise.
func (o *SelfServiceReservationItemRequest) GetMetroId() string {
	if o == nil || o.MetroId == nil {
		var ret string
		return ret
	}
	return *o.MetroId
}

// GetMetroIdOk returns a tuple with the MetroId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelfServiceReservationItemRequest) GetMetroIdOk() (*string, bool) {
	if o == nil || o.MetroId == nil {
		return nil, false
	}
	return o.MetroId, true
}

// HasMetroId returns a boolean if a field has been set.
func (o *SelfServiceReservationItemRequest) HasMetroId() bool {
	if o != nil && o.MetroId != nil {
		return true
	}

	return false
}

// SetMetroId gets a reference to the given string and assigns it to the MetroId field.
func (o *SelfServiceReservationItemRequest) SetMetroId(v string) {
	o.MetroId = &v
}

// GetPlanId returns the PlanId field value if set, zero value otherwise.
func (o *SelfServiceReservationItemRequest) GetPlanId() string {
	if o == nil || o.PlanId == nil {
		var ret string
		return ret
	}
	return *o.PlanId
}

// GetPlanIdOk returns a tuple with the PlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelfServiceReservationItemRequest) GetPlanIdOk() (*string, bool) {
	if o == nil || o.PlanId == nil {
		return nil, false
	}
	return o.PlanId, true
}

// HasPlanId returns a boolean if a field has been set.
func (o *SelfServiceReservationItemRequest) HasPlanId() bool {
	if o != nil && o.PlanId != nil {
		return true
	}

	return false
}

// SetPlanId gets a reference to the given string and assigns it to the PlanId field.
func (o *SelfServiceReservationItemRequest) SetPlanId(v string) {
	o.PlanId = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *SelfServiceReservationItemRequest) GetQuantity() int32 {
	if o == nil || o.Quantity == nil {
		var ret int32
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelfServiceReservationItemRequest) GetQuantityOk() (*int32, bool) {
	if o == nil || o.Quantity == nil {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *SelfServiceReservationItemRequest) HasQuantity() bool {
	if o != nil && o.Quantity != nil {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given int32 and assigns it to the Quantity field.
func (o *SelfServiceReservationItemRequest) SetQuantity(v int32) {
	o.Quantity = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *SelfServiceReservationItemRequest) GetAmount() float32 {
	if o == nil || o.Amount == nil {
		var ret float32
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelfServiceReservationItemRequest) GetAmountOk() (*float32, bool) {
	if o == nil || o.Amount == nil {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *SelfServiceReservationItemRequest) HasAmount() bool {
	if o != nil && o.Amount != nil {
		return true
	}

	return false
}

// SetAmount gets a reference to the given float32 and assigns it to the Amount field.
func (o *SelfServiceReservationItemRequest) SetAmount(v float32) {
	o.Amount = &v
}

// GetTerm returns the Term field value if set, zero value otherwise.
func (o *SelfServiceReservationItemRequest) GetTerm() string {
	if o == nil || o.Term == nil {
		var ret string
		return ret
	}
	return *o.Term
}

// GetTermOk returns a tuple with the Term field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelfServiceReservationItemRequest) GetTermOk() (*string, bool) {
	if o == nil || o.Term == nil {
		return nil, false
	}
	return o.Term, true
}

// HasTerm returns a boolean if a field has been set.
func (o *SelfServiceReservationItemRequest) HasTerm() bool {
	if o != nil && o.Term != nil {
		return true
	}

	return false
}

// SetTerm gets a reference to the given string and assigns it to the Term field.
func (o *SelfServiceReservationItemRequest) SetTerm(v string) {
	o.Term = &v
}

func (o SelfServiceReservationItemRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.MetroId != nil {
		toSerialize["metro_id"] = o.MetroId
	}
	if o.PlanId != nil {
		toSerialize["plan_id"] = o.PlanId
	}
	if o.Quantity != nil {
		toSerialize["quantity"] = o.Quantity
	}
	if o.Amount != nil {
		toSerialize["amount"] = o.Amount
	}
	if o.Term != nil {
		toSerialize["term"] = o.Term
	}
	return json.Marshal(toSerialize)
}

type NullableSelfServiceReservationItemRequest struct {
	value *SelfServiceReservationItemRequest
	isSet bool
}

func (v NullableSelfServiceReservationItemRequest) Get() *SelfServiceReservationItemRequest {
	return v.value
}

func (v *NullableSelfServiceReservationItemRequest) Set(val *SelfServiceReservationItemRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSelfServiceReservationItemRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSelfServiceReservationItemRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSelfServiceReservationItemRequest(val *SelfServiceReservationItemRequest) *NullableSelfServiceReservationItemRequest {
	return &NullableSelfServiceReservationItemRequest{value: val, isSet: true}
}

func (v NullableSelfServiceReservationItemRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSelfServiceReservationItemRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


