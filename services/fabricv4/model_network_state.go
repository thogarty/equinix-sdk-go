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

// NetworkState Network status
type NetworkState string

// List of NetworkState
const (
	NETWORKSTATE_ACTIVE   NetworkState = "ACTIVE"
	NETWORKSTATE_INACTIVE NetworkState = "INACTIVE"
	NETWORKSTATE_DELETED  NetworkState = "DELETED"
)

// All allowed values of NetworkState enum
var AllowedNetworkStateEnumValues = []NetworkState{
	"ACTIVE",
	"INACTIVE",
	"DELETED",
}

func (v *NetworkState) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := NetworkState(value)
	for _, existing := range AllowedNetworkStateEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid NetworkState", value)
}

// NewNetworkStateFromValue returns a pointer to a valid NetworkState
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNetworkStateFromValue(v string) (*NetworkState, error) {
	ev := NetworkState(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NetworkState: valid values are %v", v, AllowedNetworkStateEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NetworkState) IsValid() bool {
	for _, existing := range AllowedNetworkStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to NetworkState value
func (v NetworkState) Ptr() *NetworkState {
	return &v
}

type NullableNetworkState struct {
	value *NetworkState
	isSet bool
}

func (v NullableNetworkState) Get() *NetworkState {
	return v.value
}

func (v *NullableNetworkState) Set(val *NetworkState) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkState) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkState(val *NetworkState) *NullableNetworkState {
	return &NullableNetworkState{value: val, isSet: true}
}

func (v NullableNetworkState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
