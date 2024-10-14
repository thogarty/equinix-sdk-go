/*
Equinix Fabric API v4

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fabricv4

import (
	"encoding/json"
)

// checks if the PortV4SearchRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PortV4SearchRequest{}

// PortV4SearchRequest Search requests containing criteria
type PortV4SearchRequest struct {
	Filter               *PortExpression    `json:"filter,omitempty"`
	Pagination           *PaginationRequest `json:"pagination,omitempty"`
	Sort                 []PortSortCriteria `json:"sort,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PortV4SearchRequest PortV4SearchRequest

// NewPortV4SearchRequest instantiates a new PortV4SearchRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPortV4SearchRequest() *PortV4SearchRequest {
	this := PortV4SearchRequest{}
	return &this
}

// NewPortV4SearchRequestWithDefaults instantiates a new PortV4SearchRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPortV4SearchRequestWithDefaults() *PortV4SearchRequest {
	this := PortV4SearchRequest{}
	return &this
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *PortV4SearchRequest) GetFilter() PortExpression {
	if o == nil || IsNil(o.Filter) {
		var ret PortExpression
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortV4SearchRequest) GetFilterOk() (*PortExpression, bool) {
	if o == nil || IsNil(o.Filter) {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *PortV4SearchRequest) HasFilter() bool {
	if o != nil && !IsNil(o.Filter) {
		return true
	}

	return false
}

// SetFilter gets a reference to the given PortExpression and assigns it to the Filter field.
func (o *PortV4SearchRequest) SetFilter(v PortExpression) {
	o.Filter = &v
}

// GetPagination returns the Pagination field value if set, zero value otherwise.
func (o *PortV4SearchRequest) GetPagination() PaginationRequest {
	if o == nil || IsNil(o.Pagination) {
		var ret PaginationRequest
		return ret
	}
	return *o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortV4SearchRequest) GetPaginationOk() (*PaginationRequest, bool) {
	if o == nil || IsNil(o.Pagination) {
		return nil, false
	}
	return o.Pagination, true
}

// HasPagination returns a boolean if a field has been set.
func (o *PortV4SearchRequest) HasPagination() bool {
	if o != nil && !IsNil(o.Pagination) {
		return true
	}

	return false
}

// SetPagination gets a reference to the given PaginationRequest and assigns it to the Pagination field.
func (o *PortV4SearchRequest) SetPagination(v PaginationRequest) {
	o.Pagination = &v
}

// GetSort returns the Sort field value if set, zero value otherwise.
func (o *PortV4SearchRequest) GetSort() []PortSortCriteria {
	if o == nil || IsNil(o.Sort) {
		var ret []PortSortCriteria
		return ret
	}
	return o.Sort
}

// GetSortOk returns a tuple with the Sort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortV4SearchRequest) GetSortOk() ([]PortSortCriteria, bool) {
	if o == nil || IsNil(o.Sort) {
		return nil, false
	}
	return o.Sort, true
}

// HasSort returns a boolean if a field has been set.
func (o *PortV4SearchRequest) HasSort() bool {
	if o != nil && !IsNil(o.Sort) {
		return true
	}

	return false
}

// SetSort gets a reference to the given []PortSortCriteria and assigns it to the Sort field.
func (o *PortV4SearchRequest) SetSort(v []PortSortCriteria) {
	o.Sort = v
}

func (o PortV4SearchRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PortV4SearchRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Filter) {
		toSerialize["filter"] = o.Filter
	}
	if !IsNil(o.Pagination) {
		toSerialize["pagination"] = o.Pagination
	}
	if !IsNil(o.Sort) {
		toSerialize["sort"] = o.Sort
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PortV4SearchRequest) UnmarshalJSON(data []byte) (err error) {
	varPortV4SearchRequest := _PortV4SearchRequest{}

	err = json.Unmarshal(data, &varPortV4SearchRequest)

	if err != nil {
		return err
	}

	*o = PortV4SearchRequest(varPortV4SearchRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "filter")
		delete(additionalProperties, "pagination")
		delete(additionalProperties, "sort")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePortV4SearchRequest struct {
	value *PortV4SearchRequest
	isSet bool
}

func (v NullablePortV4SearchRequest) Get() *PortV4SearchRequest {
	return v.value
}

func (v *NullablePortV4SearchRequest) Set(val *PortV4SearchRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePortV4SearchRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePortV4SearchRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePortV4SearchRequest(val *PortV4SearchRequest) *NullablePortV4SearchRequest {
	return &NullablePortV4SearchRequest{value: val, isSet: true}
}

func (v NullablePortV4SearchRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePortV4SearchRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
