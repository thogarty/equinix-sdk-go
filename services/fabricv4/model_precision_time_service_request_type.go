/*
Equinix Fabric API v4

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fabricv4

import (
	"encoding/json"
	"fmt"
)

// PrecisionTimeServiceRequestType Precision Time Service Type refers to the corresponding Protocol.
type PrecisionTimeServiceRequestType string

// List of precisionTimeServiceRequest_type
const (
	PRECISIONTIMESERVICEREQUESTTYPE_NTP PrecisionTimeServiceRequestType = "NTP"
	PRECISIONTIMESERVICEREQUESTTYPE_PTP PrecisionTimeServiceRequestType = "PTP"
)

// All allowed values of PrecisionTimeServiceRequestType enum
var AllowedPrecisionTimeServiceRequestTypeEnumValues = []PrecisionTimeServiceRequestType{
	"NTP",
	"PTP",
}

func (v *PrecisionTimeServiceRequestType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PrecisionTimeServiceRequestType(value)
	for _, existing := range AllowedPrecisionTimeServiceRequestTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PrecisionTimeServiceRequestType", value)
}

// NewPrecisionTimeServiceRequestTypeFromValue returns a pointer to a valid PrecisionTimeServiceRequestType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPrecisionTimeServiceRequestTypeFromValue(v string) (*PrecisionTimeServiceRequestType, error) {
	ev := PrecisionTimeServiceRequestType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PrecisionTimeServiceRequestType: valid values are %v", v, AllowedPrecisionTimeServiceRequestTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PrecisionTimeServiceRequestType) IsValid() bool {
	for _, existing := range AllowedPrecisionTimeServiceRequestTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to precisionTimeServiceRequest_type value
func (v PrecisionTimeServiceRequestType) Ptr() *PrecisionTimeServiceRequestType {
	return &v
}

type NullablePrecisionTimeServiceRequestType struct {
	value *PrecisionTimeServiceRequestType
	isSet bool
}

func (v NullablePrecisionTimeServiceRequestType) Get() *PrecisionTimeServiceRequestType {
	return v.value
}

func (v *NullablePrecisionTimeServiceRequestType) Set(val *PrecisionTimeServiceRequestType) {
	v.value = val
	v.isSet = true
}

func (v NullablePrecisionTimeServiceRequestType) IsSet() bool {
	return v.isSet
}

func (v *NullablePrecisionTimeServiceRequestType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrecisionTimeServiceRequestType(val *PrecisionTimeServiceRequestType) *NullablePrecisionTimeServiceRequestType {
	return &NullablePrecisionTimeServiceRequestType{value: val, isSet: true}
}

func (v NullablePrecisionTimeServiceRequestType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrecisionTimeServiceRequestType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
