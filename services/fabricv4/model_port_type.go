/*
Equinix Fabric API v4

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fabricv4

import (
	"encoding/json"
	"fmt"
)

// PortType Type of Port
type PortType string

// List of PortType
const (
	PORTTYPE_XF_PORT PortType = "XF_PORT"
	PORTTYPE_IX_PORT PortType = "IX_PORT"
)

// All allowed values of PortType enum
var AllowedPortTypeEnumValues = []PortType{
	"XF_PORT",
	"IX_PORT",
}

func (v *PortType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PortType(value)
	for _, existing := range AllowedPortTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PortType", value)
}

// NewPortTypeFromValue returns a pointer to a valid PortType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPortTypeFromValue(v string) (*PortType, error) {
	ev := PortType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PortType: valid values are %v", v, AllowedPortTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PortType) IsValid() bool {
	for _, existing := range AllowedPortTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PortType value
func (v PortType) Ptr() *PortType {
	return &v
}

type NullablePortType struct {
	value *PortType
	isSet bool
}

func (v NullablePortType) Get() *PortType {
	return v.value
}

func (v *NullablePortType) Set(val *PortType) {
	v.value = val
	v.isSet = true
}

func (v NullablePortType) IsSet() bool {
	return v.isSet
}

func (v *NullablePortType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePortType(val *PortType) *NullablePortType {
	return &NullablePortType{value: val, isSet: true}
}

func (v NullablePortType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePortType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
