/*
Equinix Fabric API v4

Equinix Fabric is an advanced software-defined interconnection solution that enables you to directly, securely and dynamically connect to distributed infrastructure and digital ecosystems on platform Equinix via a single port, Customers can use Fabric to connect to: </br> 1. Cloud Service Providers - Clouds, network and other service providers.  </br> 2. Enterprises - Other Equinix customers, vendors and partners.  </br> 3. Myself - Another customer instance deployed at Equinix. </br>

API version: 4.12
Contact: api-support@equinix.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fabricv4

import (
	"encoding/json"
)

// checks if the RouteFilterRulesChangeDataResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RouteFilterRulesChangeDataResponse{}

// RouteFilterRulesChangeDataResponse List of Route Filter Rule changes
type RouteFilterRulesChangeDataResponse struct {
	Pagination           *Pagination                  `json:"pagination,omitempty"`
	Data                 []RouteFilterRulesChangeData `json:"data,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RouteFilterRulesChangeDataResponse RouteFilterRulesChangeDataResponse

// NewRouteFilterRulesChangeDataResponse instantiates a new RouteFilterRulesChangeDataResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRouteFilterRulesChangeDataResponse() *RouteFilterRulesChangeDataResponse {
	this := RouteFilterRulesChangeDataResponse{}
	return &this
}

// NewRouteFilterRulesChangeDataResponseWithDefaults instantiates a new RouteFilterRulesChangeDataResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRouteFilterRulesChangeDataResponseWithDefaults() *RouteFilterRulesChangeDataResponse {
	this := RouteFilterRulesChangeDataResponse{}
	return &this
}

// GetPagination returns the Pagination field value if set, zero value otherwise.
func (o *RouteFilterRulesChangeDataResponse) GetPagination() Pagination {
	if o == nil || IsNil(o.Pagination) {
		var ret Pagination
		return ret
	}
	return *o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RouteFilterRulesChangeDataResponse) GetPaginationOk() (*Pagination, bool) {
	if o == nil || IsNil(o.Pagination) {
		return nil, false
	}
	return o.Pagination, true
}

// HasPagination returns a boolean if a field has been set.
func (o *RouteFilterRulesChangeDataResponse) HasPagination() bool {
	if o != nil && !IsNil(o.Pagination) {
		return true
	}

	return false
}

// SetPagination gets a reference to the given Pagination and assigns it to the Pagination field.
func (o *RouteFilterRulesChangeDataResponse) SetPagination(v Pagination) {
	o.Pagination = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *RouteFilterRulesChangeDataResponse) GetData() []RouteFilterRulesChangeData {
	if o == nil || IsNil(o.Data) {
		var ret []RouteFilterRulesChangeData
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RouteFilterRulesChangeDataResponse) GetDataOk() ([]RouteFilterRulesChangeData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *RouteFilterRulesChangeDataResponse) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []RouteFilterRulesChangeData and assigns it to the Data field.
func (o *RouteFilterRulesChangeDataResponse) SetData(v []RouteFilterRulesChangeData) {
	o.Data = v
}

func (o RouteFilterRulesChangeDataResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RouteFilterRulesChangeDataResponse) ToMap() (map[string]interface{}, error) {
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

func (o *RouteFilterRulesChangeDataResponse) UnmarshalJSON(data []byte) (err error) {
	varRouteFilterRulesChangeDataResponse := _RouteFilterRulesChangeDataResponse{}

	err = json.Unmarshal(data, &varRouteFilterRulesChangeDataResponse)

	if err != nil {
		return err
	}

	*o = RouteFilterRulesChangeDataResponse(varRouteFilterRulesChangeDataResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "pagination")
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRouteFilterRulesChangeDataResponse struct {
	value *RouteFilterRulesChangeDataResponse
	isSet bool
}

func (v NullableRouteFilterRulesChangeDataResponse) Get() *RouteFilterRulesChangeDataResponse {
	return v.value
}

func (v *NullableRouteFilterRulesChangeDataResponse) Set(val *RouteFilterRulesChangeDataResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRouteFilterRulesChangeDataResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRouteFilterRulesChangeDataResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRouteFilterRulesChangeDataResponse(val *RouteFilterRulesChangeDataResponse) *NullableRouteFilterRulesChangeDataResponse {
	return &NullableRouteFilterRulesChangeDataResponse{value: val, isSet: true}
}

func (v NullableRouteFilterRulesChangeDataResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRouteFilterRulesChangeDataResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}