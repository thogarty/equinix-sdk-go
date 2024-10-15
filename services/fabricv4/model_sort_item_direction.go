/*
Equinix Fabric API v4

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fabricv4

import (
	"encoding/json"
	"fmt"
)

// SortItemDirection Sorting direction
type SortItemDirection string

// List of SortItem_direction
const (
	SORTITEMDIRECTION_DESC SortItemDirection = "DESC"
	SORTITEMDIRECTION_ASC  SortItemDirection = "ASC"
)

// All allowed values of SortItemDirection enum
var AllowedSortItemDirectionEnumValues = []SortItemDirection{
	"DESC",
	"ASC",
}

func (v *SortItemDirection) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SortItemDirection(value)
	for _, existing := range AllowedSortItemDirectionEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SortItemDirection", value)
}

// NewSortItemDirectionFromValue returns a pointer to a valid SortItemDirection
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSortItemDirectionFromValue(v string) (*SortItemDirection, error) {
	ev := SortItemDirection(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SortItemDirection: valid values are %v", v, AllowedSortItemDirectionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SortItemDirection) IsValid() bool {
	for _, existing := range AllowedSortItemDirectionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SortItem_direction value
func (v SortItemDirection) Ptr() *SortItemDirection {
	return &v
}

type NullableSortItemDirection struct {
	value *SortItemDirection
	isSet bool
}

func (v NullableSortItemDirection) Get() *SortItemDirection {
	return v.value
}

func (v *NullableSortItemDirection) Set(val *SortItemDirection) {
	v.value = val
	v.isSet = true
}

func (v NullableSortItemDirection) IsSet() bool {
	return v.isSet
}

func (v *NullableSortItemDirection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSortItemDirection(val *SortItemDirection) *NullableSortItemDirection {
	return &NullableSortItemDirection{value: val, isSet: true}
}

func (v NullableSortItemDirection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSortItemDirection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
