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

// checks if the PrecisionTimePackageResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PrecisionTimePackageResponse{}

// PrecisionTimePackageResponse EPT Service Package Information
type PrecisionTimePackageResponse struct {
	Href                 *string                                          `json:"href,omitempty"`
	Code                 GetTimeServicesPackageByCodePackageCodeParameter `json:"code"`
	Type                 *PrecisionTimePackageResponseType                `json:"type,omitempty"`
	Bandwidth            *int32                                           `json:"bandwidth,omitempty"`
	ClientsPerSecondMax  *int32                                           `json:"clientsPerSecondMax,omitempty"`
	RedundancySupported  *bool                                            `json:"redundancySupported,omitempty"`
	MultiSubnetSupported *bool                                            `json:"multiSubnetSupported,omitempty"`
	AccuracyUnit         *string                                          `json:"accuracyUnit,omitempty"`
	AccuracySla          *int32                                           `json:"accuracySla,omitempty"`
	AccuracyAvgMin       *int32                                           `json:"accuracyAvgMin,omitempty"`
	AccuracyAvgMax       *int32                                           `json:"accuracyAvgMax,omitempty"`
	Changelog            *Changelog                                       `json:"changelog,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PrecisionTimePackageResponse PrecisionTimePackageResponse

// NewPrecisionTimePackageResponse instantiates a new PrecisionTimePackageResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrecisionTimePackageResponse(code GetTimeServicesPackageByCodePackageCodeParameter) *PrecisionTimePackageResponse {
	this := PrecisionTimePackageResponse{}
	this.Code = code
	return &this
}

// NewPrecisionTimePackageResponseWithDefaults instantiates a new PrecisionTimePackageResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrecisionTimePackageResponseWithDefaults() *PrecisionTimePackageResponse {
	this := PrecisionTimePackageResponse{}
	return &this
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *PrecisionTimePackageResponse) GetHref() string {
	if o == nil || IsNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrecisionTimePackageResponse) GetHrefOk() (*string, bool) {
	if o == nil || IsNil(o.Href) {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *PrecisionTimePackageResponse) HasHref() bool {
	if o != nil && !IsNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *PrecisionTimePackageResponse) SetHref(v string) {
	o.Href = &v
}

// GetCode returns the Code field value
func (o *PrecisionTimePackageResponse) GetCode() GetTimeServicesPackageByCodePackageCodeParameter {
	if o == nil {
		var ret GetTimeServicesPackageByCodePackageCodeParameter
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *PrecisionTimePackageResponse) GetCodeOk() (*GetTimeServicesPackageByCodePackageCodeParameter, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *PrecisionTimePackageResponse) SetCode(v GetTimeServicesPackageByCodePackageCodeParameter) {
	o.Code = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PrecisionTimePackageResponse) GetType() PrecisionTimePackageResponseType {
	if o == nil || IsNil(o.Type) {
		var ret PrecisionTimePackageResponseType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrecisionTimePackageResponse) GetTypeOk() (*PrecisionTimePackageResponseType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PrecisionTimePackageResponse) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given PrecisionTimePackageResponseType and assigns it to the Type field.
func (o *PrecisionTimePackageResponse) SetType(v PrecisionTimePackageResponseType) {
	o.Type = &v
}

// GetBandwidth returns the Bandwidth field value if set, zero value otherwise.
func (o *PrecisionTimePackageResponse) GetBandwidth() int32 {
	if o == nil || IsNil(o.Bandwidth) {
		var ret int32
		return ret
	}
	return *o.Bandwidth
}

// GetBandwidthOk returns a tuple with the Bandwidth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrecisionTimePackageResponse) GetBandwidthOk() (*int32, bool) {
	if o == nil || IsNil(o.Bandwidth) {
		return nil, false
	}
	return o.Bandwidth, true
}

// HasBandwidth returns a boolean if a field has been set.
func (o *PrecisionTimePackageResponse) HasBandwidth() bool {
	if o != nil && !IsNil(o.Bandwidth) {
		return true
	}

	return false
}

// SetBandwidth gets a reference to the given int32 and assigns it to the Bandwidth field.
func (o *PrecisionTimePackageResponse) SetBandwidth(v int32) {
	o.Bandwidth = &v
}

// GetClientsPerSecondMax returns the ClientsPerSecondMax field value if set, zero value otherwise.
func (o *PrecisionTimePackageResponse) GetClientsPerSecondMax() int32 {
	if o == nil || IsNil(o.ClientsPerSecondMax) {
		var ret int32
		return ret
	}
	return *o.ClientsPerSecondMax
}

// GetClientsPerSecondMaxOk returns a tuple with the ClientsPerSecondMax field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrecisionTimePackageResponse) GetClientsPerSecondMaxOk() (*int32, bool) {
	if o == nil || IsNil(o.ClientsPerSecondMax) {
		return nil, false
	}
	return o.ClientsPerSecondMax, true
}

// HasClientsPerSecondMax returns a boolean if a field has been set.
func (o *PrecisionTimePackageResponse) HasClientsPerSecondMax() bool {
	if o != nil && !IsNil(o.ClientsPerSecondMax) {
		return true
	}

	return false
}

// SetClientsPerSecondMax gets a reference to the given int32 and assigns it to the ClientsPerSecondMax field.
func (o *PrecisionTimePackageResponse) SetClientsPerSecondMax(v int32) {
	o.ClientsPerSecondMax = &v
}

// GetRedundancySupported returns the RedundancySupported field value if set, zero value otherwise.
func (o *PrecisionTimePackageResponse) GetRedundancySupported() bool {
	if o == nil || IsNil(o.RedundancySupported) {
		var ret bool
		return ret
	}
	return *o.RedundancySupported
}

// GetRedundancySupportedOk returns a tuple with the RedundancySupported field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrecisionTimePackageResponse) GetRedundancySupportedOk() (*bool, bool) {
	if o == nil || IsNil(o.RedundancySupported) {
		return nil, false
	}
	return o.RedundancySupported, true
}

// HasRedundancySupported returns a boolean if a field has been set.
func (o *PrecisionTimePackageResponse) HasRedundancySupported() bool {
	if o != nil && !IsNil(o.RedundancySupported) {
		return true
	}

	return false
}

// SetRedundancySupported gets a reference to the given bool and assigns it to the RedundancySupported field.
func (o *PrecisionTimePackageResponse) SetRedundancySupported(v bool) {
	o.RedundancySupported = &v
}

// GetMultiSubnetSupported returns the MultiSubnetSupported field value if set, zero value otherwise.
func (o *PrecisionTimePackageResponse) GetMultiSubnetSupported() bool {
	if o == nil || IsNil(o.MultiSubnetSupported) {
		var ret bool
		return ret
	}
	return *o.MultiSubnetSupported
}

// GetMultiSubnetSupportedOk returns a tuple with the MultiSubnetSupported field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrecisionTimePackageResponse) GetMultiSubnetSupportedOk() (*bool, bool) {
	if o == nil || IsNil(o.MultiSubnetSupported) {
		return nil, false
	}
	return o.MultiSubnetSupported, true
}

// HasMultiSubnetSupported returns a boolean if a field has been set.
func (o *PrecisionTimePackageResponse) HasMultiSubnetSupported() bool {
	if o != nil && !IsNil(o.MultiSubnetSupported) {
		return true
	}

	return false
}

// SetMultiSubnetSupported gets a reference to the given bool and assigns it to the MultiSubnetSupported field.
func (o *PrecisionTimePackageResponse) SetMultiSubnetSupported(v bool) {
	o.MultiSubnetSupported = &v
}

// GetAccuracyUnit returns the AccuracyUnit field value if set, zero value otherwise.
func (o *PrecisionTimePackageResponse) GetAccuracyUnit() string {
	if o == nil || IsNil(o.AccuracyUnit) {
		var ret string
		return ret
	}
	return *o.AccuracyUnit
}

// GetAccuracyUnitOk returns a tuple with the AccuracyUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrecisionTimePackageResponse) GetAccuracyUnitOk() (*string, bool) {
	if o == nil || IsNil(o.AccuracyUnit) {
		return nil, false
	}
	return o.AccuracyUnit, true
}

// HasAccuracyUnit returns a boolean if a field has been set.
func (o *PrecisionTimePackageResponse) HasAccuracyUnit() bool {
	if o != nil && !IsNil(o.AccuracyUnit) {
		return true
	}

	return false
}

// SetAccuracyUnit gets a reference to the given string and assigns it to the AccuracyUnit field.
func (o *PrecisionTimePackageResponse) SetAccuracyUnit(v string) {
	o.AccuracyUnit = &v
}

// GetAccuracySla returns the AccuracySla field value if set, zero value otherwise.
func (o *PrecisionTimePackageResponse) GetAccuracySla() int32 {
	if o == nil || IsNil(o.AccuracySla) {
		var ret int32
		return ret
	}
	return *o.AccuracySla
}

// GetAccuracySlaOk returns a tuple with the AccuracySla field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrecisionTimePackageResponse) GetAccuracySlaOk() (*int32, bool) {
	if o == nil || IsNil(o.AccuracySla) {
		return nil, false
	}
	return o.AccuracySla, true
}

// HasAccuracySla returns a boolean if a field has been set.
func (o *PrecisionTimePackageResponse) HasAccuracySla() bool {
	if o != nil && !IsNil(o.AccuracySla) {
		return true
	}

	return false
}

// SetAccuracySla gets a reference to the given int32 and assigns it to the AccuracySla field.
func (o *PrecisionTimePackageResponse) SetAccuracySla(v int32) {
	o.AccuracySla = &v
}

// GetAccuracyAvgMin returns the AccuracyAvgMin field value if set, zero value otherwise.
func (o *PrecisionTimePackageResponse) GetAccuracyAvgMin() int32 {
	if o == nil || IsNil(o.AccuracyAvgMin) {
		var ret int32
		return ret
	}
	return *o.AccuracyAvgMin
}

// GetAccuracyAvgMinOk returns a tuple with the AccuracyAvgMin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrecisionTimePackageResponse) GetAccuracyAvgMinOk() (*int32, bool) {
	if o == nil || IsNil(o.AccuracyAvgMin) {
		return nil, false
	}
	return o.AccuracyAvgMin, true
}

// HasAccuracyAvgMin returns a boolean if a field has been set.
func (o *PrecisionTimePackageResponse) HasAccuracyAvgMin() bool {
	if o != nil && !IsNil(o.AccuracyAvgMin) {
		return true
	}

	return false
}

// SetAccuracyAvgMin gets a reference to the given int32 and assigns it to the AccuracyAvgMin field.
func (o *PrecisionTimePackageResponse) SetAccuracyAvgMin(v int32) {
	o.AccuracyAvgMin = &v
}

// GetAccuracyAvgMax returns the AccuracyAvgMax field value if set, zero value otherwise.
func (o *PrecisionTimePackageResponse) GetAccuracyAvgMax() int32 {
	if o == nil || IsNil(o.AccuracyAvgMax) {
		var ret int32
		return ret
	}
	return *o.AccuracyAvgMax
}

// GetAccuracyAvgMaxOk returns a tuple with the AccuracyAvgMax field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrecisionTimePackageResponse) GetAccuracyAvgMaxOk() (*int32, bool) {
	if o == nil || IsNil(o.AccuracyAvgMax) {
		return nil, false
	}
	return o.AccuracyAvgMax, true
}

// HasAccuracyAvgMax returns a boolean if a field has been set.
func (o *PrecisionTimePackageResponse) HasAccuracyAvgMax() bool {
	if o != nil && !IsNil(o.AccuracyAvgMax) {
		return true
	}

	return false
}

// SetAccuracyAvgMax gets a reference to the given int32 and assigns it to the AccuracyAvgMax field.
func (o *PrecisionTimePackageResponse) SetAccuracyAvgMax(v int32) {
	o.AccuracyAvgMax = &v
}

// GetChangelog returns the Changelog field value if set, zero value otherwise.
func (o *PrecisionTimePackageResponse) GetChangelog() Changelog {
	if o == nil || IsNil(o.Changelog) {
		var ret Changelog
		return ret
	}
	return *o.Changelog
}

// GetChangelogOk returns a tuple with the Changelog field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrecisionTimePackageResponse) GetChangelogOk() (*Changelog, bool) {
	if o == nil || IsNil(o.Changelog) {
		return nil, false
	}
	return o.Changelog, true
}

// HasChangelog returns a boolean if a field has been set.
func (o *PrecisionTimePackageResponse) HasChangelog() bool {
	if o != nil && !IsNil(o.Changelog) {
		return true
	}

	return false
}

// SetChangelog gets a reference to the given Changelog and assigns it to the Changelog field.
func (o *PrecisionTimePackageResponse) SetChangelog(v Changelog) {
	o.Changelog = &v
}

func (o PrecisionTimePackageResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PrecisionTimePackageResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Href) {
		toSerialize["href"] = o.Href
	}
	toSerialize["code"] = o.Code
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Bandwidth) {
		toSerialize["bandwidth"] = o.Bandwidth
	}
	if !IsNil(o.ClientsPerSecondMax) {
		toSerialize["clientsPerSecondMax"] = o.ClientsPerSecondMax
	}
	if !IsNil(o.RedundancySupported) {
		toSerialize["redundancySupported"] = o.RedundancySupported
	}
	if !IsNil(o.MultiSubnetSupported) {
		toSerialize["multiSubnetSupported"] = o.MultiSubnetSupported
	}
	if !IsNil(o.AccuracyUnit) {
		toSerialize["accuracyUnit"] = o.AccuracyUnit
	}
	if !IsNil(o.AccuracySla) {
		toSerialize["accuracySla"] = o.AccuracySla
	}
	if !IsNil(o.AccuracyAvgMin) {
		toSerialize["accuracyAvgMin"] = o.AccuracyAvgMin
	}
	if !IsNil(o.AccuracyAvgMax) {
		toSerialize["accuracyAvgMax"] = o.AccuracyAvgMax
	}
	if !IsNil(o.Changelog) {
		toSerialize["changelog"] = o.Changelog
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PrecisionTimePackageResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"code",
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

	varPrecisionTimePackageResponse := _PrecisionTimePackageResponse{}

	err = json.Unmarshal(data, &varPrecisionTimePackageResponse)

	if err != nil {
		return err
	}

	*o = PrecisionTimePackageResponse(varPrecisionTimePackageResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "href")
		delete(additionalProperties, "code")
		delete(additionalProperties, "type")
		delete(additionalProperties, "bandwidth")
		delete(additionalProperties, "clientsPerSecondMax")
		delete(additionalProperties, "redundancySupported")
		delete(additionalProperties, "multiSubnetSupported")
		delete(additionalProperties, "accuracyUnit")
		delete(additionalProperties, "accuracySla")
		delete(additionalProperties, "accuracyAvgMin")
		delete(additionalProperties, "accuracyAvgMax")
		delete(additionalProperties, "changelog")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePrecisionTimePackageResponse struct {
	value *PrecisionTimePackageResponse
	isSet bool
}

func (v NullablePrecisionTimePackageResponse) Get() *PrecisionTimePackageResponse {
	return v.value
}

func (v *NullablePrecisionTimePackageResponse) Set(val *PrecisionTimePackageResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePrecisionTimePackageResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePrecisionTimePackageResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrecisionTimePackageResponse(val *PrecisionTimePackageResponse) *NullablePrecisionTimePackageResponse {
	return &NullablePrecisionTimePackageResponse{value: val, isSet: true}
}

func (v NullablePrecisionTimePackageResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrecisionTimePackageResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
