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

// VolumeCreateInput struct for VolumeCreateInput
type VolumeCreateInput struct {
	Description *string `json:"description,omitempty"`
	// ams1, ewr1, nrt1, sjc1
	Facility string `json:"facility"`
	// storage_1, storage_2
	Plan string `json:"plan"`
	Size int32 `json:"size"`
	Locked *bool `json:"locked,omitempty"`
	// hourly
	BillingCycle *string `json:"billing_cycle,omitempty"`
	SnapshotPolicies *SnapshotPolicyInput `json:"snapshot_policies,omitempty"`
	Customdata *map[string]interface{} `json:"customdata,omitempty"`
}

// NewVolumeCreateInput instantiates a new VolumeCreateInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVolumeCreateInput(facility string, plan string, size int32) *VolumeCreateInput {
	this := VolumeCreateInput{}
	this.Facility = facility
	this.Plan = plan
	this.Size = size
	return &this
}

// NewVolumeCreateInputWithDefaults instantiates a new VolumeCreateInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVolumeCreateInputWithDefaults() *VolumeCreateInput {
	this := VolumeCreateInput{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *VolumeCreateInput) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumeCreateInput) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *VolumeCreateInput) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *VolumeCreateInput) SetDescription(v string) {
	o.Description = &v
}

// GetFacility returns the Facility field value
func (o *VolumeCreateInput) GetFacility() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Facility
}

// GetFacilityOk returns a tuple with the Facility field value
// and a boolean to check if the value has been set.
func (o *VolumeCreateInput) GetFacilityOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Facility, true
}

// SetFacility sets field value
func (o *VolumeCreateInput) SetFacility(v string) {
	o.Facility = v
}

// GetPlan returns the Plan field value
func (o *VolumeCreateInput) GetPlan() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Plan
}

// GetPlanOk returns a tuple with the Plan field value
// and a boolean to check if the value has been set.
func (o *VolumeCreateInput) GetPlanOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Plan, true
}

// SetPlan sets field value
func (o *VolumeCreateInput) SetPlan(v string) {
	o.Plan = v
}

// GetSize returns the Size field value
func (o *VolumeCreateInput) GetSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Size
}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
func (o *VolumeCreateInput) GetSizeOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Size, true
}

// SetSize sets field value
func (o *VolumeCreateInput) SetSize(v int32) {
	o.Size = v
}

// GetLocked returns the Locked field value if set, zero value otherwise.
func (o *VolumeCreateInput) GetLocked() bool {
	if o == nil || o.Locked == nil {
		var ret bool
		return ret
	}
	return *o.Locked
}

// GetLockedOk returns a tuple with the Locked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumeCreateInput) GetLockedOk() (*bool, bool) {
	if o == nil || o.Locked == nil {
		return nil, false
	}
	return o.Locked, true
}

// HasLocked returns a boolean if a field has been set.
func (o *VolumeCreateInput) HasLocked() bool {
	if o != nil && o.Locked != nil {
		return true
	}

	return false
}

// SetLocked gets a reference to the given bool and assigns it to the Locked field.
func (o *VolumeCreateInput) SetLocked(v bool) {
	o.Locked = &v
}

// GetBillingCycle returns the BillingCycle field value if set, zero value otherwise.
func (o *VolumeCreateInput) GetBillingCycle() string {
	if o == nil || o.BillingCycle == nil {
		var ret string
		return ret
	}
	return *o.BillingCycle
}

// GetBillingCycleOk returns a tuple with the BillingCycle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumeCreateInput) GetBillingCycleOk() (*string, bool) {
	if o == nil || o.BillingCycle == nil {
		return nil, false
	}
	return o.BillingCycle, true
}

// HasBillingCycle returns a boolean if a field has been set.
func (o *VolumeCreateInput) HasBillingCycle() bool {
	if o != nil && o.BillingCycle != nil {
		return true
	}

	return false
}

// SetBillingCycle gets a reference to the given string and assigns it to the BillingCycle field.
func (o *VolumeCreateInput) SetBillingCycle(v string) {
	o.BillingCycle = &v
}

// GetSnapshotPolicies returns the SnapshotPolicies field value if set, zero value otherwise.
func (o *VolumeCreateInput) GetSnapshotPolicies() SnapshotPolicyInput {
	if o == nil || o.SnapshotPolicies == nil {
		var ret SnapshotPolicyInput
		return ret
	}
	return *o.SnapshotPolicies
}

// GetSnapshotPoliciesOk returns a tuple with the SnapshotPolicies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumeCreateInput) GetSnapshotPoliciesOk() (*SnapshotPolicyInput, bool) {
	if o == nil || o.SnapshotPolicies == nil {
		return nil, false
	}
	return o.SnapshotPolicies, true
}

// HasSnapshotPolicies returns a boolean if a field has been set.
func (o *VolumeCreateInput) HasSnapshotPolicies() bool {
	if o != nil && o.SnapshotPolicies != nil {
		return true
	}

	return false
}

// SetSnapshotPolicies gets a reference to the given SnapshotPolicyInput and assigns it to the SnapshotPolicies field.
func (o *VolumeCreateInput) SetSnapshotPolicies(v SnapshotPolicyInput) {
	o.SnapshotPolicies = &v
}

// GetCustomdata returns the Customdata field value if set, zero value otherwise.
func (o *VolumeCreateInput) GetCustomdata() map[string]interface{} {
	if o == nil || o.Customdata == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Customdata
}

// GetCustomdataOk returns a tuple with the Customdata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumeCreateInput) GetCustomdataOk() (*map[string]interface{}, bool) {
	if o == nil || o.Customdata == nil {
		return nil, false
	}
	return o.Customdata, true
}

// HasCustomdata returns a boolean if a field has been set.
func (o *VolumeCreateInput) HasCustomdata() bool {
	if o != nil && o.Customdata != nil {
		return true
	}

	return false
}

// SetCustomdata gets a reference to the given map[string]interface{} and assigns it to the Customdata field.
func (o *VolumeCreateInput) SetCustomdata(v map[string]interface{}) {
	o.Customdata = &v
}

func (o VolumeCreateInput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["facility"] = o.Facility
	}
	if true {
		toSerialize["plan"] = o.Plan
	}
	if true {
		toSerialize["size"] = o.Size
	}
	if o.Locked != nil {
		toSerialize["locked"] = o.Locked
	}
	if o.BillingCycle != nil {
		toSerialize["billing_cycle"] = o.BillingCycle
	}
	if o.SnapshotPolicies != nil {
		toSerialize["snapshot_policies"] = o.SnapshotPolicies
	}
	if o.Customdata != nil {
		toSerialize["customdata"] = o.Customdata
	}
	return json.Marshal(toSerialize)
}

type NullableVolumeCreateInput struct {
	value *VolumeCreateInput
	isSet bool
}

func (v NullableVolumeCreateInput) Get() *VolumeCreateInput {
	return v.value
}

func (v *NullableVolumeCreateInput) Set(val *VolumeCreateInput) {
	v.value = val
	v.isSet = true
}

func (v NullableVolumeCreateInput) IsSet() bool {
	return v.isSet
}

func (v *NullableVolumeCreateInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVolumeCreateInput(val *VolumeCreateInput) *NullableVolumeCreateInput {
	return &NullableVolumeCreateInput{value: val, isSet: true}
}

func (v NullableVolumeCreateInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVolumeCreateInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


