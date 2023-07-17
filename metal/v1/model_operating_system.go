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

// checks if the OperatingSystem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OperatingSystem{}

// OperatingSystem struct for OperatingSystem
type OperatingSystem struct {
	Distro      *string `json:"distro,omitempty"`
	DistroLabel *string `json:"distro_label,omitempty"`
	Id          *string `json:"id,omitempty"`
	// Licenced OS is priced according to pricing property
	Licensed *bool   `json:"licensed,omitempty"`
	Name     *string `json:"name,omitempty"`
	// Servers can be already preinstalled with OS in order to shorten provision time.
	Preinstallable *bool `json:"preinstallable,omitempty"`
	// This object contains price per time unit and optional multiplier value if licence price depends on hardware plan or components (e.g. number of cores)
	Pricing         map[string]interface{} `json:"pricing,omitempty"`
	ProvisionableOn []string               `json:"provisionable_on,omitempty"`
	Slug            *string                `json:"slug,omitempty"`
	Version         *string                `json:"version,omitempty"`
	// Default operating system for the distro.
	DefaultOperatingSystem *bool `json:"default_operating_system,omitempty"`
	AdditionalProperties   map[string]interface{}
}

type _OperatingSystem OperatingSystem

// NewOperatingSystem instantiates a new OperatingSystem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOperatingSystem() *OperatingSystem {
	this := OperatingSystem{}
	return &this
}

// NewOperatingSystemWithDefaults instantiates a new OperatingSystem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOperatingSystemWithDefaults() *OperatingSystem {
	this := OperatingSystem{}
	return &this
}

// GetDistro returns the Distro field value if set, zero value otherwise.
func (o *OperatingSystem) GetDistro() string {
	if o == nil || IsNil(o.Distro) {
		var ret string
		return ret
	}
	return *o.Distro
}

// GetDistroOk returns a tuple with the Distro field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperatingSystem) GetDistroOk() (*string, bool) {
	if o == nil || IsNil(o.Distro) {
		return nil, false
	}
	return o.Distro, true
}

// HasDistro returns a boolean if a field has been set.
func (o *OperatingSystem) HasDistro() bool {
	if o != nil && !IsNil(o.Distro) {
		return true
	}

	return false
}

// SetDistro gets a reference to the given string and assigns it to the Distro field.
func (o *OperatingSystem) SetDistro(v string) {
	o.Distro = &v
}

// GetDistroLabel returns the DistroLabel field value if set, zero value otherwise.
func (o *OperatingSystem) GetDistroLabel() string {
	if o == nil || IsNil(o.DistroLabel) {
		var ret string
		return ret
	}
	return *o.DistroLabel
}

// GetDistroLabelOk returns a tuple with the DistroLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperatingSystem) GetDistroLabelOk() (*string, bool) {
	if o == nil || IsNil(o.DistroLabel) {
		return nil, false
	}
	return o.DistroLabel, true
}

// HasDistroLabel returns a boolean if a field has been set.
func (o *OperatingSystem) HasDistroLabel() bool {
	if o != nil && !IsNil(o.DistroLabel) {
		return true
	}

	return false
}

// SetDistroLabel gets a reference to the given string and assigns it to the DistroLabel field.
func (o *OperatingSystem) SetDistroLabel(v string) {
	o.DistroLabel = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OperatingSystem) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperatingSystem) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OperatingSystem) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *OperatingSystem) SetId(v string) {
	o.Id = &v
}

// GetLicensed returns the Licensed field value if set, zero value otherwise.
func (o *OperatingSystem) GetLicensed() bool {
	if o == nil || IsNil(o.Licensed) {
		var ret bool
		return ret
	}
	return *o.Licensed
}

// GetLicensedOk returns a tuple with the Licensed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperatingSystem) GetLicensedOk() (*bool, bool) {
	if o == nil || IsNil(o.Licensed) {
		return nil, false
	}
	return o.Licensed, true
}

// HasLicensed returns a boolean if a field has been set.
func (o *OperatingSystem) HasLicensed() bool {
	if o != nil && !IsNil(o.Licensed) {
		return true
	}

	return false
}

// SetLicensed gets a reference to the given bool and assigns it to the Licensed field.
func (o *OperatingSystem) SetLicensed(v bool) {
	o.Licensed = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *OperatingSystem) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperatingSystem) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *OperatingSystem) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *OperatingSystem) SetName(v string) {
	o.Name = &v
}

// GetPreinstallable returns the Preinstallable field value if set, zero value otherwise.
func (o *OperatingSystem) GetPreinstallable() bool {
	if o == nil || IsNil(o.Preinstallable) {
		var ret bool
		return ret
	}
	return *o.Preinstallable
}

// GetPreinstallableOk returns a tuple with the Preinstallable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperatingSystem) GetPreinstallableOk() (*bool, bool) {
	if o == nil || IsNil(o.Preinstallable) {
		return nil, false
	}
	return o.Preinstallable, true
}

// HasPreinstallable returns a boolean if a field has been set.
func (o *OperatingSystem) HasPreinstallable() bool {
	if o != nil && !IsNil(o.Preinstallable) {
		return true
	}

	return false
}

// SetPreinstallable gets a reference to the given bool and assigns it to the Preinstallable field.
func (o *OperatingSystem) SetPreinstallable(v bool) {
	o.Preinstallable = &v
}

// GetPricing returns the Pricing field value if set, zero value otherwise.
func (o *OperatingSystem) GetPricing() map[string]interface{} {
	if o == nil || IsNil(o.Pricing) {
		var ret map[string]interface{}
		return ret
	}
	return o.Pricing
}

