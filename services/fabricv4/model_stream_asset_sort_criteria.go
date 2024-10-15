/*
Equinix Fabric API v4

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fabricv4

import (
	"encoding/json"
)

// checks if the StreamAssetSortCriteria type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StreamAssetSortCriteria{}

// StreamAssetSortCriteria struct for StreamAssetSortCriteria
type StreamAssetSortCriteria struct {
	Direction            *StreamAssetSortDirection `json:"direction,omitempty"`
	Property             *StreamAssetSortBy        `json:"property,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StreamAssetSortCriteria StreamAssetSortCriteria

// NewStreamAssetSortCriteria instantiates a new StreamAssetSortCriteria object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStreamAssetSortCriteria() *StreamAssetSortCriteria {
	this := StreamAssetSortCriteria{}
	var direction StreamAssetSortDirection = STREAMASSETSORTDIRECTION_DESC
	this.Direction = &direction
	var property StreamAssetSortBy = STREAMASSETSORTBY_UUID
	this.Property = &property
	return &this
}

// NewStreamAssetSortCriteriaWithDefaults instantiates a new StreamAssetSortCriteria object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStreamAssetSortCriteriaWithDefaults() *StreamAssetSortCriteria {
	this := StreamAssetSortCriteria{}
	var direction StreamAssetSortDirection = STREAMASSETSORTDIRECTION_DESC
	this.Direction = &direction
	var property StreamAssetSortBy = STREAMASSETSORTBY_UUID
	this.Property = &property
	return &this
}

// GetDirection returns the Direction field value if set, zero value otherwise.
func (o *StreamAssetSortCriteria) GetDirection() StreamAssetSortDirection {
	if o == nil || IsNil(o.Direction) {
		var ret StreamAssetSortDirection
		return ret
	}
	return *o.Direction
}

// GetDirectionOk returns a tuple with the Direction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamAssetSortCriteria) GetDirectionOk() (*StreamAssetSortDirection, bool) {
	if o == nil || IsNil(o.Direction) {
		return nil, false
	}
	return o.Direction, true
}

// HasDirection returns a boolean if a field has been set.
func (o *StreamAssetSortCriteria) HasDirection() bool {
	if o != nil && !IsNil(o.Direction) {
		return true
	}

	return false
}

// SetDirection gets a reference to the given StreamAssetSortDirection and assigns it to the Direction field.
func (o *StreamAssetSortCriteria) SetDirection(v StreamAssetSortDirection) {
	o.Direction = &v
}

// GetProperty returns the Property field value if set, zero value otherwise.
func (o *StreamAssetSortCriteria) GetProperty() StreamAssetSortBy {
	if o == nil || IsNil(o.Property) {
		var ret StreamAssetSortBy
		return ret
	}
	return *o.Property
}

// GetPropertyOk returns a tuple with the Property field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamAssetSortCriteria) GetPropertyOk() (*StreamAssetSortBy, bool) {
	if o == nil || IsNil(o.Property) {
		return nil, false
	}
	return o.Property, true
}

// HasProperty returns a boolean if a field has been set.
func (o *StreamAssetSortCriteria) HasProperty() bool {
	if o != nil && !IsNil(o.Property) {
		return true
	}

	return false
}

// SetProperty gets a reference to the given StreamAssetSortBy and assigns it to the Property field.
func (o *StreamAssetSortCriteria) SetProperty(v StreamAssetSortBy) {
	o.Property = &v
}

func (o StreamAssetSortCriteria) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StreamAssetSortCriteria) ToMap() (map[string]interface{}, error) {
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

func (o *StreamAssetSortCriteria) UnmarshalJSON(data []byte) (err error) {
	varStreamAssetSortCriteria := _StreamAssetSortCriteria{}

	err = json.Unmarshal(data, &varStreamAssetSortCriteria)

	if err != nil {
		return err
	}

	*o = StreamAssetSortCriteria(varStreamAssetSortCriteria)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "direction")
		delete(additionalProperties, "property")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStreamAssetSortCriteria struct {
	value *StreamAssetSortCriteria
	isSet bool
}

func (v NullableStreamAssetSortCriteria) Get() *StreamAssetSortCriteria {
	return v.value
}

func (v *NullableStreamAssetSortCriteria) Set(val *StreamAssetSortCriteria) {
	v.value = val
	v.isSet = true
}

func (v NullableStreamAssetSortCriteria) IsSet() bool {
	return v.isSet
}

func (v *NullableStreamAssetSortCriteria) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStreamAssetSortCriteria(val *StreamAssetSortCriteria) *NullableStreamAssetSortCriteria {
	return &NullableStreamAssetSortCriteria{value: val, isSet: true}
}

func (v NullableStreamAssetSortCriteria) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStreamAssetSortCriteria) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}