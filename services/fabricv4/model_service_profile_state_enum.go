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

// ServiceProfileStateEnum Equinix assigned state.
type ServiceProfileStateEnum string

// List of ServiceProfileStateEnum
const (
	SERVICEPROFILESTATEENUM_ACTIVE           ServiceProfileStateEnum = "ACTIVE"
	SERVICEPROFILESTATEENUM_PENDING_APPROVAL ServiceProfileStateEnum = "PENDING_APPROVAL"
	SERVICEPROFILESTATEENUM_DELETED          ServiceProfileStateEnum = "DELETED"
	SERVICEPROFILESTATEENUM_REJECTED         ServiceProfileStateEnum = "REJECTED"
)

// All allowed values of ServiceProfileStateEnum enum
var AllowedServiceProfileStateEnumEnumValues = []ServiceProfileStateEnum{
	"ACTIVE",
	"PENDING_APPROVAL",
	"DELETED",
	"REJECTED",
}

func (v *ServiceProfileStateEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ServiceProfileStateEnum(value)
	for _, existing := range AllowedServiceProfileStateEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ServiceProfileStateEnum", value)
}

// NewServiceProfileStateEnumFromValue returns a pointer to a valid ServiceProfileStateEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewServiceProfileStateEnumFromValue(v string) (*ServiceProfileStateEnum, error) {
	ev := ServiceProfileStateEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ServiceProfileStateEnum: valid values are %v", v, AllowedServiceProfileStateEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ServiceProfileStateEnum) IsValid() bool {
	for _, existing := range AllowedServiceProfileStateEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ServiceProfileStateEnum value
func (v ServiceProfileStateEnum) Ptr() *ServiceProfileStateEnum {
	return &v
}

type NullableServiceProfileStateEnum struct {
	value *ServiceProfileStateEnum
	isSet bool
}

func (v NullableServiceProfileStateEnum) Get() *ServiceProfileStateEnum {
	return v.value
}

func (v *NullableServiceProfileStateEnum) Set(val *ServiceProfileStateEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceProfileStateEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceProfileStateEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceProfileStateEnum(val *ServiceProfileStateEnum) *NullableServiceProfileStateEnum {
	return &NullableServiceProfileStateEnum{value: val, isSet: true}
}

func (v NullableServiceProfileStateEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceProfileStateEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
