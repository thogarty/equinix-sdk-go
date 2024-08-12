/*
Equinix Fabric API v4

Contact: api-support@equinix.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fabricv4

import (
	"encoding/json"
)

// checks if the ServiceSearchResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceSearchResponse{}

// ServiceSearchResponse struct for ServiceSearchResponse
type ServiceSearchResponse struct {
	Pagination *Pagination `json:"pagination,omitempty"`
	// Data returned from the API call.
	Data                 []PrecisionTimeServiceResponse `json:"data,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ServiceSearchResponse ServiceSearchResponse

// NewServiceSearchResponse instantiates a new ServiceSearchResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceSearchResponse() *ServiceSearchResponse {
	this := ServiceSearchResponse{}
	return &this
}

// NewServiceSearchResponseWithDefaults instantiates a new ServiceSearchResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceSearchResponseWithDefaults() *ServiceSearchResponse {
	this := ServiceSearchResponse{}
	return &this
}

// GetPagination returns the Pagination field value if set, zero value otherwise.
func (o *ServiceSearchResponse) GetPagination() Pagination {
	if o == nil || IsNil(o.Pagination) {
		var ret Pagination
		return ret
	}
	return *o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceSearchResponse) GetPaginationOk() (*Pagination, bool) {
	if o == nil || IsNil(o.Pagination) {
		return nil, false
	}
	return o.Pagination, true
}

// HasPagination returns a boolean if a field has been set.
func (o *ServiceSearchResponse) HasPagination() bool {
	if o != nil && !IsNil(o.Pagination) {
		return true
	}

	return false
}

// SetPagination gets a reference to the given Pagination and assigns it to the Pagination field.
func (o *ServiceSearchResponse) SetPagination(v Pagination) {
	o.Pagination = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ServiceSearchResponse) GetData() []PrecisionTimeServiceResponse {
	if o == nil || IsNil(o.Data) {
		var ret []PrecisionTimeServiceResponse
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceSearchResponse) GetDataOk() ([]PrecisionTimeServiceResponse, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ServiceSearchResponse) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []PrecisionTimeServiceResponse and assigns it to the Data field.
func (o *ServiceSearchResponse) SetData(v []PrecisionTimeServiceResponse) {
	o.Data = v
}

func (o ServiceSearchResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceSearchResponse) ToMap() (map[string]interface{}, error) {
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

func (o *ServiceSearchResponse) UnmarshalJSON(data []byte) (err error) {
	varServiceSearchResponse := _ServiceSearchResponse{}

	err = json.Unmarshal(data, &varServiceSearchResponse)

	if err != nil {
		return err
	}

	*o = ServiceSearchResponse(varServiceSearchResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "pagination")
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableServiceSearchResponse struct {
	value *ServiceSearchResponse
	isSet bool
}

func (v NullableServiceSearchResponse) Get() *ServiceSearchResponse {
	return v.value
}

func (v *NullableServiceSearchResponse) Set(val *ServiceSearchResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceSearchResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceSearchResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceSearchResponse(val *ServiceSearchResponse) *NullableServiceSearchResponse {
	return &NullableServiceSearchResponse{value: val, isSet: true}
}

func (v NullableServiceSearchResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceSearchResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
