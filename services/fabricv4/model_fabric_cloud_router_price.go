/*
Equinix Fabric API v4

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fabricv4

import (
	"encoding/json"
)

// checks if the FabricCloudRouterPrice type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FabricCloudRouterPrice{}

// FabricCloudRouterPrice Cloud Router  Product configuration
type FabricCloudRouterPrice struct {
	// Unique identifier assigned to the Cloud Router
	Uuid                 *string                    `json:"uuid,omitempty"`
	Location             *PriceLocation             `json:"location,omitempty"`
	Package              *FabricCloudRouterPackages `json:"package,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FabricCloudRouterPrice FabricCloudRouterPrice

// NewFabricCloudRouterPrice instantiates a new FabricCloudRouterPrice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFabricCloudRouterPrice() *FabricCloudRouterPrice {
	this := FabricCloudRouterPrice{}
	return &this
}

// NewFabricCloudRouterPriceWithDefaults instantiates a new FabricCloudRouterPrice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFabricCloudRouterPriceWithDefaults() *FabricCloudRouterPrice {
	this := FabricCloudRouterPrice{}
	return &this
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *FabricCloudRouterPrice) GetUuid() string {
	if o == nil || IsNil(o.Uuid) {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricCloudRouterPrice) GetUuidOk() (*string, bool) {
	if o == nil || IsNil(o.Uuid) {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *FabricCloudRouterPrice) HasUuid() bool {
	if o != nil && !IsNil(o.Uuid) {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *FabricCloudRouterPrice) SetUuid(v string) {
	o.Uuid = &v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *FabricCloudRouterPrice) GetLocation() PriceLocation {
	if o == nil || IsNil(o.Location) {
		var ret PriceLocation
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricCloudRouterPrice) GetLocationOk() (*PriceLocation, bool) {
	if o == nil || IsNil(o.Location) {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *FabricCloudRouterPrice) HasLocation() bool {
	if o != nil && !IsNil(o.Location) {
		return true
	}

	return false
}

// SetLocation gets a reference to the given PriceLocation and assigns it to the Location field.
func (o *FabricCloudRouterPrice) SetLocation(v PriceLocation) {
	o.Location = &v
}

// GetPackage returns the Package field value if set, zero value otherwise.
func (o *FabricCloudRouterPrice) GetPackage() FabricCloudRouterPackages {
	if o == nil || IsNil(o.Package) {
		var ret FabricCloudRouterPackages
		return ret
	}
	return *o.Package
}

// GetPackageOk returns a tuple with the Package field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricCloudRouterPrice) GetPackageOk() (*FabricCloudRouterPackages, bool) {
	if o == nil || IsNil(o.Package) {
		return nil, false
	}
	return o.Package, true
}

// HasPackage returns a boolean if a field has been set.
func (o *FabricCloudRouterPrice) HasPackage() bool {
	if o != nil && !IsNil(o.Package) {
		return true
	}

	return false
}

// SetPackage gets a reference to the given FabricCloudRouterPackages and assigns it to the Package field.
func (o *FabricCloudRouterPrice) SetPackage(v FabricCloudRouterPackages) {
	o.Package = &v
}

func (o FabricCloudRouterPrice) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FabricCloudRouterPrice) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Uuid) {
		toSerialize["uuid"] = o.Uuid
	}
	if !IsNil(o.Location) {
		toSerialize["location"] = o.Location
	}
	if !IsNil(o.Package) {
		toSerialize["package"] = o.Package
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FabricCloudRouterPrice) UnmarshalJSON(data []byte) (err error) {
	varFabricCloudRouterPrice := _FabricCloudRouterPrice{}

	err = json.Unmarshal(data, &varFabricCloudRouterPrice)

	if err != nil {
		return err
	}

	*o = FabricCloudRouterPrice(varFabricCloudRouterPrice)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "uuid")
		delete(additionalProperties, "location")
		delete(additionalProperties, "package")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFabricCloudRouterPrice struct {
	value *FabricCloudRouterPrice
	isSet bool
}

func (v NullableFabricCloudRouterPrice) Get() *FabricCloudRouterPrice {
	return v.value
}

func (v *NullableFabricCloudRouterPrice) Set(val *FabricCloudRouterPrice) {
	v.value = val
	v.isSet = true
}

func (v NullableFabricCloudRouterPrice) IsSet() bool {
	return v.isSet
}

func (v *NullableFabricCloudRouterPrice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFabricCloudRouterPrice(val *FabricCloudRouterPrice) *NullableFabricCloudRouterPrice {
	return &NullableFabricCloudRouterPrice{value: val, isSet: true}
}

func (v NullableFabricCloudRouterPrice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFabricCloudRouterPrice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
