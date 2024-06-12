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

// checks if the SimplifiedNotification type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SimplifiedNotification{}

// SimplifiedNotification struct for SimplifiedNotification
type SimplifiedNotification struct {
	Type         SimplifiedNotificationType `json:"type"`
	SendInterval *string                    `json:"sendInterval,omitempty"`
	// Array of contact emails
	Emails []string `json:"emails"`
	// Array of registered users
	RegisteredUsers      []string `json:"registeredUsers,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SimplifiedNotification SimplifiedNotification

// NewSimplifiedNotification instantiates a new SimplifiedNotification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSimplifiedNotification(type_ SimplifiedNotificationType, emails []string) *SimplifiedNotification {
	this := SimplifiedNotification{}
	this.Type = type_
	this.Emails = emails
	return &this
}

// NewSimplifiedNotificationWithDefaults instantiates a new SimplifiedNotification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSimplifiedNotificationWithDefaults() *SimplifiedNotification {
	this := SimplifiedNotification{}
	return &this
}

// GetType returns the Type field value
func (o *SimplifiedNotification) GetType() SimplifiedNotificationType {
	if o == nil {
		var ret SimplifiedNotificationType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SimplifiedNotification) GetTypeOk() (*SimplifiedNotificationType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SimplifiedNotification) SetType(v SimplifiedNotificationType) {
	o.Type = v
}

// GetSendInterval returns the SendInterval field value if set, zero value otherwise.
func (o *SimplifiedNotification) GetSendInterval() string {
	if o == nil || IsNil(o.SendInterval) {
		var ret string
		return ret
	}
	return *o.SendInterval
}

// GetSendIntervalOk returns a tuple with the SendInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimplifiedNotification) GetSendIntervalOk() (*string, bool) {
	if o == nil || IsNil(o.SendInterval) {
		return nil, false
	}
	return o.SendInterval, true
}

// HasSendInterval returns a boolean if a field has been set.
func (o *SimplifiedNotification) HasSendInterval() bool {
	if o != nil && !IsNil(o.SendInterval) {
		return true
	}

	return false
}

// SetSendInterval gets a reference to the given string and assigns it to the SendInterval field.
func (o *SimplifiedNotification) SetSendInterval(v string) {
	o.SendInterval = &v
}

// GetEmails returns the Emails field value
func (o *SimplifiedNotification) GetEmails() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Emails
}

// GetEmailsOk returns a tuple with the Emails field value
// and a boolean to check if the value has been set.
func (o *SimplifiedNotification) GetEmailsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Emails, true
}

// SetEmails sets field value
func (o *SimplifiedNotification) SetEmails(v []string) {
	o.Emails = v
}

// GetRegisteredUsers returns the RegisteredUsers field value if set, zero value otherwise.
func (o *SimplifiedNotification) GetRegisteredUsers() []string {
	if o == nil || IsNil(o.RegisteredUsers) {
		var ret []string
		return ret
	}
	return o.RegisteredUsers
}

// GetRegisteredUsersOk returns a tuple with the RegisteredUsers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimplifiedNotification) GetRegisteredUsersOk() ([]string, bool) {
	if o == nil || IsNil(o.RegisteredUsers) {
		return nil, false
	}
	return o.RegisteredUsers, true
}

// HasRegisteredUsers returns a boolean if a field has been set.
func (o *SimplifiedNotification) HasRegisteredUsers() bool {
	if o != nil && !IsNil(o.RegisteredUsers) {
		return true
	}

	return false
}

// SetRegisteredUsers gets a reference to the given []string and assigns it to the RegisteredUsers field.
func (o *SimplifiedNotification) SetRegisteredUsers(v []string) {
	o.RegisteredUsers = v
}

func (o SimplifiedNotification) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SimplifiedNotification) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	if !IsNil(o.SendInterval) {
		toSerialize["sendInterval"] = o.SendInterval
	}
	toSerialize["emails"] = o.Emails
	if !IsNil(o.RegisteredUsers) {
		toSerialize["registeredUsers"] = o.RegisteredUsers
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SimplifiedNotification) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"emails",
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

	varSimplifiedNotification := _SimplifiedNotification{}

	err = json.Unmarshal(data, &varSimplifiedNotification)

	if err != nil {
		return err
	}

	*o = SimplifiedNotification(varSimplifiedNotification)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "sendInterval")
		delete(additionalProperties, "emails")
		delete(additionalProperties, "registeredUsers")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSimplifiedNotification struct {
	value *SimplifiedNotification
	isSet bool
}

func (v NullableSimplifiedNotification) Get() *SimplifiedNotification {
	return v.value
}

func (v *NullableSimplifiedNotification) Set(val *SimplifiedNotification) {
	v.value = val
	v.isSet = true
}

func (v NullableSimplifiedNotification) IsSet() bool {
	return v.isSet
}

func (v *NullableSimplifiedNotification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSimplifiedNotification(val *SimplifiedNotification) *NullableSimplifiedNotification {
	return &NullableSimplifiedNotification{value: val, isSet: true}
}

func (v NullableSimplifiedNotification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSimplifiedNotification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
