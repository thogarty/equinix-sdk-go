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

// checks if the ApiConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiConfig{}

// ApiConfig Configuration for API based Integration for Service Profile
type ApiConfig struct {
	// Setting indicating whether the API is available (true) or not (false).
	ApiAvailable  *bool   `json:"apiAvailable,omitempty"`
	IntegrationId *string `json:"integrationId,omitempty"`
	// Setting indicating that the port is managed by Equinix (true) or not (false).
	EquinixManagedPort *bool `json:"equinixManagedPort,omitempty"`
	// Setting indicating that the VLAN is managed by Equinix (true) or not (false).
	EquinixManagedVlan *bool `json:"equinixManagedVlan,omitempty"`
	// Setting showing that oversubscription support is available (true) or not (false). The default is false. Oversubscription is the sale of more than the available network bandwidth. This practice is common and legitimate. After all, many customers use less bandwidth than they've purchased. And network users don't consume bandwidth all at the same time. The leftover bandwidth can be sold to other customers. When demand surges, operational and engineering resources can be shifted to accommodate the load.
	AllowOverSubscription *bool `json:"allowOverSubscription,omitempty"`
	// A cap on oversubscription.
	OverSubscriptionLimit *int32 `json:"overSubscriptionLimit,omitempty"`
	BandwidthFromApi      *bool  `json:"bandwidthFromApi,omitempty"`
	AdditionalProperties  map[string]interface{}
}

type _ApiConfig ApiConfig

// NewApiConfig instantiates a new ApiConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiConfig() *ApiConfig {
	this := ApiConfig{}
	var apiAvailable bool = false
	this.ApiAvailable = &apiAvailable
	var equinixManagedPort bool = false
	this.EquinixManagedPort = &equinixManagedPort
	var equinixManagedVlan bool = false
	this.EquinixManagedVlan = &equinixManagedVlan
	var allowOverSubscription bool = false
	this.AllowOverSubscription = &allowOverSubscription
	var overSubscriptionLimit int32 = 1
	this.OverSubscriptionLimit = &overSubscriptionLimit
	var bandwidthFromApi bool = false
	this.BandwidthFromApi = &bandwidthFromApi
	return &this
}

// NewApiConfigWithDefaults instantiates a new ApiConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiConfigWithDefaults() *ApiConfig {
	this := ApiConfig{}
	var apiAvailable bool = false
	this.ApiAvailable = &apiAvailable
	var equinixManagedPort bool = false
	this.EquinixManagedPort = &equinixManagedPort
	var equinixManagedVlan bool = false
	this.EquinixManagedVlan = &equinixManagedVlan
	var allowOverSubscription bool = false
	this.AllowOverSubscription = &allowOverSubscription
	var overSubscriptionLimit int32 = 1
	this.OverSubscriptionLimit = &overSubscriptionLimit
	var bandwidthFromApi bool = false
	this.BandwidthFromApi = &bandwidthFromApi
	return &this
}

// GetApiAvailable returns the ApiAvailable field value if set, zero value otherwise.
func (o *ApiConfig) GetApiAvailable() bool {
	if o == nil || IsNil(o.ApiAvailable) {
		var ret bool
		return ret
	}
	return *o.ApiAvailable
}

// GetApiAvailableOk returns a tuple with the ApiAvailable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiConfig) GetApiAvailableOk() (*bool, bool) {
	if o == nil || IsNil(o.ApiAvailable) {
		return nil, false
	}
	return o.ApiAvailable, true
}

// HasApiAvailable returns a boolean if a field has been set.
func (o *ApiConfig) HasApiAvailable() bool {
	if o != nil && !IsNil(o.ApiAvailable) {
		return true
	}

	return false
}

// SetApiAvailable gets a reference to the given bool and assigns it to the ApiAvailable field.
func (o *ApiConfig) SetApiAvailable(v bool) {
	o.ApiAvailable = &v
}

// GetIntegrationId returns the IntegrationId field value if set, zero value otherwise.
func (o *ApiConfig) GetIntegrationId() string {
	if o == nil || IsNil(o.IntegrationId) {
		var ret string
		return ret
	}
	return *o.IntegrationId
}

