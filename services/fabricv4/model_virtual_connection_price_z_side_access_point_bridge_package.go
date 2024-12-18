/*
Equinix Fabric API v4

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fabricv4

import (
	"encoding/json"
)

// checks if the VirtualConnectionPriceZSideAccessPointBridgePackage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VirtualConnectionPriceZSideAccessPointBridgePackage{}

// VirtualConnectionPriceZSideAccessPointBridgePackage struct for VirtualConnectionPriceZSideAccessPointBridgePackage
type VirtualConnectionPriceZSideAccessPointBridgePackage struct {
	Code                 *VirtualConnectionBridgePackageCode `json:"code,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VirtualConnectionPriceZSideAccessPointBridgePackage VirtualConnectionPriceZSideAccessPointBridgePackage

// NewVirtualConnectionPriceZSideAccessPointBridgePackage instantiates a new VirtualConnectionPriceZSideAccessPointBridgePackage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualConnectionPriceZSideAccessPointBridgePackage() *VirtualConnectionPriceZSideAccessPointBridgePackage {
	this := VirtualConnectionPriceZSideAccessPointBridgePackage{}
	return &this
}

// NewVirtualConnectionPriceZSideAccessPointBridgePackageWithDefaults instantiates a new VirtualConnectionPriceZSideAccessPointBridgePackage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualConnectionPriceZSideAccessPointBridgePackageWithDefaults() *VirtualConnectionPriceZSideAccessPointBridgePackage {
	this := VirtualConnectionPriceZSideAccessPointBridgePackage{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *VirtualConnectionPriceZSideAccessPointBridgePackage) GetCode() VirtualConnectionBridgePackageCode {
	if o == nil || IsNil(o.Code) {
		var ret VirtualConnectionBridgePackageCode
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualConnectionPriceZSideAccessPointBridgePackage) GetCodeOk() (*VirtualConnectionBridgePackageCode, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *VirtualConnectionPriceZSideAccessPointBridgePackage) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given VirtualConnectionBridgePackageCode and assigns it to the Code field.
func (o *VirtualConnectionPriceZSideAccessPointBridgePackage) SetCode(v VirtualConnectionBridgePackageCode) {
	o.Code = &v
}

func (o VirtualConnectionPriceZSideAccessPointBridgePackage) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VirtualConnectionPriceZSideAccessPointBridgePackage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *VirtualConnectionPriceZSideAccessPointBridgePackage) UnmarshalJSON(data []byte) (err error) {
	varVirtualConnectionPriceZSideAccessPointBridgePackage := _VirtualConnectionPriceZSideAccessPointBridgePackage{}

	err = json.Unmarshal(data, &varVirtualConnectionPriceZSideAccessPointBridgePackage)

	if err != nil {
		return err
	}

	*o = VirtualConnectionPriceZSideAccessPointBridgePackage(varVirtualConnectionPriceZSideAccessPointBridgePackage)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "code")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVirtualConnectionPriceZSideAccessPointBridgePackage struct {
	value *VirtualConnectionPriceZSideAccessPointBridgePackage
	isSet bool
}

func (v NullableVirtualConnectionPriceZSideAccessPointBridgePackage) Get() *VirtualConnectionPriceZSideAccessPointBridgePackage {
	return v.value
}

func (v *NullableVirtualConnectionPriceZSideAccessPointBridgePackage) Set(val *VirtualConnectionPriceZSideAccessPointBridgePackage) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualConnectionPriceZSideAccessPointBridgePackage) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualConnectionPriceZSideAccessPointBridgePackage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualConnectionPriceZSideAccessPointBridgePackage(val *VirtualConnectionPriceZSideAccessPointBridgePackage) *NullableVirtualConnectionPriceZSideAccessPointBridgePackage {
	return &NullableVirtualConnectionPriceZSideAccessPointBridgePackage{value: val, isSet: true}
}

func (v NullableVirtualConnectionPriceZSideAccessPointBridgePackage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualConnectionPriceZSideAccessPointBridgePackage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
