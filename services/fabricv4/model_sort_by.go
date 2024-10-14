/*
Equinix Fabric API v4

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fabricv4

import (
	"encoding/json"
	"fmt"
)

// SortBy Possible field names to use on sorting
type SortBy string

// List of SortBy
const (
	SORTBY_NAME                                         SortBy = "/name"
	SORTBY_DIRECTION                                    SortBy = "/direction"
	SORTBY_A_SIDE_ACCESS_POINT_NAME                     SortBy = "/aSide/accessPoint/name"
	SORTBY_A_SIDE_ACCESS_POINT_TYPE                     SortBy = "/aSide/accessPoint/type"
	SORTBY_A_SIDE_ACCESS_POINT_ACCOUNT_ACCOUNT_NAME     SortBy = "/aSide/accessPoint/account/accountName"
	SORTBY_A_SIDE_ACCESS_POINT_LOCATION_METRO_NAME      SortBy = "/aSide/accessPoint/location/metroName"
	SORTBY_A_SIDE_ACCESS_POINT_LOCATION_METRO_CODE      SortBy = "/aSide/accessPoint/location/metroCode"
	SORTBY_A_SIDE_ACCESS_POINT_LINK_PROTOCOL_VLAN_C_TAG SortBy = "/aSide/accessPoint/linkProtocol/vlanCTag"
	SORTBY_A_SIDE_ACCESS_POINT_LINK_PROTOCOL_VLAN_S_TAG SortBy = "/aSide/accessPoint/linkProtocol/vlanSTag"
	SORTBY_Z_SIDE_ACCESS_POINT_NAME                     SortBy = "/zSide/accessPoint/name"
	SORTBY_Z_SIDE_ACCESS_POINT_TYPE                     SortBy = "/zSide/accessPoint/type"
	SORTBY_Z_SIDE_ACCESS_POINT_ACCOUNT_ACCOUNT_NAME     SortBy = "/zSide/accessPoint/account/accountName"
	SORTBY_Z_SIDE_ACCESS_POINT_LOCATION_METRO_NAME      SortBy = "/zSide/accessPoint/location/metroName"
	SORTBY_Z_SIDE_ACCESS_POINT_LOCATION_METRO_CODE      SortBy = "/zSide/accessPoint/location/metroCode"
	SORTBY_Z_SIDE_ACCESS_POINT_LINK_PROTOCOL_VLAN_C_TAG SortBy = "/zSide/accessPoint/linkProtocol/vlanCTag"
	SORTBY_Z_SIDE_ACCESS_POINT_LINK_PROTOCOL_VLAN_S_TAG SortBy = "/zSide/accessPoint/linkProtocol/vlanSTag"
	SORTBY_Z_SIDE_ACCESS_POINT_AUTHENTICATION_KEY       SortBy = "/zSide/accessPoint/authenticationKey"
	SORTBY_BANDWIDTH                                    SortBy = "/bandwidth"
	SORTBY_GEO_SCOPE                                    SortBy = "/geoScope"
	SORTBY_UUID                                         SortBy = "/uuid"
	SORTBY_CHANGE_LOG_CREATED_DATE_TIME                 SortBy = "/changeLog/createdDateTime"
	SORTBY_CHANGE_LOG_UPDATED_DATE_TIME                 SortBy = "/changeLog/updatedDateTime"
	SORTBY_OPERATION_EQUINIX_STATUS                     SortBy = "/operation/equinixStatus"
	SORTBY_OPERATION_PROVIDER_STATUS                    SortBy = "/operation/providerStatus"
	SORTBY_REDUNDANCY_PRIORITY                          SortBy = "/redundancy/priority"
)

// All allowed values of SortBy enum
var AllowedSortByEnumValues = []SortBy{
	"/name",
	"/direction",
	"/aSide/accessPoint/name",
	"/aSide/accessPoint/type",
	"/aSide/accessPoint/account/accountName",
	"/aSide/accessPoint/location/metroName",
	"/aSide/accessPoint/location/metroCode",
	"/aSide/accessPoint/linkProtocol/vlanCTag",
	"/aSide/accessPoint/linkProtocol/vlanSTag",
	"/zSide/accessPoint/name",
	"/zSide/accessPoint/type",
	"/zSide/accessPoint/account/accountName",
	"/zSide/accessPoint/location/metroName",
	"/zSide/accessPoint/location/metroCode",
	"/zSide/accessPoint/linkProtocol/vlanCTag",
	"/zSide/accessPoint/linkProtocol/vlanSTag",
	"/zSide/accessPoint/authenticationKey",
	"/bandwidth",
	"/geoScope",
	"/uuid",
	"/changeLog/createdDateTime",
	"/changeLog/updatedDateTime",
	"/operation/equinixStatus",
	"/operation/providerStatus",
	"/redundancy/priority",
}

func (v *SortBy) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SortBy(value)
	for _, existing := range AllowedSortByEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SortBy", value)
}

// NewSortByFromValue returns a pointer to a valid SortBy
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSortByFromValue(v string) (*SortBy, error) {
	ev := SortBy(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SortBy: valid values are %v", v, AllowedSortByEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SortBy) IsValid() bool {
	for _, existing := range AllowedSortByEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SortBy value
func (v SortBy) Ptr() *SortBy {
	return &v
}

type NullableSortBy struct {
	value *SortBy
	isSet bool
}

func (v NullableSortBy) Get() *SortBy {
	return v.value
}

func (v *NullableSortBy) Set(val *SortBy) {
	v.value = val
	v.isSet = true
}

func (v NullableSortBy) IsSet() bool {
	return v.isSet
}

func (v *NullableSortBy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSortBy(val *SortBy) *NullableSortBy {
	return &NullableSortBy{value: val, isSet: true}
}

func (v NullableSortBy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSortBy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
