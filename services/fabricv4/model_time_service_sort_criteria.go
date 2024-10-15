/*
Equinix Fabric API v4

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fabricv4

import (
	"encoding/json"
)

// checks if the TimeServiceSortCriteria type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TimeServiceSortCriteria{}

// TimeServiceSortCriteria struct for TimeServiceSortCriteria
type TimeServiceSortCriteria struct {
	Direction            *TimeServiceSortDirection `json:"direction,omitempty"`
	Property             *TimeServiceSortBy        `json:"property,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TimeServiceSortCriteria TimeServiceSortCriteria

// NewTimeServiceSortCriteria instantiates a new TimeServiceSortCriteria object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTimeServiceSortCriteria() *TimeServiceSortCriteria {
	this := TimeServiceSortCriteria{}
	var direction TimeServiceSortDirection = TIMESERVICESORTDIRECTION_DESC
	this.Direction = &direction
	var property TimeServiceSortBy = TIMESERVICESORTBY_CHANGE_LOG_UPDATED_DATE_TIME
	this.Property = &property
	return &this
}

// NewTimeServiceSortCriteriaWithDefaults instantiates a new TimeServiceSortCriteria object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTimeServiceSortCriteriaWithDefaults() *TimeServiceSortCriteria {
	this := TimeServiceSortCriteria{}
	var direction TimeServiceSortDirection = TIMESERVICESORTDIRECTION_DESC
	this.Direction = &direction
	var property TimeServiceSortBy = TIMESERVICESORTBY_CHANGE_LOG_UPDATED_DATE_TIME
	this.Property = &property
	return &this
}

// GetDirection returns the Direction field value if set, zero value otherwise.
func (o *TimeServiceSortCriteria) GetDirection() TimeServiceSortDirection {
	if o == nil || IsNil(o.Direction) {
		var ret TimeServiceSortDirection
		return ret
	}
	return *o.Direction
}

// GetDirectionOk returns a tuple with the Direction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeServiceSortCriteria) GetDirectionOk() (*TimeServiceSortDirection, bool) {
	if o == nil || IsNil(o.Direction) {
		return nil, false
	}
	return o.Direction, true
}

// HasDirection returns a boolean if a field has been set.
func (o *TimeServiceSortCriteria) HasDirection() bool {
	if o != nil && !IsNil(o.Direction) {
		return true
	}

	return false
}

// SetDirection gets a reference to the given TimeServiceSortDirection and assigns it to the Direction field.
func (o *TimeServiceSortCriteria) SetDirection(v TimeServiceSortDirection) {
	o.Direction = &v
}

// GetProperty returns the Property field value if set, zero value otherwise.
func (o *TimeServiceSortCriteria) GetProperty() TimeServiceSortBy {
	if o == nil || IsNil(o.Property) {
		var ret TimeServiceSortBy
		return ret
	}
	return *o.Property
}

// GetPropertyOk returns a tuple with the Property field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeServiceSortCriteria) GetPropertyOk() (*TimeServiceSortBy, bool) {
	if o == nil || IsNil(o.Property) {
		return nil, false
	}
	return o.Property, true
}

// HasProperty returns a boolean if a field has been set.
func (o *TimeServiceSortCriteria) HasProperty() bool {
	if o != nil && !IsNil(o.Property) {
		return true
	}

	return false
}

// SetProperty gets a reference to the given TimeServiceSortBy and assigns it to the Property field.
func (o *TimeServiceSortCriteria) SetProperty(v TimeServiceSortBy) {
	o.Property = &v
}

func (o TimeServiceSortCriteria) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TimeServiceSortCriteria) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Direction) {
		toSerialize["direction"] = o.Direction
	}
	if !IsNil(o.Property) {
		toSerialize["property"] = o.Property
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TimeServiceSortCriteria) UnmarshalJSON(data []byte) (err error) {
	varTimeServiceSortCriteria := _TimeServiceSortCriteria{}

	err = json.Unmarshal(data, &varTimeServiceSortCriteria)

	if err != nil {
		return err
	}

	*o = TimeServiceSortCriteria(varTimeServiceSortCriteria)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "direction")
		delete(additionalProperties, "property")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTimeServiceSortCriteria struct {
	value *TimeServiceSortCriteria
	isSet bool
}

func (v NullableTimeServiceSortCriteria) Get() *TimeServiceSortCriteria {
	return v.value
}

func (v *NullableTimeServiceSortCriteria) Set(val *TimeServiceSortCriteria) {
	v.value = val
	v.isSet = true
}

func (v NullableTimeServiceSortCriteria) IsSet() bool {
	return v.isSet
}

func (v *NullableTimeServiceSortCriteria) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTimeServiceSortCriteria(val *TimeServiceSortCriteria) *NullableTimeServiceSortCriteria {
	return &NullableTimeServiceSortCriteria{value: val, isSet: true}
}

func (v NullableTimeServiceSortCriteria) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTimeServiceSortCriteria) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
