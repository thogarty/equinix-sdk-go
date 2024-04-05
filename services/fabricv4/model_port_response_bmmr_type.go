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

// PortResponseBmmrType the model 'PortResponseBmmrType'
type PortResponseBmmrType string

// List of PortResponse_bmmrType
const (
	PORTRESPONSEBMMRTYPE_SELF    PortResponseBmmrType = "SELF"
	PORTRESPONSEBMMRTYPE_EQUINIX PortResponseBmmrType = "EQUINIX"
)

// All allowed values of PortResponseBmmrType enum
var AllowedPortResponseBmmrTypeEnumValues = []PortResponseBmmrType{
	"SELF",
	"EQUINIX",
}

func (v *PortResponseBmmrType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PortResponseBmmrType(value)
	for _, existing := range AllowedPortResponseBmmrTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PortResponseBmmrType", value)
}

// NewPortResponseBmmrTypeFromValue returns a pointer to a valid PortResponseBmmrType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPortResponseBmmrTypeFromValue(v string) (*PortResponseBmmrType, error) {
	ev := PortResponseBmmrType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PortResponseBmmrType: valid values are %v", v, AllowedPortResponseBmmrTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PortResponseBmmrType) IsValid() bool {
	for _, existing := range AllowedPortResponseBmmrTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PortResponse_bmmrType value
func (v PortResponseBmmrType) Ptr() *PortResponseBmmrType {
	return &v
}

type NullablePortResponseBmmrType struct {
	value *PortResponseBmmrType
	isSet bool
}

func (v NullablePortResponseBmmrType) Get() *PortResponseBmmrType {
	return v.value
}

func (v *NullablePortResponseBmmrType) Set(val *PortResponseBmmrType) {
	v.value = val
	v.isSet = true
}

func (v NullablePortResponseBmmrType) IsSet() bool {
	return v.isSet
}

func (v *NullablePortResponseBmmrType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePortResponseBmmrType(val *PortResponseBmmrType) *NullablePortResponseBmmrType {
	return &NullablePortResponseBmmrType{value: val, isSet: true}
}

func (v NullablePortResponseBmmrType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePortResponseBmmrType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}