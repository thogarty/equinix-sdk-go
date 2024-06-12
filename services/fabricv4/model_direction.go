/*
Equinix Fabric API v4

Contact: api-support@equinix.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fabricv4

import (
	"encoding/json"
)

// checks if the Direction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Direction{}

// Direction Directional statistics
type Direction struct {
	// Max bandwidth within request time range, represented in units specified by response \"units\" field
	Max *float32 `json:"max,omitempty"`
	// Mean bandwidth within request time range, represented in units specified by response \"units\" field
	Mean *float32 `json:"mean,omitempty"`
	// Bandwidth utilization statistics for a specified interval.
	Metrics              []Metrics `json:"metrics,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Direction Direction

// NewDirection instantiates a new Direction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDirection() *Direction {
	this := Direction{}
	return &this
}

// NewDirectionWithDefaults instantiates a new Direction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDirectionWithDefaults() *Direction {
	this := Direction{}
	return &this
}

// GetMax returns the Max field value if set, zero value otherwise.
func (o *Direction) GetMax() float32 {
	if o == nil || IsNil(o.Max) {
		var ret float32
		return ret
	}
	return *o.Max
}

// GetMaxOk returns a tuple with the Max field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Direction) GetMaxOk() (*float32, bool) {
	if o == nil || IsNil(o.Max) {
		return nil, false
	}
	return o.Max, true
}

// HasMax returns a boolean if a field has been set.
func (o *Direction) HasMax() bool {
	if o != nil && !IsNil(o.Max) {
		return true
	}

	return false
}

// SetMax gets a reference to the given float32 and assigns it to the Max field.
func (o *Direction) SetMax(v float32) {
	o.Max = &v
}

// GetMean returns the Mean field value if set, zero value otherwise.
func (o *Direction) GetMean() float32 {
	if o == nil || IsNil(o.Mean) {
		var ret float32
		return ret
	}
	return *o.Mean
}

// GetMeanOk returns a tuple with the Mean field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Direction) GetMeanOk() (*float32, bool) {
	if o == nil || IsNil(o.Mean) {
		return nil, false
	}
	return o.Mean, true
}

// HasMean returns a boolean if a field has been set.
func (o *Direction) HasMean() bool {
	if o != nil && !IsNil(o.Mean) {
		return true
	}

	return false
}

// SetMean gets a reference to the given float32 and assigns it to the Mean field.
func (o *Direction) SetMean(v float32) {
	o.Mean = &v
}

// GetMetrics returns the Metrics field value if set, zero value otherwise.
func (o *Direction) GetMetrics() []Metrics {
	if o == nil || IsNil(o.Metrics) {
		var ret []Metrics
		return ret
	}
	return o.Metrics
}

// GetMetricsOk returns a tuple with the Metrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Direction) GetMetricsOk() ([]Metrics, bool) {
	if o == nil || IsNil(o.Metrics) {
		return nil, false
	}
	return o.Metrics, true
}

// HasMetrics returns a boolean if a field has been set.
func (o *Direction) HasMetrics() bool {
	if o != nil && !IsNil(o.Metrics) {
		return true
	}

	return false
}

// SetMetrics gets a reference to the given []Metrics and assigns it to the Metrics field.
func (o *Direction) SetMetrics(v []Metrics) {
	o.Metrics = v
}

func (o Direction) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Direction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Max) {
		toSerialize["max"] = o.Max
	}
	if !IsNil(o.Mean) {
		toSerialize["mean"] = o.Mean
	}
	if !IsNil(o.Metrics) {
		toSerialize["metrics"] = o.Metrics
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Direction) UnmarshalJSON(data []byte) (err error) {
	varDirection := _Direction{}

	err = json.Unmarshal(data, &varDirection)

	if err != nil {
		return err
	}

	*o = Direction(varDirection)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "max")
		delete(additionalProperties, "mean")
		delete(additionalProperties, "metrics")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDirection struct {
	value *Direction
	isSet bool
}

func (v NullableDirection) Get() *Direction {
	return v.value
}

func (v *NullableDirection) Set(val *Direction) {
	v.value = val
	v.isSet = true
}

func (v NullableDirection) IsSet() bool {
	return v.isSet
}

func (v *NullableDirection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDirection(val *Direction) *NullableDirection {
	return &NullableDirection{value: val, isSet: true}
}

func (v NullableDirection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDirection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
