/*
Equinix Fabric API v4

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fabricv4

import (
	"encoding/json"
)

// checks if the ConnectionRouteTableEntrySearchResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConnectionRouteTableEntrySearchResponse{}

// ConnectionRouteTableEntrySearchResponse struct for ConnectionRouteTableEntrySearchResponse
type ConnectionRouteTableEntrySearchResponse struct {
	Pagination *Pagination `json:"pagination,omitempty"`
	// Data returned from the API call.
	Data                 []ConnectionRouteTableEntry `json:"data,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ConnectionRouteTableEntrySearchResponse ConnectionRouteTableEntrySearchResponse

// NewConnectionRouteTableEntrySearchResponse instantiates a new ConnectionRouteTableEntrySearchResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectionRouteTableEntrySearchResponse() *ConnectionRouteTableEntrySearchResponse {
	this := ConnectionRouteTableEntrySearchResponse{}
	return &this
}

// NewConnectionRouteTableEntrySearchResponseWithDefaults instantiates a new ConnectionRouteTableEntrySearchResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectionRouteTableEntrySearchResponseWithDefaults() *ConnectionRouteTableEntrySearchResponse {
	this := ConnectionRouteTableEntrySearchResponse{}
	return &this
}

// GetPagination returns the Pagination field value if set, zero value otherwise.
func (o *ConnectionRouteTableEntrySearchResponse) GetPagination() Pagination {
	if o == nil || IsNil(o.Pagination) {
		var ret Pagination
		return ret
	}
	return *o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionRouteTableEntrySearchResponse) GetPaginationOk() (*Pagination, bool) {
	if o == nil || IsNil(o.Pagination) {
		return nil, false
	}
	return o.Pagination, true
}

// HasPagination returns a boolean if a field has been set.
func (o *ConnectionRouteTableEntrySearchResponse) HasPagination() bool {
	if o != nil && !IsNil(o.Pagination) {
		return true
	}

	return false
}

// SetPagination gets a reference to the given Pagination and assigns it to the Pagination field.
func (o *ConnectionRouteTableEntrySearchResponse) SetPagination(v Pagination) {
	o.Pagination = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ConnectionRouteTableEntrySearchResponse) GetData() []ConnectionRouteTableEntry {
	if o == nil || IsNil(o.Data) {
		var ret []ConnectionRouteTableEntry
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionRouteTableEntrySearchResponse) GetDataOk() ([]ConnectionRouteTableEntry, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ConnectionRouteTableEntrySearchResponse) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []ConnectionRouteTableEntry and assigns it to the Data field.
func (o *ConnectionRouteTableEntrySearchResponse) SetData(v []ConnectionRouteTableEntry) {
	o.Data = v
}

func (o ConnectionRouteTableEntrySearchResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConnectionRouteTableEntrySearchResponse) ToMap() (map[string]interface{}, error) {
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

func (o *ConnectionRouteTableEntrySearchResponse) UnmarshalJSON(data []byte) (err error) {
	varConnectionRouteTableEntrySearchResponse := _ConnectionRouteTableEntrySearchResponse{}

	err = json.Unmarshal(data, &varConnectionRouteTableEntrySearchResponse)

	if err != nil {
		return err
	}

	*o = ConnectionRouteTableEntrySearchResponse(varConnectionRouteTableEntrySearchResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "pagination")
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableConnectionRouteTableEntrySearchResponse struct {
	value *ConnectionRouteTableEntrySearchResponse
	isSet bool
}

func (v NullableConnectionRouteTableEntrySearchResponse) Get() *ConnectionRouteTableEntrySearchResponse {
	return v.value
}

func (v *NullableConnectionRouteTableEntrySearchResponse) Set(val *ConnectionRouteTableEntrySearchResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectionRouteTableEntrySearchResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectionRouteTableEntrySearchResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectionRouteTableEntrySearchResponse(val *ConnectionRouteTableEntrySearchResponse) *NullableConnectionRouteTableEntrySearchResponse {
	return &NullableConnectionRouteTableEntrySearchResponse{value: val, isSet: true}
}

func (v NullableConnectionRouteTableEntrySearchResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectionRouteTableEntrySearchResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
