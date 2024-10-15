/*
Equinix Fabric API v4

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fabricv4

import (
	"encoding/json"
	"fmt"
)

// Asset the model 'Asset'
type Asset string

// List of Asset
const (
	ASSET_PORTS       Asset = "ports"
	ASSET_CONNECTIONS Asset = "connections"
	ASSET_ROUTERS     Asset = "routers"
	ASSET_METROS      Asset = "metros"
)

// All allowed values of Asset enum
var AllowedAssetEnumValues = []Asset{
	"ports",
	"connections",
	"routers",
	"metros",
}

func (v *Asset) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := Asset(value)
	for _, existing := range AllowedAssetEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid Asset", value)
}

// NewAssetFromValue returns a pointer to a valid Asset
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAssetFromValue(v string) (*Asset, error) {
	ev := Asset(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for Asset: valid values are %v", v, AllowedAssetEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v Asset) IsValid() bool {
	for _, existing := range AllowedAssetEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Asset value
func (v Asset) Ptr() *Asset {
	return &v
}

type NullableAsset struct {
	value *Asset
	isSet bool
}

func (v NullableAsset) Get() *Asset {
	return v.value
}

func (v *NullableAsset) Set(val *Asset) {
	v.value = val
	v.isSet = true
}

func (v NullableAsset) IsSet() bool {
	return v.isSet
}

func (v *NullableAsset) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAsset(val *Asset) *NullableAsset {
	return &NullableAsset{value: val, isSet: true}
}

func (v NullableAsset) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAsset) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}