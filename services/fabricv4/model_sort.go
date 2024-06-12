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

// Sort Key or set of keys that organizes the search payload by property (such as createdDate or metroCode) or by direction. Ascending (ASC) is the default value. The \"?\" prefix indicates descending (DESC) order.
type Sort string

// List of sort
const (
	SORT_BANDWIDTH_UTILIZATION Sort = "-bandwidthUtilization"
)

// All allowed values of Sort enum
var AllowedSortEnumValues = []Sort{
	"-bandwidthUtilization",
}

func (v *Sort) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := Sort(value)
	for _, existing := range AllowedSortEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid Sort", value)
}

// NewSortFromValue returns a pointer to a valid Sort
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSortFromValue(v string) (*Sort, error) {
	ev := Sort(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for Sort: valid values are %v", v, AllowedSortEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v Sort) IsValid() bool {
	for _, existing := range AllowedSortEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to sort value
func (v Sort) Ptr() *Sort {
	return &v
}

type NullableSort struct {
	value *Sort
	isSet bool
}

func (v NullableSort) Get() *Sort {
	return v.value
}

func (v *NullableSort) Set(val *Sort) {
	v.value = val
	v.isSet = true
}

func (v NullableSort) IsSet() bool {
	return v.isSet
}

func (v *NullableSort) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSort(val *Sort) *NullableSort {
	return &NullableSort{value: val, isSet: true}
}

func (v NullableSort) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSort) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
