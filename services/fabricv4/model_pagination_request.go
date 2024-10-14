/*
Equinix Fabric API v4

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fabricv4

import (
	"encoding/json"
)

// checks if the PaginationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaginationRequest{}

// PaginationRequest Pagination request information
type PaginationRequest struct {
	// Index of the first element.
	Offset *int32 `json:"offset,omitempty"`
	// Number of elements to be requested per page. Number must be between 1 and 100, and the default is 20.
	Limit                *int32 `json:"limit,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PaginationRequest PaginationRequest

// NewPaginationRequest instantiates a new PaginationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginationRequest() *PaginationRequest {
	this := PaginationRequest{}
	var offset int32 = 0
	this.Offset = &offset
	var limit int32 = 20
	this.Limit = &limit
	return &this
}

// NewPaginationRequestWithDefaults instantiates a new PaginationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginationRequestWithDefaults() *PaginationRequest {
	this := PaginationRequest{}
	var offset int32 = 0
	this.Offset = &offset
	var limit int32 = 20
	this.Limit = &limit
	return &this
}

// GetOffset returns the Offset field value if set, zero value otherwise.
func (o *PaginationRequest) GetOffset() int32 {
	if o == nil || IsNil(o.Offset) {
		var ret int32
		return ret
	}
	return *o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginationRequest) GetOffsetOk() (*int32, bool) {
	if o == nil || IsNil(o.Offset) {
		return nil, false
	}
	return o.Offset, true
}

// HasOffset returns a boolean if a field has been set.
func (o *PaginationRequest) HasOffset() bool {
	if o != nil && !IsNil(o.Offset) {
		return true
	}

	return false
}

// SetOffset gets a reference to the given int32 and assigns it to the Offset field.
func (o *PaginationRequest) SetOffset(v int32) {
	o.Offset = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *PaginationRequest) GetLimit() int32 {
	if o == nil || IsNil(o.Limit) {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginationRequest) GetLimitOk() (*int32, bool) {
	if o == nil || IsNil(o.Limit) {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *PaginationRequest) HasLimit() bool {
	if o != nil && !IsNil(o.Limit) {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *PaginationRequest) SetLimit(v int32) {
	o.Limit = &v
}

func (o PaginationRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaginationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Offset) {
		toSerialize["offset"] = o.Offset
	}
	if !IsNil(o.Limit) {
		toSerialize["limit"] = o.Limit
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PaginationRequest) UnmarshalJSON(data []byte) (err error) {
	varPaginationRequest := _PaginationRequest{}

	err = json.Unmarshal(data, &varPaginationRequest)

	if err != nil {
		return err
	}

	*o = PaginationRequest(varPaginationRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "offset")
		delete(additionalProperties, "limit")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePaginationRequest struct {
	value *PaginationRequest
	isSet bool
}

func (v NullablePaginationRequest) Get() *PaginationRequest {
	return v.value
}

func (v *NullablePaginationRequest) Set(val *PaginationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginationRequest(val *PaginationRequest) *NullablePaginationRequest {
	return &NullablePaginationRequest{value: val, isSet: true}
}

func (v NullablePaginationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
