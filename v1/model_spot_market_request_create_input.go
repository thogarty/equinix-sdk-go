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
	"time"
)

// SpotMarketRequestCreateInput struct for SpotMarketRequestCreateInput
type SpotMarketRequestCreateInput struct {
	InstanceAttributes *SpotMarketRequestCreateInputInstanceAttributes `json:"instance_attributes,omitempty"`
	DevicesMin *int32 `json:"devices_min,omitempty"`
	DevicesMax *int32 `json:"devices_max,omitempty"`
	MaxBidPrice *float32 `json:"max_bid_price,omitempty"`
	EndAt *time.Time `json:"end_at,omitempty"`
	Facilities *[]string `json:"facilities,omitempty"`
}

// NewSpotMarketRequestCreateInput instantiates a new SpotMarketRequestCreateInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSpotMarketRequestCreateInput() *SpotMarketRequestCreateInput {
	this := SpotMarketRequestCreateInput{}
	return &this
}

// NewSpotMarketRequestCreateInputWithDefaults instantiates a new SpotMarketRequestCreateInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSpotMarketRequestCreateInputWithDefaults() *SpotMarketRequestCreateInput {
	this := SpotMarketRequestCreateInput{}
	return &this
}

// GetInstanceAttributes returns the InstanceAttributes field value if set, zero value otherwise.
func (o *SpotMarketRequestCreateInput) GetInstanceAttributes() SpotMarketRequestCreateInputInstanceAttributes {
	if o == nil || o.InstanceAttributes == nil {
		var ret SpotMarketRequestCreateInputInstanceAttributes
		return ret
	}
	return *o.InstanceAttributes
}

// GetInstanceAttributesOk returns a tuple with the InstanceAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotMarketRequestCreateInput) GetInstanceAttributesOk() (*SpotMarketRequestCreateInputInstanceAttributes, bool) {
	if o == nil || o.InstanceAttributes == nil {
		return nil, false
	}
	return o.InstanceAttributes, true
}

// HasInstanceAttributes returns a boolean if a field has been set.
func (o *SpotMarketRequestCreateInput) HasInstanceAttributes() bool {
	if o != nil && o.InstanceAttributes != nil {
		return true
	}

	return false
}

// SetInstanceAttributes gets a reference to the given SpotMarketRequestCreateInputInstanceAttributes and assigns it to the InstanceAttributes field.
func (o *SpotMarketRequestCreateInput) SetInstanceAttributes(v SpotMarketRequestCreateInputInstanceAttributes) {
	o.InstanceAttributes = &v
}

// GetDevicesMin returns the DevicesMin field value if set, zero value otherwise.
func (o *SpotMarketRequestCreateInput) GetDevicesMin() int32 {
	if o == nil || o.DevicesMin == nil {
		var ret int32
		return ret
	}
	return *o.DevicesMin
}

// GetDevicesMinOk returns a tuple with the DevicesMin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotMarketRequestCreateInput) GetDevicesMinOk() (*int32, bool) {
	if o == nil || o.DevicesMin == nil {
		return nil, false
	}
	return o.DevicesMin, true
}

// HasDevicesMin returns a boolean if a field has been set.
func (o *SpotMarketRequestCreateInput) HasDevicesMin() bool {
	if o != nil && o.DevicesMin != nil {
		return true
	}

	return false
}

// SetDevicesMin gets a reference to the given int32 and assigns it to the DevicesMin field.
func (o *SpotMarketRequestCreateInput) SetDevicesMin(v int32) {
	o.DevicesMin = &v
}

// GetDevicesMax returns the DevicesMax field value if set, zero value otherwise.
func (o *SpotMarketRequestCreateInput) GetDevicesMax() int32 {
	if o == nil || o.DevicesMax == nil {
		var ret int32
		return ret
	}
	return *o.DevicesMax
}

// GetDevicesMaxOk returns a tuple with the DevicesMax field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotMarketRequestCreateInput) GetDevicesMaxOk() (*int32, bool) {
	if o == nil || o.DevicesMax == nil {
		return nil, false
	}
	return o.DevicesMax, true
}

// HasDevicesMax returns a boolean if a field has been set.
func (o *SpotMarketRequestCreateInput) HasDevicesMax() bool {
	if o != nil && o.DevicesMax != nil {
		return true
	}

	return false
}

// SetDevicesMax gets a reference to the given int32 and assigns it to the DevicesMax field.
func (o *SpotMarketRequestCreateInput) SetDevicesMax(v int32) {
	o.DevicesMax = &v
}

