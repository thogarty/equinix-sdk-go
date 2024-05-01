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

// ConnectionType Connection type. IA_C - internet access connection on dedicated port; IA_VC - internet access virtual connection on Fabric Port, Virtual Device, Bare Metal, etc.
type ConnectionType string

// List of ConnectionType
const (
	CONNECTIONTYPE_C  ConnectionType = "IA_C"
	CONNECTIONTYPE_VC ConnectionType = "IA_VC"
)

// All allowed values of ConnectionType enum
var AllowedConnectionTypeEnumValues = []ConnectionType{
	"IA_C",
	"IA_VC",
}

func (v *ConnectionType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ConnectionType(value)
	for _, existing := range AllowedConnectionTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ConnectionType", value)
}

// NewConnectionTypeFromValue returns a pointer to a valid ConnectionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewConnectionTypeFromValue(v string) (*ConnectionType, error) {
	ev := ConnectionType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ConnectionType: valid values are %v", v, AllowedConnectionTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ConnectionType) IsValid() bool {
	for _, existing := range AllowedConnectionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ConnectionType value
func (v ConnectionType) Ptr() *ConnectionType {
	return &v
}

type NullableConnectionType struct {
	value *ConnectionType
	isSet bool
}

func (v NullableConnectionType) Get() *ConnectionType {
	return v.value
}

func (v *NullableConnectionType) Set(val *ConnectionType) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectionType) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectionType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectionType(val *ConnectionType) *NullableConnectionType {
	return &NullableConnectionType{value: val, isSet: true}
}

func (v NullableConnectionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectionType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