// GetIntegrationIdOk returns a tuple with the IntegrationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiConfig) GetIntegrationIdOk() (*string, bool) {
	if o == nil || IsNil(o.IntegrationId) {
		return nil, false
	}
	return o.IntegrationId, true
}

// HasIntegrationId returns a boolean if a field has been set.
func (o *ApiConfig) HasIntegrationId() bool {
	if o != nil && !IsNil(o.IntegrationId) {
		return true
	}

	return false
}

// SetIntegrationId gets a reference to the given string and assigns it to the IntegrationId field.
func (o *ApiConfig) SetIntegrationId(v string) {
	o.IntegrationId = &v
}

// GetEquinixManagedPort returns the EquinixManagedPort field value if set, zero value otherwise.
func (o *ApiConfig) GetEquinixManagedPort() bool {
	if o == nil || IsNil(o.EquinixManagedPort) {
		var ret bool
		return ret
	}
	return *o.EquinixManagedPort
}

// GetEquinixManagedPortOk returns a tuple with the EquinixManagedPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiConfig) GetEquinixManagedPortOk() (*bool, bool) {
	if o == nil || IsNil(o.EquinixManagedPort) {
		return nil, false
	}
	return o.EquinixManagedPort, true
}

// HasEquinixManagedPort returns a boolean if a field has been set.
func (o *ApiConfig) HasEquinixManagedPort() bool {
	if o != nil && !IsNil(o.EquinixManagedPort) {
		return true
	}

	return false
}

// SetEquinixManagedPort gets a reference to the given bool and assigns it to the EquinixManagedPort field.
func (o *ApiConfig) SetEquinixManagedPort(v bool) {
	o.EquinixManagedPort = &v
}

// GetEquinixManagedVlan returns the EquinixManagedVlan field value if set, zero value otherwise.
func (o *ApiConfig) GetEquinixManagedVlan() bool {
	if o == nil || IsNil(o.EquinixManagedVlan) {
		var ret bool
		return ret
	}
	return *o.EquinixManagedVlan
}

// GetEquinixManagedVlanOk returns a tuple with the EquinixManagedVlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiConfig) GetEquinixManagedVlanOk() (*bool, bool) {
	if o == nil || IsNil(o.EquinixManagedVlan) {
		return nil, false
	}
	return o.EquinixManagedVlan, true
}

// HasEquinixManagedVlan returns a boolean if a field has been set.
func (o *ApiConfig) HasEquinixManagedVlan() bool {
	if o != nil && !IsNil(o.EquinixManagedVlan) {
		return true
	}

	return false
}

// SetEquinixManagedVlan gets a reference to the given bool and assigns it to the EquinixManagedVlan field.
func (o *ApiConfig) SetEquinixManagedVlan(v bool) {
	o.EquinixManagedVlan = &v
}

// GetAllowOverSubscription returns the AllowOverSubscription field value if set, zero value otherwise.
func (o *ApiConfig) GetAllowOverSubscription() bool {
	if o == nil || IsNil(o.AllowOverSubscription) {
		var ret bool
		return ret
	}
	return *o.AllowOverSubscription
}

// GetAllowOverSubscriptionOk returns a tuple with the AllowOverSubscription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiConfig) GetAllowOverSubscriptionOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowOverSubscription) {
		return nil, false
	}
	return o.AllowOverSubscription, true
}

// HasAllowOverSubscription returns a boolean if a field has been set.
func (o *ApiConfig) HasAllowOverSubscription() bool {
	if o != nil && !IsNil(o.AllowOverSubscription) {
		return true
	}

	return false
}

// SetAllowOverSubscription gets a reference to the given bool and assigns it to the AllowOverSubscription field.
func (o *ApiConfig) SetAllowOverSubscription(v bool) {
	o.AllowOverSubscription = &v
}

