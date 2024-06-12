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

// checks if the ServiceProfileAccessPointCOLO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceProfileAccessPointCOLO{}

// ServiceProfileAccessPointCOLO Colo Access Point
type ServiceProfileAccessPointCOLO struct {
	Type                    ServiceProfileAccessPointCOLOType `json:"type"`
	Uuid                    string                            `json:"uuid"`
	Location                *SimplifiedLocation               `json:"location,omitempty"`
	SellerRegion            *string                           `json:"sellerRegion,omitempty"`
	SellerRegionDescription *string                           `json:"sellerRegionDescription,omitempty"`
	CrossConnectId          *string                           `json:"crossConnectId,omitempty"`
	AdditionalProperties    map[string]interface{}
}

type _ServiceProfileAccessPointCOLO ServiceProfileAccessPointCOLO

// NewServiceProfileAccessPointCOLO instantiates a new ServiceProfileAccessPointCOLO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceProfileAccessPointCOLO(type_ ServiceProfileAccessPointCOLOType, uuid string) *ServiceProfileAccessPointCOLO {
	this := ServiceProfileAccessPointCOLO{}
	this.Type = type_
	this.Uuid = uuid
	return &this
}

// NewServiceProfileAccessPointCOLOWithDefaults instantiates a new ServiceProfileAccessPointCOLO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceProfileAccessPointCOLOWithDefaults() *ServiceProfileAccessPointCOLO {
	this := ServiceProfileAccessPointCOLO{}
	return &this
}

// GetType returns the Type field value
func (o *ServiceProfileAccessPointCOLO) GetType() ServiceProfileAccessPointCOLOType {
	if o == nil {
		var ret ServiceProfileAccessPointCOLOType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ServiceProfileAccessPointCOLO) GetTypeOk() (*ServiceProfileAccessPointCOLOType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ServiceProfileAccessPointCOLO) SetType(v ServiceProfileAccessPointCOLOType) {
	o.Type = v
}

// GetUuid returns the Uuid field value
func (o *ServiceProfileAccessPointCOLO) GetUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value
// and a boolean to check if the value has been set.
func (o *ServiceProfileAccessPointCOLO) GetUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uuid, true
}

// SetUuid sets field value
func (o *ServiceProfileAccessPointCOLO) SetUuid(v string) {
	o.Uuid = v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *ServiceProfileAccessPointCOLO) GetLocation() SimplifiedLocation {
	if o == nil || IsNil(o.Location) {
		var ret SimplifiedLocation
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceProfileAccessPointCOLO) GetLocationOk() (*SimplifiedLocation, bool) {
	if o == nil || IsNil(o.Location) {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *ServiceProfileAccessPointCOLO) HasLocation() bool {
	if o != nil && !IsNil(o.Location) {
		return true
	}

	return false
}

// SetLocation gets a reference to the given SimplifiedLocation and assigns it to the Location field.
func (o *ServiceProfileAccessPointCOLO) SetLocation(v SimplifiedLocation) {
	o.Location = &v
}

// GetSellerRegion returns the SellerRegion field value if set, zero value otherwise.
func (o *ServiceProfileAccessPointCOLO) GetSellerRegion() string {
	if o == nil || IsNil(o.SellerRegion) {
		var ret string
		return ret
	}
	return *o.SellerRegion
}

// GetSellerRegionOk returns a tuple with the SellerRegion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceProfileAccessPointCOLO) GetSellerRegionOk() (*string, bool) {
	if o == nil || IsNil(o.SellerRegion) {
		return nil, false
	}
	return o.SellerRegion, true
}

// HasSellerRegion returns a boolean if a field has been set.
func (o *ServiceProfileAccessPointCOLO) HasSellerRegion() bool {
	if o != nil && !IsNil(o.SellerRegion) {
		return true
	}

	return false
}

// SetSellerRegion gets a reference to the given string and assigns it to the SellerRegion field.
func (o *ServiceProfileAccessPointCOLO) SetSellerRegion(v string) {
	o.SellerRegion = &v
}

// GetSellerRegionDescription returns the SellerRegionDescription field value if set, zero value otherwise.
func (o *ServiceProfileAccessPointCOLO) GetSellerRegionDescription() string {
	if o == nil || IsNil(o.SellerRegionDescription) {
		var ret string
		return ret
	}
	return *o.SellerRegionDescription
}

// GetSellerRegionDescriptionOk returns a tuple with the SellerRegionDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceProfileAccessPointCOLO) GetSellerRegionDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.SellerRegionDescription) {
		return nil, false
	}
	return o.SellerRegionDescription, true
}

