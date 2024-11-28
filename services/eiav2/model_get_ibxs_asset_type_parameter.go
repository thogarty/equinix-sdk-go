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

// GetIbxsAssetTypeParameter the model 'GetIbxsAssetTypeParameter'
type GetIbxsAssetTypeParameter string

// List of getIbxs_asset_type_parameter
const (
	GETIBXSASSETTYPEPARAMETER_CABINET GetIbxsAssetTypeParameter = "CABINET"
)

// All allowed values of GetIbxsAssetTypeParameter enum
var AllowedGetIbxsAssetTypeParameterEnumValues = []GetIbxsAssetTypeParameter{
	"CABINET",
}

func (v *GetIbxsAssetTypeParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := GetIbxsAssetTypeParameter(value)
	for _, existing := range AllowedGetIbxsAssetTypeParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid GetIbxsAssetTypeParameter", value)
}

// NewGetIbxsAssetTypeParameterFromValue returns a pointer to a valid GetIbxsAssetTypeParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewGetIbxsAssetTypeParameterFromValue(v string) (*GetIbxsAssetTypeParameter, error) {
	ev := GetIbxsAssetTypeParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for GetIbxsAssetTypeParameter: valid values are %v", v, AllowedGetIbxsAssetTypeParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v GetIbxsAssetTypeParameter) IsValid() bool {
	for _, existing := range AllowedGetIbxsAssetTypeParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to getIbxs_asset_type_parameter value
func (v GetIbxsAssetTypeParameter) Ptr() *GetIbxsAssetTypeParameter {
	return &v
}

type NullableGetIbxsAssetTypeParameter struct {
	value *GetIbxsAssetTypeParameter
	isSet bool
}

func (v NullableGetIbxsAssetTypeParameter) Get() *GetIbxsAssetTypeParameter {
	return v.value
}

func (v *NullableGetIbxsAssetTypeParameter) Set(val *GetIbxsAssetTypeParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableGetIbxsAssetTypeParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableGetIbxsAssetTypeParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetIbxsAssetTypeParameter(val *GetIbxsAssetTypeParameter) *NullableGetIbxsAssetTypeParameter {
	return &NullableGetIbxsAssetTypeParameter{value: val, isSet: true}
}

func (v NullableGetIbxsAssetTypeParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetIbxsAssetTypeParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}