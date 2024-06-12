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

// BGPConnectionOperationOperationalStatus BGP IPv4 or IPv6 Connection State operational status
type BGPConnectionOperationOperationalStatus string

// List of BGPConnectionOperation_operationalStatus
const (
	BGPCONNECTIONOPERATIONOPERATIONALSTATUS_UP      BGPConnectionOperationOperationalStatus = "UP"
	BGPCONNECTIONOPERATIONOPERATIONALSTATUS_DOWN    BGPConnectionOperationOperationalStatus = "DOWN"
	BGPCONNECTIONOPERATIONOPERATIONALSTATUS_UNKNOWN BGPConnectionOperationOperationalStatus = "UNKNOWN"
)

// All allowed values of BGPConnectionOperationOperationalStatus enum
var AllowedBGPConnectionOperationOperationalStatusEnumValues = []BGPConnectionOperationOperationalStatus{
	"UP",
	"DOWN",
	"UNKNOWN",
}

func (v *BGPConnectionOperationOperationalStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := BGPConnectionOperationOperationalStatus(value)
	for _, existing := range AllowedBGPConnectionOperationOperationalStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid BGPConnectionOperationOperationalStatus", value)
}

// NewBGPConnectionOperationOperationalStatusFromValue returns a pointer to a valid BGPConnectionOperationOperationalStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBGPConnectionOperationOperationalStatusFromValue(v string) (*BGPConnectionOperationOperationalStatus, error) {
	ev := BGPConnectionOperationOperationalStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BGPConnectionOperationOperationalStatus: valid values are %v", v, AllowedBGPConnectionOperationOperationalStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BGPConnectionOperationOperationalStatus) IsValid() bool {
	for _, existing := range AllowedBGPConnectionOperationOperationalStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to BGPConnectionOperation_operationalStatus value
func (v BGPConnectionOperationOperationalStatus) Ptr() *BGPConnectionOperationOperationalStatus {
	return &v
}

type NullableBGPConnectionOperationOperationalStatus struct {
	value *BGPConnectionOperationOperationalStatus
	isSet bool
}

func (v NullableBGPConnectionOperationOperationalStatus) Get() *BGPConnectionOperationOperationalStatus {
	return v.value
}

func (v *NullableBGPConnectionOperationOperationalStatus) Set(val *BGPConnectionOperationOperationalStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableBGPConnectionOperationOperationalStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableBGPConnectionOperationOperationalStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBGPConnectionOperationOperationalStatus(val *BGPConnectionOperationOperationalStatus) *NullableBGPConnectionOperationOperationalStatus {
	return &NullableBGPConnectionOperationOperationalStatus{value: val, isSet: true}
}

func (v NullableBGPConnectionOperationOperationalStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBGPConnectionOperationOperationalStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
