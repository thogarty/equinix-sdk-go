/*
Equinix Internet Access API

Equinix Internet Access provides direct access to the Internet with scalable bandwidth options in IBX data centers.

Contact: api-support@equinix.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eiav2

import (
	"encoding/json"
	"fmt"
)

// checks if the ServiceV2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceV2{}

// ServiceV2 struct for ServiceV2
type ServiceV2 struct {
	Uuid string         `json:"uuid"`
	Type *ServiceTypeV2 `json:"type,omitempty"`
	// Service bandwidth in Mbps
	Bandwidth            *int32                  `json:"bandwidth,omitempty"`
	Account              *CustomerBillingAccount `json:"account,omitempty"`
	ChangeLog            *ServiceChangeLog       `json:"changeLog,omitempty"`
	Links                []Link                  `json:"links,omitempty"`
	Order                *ServiceOrderReference  `json:"order,omitempty"`
	Project              *ProjectReference       `json:"project,omitempty"`
	State                ServiceState            `json:"state"`
	AdditionalProperties map[string]interface{}
}

type _ServiceV2 ServiceV2

// NewServiceV2 instantiates a new ServiceV2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceV2(uuid string, state ServiceState) *ServiceV2 {
	this := ServiceV2{}
	this.Uuid = uuid
	this.State = state
	return &this
}

// NewServiceV2WithDefaults instantiates a new ServiceV2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceV2WithDefaults() *ServiceV2 {
	this := ServiceV2{}
	return &this
}

// GetUuid returns the Uuid field value
func (o *ServiceV2) GetUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value
// and a boolean to check if the value has been set.
func (o *ServiceV2) GetUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uuid, true
}

// SetUuid sets field value
func (o *ServiceV2) SetUuid(v string) {
	o.Uuid = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ServiceV2) GetType() ServiceTypeV2 {
	if o == nil || IsNil(o.Type) {
		var ret ServiceTypeV2
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceV2) GetTypeOk() (*ServiceTypeV2, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ServiceV2) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given ServiceTypeV2 and assigns it to the Type field.
func (o *ServiceV2) SetType(v ServiceTypeV2) {
	o.Type = &v
}

// GetBandwidth returns the Bandwidth field value if set, zero value otherwise.
func (o *ServiceV2) GetBandwidth() int32 {
	if o == nil || IsNil(o.Bandwidth) {
		var ret int32
		return ret
	}
	return *o.Bandwidth
}

// GetBandwidthOk returns a tuple with the Bandwidth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceV2) GetBandwidthOk() (*int32, bool) {
	if o == nil || IsNil(o.Bandwidth) {
		return nil, false
	}
	return o.Bandwidth, true
}

// HasBandwidth returns a boolean if a field has been set.
func (o *ServiceV2) HasBandwidth() bool {
	if o != nil && !IsNil(o.Bandwidth) {
		return true
	}

	return false
}

// SetBandwidth gets a reference to the given int32 and assigns it to the Bandwidth field.
func (o *ServiceV2) SetBandwidth(v int32) {
	o.Bandwidth = &v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *ServiceV2) GetAccount() CustomerBillingAccount {
	if o == nil || IsNil(o.Account) {
		var ret CustomerBillingAccount
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceV2) GetAccountOk() (*CustomerBillingAccount, bool) {
	if o == nil || IsNil(o.Account) {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *ServiceV2) HasAccount() bool {
	if o != nil && !IsNil(o.Account) {
		return true
	}

	return false
}

// SetAccount gets a reference to the given CustomerBillingAccount and assigns it to the Account field.
func (o *ServiceV2) SetAccount(v CustomerBillingAccount) {
	o.Account = &v
}

// GetChangeLog returns the ChangeLog field value if set, zero value otherwise.
func (o *ServiceV2) GetChangeLog() ServiceChangeLog {
	if o == nil || IsNil(o.ChangeLog) {
		var ret ServiceChangeLog
		return ret
	}
	return *o.ChangeLog
}

// GetChangeLogOk returns a tuple with the ChangeLog field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceV2) GetChangeLogOk() (*ServiceChangeLog, bool) {
	if o == nil || IsNil(o.ChangeLog) {
		return nil, false
	}
	return o.ChangeLog, true
}

// HasChangeLog returns a boolean if a field has been set.
func (o *ServiceV2) HasChangeLog() bool {
	if o != nil && !IsNil(o.ChangeLog) {
		return true
	}

	return false
}

// SetChangeLog gets a reference to the given ServiceChangeLog and assigns it to the ChangeLog field.
func (o *ServiceV2) SetChangeLog(v ServiceChangeLog) {
	o.ChangeLog = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ServiceV2) GetLinks() []Link {
	if o == nil || IsNil(o.Links) {
		var ret []Link
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceV2) GetLinksOk() ([]Link, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ServiceV2) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []Link and assigns it to the Links field.
func (o *ServiceV2) SetLinks(v []Link) {
	o.Links = v
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *ServiceV2) GetOrder() ServiceOrderReference {
	if o == nil || IsNil(o.Order) {
		var ret ServiceOrderReference
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceV2) GetOrderOk() (*ServiceOrderReference, bool) {
	if o == nil || IsNil(o.Order) {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *ServiceV2) HasOrder() bool {
	if o != nil && !IsNil(o.Order) {
		return true
	}

	return false
}

// SetOrder gets a reference to the given ServiceOrderReference and assigns it to the Order field.
func (o *ServiceV2) SetOrder(v ServiceOrderReference) {
	o.Order = &v
}

// GetProject returns the Project field value if set, zero value otherwise.
func (o *ServiceV2) GetProject() ProjectReference {
	if o == nil || IsNil(o.Project) {
		var ret ProjectReference
		return ret
	}
	return *o.Project
}

// GetProjectOk returns a tuple with the Project field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceV2) GetProjectOk() (*ProjectReference, bool) {
	if o == nil || IsNil(o.Project) {
		return nil, false
	}
	return o.Project, true
}

// HasProject returns a boolean if a field has been set.
func (o *ServiceV2) HasProject() bool {
	if o != nil && !IsNil(o.Project) {
		return true
	}

	return false
}

// SetProject gets a reference to the given ProjectReference and assigns it to the Project field.
func (o *ServiceV2) SetProject(v ProjectReference) {
	o.Project = &v
}

// GetState returns the State field value
func (o *ServiceV2) GetState() ServiceState {
	if o == nil {
		var ret ServiceState
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *ServiceV2) GetStateOk() (*ServiceState, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *ServiceV2) SetState(v ServiceState) {
	o.State = v
}

func (o ServiceV2) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceV2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["uuid"] = o.Uuid
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Bandwidth) {
		toSerialize["bandwidth"] = o.Bandwidth
	}
	if !IsNil(o.Account) {
		toSerialize["account"] = o.Account
	}
	if !IsNil(o.ChangeLog) {
		toSerialize["changeLog"] = o.ChangeLog
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Order) {
		toSerialize["order"] = o.Order
	}
	if !IsNil(o.Project) {
		toSerialize["project"] = o.Project
	}
	toSerialize["state"] = o.State

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ServiceV2) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"uuid",
		"state",
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

	varServiceV2 := _ServiceV2{}

	err = json.Unmarshal(data, &varServiceV2)

	if err != nil {
		return err
	}

	*o = ServiceV2(varServiceV2)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "uuid")
		delete(additionalProperties, "type")
		delete(additionalProperties, "bandwidth")
		delete(additionalProperties, "account")
		delete(additionalProperties, "changeLog")
		delete(additionalProperties, "links")
		delete(additionalProperties, "order")
		delete(additionalProperties, "project")
		delete(additionalProperties, "state")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableServiceV2 struct {
	value *ServiceV2
	isSet bool
}

func (v NullableServiceV2) Get() *ServiceV2 {
	return v.value
}

func (v *NullableServiceV2) Set(val *ServiceV2) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceV2) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceV2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceV2(val *ServiceV2) *NullableServiceV2 {
	return &NullableServiceV2{value: val, isSet: true}
}

func (v NullableServiceV2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceV2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
