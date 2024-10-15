/*
Equinix Fabric API v4

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fabricv4

import (
	"encoding/json"
)

// checks if the ConnectionRouteTableEntryConnection type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConnectionRouteTableEntryConnection{}

// ConnectionRouteTableEntryConnection struct for ConnectionRouteTableEntryConnection
type ConnectionRouteTableEntryConnection struct {
	Uuid                 *string `json:"uuid,omitempty"`
	Name                 *string `json:"name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ConnectionRouteTableEntryConnection ConnectionRouteTableEntryConnection

// NewConnectionRouteTableEntryConnection instantiates a new ConnectionRouteTableEntryConnection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectionRouteTableEntryConnection() *ConnectionRouteTableEntryConnection {
	this := ConnectionRouteTableEntryConnection{}
	return &this
}

// NewConnectionRouteTableEntryConnectionWithDefaults instantiates a new ConnectionRouteTableEntryConnection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectionRouteTableEntryConnectionWithDefaults() *ConnectionRouteTableEntryConnection {
	this := ConnectionRouteTableEntryConnection{}
	return &this
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *ConnectionRouteTableEntryConnection) GetUuid() string {
	if o == nil || IsNil(o.Uuid) {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionRouteTableEntryConnection) GetUuidOk() (*string, bool) {
	if o == nil || IsNil(o.Uuid) {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *ConnectionRouteTableEntryConnection) HasUuid() bool {
	if o != nil && !IsNil(o.Uuid) {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *ConnectionRouteTableEntryConnection) SetUuid(v string) {
	o.Uuid = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ConnectionRouteTableEntryConnection) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionRouteTableEntryConnection) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ConnectionRouteTableEntryConnection) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ConnectionRouteTableEntryConnection) SetName(v string) {
	o.Name = &v
}

func (o ConnectionRouteTableEntryConnection) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConnectionRouteTableEntryConnection) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Uuid) {
		toSerialize["uuid"] = o.Uuid
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ConnectionRouteTableEntryConnection) UnmarshalJSON(data []byte) (err error) {
	varConnectionRouteTableEntryConnection := _ConnectionRouteTableEntryConnection{}

	err = json.Unmarshal(data, &varConnectionRouteTableEntryConnection)

	if err != nil {
		return err
	}

	*o = ConnectionRouteTableEntryConnection(varConnectionRouteTableEntryConnection)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "uuid")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableConnectionRouteTableEntryConnection struct {
	value *ConnectionRouteTableEntryConnection
	isSet bool
}

func (v NullableConnectionRouteTableEntryConnection) Get() *ConnectionRouteTableEntryConnection {
	return v.value
}

func (v *NullableConnectionRouteTableEntryConnection) Set(val *ConnectionRouteTableEntryConnection) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectionRouteTableEntryConnection) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectionRouteTableEntryConnection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectionRouteTableEntryConnection(val *ConnectionRouteTableEntryConnection) *NullableConnectionRouteTableEntryConnection {
	return &NullableConnectionRouteTableEntryConnection{value: val, isSet: true}
}

func (v NullableConnectionRouteTableEntryConnection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectionRouteTableEntryConnection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}