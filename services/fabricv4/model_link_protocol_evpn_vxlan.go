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

// checks if the LinkProtocolEvpnVxlan type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LinkProtocolEvpnVxlan{}

// LinkProtocolEvpnVxlan Connection link protocol configuration - EVPN_VXLAN
type LinkProtocolEvpnVxlan struct {
	Type *LinkProtocolType `json:"type,omitempty"`
	// Virtual Network Identifier
	Vnid int32 `json:"vnid"`
	// Type 5 VNI identifier
	Type5vni             int32 `json:"type5vni"`
	AdditionalProperties map[string]interface{}
}

type _LinkProtocolEvpnVxlan LinkProtocolEvpnVxlan

// NewLinkProtocolEvpnVxlan instantiates a new LinkProtocolEvpnVxlan object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinkProtocolEvpnVxlan(vnid int32, type5vni int32) *LinkProtocolEvpnVxlan {
	this := LinkProtocolEvpnVxlan{}
	this.Vnid = vnid
	this.Type5vni = type5vni
	return &this
}

// NewLinkProtocolEvpnVxlanWithDefaults instantiates a new LinkProtocolEvpnVxlan object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinkProtocolEvpnVxlanWithDefaults() *LinkProtocolEvpnVxlan {
	this := LinkProtocolEvpnVxlan{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *LinkProtocolEvpnVxlan) GetType() LinkProtocolType {
	if o == nil || IsNil(o.Type) {
		var ret LinkProtocolType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkProtocolEvpnVxlan) GetTypeOk() (*LinkProtocolType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *LinkProtocolEvpnVxlan) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given LinkProtocolType and assigns it to the Type field.
func (o *LinkProtocolEvpnVxlan) SetType(v LinkProtocolType) {
	o.Type = &v
}

// GetVnid returns the Vnid field value
func (o *LinkProtocolEvpnVxlan) GetVnid() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Vnid
}

// GetVnidOk returns a tuple with the Vnid field value
// and a boolean to check if the value has been set.
func (o *LinkProtocolEvpnVxlan) GetVnidOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Vnid, true
}

// SetVnid sets field value
func (o *LinkProtocolEvpnVxlan) SetVnid(v int32) {
	o.Vnid = v
}

// GetType5vni returns the Type5vni field value
func (o *LinkProtocolEvpnVxlan) GetType5vni() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Type5vni
}

// GetType5vniOk returns a tuple with the Type5vni field value
// and a boolean to check if the value has been set.
func (o *LinkProtocolEvpnVxlan) GetType5vniOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type5vni, true
}

// SetType5vni sets field value
func (o *LinkProtocolEvpnVxlan) SetType5vni(v int32) {
	o.Type5vni = v
}

func (o LinkProtocolEvpnVxlan) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LinkProtocolEvpnVxlan) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	toSerialize["vnid"] = o.Vnid
	toSerialize["type5vni"] = o.Type5vni

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *LinkProtocolEvpnVxlan) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"vnid",
		"type5vni",
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

	varLinkProtocolEvpnVxlan := _LinkProtocolEvpnVxlan{}

	err = json.Unmarshal(data, &varLinkProtocolEvpnVxlan)

	if err != nil {
		return err
	}

	*o = LinkProtocolEvpnVxlan(varLinkProtocolEvpnVxlan)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "vnid")
		delete(additionalProperties, "type5vni")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLinkProtocolEvpnVxlan struct {
	value *LinkProtocolEvpnVxlan
	isSet bool
}

func (v NullableLinkProtocolEvpnVxlan) Get() *LinkProtocolEvpnVxlan {
	return v.value
}

func (v *NullableLinkProtocolEvpnVxlan) Set(val *LinkProtocolEvpnVxlan) {
	v.value = val
	v.isSet = true
}

func (v NullableLinkProtocolEvpnVxlan) IsSet() bool {
	return v.isSet
}

func (v *NullableLinkProtocolEvpnVxlan) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinkProtocolEvpnVxlan(val *LinkProtocolEvpnVxlan) *NullableLinkProtocolEvpnVxlan {
	return &NullableLinkProtocolEvpnVxlan{value: val, isSet: true}
}

func (v NullableLinkProtocolEvpnVxlan) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinkProtocolEvpnVxlan) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
