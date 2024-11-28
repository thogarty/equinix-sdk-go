/*
Equinix Internet Access API

Equinix Internet Access provides direct access to the Internet with scalable bandwidth options in IBX data centers.

Contact: api-support@equinix.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eiav2

import (
	"encoding/json"
)

// checks if the SearchRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchRequest{}

// SearchRequest struct for SearchRequest
type SearchRequest struct {
	Filter               *Filter `json:"filter,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SearchRequest SearchRequest

// NewSearchRequest instantiates a new SearchRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchRequest() *SearchRequest {
	this := SearchRequest{}
	return &this
}

// NewSearchRequestWithDefaults instantiates a new SearchRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchRequestWithDefaults() *SearchRequest {
	this := SearchRequest{}
	return &this
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *SearchRequest) GetFilter() Filter {
	if o == nil || IsNil(o.Filter) {
		var ret Filter
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchRequest) GetFilterOk() (*Filter, bool) {
	if o == nil || IsNil(o.Filter) {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *SearchRequest) HasFilter() bool {
	if o != nil && !IsNil(o.Filter) {
		return true
	}

	return false
}

// SetFilter gets a reference to the given Filter and assigns it to the Filter field.
func (o *SearchRequest) SetFilter(v Filter) {
	o.Filter = &v
}

func (o SearchRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Filter) {
		toSerialize["filter"] = o.Filter
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SearchRequest) UnmarshalJSON(data []byte) (err error) {
	varSearchRequest := _SearchRequest{}

	err = json.Unmarshal(data, &varSearchRequest)

	if err != nil {
		return err
	}

	*o = SearchRequest(varSearchRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "filter")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSearchRequest struct {
	value *SearchRequest
	isSet bool
}

func (v NullableSearchRequest) Get() *SearchRequest {
	return v.value
}

func (v *NullableSearchRequest) Set(val *SearchRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchRequest(val *SearchRequest) *NullableSearchRequest {
	return &NullableSearchRequest{value: val, isSet: true}
}

func (v NullableSearchRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}