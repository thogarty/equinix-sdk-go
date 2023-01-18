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

// checks if the HardwareReservation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HardwareReservation{}

// HardwareReservation struct for HardwareReservation
type HardwareReservation struct {
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// Amount that will be charged for every billing_cycle.
	CustomRate *float32  `json:"custom_rate,omitempty"`
	Device     *Device   `json:"device,omitempty"`
	Facility   *Facility `json:"facility,omitempty"`
	Href       *string   `json:"href,omitempty"`
	Id         *string   `json:"id,omitempty"`
	// Whether this Device requires assistance from Metal Equinix.
	NeedOfService *bool    `json:"need_of_service,omitempty"`
	Plan          *Plan    `json:"plan,omitempty"`
	Project       *Project `json:"project,omitempty"`
	// Whether the reserved server is provisionable or not. Spare devices can't be provisioned unless they are activated first.
	Provisionable *bool `json:"provisionable,omitempty"`
	// Short version of the ID.
	ShortId *string `json:"short_id,omitempty"`
	// Whether the Hardware Reservation is a spare. Spare Hardware Reservations are used when a Hardware Reservations requires service from Metal Equinix
	Spare *bool `json:"spare,omitempty"`
	// Switch short id. This can be used to determine if two devices are connected to the same switch, for example.
	SwitchUuid *string `json:"switch_uuid,omitempty"`
}

// NewHardwareReservation instantiates a new HardwareReservation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHardwareReservation() *HardwareReservation {
	this := HardwareReservation{}
	return &this
}

// NewHardwareReservationWithDefaults instantiates a new HardwareReservation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHardwareReservationWithDefaults() *HardwareReservation {
	this := HardwareReservation{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *HardwareReservation) GetCreatedAt() time.Time {
	if o == nil || isNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HardwareReservation) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *HardwareReservation) HasCreatedAt() bool {
	if o != nil && !isNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *HardwareReservation) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetCustomRate returns the CustomRate field value if set, zero value otherwise.
func (o *HardwareReservation) GetCustomRate() float32 {
	if o == nil || isNil(o.CustomRate) {
		var ret float32
		return ret
	}
	return *o.CustomRate
}

// GetCustomRateOk returns a tuple with the CustomRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HardwareReservation) GetCustomRateOk() (*float32, bool) {
	if o == nil || isNil(o.CustomRate) {
		return nil, false
	}
	return o.CustomRate, true
}

// HasCustomRate returns a boolean if a field has been set.
func (o *HardwareReservation) HasCustomRate() bool {
	if o != nil && !isNil(o.CustomRate) {
		return true
	}

	return false
}

// SetCustomRate gets a reference to the given float32 and assigns it to the CustomRate field.
func (o *HardwareReservation) SetCustomRate(v float32) {
	o.CustomRate = &v
}

// GetDevice returns the Device field value if set, zero value otherwise.
func (o *HardwareReservation) GetDevice() Device {
	if o == nil || isNil(o.Device) {
		var ret Device
		return ret
	}
	return *o.Device
}

// GetDeviceOk returns a tuple with the Device field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HardwareReservation) GetDeviceOk() (*Device, bool) {
	if o == nil || isNil(o.Device) {
		return nil, false
	}
	return o.Device, true
}

// HasDevice returns a boolean if a field has been set.
func (o *HardwareReservation) HasDevice() bool {
	if o != nil && !isNil(o.Device) {
		return true
	}

	return false
}

// SetDevice gets a reference to the given Device and assigns it to the Device field.
func (o *HardwareReservation) SetDevice(v Device) {
	o.Device = &v
}

// GetFacility returns the Facility field value if set, zero value otherwise.
func (o *HardwareReservation) GetFacility() Facility {
	if o == nil || isNil(o.Facility) {
		var ret Facility
		return ret
	}
	return *o.Facility
}

// GetFacilityOk returns a tuple with the Facility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HardwareReservation) GetFacilityOk() (*Facility, bool) {
	if o == nil || isNil(o.Facility) {
		return nil, false
	}
	return o.Facility, true
}

// HasFacility returns a boolean if a field has been set.
func (o *HardwareReservation) HasFacility() bool {
	if o != nil && !isNil(o.Facility) {
		return true
	}

	return false
}

