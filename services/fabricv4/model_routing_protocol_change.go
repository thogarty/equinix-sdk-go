/*
Equinix Fabric API v4

Contact: api-support@equinix.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fabricv4

import (
	"encoding/json"
	"fmt"
)

// checks if the RoutingProtocolChange type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RoutingProtocolChange{}

// RoutingProtocolChange Current state of latest Routing Protocol change
type RoutingProtocolChange struct {
	// Uniquely identifies a change
	Uuid string                    `json:"uuid"`
	Type RoutingProtocolChangeType `json:"type"`
	// Routing Protocol Change URI
	Href                 *string `json:"href,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RoutingProtocolChange RoutingProtocolChange

// NewRoutingProtocolChange instantiates a new RoutingProtocolChange object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoutingProtocolChange(uuid string, type_ RoutingProtocolChangeType) *RoutingProtocolChange {
	this := RoutingProtocolChange{}
	this.Uuid = uuid
	this.Type = type_
	return &this
}

// NewRoutingProtocolChangeWithDefaults instantiates a new RoutingProtocolChange object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoutingProtocolChangeWithDefaults() *RoutingProtocolChange {
	this := RoutingProtocolChange{}
	return &this
}

// GetUuid returns the Uuid field value
func (o *RoutingProtocolChange) GetUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value
// and a boolean to check if the value has been set.
func (o *RoutingProtocolChange) GetUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uuid, true
}

// SetUuid sets field value
func (o *RoutingProtocolChange) SetUuid(v string) {
	o.Uuid = v
}

// GetType returns the Type field value
func (o *RoutingProtocolChange) GetType() RoutingProtocolChangeType {
	if o == nil {
		var ret RoutingProtocolChangeType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *RoutingProtocolChange) GetTypeOk() (*RoutingProtocolChangeType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *RoutingProtocolChange) SetType(v RoutingProtocolChangeType) {
	o.Type = v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *RoutingProtocolChange) GetHref() string {
	if o == nil || IsNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoutingProtocolChange) GetHrefOk() (*string, bool) {
	if o == nil || IsNil(o.Href) {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *RoutingProtocolChange) HasHref() bool {
	if o != nil && !IsNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *RoutingProtocolChange) SetHref(v string) {
	o.Href = &v
}

func (o RoutingProtocolChange) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RoutingProtocolChange) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["uuid"] = o.Uuid
	toSerialize["type"] = o.Type
	if !IsNil(o.Href) {
		toSerialize["href"] = o.Href
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RoutingProtocolChange) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"uuid",
		"type",
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

	varRoutingProtocolChange := _RoutingProtocolChange{}

	err = json.Unmarshal(data, &varRoutingProtocolChange)

	if err != nil {
		return err
	}

	*o = RoutingProtocolChange(varRoutingProtocolChange)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "uuid")
		delete(additionalProperties, "type")
		delete(additionalProperties, "href")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRoutingProtocolChange struct {
	value *RoutingProtocolChange
	isSet bool
}

func (v NullableRoutingProtocolChange) Get() *RoutingProtocolChange {
	return v.value
}

func (v *NullableRoutingProtocolChange) Set(val *RoutingProtocolChange) {
	v.value = val
	v.isSet = true
}

func (v NullableRoutingProtocolChange) IsSet() bool {
	return v.isSet
}

func (v *NullableRoutingProtocolChange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoutingProtocolChange(val *RoutingProtocolChange) *NullableRoutingProtocolChange {
	return &NullableRoutingProtocolChange{value: val, isSet: true}
}

func (v NullableRoutingProtocolChange) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoutingProtocolChange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
