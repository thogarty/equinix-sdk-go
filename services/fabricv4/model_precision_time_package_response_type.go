/*
Equinix Fabric API v4

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fabricv4

import (
	"encoding/json"
	"fmt"
)

// PrecisionTimePackageResponseType the model 'PrecisionTimePackageResponseType'
type PrecisionTimePackageResponseType string

// List of precisionTimePackageResponse_type
const (
	PRECISIONTIMEPACKAGERESPONSETYPE_TIME_SERVICE_PACKAGE PrecisionTimePackageResponseType = "TIME_SERVICE_PACKAGE"
)

// All allowed values of PrecisionTimePackageResponseType enum
var AllowedPrecisionTimePackageResponseTypeEnumValues = []PrecisionTimePackageResponseType{
	"TIME_SERVICE_PACKAGE",
}

func (v *PrecisionTimePackageResponseType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PrecisionTimePackageResponseType(value)
	for _, existing := range AllowedPrecisionTimePackageResponseTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PrecisionTimePackageResponseType", value)
}

// NewPrecisionTimePackageResponseTypeFromValue returns a pointer to a valid PrecisionTimePackageResponseType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPrecisionTimePackageResponseTypeFromValue(v string) (*PrecisionTimePackageResponseType, error) {
	ev := PrecisionTimePackageResponseType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PrecisionTimePackageResponseType: valid values are %v", v, AllowedPrecisionTimePackageResponseTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PrecisionTimePackageResponseType) IsValid() bool {
	for _, existing := range AllowedPrecisionTimePackageResponseTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to precisionTimePackageResponse_type value
func (v PrecisionTimePackageResponseType) Ptr() *PrecisionTimePackageResponseType {
	return &v
}

type NullablePrecisionTimePackageResponseType struct {
	value *PrecisionTimePackageResponseType
	isSet bool
}

func (v NullablePrecisionTimePackageResponseType) Get() *PrecisionTimePackageResponseType {
	return v.value
}

func (v *NullablePrecisionTimePackageResponseType) Set(val *PrecisionTimePackageResponseType) {
	v.value = val
	v.isSet = true
}

func (v NullablePrecisionTimePackageResponseType) IsSet() bool {
	return v.isSet
}

func (v *NullablePrecisionTimePackageResponseType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrecisionTimePackageResponseType(val *PrecisionTimePackageResponseType) *NullablePrecisionTimePackageResponseType {
	return &NullablePrecisionTimePackageResponseType{value: val, isSet: true}
}

func (v NullablePrecisionTimePackageResponseType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrecisionTimePackageResponseType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
