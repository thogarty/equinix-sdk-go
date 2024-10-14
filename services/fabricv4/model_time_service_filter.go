/*
Equinix Fabric API v4

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fabricv4

import (
	"encoding/json"
	"fmt"
)

// TimeServiceFilter struct for TimeServiceFilter
type TimeServiceFilter struct {
	TimeServiceOrFilter         *TimeServiceOrFilter
	TimeServiceSimpleExpression *TimeServiceSimpleExpression
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *TimeServiceFilter) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into TimeServiceOrFilter
	err = json.Unmarshal(data, &dst.TimeServiceOrFilter)
	if err == nil {
		jsonTimeServiceOrFilter, _ := json.Marshal(dst.TimeServiceOrFilter)
		if string(jsonTimeServiceOrFilter) == "{}" { // empty struct
			dst.TimeServiceOrFilter = nil
		} else {
			return nil // data stored in dst.TimeServiceOrFilter, return on the first match
		}
	} else {
		dst.TimeServiceOrFilter = nil
	}

	// try to unmarshal JSON data into TimeServiceSimpleExpression
	err = json.Unmarshal(data, &dst.TimeServiceSimpleExpression)
	if err == nil {
		jsonTimeServiceSimpleExpression, _ := json.Marshal(dst.TimeServiceSimpleExpression)
		if string(jsonTimeServiceSimpleExpression) == "{}" { // empty struct
			dst.TimeServiceSimpleExpression = nil
		} else {
			return nil // data stored in dst.TimeServiceSimpleExpression, return on the first match
		}
	} else {
		dst.TimeServiceSimpleExpression = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(TimeServiceFilter)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *TimeServiceFilter) MarshalJSON() ([]byte, error) {
	if src.TimeServiceOrFilter != nil {
		return json.Marshal(&src.TimeServiceOrFilter)
	}

	if src.TimeServiceSimpleExpression != nil {
		return json.Marshal(&src.TimeServiceSimpleExpression)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableTimeServiceFilter struct {
	value *TimeServiceFilter
	isSet bool
}

func (v NullableTimeServiceFilter) Get() *TimeServiceFilter {
	return v.value
}

func (v *NullableTimeServiceFilter) Set(val *TimeServiceFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableTimeServiceFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableTimeServiceFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTimeServiceFilter(val *TimeServiceFilter) *NullableTimeServiceFilter {
	return &NullableTimeServiceFilter{value: val, isSet: true}
}

func (v NullableTimeServiceFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTimeServiceFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
