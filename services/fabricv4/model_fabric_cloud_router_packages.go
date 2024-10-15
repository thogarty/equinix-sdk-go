/*
Equinix Fabric API v4

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fabricv4

import (
	"encoding/json"
)

// checks if the FabricCloudRouterPackages type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FabricCloudRouterPackages{}

// FabricCloudRouterPackages Cloud Router  package
type FabricCloudRouterPackages struct {
	Code                 *FabricCloudRouterCode `json:"code,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FabricCloudRouterPackages FabricCloudRouterPackages

// NewFabricCloudRouterPackages instantiates a new FabricCloudRouterPackages object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFabricCloudRouterPackages() *FabricCloudRouterPackages {
	this := FabricCloudRouterPackages{}
	return &this
}

// NewFabricCloudRouterPackagesWithDefaults instantiates a new FabricCloudRouterPackages object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFabricCloudRouterPackagesWithDefaults() *FabricCloudRouterPackages {
	this := FabricCloudRouterPackages{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *FabricCloudRouterPackages) GetCode() FabricCloudRouterCode {
	if o == nil || IsNil(o.Code) {
		var ret FabricCloudRouterCode
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricCloudRouterPackages) GetCodeOk() (*FabricCloudRouterCode, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *FabricCloudRouterPackages) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given FabricCloudRouterCode and assigns it to the Code field.
func (o *FabricCloudRouterPackages) SetCode(v FabricCloudRouterCode) {
	o.Code = &v
}

func (o FabricCloudRouterPackages) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FabricCloudRouterPackages) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FabricCloudRouterPackages) UnmarshalJSON(data []byte) (err error) {
	varFabricCloudRouterPackages := _FabricCloudRouterPackages{}

	err = json.Unmarshal(data, &varFabricCloudRouterPackages)

	if err != nil {
		return err
	}

	*o = FabricCloudRouterPackages(varFabricCloudRouterPackages)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "code")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFabricCloudRouterPackages struct {
	value *FabricCloudRouterPackages
	isSet bool
}

func (v NullableFabricCloudRouterPackages) Get() *FabricCloudRouterPackages {
	return v.value
}

func (v *NullableFabricCloudRouterPackages) Set(val *FabricCloudRouterPackages) {
	v.value = val
	v.isSet = true
}

func (v NullableFabricCloudRouterPackages) IsSet() bool {
	return v.isSet
}

func (v *NullableFabricCloudRouterPackages) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFabricCloudRouterPackages(val *FabricCloudRouterPackages) *NullableFabricCloudRouterPackages {
	return &NullableFabricCloudRouterPackages{value: val, isSet: true}
}

func (v NullableFabricCloudRouterPackages) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFabricCloudRouterPackages) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