// SetFacility gets a reference to the given Facility and assigns it to the Facility field.
func (o *HardwareReservation) SetFacility(v Facility) {
	o.Facility = &v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *HardwareReservation) GetHref() string {
	if o == nil || isNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HardwareReservation) GetHrefOk() (*string, bool) {
	if o == nil || isNil(o.Href) {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *HardwareReservation) HasHref() bool {
	if o != nil && !isNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *HardwareReservation) SetHref(v string) {
	o.Href = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *HardwareReservation) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HardwareReservation) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *HardwareReservation) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *HardwareReservation) SetId(v string) {
	o.Id = &v
}

// GetNeedOfService returns the NeedOfService field value if set, zero value otherwise.
func (o *HardwareReservation) GetNeedOfService() bool {
	if o == nil || isNil(o.NeedOfService) {
		var ret bool
		return ret
	}
	return *o.NeedOfService
}

// GetNeedOfServiceOk returns a tuple with the NeedOfService field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HardwareReservation) GetNeedOfServiceOk() (*bool, bool) {
	if o == nil || isNil(o.NeedOfService) {
		return nil, false
	}
	return o.NeedOfService, true
}

// HasNeedOfService returns a boolean if a field has been set.
func (o *HardwareReservation) HasNeedOfService() bool {
	if o != nil && !isNil(o.NeedOfService) {
		return true
	}

	return false
}

// SetNeedOfService gets a reference to the given bool and assigns it to the NeedOfService field.
func (o *HardwareReservation) SetNeedOfService(v bool) {
	o.NeedOfService = &v
}

// GetPlan returns the Plan field value if set, zero value otherwise.
func (o *HardwareReservation) GetPlan() Plan {
	if o == nil || isNil(o.Plan) {
		var ret Plan
		return ret
	}
	return *o.Plan
}

// GetPlanOk returns a tuple with the Plan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HardwareReservation) GetPlanOk() (*Plan, bool) {
	if o == nil || isNil(o.Plan) {
		return nil, false
	}
	return o.Plan, true
}

// HasPlan returns a boolean if a field has been set.
func (o *HardwareReservation) HasPlan() bool {
	if o != nil && !isNil(o.Plan) {
		return true
	}

	return false
}

// SetPlan gets a reference to the given Plan and assigns it to the Plan field.
func (o *HardwareReservation) SetPlan(v Plan) {
	o.Plan = &v
}

// GetProject returns the Project field value if set, zero value otherwise.
func (o *HardwareReservation) GetProject() Project {
	if o == nil || isNil(o.Project) {
		var ret Project
		return ret
	}
	return *o.Project
}

// GetProjectOk returns a tuple with the Project field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HardwareReservation) GetProjectOk() (*Project, bool) {
	if o == nil || isNil(o.Project) {
		return nil, false
	}
	return o.Project, true
}

// HasProject returns a boolean if a field has been set.
func (o *HardwareReservation) HasProject() bool {
	if o != nil && !isNil(o.Project) {
		return true
	}

	return false
}

// SetProject gets a reference to the given Project and assigns it to the Project field.
func (o *HardwareReservation) SetProject(v Project) {
	o.Project = &v
}

// GetProvisionable returns the Provisionable field value if set, zero value otherwise.
func (o *HardwareReservation) GetProvisionable() bool {
	if o == nil || isNil(o.Provisionable) {
		var ret bool
		return ret
	}
	return *o.Provisionable
}

// GetProvisionableOk returns a tuple with the Provisionable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HardwareReservation) GetProvisionableOk() (*bool, bool) {
	if o == nil || isNil(o.Provisionable) {
		return nil, false
	}
	return o.Provisionable, true
}

// HasProvisionable returns a boolean if a field has been set.
func (o *HardwareReservation) HasProvisionable() bool {
	if o != nil && !isNil(o.Provisionable) {
		return true
	}

	return false
}

// SetProvisionable gets a reference to the given bool and assigns it to the Provisionable field.
func (o *HardwareReservation) SetProvisionable(v bool) {
	o.Provisionable = &v
}

// GetShortId returns the ShortId field value if set, zero value otherwise.
func (o *HardwareReservation) GetShortId() string {
	if o == nil || isNil(o.ShortId) {
		var ret string
		return ret
	}
	return *o.ShortId
}

