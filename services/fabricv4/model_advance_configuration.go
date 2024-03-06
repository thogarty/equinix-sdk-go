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

// checks if the AdvanceConfiguration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdvanceConfiguration{}

// AdvanceConfiguration Advance Configuration for NTP/PTP
type AdvanceConfiguration struct {
	Ntp                  []Md5                    `json:"ntp,omitempty"`
	Ptp                  *PtpAdvanceConfiguration `json:"ptp,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AdvanceConfiguration AdvanceConfiguration

// NewAdvanceConfiguration instantiates a new AdvanceConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdvanceConfiguration() *AdvanceConfiguration {
	this := AdvanceConfiguration{}
	return &this
}

// NewAdvanceConfigurationWithDefaults instantiates a new AdvanceConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdvanceConfigurationWithDefaults() *AdvanceConfiguration {
	this := AdvanceConfiguration{}
	return &this
}

// GetNtp returns the Ntp field value if set, zero value otherwise.
func (o *AdvanceConfiguration) GetNtp() []Md5 {
	if o == nil || IsNil(o.Ntp) {
		var ret []Md5
		return ret
	}
	return o.Ntp
}

// GetNtpOk returns a tuple with the Ntp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvanceConfiguration) GetNtpOk() ([]Md5, bool) {
	if o == nil || IsNil(o.Ntp) {
		return nil, false
	}
	return o.Ntp, true
}

// HasNtp returns a boolean if a field has been set.
func (o *AdvanceConfiguration) HasNtp() bool {
	if o != nil && !IsNil(o.Ntp) {
		return true
	}

	return false
}

// SetNtp gets a reference to the given []Md5 and assigns it to the Ntp field.
func (o *AdvanceConfiguration) SetNtp(v []Md5) {
	o.Ntp = v
}

// GetPtp returns the Ptp field value if set, zero value otherwise.
func (o *AdvanceConfiguration) GetPtp() PtpAdvanceConfiguration {
	if o == nil || IsNil(o.Ptp) {
		var ret PtpAdvanceConfiguration
		return ret
	}
	return *o.Ptp
}

// GetPtpOk returns a tuple with the Ptp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvanceConfiguration) GetPtpOk() (*PtpAdvanceConfiguration, bool) {
	if o == nil || IsNil(o.Ptp) {
		return nil, false
	}
	return o.Ptp, true
}

// HasPtp returns a boolean if a field has been set.
func (o *AdvanceConfiguration) HasPtp() bool {
	if o != nil && !IsNil(o.Ptp) {
		return true
	}

	return false
}

// SetPtp gets a reference to the given PtpAdvanceConfiguration and assigns it to the Ptp field.
func (o *AdvanceConfiguration) SetPtp(v PtpAdvanceConfiguration) {
	o.Ptp = &v
}

func (o AdvanceConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdvanceConfiguration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ntp) {
		toSerialize["ntp"] = o.Ntp
	}
	if !IsNil(o.Ptp) {
		toSerialize["ptp"] = o.Ptp
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AdvanceConfiguration) UnmarshalJSON(data []byte) (err error) {
	varAdvanceConfiguration := _AdvanceConfiguration{}

	err = json.Unmarshal(data, &varAdvanceConfiguration)

	if err != nil {
		return err
	}

	*o = AdvanceConfiguration(varAdvanceConfiguration)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ntp")
		delete(additionalProperties, "ptp")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAdvanceConfiguration struct {
	value *AdvanceConfiguration
	isSet bool
}

func (v NullableAdvanceConfiguration) Get() *AdvanceConfiguration {
	return v.value
}

func (v *NullableAdvanceConfiguration) Set(val *AdvanceConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableAdvanceConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableAdvanceConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdvanceConfiguration(val *AdvanceConfiguration) *NullableAdvanceConfiguration {
	return &NullableAdvanceConfiguration{value: val, isSet: true}
}

func (v NullableAdvanceConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdvanceConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}