/*
Equinix Fabric API v4

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fabricv4

import (
	"encoding/json"
	"fmt"
)

// SimplifiedTokenNetworkType Type of Network
type SimplifiedTokenNetworkType string

// List of SimplifiedTokenNetwork_type
const (
	SIMPLIFIEDTOKENNETWORKTYPE_EVPLAN SimplifiedTokenNetworkType = "EVPLAN"
	SIMPLIFIEDTOKENNETWORKTYPE_EPLAN  SimplifiedTokenNetworkType = "EPLAN"
	SIMPLIFIEDTOKENNETWORKTYPE_IPWAN  SimplifiedTokenNetworkType = "IPWAN"
)

// All allowed values of SimplifiedTokenNetworkType enum
var AllowedSimplifiedTokenNetworkTypeEnumValues = []SimplifiedTokenNetworkType{
	"EVPLAN",
	"EPLAN",
	"IPWAN",
}

func (v *SimplifiedTokenNetworkType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SimplifiedTokenNetworkType(value)
	for _, existing := range AllowedSimplifiedTokenNetworkTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SimplifiedTokenNetworkType", value)
}

// NewSimplifiedTokenNetworkTypeFromValue returns a pointer to a valid SimplifiedTokenNetworkType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSimplifiedTokenNetworkTypeFromValue(v string) (*SimplifiedTokenNetworkType, error) {
	ev := SimplifiedTokenNetworkType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SimplifiedTokenNetworkType: valid values are %v", v, AllowedSimplifiedTokenNetworkTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SimplifiedTokenNetworkType) IsValid() bool {
	for _, existing := range AllowedSimplifiedTokenNetworkTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SimplifiedTokenNetwork_type value
func (v SimplifiedTokenNetworkType) Ptr() *SimplifiedTokenNetworkType {
	return &v
}

type NullableSimplifiedTokenNetworkType struct {
	value *SimplifiedTokenNetworkType
	isSet bool
}

func (v NullableSimplifiedTokenNetworkType) Get() *SimplifiedTokenNetworkType {
	return v.value
}

func (v *NullableSimplifiedTokenNetworkType) Set(val *SimplifiedTokenNetworkType) {
	v.value = val
	v.isSet = true
}

func (v NullableSimplifiedTokenNetworkType) IsSet() bool {
	return v.isSet
}

func (v *NullableSimplifiedTokenNetworkType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSimplifiedTokenNetworkType(val *SimplifiedTokenNetworkType) *NullableSimplifiedTokenNetworkType {
	return &NullableSimplifiedTokenNetworkType{value: val, isSet: true}
}

func (v NullableSimplifiedTokenNetworkType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSimplifiedTokenNetworkType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
