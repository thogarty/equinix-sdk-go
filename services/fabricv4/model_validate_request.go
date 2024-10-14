/*
Equinix Fabric API v4

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fabricv4

import (
	"encoding/json"
)

// checks if the ValidateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ValidateRequest{}

// ValidateRequest Validate connection auth api key or vlan
type ValidateRequest struct {
	Filter               *ValidateRequestFilter `json:"filter,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ValidateRequest ValidateRequest

// NewValidateRequest instantiates a new ValidateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewValidateRequest() *ValidateRequest {
	this := ValidateRequest{}
	return &this
}

// NewValidateRequestWithDefaults instantiates a new ValidateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewValidateRequestWithDefaults() *ValidateRequest {
	this := ValidateRequest{}
	return &this
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *ValidateRequest) GetFilter() ValidateRequestFilter {
	if o == nil || IsNil(o.Filter) {
		var ret ValidateRequestFilter
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateRequest) GetFilterOk() (*ValidateRequestFilter, bool) {
	if o == nil || IsNil(o.Filter) {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *ValidateRequest) HasFilter() bool {
	if o != nil && !IsNil(o.Filter) {
		return true
	}

	return false
}

// SetFilter gets a reference to the given ValidateRequestFilter and assigns it to the Filter field.
func (o *ValidateRequest) SetFilter(v ValidateRequestFilter) {
	o.Filter = &v
}

func (o ValidateRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ValidateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Filter) {
		toSerialize["filter"] = o.Filter
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ValidateRequest) UnmarshalJSON(data []byte) (err error) {
	varValidateRequest := _ValidateRequest{}

	err = json.Unmarshal(data, &varValidateRequest)

	if err != nil {
		return err
	}

	*o = ValidateRequest(varValidateRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "filter")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableValidateRequest struct {
	value *ValidateRequest
	isSet bool
}

func (v NullableValidateRequest) Get() *ValidateRequest {
	return v.value
}

func (v *NullableValidateRequest) Set(val *ValidateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableValidateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableValidateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableValidateRequest(val *ValidateRequest) *NullableValidateRequest {
	return &NullableValidateRequest{value: val, isSet: true}
}

func (v NullableValidateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableValidateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
