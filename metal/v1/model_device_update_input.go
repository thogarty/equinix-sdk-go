/*
Metal API

# Introduction Equinix Metal provides a RESTful HTTP API which can be reached at <https://api.equinix.com/metal/v1>. This document describes the API and how to use it.  The API allows you to programmatically interact with all of your Equinix Metal resources, including devices, networks, addresses, organizations, projects, and your user account. Every feature of the Equinix Metal web interface is accessible through the API.  The API docs are generated from the Equinix Metal OpenAPI specification and are officially hosted at <https://metal.equinix.com/developers/api>.  # Common Parameters  The Equinix Metal API uses a few methods to minimize network traffic and improve throughput. These parameters are not used in all API calls, but are used often enough to warrant their own section. Look for these parameters in the documentation for the API calls that support them.  ## Pagination  Pagination is used to limit the number of results returned in a single request. The API will return a maximum of 100 results per page. To retrieve additional results, you can use the `page` and `per_page` query parameters.  The `page` parameter is used to specify the page number. The first page is `1`. The `per_page` parameter is used to specify the number of results per page. The maximum number of results differs by resource type.  ## Sorting  Where offered, the API allows you to sort results by a specific field. To sort results use the `sort_by` query parameter with the root level field name as the value. The `sort_direction` parameter is used to specify the sort direction, either either `asc` (ascending) or `desc` (descending).  ## Filtering  Filtering is used to limit the results returned in a single request. The API supports filtering by certain fields in the response. To filter results, you can use the field as a query parameter.  For example, to filter the IP list to only return public IPv4 addresses, you can filter by the `type` field, as in the following request:  ```sh curl -H 'X-Auth-Token: my_authentication_token' \\   https://api.equinix.com/metal/v1/projects/id/ips?type=public_ipv4 ```  Only IP addresses with the `type` field set to `public_ipv4` will be returned.  ## Searching  Searching is used to find matching resources using multiple field comparissons. The API supports searching in resources that define this behavior. The fields available for search differ by resource, as does the search strategy.  To search resources you can use the `search` query parameter.  ## Include and Exclude  For resources that contain references to other resources, sucha as a Device that refers to the Project it resides in, the Equinix Metal API will returns `href` values (API links) to the associated resource.  ```json {   ...   \"project\": {     \"href\": \"/metal/v1/projects/f3f131c8-f302-49ef-8c44-9405022dc6dd\"   } } ```  If you're going need the project details, you can avoid a second API request.  Specify the contained `href` resources and collections that you'd like to have included in the response using the `include` query parameter.  For example:  ```sh curl -H 'X-Auth-Token: my_authentication_token' \\   https://api.equinix.com/metal/v1/user?include=projects ```  The `include` parameter is generally accepted in `GET`, `POST`, `PUT`, and `PATCH` requests where `href` resources are presented.  To have multiple resources include, use a comma-separated list (e.g. `?include=emails,projects,memberships`).  ```sh curl -H 'X-Auth-Token: my_authentication_token' \\   https://api.equinix.com/metal/v1/user?include=emails,projects,memberships ```  You may also include nested associations up to three levels deep using dot notation (`?include=memberships.projects`):  ```sh curl -H 'X-Auth-Token: my_authentication_token' \\   https://api.equinix.com/metal/v1/user?include=memberships.projects ```  To exclude resources, and optimize response delivery, use the `exclude` query parameter. The `exclude` parameter is generally accepted in `GET`, `POST`, `PUT`, and `PATCH` requests for fields with nested object responses. When excluded, these fields will be replaced with an object that contains only an `href` field.

API version: 1.0.0
Contact: support@equinixmetal.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1

import (
	"encoding/json"
)

// checks if the DeviceUpdateInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeviceUpdateInput{}

// DeviceUpdateInput struct for DeviceUpdateInput
type DeviceUpdateInput struct {
	AlwaysPxe     *bool                  `json:"always_pxe,omitempty"`
	BillingCycle  *string                `json:"billing_cycle,omitempty"`
	Customdata    map[string]interface{} `json:"customdata,omitempty"`
	Description   *string                `json:"description,omitempty"`
	Hostname      *string                `json:"hostname,omitempty"`
	IpxeScriptUrl *string                `json:"ipxe_script_url,omitempty"`
	// Whether the device should be locked, preventing accidental deletion.
	Locked *bool `json:"locked,omitempty"`
	// If true, this instance can not be converted to a different network type.
	NetworkFrozen *bool `json:"network_frozen,omitempty"`
	// Can be set to false to convert a spot-market instance to on-demand.
	SpotInstance         *bool    `json:"spot_instance,omitempty"`
	Tags                 []string `json:"tags,omitempty"`
	Userdata             *string  `json:"userdata,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DeviceUpdateInput DeviceUpdateInput

// NewDeviceUpdateInput instantiates a new DeviceUpdateInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceUpdateInput() *DeviceUpdateInput {
	this := DeviceUpdateInput{}
	return &this
}

// NewDeviceUpdateInputWithDefaults instantiates a new DeviceUpdateInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceUpdateInputWithDefaults() *DeviceUpdateInput {
	this := DeviceUpdateInput{}
	return &this
}

// GetAlwaysPxe returns the AlwaysPxe field value if set, zero value otherwise.
func (o *DeviceUpdateInput) GetAlwaysPxe() bool {
	if o == nil || IsNil(o.AlwaysPxe) {
		var ret bool
		return ret
	}
	return *o.AlwaysPxe
}

// GetAlwaysPxeOk returns a tuple with the AlwaysPxe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceUpdateInput) GetAlwaysPxeOk() (*bool, bool) {
	if o == nil || IsNil(o.AlwaysPxe) {
		return nil, false
	}
	return o.AlwaysPxe, true
}

// HasAlwaysPxe returns a boolean if a field has been set.
func (o *DeviceUpdateInput) HasAlwaysPxe() bool {
	if o != nil && !IsNil(o.AlwaysPxe) {
		return true
	}

	return false
}

// SetAlwaysPxe gets a reference to the given bool and assigns it to the AlwaysPxe field.
func (o *DeviceUpdateInput) SetAlwaysPxe(v bool) {
	o.AlwaysPxe = &v
}

// GetBillingCycle returns the BillingCycle field value if set, zero value otherwise.
func (o *DeviceUpdateInput) GetBillingCycle() string {
	if o == nil || IsNil(o.BillingCycle) {
		var ret string
		return ret
	}
	return *o.BillingCycle
}

// GetBillingCycleOk returns a tuple with the BillingCycle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceUpdateInput) GetBillingCycleOk() (*string, bool) {
	if o == nil || IsNil(o.BillingCycle) {
		return nil, false
	}
	return o.BillingCycle, true
}

// HasBillingCycle returns a boolean if a field has been set.
func (o *DeviceUpdateInput) HasBillingCycle() bool {
	if o != nil && !IsNil(o.BillingCycle) {
		return true
	}

	return false
}

// SetBillingCycle gets a reference to the given string and assigns it to the BillingCycle field.
func (o *DeviceUpdateInput) SetBillingCycle(v string) {
	o.BillingCycle = &v
}

// GetCustomdata returns the Customdata field value if set, zero value otherwise.
func (o *DeviceUpdateInput) GetCustomdata() map[string]interface{} {
	if o == nil || IsNil(o.Customdata) {
		var ret map[string]interface{}
		return ret
	}
	return o.Customdata
}

// GetCustomdataOk returns a tuple with the Customdata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceUpdateInput) GetCustomdataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Customdata) {
		return map[string]interface{}{}, false
	}
	return o.Customdata, true
}

// HasCustomdata returns a boolean if a field has been set.
func (o *DeviceUpdateInput) HasCustomdata() bool {
	if o != nil && !IsNil(o.Customdata) {
		return true
	}

	return false
}

// SetCustomdata gets a reference to the given map[string]interface{} and assigns it to the Customdata field.
func (o *DeviceUpdateInput) SetCustomdata(v map[string]interface{}) {
	o.Customdata = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *DeviceUpdateInput) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceUpdateInput) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *DeviceUpdateInput) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *DeviceUpdateInput) SetDescription(v string) {
	o.Description = &v
}

// GetHostname returns the Hostname field value if set, zero value otherwise.
func (o *DeviceUpdateInput) GetHostname() string {
	if o == nil || IsNil(o.Hostname) {
		var ret string
		return ret
	}
	return *o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceUpdateInput) GetHostnameOk() (*string, bool) {
	if o == nil || IsNil(o.Hostname) {
		return nil, false
	}
	return o.Hostname, true
}

// HasHostname returns a boolean if a field has been set.
func (o *DeviceUpdateInput) HasHostname() bool {
	if o != nil && !IsNil(o.Hostname) {
		return true
	}

	return false
}

// SetHostname gets a reference to the given string and assigns it to the Hostname field.
func (o *DeviceUpdateInput) SetHostname(v string) {
	o.Hostname = &v
}

// GetIpxeScriptUrl returns the IpxeScriptUrl field value if set, zero value otherwise.
func (o *DeviceUpdateInput) GetIpxeScriptUrl() string {
	if o == nil || IsNil(o.IpxeScriptUrl) {
		var ret string
		return ret
	}
	return *o.IpxeScriptUrl
}

// GetIpxeScriptUrlOk returns a tuple with the IpxeScriptUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceUpdateInput) GetIpxeScriptUrlOk() (*string, bool) {
	if o == nil || IsNil(o.IpxeScriptUrl) {
		return nil, false
	}
	return o.IpxeScriptUrl, true
}

// HasIpxeScriptUrl returns a boolean if a field has been set.
func (o *DeviceUpdateInput) HasIpxeScriptUrl() bool {
	if o != nil && !IsNil(o.IpxeScriptUrl) {
		return true
	}

	return false
}

// SetIpxeScriptUrl gets a reference to the given string and assigns it to the IpxeScriptUrl field.
func (o *DeviceUpdateInput) SetIpxeScriptUrl(v string) {
	o.IpxeScriptUrl = &v
}

// GetLocked returns the Locked field value if set, zero value otherwise.
func (o *DeviceUpdateInput) GetLocked() bool {
	if o == nil || IsNil(o.Locked) {
		var ret bool
		return ret
	}
	return *o.Locked
}

// GetLockedOk returns a tuple with the Locked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceUpdateInput) GetLockedOk() (*bool, bool) {
	if o == nil || IsNil(o.Locked) {
		return nil, false
	}
	return o.Locked, true
}

// HasLocked returns a boolean if a field has been set.
func (o *DeviceUpdateInput) HasLocked() bool {
	if o != nil && !IsNil(o.Locked) {
		return true
	}

	return false
}

// SetLocked gets a reference to the given bool and assigns it to the Locked field.
func (o *DeviceUpdateInput) SetLocked(v bool) {
	o.Locked = &v
}

// GetNetworkFrozen returns the NetworkFrozen field value if set, zero value otherwise.
func (o *DeviceUpdateInput) GetNetworkFrozen() bool {
	if o == nil || IsNil(o.NetworkFrozen) {
		var ret bool
		return ret
	}
	return *o.NetworkFrozen
}

// GetNetworkFrozenOk returns a tuple with the NetworkFrozen field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceUpdateInput) GetNetworkFrozenOk() (*bool, bool) {
	if o == nil || IsNil(o.NetworkFrozen) {
		return nil, false
	}
	return o.NetworkFrozen, true
}

// HasNetworkFrozen returns a boolean if a field has been set.
func (o *DeviceUpdateInput) HasNetworkFrozen() bool {
	if o != nil && !IsNil(o.NetworkFrozen) {
		return true
	}

	return false
}

// SetNetworkFrozen gets a reference to the given bool and assigns it to the NetworkFrozen field.
func (o *DeviceUpdateInput) SetNetworkFrozen(v bool) {
	o.NetworkFrozen = &v
}

// GetSpotInstance returns the SpotInstance field value if set, zero value otherwise.
func (o *DeviceUpdateInput) GetSpotInstance() bool {
	if o == nil || IsNil(o.SpotInstance) {
		var ret bool
		return ret
	}
	return *o.SpotInstance
}

// GetSpotInstanceOk returns a tuple with the SpotInstance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceUpdateInput) GetSpotInstanceOk() (*bool, bool) {
	if o == nil || IsNil(o.SpotInstance) {
		return nil, false
	}
	return o.SpotInstance, true
}

// HasSpotInstance returns a boolean if a field has been set.
func (o *DeviceUpdateInput) HasSpotInstance() bool {
	if o != nil && !IsNil(o.SpotInstance) {
		return true
	}

	return false
}

// SetSpotInstance gets a reference to the given bool and assigns it to the SpotInstance field.
func (o *DeviceUpdateInput) SetSpotInstance(v bool) {
	o.SpotInstance = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *DeviceUpdateInput) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceUpdateInput) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *DeviceUpdateInput) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *DeviceUpdateInput) SetTags(v []string) {
	o.Tags = v
}

// GetUserdata returns the Userdata field value if set, zero value otherwise.
func (o *DeviceUpdateInput) GetUserdata() string {
	if o == nil || IsNil(o.Userdata) {
		var ret string
		return ret
	}
	return *o.Userdata
}

// GetUserdataOk returns a tuple with the Userdata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceUpdateInput) GetUserdataOk() (*string, bool) {
	if o == nil || IsNil(o.Userdata) {
		return nil, false
	}
	return o.Userdata, true
}

// HasUserdata returns a boolean if a field has been set.
func (o *DeviceUpdateInput) HasUserdata() bool {
	if o != nil && !IsNil(o.Userdata) {
		return true
	}

	return false
}

// SetUserdata gets a reference to the given string and assigns it to the Userdata field.
func (o *DeviceUpdateInput) SetUserdata(v string) {
	o.Userdata = &v
}

func (o DeviceUpdateInput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeviceUpdateInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AlwaysPxe) {
		toSerialize["always_pxe"] = o.AlwaysPxe
	}
	if !IsNil(o.BillingCycle) {
		toSerialize["billing_cycle"] = o.BillingCycle
	}
	if !IsNil(o.Customdata) {
		toSerialize["customdata"] = o.Customdata
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Hostname) {
		toSerialize["hostname"] = o.Hostname
	}
	if !IsNil(o.IpxeScriptUrl) {
		toSerialize["ipxe_script_url"] = o.IpxeScriptUrl
	}
	if !IsNil(o.Locked) {
		toSerialize["locked"] = o.Locked
	}
	if !IsNil(o.NetworkFrozen) {
		toSerialize["network_frozen"] = o.NetworkFrozen
	}
	if !IsNil(o.SpotInstance) {
		toSerialize["spot_instance"] = o.SpotInstance
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.Userdata) {
		toSerialize["userdata"] = o.Userdata
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DeviceUpdateInput) UnmarshalJSON(bytes []byte) (err error) {
	varDeviceUpdateInput := _DeviceUpdateInput{}

	if err = json.Unmarshal(bytes, &varDeviceUpdateInput); err == nil {
		*o = DeviceUpdateInput(varDeviceUpdateInput)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "always_pxe")
		delete(additionalProperties, "billing_cycle")
		delete(additionalProperties, "customdata")
		delete(additionalProperties, "description")
		delete(additionalProperties, "hostname")
		delete(additionalProperties, "ipxe_script_url")
		delete(additionalProperties, "locked")
		delete(additionalProperties, "network_frozen")
		delete(additionalProperties, "spot_instance")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "userdata")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDeviceUpdateInput struct {
	value *DeviceUpdateInput
	isSet bool
}

func (v NullableDeviceUpdateInput) Get() *DeviceUpdateInput {
	return v.value
}

func (v *NullableDeviceUpdateInput) Set(val *DeviceUpdateInput) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceUpdateInput) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceUpdateInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceUpdateInput(val *DeviceUpdateInput) *NullableDeviceUpdateInput {
	return &NullableDeviceUpdateInput{value: val, isSet: true}
}

func (v NullableDeviceUpdateInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceUpdateInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
