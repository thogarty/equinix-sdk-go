/*
Equinix Fabric API v4

Contact: api-support@equinix.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fabricv4

import (
	"encoding/json"
	"fmt"
)

// checks if the SimplifiedLocationWithoutIBX type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SimplifiedLocationWithoutIBX{}

// SimplifiedLocationWithoutIBX struct for SimplifiedLocationWithoutIBX
type SimplifiedLocationWithoutIBX struct {
	// The Canonical URL at which the resource resides.
	Href                 *string `json:"href,omitempty"`
	Region               *string `json:"region,omitempty"`
	MetroName            *string `json:"metroName,omitempty"`
	MetroCode            string  `json:"metroCode"`
	AdditionalProperties map[string]interface{}
}

type _SimplifiedLocationWithoutIBX SimplifiedLocationWithoutIBX

// NewSimplifiedLocationWithoutIBX instantiates a new SimplifiedLocationWithoutIBX object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSimplifiedLocationWithoutIBX(metroCode string) *SimplifiedLocationWithoutIBX {
	this := SimplifiedLocationWithoutIBX{}
	this.MetroCode = metroCode
	return &this
}

// NewSimplifiedLocationWithoutIBXWithDefaults instantiates a new SimplifiedLocationWithoutIBX object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSimplifiedLocationWithoutIBXWithDefaults() *SimplifiedLocationWithoutIBX {
	this := SimplifiedLocationWithoutIBX{}
	return &this
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *SimplifiedLocationWithoutIBX) GetHref() string {
	if o == nil || IsNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimplifiedLocationWithoutIBX) GetHrefOk() (*string, bool) {
	if o == nil || IsNil(o.Href) {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *SimplifiedLocationWithoutIBX) HasHref() bool {
	if o != nil && !IsNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *SimplifiedLocationWithoutIBX) SetHref(v string) {
	o.Href = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *SimplifiedLocationWithoutIBX) GetRegion() string {
	if o == nil || IsNil(o.Region) {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimplifiedLocationWithoutIBX) GetRegionOk() (*string, bool) {
	if o == nil || IsNil(o.Region) {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *SimplifiedLocationWithoutIBX) HasRegion() bool {
	if o != nil && !IsNil(o.Region) {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *SimplifiedLocationWithoutIBX) SetRegion(v string) {
	o.Region = &v
}

// GetMetroName returns the MetroName field value if set, zero value otherwise.
func (o *SimplifiedLocationWithoutIBX) GetMetroName() string {
	if o == nil || IsNil(o.MetroName) {
		var ret string
		return ret
	}
	return *o.MetroName
}

// GetMetroNameOk returns a tuple with the MetroName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimplifiedLocationWithoutIBX) GetMetroNameOk() (*string, bool) {
	if o == nil || IsNil(o.MetroName) {
		return nil, false
	}
	return o.MetroName, true
}

// HasMetroName returns a boolean if a field has been set.
func (o *SimplifiedLocationWithoutIBX) HasMetroName() bool {
	if o != nil && !IsNil(o.MetroName) {
		return true
	}

	return false
}

// SetMetroName gets a reference to the given string and assigns it to the MetroName field.
func (o *SimplifiedLocationWithoutIBX) SetMetroName(v string) {
	o.MetroName = &v
}

// GetMetroCode returns the MetroCode field value
func (o *SimplifiedLocationWithoutIBX) GetMetroCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MetroCode
}

// GetMetroCodeOk returns a tuple with the MetroCode field value
// and a boolean to check if the value has been set.
func (o *SimplifiedLocationWithoutIBX) GetMetroCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MetroCode, true
}

// SetMetroCode sets field value
func (o *SimplifiedLocationWithoutIBX) SetMetroCode(v string) {
	o.MetroCode = v
}

func (o SimplifiedLocationWithoutIBX) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SimplifiedLocationWithoutIBX) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Href) {
		toSerialize["href"] = o.Href
	}
	if !IsNil(o.Region) {
		toSerialize["region"] = o.Region
	}
	if !IsNil(o.MetroName) {
		toSerialize["metroName"] = o.MetroName
	}
	toSerialize["metroCode"] = o.MetroCode

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SimplifiedLocationWithoutIBX) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"metroCode",
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

	varSimplifiedLocationWithoutIBX := _SimplifiedLocationWithoutIBX{}

	err = json.Unmarshal(data, &varSimplifiedLocationWithoutIBX)

	if err != nil {
		return err
	}

	*o = SimplifiedLocationWithoutIBX(varSimplifiedLocationWithoutIBX)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "href")
		delete(additionalProperties, "region")
		delete(additionalProperties, "metroName")
		delete(additionalProperties, "metroCode")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSimplifiedLocationWithoutIBX struct {
	value *SimplifiedLocationWithoutIBX
	isSet bool
}

func (v NullableSimplifiedLocationWithoutIBX) Get() *SimplifiedLocationWithoutIBX {
	return v.value
}

func (v *NullableSimplifiedLocationWithoutIBX) Set(val *SimplifiedLocationWithoutIBX) {
	v.value = val
	v.isSet = true
}

func (v NullableSimplifiedLocationWithoutIBX) IsSet() bool {
	return v.isSet
}

func (v *NullableSimplifiedLocationWithoutIBX) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSimplifiedLocationWithoutIBX(val *SimplifiedLocationWithoutIBX) *NullableSimplifiedLocationWithoutIBX {
	return &NullableSimplifiedLocationWithoutIBX{value: val, isSet: true}
}

func (v NullableSimplifiedLocationWithoutIBX) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSimplifiedLocationWithoutIBX) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