// GetPricingOk returns a tuple with the Pricing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperatingSystem) GetPricingOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Pricing) {
		return map[string]interface{}{}, false
	}
	return o.Pricing, true
}

// HasPricing returns a boolean if a field has been set.
func (o *OperatingSystem) HasPricing() bool {
	if o != nil && !IsNil(o.Pricing) {
		return true
	}

	return false
}

// SetPricing gets a reference to the given map[string]interface{} and assigns it to the Pricing field.
func (o *OperatingSystem) SetPricing(v map[string]interface{}) {
	o.Pricing = v
}

// GetProvisionableOn returns the ProvisionableOn field value if set, zero value otherwise.
func (o *OperatingSystem) GetProvisionableOn() []string {
	if o == nil || IsNil(o.ProvisionableOn) {
		var ret []string
		return ret
	}
	return o.ProvisionableOn
}

// GetProvisionableOnOk returns a tuple with the ProvisionableOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperatingSystem) GetProvisionableOnOk() ([]string, bool) {
	if o == nil || IsNil(o.ProvisionableOn) {
		return nil, false
	}
	return o.ProvisionableOn, true
}

// HasProvisionableOn returns a boolean if a field has been set.
func (o *OperatingSystem) HasProvisionableOn() bool {
	if o != nil && !IsNil(o.ProvisionableOn) {
		return true
	}

	return false
}

// SetProvisionableOn gets a reference to the given []string and assigns it to the ProvisionableOn field.
func (o *OperatingSystem) SetProvisionableOn(v []string) {
	o.ProvisionableOn = v
}

// GetSlug returns the Slug field value if set, zero value otherwise.
func (o *OperatingSystem) GetSlug() string {
	if o == nil || IsNil(o.Slug) {
		var ret string
		return ret
	}
	return *o.Slug
}

// GetSlugOk returns a tuple with the Slug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperatingSystem) GetSlugOk() (*string, bool) {
	if o == nil || IsNil(o.Slug) {
		return nil, false
	}
	return o.Slug, true
}

// HasSlug returns a boolean if a field has been set.
func (o *OperatingSystem) HasSlug() bool {
	if o != nil && !IsNil(o.Slug) {
		return true
	}

	return false
}

// SetSlug gets a reference to the given string and assigns it to the Slug field.
func (o *OperatingSystem) SetSlug(v string) {
	o.Slug = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *OperatingSystem) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperatingSystem) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *OperatingSystem) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *OperatingSystem) SetVersion(v string) {
	o.Version = &v
}

// GetDefaultOperatingSystem returns the DefaultOperatingSystem field value if set, zero value otherwise.
func (o *OperatingSystem) GetDefaultOperatingSystem() bool {
	if o == nil || IsNil(o.DefaultOperatingSystem) {
		var ret bool
		return ret
	}
	return *o.DefaultOperatingSystem
}

// GetDefaultOperatingSystemOk returns a tuple with the DefaultOperatingSystem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperatingSystem) GetDefaultOperatingSystemOk() (*bool, bool) {
	if o == nil || IsNil(o.DefaultOperatingSystem) {
		return nil, false
	}
	return o.DefaultOperatingSystem, true
}

// HasDefaultOperatingSystem returns a boolean if a field has been set.
func (o *OperatingSystem) HasDefaultOperatingSystem() bool {
	if o != nil && !IsNil(o.DefaultOperatingSystem) {
		return true
	}

	return false
}

// SetDefaultOperatingSystem gets a reference to the given bool and assigns it to the DefaultOperatingSystem field.
func (o *OperatingSystem) SetDefaultOperatingSystem(v bool) {
	o.DefaultOperatingSystem = &v
}

func (o OperatingSystem) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OperatingSystem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Distro) {
		toSerialize["distro"] = o.Distro
	}
	if !IsNil(o.DistroLabel) {
		toSerialize["distro_label"] = o.DistroLabel
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Licensed) {
		toSerialize["licensed"] = o.Licensed
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Preinstallable) {
		toSerialize["preinstallable"] = o.Preinstallable
	}
	if !IsNil(o.Pricing) {
		toSerialize["pricing"] = o.Pricing
	}
	if !IsNil(o.ProvisionableOn) {
		toSerialize["provisionable_on"] = o.ProvisionableOn
	}
	if !IsNil(o.Slug) {
		toSerialize["slug"] = o.Slug
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	// skip: default_operating_system is readOnly

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OperatingSystem) UnmarshalJSON(bytes []byte) (err error) {
	varOperatingSystem := _OperatingSystem{}

	if err = json.Unmarshal(bytes, &varOperatingSystem); err == nil {
		*o = OperatingSystem(varOperatingSystem)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "distro")
		delete(additionalProperties, "distro_label")
		delete(additionalProperties, "id")
		delete(additionalProperties, "licensed")
		delete(additionalProperties, "name")
		delete(additionalProperties, "preinstallable")
		delete(additionalProperties, "pricing")
		delete(additionalProperties, "provisionable_on")
		delete(additionalProperties, "slug")
		delete(additionalProperties, "version")
		delete(additionalProperties, "default_operating_system")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOperatingSystem struct {
	value *OperatingSystem
	isSet bool
}

func (v NullableOperatingSystem) Get() *OperatingSystem {
	return v.value
}

func (v *NullableOperatingSystem) Set(val *OperatingSystem) {
	v.value = val
	v.isSet = true
}

func (v NullableOperatingSystem) IsSet() bool {
	return v.isSet
}

func (v *NullableOperatingSystem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOperatingSystem(val *OperatingSystem) *NullableOperatingSystem {
	return &NullableOperatingSystem{value: val, isSet: true}
}

func (v NullableOperatingSystem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOperatingSystem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
