/*
Equinix Fabric API v4

Contact: api-support@equinix.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fabricv4

import (
	"encoding/json"
	"fmt"
)

// SubscriptionStatus the model 'SubscriptionStatus'
type SubscriptionStatus string

// List of SubscriptionStatus
const (
	SUBSCRIPTIONSTATUS_ACTIVE       SubscriptionStatus = "ACTIVE"
	SUBSCRIPTIONSTATUS_EXPIRED      SubscriptionStatus = "EXPIRED"
	SUBSCRIPTIONSTATUS_CANCELLED    SubscriptionStatus = "CANCELLED"
	SUBSCRIPTIONSTATUS_GRACE_PERIOD SubscriptionStatus = "GRACE_PERIOD"
)

// All allowed values of SubscriptionStatus enum
var AllowedSubscriptionStatusEnumValues = []SubscriptionStatus{
	"ACTIVE",
	"EXPIRED",
	"CANCELLED",
	"GRACE_PERIOD",
}

func (v *SubscriptionStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SubscriptionStatus(value)
	for _, existing := range AllowedSubscriptionStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SubscriptionStatus", value)
}

// NewSubscriptionStatusFromValue returns a pointer to a valid SubscriptionStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSubscriptionStatusFromValue(v string) (*SubscriptionStatus, error) {
	ev := SubscriptionStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SubscriptionStatus: valid values are %v", v, AllowedSubscriptionStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SubscriptionStatus) IsValid() bool {
	for _, existing := range AllowedSubscriptionStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SubscriptionStatus value
func (v SubscriptionStatus) Ptr() *SubscriptionStatus {
	return &v
}

type NullableSubscriptionStatus struct {
	value *SubscriptionStatus
	isSet bool
}

func (v NullableSubscriptionStatus) Get() *SubscriptionStatus {
	return v.value
}

func (v *NullableSubscriptionStatus) Set(val *SubscriptionStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionStatus(val *SubscriptionStatus) *NullableSubscriptionStatus {
	return &NullableSubscriptionStatus{value: val, isSet: true}
}

func (v NullableSubscriptionStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}