// GetMaxBidPrice returns the MaxBidPrice field value if set, zero value otherwise.
func (o *SpotMarketRequestCreateInput) GetMaxBidPrice() float32 {
	if o == nil || o.MaxBidPrice == nil {
		var ret float32
		return ret
	}
	return *o.MaxBidPrice
}

// GetMaxBidPriceOk returns a tuple with the MaxBidPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotMarketRequestCreateInput) GetMaxBidPriceOk() (*float32, bool) {
	if o == nil || o.MaxBidPrice == nil {
		return nil, false
	}
	return o.MaxBidPrice, true
}

// HasMaxBidPrice returns a boolean if a field has been set.
func (o *SpotMarketRequestCreateInput) HasMaxBidPrice() bool {
	if o != nil && o.MaxBidPrice != nil {
		return true
	}

	return false
}

// SetMaxBidPrice gets a reference to the given float32 and assigns it to the MaxBidPrice field.
func (o *SpotMarketRequestCreateInput) SetMaxBidPrice(v float32) {
	o.MaxBidPrice = &v
}

// GetEndAt returns the EndAt field value if set, zero value otherwise.
func (o *SpotMarketRequestCreateInput) GetEndAt() time.Time {
	if o == nil || o.EndAt == nil {
		var ret time.Time
		return ret
	}
	return *o.EndAt
}

// GetEndAtOk returns a tuple with the EndAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotMarketRequestCreateInput) GetEndAtOk() (*time.Time, bool) {
	if o == nil || o.EndAt == nil {
		return nil, false
	}
	return o.EndAt, true
}

// HasEndAt returns a boolean if a field has been set.
func (o *SpotMarketRequestCreateInput) HasEndAt() bool {
	if o != nil && o.EndAt != nil {
		return true
	}

	return false
}

// SetEndAt gets a reference to the given time.Time and assigns it to the EndAt field.
func (o *SpotMarketRequestCreateInput) SetEndAt(v time.Time) {
	o.EndAt = &v
}

// GetFacilities returns the Facilities field value if set, zero value otherwise.
func (o *SpotMarketRequestCreateInput) GetFacilities() []string {
	if o == nil || o.Facilities == nil {
		var ret []string
		return ret
	}
	return *o.Facilities
}

// GetFacilitiesOk returns a tuple with the Facilities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotMarketRequestCreateInput) GetFacilitiesOk() (*[]string, bool) {
	if o == nil || o.Facilities == nil {
		return nil, false
	}
	return o.Facilities, true
}

// HasFacilities returns a boolean if a field has been set.
func (o *SpotMarketRequestCreateInput) HasFacilities() bool {
	if o != nil && o.Facilities != nil {
		return true
	}

	return false
}

// SetFacilities gets a reference to the given []string and assigns it to the Facilities field.
func (o *SpotMarketRequestCreateInput) SetFacilities(v []string) {
	o.Facilities = &v
}

func (o SpotMarketRequestCreateInput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.InstanceAttributes != nil {
		toSerialize["instance_attributes"] = o.InstanceAttributes
	}
	if o.DevicesMin != nil {
		toSerialize["devices_min"] = o.DevicesMin
	}
	if o.DevicesMax != nil {
		toSerialize["devices_max"] = o.DevicesMax
	}
	if o.MaxBidPrice != nil {
		toSerialize["max_bid_price"] = o.MaxBidPrice
	}
	if o.EndAt != nil {
		toSerialize["end_at"] = o.EndAt
	}
	if o.Facilities != nil {
		toSerialize["facilities"] = o.Facilities
	}
	return json.Marshal(toSerialize)
}

type NullableSpotMarketRequestCreateInput struct {
	value *SpotMarketRequestCreateInput
	isSet bool
}

func (v NullableSpotMarketRequestCreateInput) Get() *SpotMarketRequestCreateInput {
	return v.value
}

func (v *NullableSpotMarketRequestCreateInput) Set(val *SpotMarketRequestCreateInput) {
	v.value = val
	v.isSet = true
}

func (v NullableSpotMarketRequestCreateInput) IsSet() bool {
	return v.isSet
}

func (v *NullableSpotMarketRequestCreateInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpotMarketRequestCreateInput(val *SpotMarketRequestCreateInput) *NullableSpotMarketRequestCreateInput {
	return &NullableSpotMarketRequestCreateInput{value: val, isSet: true}
}

func (v NullableSpotMarketRequestCreateInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpotMarketRequestCreateInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


