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

// SpotPricesDatapoints struct for SpotPricesDatapoints
type SpotPricesDatapoints struct {
	Datapoints *[][]float32 `json:"datapoints,omitempty"`
}

// NewSpotPricesDatapoints instantiates a new SpotPricesDatapoints object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSpotPricesDatapoints() *SpotPricesDatapoints {
	this := SpotPricesDatapoints{}
	return &this
}

// NewSpotPricesDatapointsWithDefaults instantiates a new SpotPricesDatapoints object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSpotPricesDatapointsWithDefaults() *SpotPricesDatapoints {
	this := SpotPricesDatapoints{}
	return &this
}

// GetDatapoints returns the Datapoints field value if set, zero value otherwise.
func (o *SpotPricesDatapoints) GetDatapoints() [][]float32 {
	if o == nil || o.Datapoints == nil {
		var ret [][]float32
		return ret
	}
	return *o.Datapoints
}

// GetDatapointsOk returns a tuple with the Datapoints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotPricesDatapoints) GetDatapointsOk() (*[][]float32, bool) {
	if o == nil || o.Datapoints == nil {
		return nil, false
	}
	return o.Datapoints, true
}

// HasDatapoints returns a boolean if a field has been set.
func (o *SpotPricesDatapoints) HasDatapoints() bool {
	if o != nil && o.Datapoints != nil {
		return true
	}

	return false
}

// SetDatapoints gets a reference to the given [][]float32 and assigns it to the Datapoints field.
func (o *SpotPricesDatapoints) SetDatapoints(v [][]float32) {
	o.Datapoints = &v
}

func (o SpotPricesDatapoints) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Datapoints != nil {
		toSerialize["datapoints"] = o.Datapoints
	}
	return json.Marshal(toSerialize)
}

type NullableSpotPricesDatapoints struct {
	value *SpotPricesDatapoints
	isSet bool
}

func (v NullableSpotPricesDatapoints) Get() *SpotPricesDatapoints {
	return v.value
}

func (v *NullableSpotPricesDatapoints) Set(val *SpotPricesDatapoints) {
	v.value = val
	v.isSet = true
}

func (v NullableSpotPricesDatapoints) IsSet() bool {
	return v.isSet
}

func (v *NullableSpotPricesDatapoints) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpotPricesDatapoints(val *SpotPricesDatapoints) *NullableSpotPricesDatapoints {
	return &NullableSpotPricesDatapoints{value: val, isSet: true}
}

func (v NullableSpotPricesDatapoints) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpotPricesDatapoints) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