// GetOverSubscriptionLimit returns the OverSubscriptionLimit field value if set, zero value otherwise.
func (o *ApiConfig) GetOverSubscriptionLimit() int32 {
	if o == nil || IsNil(o.OverSubscriptionLimit) {
		var ret int32
		return ret
	}
	return *o.OverSubscriptionLimit
}

// GetOverSubscriptionLimitOk returns a tuple with the OverSubscriptionLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiConfig) GetOverSubscriptionLimitOk() (*int32, bool) {
	if o == nil || IsNil(o.OverSubscriptionLimit) {
		return nil, false
	}
	return o.OverSubscriptionLimit, true
}

// HasOverSubscriptionLimit returns a boolean if a field has been set.
func (o *ApiConfig) HasOverSubscriptionLimit() bool {
	if o != nil && !IsNil(o.OverSubscriptionLimit) {
		return true
	}

	return false
}

// SetOverSubscriptionLimit gets a reference to the given int32 and assigns it to the OverSubscriptionLimit field.
func (o *ApiConfig) SetOverSubscriptionLimit(v int32) {
	o.OverSubscriptionLimit = &v
}

// GetBandwidthFromApi returns the BandwidthFromApi field value if set, zero value otherwise.
func (o *ApiConfig) GetBandwidthFromApi() bool {
	if o == nil || IsNil(o.BandwidthFromApi) {
		var ret bool
		return ret
	}
	return *o.BandwidthFromApi
}

// GetBandwidthFromApiOk returns a tuple with the BandwidthFromApi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiConfig) GetBandwidthFromApiOk() (*bool, bool) {
	if o == nil || IsNil(o.BandwidthFromApi) {
		return nil, false
	}
	return o.BandwidthFromApi, true
}

// HasBandwidthFromApi returns a boolean if a field has been set.
func (o *ApiConfig) HasBandwidthFromApi() bool {
	if o != nil && !IsNil(o.BandwidthFromApi) {
		return true
	}

	return false
}

// SetBandwidthFromApi gets a reference to the given bool and assigns it to the BandwidthFromApi field.
func (o *ApiConfig) SetBandwidthFromApi(v bool) {
	o.BandwidthFromApi = &v
}

func (o ApiConfig) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ApiAvailable) {
		toSerialize["apiAvailable"] = o.ApiAvailable
	}
	if !IsNil(o.IntegrationId) {
		toSerialize["integrationId"] = o.IntegrationId
	}
	if !IsNil(o.EquinixManagedPort) {
		toSerialize["equinixManagedPort"] = o.EquinixManagedPort
	}
	if !IsNil(o.EquinixManagedVlan) {
		toSerialize["equinixManagedVlan"] = o.EquinixManagedVlan
	}
	if !IsNil(o.AllowOverSubscription) {
		toSerialize["allowOverSubscription"] = o.AllowOverSubscription
	}
	if !IsNil(o.OverSubscriptionLimit) {
		toSerialize["overSubscriptionLimit"] = o.OverSubscriptionLimit
	}
	if !IsNil(o.BandwidthFromApi) {
		toSerialize["bandwidthFromApi"] = o.BandwidthFromApi
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ApiConfig) UnmarshalJSON(data []byte) (err error) {
	varApiConfig := _ApiConfig{}

	err = json.Unmarshal(data, &varApiConfig)

	if err != nil {
		return err
	}

	*o = ApiConfig(varApiConfig)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "apiAvailable")
		delete(additionalProperties, "integrationId")
		delete(additionalProperties, "equinixManagedPort")
		delete(additionalProperties, "equinixManagedVlan")
		delete(additionalProperties, "allowOverSubscription")
		delete(additionalProperties, "overSubscriptionLimit")
		delete(additionalProperties, "bandwidthFromApi")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableApiConfig struct {
	value *ApiConfig
	isSet bool
}

func (v NullableApiConfig) Get() *ApiConfig {
	return v.value
}

func (v *NullableApiConfig) Set(val *ApiConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableApiConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableApiConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiConfig(val *ApiConfig) *NullableApiConfig {
	return &NullableApiConfig{value: val, isSet: true}
}

func (v NullableApiConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}