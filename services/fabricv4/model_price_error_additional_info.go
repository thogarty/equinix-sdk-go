/*
Equinix Fabric API v4

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fabricv4

import (
	"encoding/json"
)

// checks if the PriceErrorAdditionalInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PriceErrorAdditionalInfo{}

// PriceErrorAdditionalInfo struct for PriceErrorAdditionalInfo
type PriceErrorAdditionalInfo struct {
	Property             *string `json:"property,omitempty"`
	Reason               *string `json:"reason,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PriceErrorAdditionalInfo PriceErrorAdditionalInfo

// NewPriceErrorAdditionalInfo instantiates a new PriceErrorAdditionalInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPriceErrorAdditionalInfo() *PriceErrorAdditionalInfo {
	this := PriceErrorAdditionalInfo{}
	return &this
}

// NewPriceErrorAdditionalInfoWithDefaults instantiates a new PriceErrorAdditionalInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPriceErrorAdditionalInfoWithDefaults() *PriceErrorAdditionalInfo {
	this := PriceErrorAdditionalInfo{}
	return &this
}

// GetProperty returns the Property field value if set, zero value otherwise.
func (o *PriceErrorAdditionalInfo) GetProperty() string {
	if o == nil || IsNil(o.Property) {
		var ret string
		return ret
	}
	return *o.Property
}

// GetPropertyOk returns a tuple with the Property field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceErrorAdditionalInfo) GetPropertyOk() (*string, bool) {
	if o == nil || IsNil(o.Property) {
		return nil, false
	}
	return o.Property, true
}

// HasProperty returns a boolean if a field has been set.
func (o *PriceErrorAdditionalInfo) HasProperty() bool {
	if o != nil && !IsNil(o.Property) {
		return true
	}

	return false
}

// SetProperty gets a reference to the given string and assigns it to the Property field.
func (o *PriceErrorAdditionalInfo) SetProperty(v string) {
	o.Property = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *PriceErrorAdditionalInfo) GetReason() string {
	if o == nil || IsNil(o.Reason) {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceErrorAdditionalInfo) GetReasonOk() (*string, bool) {
	if o == nil || IsNil(o.Reason) {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *PriceErrorAdditionalInfo) HasReason() bool {
	if o != nil && !IsNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *PriceErrorAdditionalInfo) SetReason(v string) {
	o.Reason = &v
}

func (o PriceErrorAdditionalInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PriceErrorAdditionalInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Property) {
		toSerialize["property"] = o.Property
	}
	if !IsNil(o.Reason) {
		toSerialize["reason"] = o.Reason
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PriceErrorAdditionalInfo) UnmarshalJSON(data []byte) (err error) {
	varPriceErrorAdditionalInfo := _PriceErrorAdditionalInfo{}

	err = json.Unmarshal(data, &varPriceErrorAdditionalInfo)

	if err != nil {
		return err
	}

	*o = PriceErrorAdditionalInfo(varPriceErrorAdditionalInfo)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "property")
		delete(additionalProperties, "reason")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePriceErrorAdditionalInfo struct {
	value *PriceErrorAdditionalInfo
	isSet bool
}

func (v NullablePriceErrorAdditionalInfo) Get() *PriceErrorAdditionalInfo {
	return v.value
}

func (v *NullablePriceErrorAdditionalInfo) Set(val *PriceErrorAdditionalInfo) {
	v.value = val
	v.isSet = true
}

func (v NullablePriceErrorAdditionalInfo) IsSet() bool {
	return v.isSet
}

func (v *NullablePriceErrorAdditionalInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePriceErrorAdditionalInfo(val *PriceErrorAdditionalInfo) *NullablePriceErrorAdditionalInfo {
	return &NullablePriceErrorAdditionalInfo{value: val, isSet: true}
}

func (v NullablePriceErrorAdditionalInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePriceErrorAdditionalInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
