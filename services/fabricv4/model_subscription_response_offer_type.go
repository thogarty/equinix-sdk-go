/*
Equinix Fabric API v4

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fabricv4

import (
	"encoding/json"
	"fmt"
)

// SubscriptionResponseOfferType Marketplace Offer Type
type SubscriptionResponseOfferType string

// List of SubscriptionResponse_offerType
const (
	SUBSCRIPTIONRESPONSEOFFERTYPE_PUBLIC        SubscriptionResponseOfferType = "PUBLIC"
	SUBSCRIPTIONRESPONSEOFFERTYPE_PRIVATE_OFFER SubscriptionResponseOfferType = "PRIVATE_OFFER"
)

// All allowed values of SubscriptionResponseOfferType enum
var AllowedSubscriptionResponseOfferTypeEnumValues = []SubscriptionResponseOfferType{
	"PUBLIC",
	"PRIVATE_OFFER",
}

func (v *SubscriptionResponseOfferType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SubscriptionResponseOfferType(value)
	for _, existing := range AllowedSubscriptionResponseOfferTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SubscriptionResponseOfferType", value)
}

// NewSubscriptionResponseOfferTypeFromValue returns a pointer to a valid SubscriptionResponseOfferType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSubscriptionResponseOfferTypeFromValue(v string) (*SubscriptionResponseOfferType, error) {
	ev := SubscriptionResponseOfferType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SubscriptionResponseOfferType: valid values are %v", v, AllowedSubscriptionResponseOfferTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SubscriptionResponseOfferType) IsValid() bool {
	for _, existing := range AllowedSubscriptionResponseOfferTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SubscriptionResponse_offerType value
func (v SubscriptionResponseOfferType) Ptr() *SubscriptionResponseOfferType {
	return &v
}

type NullableSubscriptionResponseOfferType struct {
	value *SubscriptionResponseOfferType
	isSet bool
}

func (v NullableSubscriptionResponseOfferType) Get() *SubscriptionResponseOfferType {
	return v.value
}

func (v *NullableSubscriptionResponseOfferType) Set(val *SubscriptionResponseOfferType) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionResponseOfferType) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionResponseOfferType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionResponseOfferType(val *SubscriptionResponseOfferType) *NullableSubscriptionResponseOfferType {
	return &NullableSubscriptionResponseOfferType{value: val, isSet: true}
}

func (v NullableSubscriptionResponseOfferType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionResponseOfferType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
