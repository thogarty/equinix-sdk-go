/*
Equinix Internet Access API

Equinix Internet Access provides direct access to the Internet with scalable bandwidth options in IBX data centers.

Contact: api-support@equinix.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eiav2

import (
	"encoding/json"
	"fmt"
)

// checks if the RoutingProtocolIpv4 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RoutingProtocolIpv4{}

// RoutingProtocolIpv4 struct for RoutingProtocolIpv4
type RoutingProtocolIpv4 struct {
	CustomerRoutes       []RoutingProtocolCustomerRouteIpv4 `json:"customerRoutes"`
	Peerings             []RoutingProtocolPeeringIpv4       `json:"peerings"`
	AdditionalProperties map[string]interface{}
}

type _RoutingProtocolIpv4 RoutingProtocolIpv4

// NewRoutingProtocolIpv4 instantiates a new RoutingProtocolIpv4 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoutingProtocolIpv4(customerRoutes []RoutingProtocolCustomerRouteIpv4, peerings []RoutingProtocolPeeringIpv4) *RoutingProtocolIpv4 {
	this := RoutingProtocolIpv4{}
	this.CustomerRoutes = customerRoutes
	this.Peerings = peerings
	return &this
}

// NewRoutingProtocolIpv4WithDefaults instantiates a new RoutingProtocolIpv4 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoutingProtocolIpv4WithDefaults() *RoutingProtocolIpv4 {
	this := RoutingProtocolIpv4{}
	return &this
}

// GetCustomerRoutes returns the CustomerRoutes field value
func (o *RoutingProtocolIpv4) GetCustomerRoutes() []RoutingProtocolCustomerRouteIpv4 {
	if o == nil {
		var ret []RoutingProtocolCustomerRouteIpv4
		return ret
	}

	return o.CustomerRoutes
}

// GetCustomerRoutesOk returns a tuple with the CustomerRoutes field value
// and a boolean to check if the value has been set.
func (o *RoutingProtocolIpv4) GetCustomerRoutesOk() ([]RoutingProtocolCustomerRouteIpv4, bool) {
	if o == nil {
		return nil, false
	}
	return o.CustomerRoutes, true
}

// SetCustomerRoutes sets field value
func (o *RoutingProtocolIpv4) SetCustomerRoutes(v []RoutingProtocolCustomerRouteIpv4) {
	o.CustomerRoutes = v
}

// GetPeerings returns the Peerings field value
func (o *RoutingProtocolIpv4) GetPeerings() []RoutingProtocolPeeringIpv4 {
	if o == nil {
		var ret []RoutingProtocolPeeringIpv4
		return ret
	}

	return o.Peerings
}

// GetPeeringsOk returns a tuple with the Peerings field value
// and a boolean to check if the value has been set.
func (o *RoutingProtocolIpv4) GetPeeringsOk() ([]RoutingProtocolPeeringIpv4, bool) {
	if o == nil {
		return nil, false
	}
	return o.Peerings, true
}

// SetPeerings sets field value
func (o *RoutingProtocolIpv4) SetPeerings(v []RoutingProtocolPeeringIpv4) {
	o.Peerings = v
}

func (o RoutingProtocolIpv4) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RoutingProtocolIpv4) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["customerRoutes"] = o.CustomerRoutes
	toSerialize["peerings"] = o.Peerings

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RoutingProtocolIpv4) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"customerRoutes",
		"peerings",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varRoutingProtocolIpv4 := _RoutingProtocolIpv4{}

	err = json.Unmarshal(data, &varRoutingProtocolIpv4)

	if err != nil {
		return err
	}

	*o = RoutingProtocolIpv4(varRoutingProtocolIpv4)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "customerRoutes")
		delete(additionalProperties, "peerings")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRoutingProtocolIpv4 struct {
	value *RoutingProtocolIpv4
	isSet bool
}

func (v NullableRoutingProtocolIpv4) Get() *RoutingProtocolIpv4 {
	return v.value
}

func (v *NullableRoutingProtocolIpv4) Set(val *RoutingProtocolIpv4) {
	v.value = val
	v.isSet = true
}

func (v NullableRoutingProtocolIpv4) IsSet() bool {
	return v.isSet
}

func (v *NullableRoutingProtocolIpv4) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoutingProtocolIpv4(val *RoutingProtocolIpv4) *NullableRoutingProtocolIpv4 {
	return &NullableRoutingProtocolIpv4{value: val, isSet: true}
}

func (v NullableRoutingProtocolIpv4) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoutingProtocolIpv4) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}