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
	"time"
)

// checks if the PortVlanAssignmentBatch type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PortVlanAssignmentBatch{}

// PortVlanAssignmentBatch struct for PortVlanAssignmentBatch
type PortVlanAssignmentBatch struct {
	CreatedAt            *time.Time                                    `json:"created_at,omitempty"`
	ErrorMessages        []string                                      `json:"error_messages,omitempty"`
	Id                   *string                                       `json:"id,omitempty"`
	Port                 *Port                                         `json:"port,omitempty"`
	Quantity             *int32                                        `json:"quantity,omitempty"`
	State                *string                                       `json:"state,omitempty"`
	UpdatedAt            *time.Time                                    `json:"updated_at,omitempty"`
	VlanAssignments      []PortVlanAssignmentBatchVlanAssignmentsInner `json:"vlan_assignments,omitempty"`
	Project              *Href                                         `json:"project,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PortVlanAssignmentBatch PortVlanAssignmentBatch

// NewPortVlanAssignmentBatch instantiates a new PortVlanAssignmentBatch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPortVlanAssignmentBatch() *PortVlanAssignmentBatch {
	this := PortVlanAssignmentBatch{}
	return &this
}

// NewPortVlanAssignmentBatchWithDefaults instantiates a new PortVlanAssignmentBatch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPortVlanAssignmentBatchWithDefaults() *PortVlanAssignmentBatch {
	this := PortVlanAssignmentBatch{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *PortVlanAssignmentBatch) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortVlanAssignmentBatch) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *PortVlanAssignmentBatch) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *PortVlanAssignmentBatch) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetErrorMessages returns the ErrorMessages field value if set, zero value otherwise.
func (o *PortVlanAssignmentBatch) GetErrorMessages() []string {
	if o == nil || IsNil(o.ErrorMessages) {
		var ret []string
		return ret
	}
	return o.ErrorMessages
}

// GetErrorMessagesOk returns a tuple with the ErrorMessages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortVlanAssignmentBatch) GetErrorMessagesOk() ([]string, bool) {
	if o == nil || IsNil(o.ErrorMessages) {
		return nil, false
	}
	return o.ErrorMessages, true
}

// HasErrorMessages returns a boolean if a field has been set.
func (o *PortVlanAssignmentBatch) HasErrorMessages() bool {
	if o != nil && !IsNil(o.ErrorMessages) {
		return true
	}

	return false
}

// SetErrorMessages gets a reference to the given []string and assigns it to the ErrorMessages field.
func (o *PortVlanAssignmentBatch) SetErrorMessages(v []string) {
	o.ErrorMessages = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PortVlanAssignmentBatch) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortVlanAssignmentBatch) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PortVlanAssignmentBatch) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PortVlanAssignmentBatch) SetId(v string) {
	o.Id = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *PortVlanAssignmentBatch) GetPort() Port {
	if o == nil || IsNil(o.Port) {
		var ret Port
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortVlanAssignmentBatch) GetPortOk() (*Port, bool) {
	if o == nil || IsNil(o.Port) {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *PortVlanAssignmentBatch) HasPort() bool {
	if o != nil && !IsNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given Port and assigns it to the Port field.
func (o *PortVlanAssignmentBatch) SetPort(v Port) {
	o.Port = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *PortVlanAssignmentBatch) GetQuantity() int32 {
	if o == nil || IsNil(o.Quantity) {
		var ret int32
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortVlanAssignmentBatch) GetQuantityOk() (*int32, bool) {
	if o == nil || IsNil(o.Quantity) {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *PortVlanAssignmentBatch) HasQuantity() bool {
	if o != nil && !IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given int32 and assigns it to the Quantity field.
func (o *PortVlanAssignmentBatch) SetQuantity(v int32) {
	o.Quantity = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *PortVlanAssignmentBatch) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortVlanAssignmentBatch) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *PortVlanAssignmentBatch) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *PortVlanAssignmentBatch) SetState(v string) {
	o.State = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *PortVlanAssignmentBatch) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortVlanAssignmentBatch) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *PortVlanAssignmentBatch) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *PortVlanAssignmentBatch) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetVlanAssignments returns the VlanAssignments field value if set, zero value otherwise.
func (o *PortVlanAssignmentBatch) GetVlanAssignments() []PortVlanAssignmentBatchVlanAssignmentsInner {
	if o == nil || IsNil(o.VlanAssignments) {
		var ret []PortVlanAssignmentBatchVlanAssignmentsInner
		return ret
	}
	return o.VlanAssignments
}

// GetVlanAssignmentsOk returns a tuple with the VlanAssignments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortVlanAssignmentBatch) GetVlanAssignmentsOk() ([]PortVlanAssignmentBatchVlanAssignmentsInner, bool) {
	if o == nil || IsNil(o.VlanAssignments) {
		return nil, false
	}
	return o.VlanAssignments, true
}

// HasVlanAssignments returns a boolean if a field has been set.
func (o *PortVlanAssignmentBatch) HasVlanAssignments() bool {
	if o != nil && !IsNil(o.VlanAssignments) {
		return true
	}

	return false
}

// SetVlanAssignments gets a reference to the given []PortVlanAssignmentBatchVlanAssignmentsInner and assigns it to the VlanAssignments field.
func (o *PortVlanAssignmentBatch) SetVlanAssignments(v []PortVlanAssignmentBatchVlanAssignmentsInner) {
	o.VlanAssignments = v
}

// GetProject returns the Project field value if set, zero value otherwise.
func (o *PortVlanAssignmentBatch) GetProject() Href {
	if o == nil || IsNil(o.Project) {
		var ret Href
		return ret
	}
	return *o.Project
}

// GetProjectOk returns a tuple with the Project field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortVlanAssignmentBatch) GetProjectOk() (*Href, bool) {
	if o == nil || IsNil(o.Project) {
		return nil, false
	}
	return o.Project, true
}

// HasProject returns a boolean if a field has been set.
func (o *PortVlanAssignmentBatch) HasProject() bool {
	if o != nil && !IsNil(o.Project) {
		return true
	}

	return false
}

// SetProject gets a reference to the given Href and assigns it to the Project field.
func (o *PortVlanAssignmentBatch) SetProject(v Href) {
	o.Project = &v
}

func (o PortVlanAssignmentBatch) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PortVlanAssignmentBatch) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.ErrorMessages) {
		toSerialize["error_messages"] = o.ErrorMessages
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Port) {
		toSerialize["port"] = o.Port
	}
	if !IsNil(o.Quantity) {
		toSerialize["quantity"] = o.Quantity
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if !IsNil(o.VlanAssignments) {
		toSerialize["vlan_assignments"] = o.VlanAssignments
	}
	if !IsNil(o.Project) {
		toSerialize["project"] = o.Project
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PortVlanAssignmentBatch) UnmarshalJSON(bytes []byte) (err error) {
	varPortVlanAssignmentBatch := _PortVlanAssignmentBatch{}

	if err = json.Unmarshal(bytes, &varPortVlanAssignmentBatch); err == nil {
		*o = PortVlanAssignmentBatch(varPortVlanAssignmentBatch)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "error_messages")
		delete(additionalProperties, "id")
		delete(additionalProperties, "port")
		delete(additionalProperties, "quantity")
		delete(additionalProperties, "state")
		delete(additionalProperties, "updated_at")
		delete(additionalProperties, "vlan_assignments")
		delete(additionalProperties, "project")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePortVlanAssignmentBatch struct {
	value *PortVlanAssignmentBatch
	isSet bool
}

func (v NullablePortVlanAssignmentBatch) Get() *PortVlanAssignmentBatch {
	return v.value
}

func (v *NullablePortVlanAssignmentBatch) Set(val *PortVlanAssignmentBatch) {
	v.value = val
	v.isSet = true
}

func (v NullablePortVlanAssignmentBatch) IsSet() bool {
	return v.isSet
}

func (v *NullablePortVlanAssignmentBatch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePortVlanAssignmentBatch(val *PortVlanAssignmentBatch) *NullablePortVlanAssignmentBatch {
	return &NullablePortVlanAssignmentBatch{value: val, isSet: true}
}

func (v NullablePortVlanAssignmentBatch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePortVlanAssignmentBatch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
