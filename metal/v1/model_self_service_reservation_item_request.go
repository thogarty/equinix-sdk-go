/*
Metal API

# Introduction Equinix Metal provides a RESTful HTTP API which can be reached at <https://api.equinix.com/metal/v1>. This document describes the API and how to use it.  The API allows you to programmatically interact with all of your Equinix Metal resources, including devices, networks, addresses, organizations, projects, and your user account. Every feature of the Equinix Metal web interface is accessible through the API.  The API docs are generated from the Equinix Metal OpenAPI specification and are officially hosted at <https://metal.equinix.com/developers/api>.  # Common Parameters  The Equinix Metal API uses a few methods to minimize network traffic and improve throughput. These parameters are not used in all API calls, but are used often enough to warrant their own section. Look for these parameters in the documentation for the API calls that support them.  ## Pagination  Pagination is used to limit the number of results returned in a single request. The API will return a maximum of 100 results per page. To retrieve additional results, you can use the `page` and `per_page` query parameters.  The `page` parameter is used to specify the page number. The first page is `1`. The `per_page` parameter is used to specify the number of results per page. The maximum number of results differs by resource type.  ## Sorting  Where offered, the API allows you to sort results by a specific field. To sort results use the `sort_by` query parameter with the root level field name as the value. The `sort_direction` parameter is used to specify the sort direction, either either `asc` (ascending) or `desc` (descending).  ## Filtering  Filtering is used to limit the results returned in a single request. The API supports filtering by certain fields in the response. To filter results, you can use the field as a query parameter.  For example, to filter the IP list to only return public IPv4 addresses, you can filter by the `type` field, as in the following request:  ```sh curl -H 'X-Auth-Token: my_authentication_token' \\   https://api.equinix.com/metal/v1/projects/id/ips?type=public_ipv4 ```  Only IP addresses with the `type` field set to `public_ipv4` will be returned.  ## Searching  Searching is used to find matching resources using multiple field comparissons. The API supports searching in resources that define this behavior. Currently the search parameter is only available on devices, ssh_keys, api_keys and memberships endpoints.  To search resources you can use the `search` query parameter.  ## Include and Exclude  For resources that contain references to other resources, sucha as a Device that refers to the Project it resides in, the Equinix Metal API will returns `href` values (API links) to the associated resource.  ```json {   ...   \"project\": {     \"href\": \"/metal/v1/projects/f3f131c8-f302-49ef-8c44-9405022dc6dd\"   } } ```  If you're going need the project details, you can avoid a second API request.  Specify the contained `href` resources and collections that you'd like to have included in the response using the `include` query parameter.  For example:  ```sh curl -H 'X-Auth-Token: my_authentication_token' \\   https://api.equinix.com/metal/v1/user?include=projects ```  The `include` parameter is generally accepted in `GET`, `POST`, `PUT`, and `PATCH` requests where `href` resources are presented.  To have multiple resources include, use a comma-separated list (e.g. `?include=emails,projects,memberships`).  ```sh curl -H 'X-Auth-Token: my_authentication_token' \\   https://api.equinix.com/metal/v1/user?include=emails,projects,memberships ```  You may also include nested associations up to three levels deep using dot notation (`?include=memberships.projects`):  ```sh curl -H 'X-Auth-Token: my_authentication_token' \\   https://api.equinix.com/metal/v1/user?include=memberships.projects ```  To exclude resources, and optimize response delivery, use the `exclude` query parameter. The `exclude` parameter is generally accepted in `GET`, `POST`, `PUT`, and `PATCH` requests for fields with nested object responses. When excluded, these fields will be replaced with an object that contains only an `href` field.

API version: 1.0.0
Contact: support@equinixmetal.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1

import (
	"encoding/json"
)

// checks if the SelfServiceReservationItemRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SelfServiceReservationItemRequest{}

// SelfServiceReservationItemRequest struct for SelfServiceReservationItemRequest
type SelfServiceReservationItemRequest struct {
	Amount               *float32 `json:"amount,omitempty"`
	MetroId              *string  `json:"metro_id,omitempty"`
	PlanId               *string  `json:"plan_id,omitempty"`
	Quantity             *int32   `json:"quantity,omitempty"`
	Term                 *string  `json:"term,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SelfServiceReservationItemRequest SelfServiceReservationItemRequest

// NewSelfServiceReservationItemRequest instantiates a new SelfServiceReservationItemRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSelfServiceReservationItemRequest() *SelfServiceReservationItemRequest {
	this := SelfServiceReservationItemRequest{}
	return &this
}

// NewSelfServiceReservationItemRequestWithDefaults instantiates a new SelfServiceReservationItemRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSelfServiceReservationItemRequestWithDefaults() *SelfServiceReservationItemRequest {
	this := SelfServiceReservationItemRequest{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *SelfServiceReservationItemRequest) GetAmount() float32 {
	if o == nil || IsNil(o.Amount) {
		var ret float32
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelfServiceReservationItemRequest) GetAmountOk() (*float32, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *SelfServiceReservationItemRequest) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given float32 and assigns it to the Amount field.
func (o *SelfServiceReservationItemRequest) SetAmount(v float32) {
	o.Amount = &v
}

// GetMetroId returns the MetroId field value if set, zero value otherwise.
func (o *SelfServiceReservationItemRequest) GetMetroId() string {
	if o == nil || IsNil(o.MetroId) {
		var ret string
		return ret
	}
	return *o.MetroId
}

// GetMetroIdOk returns a tuple with the MetroId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelfServiceReservationItemRequest) GetMetroIdOk() (*string, bool) {
	if o == nil || IsNil(o.MetroId) {
		return nil, false
	}
	return o.MetroId, true
}

// HasMetroId returns a boolean if a field has been set.
func (o *SelfServiceReservationItemRequest) HasMetroId() bool {
	if o != nil && !IsNil(o.MetroId) {
		return true
	}

	return false
}

// SetMetroId gets a reference to the given string and assigns it to the MetroId field.
func (o *SelfServiceReservationItemRequest) SetMetroId(v string) {
	o.MetroId = &v
}

// GetPlanId returns the PlanId field value if set, zero value otherwise.
func (o *SelfServiceReservationItemRequest) GetPlanId() string {
	if o == nil || IsNil(o.PlanId) {
		var ret string
		return ret
	}
	return *o.PlanId
}

// GetPlanIdOk returns a tuple with the PlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelfServiceReservationItemRequest) GetPlanIdOk() (*string, bool) {
	if o == nil || IsNil(o.PlanId) {
		return nil, false
	}
	return o.PlanId, true
}

// HasPlanId returns a boolean if a field has been set.
func (o *SelfServiceReservationItemRequest) HasPlanId() bool {
	if o != nil && !IsNil(o.PlanId) {
		return true
	}

	return false
}

// SetPlanId gets a reference to the given string and assigns it to the PlanId field.
func (o *SelfServiceReservationItemRequest) SetPlanId(v string) {
	o.PlanId = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *SelfServiceReservationItemRequest) GetQuantity() int32 {
	if o == nil || IsNil(o.Quantity) {
		var ret int32
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelfServiceReservationItemRequest) GetQuantityOk() (*int32, bool) {
	if o == nil || IsNil(o.Quantity) {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *SelfServiceReservationItemRequest) HasQuantity() bool {
	if o != nil && !IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given int32 and assigns it to the Quantity field.
func (o *SelfServiceReservationItemRequest) SetQuantity(v int32) {
	o.Quantity = &v
}

// GetTerm returns the Term field value if set, zero value otherwise.
func (o *SelfServiceReservationItemRequest) GetTerm() string {
	if o == nil || IsNil(o.Term) {
		var ret string
		return ret
	}
	return *o.Term
}

// GetTermOk returns a tuple with the Term field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelfServiceReservationItemRequest) GetTermOk() (*string, bool) {
	if o == nil || IsNil(o.Term) {
		return nil, false
	}
	return o.Term, true
}

// HasTerm returns a boolean if a field has been set.
func (o *SelfServiceReservationItemRequest) HasTerm() bool {
	if o != nil && !IsNil(o.Term) {
		return true
	}

	return false
}

// SetTerm gets a reference to the given string and assigns it to the Term field.
func (o *SelfServiceReservationItemRequest) SetTerm(v string) {
	o.Term = &v
}

func (o SelfServiceReservationItemRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SelfServiceReservationItemRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !IsNil(o.MetroId) {
		toSerialize["metro_id"] = o.MetroId
	}
	if !IsNil(o.PlanId) {
		toSerialize["plan_id"] = o.PlanId
	}
	if !IsNil(o.Quantity) {
		toSerialize["quantity"] = o.Quantity
	}
	if !IsNil(o.Term) {
		toSerialize["term"] = o.Term
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SelfServiceReservationItemRequest) UnmarshalJSON(bytes []byte) (err error) {
	varSelfServiceReservationItemRequest := _SelfServiceReservationItemRequest{}

	if err = json.Unmarshal(bytes, &varSelfServiceReservationItemRequest); err == nil {
		*o = SelfServiceReservationItemRequest(varSelfServiceReservationItemRequest)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "amount")
		delete(additionalProperties, "metro_id")
		delete(additionalProperties, "plan_id")
		delete(additionalProperties, "quantity")
		delete(additionalProperties, "term")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSelfServiceReservationItemRequest struct {
	value *SelfServiceReservationItemRequest
	isSet bool
}

func (v NullableSelfServiceReservationItemRequest) Get() *SelfServiceReservationItemRequest {
	return v.value
}

func (v *NullableSelfServiceReservationItemRequest) Set(val *SelfServiceReservationItemRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSelfServiceReservationItemRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSelfServiceReservationItemRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSelfServiceReservationItemRequest(val *SelfServiceReservationItemRequest) *NullableSelfServiceReservationItemRequest {
	return &NullableSelfServiceReservationItemRequest{value: val, isSet: true}
}

func (v NullableSelfServiceReservationItemRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSelfServiceReservationItemRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
