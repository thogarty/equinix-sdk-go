/*
Equinix Fabric API v4

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fabricv4

import (
	"encoding/json"
)

// checks if the CloudRouterOrFilter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CloudRouterOrFilter{}

// CloudRouterOrFilter struct for CloudRouterOrFilter
type CloudRouterOrFilter struct {
	Or                   []CloudRouterSimpleExpression `json:"or,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CloudRouterOrFilter CloudRouterOrFilter

// NewCloudRouterOrFilter instantiates a new CloudRouterOrFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudRouterOrFilter() *CloudRouterOrFilter {
	this := CloudRouterOrFilter{}
	return &this
}

// NewCloudRouterOrFilterWithDefaults instantiates a new CloudRouterOrFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudRouterOrFilterWithDefaults() *CloudRouterOrFilter {
	this := CloudRouterOrFilter{}
	return &this
}

// GetOr returns the Or field value if set, zero value otherwise.
func (o *CloudRouterOrFilter) GetOr() []CloudRouterSimpleExpression {
	if o == nil || IsNil(o.Or) {
		var ret []CloudRouterSimpleExpression
		return ret
	}
	return o.Or
}

// GetOrOk returns a tuple with the Or field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudRouterOrFilter) GetOrOk() ([]CloudRouterSimpleExpression, bool) {
	if o == nil || IsNil(o.Or) {
		return nil, false
	}
	return o.Or, true
}

// HasOr returns a boolean if a field has been set.
func (o *CloudRouterOrFilter) HasOr() bool {
	if o != nil && !IsNil(o.Or) {
		return true
	}

	return false
}

// SetOr gets a reference to the given []CloudRouterSimpleExpression and assigns it to the Or field.
func (o *CloudRouterOrFilter) SetOr(v []CloudRouterSimpleExpression) {
	o.Or = v
}

func (o CloudRouterOrFilter) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CloudRouterOrFilter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Or) {
		toSerialize["or"] = o.Or
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CloudRouterOrFilter) UnmarshalJSON(data []byte) (err error) {
	varCloudRouterOrFilter := _CloudRouterOrFilter{}

	err = json.Unmarshal(data, &varCloudRouterOrFilter)

	if err != nil {
		return err
	}

	*o = CloudRouterOrFilter(varCloudRouterOrFilter)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "or")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCloudRouterOrFilter struct {
	value *CloudRouterOrFilter
	isSet bool
}

func (v NullableCloudRouterOrFilter) Get() *CloudRouterOrFilter {
	return v.value
}

func (v *NullableCloudRouterOrFilter) Set(val *CloudRouterOrFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudRouterOrFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudRouterOrFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudRouterOrFilter(val *CloudRouterOrFilter) *NullableCloudRouterOrFilter {
	return &NullableCloudRouterOrFilter{value: val, isSet: true}
}

func (v NullableCloudRouterOrFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudRouterOrFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
