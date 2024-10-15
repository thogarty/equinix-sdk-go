/*
Equinix Fabric API v4

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fabricv4

import (
	"encoding/json"
	"fmt"
)

// ChangeStatus Current outcome of the change flow
type ChangeStatus string

// List of Change_status
const (
	CHANGESTATUS_APPROVED               ChangeStatus = "APPROVED"
	CHANGESTATUS_COMPLETED              ChangeStatus = "COMPLETED"
	CHANGESTATUS_FAILED                 ChangeStatus = "FAILED"
	CHANGESTATUS_REJECTED               ChangeStatus = "REJECTED"
	CHANGESTATUS_REQUESTED              ChangeStatus = "REQUESTED"
	CHANGESTATUS_SUBMITTED_FOR_APPROVAL ChangeStatus = "SUBMITTED_FOR_APPROVAL"
)

// All allowed values of ChangeStatus enum
var AllowedChangeStatusEnumValues = []ChangeStatus{
	"APPROVED",
	"COMPLETED",
	"FAILED",
	"REJECTED",
	"REQUESTED",
	"SUBMITTED_FOR_APPROVAL",
}

func (v *ChangeStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ChangeStatus(value)
	for _, existing := range AllowedChangeStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ChangeStatus", value)
}

// NewChangeStatusFromValue returns a pointer to a valid ChangeStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewChangeStatusFromValue(v string) (*ChangeStatus, error) {
	ev := ChangeStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ChangeStatus: valid values are %v", v, AllowedChangeStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ChangeStatus) IsValid() bool {
	for _, existing := range AllowedChangeStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Change_status value
func (v ChangeStatus) Ptr() *ChangeStatus {
	return &v
}

type NullableChangeStatus struct {
	value *ChangeStatus
	isSet bool
}

func (v NullableChangeStatus) Get() *ChangeStatus {
	return v.value
}

func (v *NullableChangeStatus) Set(val *ChangeStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableChangeStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableChangeStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChangeStatus(val *ChangeStatus) *NullableChangeStatus {
	return &NullableChangeStatus{value: val, isSet: true}
}

func (v NullableChangeStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChangeStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
