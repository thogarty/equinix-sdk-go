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

// checks if the MarketingInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MarketingInfo{}

// MarketingInfo struct for MarketingInfo
type MarketingInfo struct {
	// Logo file name
	Logo *string `json:"logo,omitempty"`
	// Profile promotion on marketplace
	Promotion            *bool         `json:"promotion,omitempty"`
	ProcessSteps         []ProcessStep `json:"processSteps,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MarketingInfo MarketingInfo

// NewMarketingInfo instantiates a new MarketingInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMarketingInfo() *MarketingInfo {
	this := MarketingInfo{}
	return &this
}

// NewMarketingInfoWithDefaults instantiates a new MarketingInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMarketingInfoWithDefaults() *MarketingInfo {
	this := MarketingInfo{}
	return &this
}

// GetLogo returns the Logo field value if set, zero value otherwise.
func (o *MarketingInfo) GetLogo() string {
	if o == nil || IsNil(o.Logo) {
		var ret string
		return ret
	}
	return *o.Logo
}

// GetLogoOk returns a tuple with the Logo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketingInfo) GetLogoOk() (*string, bool) {
	if o == nil || IsNil(o.Logo) {
		return nil, false
	}
	return o.Logo, true
}

// HasLogo returns a boolean if a field has been set.
func (o *MarketingInfo) HasLogo() bool {
	if o != nil && !IsNil(o.Logo) {
		return true
	}

	return false
}

// SetLogo gets a reference to the given string and assigns it to the Logo field.
func (o *MarketingInfo) SetLogo(v string) {
	o.Logo = &v
}

// GetPromotion returns the Promotion field value if set, zero value otherwise.
func (o *MarketingInfo) GetPromotion() bool {
	if o == nil || IsNil(o.Promotion) {
		var ret bool
		return ret
	}
	return *o.Promotion
}

// GetPromotionOk returns a tuple with the Promotion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketingInfo) GetPromotionOk() (*bool, bool) {
	if o == nil || IsNil(o.Promotion) {
		return nil, false
	}
	return o.Promotion, true
}

// HasPromotion returns a boolean if a field has been set.
func (o *MarketingInfo) HasPromotion() bool {
	if o != nil && !IsNil(o.Promotion) {
		return true
	}

	return false
}

// SetPromotion gets a reference to the given bool and assigns it to the Promotion field.
func (o *MarketingInfo) SetPromotion(v bool) {
	o.Promotion = &v
}

// GetProcessSteps returns the ProcessSteps field value if set, zero value otherwise.
func (o *MarketingInfo) GetProcessSteps() []ProcessStep {
	if o == nil || IsNil(o.ProcessSteps) {
		var ret []ProcessStep
		return ret
	}
	return o.ProcessSteps
}

// GetProcessStepsOk returns a tuple with the ProcessSteps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketingInfo) GetProcessStepsOk() ([]ProcessStep, bool) {
	if o == nil || IsNil(o.ProcessSteps) {
		return nil, false
	}
	return o.ProcessSteps, true
}

// HasProcessSteps returns a boolean if a field has been set.
func (o *MarketingInfo) HasProcessSteps() bool {
	if o != nil && !IsNil(o.ProcessSteps) {
		return true
	}

	return false
}

// SetProcessSteps gets a reference to the given []ProcessStep and assigns it to the ProcessSteps field.
func (o *MarketingInfo) SetProcessSteps(v []ProcessStep) {
	o.ProcessSteps = v
}

func (o MarketingInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MarketingInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Logo) {
		toSerialize["logo"] = o.Logo
	}
	if !IsNil(o.Promotion) {
		toSerialize["promotion"] = o.Promotion
	}
	if !IsNil(o.ProcessSteps) {
		toSerialize["processSteps"] = o.ProcessSteps
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MarketingInfo) UnmarshalJSON(data []byte) (err error) {
	varMarketingInfo := _MarketingInfo{}

	err = json.Unmarshal(data, &varMarketingInfo)

	if err != nil {
		return err
	}

	*o = MarketingInfo(varMarketingInfo)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "logo")
		delete(additionalProperties, "promotion")
		delete(additionalProperties, "processSteps")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMarketingInfo struct {
	value *MarketingInfo
	isSet bool
}

func (v NullableMarketingInfo) Get() *MarketingInfo {
	return v.value
}

func (v *NullableMarketingInfo) Set(val *MarketingInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableMarketingInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableMarketingInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarketingInfo(val *MarketingInfo) *NullableMarketingInfo {
	return &NullableMarketingInfo{value: val, isSet: true}
}

func (v NullableMarketingInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarketingInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}