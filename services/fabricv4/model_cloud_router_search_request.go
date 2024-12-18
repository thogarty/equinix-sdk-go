/*
Equinix Fabric API v4

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fabricv4

import (
	"encoding/json"
)

// checks if the CloudRouterSearchRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CloudRouterSearchRequest{}

// CloudRouterSearchRequest Search requests containing criteria
type CloudRouterSearchRequest struct {
	Filter               *CloudRouterFilters       `json:"filter,omitempty"`
	Pagination           *PaginationRequest        `json:"pagination,omitempty"`
	Sort                 []CloudRouterSortCriteria `json:"sort,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CloudRouterSearchRequest CloudRouterSearchRequest

// NewCloudRouterSearchRequest instantiates a new CloudRouterSearchRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudRouterSearchRequest() *CloudRouterSearchRequest {
	this := CloudRouterSearchRequest{}
	return &this
}

// NewCloudRouterSearchRequestWithDefaults instantiates a new CloudRouterSearchRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudRouterSearchRequestWithDefaults() *CloudRouterSearchRequest {
	this := CloudRouterSearchRequest{}
	return &this
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *CloudRouterSearchRequest) GetFilter() CloudRouterFilters {
	if o == nil || IsNil(o.Filter) {
		var ret CloudRouterFilters
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudRouterSearchRequest) GetFilterOk() (*CloudRouterFilters, bool) {
	if o == nil || IsNil(o.Filter) {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *CloudRouterSearchRequest) HasFilter() bool {
	if o != nil && !IsNil(o.Filter) {
		return true
	}

	return false
}

// SetFilter gets a reference to the given CloudRouterFilters and assigns it to the Filter field.
func (o *CloudRouterSearchRequest) SetFilter(v CloudRouterFilters) {
	o.Filter = &v
}

// GetPagination returns the Pagination field value if set, zero value otherwise.
func (o *CloudRouterSearchRequest) GetPagination() PaginationRequest {
	if o == nil || IsNil(o.Pagination) {
		var ret PaginationRequest
		return ret
	}
	return *o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudRouterSearchRequest) GetPaginationOk() (*PaginationRequest, bool) {
	if o == nil || IsNil(o.Pagination) {
		return nil, false
	}
	return o.Pagination, true
}

// HasPagination returns a boolean if a field has been set.
func (o *CloudRouterSearchRequest) HasPagination() bool {
	if o != nil && !IsNil(o.Pagination) {
		return true
	}

	return false
}

// SetPagination gets a reference to the given PaginationRequest and assigns it to the Pagination field.
func (o *CloudRouterSearchRequest) SetPagination(v PaginationRequest) {
	o.Pagination = &v
}

// GetSort returns the Sort field value if set, zero value otherwise.
func (o *CloudRouterSearchRequest) GetSort() []CloudRouterSortCriteria {
	if o == nil || IsNil(o.Sort) {
		var ret []CloudRouterSortCriteria
		return ret
	}
	return o.Sort
}

// GetSortOk returns a tuple with the Sort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudRouterSearchRequest) GetSortOk() ([]CloudRouterSortCriteria, bool) {
	if o == nil || IsNil(o.Sort) {
		return nil, false
	}
	return o.Sort, true
}

// HasSort returns a boolean if a field has been set.
func (o *CloudRouterSearchRequest) HasSort() bool {
	if o != nil && !IsNil(o.Sort) {
		return true
	}

	return false
}

// SetSort gets a reference to the given []CloudRouterSortCriteria and assigns it to the Sort field.
func (o *CloudRouterSearchRequest) SetSort(v []CloudRouterSortCriteria) {
	o.Sort = v
}

func (o CloudRouterSearchRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CloudRouterSearchRequest) ToMap() (map[string]interface{}, error) {
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

func (o *CloudRouterSearchRequest) UnmarshalJSON(data []byte) (err error) {
	varCloudRouterSearchRequest := _CloudRouterSearchRequest{}

	err = json.Unmarshal(data, &varCloudRouterSearchRequest)

	if err != nil {
		return err
	}

	*o = CloudRouterSearchRequest(varCloudRouterSearchRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "filter")
		delete(additionalProperties, "pagination")
		delete(additionalProperties, "sort")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCloudRouterSearchRequest struct {
	value *CloudRouterSearchRequest
	isSet bool
}

func (v NullableCloudRouterSearchRequest) Get() *CloudRouterSearchRequest {
	return v.value
}

func (v *NullableCloudRouterSearchRequest) Set(val *CloudRouterSearchRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudRouterSearchRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudRouterSearchRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudRouterSearchRequest(val *CloudRouterSearchRequest) *NullableCloudRouterSearchRequest {
	return &NullableCloudRouterSearchRequest{value: val, isSet: true}
}

func (v NullableCloudRouterSearchRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudRouterSearchRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
