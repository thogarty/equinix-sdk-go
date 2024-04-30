/*
Equinix Internet Access API

Equinix Internet Access provides direct access to the Internet with scalable bandwidth options in IBX data centers.

API version: 2.0
Contact: api-support@equinix.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eiav2

import (
	"encoding/json"
	"fmt"
)

// UseCaseType Service use case
type UseCaseType string

// List of UseCaseType
const (
	USECASETYPE_MAIN              UseCaseType = "MAIN"
	USECASETYPE_BACKUP            UseCaseType = "BACKUP"
	USECASETYPE_MANAGEMENT_ACCESS UseCaseType = "MANAGEMENT_ACCESS"
)

// All allowed values of UseCaseType enum
var AllowedUseCaseTypeEnumValues = []UseCaseType{
	"MAIN",
	"BACKUP",
	"MANAGEMENT_ACCESS",
}

func (v *UseCaseType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := UseCaseType(value)
	for _, existing := range AllowedUseCaseTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid UseCaseType", value)
}

// NewUseCaseTypeFromValue returns a pointer to a valid UseCaseType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewUseCaseTypeFromValue(v string) (*UseCaseType, error) {
	ev := UseCaseType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for UseCaseType: valid values are %v", v, AllowedUseCaseTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v UseCaseType) IsValid() bool {
	for _, existing := range AllowedUseCaseTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to UseCaseType value
func (v UseCaseType) Ptr() *UseCaseType {
	return &v
}

type NullableUseCaseType struct {
	value *UseCaseType
	isSet bool
}

func (v NullableUseCaseType) Get() *UseCaseType {
	return v.value
}

func (v *NullableUseCaseType) Set(val *UseCaseType) {
	v.value = val
	v.isSet = true
}

func (v NullableUseCaseType) IsSet() bool {
	return v.isSet
}

func (v *NullableUseCaseType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUseCaseType(val *UseCaseType) *NullableUseCaseType {
	return &NullableUseCaseType{value: val, isSet: true}
}

func (v NullableUseCaseType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUseCaseType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
