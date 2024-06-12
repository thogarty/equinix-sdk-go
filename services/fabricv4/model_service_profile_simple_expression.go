/*
Equinix Fabric API v4

Contact: api-support@equinix.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fabricv4

import (
	"encoding/json"
)

// checks if the ServiceProfileSimpleExpression type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceProfileSimpleExpression{}

// ServiceProfileSimpleExpression struct for ServiceProfileSimpleExpression
type ServiceProfileSimpleExpression struct {
	// Possible field names to use on filters:  * `/name` - Service Profile name  * `/uuid` - Service Profile uuid  * `/state` - Service Profile status  * `/metros/code` - Service Profile metro code  * `/visibility` - Service Profile package  * `/type` - Service Profile package  * `/project/projectId` - Service Profile project id
	Property *string `json:"property,omitempty"`
	// Possible operators to use on filters:  * `=` - equal
	Operator             *string  `json:"operator,omitempty"`
	Values               []string `json:"values,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ServiceProfileSimpleExpression ServiceProfileSimpleExpression

// NewServiceProfileSimpleExpression instantiates a new ServiceProfileSimpleExpression object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceProfileSimpleExpression() *ServiceProfileSimpleExpression {
	this := ServiceProfileSimpleExpression{}
	return &this
}

// NewServiceProfileSimpleExpressionWithDefaults instantiates a new ServiceProfileSimpleExpression object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceProfileSimpleExpressionWithDefaults() *ServiceProfileSimpleExpression {
	this := ServiceProfileSimpleExpression{}
	return &this
}

// GetProperty returns the Property field value if set, zero value otherwise.
func (o *ServiceProfileSimpleExpression) GetProperty() string {
	if o == nil || IsNil(o.Property) {
		var ret string
		return ret
	}
	return *o.Property
}

// GetPropertyOk returns a tuple with the Property field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceProfileSimpleExpression) GetPropertyOk() (*string, bool) {
	if o == nil || IsNil(o.Property) {
		return nil, false
	}
	return o.Property, true
}

// HasProperty returns a boolean if a field has been set.
func (o *ServiceProfileSimpleExpression) HasProperty() bool {
	if o != nil && !IsNil(o.Property) {
		return true
	}

	return false
}

// SetProperty gets a reference to the given string and assigns it to the Property field.
func (o *ServiceProfileSimpleExpression) SetProperty(v string) {
	o.Property = &v
}

// GetOperator returns the Operator field value if set, zero value otherwise.
func (o *ServiceProfileSimpleExpression) GetOperator() string {
	if o == nil || IsNil(o.Operator) {
		var ret string
		return ret
	}
	return *o.Operator
}

// GetOperatorOk returns a tuple with the Operator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceProfileSimpleExpression) GetOperatorOk() (*string, bool) {
	if o == nil || IsNil(o.Operator) {
		return nil, false
	}
	return o.Operator, true
}

// HasOperator returns a boolean if a field has been set.
func (o *ServiceProfileSimpleExpression) HasOperator() bool {
	if o != nil && !IsNil(o.Operator) {
		return true
	}

	return false
}

// SetOperator gets a reference to the given string and assigns it to the Operator field.
func (o *ServiceProfileSimpleExpression) SetOperator(v string) {
	o.Operator = &v
}

// GetValues returns the Values field value if set, zero value otherwise.
func (o *ServiceProfileSimpleExpression) GetValues() []string {
	if o == nil || IsNil(o.Values) {
		var ret []string
		return ret
	}
	return o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceProfileSimpleExpression) GetValuesOk() ([]string, bool) {
	if o == nil || IsNil(o.Values) {
		return nil, false
	}
	return o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *ServiceProfileSimpleExpression) HasValues() bool {
	if o != nil && !IsNil(o.Values) {
		return true
	}

	return false
}

// SetValues gets a reference to the given []string and assigns it to the Values field.
func (o *ServiceProfileSimpleExpression) SetValues(v []string) {
	o.Values = v
}

func (o ServiceProfileSimpleExpression) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceProfileSimpleExpression) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Property) {
		toSerialize["property"] = o.Property
	}
	if !IsNil(o.Operator) {
		toSerialize["operator"] = o.Operator
	}
	if !IsNil(o.Values) {
		toSerialize["values"] = o.Values
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ServiceProfileSimpleExpression) UnmarshalJSON(data []byte) (err error) {
	varServiceProfileSimpleExpression := _ServiceProfileSimpleExpression{}

	err = json.Unmarshal(data, &varServiceProfileSimpleExpression)

	if err != nil {
		return err
	}

	*o = ServiceProfileSimpleExpression(varServiceProfileSimpleExpression)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "property")
		delete(additionalProperties, "operator")
		delete(additionalProperties, "values")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableServiceProfileSimpleExpression struct {
	value *ServiceProfileSimpleExpression
	isSet bool
}

func (v NullableServiceProfileSimpleExpression) Get() *ServiceProfileSimpleExpression {
	return v.value
}

func (v *NullableServiceProfileSimpleExpression) Set(val *ServiceProfileSimpleExpression) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceProfileSimpleExpression) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceProfileSimpleExpression) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceProfileSimpleExpression(val *ServiceProfileSimpleExpression) *NullableServiceProfileSimpleExpression {
	return &NullableServiceProfileSimpleExpression{value: val, isSet: true}
}

func (v NullableServiceProfileSimpleExpression) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceProfileSimpleExpression) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
