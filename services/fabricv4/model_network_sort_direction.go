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

// NetworkSortDirection Sorting direction
type NetworkSortDirection string

// List of NetworkSortDirection
const (
	NETWORKSORTDIRECTION_DESC NetworkSortDirection = "DESC"
	NETWORKSORTDIRECTION_ASC  NetworkSortDirection = "ASC"
)

// All allowed values of NetworkSortDirection enum
var AllowedNetworkSortDirectionEnumValues = []NetworkSortDirection{
	"DESC",
	"ASC",
}

func (v *NetworkSortDirection) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := NetworkSortDirection(value)
	for _, existing := range AllowedNetworkSortDirectionEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid NetworkSortDirection", value)
}

// NewNetworkSortDirectionFromValue returns a pointer to a valid NetworkSortDirection
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNetworkSortDirectionFromValue(v string) (*NetworkSortDirection, error) {
	ev := NetworkSortDirection(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NetworkSortDirection: valid values are %v", v, AllowedNetworkSortDirectionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NetworkSortDirection) IsValid() bool {
	for _, existing := range AllowedNetworkSortDirectionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to NetworkSortDirection value
func (v NetworkSortDirection) Ptr() *NetworkSortDirection {
	return &v
}

type NullableNetworkSortDirection struct {
	value *NetworkSortDirection
	isSet bool
}

func (v NullableNetworkSortDirection) Get() *NetworkSortDirection {
	return v.value
}

func (v *NullableNetworkSortDirection) Set(val *NetworkSortDirection) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkSortDirection) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkSortDirection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkSortDirection(val *NetworkSortDirection) *NullableNetworkSortDirection {
	return &NullableNetworkSortDirection{value: val, isSet: true}
}

func (v NullableNetworkSortDirection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkSortDirection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
