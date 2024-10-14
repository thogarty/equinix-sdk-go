/*
Equinix Fabric API v4

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fabricv4

import (
	"encoding/json"
)

// checks if the TimeServiceFilters type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TimeServiceFilters{}

// TimeServiceFilters struct for TimeServiceFilters
type TimeServiceFilters struct {
	And                  []TimeServiceFilter `json:"and,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TimeServiceFilters TimeServiceFilters

// NewTimeServiceFilters instantiates a new TimeServiceFilters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTimeServiceFilters() *TimeServiceFilters {
	this := TimeServiceFilters{}
	return &this
}

// NewTimeServiceFiltersWithDefaults instantiates a new TimeServiceFilters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTimeServiceFiltersWithDefaults() *TimeServiceFilters {
	this := TimeServiceFilters{}
	return &this
}

// GetAnd returns the And field value if set, zero value otherwise.
func (o *TimeServiceFilters) GetAnd() []TimeServiceFilter {
	if o == nil || IsNil(o.And) {
		var ret []TimeServiceFilter
		return ret
	}
	return o.And
}

// GetAndOk returns a tuple with the And field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeServiceFilters) GetAndOk() ([]TimeServiceFilter, bool) {
	if o == nil || IsNil(o.And) {
		return nil, false
	}
	return o.And, true
}

// HasAnd returns a boolean if a field has been set.
func (o *TimeServiceFilters) HasAnd() bool {
	if o != nil && !IsNil(o.And) {
		return true
	}

	return false
}

// SetAnd gets a reference to the given []TimeServiceFilter and assigns it to the And field.
func (o *TimeServiceFilters) SetAnd(v []TimeServiceFilter) {
	o.And = v
}

func (o TimeServiceFilters) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TimeServiceFilters) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.And) {
		toSerialize["and"] = o.And
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TimeServiceFilters) UnmarshalJSON(data []byte) (err error) {
	varTimeServiceFilters := _TimeServiceFilters{}

	err = json.Unmarshal(data, &varTimeServiceFilters)

	if err != nil {
		return err
	}

	*o = TimeServiceFilters(varTimeServiceFilters)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "and")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTimeServiceFilters struct {
	value *TimeServiceFilters
	isSet bool
}

func (v NullableTimeServiceFilters) Get() *TimeServiceFilters {
	return v.value
}

func (v *NullableTimeServiceFilters) Set(val *TimeServiceFilters) {
	v.value = val
	v.isSet = true
}

func (v NullableTimeServiceFilters) IsSet() bool {
	return v.isSet
}

func (v *NullableTimeServiceFilters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTimeServiceFilters(val *TimeServiceFilters) *NullableTimeServiceFilters {
	return &NullableTimeServiceFilters{value: val, isSet: true}
}

func (v NullableTimeServiceFilters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTimeServiceFilters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
