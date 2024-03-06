/*
Equinix Fabric API v4

Equinix Fabric is an advanced software-defined interconnection solution that enables you to directly, securely and dynamically connect to distributed infrastructure and digital ecosystems on platform Equinix via a single port, Customers can use Fabric to connect to: </br> 1. Cloud Service Providers - Clouds, network and other service providers.  </br> 2. Enterprises - Other Equinix customers, vendors and partners.  </br> 3. Myself - Another customer instance deployed at Equinix. </br>

API version: 4.12
Contact: api-support@equinix.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fabricv4

import (
	"encoding/json"
	"fmt"
)

// CloudRouterAccessPointState Access point lifecycle state
type CloudRouterAccessPointState string

// List of CloudRouterAccessPointState
const (
	CLOUDROUTERACCESSPOINTSTATE_PROVISIONED       CloudRouterAccessPointState = "PROVISIONED"
	CLOUDROUTERACCESSPOINTSTATE_PROVISIONING      CloudRouterAccessPointState = "PROVISIONING"
	CLOUDROUTERACCESSPOINTSTATE_DEPROVISIONING    CloudRouterAccessPointState = "DEPROVISIONING"
	CLOUDROUTERACCESSPOINTSTATE_DEPROVISIONED     CloudRouterAccessPointState = "DEPROVISIONED"
	CLOUDROUTERACCESSPOINTSTATE_LOCKED            CloudRouterAccessPointState = "LOCKED"
	CLOUDROUTERACCESSPOINTSTATE_NOT_PROVISIONED   CloudRouterAccessPointState = "NOT_PROVISIONED"
	CLOUDROUTERACCESSPOINTSTATE_NOT_DEPROVISIONED CloudRouterAccessPointState = "NOT_DEPROVISIONED"
)

// All allowed values of CloudRouterAccessPointState enum
var AllowedCloudRouterAccessPointStateEnumValues = []CloudRouterAccessPointState{
	"PROVISIONED",
	"PROVISIONING",
	"DEPROVISIONING",
	"DEPROVISIONED",
	"LOCKED",
	"NOT_PROVISIONED",
	"NOT_DEPROVISIONED",
}

func (v *CloudRouterAccessPointState) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CloudRouterAccessPointState(value)
	for _, existing := range AllowedCloudRouterAccessPointStateEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CloudRouterAccessPointState", value)
}

// NewCloudRouterAccessPointStateFromValue returns a pointer to a valid CloudRouterAccessPointState
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCloudRouterAccessPointStateFromValue(v string) (*CloudRouterAccessPointState, error) {
	ev := CloudRouterAccessPointState(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CloudRouterAccessPointState: valid values are %v", v, AllowedCloudRouterAccessPointStateEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CloudRouterAccessPointState) IsValid() bool {
	for _, existing := range AllowedCloudRouterAccessPointStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CloudRouterAccessPointState value
func (v CloudRouterAccessPointState) Ptr() *CloudRouterAccessPointState {
	return &v
}

type NullableCloudRouterAccessPointState struct {
	value *CloudRouterAccessPointState
	isSet bool
}

func (v NullableCloudRouterAccessPointState) Get() *CloudRouterAccessPointState {
	return v.value
}

func (v *NullableCloudRouterAccessPointState) Set(val *CloudRouterAccessPointState) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudRouterAccessPointState) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudRouterAccessPointState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudRouterAccessPointState(val *CloudRouterAccessPointState) *NullableCloudRouterAccessPointState {
	return &NullableCloudRouterAccessPointState{value: val, isSet: true}
}

func (v NullableCloudRouterAccessPointState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudRouterAccessPointState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}