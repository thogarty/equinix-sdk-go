/*
Equinix Fabric API v4

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fabricv4

import (
	"encoding/json"
	"fmt"
)

// checks if the PrecisionTimeServiceResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PrecisionTimeServiceResponse{}

// PrecisionTimeServiceResponse Precision Time Service Response Schema.
type PrecisionTimeServiceResponse struct {
	// Precision Time Service URI.
	Href string                           `json:"href"`
	Type PrecisionTimeServiceResponseType `json:"type"`
	// Precision Time Service Name.
	Name *string `json:"name,omitempty"`
	// Precision Time Service UUID.
	Uuid    string                            `json:"uuid"`
	State   PrecisionTimeServiceResponseState `json:"state"`
	Package PrecisionTimePackageResponse      `json:"package"`
	// Fabric Connections associated with Precision Time Service.
	Connections []VirtualConnectionTimeServiceResponse `json:"connections,omitempty"`
	Ipv4        *Ipv4                                  `json:"ipv4,omitempty"`
	// NTP Advanced configuration - MD5 Authentication.
	NtpAdvancedConfiguration []Md5                    `json:"ntpAdvancedConfiguration,omitempty"`
	PtpAdvancedConfiguration *PtpAdvanceConfiguration `json:"ptpAdvancedConfiguration,omitempty"`
	Project                  *Project                 `json:"project,omitempty"`
	Account                  *SimplifiedAccount       `json:"account,omitempty"`
	Order                    *PrecisionTimeOrder      `json:"order,omitempty"`
	ChangeLog                *Changelog               `json:"changeLog,omitempty"`
	AdditionalProperties     map[string]interface{}
}

type _PrecisionTimeServiceResponse PrecisionTimeServiceResponse

// NewPrecisionTimeServiceResponse instantiates a new PrecisionTimeServiceResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrecisionTimeServiceResponse(href string, type_ PrecisionTimeServiceResponseType, uuid string, state PrecisionTimeServiceResponseState, package_ PrecisionTimePackageResponse) *PrecisionTimeServiceResponse {
	this := PrecisionTimeServiceResponse{}
	this.Href = href
	this.Type = type_
	this.Uuid = uuid
	this.State = state
	this.Package = package_
	return &this
}

// NewPrecisionTimeServiceResponseWithDefaults instantiates a new PrecisionTimeServiceResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrecisionTimeServiceResponseWithDefaults() *PrecisionTimeServiceResponse {
	this := PrecisionTimeServiceResponse{}
	return &this
}

// GetHref returns the Href field value
func (o *PrecisionTimeServiceResponse) GetHref() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Href
}

// GetHrefOk returns a tuple with the Href field value
// and a boolean to check if the value has been set.
func (o *PrecisionTimeServiceResponse) GetHrefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Href, true
}

// SetHref sets field value
func (o *PrecisionTimeServiceResponse) SetHref(v string) {
	o.Href = v
}

// GetType returns the Type field value
func (o *PrecisionTimeServiceResponse) GetType() PrecisionTimeServiceResponseType {
	if o == nil {
		var ret PrecisionTimeServiceResponseType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *PrecisionTimeServiceResponse) GetTypeOk() (*PrecisionTimeServiceResponseType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PrecisionTimeServiceResponse) SetType(v PrecisionTimeServiceResponseType) {
	o.Type = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PrecisionTimeServiceResponse) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrecisionTimeServiceResponse) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PrecisionTimeServiceResponse) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PrecisionTimeServiceResponse) SetName(v string) {
	o.Name = &v
}

// GetUuid returns the Uuid field value
func (o *PrecisionTimeServiceResponse) GetUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value
// and a boolean to check if the value has been set.
func (o *PrecisionTimeServiceResponse) GetUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uuid, true
}

// SetUuid sets field value
func (o *PrecisionTimeServiceResponse) SetUuid(v string) {
	o.Uuid = v
}

// GetState returns the State field value
func (o *PrecisionTimeServiceResponse) GetState() PrecisionTimeServiceResponseState {
	if o == nil {
		var ret PrecisionTimeServiceResponseState
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *PrecisionTimeServiceResponse) GetStateOk() (*PrecisionTimeServiceResponseState, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *PrecisionTimeServiceResponse) SetState(v PrecisionTimeServiceResponseState) {
	o.State = v
}

// GetPackage returns the Package field value
func (o *PrecisionTimeServiceResponse) GetPackage() PrecisionTimePackageResponse {
	if o == nil {
		var ret PrecisionTimePackageResponse
		return ret
	}

	return o.Package
}

// GetPackageOk returns a tuple with the Package field value
// and a boolean to check if the value has been set.
func (o *PrecisionTimeServiceResponse) GetPackageOk() (*PrecisionTimePackageResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Package, true
}

// SetPackage sets field value
func (o *PrecisionTimeServiceResponse) SetPackage(v PrecisionTimePackageResponse) {
	o.Package = v
}

// GetConnections returns the Connections field value if set, zero value otherwise.
func (o *PrecisionTimeServiceResponse) GetConnections() []VirtualConnectionTimeServiceResponse {
	if o == nil || IsNil(o.Connections) {
		var ret []VirtualConnectionTimeServiceResponse
		return ret
	}
	return o.Connections
}

// GetConnectionsOk returns a tuple with the Connections field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrecisionTimeServiceResponse) GetConnectionsOk() ([]VirtualConnectionTimeServiceResponse, bool) {
	if o == nil || IsNil(o.Connections) {
		return nil, false
	}
	return o.Connections, true
}

// HasConnections returns a boolean if a field has been set.
func (o *PrecisionTimeServiceResponse) HasConnections() bool {
	if o != nil && !IsNil(o.Connections) {
		return true
	}

	return false
}

// SetConnections gets a reference to the given []VirtualConnectionTimeServiceResponse and assigns it to the Connections field.
func (o *PrecisionTimeServiceResponse) SetConnections(v []VirtualConnectionTimeServiceResponse) {
	o.Connections = v
}

// GetIpv4 returns the Ipv4 field value if set, zero value otherwise.
func (o *PrecisionTimeServiceResponse) GetIpv4() Ipv4 {
	if o == nil || IsNil(o.Ipv4) {
		var ret Ipv4
		return ret
	}
	return *o.Ipv4
}

// GetIpv4Ok returns a tuple with the Ipv4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrecisionTimeServiceResponse) GetIpv4Ok() (*Ipv4, bool) {
	if o == nil || IsNil(o.Ipv4) {
		return nil, false
	}
	return o.Ipv4, true
}

// HasIpv4 returns a boolean if a field has been set.
func (o *PrecisionTimeServiceResponse) HasIpv4() bool {
	if o != nil && !IsNil(o.Ipv4) {
		return true
	}

	return false
}

// SetIpv4 gets a reference to the given Ipv4 and assigns it to the Ipv4 field.
func (o *PrecisionTimeServiceResponse) SetIpv4(v Ipv4) {
	o.Ipv4 = &v
}

// GetNtpAdvancedConfiguration returns the NtpAdvancedConfiguration field value if set, zero value otherwise.
func (o *PrecisionTimeServiceResponse) GetNtpAdvancedConfiguration() []Md5 {
	if o == nil || IsNil(o.NtpAdvancedConfiguration) {
		var ret []Md5
		return ret
	}
	return o.NtpAdvancedConfiguration
}

// GetNtpAdvancedConfigurationOk returns a tuple with the NtpAdvancedConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrecisionTimeServiceResponse) GetNtpAdvancedConfigurationOk() ([]Md5, bool) {
	if o == nil || IsNil(o.NtpAdvancedConfiguration) {
		return nil, false
	}
	return o.NtpAdvancedConfiguration, true
}

// HasNtpAdvancedConfiguration returns a boolean if a field has been set.
func (o *PrecisionTimeServiceResponse) HasNtpAdvancedConfiguration() bool {
	if o != nil && !IsNil(o.NtpAdvancedConfiguration) {
		return true
	}

	return false
}

// SetNtpAdvancedConfiguration gets a reference to the given []Md5 and assigns it to the NtpAdvancedConfiguration field.
func (o *PrecisionTimeServiceResponse) SetNtpAdvancedConfiguration(v []Md5) {
	o.NtpAdvancedConfiguration = v
}

// GetPtpAdvancedConfiguration returns the PtpAdvancedConfiguration field value if set, zero value otherwise.
func (o *PrecisionTimeServiceResponse) GetPtpAdvancedConfiguration() PtpAdvanceConfiguration {
	if o == nil || IsNil(o.PtpAdvancedConfiguration) {
		var ret PtpAdvanceConfiguration
		return ret
	}
	return *o.PtpAdvancedConfiguration
}

// GetPtpAdvancedConfigurationOk returns a tuple with the PtpAdvancedConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrecisionTimeServiceResponse) GetPtpAdvancedConfigurationOk() (*PtpAdvanceConfiguration, bool) {
	if o == nil || IsNil(o.PtpAdvancedConfiguration) {
		return nil, false
	}
	return o.PtpAdvancedConfiguration, true
}

// HasPtpAdvancedConfiguration returns a boolean if a field has been set.
func (o *PrecisionTimeServiceResponse) HasPtpAdvancedConfiguration() bool {
	if o != nil && !IsNil(o.PtpAdvancedConfiguration) {
		return true
	}

	return false
}

// SetPtpAdvancedConfiguration gets a reference to the given PtpAdvanceConfiguration and assigns it to the PtpAdvancedConfiguration field.
func (o *PrecisionTimeServiceResponse) SetPtpAdvancedConfiguration(v PtpAdvanceConfiguration) {
	o.PtpAdvancedConfiguration = &v
}

// GetProject returns the Project field value if set, zero value otherwise.
func (o *PrecisionTimeServiceResponse) GetProject() Project {
	if o == nil || IsNil(o.Project) {
		var ret Project
		return ret
	}
	return *o.Project
}

// GetProjectOk returns a tuple with the Project field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrecisionTimeServiceResponse) GetProjectOk() (*Project, bool) {
	if o == nil || IsNil(o.Project) {
		return nil, false
	}
	return o.Project, true
}

// HasProject returns a boolean if a field has been set.
func (o *PrecisionTimeServiceResponse) HasProject() bool {
	if o != nil && !IsNil(o.Project) {
		return true
	}

	return false
}

// SetProject gets a reference to the given Project and assigns it to the Project field.
func (o *PrecisionTimeServiceResponse) SetProject(v Project) {
	o.Project = &v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *PrecisionTimeServiceResponse) GetAccount() SimplifiedAccount {
	if o == nil || IsNil(o.Account) {
		var ret SimplifiedAccount
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrecisionTimeServiceResponse) GetAccountOk() (*SimplifiedAccount, bool) {
	if o == nil || IsNil(o.Account) {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *PrecisionTimeServiceResponse) HasAccount() bool {
	if o != nil && !IsNil(o.Account) {
		return true
	}

	return false
}

// SetAccount gets a reference to the given SimplifiedAccount and assigns it to the Account field.
func (o *PrecisionTimeServiceResponse) SetAccount(v SimplifiedAccount) {
	o.Account = &v
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *PrecisionTimeServiceResponse) GetOrder() PrecisionTimeOrder {
	if o == nil || IsNil(o.Order) {
		var ret PrecisionTimeOrder
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrecisionTimeServiceResponse) GetOrderOk() (*PrecisionTimeOrder, bool) {
	if o == nil || IsNil(o.Order) {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *PrecisionTimeServiceResponse) HasOrder() bool {
	if o != nil && !IsNil(o.Order) {
		return true
	}

	return false
}

// SetOrder gets a reference to the given PrecisionTimeOrder and assigns it to the Order field.
func (o *PrecisionTimeServiceResponse) SetOrder(v PrecisionTimeOrder) {
	o.Order = &v
}

// GetChangeLog returns the ChangeLog field value if set, zero value otherwise.
func (o *PrecisionTimeServiceResponse) GetChangeLog() Changelog {
	if o == nil || IsNil(o.ChangeLog) {
		var ret Changelog
		return ret
	}
	return *o.ChangeLog
}

// GetChangeLogOk returns a tuple with the ChangeLog field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrecisionTimeServiceResponse) GetChangeLogOk() (*Changelog, bool) {
	if o == nil || IsNil(o.ChangeLog) {
		return nil, false
	}
	return o.ChangeLog, true
}

// HasChangeLog returns a boolean if a field has been set.
func (o *PrecisionTimeServiceResponse) HasChangeLog() bool {
	if o != nil && !IsNil(o.ChangeLog) {
		return true
	}

	return false
}

// SetChangeLog gets a reference to the given Changelog and assigns it to the ChangeLog field.
func (o *PrecisionTimeServiceResponse) SetChangeLog(v Changelog) {
	o.ChangeLog = &v
}

func (o PrecisionTimeServiceResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PrecisionTimeServiceResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["href"] = o.Href
	toSerialize["type"] = o.Type
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	toSerialize["uuid"] = o.Uuid
	toSerialize["state"] = o.State
	toSerialize["package"] = o.Package
	if !IsNil(o.Connections) {
		toSerialize["connections"] = o.Connections
	}
	if !IsNil(o.Ipv4) {
		toSerialize["ipv4"] = o.Ipv4
	}
	if !IsNil(o.NtpAdvancedConfiguration) {
		toSerialize["ntpAdvancedConfiguration"] = o.NtpAdvancedConfiguration
	}
	if !IsNil(o.PtpAdvancedConfiguration) {
		toSerialize["ptpAdvancedConfiguration"] = o.PtpAdvancedConfiguration
	}
	if !IsNil(o.Project) {
		toSerialize["project"] = o.Project
	}
	if !IsNil(o.Account) {
		toSerialize["account"] = o.Account
	}
	if !IsNil(o.Order) {
		toSerialize["order"] = o.Order
	}
	if !IsNil(o.ChangeLog) {
		toSerialize["changeLog"] = o.ChangeLog
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PrecisionTimeServiceResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"href",
		"type",
		"uuid",
		"state",
		"package",
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

	varPrecisionTimeServiceResponse := _PrecisionTimeServiceResponse{}

	err = json.Unmarshal(data, &varPrecisionTimeServiceResponse)

	if err != nil {
		return err
	}

	*o = PrecisionTimeServiceResponse(varPrecisionTimeServiceResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "href")
		delete(additionalProperties, "type")
		delete(additionalProperties, "name")
		delete(additionalProperties, "uuid")
		delete(additionalProperties, "state")
		delete(additionalProperties, "package")
		delete(additionalProperties, "connections")
		delete(additionalProperties, "ipv4")
		delete(additionalProperties, "ntpAdvancedConfiguration")
		delete(additionalProperties, "ptpAdvancedConfiguration")
		delete(additionalProperties, "project")
		delete(additionalProperties, "account")
		delete(additionalProperties, "order")
		delete(additionalProperties, "changeLog")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePrecisionTimeServiceResponse struct {
	value *PrecisionTimeServiceResponse
	isSet bool
}

func (v NullablePrecisionTimeServiceResponse) Get() *PrecisionTimeServiceResponse {
	return v.value
}

func (v *NullablePrecisionTimeServiceResponse) Set(val *PrecisionTimeServiceResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePrecisionTimeServiceResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePrecisionTimeServiceResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrecisionTimeServiceResponse(val *PrecisionTimeServiceResponse) *NullablePrecisionTimeServiceResponse {
	return &NullablePrecisionTimeServiceResponse{value: val, isSet: true}
}

func (v NullablePrecisionTimeServiceResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrecisionTimeServiceResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
