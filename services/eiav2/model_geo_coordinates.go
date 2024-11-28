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

// checks if the GeoCoordinates type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GeoCoordinates{}

// GeoCoordinates struct for GeoCoordinates
type GeoCoordinates struct {
	Latitude             float32 `json:"latitude"`
	Longitude            float32 `json:"longitude"`
	AdditionalProperties map[string]interface{}
}

type _GeoCoordinates GeoCoordinates

// NewGeoCoordinates instantiates a new GeoCoordinates object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGeoCoordinates(latitude float32, longitude float32) *GeoCoordinates {
	this := GeoCoordinates{}
	this.Latitude = latitude
	this.Longitude = longitude
	return &this
}

// NewGeoCoordinatesWithDefaults instantiates a new GeoCoordinates object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGeoCoordinatesWithDefaults() *GeoCoordinates {
	this := GeoCoordinates{}
	return &this
}

// GetLatitude returns the Latitude field value
func (o *GeoCoordinates) GetLatitude() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Latitude
}

// GetLatitudeOk returns a tuple with the Latitude field value
// and a boolean to check if the value has been set.
func (o *GeoCoordinates) GetLatitudeOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Latitude, true
}

// SetLatitude sets field value
func (o *GeoCoordinates) SetLatitude(v float32) {
	o.Latitude = v
}

// GetLongitude returns the Longitude field value
func (o *GeoCoordinates) GetLongitude() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Longitude
}

// GetLongitudeOk returns a tuple with the Longitude field value
// and a boolean to check if the value has been set.
func (o *GeoCoordinates) GetLongitudeOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Longitude, true
}

// SetLongitude sets field value
func (o *GeoCoordinates) SetLongitude(v float32) {
	o.Longitude = v
}

func (o GeoCoordinates) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GeoCoordinates) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["latitude"] = o.Latitude
	toSerialize["longitude"] = o.Longitude

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GeoCoordinates) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"latitude",
		"longitude",
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

	varGeoCoordinates := _GeoCoordinates{}

	err = json.Unmarshal(data, &varGeoCoordinates)

	if err != nil {
		return err
	}

	*o = GeoCoordinates(varGeoCoordinates)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "latitude")
		delete(additionalProperties, "longitude")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGeoCoordinates struct {
	value *GeoCoordinates
	isSet bool
}

func (v NullableGeoCoordinates) Get() *GeoCoordinates {
	return v.value
}

func (v *NullableGeoCoordinates) Set(val *GeoCoordinates) {
	v.value = val
	v.isSet = true
}

func (v NullableGeoCoordinates) IsSet() bool {
	return v.isSet
}

func (v *NullableGeoCoordinates) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGeoCoordinates(val *GeoCoordinates) *NullableGeoCoordinates {
	return &NullableGeoCoordinates{value: val, isSet: true}
}

func (v NullableGeoCoordinates) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGeoCoordinates) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
