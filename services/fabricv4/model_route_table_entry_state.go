/*
Equinix Fabric API v4

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fabricv4

import (
	"encoding/json"
	"fmt"
)

// RouteTableEntryState Route table entry state
type RouteTableEntryState string

// List of RouteTableEntryState
const (
	ROUTETABLEENTRYSTATE_ACTIVE RouteTableEntryState = "ACTIVE"
)

// All allowed values of RouteTableEntryState enum
var AllowedRouteTableEntryStateEnumValues = []RouteTableEntryState{
	"ACTIVE",
}

func (v *RouteTableEntryState) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RouteTableEntryState(value)
	for _, existing := range AllowedRouteTableEntryStateEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RouteTableEntryState", value)
}

// NewRouteTableEntryStateFromValue returns a pointer to a valid RouteTableEntryState
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRouteTableEntryStateFromValue(v string) (*RouteTableEntryState, error) {
	ev := RouteTableEntryState(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RouteTableEntryState: valid values are %v", v, AllowedRouteTableEntryStateEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RouteTableEntryState) IsValid() bool {
	for _, existing := range AllowedRouteTableEntryStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RouteTableEntryState value
func (v RouteTableEntryState) Ptr() *RouteTableEntryState {
	return &v
}

type NullableRouteTableEntryState struct {
	value *RouteTableEntryState
	isSet bool
}

func (v NullableRouteTableEntryState) Get() *RouteTableEntryState {
	return v.value
}

func (v *NullableRouteTableEntryState) Set(val *RouteTableEntryState) {
	v.value = val
	v.isSet = true
}

func (v NullableRouteTableEntryState) IsSet() bool {
	return v.isSet
}

func (v *NullableRouteTableEntryState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRouteTableEntryState(val *RouteTableEntryState) *NullableRouteTableEntryState {
	return &NullableRouteTableEntryState{value: val, isSet: true}
}

func (v NullableRouteTableEntryState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRouteTableEntryState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
