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

// checks if the SelfServiceReservationResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SelfServiceReservationResponse{}

// SelfServiceReservationResponse struct for SelfServiceReservationResponse
type SelfServiceReservationResponse struct {
	CreatedAt      *time.Time                                 `json:"created_at,omitempty"`
	Item           []SelfServiceReservationItemResponse       `json:"item,omitempty"`
	Notes          *string                                    `json:"notes,omitempty"`
	Organization   *string                                    `json:"organization,omitempty"`
	OrganizationId *string                                    `json:"organization_id,omitempty"`
	Period         *CreateSelfServiceReservationRequestPeriod `json:"period,omitempty"`
	Project        *string                                    `json:"project,omitempty"`
	ProjectId      *string                                    `json:"project_id,omitempty"`
	StartDate      *time.Time                                 `json:"start_date,omitempty"`
	Status         *string                                    `json:"status,omitempty"`
	TotalCost      *int32                                     `json:"total_cost,omitempty"`
}

// NewSelfServiceReservationResponse instantiates a new SelfServiceReservationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSelfServiceReservationResponse() *SelfServiceReservationResponse {
	this := SelfServiceReservationResponse{}
	return &this
}

// NewSelfServiceReservationResponseWithDefaults instantiates a new SelfServiceReservationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSelfServiceReservationResponseWithDefaults() *SelfServiceReservationResponse {
	this := SelfServiceReservationResponse{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *SelfServiceReservationResponse) GetCreatedAt() time.Time {
	if o == nil || isNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelfServiceReservationResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *SelfServiceReservationResponse) HasCreatedAt() bool {
	if o != nil && !isNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *SelfServiceReservationResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetItem returns the Item field value if set, zero value otherwise.
func (o *SelfServiceReservationResponse) GetItem() []SelfServiceReservationItemResponse {
	if o == nil || isNil(o.Item) {
		var ret []SelfServiceReservationItemResponse
		return ret
	}
	return o.Item
}

// GetItemOk returns a tuple with the Item field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelfServiceReservationResponse) GetItemOk() ([]SelfServiceReservationItemResponse, bool) {
	if o == nil || isNil(o.Item) {
		return nil, false
	}
	return o.Item, true
}

// HasItem returns a boolean if a field has been set.
func (o *SelfServiceReservationResponse) HasItem() bool {
	if o != nil && !isNil(o.Item) {
		return true
	}

	return false
}

