/*
Equinix Fabric API v4

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fabricv4

import (
	"encoding/json"
	"fmt"
)

// Md5Type the model 'Md5Type'
type Md5Type string

// List of md5_type
const (
	MD5TYPE_ASCII Md5Type = "ASCII"
	MD5TYPE_HEX   Md5Type = "HEX"
)

// All allowed values of Md5Type enum
var AllowedMd5TypeEnumValues = []Md5Type{
	"ASCII",
	"HEX",
}

func (v *Md5Type) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := Md5Type(value)
	for _, existing := range AllowedMd5TypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid Md5Type", value)
}

// NewMd5TypeFromValue returns a pointer to a valid Md5Type
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMd5TypeFromValue(v string) (*Md5Type, error) {
	ev := Md5Type(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for Md5Type: valid values are %v", v, AllowedMd5TypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v Md5Type) IsValid() bool {
	for _, existing := range AllowedMd5TypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to md5_type value
func (v Md5Type) Ptr() *Md5Type {
	return &v
}

type NullableMd5Type struct {
	value *Md5Type
	isSet bool
}

func (v NullableMd5Type) Get() *Md5Type {
	return v.value
}

func (v *NullableMd5Type) Set(val *Md5Type) {
	v.value = val
	v.isSet = true
}

func (v NullableMd5Type) IsSet() bool {
	return v.isSet
}

func (v *NullableMd5Type) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMd5Type(val *Md5Type) *NullableMd5Type {
	return &NullableMd5Type{value: val, isSet: true}
}

func (v NullableMd5Type) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMd5Type) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
