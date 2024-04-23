/*
Equinix Fabric API v4

Equinix Fabric is an advanced software-defined interconnection solution that enables you to directly, securely and dynamically connect to distributed infrastructure and digital ecosystems on platform Equinix via a single port, Customers can use Fabric to connect to: </br> 1. Cloud Service Providers - Clouds, network and other service providers.  </br> 2. Enterprises - Other Equinix customers, vendors and partners.  </br> 3. Myself - Another customer instance deployed at Equinix. </br>

Contact: api-support@equinix.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fabricv4

import (
	"encoding/json"
)

// checks if the AllPortsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AllPortsResponse{}

// AllPortsResponse GET All User Port Across Fabric Metros
type AllPortsResponse struct {
	Pagination *Pagination `json:"pagination,omitempty"`
	// GET All User Port Across Fabric Metros
	Data                 []PortResponse `json:"data,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AllPortsResponse AllPortsResponse

// NewAllPortsResponse instantiates a new AllPortsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAllPortsResponse() *AllPortsResponse {
	this := AllPortsResponse{}
	return &this
}

// NewAllPortsResponseWithDefaults instantiates a new AllPortsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAllPortsResponseWithDefaults() *AllPortsResponse {
	this := AllPortsResponse{}
	return &this
}

// GetPagination returns the Pagination field value if set, zero value otherwise.
func (o *AllPortsResponse) GetPagination() Pagination {
	if o == nil || IsNil(o.Pagination) {
		var ret Pagination
		return ret
	}
	return *o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllPortsResponse) GetPaginationOk() (*Pagination, bool) {
	if o == nil || IsNil(o.Pagination) {
		return nil, false
	}
	return o.Pagination, true
}

// HasPagination returns a boolean if a field has been set.
func (o *AllPortsResponse) HasPagination() bool {
	if o != nil && !IsNil(o.Pagination) {
		return true
	}

	return false
}

// SetPagination gets a reference to the given Pagination and assigns it to the Pagination field.
func (o *AllPortsResponse) SetPagination(v Pagination) {
	o.Pagination = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *AllPortsResponse) GetData() []PortResponse {
	if o == nil || IsNil(o.Data) {
		var ret []PortResponse
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllPortsResponse) GetDataOk() ([]PortResponse, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *AllPortsResponse) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []PortResponse and assigns it to the Data field.
func (o *AllPortsResponse) SetData(v []PortResponse) {
	o.Data = v
}

func (o AllPortsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AllPortsResponse) ToMap() (map[string]interface{}, error) {
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

func (o *AllPortsResponse) UnmarshalJSON(data []byte) (err error) {
	varAllPortsResponse := _AllPortsResponse{}

	err = json.Unmarshal(data, &varAllPortsResponse)

	if err != nil {
		return err
	}

	*o = AllPortsResponse(varAllPortsResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "pagination")
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAllPortsResponse struct {
	value *AllPortsResponse
	isSet bool
}

func (v NullableAllPortsResponse) Get() *AllPortsResponse {
	return v.value
}

func (v *NullableAllPortsResponse) Set(val *AllPortsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAllPortsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAllPortsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAllPortsResponse(val *AllPortsResponse) *NullableAllPortsResponse {
	return &NullableAllPortsResponse{value: val, isSet: true}
}

func (v NullableAllPortsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAllPortsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
