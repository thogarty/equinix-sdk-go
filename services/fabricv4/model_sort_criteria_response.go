/*
Equinix Fabric API v4

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fabricv4

import (
	"encoding/json"
)

// checks if the SortCriteriaResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SortCriteriaResponse{}

// SortCriteriaResponse struct for SortCriteriaResponse
type SortCriteriaResponse struct {
	Direction            *SortDirection `json:"direction,omitempty"`
	Property             *SortBy        `json:"property,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SortCriteriaResponse SortCriteriaResponse

// NewSortCriteriaResponse instantiates a new SortCriteriaResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSortCriteriaResponse() *SortCriteriaResponse {
	this := SortCriteriaResponse{}
	var direction SortDirection = SORTDIRECTION_DESC
	this.Direction = &direction
	var property SortBy = SORTBY_CHANGE_LOG_UPDATED_DATE_TIME
	this.Property = &property
	return &this
}

// NewSortCriteriaResponseWithDefaults instantiates a new SortCriteriaResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSortCriteriaResponseWithDefaults() *SortCriteriaResponse {
	this := SortCriteriaResponse{}
	var direction SortDirection = SORTDIRECTION_DESC
	this.Direction = &direction
	var property SortBy = SORTBY_CHANGE_LOG_UPDATED_DATE_TIME
	this.Property = &property
	return &this
}

// GetDirection returns the Direction field value if set, zero value otherwise.
func (o *SortCriteriaResponse) GetDirection() SortDirection {
	if o == nil || IsNil(o.Direction) {
		var ret SortDirection
		return ret
	}
	return *o.Direction
}

// GetDirectionOk returns a tuple with the Direction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SortCriteriaResponse) GetDirectionOk() (*SortDirection, bool) {
	if o == nil || IsNil(o.Direction) {
		return nil, false
	}
	return o.Direction, true
}

// HasDirection returns a boolean if a field has been set.
func (o *SortCriteriaResponse) HasDirection() bool {
	if o != nil && !IsNil(o.Direction) {
		return true
	}

	return false
}

// SetDirection gets a reference to the given SortDirection and assigns it to the Direction field.
func (o *SortCriteriaResponse) SetDirection(v SortDirection) {
	o.Direction = &v
}

// GetProperty returns the Property field value if set, zero value otherwise.
func (o *SortCriteriaResponse) GetProperty() SortBy {
	if o == nil || IsNil(o.Property) {
		var ret SortBy
		return ret
	}
	return *o.Property
}

// GetPropertyOk returns a tuple with the Property field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SortCriteriaResponse) GetPropertyOk() (*SortBy, bool) {
	if o == nil || IsNil(o.Property) {
		return nil, false
	}
	return o.Property, true
}

// HasProperty returns a boolean if a field has been set.
func (o *SortCriteriaResponse) HasProperty() bool {
	if o != nil && !IsNil(o.Property) {
		return true
	}

	return false
}

// SetProperty gets a reference to the given SortBy and assigns it to the Property field.
func (o *SortCriteriaResponse) SetProperty(v SortBy) {
	o.Property = &v
}

func (o SortCriteriaResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SortCriteriaResponse) ToMap() (map[string]interface{}, error) {
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

func (o *SortCriteriaResponse) UnmarshalJSON(data []byte) (err error) {
	varSortCriteriaResponse := _SortCriteriaResponse{}

	err = json.Unmarshal(data, &varSortCriteriaResponse)

	if err != nil {
		return err
	}

	*o = SortCriteriaResponse(varSortCriteriaResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "direction")
		delete(additionalProperties, "property")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSortCriteriaResponse struct {
	value *SortCriteriaResponse
	isSet bool
}

func (v NullableSortCriteriaResponse) Get() *SortCriteriaResponse {
	return v.value
}

func (v *NullableSortCriteriaResponse) Set(val *SortCriteriaResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSortCriteriaResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSortCriteriaResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSortCriteriaResponse(val *SortCriteriaResponse) *NullableSortCriteriaResponse {
	return &NullableSortCriteriaResponse{value: val, isSet: true}
}

func (v NullableSortCriteriaResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSortCriteriaResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
