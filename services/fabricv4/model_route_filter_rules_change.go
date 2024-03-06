/*
Equinix Fabric API v4

Equinix Fabric is an advanced software-defined interconnection solution that enables you to directly, securely and dynamically connect to distributed infrastructure and digital ecosystems on platform Equinix via a single port, Customers can use Fabric to connect to: </br> 1. Cloud Service Providers - Clouds, network and other service providers.  </br> 2. Enterprises - Other Equinix customers, vendors and partners.  </br> 3. Myself - Another customer instance deployed at Equinix. </br>

API version: 4.12
Contact: api-support@equinix.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fabricv4

import (
	"encoding/json"
	"fmt"
)

// checks if the RouteFilterRulesChange type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RouteFilterRulesChange{}

// RouteFilterRulesChange Current state of latest Route Filter Rule change
type RouteFilterRulesChange struct {
	// Uniquely identifies a change
	Uuid string                     `json:"uuid"`
	Type RouteFilterRulesChangeType `json:"type"`
	// Route Filter Change URI
	Href                 *string `json:"href,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RouteFilterRulesChange RouteFilterRulesChange

// NewRouteFilterRulesChange instantiates a new RouteFilterRulesChange object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRouteFilterRulesChange(uuid string, type_ RouteFilterRulesChangeType) *RouteFilterRulesChange {
	this := RouteFilterRulesChange{}
	this.Uuid = uuid
	this.Type = type_
	return &this
}

// NewRouteFilterRulesChangeWithDefaults instantiates a new RouteFilterRulesChange object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRouteFilterRulesChangeWithDefaults() *RouteFilterRulesChange {
	this := RouteFilterRulesChange{}
	return &this
}

// GetUuid returns the Uuid field value
func (o *RouteFilterRulesChange) GetUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value
// and a boolean to check if the value has been set.
func (o *RouteFilterRulesChange) GetUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uuid, true
}

// SetUuid sets field value
func (o *RouteFilterRulesChange) SetUuid(v string) {
	o.Uuid = v
}

// GetType returns the Type field value
func (o *RouteFilterRulesChange) GetType() RouteFilterRulesChangeType {
	if o == nil {
		var ret RouteFilterRulesChangeType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *RouteFilterRulesChange) GetTypeOk() (*RouteFilterRulesChangeType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *RouteFilterRulesChange) SetType(v RouteFilterRulesChangeType) {
	o.Type = v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *RouteFilterRulesChange) GetHref() string {
	if o == nil || IsNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RouteFilterRulesChange) GetHrefOk() (*string, bool) {
	if o == nil || IsNil(o.Href) {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *RouteFilterRulesChange) HasHref() bool {
	if o != nil && !IsNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *RouteFilterRulesChange) SetHref(v string) {
	o.Href = &v
}

func (o RouteFilterRulesChange) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RouteFilterRulesChange) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["uuid"] = o.Uuid
	toSerialize["type"] = o.Type
	if !IsNil(o.Href) {
		toSerialize["href"] = o.Href
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RouteFilterRulesChange) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"uuid",
		"type",
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

	varRouteFilterRulesChange := _RouteFilterRulesChange{}

	err = json.Unmarshal(data, &varRouteFilterRulesChange)

	if err != nil {
		return err
	}

	*o = RouteFilterRulesChange(varRouteFilterRulesChange)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "uuid")
		delete(additionalProperties, "type")
		delete(additionalProperties, "href")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRouteFilterRulesChange struct {
	value *RouteFilterRulesChange
	isSet bool
}

func (v NullableRouteFilterRulesChange) Get() *RouteFilterRulesChange {
	return v.value
}

func (v *NullableRouteFilterRulesChange) Set(val *RouteFilterRulesChange) {
	v.value = val
	v.isSet = true
}

func (v NullableRouteFilterRulesChange) IsSet() bool {
	return v.isSet
}

func (v *NullableRouteFilterRulesChange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRouteFilterRulesChange(val *RouteFilterRulesChange) *NullableRouteFilterRulesChange {
	return &NullableRouteFilterRulesChange{value: val, isSet: true}
}

func (v NullableRouteFilterRulesChange) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRouteFilterRulesChange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}