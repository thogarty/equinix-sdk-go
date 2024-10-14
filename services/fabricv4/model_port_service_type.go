/*
Equinix Fabric API v4

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fabricv4

import (
	"encoding/json"
	"fmt"
)

// PortServiceType Port service Type
type PortServiceType string

// List of Port_serviceType
const (
	PORTSERVICETYPE_EPL PortServiceType = "EPL"
	PORTSERVICETYPE_MSP PortServiceType = "MSP"
)

// All allowed values of PortServiceType enum
var AllowedPortServiceTypeEnumValues = []PortServiceType{
	"EPL",
	"MSP",
}

func (v *PortServiceType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PortServiceType(value)
	for _, existing := range AllowedPortServiceTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PortServiceType", value)
}

// NewPortServiceTypeFromValue returns a pointer to a valid PortServiceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPortServiceTypeFromValue(v string) (*PortServiceType, error) {
	ev := PortServiceType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PortServiceType: valid values are %v", v, AllowedPortServiceTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PortServiceType) IsValid() bool {
	for _, existing := range AllowedPortServiceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Port_serviceType value
func (v PortServiceType) Ptr() *PortServiceType {
	return &v
}

type NullablePortServiceType struct {
	value *PortServiceType
	isSet bool
}

func (v NullablePortServiceType) Get() *PortServiceType {
	return v.value
}

func (v *NullablePortServiceType) Set(val *PortServiceType) {
	v.value = val
	v.isSet = true
}

func (v NullablePortServiceType) IsSet() bool {
	return v.isSet
}

func (v *NullablePortServiceType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePortServiceType(val *PortServiceType) *NullablePortServiceType {
	return &NullablePortServiceType{value: val, isSet: true}
}

func (v NullablePortServiceType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePortServiceType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
