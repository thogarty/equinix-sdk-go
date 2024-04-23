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

// checks if the LinkProtocolVxlan type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LinkProtocolVxlan{}

// LinkProtocolVxlan Connection link protocol configuration - VXLAN
type LinkProtocolVxlan struct {
	Type *LinkProtocolType `json:"type,omitempty"`
	// Virtual Network Identifier
	Vni                  int32 `json:"vni"`
	AdditionalProperties map[string]interface{}
}

type _LinkProtocolVxlan LinkProtocolVxlan

// NewLinkProtocolVxlan instantiates a new LinkProtocolVxlan object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinkProtocolVxlan(vni int32) *LinkProtocolVxlan {
	this := LinkProtocolVxlan{}
	this.Vni = vni
	return &this
}

// NewLinkProtocolVxlanWithDefaults instantiates a new LinkProtocolVxlan object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinkProtocolVxlanWithDefaults() *LinkProtocolVxlan {
	this := LinkProtocolVxlan{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *LinkProtocolVxlan) GetType() LinkProtocolType {
	if o == nil || IsNil(o.Type) {
		var ret LinkProtocolType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkProtocolVxlan) GetTypeOk() (*LinkProtocolType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *LinkProtocolVxlan) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given LinkProtocolType and assigns it to the Type field.
func (o *LinkProtocolVxlan) SetType(v LinkProtocolType) {
	o.Type = &v
}

// GetVni returns the Vni field value
func (o *LinkProtocolVxlan) GetVni() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Vni
}

// GetVniOk returns a tuple with the Vni field value
// and a boolean to check if the value has been set.
func (o *LinkProtocolVxlan) GetVniOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Vni, true
}

// SetVni sets field value
func (o *LinkProtocolVxlan) SetVni(v int32) {
	o.Vni = v
}

func (o LinkProtocolVxlan) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LinkProtocolVxlan) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	toSerialize["vni"] = o.Vni

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *LinkProtocolVxlan) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"vni",
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

	varLinkProtocolVxlan := _LinkProtocolVxlan{}

	err = json.Unmarshal(data, &varLinkProtocolVxlan)

	if err != nil {
		return err
	}

	*o = LinkProtocolVxlan(varLinkProtocolVxlan)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "vni")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLinkProtocolVxlan struct {
	value *LinkProtocolVxlan
	isSet bool
}

func (v NullableLinkProtocolVxlan) Get() *LinkProtocolVxlan {
	return v.value
}

func (v *NullableLinkProtocolVxlan) Set(val *LinkProtocolVxlan) {
	v.value = val
	v.isSet = true
}

func (v NullableLinkProtocolVxlan) IsSet() bool {
	return v.isSet
}

func (v *NullableLinkProtocolVxlan) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinkProtocolVxlan(val *LinkProtocolVxlan) *NullableLinkProtocolVxlan {
	return &NullableLinkProtocolVxlan{value: val, isSet: true}
}

func (v NullableLinkProtocolVxlan) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinkProtocolVxlan) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
