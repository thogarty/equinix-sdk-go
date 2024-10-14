/*
Equinix Fabric API v4

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fabricv4

import (
	"encoding/json"
)

// checks if the TopUtilizedStatistics type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TopUtilizedStatistics{}

// TopUtilizedStatistics This API provides service-level traffic metrics for the top utilized ports so that you can view access and gather key information required to manage service subscription sizing and capacity.
type TopUtilizedStatistics struct {
	Pagination *Pagination `json:"pagination,omitempty"`
	// Data returned from the API call.
	Data                 []Statistics `json:"data,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TopUtilizedStatistics TopUtilizedStatistics

// NewTopUtilizedStatistics instantiates a new TopUtilizedStatistics object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTopUtilizedStatistics() *TopUtilizedStatistics {
	this := TopUtilizedStatistics{}
	return &this
}

// NewTopUtilizedStatisticsWithDefaults instantiates a new TopUtilizedStatistics object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTopUtilizedStatisticsWithDefaults() *TopUtilizedStatistics {
	this := TopUtilizedStatistics{}
	return &this
}

// GetPagination returns the Pagination field value if set, zero value otherwise.
func (o *TopUtilizedStatistics) GetPagination() Pagination {
	if o == nil || IsNil(o.Pagination) {
		var ret Pagination
		return ret
	}
	return *o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TopUtilizedStatistics) GetPaginationOk() (*Pagination, bool) {
	if o == nil || IsNil(o.Pagination) {
		return nil, false
	}
	return o.Pagination, true
}

// HasPagination returns a boolean if a field has been set.
func (o *TopUtilizedStatistics) HasPagination() bool {
	if o != nil && !IsNil(o.Pagination) {
		return true
	}

	return false
}

// SetPagination gets a reference to the given Pagination and assigns it to the Pagination field.
func (o *TopUtilizedStatistics) SetPagination(v Pagination) {
	o.Pagination = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *TopUtilizedStatistics) GetData() []Statistics {
	if o == nil || IsNil(o.Data) {
		var ret []Statistics
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TopUtilizedStatistics) GetDataOk() ([]Statistics, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *TopUtilizedStatistics) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []Statistics and assigns it to the Data field.
func (o *TopUtilizedStatistics) SetData(v []Statistics) {
	o.Data = v
}

func (o TopUtilizedStatistics) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TopUtilizedStatistics) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Pagination) {
		toSerialize["pagination"] = o.Pagination
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TopUtilizedStatistics) UnmarshalJSON(data []byte) (err error) {
	varTopUtilizedStatistics := _TopUtilizedStatistics{}

	err = json.Unmarshal(data, &varTopUtilizedStatistics)

	if err != nil {
		return err
	}

	*o = TopUtilizedStatistics(varTopUtilizedStatistics)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "pagination")
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTopUtilizedStatistics struct {
	value *TopUtilizedStatistics
	isSet bool
}

func (v NullableTopUtilizedStatistics) Get() *TopUtilizedStatistics {
	return v.value
}

func (v *NullableTopUtilizedStatistics) Set(val *TopUtilizedStatistics) {
	v.value = val
	v.isSet = true
}

func (v NullableTopUtilizedStatistics) IsSet() bool {
	return v.isSet
}

func (v *NullableTopUtilizedStatistics) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTopUtilizedStatistics(val *TopUtilizedStatistics) *NullableTopUtilizedStatistics {
	return &NullableTopUtilizedStatistics{value: val, isSet: true}
}

func (v NullableTopUtilizedStatistics) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTopUtilizedStatistics) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
