/*
 * Metal API
 *
 * This is the API for Equinix Metal. The API allows you to programmatically interact with all of your Equinix Metal resources, including devices, networks, addresses, organizations, projects, and your user account.  The official API docs are hosted at <https://metal.equinix.com/developers/api>. 
 *
 * API version: 1.0.0
 * Contact: support@equinixmetal.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1

import (
	"encoding/json"
	"time"
)

// HardwareReservation struct for HardwareReservation
type HardwareReservation struct {
	Id *string `json:"id,omitempty"`
	// Short version of the ID.
	ShortId *string `json:"short_id,omitempty"`
	Facility *Facility `json:"facility,omitempty"`
	Plan *Plan `json:"plan,omitempty"`
	Href *string `json:"href,omitempty"`
	Project *Project `json:"project,omitempty"`
	Device *Device `json:"device,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// Whether the Hardware Reservation is a spare. Spare Hardware Reservations are used when a Hardware Reservations requires service from Metal Equinix
	Spare *bool `json:"spare,omitempty"`
	// Whether this Device requires assistance from Metal Equinix.
	NeedOfService *bool `json:"need_of_service,omitempty"`
	// Whether the reserved server is provisionable or not. Spare devices can't be provisioned unless they are activated first.
	Provisionable *bool `json:"provisionable,omitempty"`
	// Amount that will be charged for every billing_cycle.
	CustomRate *float32 `json:"custom_rate,omitempty"`
}

// NewHardwareReservation instantiates a new HardwareReservation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHardwareReservation() *HardwareReservation {
	this := HardwareReservation{}
	return &this
}

// NewHardwareReservationWithDefaults instantiates a new HardwareReservation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHardwareReservationWithDefaults() *HardwareReservation {
	this := HardwareReservation{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *HardwareReservation) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HardwareReservation) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *HardwareReservation) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *HardwareReservation) SetId(v string) {
	o.Id = &v
}

// GetShortId returns the ShortId field value if set, zero value otherwise.
func (o *HardwareReservation) GetShortId() string {
	if o == nil || o.ShortId == nil {
		var ret string
		return ret
	}
	return *o.ShortId
}

// GetShortIdOk returns a tuple with the ShortId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HardwareReservation) GetShortIdOk() (*string, bool) {
	if o == nil || o.ShortId == nil {
		return nil, false
	}
	return o.ShortId, true
}

// HasShortId returns a boolean if a field has been set.
func (o *HardwareReservation) HasShortId() bool {
	if o != nil && o.ShortId != nil {
		return true
	}

	return false
}

// SetShortId gets a reference to the given string and assigns it to the ShortId field.
func (o *HardwareReservation) SetShortId(v string) {
	o.ShortId = &v
}

// GetFacility returns the Facility field value if set, zero value otherwise.
func (o *HardwareReservation) GetFacility() Facility {
	if o == nil || o.Facility == nil {
		var ret Facility
		return ret
	}
	return *o.Facility
}

// GetFacilityOk returns a tuple with the Facility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HardwareReservation) GetFacilityOk() (*Facility, bool) {
	if o == nil || o.Facility == nil {
		return nil, false
	}
	return o.Facility, true
}

// HasFacility returns a boolean if a field has been set.
func (o *HardwareReservation) HasFacility() bool {
	if o != nil && o.Facility != nil {
		return true
	}

	return false
}

// SetFacility gets a reference to the given Facility and assigns it to the Facility field.
func (o *HardwareReservation) SetFacility(v Facility) {
	o.Facility = &v
}

// GetPlan returns the Plan field value if set, zero value otherwise.
func (o *HardwareReservation) GetPlan() Plan {
	if o == nil || o.Plan == nil {
		var ret Plan
		return ret
	}
	return *o.Plan
}

// GetPlanOk returns a tuple with the Plan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HardwareReservation) GetPlanOk() (*Plan, bool) {
	if o == nil || o.Plan == nil {
		return nil, false
	}
	return o.Plan, true
}

// HasPlan returns a boolean if a field has been set.
func (o *HardwareReservation) HasPlan() bool {
	if o != nil && o.Plan != nil {
		return true
	}

	return false
}

// SetPlan gets a reference to the given Plan and assigns it to the Plan field.
func (o *HardwareReservation) SetPlan(v Plan) {
	o.Plan = &v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *HardwareReservation) GetHref() string {
	if o == nil || o.Href == nil {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HardwareReservation) GetHrefOk() (*string, bool) {
	if o == nil || o.Href == nil {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *HardwareReservation) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *HardwareReservation) SetHref(v string) {
	o.Href = &v
}

// GetProject returns the Project field value if set, zero value otherwise.
func (o *HardwareReservation) GetProject() Project {
	if o == nil || o.Project == nil {
		var ret Project
		return ret
	}
	return *o.Project
}

// GetProjectOk returns a tuple with the Project field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HardwareReservation) GetProjectOk() (*Project, bool) {
	if o == nil || o.Project == nil {
		return nil, false
	}
	return o.Project, true
}

// HasProject returns a boolean if a field has been set.
func (o *HardwareReservation) HasProject() bool {
	if o != nil && o.Project != nil {
		return true
	}

	return false
}

// SetProject gets a reference to the given Project and assigns it to the Project field.
func (o *HardwareReservation) SetProject(v Project) {
	o.Project = &v
}

// GetDevice returns the Device field value if set, zero value otherwise.
func (o *HardwareReservation) GetDevice() Device {
	if o == nil || o.Device == nil {
		var ret Device
		return ret
	}
	return *o.Device
}

// GetDeviceOk returns a tuple with the Device field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HardwareReservation) GetDeviceOk() (*Device, bool) {
	if o == nil || o.Device == nil {
		return nil, false
	}
	return o.Device, true
}

// HasDevice returns a boolean if a field has been set.
func (o *HardwareReservation) HasDevice() bool {
	if o != nil && o.Device != nil {
		return true
	}

	return false
}

// SetDevice gets a reference to the given Device and assigns it to the Device field.
func (o *HardwareReservation) SetDevice(v Device) {
	o.Device = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *HardwareReservation) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HardwareReservation) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *HardwareReservation) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *HardwareReservation) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetSpare returns the Spare field value if set, zero value otherwise.
func (o *HardwareReservation) GetSpare() bool {
	if o == nil || o.Spare == nil {
		var ret bool
		return ret
	}
	return *o.Spare
}

// GetSpareOk returns a tuple with the Spare field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HardwareReservation) GetSpareOk() (*bool, bool) {
	if o == nil || o.Spare == nil {
		return nil, false
	}
	return o.Spare, true
}

// HasSpare returns a boolean if a field has been set.
func (o *HardwareReservation) HasSpare() bool {
	if o != nil && o.Spare != nil {
		return true
	}

	return false
}

// SetSpare gets a reference to the given bool and assigns it to the Spare field.
func (o *HardwareReservation) SetSpare(v bool) {
	o.Spare = &v
}

// GetNeedOfService returns the NeedOfService field value if set, zero value otherwise.
func (o *HardwareReservation) GetNeedOfService() bool {
	if o == nil || o.NeedOfService == nil {
		var ret bool
		return ret
	}
	return *o.NeedOfService
}

// GetNeedOfServiceOk returns a tuple with the NeedOfService field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HardwareReservation) GetNeedOfServiceOk() (*bool, bool) {
	if o == nil || o.NeedOfService == nil {
		return nil, false
	}
	return o.NeedOfService, true
}

// HasNeedOfService returns a boolean if a field has been set.
func (o *HardwareReservation) HasNeedOfService() bool {
	if o != nil && o.NeedOfService != nil {
		return true
	}

	return false
}

// SetNeedOfService gets a reference to the given bool and assigns it to the NeedOfService field.
func (o *HardwareReservation) SetNeedOfService(v bool) {
	o.NeedOfService = &v
}

// GetProvisionable returns the Provisionable field value if set, zero value otherwise.
func (o *HardwareReservation) GetProvisionable() bool {
	if o == nil || o.Provisionable == nil {
		var ret bool
		return ret
	}
	return *o.Provisionable
}

// GetProvisionableOk returns a tuple with the Provisionable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HardwareReservation) GetProvisionableOk() (*bool, bool) {
	if o == nil || o.Provisionable == nil {
		return nil, false
	}
	return o.Provisionable, true
}

// HasProvisionable returns a boolean if a field has been set.
func (o *HardwareReservation) HasProvisionable() bool {
	if o != nil && o.Provisionable != nil {
		return true
	}

	return false
}

// SetProvisionable gets a reference to the given bool and assigns it to the Provisionable field.
func (o *HardwareReservation) SetProvisionable(v bool) {
	o.Provisionable = &v
}

// GetCustomRate returns the CustomRate field value if set, zero value otherwise.
func (o *HardwareReservation) GetCustomRate() float32 {
	if o == nil || o.CustomRate == nil {
		var ret float32
		return ret
	}
	return *o.CustomRate
}

// GetCustomRateOk returns a tuple with the CustomRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HardwareReservation) GetCustomRateOk() (*float32, bool) {
	if o == nil || o.CustomRate == nil {
		return nil, false
	}
	return o.CustomRate, true
}

// HasCustomRate returns a boolean if a field has been set.
func (o *HardwareReservation) HasCustomRate() bool {
	if o != nil && o.CustomRate != nil {
		return true
	}

	return false
}

// SetCustomRate gets a reference to the given float32 and assigns it to the CustomRate field.
func (o *HardwareReservation) SetCustomRate(v float32) {
	o.CustomRate = &v
}

func (o HardwareReservation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.ShortId != nil {
		toSerialize["short_id"] = o.ShortId
	}
	if o.Facility != nil {
		toSerialize["facility"] = o.Facility
	}
	if o.Plan != nil {
		toSerialize["plan"] = o.Plan
	}
	if o.Href != nil {
		toSerialize["href"] = o.Href
	}
	if o.Project != nil {
		toSerialize["project"] = o.Project
	}
	if o.Device != nil {
		toSerialize["device"] = o.Device
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.Spare != nil {
		toSerialize["spare"] = o.Spare
	}
	if o.NeedOfService != nil {
		toSerialize["need_of_service"] = o.NeedOfService
	}
	if o.Provisionable != nil {
		toSerialize["provisionable"] = o.Provisionable
	}
	if o.CustomRate != nil {
		toSerialize["custom_rate"] = o.CustomRate
	}
	return json.Marshal(toSerialize)
}

type NullableHardwareReservation struct {
	value *HardwareReservation
	isSet bool
}

func (v NullableHardwareReservation) Get() *HardwareReservation {
	return v.value
}

func (v *NullableHardwareReservation) Set(val *HardwareReservation) {
	v.value = val
	v.isSet = true
}

func (v NullableHardwareReservation) IsSet() bool {
	return v.isSet
}

func (v *NullableHardwareReservation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHardwareReservation(val *HardwareReservation) *NullableHardwareReservation {
	return &NullableHardwareReservation{value: val, isSet: true}
}

func (v NullableHardwareReservation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHardwareReservation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


