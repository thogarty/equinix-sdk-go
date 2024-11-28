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

// PurchaseOrderReadModelType Order type
type PurchaseOrderReadModelType string

// List of PurchaseOrderReadModel_type
const (
	PURCHASEORDERREADMODELTYPE_STANDARD_PURCHASE_ORDER PurchaseOrderReadModelType = "STANDARD_PURCHASE_ORDER"
	PURCHASEORDERREADMODELTYPE_BLANKET_PURCHASE_ORDER  PurchaseOrderReadModelType = "BLANKET_PURCHASE_ORDER"
)

// All allowed values of PurchaseOrderReadModelType enum
var AllowedPurchaseOrderReadModelTypeEnumValues = []PurchaseOrderReadModelType{
	"STANDARD_PURCHASE_ORDER",
	"BLANKET_PURCHASE_ORDER",
}

func (v *PurchaseOrderReadModelType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PurchaseOrderReadModelType(value)
	for _, existing := range AllowedPurchaseOrderReadModelTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PurchaseOrderReadModelType", value)
}

// NewPurchaseOrderReadModelTypeFromValue returns a pointer to a valid PurchaseOrderReadModelType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPurchaseOrderReadModelTypeFromValue(v string) (*PurchaseOrderReadModelType, error) {
	ev := PurchaseOrderReadModelType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PurchaseOrderReadModelType: valid values are %v", v, AllowedPurchaseOrderReadModelTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PurchaseOrderReadModelType) IsValid() bool {
	for _, existing := range AllowedPurchaseOrderReadModelTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PurchaseOrderReadModel_type value
func (v PurchaseOrderReadModelType) Ptr() *PurchaseOrderReadModelType {
	return &v
}

type NullablePurchaseOrderReadModelType struct {
	value *PurchaseOrderReadModelType
	isSet bool
}

func (v NullablePurchaseOrderReadModelType) Get() *PurchaseOrderReadModelType {
	return v.value
}

func (v *NullablePurchaseOrderReadModelType) Set(val *PurchaseOrderReadModelType) {
	v.value = val
	v.isSet = true
}

func (v NullablePurchaseOrderReadModelType) IsSet() bool {
	return v.isSet
}

func (v *NullablePurchaseOrderReadModelType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePurchaseOrderReadModelType(val *PurchaseOrderReadModelType) *NullablePurchaseOrderReadModelType {
	return &NullablePurchaseOrderReadModelType{value: val, isSet: true}
}

func (v NullablePurchaseOrderReadModelType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePurchaseOrderReadModelType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
