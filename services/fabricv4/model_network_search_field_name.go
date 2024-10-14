/*
Equinix Fabric API v4

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fabricv4

import (
	"encoding/json"
	"fmt"
)

// NetworkSearchFieldName Possible field names to use on filters
type NetworkSearchFieldName string

// List of NetworkSearchFieldName
const (
	NETWORKSEARCHFIELDNAME_NAME                     NetworkSearchFieldName = "/name"
	NETWORKSEARCHFIELDNAME_UUID                     NetworkSearchFieldName = "/uuid"
	NETWORKSEARCHFIELDNAME_SCOPE                    NetworkSearchFieldName = "/scope"
	NETWORKSEARCHFIELDNAME_TYPE                     NetworkSearchFieldName = "/type"
	NETWORKSEARCHFIELDNAME_OPERATION_EQUINIX_STATUS NetworkSearchFieldName = "/operation/equinixStatus"
	NETWORKSEARCHFIELDNAME_LOCATION_REGION          NetworkSearchFieldName = "/location/region"
	NETWORKSEARCHFIELDNAME_PROJECT_PROJECT_ID       NetworkSearchFieldName = "/project/projectId"
	NETWORKSEARCHFIELDNAME_ACCOUNT_GLOBAL_CUST_ID   NetworkSearchFieldName = "/account/globalCustId"
	NETWORKSEARCHFIELDNAME_ACCOUNT_ORG_ID           NetworkSearchFieldName = "/account/orgId"
	NETWORKSEARCHFIELDNAME_DELETED_DATE             NetworkSearchFieldName = "/deletedDate"
	NETWORKSEARCHFIELDNAME_STAR                     NetworkSearchFieldName = "/_*"
)

// All allowed values of NetworkSearchFieldName enum
var AllowedNetworkSearchFieldNameEnumValues = []NetworkSearchFieldName{
	"/name",
	"/uuid",
	"/scope",
	"/type",
	"/operation/equinixStatus",
	"/location/region",
	"/project/projectId",
	"/account/globalCustId",
	"/account/orgId",
	"/deletedDate",
	"/_*",
}

func (v *NetworkSearchFieldName) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := NetworkSearchFieldName(value)
	for _, existing := range AllowedNetworkSearchFieldNameEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid NetworkSearchFieldName", value)
}

// NewNetworkSearchFieldNameFromValue returns a pointer to a valid NetworkSearchFieldName
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNetworkSearchFieldNameFromValue(v string) (*NetworkSearchFieldName, error) {
	ev := NetworkSearchFieldName(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NetworkSearchFieldName: valid values are %v", v, AllowedNetworkSearchFieldNameEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NetworkSearchFieldName) IsValid() bool {
	for _, existing := range AllowedNetworkSearchFieldNameEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to NetworkSearchFieldName value
func (v NetworkSearchFieldName) Ptr() *NetworkSearchFieldName {
	return &v
}

type NullableNetworkSearchFieldName struct {
	value *NetworkSearchFieldName
	isSet bool
}

func (v NullableNetworkSearchFieldName) Get() *NetworkSearchFieldName {
	return v.value
}

func (v *NullableNetworkSearchFieldName) Set(val *NetworkSearchFieldName) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkSearchFieldName) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkSearchFieldName) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkSearchFieldName(val *NetworkSearchFieldName) *NullableNetworkSearchFieldName {
	return &NullableNetworkSearchFieldName{value: val, isSet: true}
}

func (v NullableNetworkSearchFieldName) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkSearchFieldName) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
