/*
Equinix Fabric API v4

Contact: api-support@equinix.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fabricv4

import (
	"encoding/json"
	"fmt"
)

// checks if the ServiceProfileAccessPointTypeCOLO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceProfileAccessPointTypeCOLO{}

// ServiceProfileAccessPointTypeCOLO Colo Access Point Type
type ServiceProfileAccessPointTypeCOLO struct {
	Uuid                *string                           `json:"uuid,omitempty"`
	Type                ServiceProfileAccessPointTypeEnum `json:"type"`
	SupportedBandwidths []int32                           `json:"supportedBandwidths,omitempty"`
	// Setting to allow or prohibit remote connections to the service profile.
	AllowRemoteConnections *bool `json:"allowRemoteConnections,omitempty"`
	// Setting to enable or disable the ability of the buyer to customize the bandwidth.
	AllowCustomBandwidth *bool `json:"allowCustomBandwidth,omitempty"`
	// percentage of port bandwidth at which an allocation alert is generated - missing on wiki.
	BandwidthAlertThreshold *float32 `json:"bandwidthAlertThreshold,omitempty"`
	// Setting to enable or disable the ability of the buyer to change connection bandwidth without approval of the seller.
	AllowBandwidthAutoApproval *bool `json:"allowBandwidthAutoApproval,omitempty"`
	// Availability of a bandwidth upgrade. The default is false.
	AllowBandwidthUpgrade *bool                             `json:"allowBandwidthUpgrade,omitempty"`
	LinkProtocolConfig    *ServiceProfileLinkProtocolConfig `json:"linkProtocolConfig,omitempty"`
	// for verizon only.
	EnableAutoGenerateServiceKey *bool `json:"enableAutoGenerateServiceKey,omitempty"`
	// Mandate redundant connections
	ConnectionRedundancyRequired *bool      `json:"connectionRedundancyRequired,omitempty"`
	ApiConfig                    *ApiConfig `json:"apiConfig,omitempty"`
	// custom name for \"Connection\"
	ConnectionLabel      *string                 `json:"connectionLabel,omitempty"`
	AuthenticationKey    *AuthenticationKey      `json:"authenticationKey,omitempty"`
	Metadata             *ServiceProfileMetadata `json:"metadata,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ServiceProfileAccessPointTypeCOLO ServiceProfileAccessPointTypeCOLO

// NewServiceProfileAccessPointTypeCOLO instantiates a new ServiceProfileAccessPointTypeCOLO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceProfileAccessPointTypeCOLO(type_ ServiceProfileAccessPointTypeEnum) *ServiceProfileAccessPointTypeCOLO {
	this := ServiceProfileAccessPointTypeCOLO{}
	this.Type = type_
	var allowRemoteConnections bool = false
	this.AllowRemoteConnections = &allowRemoteConnections
	var allowCustomBandwidth bool = false
	this.AllowCustomBandwidth = &allowCustomBandwidth
	var allowBandwidthAutoApproval bool = false
	this.AllowBandwidthAutoApproval = &allowBandwidthAutoApproval
	var connectionRedundancyRequired bool = false
	this.ConnectionRedundancyRequired = &connectionRedundancyRequired
	return &this
}

// NewServiceProfileAccessPointTypeCOLOWithDefaults instantiates a new ServiceProfileAccessPointTypeCOLO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceProfileAccessPointTypeCOLOWithDefaults() *ServiceProfileAccessPointTypeCOLO {
	this := ServiceProfileAccessPointTypeCOLO{}
	var allowRemoteConnections bool = false
	this.AllowRemoteConnections = &allowRemoteConnections
	var allowCustomBandwidth bool = false
	this.AllowCustomBandwidth = &allowCustomBandwidth
	var allowBandwidthAutoApproval bool = false
	this.AllowBandwidthAutoApproval = &allowBandwidthAutoApproval
	var connectionRedundancyRequired bool = false
	this.ConnectionRedundancyRequired = &connectionRedundancyRequired
	return &this
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *ServiceProfileAccessPointTypeCOLO) GetUuid() string {
	if o == nil || IsNil(o.Uuid) {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceProfileAccessPointTypeCOLO) GetUuidOk() (*string, bool) {
	if o == nil || IsNil(o.Uuid) {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *ServiceProfileAccessPointTypeCOLO) HasUuid() bool {
	if o != nil && !IsNil(o.Uuid) {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *ServiceProfileAccessPointTypeCOLO) SetUuid(v string) {
	o.Uuid = &v
}

// GetType returns the Type field value
func (o *ServiceProfileAccessPointTypeCOLO) GetType() ServiceProfileAccessPointTypeEnum {
	if o == nil {
		var ret ServiceProfileAccessPointTypeEnum
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ServiceProfileAccessPointTypeCOLO) GetTypeOk() (*ServiceProfileAccessPointTypeEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ServiceProfileAccessPointTypeCOLO) SetType(v ServiceProfileAccessPointTypeEnum) {
	o.Type = v
}

// GetSupportedBandwidths returns the SupportedBandwidths field value if set, zero value otherwise.
func (o *ServiceProfileAccessPointTypeCOLO) GetSupportedBandwidths() []int32 {
	if o == nil || IsNil(o.SupportedBandwidths) {
		var ret []int32
		return ret
	}
	return o.SupportedBandwidths
}

// GetSupportedBandwidthsOk returns a tuple with the SupportedBandwidths field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceProfileAccessPointTypeCOLO) GetSupportedBandwidthsOk() ([]int32, bool) {
	if o == nil || IsNil(o.SupportedBandwidths) {
		return nil, false
	}
	return o.SupportedBandwidths, true
}

// HasSupportedBandwidths returns a boolean if a field has been set.
func (o *ServiceProfileAccessPointTypeCOLO) HasSupportedBandwidths() bool {
	if o != nil && !IsNil(o.SupportedBandwidths) {
		return true
	}

	return false
}

// SetSupportedBandwidths gets a reference to the given []int32 and assigns it to the SupportedBandwidths field.
func (o *ServiceProfileAccessPointTypeCOLO) SetSupportedBandwidths(v []int32) {
	o.SupportedBandwidths = v
}

// GetAllowRemoteConnections returns the AllowRemoteConnections field value if set, zero value otherwise.
func (o *ServiceProfileAccessPointTypeCOLO) GetAllowRemoteConnections() bool {
	if o == nil || IsNil(o.AllowRemoteConnections) {
		var ret bool
		return ret
	}
	return *o.AllowRemoteConnections
}

// GetAllowRemoteConnectionsOk returns a tuple with the AllowRemoteConnections field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceProfileAccessPointTypeCOLO) GetAllowRemoteConnectionsOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowRemoteConnections) {
		return nil, false
	}
	return o.AllowRemoteConnections, true
}

// HasAllowRemoteConnections returns a boolean if a field has been set.
func (o *ServiceProfileAccessPointTypeCOLO) HasAllowRemoteConnections() bool {
	if o != nil && !IsNil(o.AllowRemoteConnections) {
		return true
	}

	return false
}

// SetAllowRemoteConnections gets a reference to the given bool and assigns it to the AllowRemoteConnections field.
func (o *ServiceProfileAccessPointTypeCOLO) SetAllowRemoteConnections(v bool) {
	o.AllowRemoteConnections = &v
}

// GetAllowCustomBandwidth returns the AllowCustomBandwidth field value if set, zero value otherwise.
func (o *ServiceProfileAccessPointTypeCOLO) GetAllowCustomBandwidth() bool {
	if o == nil || IsNil(o.AllowCustomBandwidth) {
		var ret bool
		return ret
	}
	return *o.AllowCustomBandwidth
}

// GetAllowCustomBandwidthOk returns a tuple with the AllowCustomBandwidth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceProfileAccessPointTypeCOLO) GetAllowCustomBandwidthOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowCustomBandwidth) {
		return nil, false
	}
	return o.AllowCustomBandwidth, true
}

// HasAllowCustomBandwidth returns a boolean if a field has been set.
func (o *ServiceProfileAccessPointTypeCOLO) HasAllowCustomBandwidth() bool {
	if o != nil && !IsNil(o.AllowCustomBandwidth) {
		return true
	}

	return false
}

// SetAllowCustomBandwidth gets a reference to the given bool and assigns it to the AllowCustomBandwidth field.
func (o *ServiceProfileAccessPointTypeCOLO) SetAllowCustomBandwidth(v bool) {
	o.AllowCustomBandwidth = &v
}

// GetBandwidthAlertThreshold returns the BandwidthAlertThreshold field value if set, zero value otherwise.
func (o *ServiceProfileAccessPointTypeCOLO) GetBandwidthAlertThreshold() float32 {
	if o == nil || IsNil(o.BandwidthAlertThreshold) {
		var ret float32
		return ret
	}
	return *o.BandwidthAlertThreshold
}

// GetBandwidthAlertThresholdOk returns a tuple with the BandwidthAlertThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceProfileAccessPointTypeCOLO) GetBandwidthAlertThresholdOk() (*float32, bool) {
	if o == nil || IsNil(o.BandwidthAlertThreshold) {
		return nil, false
	}
	return o.BandwidthAlertThreshold, true
}

// HasBandwidthAlertThreshold returns a boolean if a field has been set.
func (o *ServiceProfileAccessPointTypeCOLO) HasBandwidthAlertThreshold() bool {
	if o != nil && !IsNil(o.BandwidthAlertThreshold) {
		return true
	}

	return false
}

// SetBandwidthAlertThreshold gets a reference to the given float32 and assigns it to the BandwidthAlertThreshold field.
func (o *ServiceProfileAccessPointTypeCOLO) SetBandwidthAlertThreshold(v float32) {
	o.BandwidthAlertThreshold = &v
}

// GetAllowBandwidthAutoApproval returns the AllowBandwidthAutoApproval field value if set, zero value otherwise.
func (o *ServiceProfileAccessPointTypeCOLO) GetAllowBandwidthAutoApproval() bool {
	if o == nil || IsNil(o.AllowBandwidthAutoApproval) {
		var ret bool
		return ret
	}
	return *o.AllowBandwidthAutoApproval
}

// GetAllowBandwidthAutoApprovalOk returns a tuple with the AllowBandwidthAutoApproval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceProfileAccessPointTypeCOLO) GetAllowBandwidthAutoApprovalOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowBandwidthAutoApproval) {
		return nil, false
	}
	return o.AllowBandwidthAutoApproval, true
}

// HasAllowBandwidthAutoApproval returns a boolean if a field has been set.
func (o *ServiceProfileAccessPointTypeCOLO) HasAllowBandwidthAutoApproval() bool {
	if o != nil && !IsNil(o.AllowBandwidthAutoApproval) {
		return true
	}

	return false
}

// SetAllowBandwidthAutoApproval gets a reference to the given bool and assigns it to the AllowBandwidthAutoApproval field.
func (o *ServiceProfileAccessPointTypeCOLO) SetAllowBandwidthAutoApproval(v bool) {
	o.AllowBandwidthAutoApproval = &v
}

// GetAllowBandwidthUpgrade returns the AllowBandwidthUpgrade field value if set, zero value otherwise.
func (o *ServiceProfileAccessPointTypeCOLO) GetAllowBandwidthUpgrade() bool {
	if o == nil || IsNil(o.AllowBandwidthUpgrade) {
		var ret bool
		return ret
	}
	return *o.AllowBandwidthUpgrade
}

// GetAllowBandwidthUpgradeOk returns a tuple with the AllowBandwidthUpgrade field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceProfileAccessPointTypeCOLO) GetAllowBandwidthUpgradeOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowBandwidthUpgrade) {
		return nil, false
	}
	return o.AllowBandwidthUpgrade, true
}

// HasAllowBandwidthUpgrade returns a boolean if a field has been set.
func (o *ServiceProfileAccessPointTypeCOLO) HasAllowBandwidthUpgrade() bool {
	if o != nil && !IsNil(o.AllowBandwidthUpgrade) {
		return true
	}

	return false
}

// SetAllowBandwidthUpgrade gets a reference to the given bool and assigns it to the AllowBandwidthUpgrade field.
func (o *ServiceProfileAccessPointTypeCOLO) SetAllowBandwidthUpgrade(v bool) {
	o.AllowBandwidthUpgrade = &v
}

// GetLinkProtocolConfig returns the LinkProtocolConfig field value if set, zero value otherwise.
func (o *ServiceProfileAccessPointTypeCOLO) GetLinkProtocolConfig() ServiceProfileLinkProtocolConfig {
	if o == nil || IsNil(o.LinkProtocolConfig) {
		var ret ServiceProfileLinkProtocolConfig
		return ret
	}
	return *o.LinkProtocolConfig
}

// GetLinkProtocolConfigOk returns a tuple with the LinkProtocolConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceProfileAccessPointTypeCOLO) GetLinkProtocolConfigOk() (*ServiceProfileLinkProtocolConfig, bool) {
	if o == nil || IsNil(o.LinkProtocolConfig) {
		return nil, false
	}
	return o.LinkProtocolConfig, true
}

// HasLinkProtocolConfig returns a boolean if a field has been set.
func (o *ServiceProfileAccessPointTypeCOLO) HasLinkProtocolConfig() bool {
	if o != nil && !IsNil(o.LinkProtocolConfig) {
		return true
	}

	return false
}

// SetLinkProtocolConfig gets a reference to the given ServiceProfileLinkProtocolConfig and assigns it to the LinkProtocolConfig field.
func (o *ServiceProfileAccessPointTypeCOLO) SetLinkProtocolConfig(v ServiceProfileLinkProtocolConfig) {
	o.LinkProtocolConfig = &v
}

// GetEnableAutoGenerateServiceKey returns the EnableAutoGenerateServiceKey field value if set, zero value otherwise.
func (o *ServiceProfileAccessPointTypeCOLO) GetEnableAutoGenerateServiceKey() bool {
	if o == nil || IsNil(o.EnableAutoGenerateServiceKey) {
		var ret bool
		return ret
	}
	return *o.EnableAutoGenerateServiceKey
}

// GetEnableAutoGenerateServiceKeyOk returns a tuple with the EnableAutoGenerateServiceKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceProfileAccessPointTypeCOLO) GetEnableAutoGenerateServiceKeyOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableAutoGenerateServiceKey) {
		return nil, false
	}
	return o.EnableAutoGenerateServiceKey, true
}

// HasEnableAutoGenerateServiceKey returns a boolean if a field has been set.
func (o *ServiceProfileAccessPointTypeCOLO) HasEnableAutoGenerateServiceKey() bool {
	if o != nil && !IsNil(o.EnableAutoGenerateServiceKey) {
		return true
	}

	return false
}

// SetEnableAutoGenerateServiceKey gets a reference to the given bool and assigns it to the EnableAutoGenerateServiceKey field.
func (o *ServiceProfileAccessPointTypeCOLO) SetEnableAutoGenerateServiceKey(v bool) {
	o.EnableAutoGenerateServiceKey = &v
}

// GetConnectionRedundancyRequired returns the ConnectionRedundancyRequired field value if set, zero value otherwise.
func (o *ServiceProfileAccessPointTypeCOLO) GetConnectionRedundancyRequired() bool {
	if o == nil || IsNil(o.ConnectionRedundancyRequired) {
		var ret bool
		return ret
	}
	return *o.ConnectionRedundancyRequired
}

// GetConnectionRedundancyRequiredOk returns a tuple with the ConnectionRedundancyRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceProfileAccessPointTypeCOLO) GetConnectionRedundancyRequiredOk() (*bool, bool) {
	if o == nil || IsNil(o.ConnectionRedundancyRequired) {
		return nil, false
	}
	return o.ConnectionRedundancyRequired, true
}

// HasConnectionRedundancyRequired returns a boolean if a field has been set.
func (o *ServiceProfileAccessPointTypeCOLO) HasConnectionRedundancyRequired() bool {
	if o != nil && !IsNil(o.ConnectionRedundancyRequired) {
		return true
	}

	return false
}

// SetConnectionRedundancyRequired gets a reference to the given bool and assigns it to the ConnectionRedundancyRequired field.
func (o *ServiceProfileAccessPointTypeCOLO) SetConnectionRedundancyRequired(v bool) {
	o.ConnectionRedundancyRequired = &v
}

// GetApiConfig returns the ApiConfig field value if set, zero value otherwise.
func (o *ServiceProfileAccessPointTypeCOLO) GetApiConfig() ApiConfig {
	if o == nil || IsNil(o.ApiConfig) {
		var ret ApiConfig
		return ret
	}
	return *o.ApiConfig
}

// GetApiConfigOk returns a tuple with the ApiConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceProfileAccessPointTypeCOLO) GetApiConfigOk() (*ApiConfig, bool) {
	if o == nil || IsNil(o.ApiConfig) {
		return nil, false
	}
	return o.ApiConfig, true
}

// HasApiConfig returns a boolean if a field has been set.
func (o *ServiceProfileAccessPointTypeCOLO) HasApiConfig() bool {
	if o != nil && !IsNil(o.ApiConfig) {
		return true
	}

	return false
}

// SetApiConfig gets a reference to the given ApiConfig and assigns it to the ApiConfig field.
func (o *ServiceProfileAccessPointTypeCOLO) SetApiConfig(v ApiConfig) {
	o.ApiConfig = &v
}

// GetConnectionLabel returns the ConnectionLabel field value if set, zero value otherwise.
func (o *ServiceProfileAccessPointTypeCOLO) GetConnectionLabel() string {
	if o == nil || IsNil(o.ConnectionLabel) {
		var ret string
		return ret
	}
	return *o.ConnectionLabel
}

// GetConnectionLabelOk returns a tuple with the ConnectionLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceProfileAccessPointTypeCOLO) GetConnectionLabelOk() (*string, bool) {
	if o == nil || IsNil(o.ConnectionLabel) {
		return nil, false
	}
	return o.ConnectionLabel, true
}

// HasConnectionLabel returns a boolean if a field has been set.
func (o *ServiceProfileAccessPointTypeCOLO) HasConnectionLabel() bool {
	if o != nil && !IsNil(o.ConnectionLabel) {
		return true
	}

	return false
}

// SetConnectionLabel gets a reference to the given string and assigns it to the ConnectionLabel field.
func (o *ServiceProfileAccessPointTypeCOLO) SetConnectionLabel(v string) {
	o.ConnectionLabel = &v
}

// GetAuthenticationKey returns the AuthenticationKey field value if set, zero value otherwise.
func (o *ServiceProfileAccessPointTypeCOLO) GetAuthenticationKey() AuthenticationKey {
	if o == nil || IsNil(o.AuthenticationKey) {
		var ret AuthenticationKey
		return ret
	}
	return *o.AuthenticationKey
}

// GetAuthenticationKeyOk returns a tuple with the AuthenticationKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceProfileAccessPointTypeCOLO) GetAuthenticationKeyOk() (*AuthenticationKey, bool) {
	if o == nil || IsNil(o.AuthenticationKey) {
		return nil, false
	}
	return o.AuthenticationKey, true
}

// HasAuthenticationKey returns a boolean if a field has been set.
func (o *ServiceProfileAccessPointTypeCOLO) HasAuthenticationKey() bool {
	if o != nil && !IsNil(o.AuthenticationKey) {
		return true
	}

	return false
}

// SetAuthenticationKey gets a reference to the given AuthenticationKey and assigns it to the AuthenticationKey field.
func (o *ServiceProfileAccessPointTypeCOLO) SetAuthenticationKey(v AuthenticationKey) {
	o.AuthenticationKey = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *ServiceProfileAccessPointTypeCOLO) GetMetadata() ServiceProfileMetadata {
	if o == nil || IsNil(o.Metadata) {
		var ret ServiceProfileMetadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceProfileAccessPointTypeCOLO) GetMetadataOk() (*ServiceProfileMetadata, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *ServiceProfileAccessPointTypeCOLO) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given ServiceProfileMetadata and assigns it to the Metadata field.
func (o *ServiceProfileAccessPointTypeCOLO) SetMetadata(v ServiceProfileMetadata) {
	o.Metadata = &v
}

func (o ServiceProfileAccessPointTypeCOLO) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceProfileAccessPointTypeCOLO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Uuid) {
		toSerialize["uuid"] = o.Uuid
	}
	toSerialize["type"] = o.Type
	if !IsNil(o.SupportedBandwidths) {
		toSerialize["supportedBandwidths"] = o.SupportedBandwidths
	}
	if !IsNil(o.AllowRemoteConnections) {
		toSerialize["allowRemoteConnections"] = o.AllowRemoteConnections
	}
	if !IsNil(o.AllowCustomBandwidth) {
		toSerialize["allowCustomBandwidth"] = o.AllowCustomBandwidth
	}
	if !IsNil(o.BandwidthAlertThreshold) {
		toSerialize["bandwidthAlertThreshold"] = o.BandwidthAlertThreshold
	}
	if !IsNil(o.AllowBandwidthAutoApproval) {
		toSerialize["allowBandwidthAutoApproval"] = o.AllowBandwidthAutoApproval
	}
	if !IsNil(o.AllowBandwidthUpgrade) {
		toSerialize["allowBandwidthUpgrade"] = o.AllowBandwidthUpgrade
	}
	if !IsNil(o.LinkProtocolConfig) {
		toSerialize["linkProtocolConfig"] = o.LinkProtocolConfig
	}
	if !IsNil(o.EnableAutoGenerateServiceKey) {
		toSerialize["enableAutoGenerateServiceKey"] = o.EnableAutoGenerateServiceKey
	}
	if !IsNil(o.ConnectionRedundancyRequired) {
		toSerialize["connectionRedundancyRequired"] = o.ConnectionRedundancyRequired
	}
	if !IsNil(o.ApiConfig) {
		toSerialize["apiConfig"] = o.ApiConfig
	}
	if !IsNil(o.ConnectionLabel) {
		toSerialize["connectionLabel"] = o.ConnectionLabel
	}
	if !IsNil(o.AuthenticationKey) {
		toSerialize["authenticationKey"] = o.AuthenticationKey
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ServiceProfileAccessPointTypeCOLO) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varServiceProfileAccessPointTypeCOLO := _ServiceProfileAccessPointTypeCOLO{}

	err = json.Unmarshal(data, &varServiceProfileAccessPointTypeCOLO)

	if err != nil {
		return err
	}

	*o = ServiceProfileAccessPointTypeCOLO(varServiceProfileAccessPointTypeCOLO)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "uuid")
		delete(additionalProperties, "type")
		delete(additionalProperties, "supportedBandwidths")
		delete(additionalProperties, "allowRemoteConnections")
		delete(additionalProperties, "allowCustomBandwidth")
		delete(additionalProperties, "bandwidthAlertThreshold")
		delete(additionalProperties, "allowBandwidthAutoApproval")
		delete(additionalProperties, "allowBandwidthUpgrade")
		delete(additionalProperties, "linkProtocolConfig")
		delete(additionalProperties, "enableAutoGenerateServiceKey")
		delete(additionalProperties, "connectionRedundancyRequired")
		delete(additionalProperties, "apiConfig")
		delete(additionalProperties, "connectionLabel")
		delete(additionalProperties, "authenticationKey")
		delete(additionalProperties, "metadata")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableServiceProfileAccessPointTypeCOLO struct {
	value *ServiceProfileAccessPointTypeCOLO
	isSet bool
}

func (v NullableServiceProfileAccessPointTypeCOLO) Get() *ServiceProfileAccessPointTypeCOLO {
	return v.value
}

func (v *NullableServiceProfileAccessPointTypeCOLO) Set(val *ServiceProfileAccessPointTypeCOLO) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceProfileAccessPointTypeCOLO) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceProfileAccessPointTypeCOLO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceProfileAccessPointTypeCOLO(val *ServiceProfileAccessPointTypeCOLO) *NullableServiceProfileAccessPointTypeCOLO {
	return &NullableServiceProfileAccessPointTypeCOLO{value: val, isSet: true}
}

func (v NullableServiceProfileAccessPointTypeCOLO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceProfileAccessPointTypeCOLO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
