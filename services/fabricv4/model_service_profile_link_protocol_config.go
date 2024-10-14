/*
Equinix Fabric API v4

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fabricv4

import (
	"encoding/json"
)

// checks if the ServiceProfileLinkProtocolConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceProfileLinkProtocolConfig{}

// ServiceProfileLinkProtocolConfig Configuration for dot1q to qinq translation support
type ServiceProfileLinkProtocolConfig struct {
	EncapsulationStrategy *ServiceProfileLinkProtocolConfigEncapsulationStrategy `json:"encapsulationStrategy,omitempty"`
	NamedTags             []string                                               `json:"namedTags,omitempty"`
	// was ctagLabel
	VlanCTagLabel        *string                                        `json:"vlanCTagLabel,omitempty"`
	ReuseVlanSTag        *bool                                          `json:"reuseVlanSTag,omitempty"`
	Encapsulation        *ServiceProfileLinkProtocolConfigEncapsulation `json:"encapsulation,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ServiceProfileLinkProtocolConfig ServiceProfileLinkProtocolConfig

// NewServiceProfileLinkProtocolConfig instantiates a new ServiceProfileLinkProtocolConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceProfileLinkProtocolConfig() *ServiceProfileLinkProtocolConfig {
	this := ServiceProfileLinkProtocolConfig{}
	var reuseVlanSTag bool = false
	this.ReuseVlanSTag = &reuseVlanSTag
	return &this
}

// NewServiceProfileLinkProtocolConfigWithDefaults instantiates a new ServiceProfileLinkProtocolConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceProfileLinkProtocolConfigWithDefaults() *ServiceProfileLinkProtocolConfig {
	this := ServiceProfileLinkProtocolConfig{}
	var reuseVlanSTag bool = false
	this.ReuseVlanSTag = &reuseVlanSTag
	return &this
}

// GetEncapsulationStrategy returns the EncapsulationStrategy field value if set, zero value otherwise.
func (o *ServiceProfileLinkProtocolConfig) GetEncapsulationStrategy() ServiceProfileLinkProtocolConfigEncapsulationStrategy {
	if o == nil || IsNil(o.EncapsulationStrategy) {
		var ret ServiceProfileLinkProtocolConfigEncapsulationStrategy
		return ret
	}
	return *o.EncapsulationStrategy
}

// GetEncapsulationStrategyOk returns a tuple with the EncapsulationStrategy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceProfileLinkProtocolConfig) GetEncapsulationStrategyOk() (*ServiceProfileLinkProtocolConfigEncapsulationStrategy, bool) {
	if o == nil || IsNil(o.EncapsulationStrategy) {
		return nil, false
	}
	return o.EncapsulationStrategy, true
}

// HasEncapsulationStrategy returns a boolean if a field has been set.
func (o *ServiceProfileLinkProtocolConfig) HasEncapsulationStrategy() bool {
	if o != nil && !IsNil(o.EncapsulationStrategy) {
		return true
	}

	return false
}

// SetEncapsulationStrategy gets a reference to the given ServiceProfileLinkProtocolConfigEncapsulationStrategy and assigns it to the EncapsulationStrategy field.
func (o *ServiceProfileLinkProtocolConfig) SetEncapsulationStrategy(v ServiceProfileLinkProtocolConfigEncapsulationStrategy) {
	o.EncapsulationStrategy = &v
}

// GetNamedTags returns the NamedTags field value if set, zero value otherwise.
func (o *ServiceProfileLinkProtocolConfig) GetNamedTags() []string {
	if o == nil || IsNil(o.NamedTags) {
		var ret []string
		return ret
	}
	return o.NamedTags
}

// GetNamedTagsOk returns a tuple with the NamedTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceProfileLinkProtocolConfig) GetNamedTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.NamedTags) {
		return nil, false
	}
	return o.NamedTags, true
}

// HasNamedTags returns a boolean if a field has been set.
func (o *ServiceProfileLinkProtocolConfig) HasNamedTags() bool {
	if o != nil && !IsNil(o.NamedTags) {
		return true
	}

	return false
}

// SetNamedTags gets a reference to the given []string and assigns it to the NamedTags field.
func (o *ServiceProfileLinkProtocolConfig) SetNamedTags(v []string) {
	o.NamedTags = v
}

// GetVlanCTagLabel returns the VlanCTagLabel field value if set, zero value otherwise.
func (o *ServiceProfileLinkProtocolConfig) GetVlanCTagLabel() string {
	if o == nil || IsNil(o.VlanCTagLabel) {
		var ret string
		return ret
	}
	return *o.VlanCTagLabel
}

// GetVlanCTagLabelOk returns a tuple with the VlanCTagLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceProfileLinkProtocolConfig) GetVlanCTagLabelOk() (*string, bool) {
	if o == nil || IsNil(o.VlanCTagLabel) {
		return nil, false
	}
	return o.VlanCTagLabel, true
}

// HasVlanCTagLabel returns a boolean if a field has been set.
func (o *ServiceProfileLinkProtocolConfig) HasVlanCTagLabel() bool {
	if o != nil && !IsNil(o.VlanCTagLabel) {
		return true
	}

	return false
}

// SetVlanCTagLabel gets a reference to the given string and assigns it to the VlanCTagLabel field.
func (o *ServiceProfileLinkProtocolConfig) SetVlanCTagLabel(v string) {
	o.VlanCTagLabel = &v
}

// GetReuseVlanSTag returns the ReuseVlanSTag field value if set, zero value otherwise.
func (o *ServiceProfileLinkProtocolConfig) GetReuseVlanSTag() bool {
	if o == nil || IsNil(o.ReuseVlanSTag) {
		var ret bool
		return ret
	}
	return *o.ReuseVlanSTag
}

// GetReuseVlanSTagOk returns a tuple with the ReuseVlanSTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceProfileLinkProtocolConfig) GetReuseVlanSTagOk() (*bool, bool) {
	if o == nil || IsNil(o.ReuseVlanSTag) {
		return nil, false
	}
	return o.ReuseVlanSTag, true
}

// HasReuseVlanSTag returns a boolean if a field has been set.
func (o *ServiceProfileLinkProtocolConfig) HasReuseVlanSTag() bool {
	if o != nil && !IsNil(o.ReuseVlanSTag) {
		return true
	}

	return false
}

// SetReuseVlanSTag gets a reference to the given bool and assigns it to the ReuseVlanSTag field.
func (o *ServiceProfileLinkProtocolConfig) SetReuseVlanSTag(v bool) {
	o.ReuseVlanSTag = &v
}

// GetEncapsulation returns the Encapsulation field value if set, zero value otherwise.
func (o *ServiceProfileLinkProtocolConfig) GetEncapsulation() ServiceProfileLinkProtocolConfigEncapsulation {
	if o == nil || IsNil(o.Encapsulation) {
		var ret ServiceProfileLinkProtocolConfigEncapsulation
		return ret
	}
	return *o.Encapsulation
}

// GetEncapsulationOk returns a tuple with the Encapsulation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceProfileLinkProtocolConfig) GetEncapsulationOk() (*ServiceProfileLinkProtocolConfigEncapsulation, bool) {
	if o == nil || IsNil(o.Encapsulation) {
		return nil, false
	}
	return o.Encapsulation, true
}

// HasEncapsulation returns a boolean if a field has been set.
func (o *ServiceProfileLinkProtocolConfig) HasEncapsulation() bool {
	if o != nil && !IsNil(o.Encapsulation) {
		return true
	}

	return false
}

// SetEncapsulation gets a reference to the given ServiceProfileLinkProtocolConfigEncapsulation and assigns it to the Encapsulation field.
func (o *ServiceProfileLinkProtocolConfig) SetEncapsulation(v ServiceProfileLinkProtocolConfigEncapsulation) {
	o.Encapsulation = &v
}

func (o ServiceProfileLinkProtocolConfig) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceProfileLinkProtocolConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EncapsulationStrategy) {
		toSerialize["encapsulationStrategy"] = o.EncapsulationStrategy
	}
	if !IsNil(o.NamedTags) {
		toSerialize["namedTags"] = o.NamedTags
	}
	if !IsNil(o.VlanCTagLabel) {
		toSerialize["vlanCTagLabel"] = o.VlanCTagLabel
	}
	if !IsNil(o.ReuseVlanSTag) {
		toSerialize["reuseVlanSTag"] = o.ReuseVlanSTag
	}
	if !IsNil(o.Encapsulation) {
		toSerialize["encapsulation"] = o.Encapsulation
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ServiceProfileLinkProtocolConfig) UnmarshalJSON(data []byte) (err error) {
	varServiceProfileLinkProtocolConfig := _ServiceProfileLinkProtocolConfig{}

	err = json.Unmarshal(data, &varServiceProfileLinkProtocolConfig)

	if err != nil {
		return err
	}

	*o = ServiceProfileLinkProtocolConfig(varServiceProfileLinkProtocolConfig)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "encapsulationStrategy")
		delete(additionalProperties, "namedTags")
		delete(additionalProperties, "vlanCTagLabel")
		delete(additionalProperties, "reuseVlanSTag")
		delete(additionalProperties, "encapsulation")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableServiceProfileLinkProtocolConfig struct {
	value *ServiceProfileLinkProtocolConfig
	isSet bool
}

func (v NullableServiceProfileLinkProtocolConfig) Get() *ServiceProfileLinkProtocolConfig {
	return v.value
}

func (v *NullableServiceProfileLinkProtocolConfig) Set(val *ServiceProfileLinkProtocolConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceProfileLinkProtocolConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceProfileLinkProtocolConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceProfileLinkProtocolConfig(val *ServiceProfileLinkProtocolConfig) *NullableServiceProfileLinkProtocolConfig {
	return &NullableServiceProfileLinkProtocolConfig{value: val, isSet: true}
}

func (v NullableServiceProfileLinkProtocolConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceProfileLinkProtocolConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
