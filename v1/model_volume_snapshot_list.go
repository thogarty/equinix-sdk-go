/*
 * Metal API
 *
 * This is the API for Equinix Metal Product. Interact with your devices, user account, and projects.
 *
 * API version: 1.0.0
 * Contact: support@equinixmetal.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1

import (
	"encoding/json"
)

// VolumeSnapshotList struct for VolumeSnapshotList
type VolumeSnapshotList struct {
	Snapshots *[]VolumeSnapshot `json:"snapshots,omitempty"`
}

// NewVolumeSnapshotList instantiates a new VolumeSnapshotList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVolumeSnapshotList() *VolumeSnapshotList {
	this := VolumeSnapshotList{}
	return &this
}

// NewVolumeSnapshotListWithDefaults instantiates a new VolumeSnapshotList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVolumeSnapshotListWithDefaults() *VolumeSnapshotList {
	this := VolumeSnapshotList{}
	return &this
}

// GetSnapshots returns the Snapshots field value if set, zero value otherwise.
func (o *VolumeSnapshotList) GetSnapshots() []VolumeSnapshot {
	if o == nil || o.Snapshots == nil {
		var ret []VolumeSnapshot
		return ret
	}
	return *o.Snapshots
}

// GetSnapshotsOk returns a tuple with the Snapshots field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumeSnapshotList) GetSnapshotsOk() (*[]VolumeSnapshot, bool) {
	if o == nil || o.Snapshots == nil {
		return nil, false
	}
	return o.Snapshots, true
}

// HasSnapshots returns a boolean if a field has been set.
func (o *VolumeSnapshotList) HasSnapshots() bool {
	if o != nil && o.Snapshots != nil {
		return true
	}

	return false
}

// SetSnapshots gets a reference to the given []VolumeSnapshot and assigns it to the Snapshots field.
func (o *VolumeSnapshotList) SetSnapshots(v []VolumeSnapshot) {
	o.Snapshots = &v
}

func (o VolumeSnapshotList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Snapshots != nil {
		toSerialize["snapshots"] = o.Snapshots
	}
	return json.Marshal(toSerialize)
}

type NullableVolumeSnapshotList struct {
	value *VolumeSnapshotList
	isSet bool
}

func (v NullableVolumeSnapshotList) Get() *VolumeSnapshotList {
	return v.value
}

func (v *NullableVolumeSnapshotList) Set(val *VolumeSnapshotList) {
	v.value = val
	v.isSet = true
}

func (v NullableVolumeSnapshotList) IsSet() bool {
	return v.isSet
}

func (v *NullableVolumeSnapshotList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVolumeSnapshotList(val *VolumeSnapshotList) *NullableVolumeSnapshotList {
	return &NullableVolumeSnapshotList{value: val, isSet: true}
}

func (v NullableVolumeSnapshotList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVolumeSnapshotList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


