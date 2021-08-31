/*
Metal API

This is the API for Equinix Metal. The API allows you to programmatically interact with all of your Equinix Metal resources, including devices, networks, addresses, organizations, projects, and your user account.  The official API docs are hosted at <https://metal.equinix.com/developers/api>. 

API version: 1.0.0
Contact: support@equinixmetal.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1

import (
	"encoding/json"
)

// MetroCapacityReport struct for MetroCapacityReport
type MetroCapacityReport struct {
	Ny *CapacityPerFacility `json:"ny,omitempty"`
	Sv *CapacityPerFacility `json:"sv,omitempty"`
	Am *CapacityPerFacility `json:"am,omitempty"`
	Ch *CapacityPerFacility `json:"ch,omitempty"`
	La *CapacityPerFacility `json:"la,omitempty"`
	Sg *CapacityPerFacility `json:"sg,omitempty"`
	Da *CapacityPerFacility `json:"da,omitempty"`
}

// NewMetroCapacityReport instantiates a new MetroCapacityReport object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetroCapacityReport() *MetroCapacityReport {
	this := MetroCapacityReport{}
	return &this
}

// NewMetroCapacityReportWithDefaults instantiates a new MetroCapacityReport object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetroCapacityReportWithDefaults() *MetroCapacityReport {
	this := MetroCapacityReport{}
	return &this
}

// GetNy returns the Ny field value if set, zero value otherwise.
func (o *MetroCapacityReport) GetNy() CapacityPerFacility {
	if o == nil || o.Ny == nil {
		var ret CapacityPerFacility
		return ret
	}
	return *o.Ny
}

// GetNyOk returns a tuple with the Ny field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetroCapacityReport) GetNyOk() (*CapacityPerFacility, bool) {
	if o == nil || o.Ny == nil {
		return nil, false
	}
	return o.Ny, true
}

// HasNy returns a boolean if a field has been set.
func (o *MetroCapacityReport) HasNy() bool {
	if o != nil && o.Ny != nil {
		return true
	}

	return false
}

// SetNy gets a reference to the given CapacityPerFacility and assigns it to the Ny field.
func (o *MetroCapacityReport) SetNy(v CapacityPerFacility) {
	o.Ny = &v
}

// GetSv returns the Sv field value if set, zero value otherwise.
func (o *MetroCapacityReport) GetSv() CapacityPerFacility {
	if o == nil || o.Sv == nil {
		var ret CapacityPerFacility
		return ret
	}
	return *o.Sv
}

// GetSvOk returns a tuple with the Sv field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetroCapacityReport) GetSvOk() (*CapacityPerFacility, bool) {
	if o == nil || o.Sv == nil {
		return nil, false
	}
	return o.Sv, true
}

// HasSv returns a boolean if a field has been set.
func (o *MetroCapacityReport) HasSv() bool {
	if o != nil && o.Sv != nil {
		return true
	}

	return false
}

// SetSv gets a reference to the given CapacityPerFacility and assigns it to the Sv field.
func (o *MetroCapacityReport) SetSv(v CapacityPerFacility) {
	o.Sv = &v
}

// GetAm returns the Am field value if set, zero value otherwise.
func (o *MetroCapacityReport) GetAm() CapacityPerFacility {
	if o == nil || o.Am == nil {
		var ret CapacityPerFacility
		return ret
	}
	return *o.Am
}

// GetAmOk returns a tuple with the Am field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetroCapacityReport) GetAmOk() (*CapacityPerFacility, bool) {
	if o == nil || o.Am == nil {
		return nil, false
	}
	return o.Am, true
}

// HasAm returns a boolean if a field has been set.
func (o *MetroCapacityReport) HasAm() bool {
	if o != nil && o.Am != nil {
		return true
	}

	return false
}

// SetAm gets a reference to the given CapacityPerFacility and assigns it to the Am field.
func (o *MetroCapacityReport) SetAm(v CapacityPerFacility) {
	o.Am = &v
}

// GetCh returns the Ch field value if set, zero value otherwise.
func (o *MetroCapacityReport) GetCh() CapacityPerFacility {
	if o == nil || o.Ch == nil {
		var ret CapacityPerFacility
		return ret
	}
	return *o.Ch
}

// GetChOk returns a tuple with the Ch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetroCapacityReport) GetChOk() (*CapacityPerFacility, bool) {
	if o == nil || o.Ch == nil {
		return nil, false
	}
	return o.Ch, true
}

// HasCh returns a boolean if a field has been set.
func (o *MetroCapacityReport) HasCh() bool {
	if o != nil && o.Ch != nil {
		return true
	}

	return false
}

// SetCh gets a reference to the given CapacityPerFacility and assigns it to the Ch field.
func (o *MetroCapacityReport) SetCh(v CapacityPerFacility) {
	o.Ch = &v
}

// GetLa returns the La field value if set, zero value otherwise.
func (o *MetroCapacityReport) GetLa() CapacityPerFacility {
	if o == nil || o.La == nil {
		var ret CapacityPerFacility
		return ret
	}
	return *o.La
}

// GetLaOk returns a tuple with the La field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetroCapacityReport) GetLaOk() (*CapacityPerFacility, bool) {
	if o == nil || o.La == nil {
		return nil, false
	}
	return o.La, true
}

// HasLa returns a boolean if a field has been set.
func (o *MetroCapacityReport) HasLa() bool {
	if o != nil && o.La != nil {
		return true
	}

	return false
}

// SetLa gets a reference to the given CapacityPerFacility and assigns it to the La field.
func (o *MetroCapacityReport) SetLa(v CapacityPerFacility) {
	o.La = &v
}

// GetSg returns the Sg field value if set, zero value otherwise.
func (o *MetroCapacityReport) GetSg() CapacityPerFacility {
	if o == nil || o.Sg == nil {
		var ret CapacityPerFacility
		return ret
	}
	return *o.Sg
}

// GetSgOk returns a tuple with the Sg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetroCapacityReport) GetSgOk() (*CapacityPerFacility, bool) {
	if o == nil || o.Sg == nil {
		return nil, false
	}
	return o.Sg, true
}

// HasSg returns a boolean if a field has been set.
func (o *MetroCapacityReport) HasSg() bool {
	if o != nil && o.Sg != nil {
		return true
	}

	return false
}

// SetSg gets a reference to the given CapacityPerFacility and assigns it to the Sg field.
func (o *MetroCapacityReport) SetSg(v CapacityPerFacility) {
	o.Sg = &v
}

// GetDa returns the Da field value if set, zero value otherwise.
func (o *MetroCapacityReport) GetDa() CapacityPerFacility {
	if o == nil || o.Da == nil {
		var ret CapacityPerFacility
		return ret
	}
	return *o.Da
}

// GetDaOk returns a tuple with the Da field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetroCapacityReport) GetDaOk() (*CapacityPerFacility, bool) {
	if o == nil || o.Da == nil {
		return nil, false
	}
	return o.Da, true
}

// HasDa returns a boolean if a field has been set.
func (o *MetroCapacityReport) HasDa() bool {
	if o != nil && o.Da != nil {
		return true
	}

	return false
}

// SetDa gets a reference to the given CapacityPerFacility and assigns it to the Da field.
func (o *MetroCapacityReport) SetDa(v CapacityPerFacility) {
	o.Da = &v
}

func (o MetroCapacityReport) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Ny != nil {
		toSerialize["ny"] = o.Ny
	}
	if o.Sv != nil {
		toSerialize["sv"] = o.Sv
	}
	if o.Am != nil {
		toSerialize["am"] = o.Am
	}
	if o.Ch != nil {
		toSerialize["ch"] = o.Ch
	}
	if o.La != nil {
		toSerialize["la"] = o.La
	}
	if o.Sg != nil {
		toSerialize["sg"] = o.Sg
	}
	if o.Da != nil {
		toSerialize["da"] = o.Da
	}
	return json.Marshal(toSerialize)
}

type NullableMetroCapacityReport struct {
	value *MetroCapacityReport
	isSet bool
}

func (v NullableMetroCapacityReport) Get() *MetroCapacityReport {
	return v.value
}

func (v *NullableMetroCapacityReport) Set(val *MetroCapacityReport) {
	v.value = val
	v.isSet = true
}

func (v NullableMetroCapacityReport) IsSet() bool {
	return v.isSet
}

func (v *NullableMetroCapacityReport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetroCapacityReport(val *MetroCapacityReport) *NullableMetroCapacityReport {
	return &NullableMetroCapacityReport{value: val, isSet: true}
}

func (v NullableMetroCapacityReport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetroCapacityReport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


