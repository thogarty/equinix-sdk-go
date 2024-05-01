/*
Equinix Internet Access API

Equinix Internet Access provides direct access to the Internet with scalable bandwidth options in IBX data centers.

Contact: api-support@equinix.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eiav2

import (
	"encoding/json"
	"fmt"
)

// ServiceUseCase use case of the service
type ServiceUseCase string

// List of ServiceUseCase
const (
	SERVICEUSECASE_MAIN              ServiceUseCase = "MAIN"
	SERVICEUSECASE_BACKUP            ServiceUseCase = "BACKUP"
	SERVICEUSECASE_MANAGEMENT_ACCESS ServiceUseCase = "MANAGEMENT_ACCESS"
)

// All allowed values of ServiceUseCase enum
var AllowedServiceUseCaseEnumValues = []ServiceUseCase{
	"MAIN",
	"BACKUP",
	"MANAGEMENT_ACCESS",
}

func (v *ServiceUseCase) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ServiceUseCase(value)
	for _, existing := range AllowedServiceUseCaseEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ServiceUseCase", value)
}

// NewServiceUseCaseFromValue returns a pointer to a valid ServiceUseCase
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewServiceUseCaseFromValue(v string) (*ServiceUseCase, error) {
	ev := ServiceUseCase(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ServiceUseCase: valid values are %v", v, AllowedServiceUseCaseEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ServiceUseCase) IsValid() bool {
	for _, existing := range AllowedServiceUseCaseEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ServiceUseCase value
func (v ServiceUseCase) Ptr() *ServiceUseCase {
	return &v
}

type NullableServiceUseCase struct {
	value *ServiceUseCase
	isSet bool
}

func (v NullableServiceUseCase) Get() *ServiceUseCase {
	return v.value
}

func (v *NullableServiceUseCase) Set(val *ServiceUseCase) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceUseCase) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceUseCase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceUseCase(val *ServiceUseCase) *NullableServiceUseCase {
	return &NullableServiceUseCase{value: val, isSet: true}
}

func (v NullableServiceUseCase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceUseCase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
