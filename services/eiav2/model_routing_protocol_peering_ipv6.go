/*
Equinix Internet Access API

Equinix Internet Access provides direct access to the Internet with scalable bandwidth options in IBX data centers.

API version: 2.0
Contact: api-support@equinix.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eiav2

import (
	"encoding/json"
	"fmt"
)

// checks if the RoutingProtocolPeeringIpv6 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RoutingProtocolPeeringIpv6{}

// RoutingProtocolPeeringIpv6 struct for RoutingProtocolPeeringIpv6
type RoutingProtocolPeeringIpv6 struct {
	Connections     []RoutingProtocolPeeringConnectionItem `json:"connections,omitempty"`
	CustomerPeerIps []string                               `json:"customerPeerIps,omitempty"`
	PeerSubnet      RoutingProtocolPeeringIpv6PeerSubnet   `json:"peerSubnet"`
	// Indicates if VRRP is enabled.
	VrrpEnabled          bool     `json:"vrrpEnabled"`
	EquinixPeerIps       []string `json:"equinixPeerIps"`
	EquinixVRRPIp        *string  `json:"equinixVRRPIp,omitempty"`
	CustomerVRRPIp       *string  `json:"customerVRRPIp,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RoutingProtocolPeeringIpv6 RoutingProtocolPeeringIpv6

// NewRoutingProtocolPeeringIpv6 instantiates a new RoutingProtocolPeeringIpv6 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoutingProtocolPeeringIpv6(peerSubnet RoutingProtocolPeeringIpv6PeerSubnet, vrrpEnabled bool, equinixPeerIps []string) *RoutingProtocolPeeringIpv6 {
	this := RoutingProtocolPeeringIpv6{}
	this.PeerSubnet = peerSubnet
	this.VrrpEnabled = vrrpEnabled
	this.EquinixPeerIps = equinixPeerIps
	return &this
}

// NewRoutingProtocolPeeringIpv6WithDefaults instantiates a new RoutingProtocolPeeringIpv6 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoutingProtocolPeeringIpv6WithDefaults() *RoutingProtocolPeeringIpv6 {
	this := RoutingProtocolPeeringIpv6{}
	return &this
}

// GetConnections returns the Connections field value if set, zero value otherwise.
func (o *RoutingProtocolPeeringIpv6) GetConnections() []RoutingProtocolPeeringConnectionItem {
	if o == nil || IsNil(o.Connections) {
		var ret []RoutingProtocolPeeringConnectionItem
		return ret
	}
	return o.Connections
}

// GetConnectionsOk returns a tuple with the Connections field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoutingProtocolPeeringIpv6) GetConnectionsOk() ([]RoutingProtocolPeeringConnectionItem, bool) {
	if o == nil || IsNil(o.Connections) {
		return nil, false
	}
	return o.Connections, true
}

// HasConnections returns a boolean if a field has been set.
func (o *RoutingProtocolPeeringIpv6) HasConnections() bool {
	if o != nil && !IsNil(o.Connections) {
		return true
	}

	return false
}

// SetConnections gets a reference to the given []RoutingProtocolPeeringConnectionItem and assigns it to the Connections field.
func (o *RoutingProtocolPeeringIpv6) SetConnections(v []RoutingProtocolPeeringConnectionItem) {
	o.Connections = v
}

// GetCustomerPeerIps returns the CustomerPeerIps field value if set, zero value otherwise.
func (o *RoutingProtocolPeeringIpv6) GetCustomerPeerIps() []string {
	if o == nil || IsNil(o.CustomerPeerIps) {
		var ret []string
		return ret
	}
	return o.CustomerPeerIps
}

// GetCustomerPeerIpsOk returns a tuple with the CustomerPeerIps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoutingProtocolPeeringIpv6) GetCustomerPeerIpsOk() ([]string, bool) {
	if o == nil || IsNil(o.CustomerPeerIps) {
		return nil, false
	}
	return o.CustomerPeerIps, true
}

// HasCustomerPeerIps returns a boolean if a field has been set.
func (o *RoutingProtocolPeeringIpv6) HasCustomerPeerIps() bool {
	if o != nil && !IsNil(o.CustomerPeerIps) {
		return true
	}

	return false
}

// SetCustomerPeerIps gets a reference to the given []string and assigns it to the CustomerPeerIps field.
func (o *RoutingProtocolPeeringIpv6) SetCustomerPeerIps(v []string) {
	o.CustomerPeerIps = v
}

// GetPeerSubnet returns the PeerSubnet field value
func (o *RoutingProtocolPeeringIpv6) GetPeerSubnet() RoutingProtocolPeeringIpv6PeerSubnet {
	if o == nil {
		var ret RoutingProtocolPeeringIpv6PeerSubnet
		return ret
	}

	return o.PeerSubnet
}

// GetPeerSubnetOk returns a tuple with the PeerSubnet field value
// and a boolean to check if the value has been set.
func (o *RoutingProtocolPeeringIpv6) GetPeerSubnetOk() (*RoutingProtocolPeeringIpv6PeerSubnet, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PeerSubnet, true
}

// SetPeerSubnet sets field value
func (o *RoutingProtocolPeeringIpv6) SetPeerSubnet(v RoutingProtocolPeeringIpv6PeerSubnet) {
	o.PeerSubnet = v
}

// GetVrrpEnabled returns the VrrpEnabled field value
func (o *RoutingProtocolPeeringIpv6) GetVrrpEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.VrrpEnabled
}

// GetVrrpEnabledOk returns a tuple with the VrrpEnabled field value
// and a boolean to check if the value has been set.
func (o *RoutingProtocolPeeringIpv6) GetVrrpEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VrrpEnabled, true
}

// SetVrrpEnabled sets field value
func (o *RoutingProtocolPeeringIpv6) SetVrrpEnabled(v bool) {
	o.VrrpEnabled = v
}

// GetEquinixPeerIps returns the EquinixPeerIps field value
func (o *RoutingProtocolPeeringIpv6) GetEquinixPeerIps() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.EquinixPeerIps
}

// GetEquinixPeerIpsOk returns a tuple with the EquinixPeerIps field value
// and a boolean to check if the value has been set.
func (o *RoutingProtocolPeeringIpv6) GetEquinixPeerIpsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EquinixPeerIps, true
}

// SetEquinixPeerIps sets field value
func (o *RoutingProtocolPeeringIpv6) SetEquinixPeerIps(v []string) {
	o.EquinixPeerIps = v
}

// GetEquinixVRRPIp returns the EquinixVRRPIp field value if set, zero value otherwise.
func (o *RoutingProtocolPeeringIpv6) GetEquinixVRRPIp() string {
	if o == nil || IsNil(o.EquinixVRRPIp) {
		var ret string
		return ret
	}
	return *o.EquinixVRRPIp
}

// GetEquinixVRRPIpOk returns a tuple with the EquinixVRRPIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoutingProtocolPeeringIpv6) GetEquinixVRRPIpOk() (*string, bool) {
	if o == nil || IsNil(o.EquinixVRRPIp) {
		return nil, false
	}
	return o.EquinixVRRPIp, true
}

// HasEquinixVRRPIp returns a boolean if a field has been set.
func (o *RoutingProtocolPeeringIpv6) HasEquinixVRRPIp() bool {
	if o != nil && !IsNil(o.EquinixVRRPIp) {
		return true
	}

	return false
}

// SetEquinixVRRPIp gets a reference to the given string and assigns it to the EquinixVRRPIp field.
func (o *RoutingProtocolPeeringIpv6) SetEquinixVRRPIp(v string) {
	o.EquinixVRRPIp = &v
}

// GetCustomerVRRPIp returns the CustomerVRRPIp field value if set, zero value otherwise.
func (o *RoutingProtocolPeeringIpv6) GetCustomerVRRPIp() string {
	if o == nil || IsNil(o.CustomerVRRPIp) {
		var ret string
		return ret
	}
	return *o.CustomerVRRPIp
}

// GetCustomerVRRPIpOk returns a tuple with the CustomerVRRPIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoutingProtocolPeeringIpv6) GetCustomerVRRPIpOk() (*string, bool) {
	if o == nil || IsNil(o.CustomerVRRPIp) {
		return nil, false
	}
	return o.CustomerVRRPIp, true
}

// HasCustomerVRRPIp returns a boolean if a field has been set.
func (o *RoutingProtocolPeeringIpv6) HasCustomerVRRPIp() bool {
	if o != nil && !IsNil(o.CustomerVRRPIp) {
		return true
	}

	return false
}

// SetCustomerVRRPIp gets a reference to the given string and assigns it to the CustomerVRRPIp field.
func (o *RoutingProtocolPeeringIpv6) SetCustomerVRRPIp(v string) {
	o.CustomerVRRPIp = &v
}

func (o RoutingProtocolPeeringIpv6) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RoutingProtocolPeeringIpv6) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Connections) {
		toSerialize["connections"] = o.Connections
	}
	if !IsNil(o.CustomerPeerIps) {
		toSerialize["customerPeerIps"] = o.CustomerPeerIps
	}
	toSerialize["peerSubnet"] = o.PeerSubnet
	toSerialize["vrrpEnabled"] = o.VrrpEnabled
	toSerialize["equinixPeerIps"] = o.EquinixPeerIps
	if !IsNil(o.EquinixVRRPIp) {
		toSerialize["equinixVRRPIp"] = o.EquinixVRRPIp
	}
	if !IsNil(o.CustomerVRRPIp) {
		toSerialize["customerVRRPIp"] = o.CustomerVRRPIp
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RoutingProtocolPeeringIpv6) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"peerSubnet",
		"vrrpEnabled",
		"equinixPeerIps",
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

	varRoutingProtocolPeeringIpv6 := _RoutingProtocolPeeringIpv6{}

	err = json.Unmarshal(data, &varRoutingProtocolPeeringIpv6)

	if err != nil {
		return err
	}

	*o = RoutingProtocolPeeringIpv6(varRoutingProtocolPeeringIpv6)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "connections")
		delete(additionalProperties, "customerPeerIps")
		delete(additionalProperties, "peerSubnet")
		delete(additionalProperties, "vrrpEnabled")
		delete(additionalProperties, "equinixPeerIps")
		delete(additionalProperties, "equinixVRRPIp")
		delete(additionalProperties, "customerVRRPIp")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRoutingProtocolPeeringIpv6 struct {
	value *RoutingProtocolPeeringIpv6
	isSet bool
}

func (v NullableRoutingProtocolPeeringIpv6) Get() *RoutingProtocolPeeringIpv6 {
	return v.value
}

func (v *NullableRoutingProtocolPeeringIpv6) Set(val *RoutingProtocolPeeringIpv6) {
	v.value = val
	v.isSet = true
}

func (v NullableRoutingProtocolPeeringIpv6) IsSet() bool {
	return v.isSet
}

func (v *NullableRoutingProtocolPeeringIpv6) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoutingProtocolPeeringIpv6(val *RoutingProtocolPeeringIpv6) *NullableRoutingProtocolPeeringIpv6 {
	return &NullableRoutingProtocolPeeringIpv6{value: val, isSet: true}
}

func (v NullableRoutingProtocolPeeringIpv6) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoutingProtocolPeeringIpv6) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