// SetItem gets a reference to the given []SelfServiceReservationItemResponse and assigns it to the Item field.
func (o *SelfServiceReservationResponse) SetItem(v []SelfServiceReservationItemResponse) {
	o.Item = v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *SelfServiceReservationResponse) GetNotes() string {
	if o == nil || isNil(o.Notes) {
		var ret string
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelfServiceReservationResponse) GetNotesOk() (*string, bool) {
	if o == nil || isNil(o.Notes) {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *SelfServiceReservationResponse) HasNotes() bool {
	if o != nil && !isNil(o.Notes) {
		return true
	}

	return false
}

// SetNotes gets a reference to the given string and assigns it to the Notes field.
func (o *SelfServiceReservationResponse) SetNotes(v string) {
	o.Notes = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *SelfServiceReservationResponse) GetOrganization() string {
	if o == nil || isNil(o.Organization) {
		var ret string
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelfServiceReservationResponse) GetOrganizationOk() (*string, bool) {
	if o == nil || isNil(o.Organization) {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *SelfServiceReservationResponse) HasOrganization() bool {
	if o != nil && !isNil(o.Organization) {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given string and assigns it to the Organization field.
func (o *SelfServiceReservationResponse) SetOrganization(v string) {
	o.Organization = &v
}

// GetOrganizationId returns the OrganizationId field value if set, zero value otherwise.
func (o *SelfServiceReservationResponse) GetOrganizationId() string {
	if o == nil || isNil(o.OrganizationId) {
		var ret string
		return ret
	}
	return *o.OrganizationId
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelfServiceReservationResponse) GetOrganizationIdOk() (*string, bool) {
	if o == nil || isNil(o.OrganizationId) {
		return nil, false
	}
	return o.OrganizationId, true
}

// HasOrganizationId returns a boolean if a field has been set.
func (o *SelfServiceReservationResponse) HasOrganizationId() bool {
	if o != nil && !isNil(o.OrganizationId) {
		return true
	}

	return false
}

// SetOrganizationId gets a reference to the given string and assigns it to the OrganizationId field.
func (o *SelfServiceReservationResponse) SetOrganizationId(v string) {
	o.OrganizationId = &v
}

// GetPeriod returns the Period field value if set, zero value otherwise.
func (o *SelfServiceReservationResponse) GetPeriod() CreateSelfServiceReservationRequestPeriod {
	if o == nil || isNil(o.Period) {
		var ret CreateSelfServiceReservationRequestPeriod
		return ret
	}
	return *o.Period
}

// GetPeriodOk returns a tuple with the Period field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelfServiceReservationResponse) GetPeriodOk() (*CreateSelfServiceReservationRequestPeriod, bool) {
	if o == nil || isNil(o.Period) {
		return nil, false
	}
	return o.Period, true
}

// HasPeriod returns a boolean if a field has been set.
func (o *SelfServiceReservationResponse) HasPeriod() bool {
	if o != nil && !isNil(o.Period) {
		return true
	}

	return false
}

// SetPeriod gets a reference to the given CreateSelfServiceReservationRequestPeriod and assigns it to the Period field.
func (o *SelfServiceReservationResponse) SetPeriod(v CreateSelfServiceReservationRequestPeriod) {
	o.Period = &v
}

// GetProject returns the Project field value if set, zero value otherwise.
func (o *SelfServiceReservationResponse) GetProject() string {
	if o == nil || isNil(o.Project) {
		var ret string
		return ret
	}
	return *o.Project
}

// GetProjectOk returns a tuple with the Project field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelfServiceReservationResponse) GetProjectOk() (*string, bool) {
	if o == nil || isNil(o.Project) {
		return nil, false
	}
	return o.Project, true
}

// HasProject returns a boolean if a field has been set.
func (o *SelfServiceReservationResponse) HasProject() bool {
	if o != nil && !isNil(o.Project) {
		return true
	}

	return false
}

// SetProject gets a reference to the given string and assigns it to the Project field.
func (o *SelfServiceReservationResponse) SetProject(v string) {
	o.Project = &v
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *SelfServiceReservationResponse) GetProjectId() string {
	if o == nil || isNil(o.ProjectId) {
		var ret string
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelfServiceReservationResponse) GetProjectIdOk() (*string, bool) {
	if o == nil || isNil(o.ProjectId) {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *SelfServiceReservationResponse) HasProjectId() bool {
	if o != nil && !isNil(o.ProjectId) {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given string and assigns it to the ProjectId field.
func (o *SelfServiceReservationResponse) SetProjectId(v string) {
	o.ProjectId = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *SelfServiceReservationResponse) GetStartDate() time.Time {
	if o == nil || isNil(o.StartDate) {
		var ret time.Time
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelfServiceReservationResponse) GetStartDateOk() (*time.Time, bool) {
	if o == nil || isNil(o.StartDate) {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *SelfServiceReservationResponse) HasStartDate() bool {
	if o != nil && !isNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given time.Time and assigns it to the StartDate field.
func (o *SelfServiceReservationResponse) SetStartDate(v time.Time) {
	o.StartDate = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *SelfServiceReservationResponse) GetStatus() string {
	if o == nil || isNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelfServiceReservationResponse) GetStatusOk() (*string, bool) {
	if o == nil || isNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *SelfServiceReservationResponse) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *SelfServiceReservationResponse) SetStatus(v string) {
	o.Status = &v
}

// GetTotalCost returns the TotalCost field value if set, zero value otherwise.
func (o *SelfServiceReservationResponse) GetTotalCost() int32 {
	if o == nil || isNil(o.TotalCost) {
		var ret int32
		return ret
	}
	return *o.TotalCost
}

// GetTotalCostOk returns a tuple with the TotalCost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelfServiceReservationResponse) GetTotalCostOk() (*int32, bool) {
	if o == nil || isNil(o.TotalCost) {
		return nil, false
	}
	return o.TotalCost, true
}

// HasTotalCost returns a boolean if a field has been set.
func (o *SelfServiceReservationResponse) HasTotalCost() bool {
	if o != nil && !isNil(o.TotalCost) {
		return true
	}

	return false
}

// SetTotalCost gets a reference to the given int32 and assigns it to the TotalCost field.
func (o *SelfServiceReservationResponse) SetTotalCost(v int32) {
	o.TotalCost = &v
}

func (o SelfServiceReservationResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SelfServiceReservationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !isNil(o.Item) {
		toSerialize["item"] = o.Item
	}
	if !isNil(o.Notes) {
		toSerialize["notes"] = o.Notes
	}
	if !isNil(o.Organization) {
		toSerialize["organization"] = o.Organization
	}
	if !isNil(o.OrganizationId) {
		toSerialize["organization_id"] = o.OrganizationId
	}
	if !isNil(o.Period) {
		toSerialize["period"] = o.Period
	}
	if !isNil(o.Project) {
		toSerialize["project"] = o.Project
	}
	if !isNil(o.ProjectId) {
		toSerialize["project_id"] = o.ProjectId
	}
	if !isNil(o.StartDate) {
		toSerialize["start_date"] = o.StartDate
	}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !isNil(o.TotalCost) {
		toSerialize["total_cost"] = o.TotalCost
	}
	return toSerialize, nil
}

type NullableSelfServiceReservationResponse struct {
	value *SelfServiceReservationResponse
	isSet bool
}

func (v NullableSelfServiceReservationResponse) Get() *SelfServiceReservationResponse {
	return v.value
}

func (v *NullableSelfServiceReservationResponse) Set(val *SelfServiceReservationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSelfServiceReservationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSelfServiceReservationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSelfServiceReservationResponse(val *SelfServiceReservationResponse) *NullableSelfServiceReservationResponse {
	return &NullableSelfServiceReservationResponse{value: val, isSet: true}
}

func (v NullableSelfServiceReservationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSelfServiceReservationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
