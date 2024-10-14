/*
Equinix Fabric API v4

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fabricv4

import (
	"encoding/json"
)

// checks if the ValidateSubnetResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ValidateSubnetResponse{}

// ValidateSubnetResponse ValidateResponse
type ValidateSubnetResponse struct {
	// Additional information
	AdditionalInfo       []ConnectionSideAdditionalInfo `json:"additionalInfo,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ValidateSubnetResponse ValidateSubnetResponse

// NewValidateSubnetResponse instantiates a new ValidateSubnetResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewValidateSubnetResponse() *ValidateSubnetResponse {
	this := ValidateSubnetResponse{}
	return &this
}

// NewValidateSubnetResponseWithDefaults instantiates a new ValidateSubnetResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewValidateSubnetResponseWithDefaults() *ValidateSubnetResponse {
	this := ValidateSubnetResponse{}
	return &this
}

// GetAdditionalInfo returns the AdditionalInfo field value if set, zero value otherwise.
func (o *ValidateSubnetResponse) GetAdditionalInfo() []ConnectionSideAdditionalInfo {
	if o == nil || IsNil(o.AdditionalInfo) {
		var ret []ConnectionSideAdditionalInfo
		return ret
	}
	return o.AdditionalInfo
}

// GetAdditionalInfoOk returns a tuple with the AdditionalInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateSubnetResponse) GetAdditionalInfoOk() ([]ConnectionSideAdditionalInfo, bool) {
	if o == nil || IsNil(o.AdditionalInfo) {
		return nil, false
	}
	return o.AdditionalInfo, true
}

// HasAdditionalInfo returns a boolean if a field has been set.
func (o *ValidateSubnetResponse) HasAdditionalInfo() bool {
	if o != nil && !IsNil(o.AdditionalInfo) {
		return true
	}

	return false
}

// SetAdditionalInfo gets a reference to the given []ConnectionSideAdditionalInfo and assigns it to the AdditionalInfo field.
func (o *ValidateSubnetResponse) SetAdditionalInfo(v []ConnectionSideAdditionalInfo) {
	o.AdditionalInfo = v
}

func (o ValidateSubnetResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ValidateSubnetResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AdditionalInfo) {
		toSerialize["additionalInfo"] = o.AdditionalInfo
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ValidateSubnetResponse) UnmarshalJSON(data []byte) (err error) {
	varValidateSubnetResponse := _ValidateSubnetResponse{}

	err = json.Unmarshal(data, &varValidateSubnetResponse)

	if err != nil {
		return err
	}

	*o = ValidateSubnetResponse(varValidateSubnetResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "additionalInfo")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableValidateSubnetResponse struct {
	value *ValidateSubnetResponse
	isSet bool
}

func (v NullableValidateSubnetResponse) Get() *ValidateSubnetResponse {
	return v.value
}

func (v *NullableValidateSubnetResponse) Set(val *ValidateSubnetResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableValidateSubnetResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableValidateSubnetResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableValidateSubnetResponse(val *ValidateSubnetResponse) *NullableValidateSubnetResponse {
	return &NullableValidateSubnetResponse{value: val, isSet: true}
}

func (v NullableValidateSubnetResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableValidateSubnetResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
