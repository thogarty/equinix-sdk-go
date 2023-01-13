/*
Metal API

# Introduction Equinix Metal provides a RESTful HTTP API which can be reached at <https://api.equinix.com/metal/v1>. This document describes the API and how to use it.  The API allows you to programmatically interact with all of your Equinix Metal resources, including devices, networks, addresses, organizations, projects, and your user account. Every feature of the Equinix Metal web interface is accessible through the API.  The API docs are generated from the Equinix Metal OpenAPI specification and are officially hosted at <https://metal.equinix.com/developers/api>.  # Common Parameters  The Equinix Metal API uses a few methods to minimize network traffic and improve throughput. These parameters are not used in all API calls, but are used often enough to warrant their own section. Look for these parameters in the documentation for the API calls that support them.  ## Pagination  Pagination is used to limit the number of results returned in a single request. The API will return a maximum of 100 results per page. To retrieve additional results, you can use the `page` and `per_page` query parameters.  The `page` parameter is used to specify the page number. The first page is `1`. The `per_page` parameter is used to specify the number of results per page. The maximum number of results differs by resource type.  ## Sorting  Where offered, the API allows you to sort results by a specific field. To sort results use the `sort_by` query parameter with the root level field name as the value. The `sort_direction` parameter is used to specify the sort direction, either either `asc` (ascending) or `desc` (descending).  ## Filtering  Filtering is used to limit the results returned in a single request. The API supports filtering by certain fields in the response. To filter results, you can use the field as a query parameter.  For example, to filter the IP list to only return public IPv4 addresses, you can filter by the `type` field, as in the following request:  ```sh curl -H 'X-Auth-Token: my_authentication_token' \\   https://api.equinix.com/metal/v1/projects/id/ips?type=public_ipv4 ```  Only IP addresses with the `type` field set to `public_ipv4` will be returned.  ## Searching  Searching is used to find matching resources using multiple field comparissons. The API supports searching in resources that define this behavior. The fields available for search differ by resource, as does the search strategy.  To search resources you can use the `search` query parameter.  ## Include and Exclude  For resources that contain references to other resources, sucha as a Device that refers to the Project it resides in, the Equinix Metal API will returns `href` values (API links) to the associated resource.  ```json {   ...   \"project\": {     \"href\": \"/metal/v1/projects/f3f131c8-f302-49ef-8c44-9405022dc6dd\"   } } ```  If you're going need the project details, you can avoid a second API request.  Specify the contained `href` resources and collections that you'd like to have included in the response using the `include` query parameter.  For example:    ```sh curl -H 'X-Auth-Token: my_authentication_token' \\   https://api.equinix.com/metal/v1/user?include=projects ```  The `include` parameter is generally accepted in `GET`, `POST`, `PUT`, and `PATCH` requests where `href` resources are presented.  To have multiple resources include, use a comma-separated list (e.g. `?include=emails,projects,memberships`).  ```sh curl -H 'X-Auth-Token: my_authentication_token' \\   https://api.equinix.com/metal/v1/user?include=emails,projects,memberships ```  You may also include nested associations up to three levels deep using dot notation (`?include=memberships.projects`):  ```sh curl -H 'X-Auth-Token: my_authentication_token' \\   https://api.equinix.com/metal/v1/user?include=memberships.projects ```  To exclude resources, and optimize response delivery, use the `exclude` query parameter. The `exclude` parameter is generally accepted in `GET`, `POST`, `PUT`, and `PATCH` requests for fields with nested object responses. When excluded, these fields will be replaced with an object that contains only an `href` field.

API version: 1.0.0
Contact: support@equinixmetal.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1

import (
	"encoding/json"
	"time"
)

// checks if the SpotMarketRequestCreateInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SpotMarketRequestCreateInput{}

// SpotMarketRequestCreateInput struct for SpotMarketRequestCreateInput
type SpotMarketRequestCreateInput struct {
	DevicesMax         *int32                                          `json:"devices_max,omitempty"`
	DevicesMin         *int32                                          `json:"devices_min,omitempty"`
	EndAt              *time.Time                                      `json:"end_at,omitempty"`
	Facilities         []string                                        `json:"facilities,omitempty"`
	InstanceAttributes *SpotMarketRequestCreateInputInstanceAttributes `json:"instance_attributes,omitempty"`
	MaxBidPrice        *float32                                        `json:"max_bid_price,omitempty"`
	// The metro ID or code the spot market request will be created in.
	Metro *string `json:"metro,omitempty"`
}

// NewSpotMarketRequestCreateInput instantiates a new SpotMarketRequestCreateInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSpotMarketRequestCreateInput() *SpotMarketRequestCreateInput {
	this := SpotMarketRequestCreateInput{}
	return &this
}

// NewSpotMarketRequestCreateInputWithDefaults instantiates a new SpotMarketRequestCreateInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSpotMarketRequestCreateInputWithDefaults() *SpotMarketRequestCreateInput {
	this := SpotMarketRequestCreateInput{}
	return &this
}

// GetDevicesMax returns the DevicesMax field value if set, zero value otherwise.
func (o *SpotMarketRequestCreateInput) GetDevicesMax() int32 {
	if o == nil || isNil(o.DevicesMax) {
		var ret int32
		return ret
	}
	return *o.DevicesMax
}

// GetDevicesMaxOk returns a tuple with the DevicesMax field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotMarketRequestCreateInput) GetDevicesMaxOk() (*int32, bool) {
	if o == nil || isNil(o.DevicesMax) {
		return nil, false
	}
	return o.DevicesMax, true
}

// HasDevicesMax returns a boolean if a field has been set.
func (o *SpotMarketRequestCreateInput) HasDevicesMax() bool {
	if o != nil && !isNil(o.DevicesMax) {
		return true
	}

	return false
}

// SetDevicesMax gets a reference to the given int32 and assigns it to the DevicesMax field.
func (o *SpotMarketRequestCreateInput) SetDevicesMax(v int32) {
	o.DevicesMax = &v
}

// GetDevicesMin returns the DevicesMin field value if set, zero value otherwise.
func (o *SpotMarketRequestCreateInput) GetDevicesMin() int32 {
	if o == nil || isNil(o.DevicesMin) {
		var ret int32
		return ret
	}
	return *o.DevicesMin
}

// GetDevicesMinOk returns a tuple with the DevicesMin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotMarketRequestCreateInput) GetDevicesMinOk() (*int32, bool) {
	if o == nil || isNil(o.DevicesMin) {
		return nil, false
	}
	return o.DevicesMin, true
}

// HasDevicesMin returns a boolean if a field has been set.
func (o *SpotMarketRequestCreateInput) HasDevicesMin() bool {
	if o != nil && !isNil(o.DevicesMin) {
		return true
	}

	return false
}

// SetDevicesMin gets a reference to the given int32 and assigns it to the DevicesMin field.
func (o *SpotMarketRequestCreateInput) SetDevicesMin(v int32) {
	o.DevicesMin = &v
}

// GetEndAt returns the EndAt field value if set, zero value otherwise.
func (o *SpotMarketRequestCreateInput) GetEndAt() time.Time {
	if o == nil || isNil(o.EndAt) {
		var ret time.Time
		return ret
	}
	return *o.EndAt
}

// GetEndAtOk returns a tuple with the EndAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotMarketRequestCreateInput) GetEndAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.EndAt) {
		return nil, false
	}
	return o.EndAt, true
}

// HasEndAt returns a boolean if a field has been set.
func (o *SpotMarketRequestCreateInput) HasEndAt() bool {
	if o != nil && !isNil(o.EndAt) {
		return true
	}

	return false
}

// SetEndAt gets a reference to the given time.Time and assigns it to the EndAt field.
func (o *SpotMarketRequestCreateInput) SetEndAt(v time.Time) {
	o.EndAt = &v
}

// GetFacilities returns the Facilities field value if set, zero value otherwise.
func (o *SpotMarketRequestCreateInput) GetFacilities() []string {
	if o == nil || isNil(o.Facilities) {
		var ret []string
		return ret
	}
	return o.Facilities
}

// GetFacilitiesOk returns a tuple with the Facilities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotMarketRequestCreateInput) GetFacilitiesOk() ([]string, bool) {
	if o == nil || isNil(o.Facilities) {
		return nil, false
	}
	return o.Facilities, true
}

// HasFacilities returns a boolean if a field has been set.
func (o *SpotMarketRequestCreateInput) HasFacilities() bool {
	if o != nil && !isNil(o.Facilities) {
		return true
	}

	return false
}

// SetFacilities gets a reference to the given []string and assigns it to the Facilities field.
func (o *SpotMarketRequestCreateInput) SetFacilities(v []string) {
	o.Facilities = v
}

// GetInstanceAttributes returns the InstanceAttributes field value if set, zero value otherwise.
func (o *SpotMarketRequestCreateInput) GetInstanceAttributes() SpotMarketRequestCreateInputInstanceAttributes {
	if o == nil || isNil(o.InstanceAttributes) {
		var ret SpotMarketRequestCreateInputInstanceAttributes
		return ret
	}
	return *o.InstanceAttributes
}

// GetInstanceAttributesOk returns a tuple with the InstanceAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotMarketRequestCreateInput) GetInstanceAttributesOk() (*SpotMarketRequestCreateInputInstanceAttributes, bool) {
	if o == nil || isNil(o.InstanceAttributes) {
		return nil, false
	}
	return o.InstanceAttributes, true
}

// HasInstanceAttributes returns a boolean if a field has been set.
func (o *SpotMarketRequestCreateInput) HasInstanceAttributes() bool {
	if o != nil && !isNil(o.InstanceAttributes) {
		return true
	}

	return false
}

// SetInstanceAttributes gets a reference to the given SpotMarketRequestCreateInputInstanceAttributes and assigns it to the InstanceAttributes field.
func (o *SpotMarketRequestCreateInput) SetInstanceAttributes(v SpotMarketRequestCreateInputInstanceAttributes) {
	o.InstanceAttributes = &v
}

// GetMaxBidPrice returns the MaxBidPrice field value if set, zero value otherwise.
func (o *SpotMarketRequestCreateInput) GetMaxBidPrice() float32 {
	if o == nil || isNil(o.MaxBidPrice) {
		var ret float32
		return ret
	}
	return *o.MaxBidPrice
}

// GetMaxBidPriceOk returns a tuple with the MaxBidPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotMarketRequestCreateInput) GetMaxBidPriceOk() (*float32, bool) {
	if o == nil || isNil(o.MaxBidPrice) {
		return nil, false
	}
	return o.MaxBidPrice, true
}

// HasMaxBidPrice returns a boolean if a field has been set.
func (o *SpotMarketRequestCreateInput) HasMaxBidPrice() bool {
	if o != nil && !isNil(o.MaxBidPrice) {
		return true
	}

	return false
}

// SetMaxBidPrice gets a reference to the given float32 and assigns it to the MaxBidPrice field.
func (o *SpotMarketRequestCreateInput) SetMaxBidPrice(v float32) {
	o.MaxBidPrice = &v
}

// GetMetro returns the Metro field value if set, zero value otherwise.
func (o *SpotMarketRequestCreateInput) GetMetro() string {
	if o == nil || isNil(o.Metro) {
		var ret string
		return ret
	}
	return *o.Metro
}

// GetMetroOk returns a tuple with the Metro field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotMarketRequestCreateInput) GetMetroOk() (*string, bool) {
	if o == nil || isNil(o.Metro) {
		return nil, false
	}
	return o.Metro, true
}

// HasMetro returns a boolean if a field has been set.
func (o *SpotMarketRequestCreateInput) HasMetro() bool {
	if o != nil && !isNil(o.Metro) {
		return true
	}

	return false
}

// SetMetro gets a reference to the given string and assigns it to the Metro field.
func (o *SpotMarketRequestCreateInput) SetMetro(v string) {
	o.Metro = &v
}

func (o SpotMarketRequestCreateInput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SpotMarketRequestCreateInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.DevicesMax) {
		toSerialize["devices_max"] = o.DevicesMax
	}
	if !isNil(o.DevicesMin) {
		toSerialize["devices_min"] = o.DevicesMin
	}
	if !isNil(o.EndAt) {
		toSerialize["end_at"] = o.EndAt
	}
	if !isNil(o.Facilities) {
		toSerialize["facilities"] = o.Facilities
	}
	if !isNil(o.InstanceAttributes) {
		toSerialize["instance_attributes"] = o.InstanceAttributes
	}
	if !isNil(o.MaxBidPrice) {
		toSerialize["max_bid_price"] = o.MaxBidPrice
	}
	if !isNil(o.Metro) {
		toSerialize["metro"] = o.Metro
	}
	return toSerialize, nil
}

type NullableSpotMarketRequestCreateInput struct {
	value *SpotMarketRequestCreateInput
	isSet bool
}

func (v NullableSpotMarketRequestCreateInput) Get() *SpotMarketRequestCreateInput {
	return v.value
}

func (v *NullableSpotMarketRequestCreateInput) Set(val *SpotMarketRequestCreateInput) {
	v.value = val
	v.isSet = true
}

func (v NullableSpotMarketRequestCreateInput) IsSet() bool {
	return v.isSet
}

func (v *NullableSpotMarketRequestCreateInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpotMarketRequestCreateInput(val *SpotMarketRequestCreateInput) *NullableSpotMarketRequestCreateInput {
	return &NullableSpotMarketRequestCreateInput{value: val, isSet: true}
}

func (v NullableSpotMarketRequestCreateInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpotMarketRequestCreateInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
