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
	"time"
)

// VolumeAttachment struct for VolumeAttachment
type VolumeAttachment struct {
	Id *string `json:"id,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	Volume *Href `json:"volume,omitempty"`
	Device *Href `json:"device,omitempty"`
	Href *string `json:"href,omitempty"`
}

// NewVolumeAttachment instantiates a new VolumeAttachment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVolumeAttachment() *VolumeAttachment {
	this := VolumeAttachment{}
	return &this
}

// NewVolumeAttachmentWithDefaults instantiates a new VolumeAttachment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVolumeAttachmentWithDefaults() *VolumeAttachment {
	this := VolumeAttachment{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *VolumeAttachment) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumeAttachment) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *VolumeAttachment) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *VolumeAttachment) SetId(v string) {
	o.Id = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *VolumeAttachment) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumeAttachment) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *VolumeAttachment) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *VolumeAttachment) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetVolume returns the Volume field value if set, zero value otherwise.
func (o *VolumeAttachment) GetVolume() Href {
	if o == nil || o.Volume == nil {
		var ret Href
		return ret
	}
	return *o.Volume
}

// GetVolumeOk returns a tuple with the Volume field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumeAttachment) GetVolumeOk() (*Href, bool) {
	if o == nil || o.Volume == nil {
		return nil, false
	}
	return o.Volume, true
}

// HasVolume returns a boolean if a field has been set.
func (o *VolumeAttachment) HasVolume() bool {
	if o != nil && o.Volume != nil {
		return true
	}

	return false
}

// SetVolume gets a reference to the given Href and assigns it to the Volume field.
func (o *VolumeAttachment) SetVolume(v Href) {
	o.Volume = &v
}

// GetDevice returns the Device field value if set, zero value otherwise.
func (o *VolumeAttachment) GetDevice() Href {
	if o == nil || o.Device == nil {
		var ret Href
		return ret
	}
	return *o.Device
}

// GetDeviceOk returns a tuple with the Device field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumeAttachment) GetDeviceOk() (*Href, bool) {
	if o == nil || o.Device == nil {
		return nil, false
	}
	return o.Device, true
}

// HasDevice returns a boolean if a field has been set.
func (o *VolumeAttachment) HasDevice() bool {
	if o != nil && o.Device != nil {
		return true
	}

	return false
}

// SetDevice gets a reference to the given Href and assigns it to the Device field.
func (o *VolumeAttachment) SetDevice(v Href) {
	o.Device = &v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *VolumeAttachment) GetHref() string {
	if o == nil || o.Href == nil {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumeAttachment) GetHrefOk() (*string, bool) {
	if o == nil || o.Href == nil {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *VolumeAttachment) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *VolumeAttachment) SetHref(v string) {
	o.Href = &v
}

func (o VolumeAttachment) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.Volume != nil {
		toSerialize["volume"] = o.Volume
	}
	if o.Device != nil {
		toSerialize["device"] = o.Device
	}
	if o.Href != nil {
		toSerialize["href"] = o.Href
	}
	return json.Marshal(toSerialize)
}

type NullableVolumeAttachment struct {
	value *VolumeAttachment
	isSet bool
}

func (v NullableVolumeAttachment) Get() *VolumeAttachment {
	return v.value
}

func (v *NullableVolumeAttachment) Set(val *VolumeAttachment) {
	v.value = val
	v.isSet = true
}

func (v NullableVolumeAttachment) IsSet() bool {
	return v.isSet
}

func (v *NullableVolumeAttachment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVolumeAttachment(val *VolumeAttachment) *NullableVolumeAttachment {
	return &NullableVolumeAttachment{value: val, isSet: true}
}

func (v NullableVolumeAttachment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVolumeAttachment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


