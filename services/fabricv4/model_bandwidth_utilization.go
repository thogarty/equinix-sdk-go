/*
Equinix Fabric API v4

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fabricv4

import (
	"encoding/json"
)

// checks if the BandwidthUtilization type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BandwidthUtilization{}

// BandwidthUtilization Bandwidth utilization statistics (octet counters-based)
type BandwidthUtilization struct {
	Unit *BandwidthUtilizationUnit `json:"unit,omitempty"`
	// An interval formatted value, indicating the time-interval the metric objects within the response represent
	MetricInterval       *string    `json:"metricInterval,omitempty"`
	Inbound              *Direction `json:"inbound,omitempty"`
	Outbound             *Direction `json:"outbound,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BandwidthUtilization BandwidthUtilization

// NewBandwidthUtilization instantiates a new BandwidthUtilization object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBandwidthUtilization() *BandwidthUtilization {
	this := BandwidthUtilization{}
	return &this
}

// NewBandwidthUtilizationWithDefaults instantiates a new BandwidthUtilization object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBandwidthUtilizationWithDefaults() *BandwidthUtilization {
	this := BandwidthUtilization{}
	return &this
}

// GetUnit returns the Unit field value if set, zero value otherwise.
func (o *BandwidthUtilization) GetUnit() BandwidthUtilizationUnit {
	if o == nil || IsNil(o.Unit) {
		var ret BandwidthUtilizationUnit
		return ret
	}
	return *o.Unit
}

// GetUnitOk returns a tuple with the Unit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BandwidthUtilization) GetUnitOk() (*BandwidthUtilizationUnit, bool) {
	if o == nil || IsNil(o.Unit) {
		return nil, false
	}
	return o.Unit, true
}

// HasUnit returns a boolean if a field has been set.
func (o *BandwidthUtilization) HasUnit() bool {
	if o != nil && !IsNil(o.Unit) {
		return true
	}

	return false
}

// SetUnit gets a reference to the given BandwidthUtilizationUnit and assigns it to the Unit field.
func (o *BandwidthUtilization) SetUnit(v BandwidthUtilizationUnit) {
	o.Unit = &v
}

// GetMetricInterval returns the MetricInterval field value if set, zero value otherwise.
func (o *BandwidthUtilization) GetMetricInterval() string {
	if o == nil || IsNil(o.MetricInterval) {
		var ret string
		return ret
	}
	return *o.MetricInterval
}

// GetMetricIntervalOk returns a tuple with the MetricInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BandwidthUtilization) GetMetricIntervalOk() (*string, bool) {
	if o == nil || IsNil(o.MetricInterval) {
		return nil, false
	}
	return o.MetricInterval, true
}

// HasMetricInterval returns a boolean if a field has been set.
func (o *BandwidthUtilization) HasMetricInterval() bool {
	if o != nil && !IsNil(o.MetricInterval) {
		return true
	}

	return false
}

// SetMetricInterval gets a reference to the given string and assigns it to the MetricInterval field.
func (o *BandwidthUtilization) SetMetricInterval(v string) {
	o.MetricInterval = &v
}

// GetInbound returns the Inbound field value if set, zero value otherwise.
func (o *BandwidthUtilization) GetInbound() Direction {
	if o == nil || IsNil(o.Inbound) {
		var ret Direction
		return ret
	}
	return *o.Inbound
}

// GetInboundOk returns a tuple with the Inbound field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BandwidthUtilization) GetInboundOk() (*Direction, bool) {
	if o == nil || IsNil(o.Inbound) {
		return nil, false
	}
	return o.Inbound, true
}

// HasInbound returns a boolean if a field has been set.
func (o *BandwidthUtilization) HasInbound() bool {
	if o != nil && !IsNil(o.Inbound) {
		return true
	}

	return false
}

// SetInbound gets a reference to the given Direction and assigns it to the Inbound field.
func (o *BandwidthUtilization) SetInbound(v Direction) {
	o.Inbound = &v
}

// GetOutbound returns the Outbound field value if set, zero value otherwise.
func (o *BandwidthUtilization) GetOutbound() Direction {
	if o == nil || IsNil(o.Outbound) {
		var ret Direction
		return ret
	}
	return *o.Outbound
}

// GetOutboundOk returns a tuple with the Outbound field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BandwidthUtilization) GetOutboundOk() (*Direction, bool) {
	if o == nil || IsNil(o.Outbound) {
		return nil, false
	}
	return o.Outbound, true
}

// HasOutbound returns a boolean if a field has been set.
func (o *BandwidthUtilization) HasOutbound() bool {
	if o != nil && !IsNil(o.Outbound) {
		return true
	}

	return false
}

// SetOutbound gets a reference to the given Direction and assigns it to the Outbound field.
func (o *BandwidthUtilization) SetOutbound(v Direction) {
	o.Outbound = &v
}

func (o BandwidthUtilization) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BandwidthUtilization) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Unit) {
		toSerialize["unit"] = o.Unit
	}
	if !IsNil(o.MetricInterval) {
		toSerialize["metricInterval"] = o.MetricInterval
	}
	if !IsNil(o.Inbound) {
		toSerialize["inbound"] = o.Inbound
	}
	if !IsNil(o.Outbound) {
		toSerialize["outbound"] = o.Outbound
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BandwidthUtilization) UnmarshalJSON(data []byte) (err error) {
	varBandwidthUtilization := _BandwidthUtilization{}

	err = json.Unmarshal(data, &varBandwidthUtilization)

	if err != nil {
		return err
	}

	*o = BandwidthUtilization(varBandwidthUtilization)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "unit")
		delete(additionalProperties, "metricInterval")
		delete(additionalProperties, "inbound")
		delete(additionalProperties, "outbound")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBandwidthUtilization struct {
	value *BandwidthUtilization
	isSet bool
}

func (v NullableBandwidthUtilization) Get() *BandwidthUtilization {
	return v.value
}

func (v *NullableBandwidthUtilization) Set(val *BandwidthUtilization) {
	v.value = val
	v.isSet = true
}

func (v NullableBandwidthUtilization) IsSet() bool {
	return v.isSet
}

func (v *NullableBandwidthUtilization) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBandwidthUtilization(val *BandwidthUtilization) *NullableBandwidthUtilization {
	return &NullableBandwidthUtilization{value: val, isSet: true}
}

func (v NullableBandwidthUtilization) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBandwidthUtilization) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
