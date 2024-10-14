/*
Equinix Fabric API v4

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fabricv4

import (
	"encoding/json"
)

// checks if the PortRedundancy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PortRedundancy{}

// PortRedundancy Port redundancy configuration
type PortRedundancy struct {
	// Access point redundancy
	Enabled *bool `json:"enabled,omitempty"`
	// Port UUID of respective primary port
	// Deprecated
	Group                *string       `json:"group,omitempty"`
	Priority             *PortPriority `json:"priority,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PortRedundancy PortRedundancy

// NewPortRedundancy instantiates a new PortRedundancy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPortRedundancy() *PortRedundancy {
	this := PortRedundancy{}
	return &this
}

// NewPortRedundancyWithDefaults instantiates a new PortRedundancy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPortRedundancyWithDefaults() *PortRedundancy {
	this := PortRedundancy{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *PortRedundancy) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortRedundancy) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *PortRedundancy) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *PortRedundancy) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetGroup returns the Group field value if set, zero value otherwise.
// Deprecated
func (o *PortRedundancy) GetGroup() string {
	if o == nil || IsNil(o.Group) {
		var ret string
		return ret
	}
	return *o.Group
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *PortRedundancy) GetGroupOk() (*string, bool) {
	if o == nil || IsNil(o.Group) {
		return nil, false
	}
	return o.Group, true
}

// HasGroup returns a boolean if a field has been set.
func (o *PortRedundancy) HasGroup() bool {
	if o != nil && !IsNil(o.Group) {
		return true
	}

	return false
}

// SetGroup gets a reference to the given string and assigns it to the Group field.
// Deprecated
func (o *PortRedundancy) SetGroup(v string) {
	o.Group = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *PortRedundancy) GetPriority() PortPriority {
	if o == nil || IsNil(o.Priority) {
		var ret PortPriority
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortRedundancy) GetPriorityOk() (*PortPriority, bool) {
	if o == nil || IsNil(o.Priority) {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *PortRedundancy) HasPriority() bool {
	if o != nil && !IsNil(o.Priority) {
		return true
	}

	return false
}

// SetPriority gets a reference to the given PortPriority and assigns it to the Priority field.
func (o *PortRedundancy) SetPriority(v PortPriority) {
	o.Priority = &v
}

func (o PortRedundancy) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PortRedundancy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.Group) {
		toSerialize["group"] = o.Group
	}
	if !IsNil(o.Priority) {
		toSerialize["priority"] = o.Priority
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PortRedundancy) UnmarshalJSON(data []byte) (err error) {
	varPortRedundancy := _PortRedundancy{}

	err = json.Unmarshal(data, &varPortRedundancy)

	if err != nil {
		return err
	}

	*o = PortRedundancy(varPortRedundancy)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "enabled")
		delete(additionalProperties, "group")
		delete(additionalProperties, "priority")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePortRedundancy struct {
	value *PortRedundancy
	isSet bool
}

func (v NullablePortRedundancy) Get() *PortRedundancy {
	return v.value
}

func (v *NullablePortRedundancy) Set(val *PortRedundancy) {
	v.value = val
	v.isSet = true
}

func (v NullablePortRedundancy) IsSet() bool {
	return v.isSet
}

func (v *NullablePortRedundancy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePortRedundancy(val *PortRedundancy) *NullablePortRedundancy {
	return &NullablePortRedundancy{value: val, isSet: true}
}

func (v NullablePortRedundancy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePortRedundancy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
