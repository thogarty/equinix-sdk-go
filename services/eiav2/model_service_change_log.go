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
	"time"
)

// checks if the ServiceChangeLog type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceChangeLog{}

// ServiceChangeLog struct for ServiceChangeLog
type ServiceChangeLog struct {
	CreatedBy            string     `json:"createdBy"`
	CreatedByFullName    string     `json:"createdByFullName"`
	CreatedByEmail       string     `json:"createdByEmail"`
	CreatedDateTime      time.Time  `json:"createdDateTime"`
	UpdatedBy            string     `json:"updatedBy"`
	UpdatedByFullName    string     `json:"updatedByFullName"`
	UpdatedByEmail       string     `json:"updatedByEmail"`
	UpdatedDateTime      time.Time  `json:"updatedDateTime"`
	DeletedBy            *string    `json:"deletedBy,omitempty"`
	DeletedByFullName    *string    `json:"deletedByFullName,omitempty"`
	DeletedByEmail       *string    `json:"deletedByEmail,omitempty"`
	DeletedDateTime      *time.Time `json:"deletedDateTime,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ServiceChangeLog ServiceChangeLog

// NewServiceChangeLog instantiates a new ServiceChangeLog object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceChangeLog(createdBy string, createdByFullName string, createdByEmail string, createdDateTime time.Time, updatedBy string, updatedByFullName string, updatedByEmail string, updatedDateTime time.Time) *ServiceChangeLog {
	this := ServiceChangeLog{}
	this.CreatedBy = createdBy
	this.CreatedByFullName = createdByFullName
	this.CreatedByEmail = createdByEmail
	this.CreatedDateTime = createdDateTime
	this.UpdatedBy = updatedBy
	this.UpdatedByFullName = updatedByFullName
	this.UpdatedByEmail = updatedByEmail
	this.UpdatedDateTime = updatedDateTime
	return &this
}

// NewServiceChangeLogWithDefaults instantiates a new ServiceChangeLog object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceChangeLogWithDefaults() *ServiceChangeLog {
	this := ServiceChangeLog{}
	return &this
}

// GetCreatedBy returns the CreatedBy field value
func (o *ServiceChangeLog) GetCreatedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *ServiceChangeLog) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value
func (o *ServiceChangeLog) SetCreatedBy(v string) {
	o.CreatedBy = v
}

// GetCreatedByFullName returns the CreatedByFullName field value
func (o *ServiceChangeLog) GetCreatedByFullName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedByFullName
}

// GetCreatedByFullNameOk returns a tuple with the CreatedByFullName field value
// and a boolean to check if the value has been set.
func (o *ServiceChangeLog) GetCreatedByFullNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedByFullName, true
}

// SetCreatedByFullName sets field value
func (o *ServiceChangeLog) SetCreatedByFullName(v string) {
	o.CreatedByFullName = v
}

// GetCreatedByEmail returns the CreatedByEmail field value
func (o *ServiceChangeLog) GetCreatedByEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedByEmail
}

// GetCreatedByEmailOk returns a tuple with the CreatedByEmail field value
// and a boolean to check if the value has been set.
func (o *ServiceChangeLog) GetCreatedByEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedByEmail, true
}

// SetCreatedByEmail sets field value
func (o *ServiceChangeLog) SetCreatedByEmail(v string) {
	o.CreatedByEmail = v
}

// GetCreatedDateTime returns the CreatedDateTime field value
func (o *ServiceChangeLog) GetCreatedDateTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedDateTime
}

// GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field value
// and a boolean to check if the value has been set.
func (o *ServiceChangeLog) GetCreatedDateTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedDateTime, true
}

// SetCreatedDateTime sets field value
func (o *ServiceChangeLog) SetCreatedDateTime(v time.Time) {
	o.CreatedDateTime = v
}

// GetUpdatedBy returns the UpdatedBy field value
func (o *ServiceChangeLog) GetUpdatedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UpdatedBy
}

// GetUpdatedByOk returns a tuple with the UpdatedBy field value
// and a boolean to check if the value has been set.
func (o *ServiceChangeLog) GetUpdatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedBy, true
}

// SetUpdatedBy sets field value
func (o *ServiceChangeLog) SetUpdatedBy(v string) {
	o.UpdatedBy = v
}

// GetUpdatedByFullName returns the UpdatedByFullName field value
func (o *ServiceChangeLog) GetUpdatedByFullName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UpdatedByFullName
}

// GetUpdatedByFullNameOk returns a tuple with the UpdatedByFullName field value
// and a boolean to check if the value has been set.
func (o *ServiceChangeLog) GetUpdatedByFullNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedByFullName, true
}

// SetUpdatedByFullName sets field value
func (o *ServiceChangeLog) SetUpdatedByFullName(v string) {
	o.UpdatedByFullName = v
}

// GetUpdatedByEmail returns the UpdatedByEmail field value
func (o *ServiceChangeLog) GetUpdatedByEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UpdatedByEmail
}

// GetUpdatedByEmailOk returns a tuple with the UpdatedByEmail field value
// and a boolean to check if the value has been set.
func (o *ServiceChangeLog) GetUpdatedByEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedByEmail, true
}

// SetUpdatedByEmail sets field value
func (o *ServiceChangeLog) SetUpdatedByEmail(v string) {
	o.UpdatedByEmail = v
}

// GetUpdatedDateTime returns the UpdatedDateTime field value
func (o *ServiceChangeLog) GetUpdatedDateTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedDateTime
}

// GetUpdatedDateTimeOk returns a tuple with the UpdatedDateTime field value
// and a boolean to check if the value has been set.
func (o *ServiceChangeLog) GetUpdatedDateTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedDateTime, true
}

// SetUpdatedDateTime sets field value
func (o *ServiceChangeLog) SetUpdatedDateTime(v time.Time) {
	o.UpdatedDateTime = v
}

// GetDeletedBy returns the DeletedBy field value if set, zero value otherwise.
func (o *ServiceChangeLog) GetDeletedBy() string {
	if o == nil || IsNil(o.DeletedBy) {
		var ret string
		return ret
	}
	return *o.DeletedBy
}

// GetDeletedByOk returns a tuple with the DeletedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceChangeLog) GetDeletedByOk() (*string, bool) {
	if o == nil || IsNil(o.DeletedBy) {
		return nil, false
	}
	return o.DeletedBy, true
}

// HasDeletedBy returns a boolean if a field has been set.
func (o *ServiceChangeLog) HasDeletedBy() bool {
	if o != nil && !IsNil(o.DeletedBy) {
		return true
	}

	return false
}

// SetDeletedBy gets a reference to the given string and assigns it to the DeletedBy field.
func (o *ServiceChangeLog) SetDeletedBy(v string) {
	o.DeletedBy = &v
}

// GetDeletedByFullName returns the DeletedByFullName field value if set, zero value otherwise.
func (o *ServiceChangeLog) GetDeletedByFullName() string {
	if o == nil || IsNil(o.DeletedByFullName) {
		var ret string
		return ret
	}
	return *o.DeletedByFullName
}

// GetDeletedByFullNameOk returns a tuple with the DeletedByFullName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceChangeLog) GetDeletedByFullNameOk() (*string, bool) {
	if o == nil || IsNil(o.DeletedByFullName) {
		return nil, false
	}
	return o.DeletedByFullName, true
}

// HasDeletedByFullName returns a boolean if a field has been set.
func (o *ServiceChangeLog) HasDeletedByFullName() bool {
	if o != nil && !IsNil(o.DeletedByFullName) {
		return true
	}

	return false
}

// SetDeletedByFullName gets a reference to the given string and assigns it to the DeletedByFullName field.
func (o *ServiceChangeLog) SetDeletedByFullName(v string) {
	o.DeletedByFullName = &v
}

// GetDeletedByEmail returns the DeletedByEmail field value if set, zero value otherwise.
func (o *ServiceChangeLog) GetDeletedByEmail() string {
	if o == nil || IsNil(o.DeletedByEmail) {
		var ret string
		return ret
	}
	return *o.DeletedByEmail
}

// GetDeletedByEmailOk returns a tuple with the DeletedByEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceChangeLog) GetDeletedByEmailOk() (*string, bool) {
	if o == nil || IsNil(o.DeletedByEmail) {
		return nil, false
	}
	return o.DeletedByEmail, true
}

// HasDeletedByEmail returns a boolean if a field has been set.
func (o *ServiceChangeLog) HasDeletedByEmail() bool {
	if o != nil && !IsNil(o.DeletedByEmail) {
		return true
	}

	return false
}

// SetDeletedByEmail gets a reference to the given string and assigns it to the DeletedByEmail field.
func (o *ServiceChangeLog) SetDeletedByEmail(v string) {
	o.DeletedByEmail = &v
}

// GetDeletedDateTime returns the DeletedDateTime field value if set, zero value otherwise.
func (o *ServiceChangeLog) GetDeletedDateTime() time.Time {
	if o == nil || IsNil(o.DeletedDateTime) {
		var ret time.Time
		return ret
	}
	return *o.DeletedDateTime
}

// GetDeletedDateTimeOk returns a tuple with the DeletedDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceChangeLog) GetDeletedDateTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.DeletedDateTime) {
		return nil, false
	}
	return o.DeletedDateTime, true
}

// HasDeletedDateTime returns a boolean if a field has been set.
func (o *ServiceChangeLog) HasDeletedDateTime() bool {
	if o != nil && !IsNil(o.DeletedDateTime) {
		return true
	}

	return false
}

// SetDeletedDateTime gets a reference to the given time.Time and assigns it to the DeletedDateTime field.
func (o *ServiceChangeLog) SetDeletedDateTime(v time.Time) {
	o.DeletedDateTime = &v
}

func (o ServiceChangeLog) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceChangeLog) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["createdBy"] = o.CreatedBy
	toSerialize["createdByFullName"] = o.CreatedByFullName
	toSerialize["createdByEmail"] = o.CreatedByEmail
	toSerialize["createdDateTime"] = o.CreatedDateTime
	toSerialize["updatedBy"] = o.UpdatedBy
	toSerialize["updatedByFullName"] = o.UpdatedByFullName
	toSerialize["updatedByEmail"] = o.UpdatedByEmail
	toSerialize["updatedDateTime"] = o.UpdatedDateTime
	if !IsNil(o.DeletedBy) {
		toSerialize["deletedBy"] = o.DeletedBy
	}
	if !IsNil(o.DeletedByFullName) {
		toSerialize["deletedByFullName"] = o.DeletedByFullName
	}
	if !IsNil(o.DeletedByEmail) {
		toSerialize["deletedByEmail"] = o.DeletedByEmail
	}
	if !IsNil(o.DeletedDateTime) {
		toSerialize["deletedDateTime"] = o.DeletedDateTime
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ServiceChangeLog) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"createdBy",
		"createdByFullName",
		"createdByEmail",
		"createdDateTime",
		"updatedBy",
		"updatedByFullName",
		"updatedByEmail",
		"updatedDateTime",
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

	varServiceChangeLog := _ServiceChangeLog{}

	err = json.Unmarshal(data, &varServiceChangeLog)

	if err != nil {
		return err
	}

	*o = ServiceChangeLog(varServiceChangeLog)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "createdBy")
		delete(additionalProperties, "createdByFullName")
		delete(additionalProperties, "createdByEmail")
		delete(additionalProperties, "createdDateTime")
		delete(additionalProperties, "updatedBy")
		delete(additionalProperties, "updatedByFullName")
		delete(additionalProperties, "updatedByEmail")
		delete(additionalProperties, "updatedDateTime")
		delete(additionalProperties, "deletedBy")
		delete(additionalProperties, "deletedByFullName")
		delete(additionalProperties, "deletedByEmail")
		delete(additionalProperties, "deletedDateTime")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableServiceChangeLog struct {
	value *ServiceChangeLog
	isSet bool
}

func (v NullableServiceChangeLog) Get() *ServiceChangeLog {
	return v.value
}

func (v *NullableServiceChangeLog) Set(val *ServiceChangeLog) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceChangeLog) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceChangeLog) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceChangeLog(val *ServiceChangeLog) *NullableServiceChangeLog {
	return &NullableServiceChangeLog{value: val, isSet: true}
}

func (v NullableServiceChangeLog) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceChangeLog) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
