/*
Equinix Fabric API v4

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fabricv4

import (
	"encoding/json"
	"fmt"
)

// ChangeType Type of change
type ChangeType string

// List of Change_type
const (
	CHANGETYPE_CREATION                ChangeType = "CONNECTION_CREATION"
	CHANGETYPE_UPDATE                  ChangeType = "CONNECTION_UPDATE"
	CHANGETYPE_DELETION                ChangeType = "CONNECTION_DELETION"
	CHANGETYPE_PROVIDER_STATUS_REQUEST ChangeType = "CONNECTION_PROVIDER_STATUS_REQUEST"
)

// All allowed values of ChangeType enum
var AllowedChangeTypeEnumValues = []ChangeType{
	"CONNECTION_CREATION",
	"CONNECTION_UPDATE",
	"CONNECTION_DELETION",
	"CONNECTION_PROVIDER_STATUS_REQUEST",
}

func (v *ChangeType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ChangeType(value)
	for _, existing := range AllowedChangeTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ChangeType", value)
}

// NewChangeTypeFromValue returns a pointer to a valid ChangeType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewChangeTypeFromValue(v string) (*ChangeType, error) {
	ev := ChangeType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ChangeType: valid values are %v", v, AllowedChangeTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ChangeType) IsValid() bool {
	for _, existing := range AllowedChangeTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Change_type value
func (v ChangeType) Ptr() *ChangeType {
	return &v
}

type NullableChangeType struct {
	value *ChangeType
	isSet bool
}

func (v NullableChangeType) Get() *ChangeType {
	return v.value
}

func (v *NullableChangeType) Set(val *ChangeType) {
	v.value = val
	v.isSet = true
}

func (v NullableChangeType) IsSet() bool {
	return v.isSet
}

func (v *NullableChangeType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChangeType(val *ChangeType) *NullableChangeType {
	return &NullableChangeType{value: val, isSet: true}
}

func (v NullableChangeType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChangeType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
