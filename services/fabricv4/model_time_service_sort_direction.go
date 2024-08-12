/*
Equinix Fabric API v4

Contact: api-support@equinix.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fabricv4

import (
	"encoding/json"
	"fmt"
)

// TimeServiceSortDirection Sorting direction
type TimeServiceSortDirection string

// List of TimeServiceSortDirection
const (
	TIMESERVICESORTDIRECTION_DESC TimeServiceSortDirection = "DESC"
	TIMESERVICESORTDIRECTION_ASC  TimeServiceSortDirection = "ASC"
)

// All allowed values of TimeServiceSortDirection enum
var AllowedTimeServiceSortDirectionEnumValues = []TimeServiceSortDirection{
	"DESC",
	"ASC",
}

func (v *TimeServiceSortDirection) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TimeServiceSortDirection(value)
	for _, existing := range AllowedTimeServiceSortDirectionEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TimeServiceSortDirection", value)
}

// NewTimeServiceSortDirectionFromValue returns a pointer to a valid TimeServiceSortDirection
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTimeServiceSortDirectionFromValue(v string) (*TimeServiceSortDirection, error) {
	ev := TimeServiceSortDirection(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TimeServiceSortDirection: valid values are %v", v, AllowedTimeServiceSortDirectionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TimeServiceSortDirection) IsValid() bool {
	for _, existing := range AllowedTimeServiceSortDirectionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TimeServiceSortDirection value
func (v TimeServiceSortDirection) Ptr() *TimeServiceSortDirection {
	return &v
}

type NullableTimeServiceSortDirection struct {
	value *TimeServiceSortDirection
	isSet bool
}

func (v NullableTimeServiceSortDirection) Get() *TimeServiceSortDirection {
	return v.value
}

func (v *NullableTimeServiceSortDirection) Set(val *TimeServiceSortDirection) {
	v.value = val
	v.isSet = true
}

func (v NullableTimeServiceSortDirection) IsSet() bool {
	return v.isSet
}

func (v *NullableTimeServiceSortDirection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTimeServiceSortDirection(val *TimeServiceSortDirection) *NullableTimeServiceSortDirection {
	return &NullableTimeServiceSortDirection{value: val, isSet: true}
}

func (v NullableTimeServiceSortDirection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTimeServiceSortDirection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
