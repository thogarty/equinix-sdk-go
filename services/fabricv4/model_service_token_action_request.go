/*
Equinix Fabric API v4

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fabricv4

import (
	"encoding/json"
	"fmt"
)

// checks if the ServiceTokenActionRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceTokenActionRequest{}

// ServiceTokenActionRequest Service Token action request
type ServiceTokenActionRequest struct {
	Type                 ServiceTokenActions `json:"type"`
	AdditionalProperties map[string]interface{}
}

type _ServiceTokenActionRequest ServiceTokenActionRequest

// NewServiceTokenActionRequest instantiates a new ServiceTokenActionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceTokenActionRequest(type_ ServiceTokenActions) *ServiceTokenActionRequest {
	this := ServiceTokenActionRequest{}
	this.Type = type_
	return &this
}

// NewServiceTokenActionRequestWithDefaults instantiates a new ServiceTokenActionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceTokenActionRequestWithDefaults() *ServiceTokenActionRequest {
	this := ServiceTokenActionRequest{}
	return &this
}

// GetType returns the Type field value
func (o *ServiceTokenActionRequest) GetType() ServiceTokenActions {
	if o == nil {
		var ret ServiceTokenActions
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ServiceTokenActionRequest) GetTypeOk() (*ServiceTokenActions, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ServiceTokenActionRequest) SetType(v ServiceTokenActions) {
	o.Type = v
}

func (o ServiceTokenActionRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceTokenActionRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ServiceTokenActionRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
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

	varServiceTokenActionRequest := _ServiceTokenActionRequest{}

	err = json.Unmarshal(data, &varServiceTokenActionRequest)

	if err != nil {
		return err
	}

	*o = ServiceTokenActionRequest(varServiceTokenActionRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableServiceTokenActionRequest struct {
	value *ServiceTokenActionRequest
	isSet bool
}

func (v NullableServiceTokenActionRequest) Get() *ServiceTokenActionRequest {
	return v.value
}

func (v *NullableServiceTokenActionRequest) Set(val *ServiceTokenActionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceTokenActionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceTokenActionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceTokenActionRequest(val *ServiceTokenActionRequest) *NullableServiceTokenActionRequest {
	return &NullableServiceTokenActionRequest{value: val, isSet: true}
}

func (v NullableServiceTokenActionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceTokenActionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
