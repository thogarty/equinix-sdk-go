/*
Equinix Fabric API v4

Equinix Fabric is an advanced software-defined interconnection solution that enables you to directly, securely and dynamically connect to distributed infrastructure and digital ecosystems on platform Equinix via a single port, Customers can use Fabric to connect to: </br> 1. Cloud Service Providers - Clouds, network and other service providers.  </br> 2. Enterprises - Other Equinix customers, vendors and partners.  </br> 3. Myself - Another customer instance deployed at Equinix. </br>

Contact: api-support@equinix.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fabricv4

import (
	"encoding/json"
	"fmt"
)

// RoutingProtocolBGPDataState the model 'RoutingProtocolBGPDataState'
type RoutingProtocolBGPDataState string

// List of RoutingProtocolBGPData_state
const (
	ROUTINGPROTOCOLBGPDATASTATE_PROVISIONED    RoutingProtocolBGPDataState = "PROVISIONED"
	ROUTINGPROTOCOLBGPDATASTATE_DEPROVISIONED  RoutingProtocolBGPDataState = "DEPROVISIONED"
	ROUTINGPROTOCOLBGPDATASTATE_PROVISIONING   RoutingProtocolBGPDataState = "PROVISIONING"
	ROUTINGPROTOCOLBGPDATASTATE_DEPROVISIONING RoutingProtocolBGPDataState = "DEPROVISIONING"
	ROUTINGPROTOCOLBGPDATASTATE_REPROVISIONING RoutingProtocolBGPDataState = "REPROVISIONING"
	ROUTINGPROTOCOLBGPDATASTATE_FAILED         RoutingProtocolBGPDataState = "FAILED"
)

// All allowed values of RoutingProtocolBGPDataState enum
var AllowedRoutingProtocolBGPDataStateEnumValues = []RoutingProtocolBGPDataState{
	"PROVISIONED",
	"DEPROVISIONED",
	"PROVISIONING",
	"DEPROVISIONING",
	"REPROVISIONING",
	"FAILED",
}

func (v *RoutingProtocolBGPDataState) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RoutingProtocolBGPDataState(value)
	for _, existing := range AllowedRoutingProtocolBGPDataStateEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RoutingProtocolBGPDataState", value)
}

// NewRoutingProtocolBGPDataStateFromValue returns a pointer to a valid RoutingProtocolBGPDataState
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRoutingProtocolBGPDataStateFromValue(v string) (*RoutingProtocolBGPDataState, error) {
	ev := RoutingProtocolBGPDataState(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RoutingProtocolBGPDataState: valid values are %v", v, AllowedRoutingProtocolBGPDataStateEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RoutingProtocolBGPDataState) IsValid() bool {
	for _, existing := range AllowedRoutingProtocolBGPDataStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RoutingProtocolBGPData_state value
func (v RoutingProtocolBGPDataState) Ptr() *RoutingProtocolBGPDataState {
	return &v
}

type NullableRoutingProtocolBGPDataState struct {
	value *RoutingProtocolBGPDataState
	isSet bool
}

func (v NullableRoutingProtocolBGPDataState) Get() *RoutingProtocolBGPDataState {
	return v.value
}

func (v *NullableRoutingProtocolBGPDataState) Set(val *RoutingProtocolBGPDataState) {
	v.value = val
	v.isSet = true
}

func (v NullableRoutingProtocolBGPDataState) IsSet() bool {
	return v.isSet
}

func (v *NullableRoutingProtocolBGPDataState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoutingProtocolBGPDataState(val *RoutingProtocolBGPDataState) *NullableRoutingProtocolBGPDataState {
	return &NullableRoutingProtocolBGPDataState{value: val, isSet: true}
}

func (v NullableRoutingProtocolBGPDataState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoutingProtocolBGPDataState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