// HasSellerRegionDescription returns a boolean if a field has been set.
func (o *ServiceProfileAccessPointCOLO) HasSellerRegionDescription() bool {
	if o != nil && !IsNil(o.SellerRegionDescription) {
		return true
	}

	return false
}

// SetSellerRegionDescription gets a reference to the given string and assigns it to the SellerRegionDescription field.
func (o *ServiceProfileAccessPointCOLO) SetSellerRegionDescription(v string) {
	o.SellerRegionDescription = &v
}

// GetCrossConnectId returns the CrossConnectId field value if set, zero value otherwise.
func (o *ServiceProfileAccessPointCOLO) GetCrossConnectId() string {
	if o == nil || IsNil(o.CrossConnectId) {
		var ret string
		return ret
	}
	return *o.CrossConnectId
}

// GetCrossConnectIdOk returns a tuple with the CrossConnectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceProfileAccessPointCOLO) GetCrossConnectIdOk() (*string, bool) {
	if o == nil || IsNil(o.CrossConnectId) {
		return nil, false
	}
	return o.CrossConnectId, true
}

// HasCrossConnectId returns a boolean if a field has been set.
func (o *ServiceProfileAccessPointCOLO) HasCrossConnectId() bool {
	if o != nil && !IsNil(o.CrossConnectId) {
		return true
	}

	return false
}

// SetCrossConnectId gets a reference to the given string and assigns it to the CrossConnectId field.
func (o *ServiceProfileAccessPointCOLO) SetCrossConnectId(v string) {
	o.CrossConnectId = &v
}

func (o ServiceProfileAccessPointCOLO) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceProfileAccessPointCOLO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["uuid"] = o.Uuid
	if !IsNil(o.Location) {
		toSerialize["location"] = o.Location
	}
	if !IsNil(o.SellerRegion) {
		toSerialize["sellerRegion"] = o.SellerRegion
	}
	if !IsNil(o.SellerRegionDescription) {
		toSerialize["sellerRegionDescription"] = o.SellerRegionDescription
	}
	if !IsNil(o.CrossConnectId) {
		toSerialize["crossConnectId"] = o.CrossConnectId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ServiceProfileAccessPointCOLO) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"uuid",
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

	varServiceProfileAccessPointCOLO := _ServiceProfileAccessPointCOLO{}

	err = json.Unmarshal(data, &varServiceProfileAccessPointCOLO)

	if err != nil {
		return err
	}

	*o = ServiceProfileAccessPointCOLO(varServiceProfileAccessPointCOLO)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "uuid")
		delete(additionalProperties, "location")
		delete(additionalProperties, "sellerRegion")
		delete(additionalProperties, "sellerRegionDescription")
		delete(additionalProperties, "crossConnectId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableServiceProfileAccessPointCOLO struct {
	value *ServiceProfileAccessPointCOLO
	isSet bool
}

func (v NullableServiceProfileAccessPointCOLO) Get() *ServiceProfileAccessPointCOLO {
	return v.value
}

func (v *NullableServiceProfileAccessPointCOLO) Set(val *ServiceProfileAccessPointCOLO) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceProfileAccessPointCOLO) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceProfileAccessPointCOLO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceProfileAccessPointCOLO(val *ServiceProfileAccessPointCOLO) *NullableServiceProfileAccessPointCOLO {
	return &NullableServiceProfileAccessPointCOLO{value: val, isSet: true}
}

func (v NullableServiceProfileAccessPointCOLO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceProfileAccessPointCOLO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
