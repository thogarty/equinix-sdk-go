/*
Equinix Fabric API v4

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fabricv4

import (
	"encoding/json"
)

// checks if the ConnectionInvitation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConnectionInvitation{}

// ConnectionInvitation Connection Invitation Details
type ConnectionInvitation struct {
	// invitee email
	Email *string `json:"email,omitempty"`
	// invitation message
	Message *string `json:"message,omitempty"`
	// draft order id for invitation
	CtrDraftOrderId      *string `json:"ctrDraftOrderId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ConnectionInvitation ConnectionInvitation

// NewConnectionInvitation instantiates a new ConnectionInvitation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectionInvitation() *ConnectionInvitation {
	this := ConnectionInvitation{}
	return &this
}

// NewConnectionInvitationWithDefaults instantiates a new ConnectionInvitation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectionInvitationWithDefaults() *ConnectionInvitation {
	this := ConnectionInvitation{}
	return &this
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *ConnectionInvitation) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionInvitation) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *ConnectionInvitation) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *ConnectionInvitation) SetEmail(v string) {
	o.Email = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ConnectionInvitation) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionInvitation) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ConnectionInvitation) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ConnectionInvitation) SetMessage(v string) {
	o.Message = &v
}

// GetCtrDraftOrderId returns the CtrDraftOrderId field value if set, zero value otherwise.
func (o *ConnectionInvitation) GetCtrDraftOrderId() string {
	if o == nil || IsNil(o.CtrDraftOrderId) {
		var ret string
		return ret
	}
	return *o.CtrDraftOrderId
}

// GetCtrDraftOrderIdOk returns a tuple with the CtrDraftOrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionInvitation) GetCtrDraftOrderIdOk() (*string, bool) {
	if o == nil || IsNil(o.CtrDraftOrderId) {
		return nil, false
	}
	return o.CtrDraftOrderId, true
}

// HasCtrDraftOrderId returns a boolean if a field has been set.
func (o *ConnectionInvitation) HasCtrDraftOrderId() bool {
	if o != nil && !IsNil(o.CtrDraftOrderId) {
		return true
	}

	return false
}

// SetCtrDraftOrderId gets a reference to the given string and assigns it to the CtrDraftOrderId field.
func (o *ConnectionInvitation) SetCtrDraftOrderId(v string) {
	o.CtrDraftOrderId = &v
}

func (o ConnectionInvitation) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConnectionInvitation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.CtrDraftOrderId) {
		toSerialize["ctrDraftOrderId"] = o.CtrDraftOrderId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ConnectionInvitation) UnmarshalJSON(data []byte) (err error) {
	varConnectionInvitation := _ConnectionInvitation{}

	err = json.Unmarshal(data, &varConnectionInvitation)

	if err != nil {
		return err
	}

	*o = ConnectionInvitation(varConnectionInvitation)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "email")
		delete(additionalProperties, "message")
		delete(additionalProperties, "ctrDraftOrderId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableConnectionInvitation struct {
	value *ConnectionInvitation
	isSet bool
}

func (v NullableConnectionInvitation) Get() *ConnectionInvitation {
	return v.value
}

func (v *NullableConnectionInvitation) Set(val *ConnectionInvitation) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectionInvitation) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectionInvitation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectionInvitation(val *ConnectionInvitation) *NullableConnectionInvitation {
	return &NullableConnectionInvitation{value: val, isSet: true}
}

func (v NullableConnectionInvitation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectionInvitation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
