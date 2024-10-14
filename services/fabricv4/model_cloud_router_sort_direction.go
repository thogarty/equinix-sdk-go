/*
Equinix Fabric API v4

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fabricv4

import (
	"encoding/json"
	"fmt"
)

// CloudRouterSortDirection Sorting direction
type CloudRouterSortDirection string

// List of CloudRouterSortDirection
const (
	CLOUDROUTERSORTDIRECTION_DESC CloudRouterSortDirection = "DESC"
	CLOUDROUTERSORTDIRECTION_ASC  CloudRouterSortDirection = "ASC"
)

// All allowed values of CloudRouterSortDirection enum
var AllowedCloudRouterSortDirectionEnumValues = []CloudRouterSortDirection{
	"DESC",
	"ASC",
}

func (v *CloudRouterSortDirection) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CloudRouterSortDirection(value)
	for _, existing := range AllowedCloudRouterSortDirectionEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CloudRouterSortDirection", value)
}

// NewCloudRouterSortDirectionFromValue returns a pointer to a valid CloudRouterSortDirection
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCloudRouterSortDirectionFromValue(v string) (*CloudRouterSortDirection, error) {
	ev := CloudRouterSortDirection(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CloudRouterSortDirection: valid values are %v", v, AllowedCloudRouterSortDirectionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CloudRouterSortDirection) IsValid() bool {
	for _, existing := range AllowedCloudRouterSortDirectionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CloudRouterSortDirection value
func (v CloudRouterSortDirection) Ptr() *CloudRouterSortDirection {
	return &v
}

type NullableCloudRouterSortDirection struct {
	value *CloudRouterSortDirection
	isSet bool
}

func (v NullableCloudRouterSortDirection) Get() *CloudRouterSortDirection {
	return v.value
}

func (v *NullableCloudRouterSortDirection) Set(val *CloudRouterSortDirection) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudRouterSortDirection) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudRouterSortDirection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudRouterSortDirection(val *CloudRouterSortDirection) *NullableCloudRouterSortDirection {
	return &NullableCloudRouterSortDirection{value: val, isSet: true}
}

func (v NullableCloudRouterSortDirection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudRouterSortDirection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
