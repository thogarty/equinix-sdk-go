/*
Equinix Fabric API v4

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fabricv4

import (
	"encoding/json"
	"fmt"
)

// ServiceTokenSearchFieldName Possible field names to use on filters
type ServiceTokenSearchFieldName string

// List of ServiceTokenSearchFieldName
const (
	SERVICETOKENSEARCHFIELDNAME_UUID               ServiceTokenSearchFieldName = "/uuid"
	SERVICETOKENSEARCHFIELDNAME_STATE              ServiceTokenSearchFieldName = "/state"
	SERVICETOKENSEARCHFIELDNAME_NAME               ServiceTokenSearchFieldName = "/name"
	SERVICETOKENSEARCHFIELDNAME_PROJECT_PROJECT_ID ServiceTokenSearchFieldName = "/project/projectId"
)

// All allowed values of ServiceTokenSearchFieldName enum
var AllowedServiceTokenSearchFieldNameEnumValues = []ServiceTokenSearchFieldName{
	"/uuid",
	"/state",
	"/name",
	"/project/projectId",
}

func (v *ServiceTokenSearchFieldName) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ServiceTokenSearchFieldName(value)
	for _, existing := range AllowedServiceTokenSearchFieldNameEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ServiceTokenSearchFieldName", value)
}

// NewServiceTokenSearchFieldNameFromValue returns a pointer to a valid ServiceTokenSearchFieldName
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewServiceTokenSearchFieldNameFromValue(v string) (*ServiceTokenSearchFieldName, error) {
	ev := ServiceTokenSearchFieldName(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ServiceTokenSearchFieldName: valid values are %v", v, AllowedServiceTokenSearchFieldNameEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ServiceTokenSearchFieldName) IsValid() bool {
	for _, existing := range AllowedServiceTokenSearchFieldNameEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ServiceTokenSearchFieldName value
func (v ServiceTokenSearchFieldName) Ptr() *ServiceTokenSearchFieldName {
	return &v
}

type NullableServiceTokenSearchFieldName struct {
	value *ServiceTokenSearchFieldName
	isSet bool
}

func (v NullableServiceTokenSearchFieldName) Get() *ServiceTokenSearchFieldName {
	return v.value
}

func (v *NullableServiceTokenSearchFieldName) Set(val *ServiceTokenSearchFieldName) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceTokenSearchFieldName) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceTokenSearchFieldName) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceTokenSearchFieldName(val *ServiceTokenSearchFieldName) *NullableServiceTokenSearchFieldName {
	return &NullableServiceTokenSearchFieldName{value: val, isSet: true}
}

func (v NullableServiceTokenSearchFieldName) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceTokenSearchFieldName) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
