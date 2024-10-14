/*
Equinix Fabric API v4

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fabricv4

import (
	"encoding/json"
	"fmt"
)

// PriceChargeType Price charge type
type PriceChargeType string

// List of PriceCharge_type
const (
	PRICECHARGETYPE_MONTHLY_RECURRING PriceChargeType = "MONTHLY_RECURRING"
	PRICECHARGETYPE_NON_RECURRING     PriceChargeType = "NON_RECURRING"
)

// All allowed values of PriceChargeType enum
var AllowedPriceChargeTypeEnumValues = []PriceChargeType{
	"MONTHLY_RECURRING",
	"NON_RECURRING",
}

func (v *PriceChargeType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PriceChargeType(value)
	for _, existing := range AllowedPriceChargeTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PriceChargeType", value)
}

// NewPriceChargeTypeFromValue returns a pointer to a valid PriceChargeType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPriceChargeTypeFromValue(v string) (*PriceChargeType, error) {
	ev := PriceChargeType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PriceChargeType: valid values are %v", v, AllowedPriceChargeTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PriceChargeType) IsValid() bool {
	for _, existing := range AllowedPriceChargeTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PriceCharge_type value
func (v PriceChargeType) Ptr() *PriceChargeType {
	return &v
}

type NullablePriceChargeType struct {
	value *PriceChargeType
	isSet bool
}

func (v NullablePriceChargeType) Get() *PriceChargeType {
	return v.value
}

func (v *NullablePriceChargeType) Set(val *PriceChargeType) {
	v.value = val
	v.isSet = true
}

func (v NullablePriceChargeType) IsSet() bool {
	return v.isSet
}

func (v *NullablePriceChargeType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePriceChargeType(val *PriceChargeType) *NullablePriceChargeType {
	return &NullablePriceChargeType{value: val, isSet: true}
}

func (v NullablePriceChargeType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePriceChargeType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
