/*
Equinix Fabric API v4

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fabricv4

import (
	"encoding/json"
	"fmt"
)

// VirtualPortServiceType Port service type. The default is managed-service provider (MSP).
type VirtualPortServiceType string

// List of VirtualPortServiceType
const (
	VIRTUALPORTSERVICETYPE_MSP VirtualPortServiceType = "MSP"
	VIRTUALPORTSERVICETYPE_EPL VirtualPortServiceType = "EPL"
)

// All allowed values of VirtualPortServiceType enum
var AllowedVirtualPortServiceTypeEnumValues = []VirtualPortServiceType{
	"MSP",
	"EPL",
}

func (v *VirtualPortServiceType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := VirtualPortServiceType(value)
	for _, existing := range AllowedVirtualPortServiceTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid VirtualPortServiceType", value)
}

// NewVirtualPortServiceTypeFromValue returns a pointer to a valid VirtualPortServiceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewVirtualPortServiceTypeFromValue(v string) (*VirtualPortServiceType, error) {
	ev := VirtualPortServiceType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for VirtualPortServiceType: valid values are %v", v, AllowedVirtualPortServiceTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v VirtualPortServiceType) IsValid() bool {
	for _, existing := range AllowedVirtualPortServiceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to VirtualPortServiceType value
func (v VirtualPortServiceType) Ptr() *VirtualPortServiceType {
	return &v
}

type NullableVirtualPortServiceType struct {
	value *VirtualPortServiceType
	isSet bool
}

func (v NullableVirtualPortServiceType) Get() *VirtualPortServiceType {
	return v.value
}

func (v *NullableVirtualPortServiceType) Set(val *VirtualPortServiceType) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualPortServiceType) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualPortServiceType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualPortServiceType(val *VirtualPortServiceType) *NullableVirtualPortServiceType {
	return &NullableVirtualPortServiceType{value: val, isSet: true}
}

func (v NullableVirtualPortServiceType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualPortServiceType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
