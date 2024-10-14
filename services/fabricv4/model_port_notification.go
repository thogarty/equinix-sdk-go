/*
Equinix Fabric API v4

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fabricv4

import (
	"encoding/json"
	"fmt"
)

// checks if the PortNotification type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PortNotification{}

// PortNotification struct for PortNotification
type PortNotification struct {
	Type PortNotificationType `json:"type"`
	// Array of registered users
	RegisteredUsers      []string `json:"registeredUsers"`
	AdditionalProperties map[string]interface{}
}

type _PortNotification PortNotification

// NewPortNotification instantiates a new PortNotification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPortNotification(type_ PortNotificationType, registeredUsers []string) *PortNotification {
	this := PortNotification{}
	this.Type = type_
	this.RegisteredUsers = registeredUsers
	return &this
}

// NewPortNotificationWithDefaults instantiates a new PortNotification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPortNotificationWithDefaults() *PortNotification {
	this := PortNotification{}
	return &this
}

// GetType returns the Type field value
func (o *PortNotification) GetType() PortNotificationType {
	if o == nil {
		var ret PortNotificationType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *PortNotification) GetTypeOk() (*PortNotificationType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PortNotification) SetType(v PortNotificationType) {
	o.Type = v
}

// GetRegisteredUsers returns the RegisteredUsers field value
func (o *PortNotification) GetRegisteredUsers() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.RegisteredUsers
}

// GetRegisteredUsersOk returns a tuple with the RegisteredUsers field value
// and a boolean to check if the value has been set.
func (o *PortNotification) GetRegisteredUsersOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RegisteredUsers, true
}

// SetRegisteredUsers sets field value
func (o *PortNotification) SetRegisteredUsers(v []string) {
	o.RegisteredUsers = v
}

func (o PortNotification) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PortNotification) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["registeredUsers"] = o.RegisteredUsers

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PortNotification) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"registeredUsers",
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

	varPortNotification := _PortNotification{}

	err = json.Unmarshal(data, &varPortNotification)

	if err != nil {
		return err
	}

	*o = PortNotification(varPortNotification)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "registeredUsers")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePortNotification struct {
	value *PortNotification
	isSet bool
}

func (v NullablePortNotification) Get() *PortNotification {
	return v.value
}

func (v *NullablePortNotification) Set(val *PortNotification) {
	v.value = val
	v.isSet = true
}

func (v NullablePortNotification) IsSet() bool {
	return v.isSet
}

func (v *NullablePortNotification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePortNotification(val *PortNotification) *NullablePortNotification {
	return &NullablePortNotification{value: val, isSet: true}
}

func (v NullablePortNotification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePortNotification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