// GetShortIdOk returns a tuple with the ShortId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HardwareReservation) GetShortIdOk() (*string, bool) {
	if o == nil || isNil(o.ShortId) {
		return nil, false
	}
	return o.ShortId, true
}

// HasShortId returns a boolean if a field has been set.
func (o *HardwareReservation) HasShortId() bool {
	if o != nil && !isNil(o.ShortId) {
		return true
	}

	return false
}

// SetShortId gets a reference to the given string and assigns it to the ShortId field.
func (o *HardwareReservation) SetShortId(v string) {
	o.ShortId = &v
}

// GetSpare returns the Spare field value if set, zero value otherwise.
func (o *HardwareReservation) GetSpare() bool {
	if o == nil || isNil(o.Spare) {
		var ret bool
		return ret
	}
	return *o.Spare
}

// GetSpareOk returns a tuple with the Spare field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HardwareReservation) GetSpareOk() (*bool, bool) {
	if o == nil || isNil(o.Spare) {
		return nil, false
	}
	return o.Spare, true
}

// HasSpare returns a boolean if a field has been set.
func (o *HardwareReservation) HasSpare() bool {
	if o != nil && !isNil(o.Spare) {
		return true
	}

	return false
}

// SetSpare gets a reference to the given bool and assigns it to the Spare field.
func (o *HardwareReservation) SetSpare(v bool) {
	o.Spare = &v
}

// GetSwitchUuid returns the SwitchUuid field value if set, zero value otherwise.
func (o *HardwareReservation) GetSwitchUuid() string {
	if o == nil || isNil(o.SwitchUuid) {
		var ret string
		return ret
	}
	return *o.SwitchUuid
}

// GetSwitchUuidOk returns a tuple with the SwitchUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HardwareReservation) GetSwitchUuidOk() (*string, bool) {
	if o == nil || isNil(o.SwitchUuid) {
		return nil, false
	}
	return o.SwitchUuid, true
}

// HasSwitchUuid returns a boolean if a field has been set.
func (o *HardwareReservation) HasSwitchUuid() bool {
	if o != nil && !isNil(o.SwitchUuid) {
		return true
	}

	return false
}

// SetSwitchUuid gets a reference to the given string and assigns it to the SwitchUuid field.
func (o *HardwareReservation) SetSwitchUuid(v string) {
	o.SwitchUuid = &v
}

func (o HardwareReservation) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HardwareReservation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !isNil(o.CustomRate) {
		toSerialize["custom_rate"] = o.CustomRate
	}
	if !isNil(o.Device) {
		toSerialize["device"] = o.Device
	}
	if !isNil(o.Facility) {
		toSerialize["facility"] = o.Facility
	}
	if !isNil(o.Href) {
		toSerialize["href"] = o.Href
	}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.NeedOfService) {
		toSerialize["need_of_service"] = o.NeedOfService
	}
	if !isNil(o.Plan) {
		toSerialize["plan"] = o.Plan
	}
	if !isNil(o.Project) {
		toSerialize["project"] = o.Project
	}
	if !isNil(o.Provisionable) {
		toSerialize["provisionable"] = o.Provisionable
	}
	if !isNil(o.ShortId) {
		toSerialize["short_id"] = o.ShortId
	}
	if !isNil(o.Spare) {
		toSerialize["spare"] = o.Spare
	}
	if !isNil(o.SwitchUuid) {
		toSerialize["switch_uuid"] = o.SwitchUuid
	}
	return toSerialize, nil
}

type NullableHardwareReservation struct {
	value *HardwareReservation
	isSet bool
}

func (v NullableHardwareReservation) Get() *HardwareReservation {
	return v.value
}

func (v *NullableHardwareReservation) Set(val *HardwareReservation) {
	v.value = val
	v.isSet = true
}

func (v NullableHardwareReservation) IsSet() bool {
	return v.isSet
}

func (v *NullableHardwareReservation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHardwareReservation(val *HardwareReservation) *NullableHardwareReservation {
	return &NullableHardwareReservation{value: val, isSet: true}
}

func (v NullableHardwareReservation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHardwareReservation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
