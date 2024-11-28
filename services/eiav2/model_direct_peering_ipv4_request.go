/*
Equinix Internet Access API

Equinix Internet Access provides direct access to the Internet with scalable bandwidth options in IBX data centers.

Contact: api-support@equinix.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eiav2

import (
	"encoding/json"
)

// checks if the DirectPeeringIpv4Request type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DirectPeeringIpv4Request{}

// DirectPeeringIpv4Request struct for DirectPeeringIpv4Request
type DirectPeeringIpv4Request struct {
	Connection *Connection `json:"connection,omitempty"`
	// Peering IP addresses in Version 4 (IPv4)
	EquinixPeerIps []string `json:"equinixPeerIps,omitempty"`
	// Virtual router group IP addresses in Version 4 (IPv4)
	EquinixVRRPIp        *string `json:"equinixVRRPIp,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DirectPeeringIpv4Request DirectPeeringIpv4Request

// NewDirectPeeringIpv4Request instantiates a new DirectPeeringIpv4Request object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDirectPeeringIpv4Request() *DirectPeeringIpv4Request {
	this := DirectPeeringIpv4Request{}
	return &this
}

// NewDirectPeeringIpv4RequestWithDefaults instantiates a new DirectPeeringIpv4Request object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDirectPeeringIpv4RequestWithDefaults() *DirectPeeringIpv4Request {
	this := DirectPeeringIpv4Request{}
	return &this
}

// GetConnection returns the Connection field value if set, zero value otherwise.
func (o *DirectPeeringIpv4Request) GetConnection() Connection {
	if o == nil || IsNil(o.Connection) {
		var ret Connection
		return ret
	}
	return *o.Connection
}

// GetConnectionOk returns a tuple with the Connection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DirectPeeringIpv4Request) GetConnectionOk() (*Connection, bool) {
	if o == nil || IsNil(o.Connection) {
		return nil, false
	}
	return o.Connection, true
}

// HasConnection returns a boolean if a field has been set.
func (o *DirectPeeringIpv4Request) HasConnection() bool {
	if o != nil && !IsNil(o.Connection) {
		return true
	}

	return false
}

// SetConnection gets a reference to the given Connection and assigns it to the Connection field.
func (o *DirectPeeringIpv4Request) SetConnection(v Connection) {
	o.Connection = &v
}

// GetEquinixPeerIps returns the EquinixPeerIps field value if set, zero value otherwise.
func (o *DirectPeeringIpv4Request) GetEquinixPeerIps() []string {
	if o == nil || IsNil(o.EquinixPeerIps) {
		var ret []string
		return ret
	}
	return o.EquinixPeerIps
}

// GetEquinixPeerIpsOk returns a tuple with the EquinixPeerIps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DirectPeeringIpv4Request) GetEquinixPeerIpsOk() ([]string, bool) {
	if o == nil || IsNil(o.EquinixPeerIps) {
		return nil, false
	}
	return o.EquinixPeerIps, true
}

// HasEquinixPeerIps returns a boolean if a field has been set.
func (o *DirectPeeringIpv4Request) HasEquinixPeerIps() bool {
	if o != nil && !IsNil(o.EquinixPeerIps) {
		return true
	}

	return false
}

// SetEquinixPeerIps gets a reference to the given []string and assigns it to the EquinixPeerIps field.
func (o *DirectPeeringIpv4Request) SetEquinixPeerIps(v []string) {
	o.EquinixPeerIps = v
}

// GetEquinixVRRPIp returns the EquinixVRRPIp field value if set, zero value otherwise.
func (o *DirectPeeringIpv4Request) GetEquinixVRRPIp() string {
	if o == nil || IsNil(o.EquinixVRRPIp) {
		var ret string
		return ret
	}
	return *o.EquinixVRRPIp
}

// GetEquinixVRRPIpOk returns a tuple with the EquinixVRRPIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DirectPeeringIpv4Request) GetEquinixVRRPIpOk() (*string, bool) {
	if o == nil || IsNil(o.EquinixVRRPIp) {
		return nil, false
	}
	return o.EquinixVRRPIp, true
}

// HasEquinixVRRPIp returns a boolean if a field has been set.
func (o *DirectPeeringIpv4Request) HasEquinixVRRPIp() bool {
	if o != nil && !IsNil(o.EquinixVRRPIp) {
		return true
	}

	return false
}

// SetEquinixVRRPIp gets a reference to the given string and assigns it to the EquinixVRRPIp field.
func (o *DirectPeeringIpv4Request) SetEquinixVRRPIp(v string) {
	o.EquinixVRRPIp = &v
}

func (o DirectPeeringIpv4Request) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DirectPeeringIpv4Request) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Connection) {
		toSerialize["connection"] = o.Connection
	}
	if !IsNil(o.EquinixPeerIps) {
		toSerialize["equinixPeerIps"] = o.EquinixPeerIps
	}
	if !IsNil(o.EquinixVRRPIp) {
		toSerialize["equinixVRRPIp"] = o.EquinixVRRPIp
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DirectPeeringIpv4Request) UnmarshalJSON(data []byte) (err error) {
	varDirectPeeringIpv4Request := _DirectPeeringIpv4Request{}

	err = json.Unmarshal(data, &varDirectPeeringIpv4Request)

	if err != nil {
		return err
	}

	*o = DirectPeeringIpv4Request(varDirectPeeringIpv4Request)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "connection")
		delete(additionalProperties, "equinixPeerIps")
		delete(additionalProperties, "equinixVRRPIp")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDirectPeeringIpv4Request struct {
	value *DirectPeeringIpv4Request
	isSet bool
}

func (v NullableDirectPeeringIpv4Request) Get() *DirectPeeringIpv4Request {
	return v.value
}

func (v *NullableDirectPeeringIpv4Request) Set(val *DirectPeeringIpv4Request) {
	v.value = val
	v.isSet = true
}

func (v NullableDirectPeeringIpv4Request) IsSet() bool {
	return v.isSet
}

func (v *NullableDirectPeeringIpv4Request) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDirectPeeringIpv4Request(val *DirectPeeringIpv4Request) *NullableDirectPeeringIpv4Request {
	return &NullableDirectPeeringIpv4Request{value: val, isSet: true}
}

func (v NullableDirectPeeringIpv4Request) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDirectPeeringIpv4Request) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
