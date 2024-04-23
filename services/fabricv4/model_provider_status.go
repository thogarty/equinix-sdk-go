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

// ProviderStatus Connection provider readiness status
type ProviderStatus string

// List of ProviderStatus
const (
	PROVIDERSTATUS_AVAILABLE             ProviderStatus = "AVAILABLE"
	PROVIDERSTATUS_DEPROVISIONED         ProviderStatus = "DEPROVISIONED"
	PROVIDERSTATUS_DEPROVISIONING        ProviderStatus = "DEPROVISIONING"
	PROVIDERSTATUS_FAILED                ProviderStatus = "FAILED"
	PROVIDERSTATUS_NOT_AVAILABLE         ProviderStatus = "NOT_AVAILABLE"
	PROVIDERSTATUS_PENDING_APPROVAL      ProviderStatus = "PENDING_APPROVAL"
	PROVIDERSTATUS_PENDING_CONFIGURATION ProviderStatus = "PENDING_CONFIGURATION"
	PROVIDERSTATUS_PROVISIONED           ProviderStatus = "PROVISIONED"
	PROVIDERSTATUS_PROVISIONING          ProviderStatus = "PROVISIONING"
	PROVIDERSTATUS_REJECTED              ProviderStatus = "REJECTED"
	PROVIDERSTATUS_PENDING_BGP           ProviderStatus = "PENDING_BGP"
	PROVIDERSTATUS_OUT_OF_BANDWIDTH      ProviderStatus = "OUT_OF_BANDWIDTH"
	PROVIDERSTATUS_DELETED               ProviderStatus = "DELETED"
	PROVIDERSTATUS_ERROR                 ProviderStatus = "ERROR"
	PROVIDERSTATUS_ERRORED               ProviderStatus = "ERRORED"
	PROVIDERSTATUS_NOTPROVISIONED        ProviderStatus = "NOTPROVISIONED"
	PROVIDERSTATUS_NOT_PROVISIONED       ProviderStatus = "NOT_PROVISIONED"
	PROVIDERSTATUS_ORDERING              ProviderStatus = "ORDERING"
	PROVIDERSTATUS_DELETING              ProviderStatus = "DELETING"
	PROVIDERSTATUS_PENDING_DELETE        ProviderStatus = "PENDING DELETE"
	PROVIDERSTATUS_N_A                   ProviderStatus = "N/A"
)

// All allowed values of ProviderStatus enum
var AllowedProviderStatusEnumValues = []ProviderStatus{
	"AVAILABLE",
	"DEPROVISIONED",
	"DEPROVISIONING",
	"FAILED",
	"NOT_AVAILABLE",
	"PENDING_APPROVAL",
	"PENDING_CONFIGURATION",
	"PROVISIONED",
	"PROVISIONING",
	"REJECTED",
	"PENDING_BGP",
	"OUT_OF_BANDWIDTH",
	"DELETED",
	"ERROR",
	"ERRORED",
	"NOTPROVISIONED",
	"NOT_PROVISIONED",
	"ORDERING",
	"DELETING",
	"PENDING DELETE",
	"N/A",
}

func (v *ProviderStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ProviderStatus(value)
	for _, existing := range AllowedProviderStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ProviderStatus", value)
}

// NewProviderStatusFromValue returns a pointer to a valid ProviderStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProviderStatusFromValue(v string) (*ProviderStatus, error) {
	ev := ProviderStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProviderStatus: valid values are %v", v, AllowedProviderStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProviderStatus) IsValid() bool {
	for _, existing := range AllowedProviderStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ProviderStatus value
func (v ProviderStatus) Ptr() *ProviderStatus {
	return &v
}

type NullableProviderStatus struct {
	value *ProviderStatus
	isSet bool
}

func (v NullableProviderStatus) Get() *ProviderStatus {
	return v.value
}

func (v *NullableProviderStatus) Set(val *ProviderStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableProviderStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableProviderStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProviderStatus(val *ProviderStatus) *NullableProviderStatus {
	return &NullableProviderStatus{value: val, isSet: true}
}

func (v NullableProviderStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProviderStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
