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

// ServiceState the model 'ServiceState'
type ServiceState string

// List of ServiceState
const (
	SERVICESTATE_PENDING        ServiceState = "PENDING"
	SERVICESTATE_PROVISIONING   ServiceState = "PROVISIONING"
	SERVICESTATE_ACTIVE         ServiceState = "ACTIVE"
	SERVICESTATE_INACTIVE       ServiceState = "INACTIVE"
	SERVICESTATE_DEPROVISIONING ServiceState = "DEPROVISIONING"
	SERVICESTATE_DEPROVISIONED  ServiceState = "DEPROVISIONED"
	SERVICESTATE_FAILED         ServiceState = "FAILED"
)

// All allowed values of ServiceState enum
var AllowedServiceStateEnumValues = []ServiceState{
	"PENDING",
	"PROVISIONING",
	"ACTIVE",
	"INACTIVE",
	"DEPROVISIONING",
	"DEPROVISIONED",
	"FAILED",
}

func (v *ServiceState) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ServiceState(value)
	for _, existing := range AllowedServiceStateEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ServiceState", value)
}

// NewServiceStateFromValue returns a pointer to a valid ServiceState
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewServiceStateFromValue(v string) (*ServiceState, error) {
	ev := ServiceState(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ServiceState: valid values are %v", v, AllowedServiceStateEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ServiceState) IsValid() bool {
	for _, existing := range AllowedServiceStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ServiceState value
func (v ServiceState) Ptr() *ServiceState {
	return &v
}

type NullableServiceState struct {
	value *ServiceState
	isSet bool
}

func (v NullableServiceState) Get() *ServiceState {
	return v.value
}

func (v *NullableServiceState) Set(val *ServiceState) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceState) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceState(val *ServiceState) *NullableServiceState {
	return &NullableServiceState{value: val, isSet: true}
}

func (v NullableServiceState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
