/*
Equinix Fabric API v4

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fabricv4

import (
	"encoding/json"
	"fmt"
)

// ServiceTokenConnectionType Type of Connection
type ServiceTokenConnectionType string

// List of ServiceTokenConnection_type
const (
	SERVICETOKENCONNECTIONTYPE_EVPL_VC   ServiceTokenConnectionType = "EVPL_VC"
	SERVICETOKENCONNECTIONTYPE_EVPLAN_VC ServiceTokenConnectionType = "EVPLAN_VC"
	SERVICETOKENCONNECTIONTYPE_EPLAN_VC  ServiceTokenConnectionType = "EPLAN_VC"
	SERVICETOKENCONNECTIONTYPE_IPWAN_VC  ServiceTokenConnectionType = "IPWAN_VC"
)

// All allowed values of ServiceTokenConnectionType enum
var AllowedServiceTokenConnectionTypeEnumValues = []ServiceTokenConnectionType{
	"EVPL_VC",
	"EVPLAN_VC",
	"EPLAN_VC",
	"IPWAN_VC",
}

func (v *ServiceTokenConnectionType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ServiceTokenConnectionType(value)
	for _, existing := range AllowedServiceTokenConnectionTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ServiceTokenConnectionType", value)
}

// NewServiceTokenConnectionTypeFromValue returns a pointer to a valid ServiceTokenConnectionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewServiceTokenConnectionTypeFromValue(v string) (*ServiceTokenConnectionType, error) {
	ev := ServiceTokenConnectionType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ServiceTokenConnectionType: valid values are %v", v, AllowedServiceTokenConnectionTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ServiceTokenConnectionType) IsValid() bool {
	for _, existing := range AllowedServiceTokenConnectionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ServiceTokenConnection_type value
func (v ServiceTokenConnectionType) Ptr() *ServiceTokenConnectionType {
	return &v
}

type NullableServiceTokenConnectionType struct {
	value *ServiceTokenConnectionType
	isSet bool
}

func (v NullableServiceTokenConnectionType) Get() *ServiceTokenConnectionType {
	return v.value
}

func (v *NullableServiceTokenConnectionType) Set(val *ServiceTokenConnectionType) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceTokenConnectionType) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceTokenConnectionType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceTokenConnectionType(val *ServiceTokenConnectionType) *NullableServiceTokenConnectionType {
	return &NullableServiceTokenConnectionType{value: val, isSet: true}
}

func (v NullableServiceTokenConnectionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceTokenConnectionType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
