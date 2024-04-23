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

// checks if the ConnectionAction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConnectionAction{}

// ConnectionAction Connection action
type ConnectionAction struct {
	Type Actions `json:"type"`
	// Connection action URI
	Href string `json:"href"`
	// Equinix-assigned connection identifier
	Uuid string `json:"uuid"`
	// Connection rejection reason detail
	Description          *string                  `json:"description,omitempty"`
	Data                 ConnectionAcceptanceData `json:"data"`
	ChangeLog            *Changelog               `json:"changeLog,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ConnectionAction ConnectionAction

// NewConnectionAction instantiates a new ConnectionAction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectionAction(type_ Actions, href string, uuid string, data ConnectionAcceptanceData) *ConnectionAction {
	this := ConnectionAction{}
	this.Type = type_
	this.Href = href
	this.Uuid = uuid
	this.Data = data
	return &this
}

// NewConnectionActionWithDefaults instantiates a new ConnectionAction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectionActionWithDefaults() *ConnectionAction {
	this := ConnectionAction{}
	return &this
}

// GetType returns the Type field value
func (o *ConnectionAction) GetType() Actions {
	if o == nil {
		var ret Actions
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ConnectionAction) GetTypeOk() (*Actions, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ConnectionAction) SetType(v Actions) {
	o.Type = v
}

// GetHref returns the Href field value
func (o *ConnectionAction) GetHref() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Href
}

// GetHrefOk returns a tuple with the Href field value
// and a boolean to check if the value has been set.
func (o *ConnectionAction) GetHrefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Href, true
}

// SetHref sets field value
func (o *ConnectionAction) SetHref(v string) {
	o.Href = v
}

// GetUuid returns the Uuid field value
func (o *ConnectionAction) GetUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value
// and a boolean to check if the value has been set.
func (o *ConnectionAction) GetUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uuid, true
}

// SetUuid sets field value
func (o *ConnectionAction) SetUuid(v string) {
	o.Uuid = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ConnectionAction) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionAction) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ConnectionAction) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ConnectionAction) SetDescription(v string) {
	o.Description = &v
}

// GetData returns the Data field value
func (o *ConnectionAction) GetData() ConnectionAcceptanceData {
	if o == nil {
		var ret ConnectionAcceptanceData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ConnectionAction) GetDataOk() (*ConnectionAcceptanceData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *ConnectionAction) SetData(v ConnectionAcceptanceData) {
	o.Data = v
}

// GetChangeLog returns the ChangeLog field value if set, zero value otherwise.
func (o *ConnectionAction) GetChangeLog() Changelog {
	if o == nil || IsNil(o.ChangeLog) {
		var ret Changelog
		return ret
	}
	return *o.ChangeLog
}

// GetChangeLogOk returns a tuple with the ChangeLog field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionAction) GetChangeLogOk() (*Changelog, bool) {
	if o == nil || IsNil(o.ChangeLog) {
		return nil, false
	}
	return o.ChangeLog, true
}

// HasChangeLog returns a boolean if a field has been set.
func (o *ConnectionAction) HasChangeLog() bool {
	if o != nil && !IsNil(o.ChangeLog) {
		return true
	}

	return false
}

// SetChangeLog gets a reference to the given Changelog and assigns it to the ChangeLog field.
func (o *ConnectionAction) SetChangeLog(v Changelog) {
	o.ChangeLog = &v
}

func (o ConnectionAction) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConnectionAction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["href"] = o.Href
	toSerialize["uuid"] = o.Uuid
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["data"] = o.Data
	if !IsNil(o.ChangeLog) {
		toSerialize["changeLog"] = o.ChangeLog
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ConnectionAction) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"href",
		"uuid",
		"data",
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

	varConnectionAction := _ConnectionAction{}

	err = json.Unmarshal(data, &varConnectionAction)

	if err != nil {
		return err
	}

	*o = ConnectionAction(varConnectionAction)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "href")
		delete(additionalProperties, "uuid")
		delete(additionalProperties, "description")
		delete(additionalProperties, "data")
		delete(additionalProperties, "changeLog")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableConnectionAction struct {
	value *ConnectionAction
	isSet bool
}

func (v NullableConnectionAction) Get() *ConnectionAction {
	return v.value
}

func (v *NullableConnectionAction) Set(val *ConnectionAction) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectionAction) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectionAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectionAction(val *ConnectionAction) *NullableConnectionAction {
	return &NullableConnectionAction{value: val, isSet: true}
}

func (v NullableConnectionAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectionAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
