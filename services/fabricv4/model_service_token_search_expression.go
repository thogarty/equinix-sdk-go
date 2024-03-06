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
)

// checks if the ServiceTokenSearchExpression type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceTokenSearchExpression{}

// ServiceTokenSearchExpression struct for ServiceTokenSearchExpression
type ServiceTokenSearchExpression struct {
	And                  []ServiceTokenSearchExpression        `json:"and,omitempty"`
	Property             *ServiceTokenSearchFieldName          `json:"property,omitempty"`
	Operator             *ServiceTokenSearchExpressionOperator `json:"operator,omitempty"`
	Values               []string                              `json:"values,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ServiceTokenSearchExpression ServiceTokenSearchExpression

// NewServiceTokenSearchExpression instantiates a new ServiceTokenSearchExpression object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceTokenSearchExpression() *ServiceTokenSearchExpression {
	this := ServiceTokenSearchExpression{}
	return &this
}

// NewServiceTokenSearchExpressionWithDefaults instantiates a new ServiceTokenSearchExpression object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceTokenSearchExpressionWithDefaults() *ServiceTokenSearchExpression {
	this := ServiceTokenSearchExpression{}
	return &this
}

// GetAnd returns the And field value if set, zero value otherwise.
func (o *ServiceTokenSearchExpression) GetAnd() []ServiceTokenSearchExpression {
	if o == nil || IsNil(o.And) {
		var ret []ServiceTokenSearchExpression
		return ret
	}
	return o.And
}

// GetAndOk returns a tuple with the And field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceTokenSearchExpression) GetAndOk() ([]ServiceTokenSearchExpression, bool) {
	if o == nil || IsNil(o.And) {
		return nil, false
	}
	return o.And, true
}

// HasAnd returns a boolean if a field has been set.
func (o *ServiceTokenSearchExpression) HasAnd() bool {
	if o != nil && !IsNil(o.And) {
		return true
	}

	return false
}

// SetAnd gets a reference to the given []ServiceTokenSearchExpression and assigns it to the And field.
func (o *ServiceTokenSearchExpression) SetAnd(v []ServiceTokenSearchExpression) {
	o.And = v
}

// GetProperty returns the Property field value if set, zero value otherwise.
func (o *ServiceTokenSearchExpression) GetProperty() ServiceTokenSearchFieldName {
	if o == nil || IsNil(o.Property) {
		var ret ServiceTokenSearchFieldName
		return ret
	}
	return *o.Property
}

// GetPropertyOk returns a tuple with the Property field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceTokenSearchExpression) GetPropertyOk() (*ServiceTokenSearchFieldName, bool) {
	if o == nil || IsNil(o.Property) {
		return nil, false
	}
	return o.Property, true
}

// HasProperty returns a boolean if a field has been set.
func (o *ServiceTokenSearchExpression) HasProperty() bool {
	if o != nil && !IsNil(o.Property) {
		return true
	}

	return false
}

// SetProperty gets a reference to the given ServiceTokenSearchFieldName and assigns it to the Property field.
func (o *ServiceTokenSearchExpression) SetProperty(v ServiceTokenSearchFieldName) {
	o.Property = &v
}

// GetOperator returns the Operator field value if set, zero value otherwise.
func (o *ServiceTokenSearchExpression) GetOperator() ServiceTokenSearchExpressionOperator {
	if o == nil || IsNil(o.Operator) {
		var ret ServiceTokenSearchExpressionOperator
		return ret
	}
	return *o.Operator
}

// GetOperatorOk returns a tuple with the Operator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceTokenSearchExpression) GetOperatorOk() (*ServiceTokenSearchExpressionOperator, bool) {
	if o == nil || IsNil(o.Operator) {
		return nil, false
	}
	return o.Operator, true
}

// HasOperator returns a boolean if a field has been set.
func (o *ServiceTokenSearchExpression) HasOperator() bool {
	if o != nil && !IsNil(o.Operator) {
		return true
	}

	return false
}

// SetOperator gets a reference to the given ServiceTokenSearchExpressionOperator and assigns it to the Operator field.
func (o *ServiceTokenSearchExpression) SetOperator(v ServiceTokenSearchExpressionOperator) {
	o.Operator = &v
}

// GetValues returns the Values field value if set, zero value otherwise.
func (o *ServiceTokenSearchExpression) GetValues() []string {
	if o == nil || IsNil(o.Values) {
		var ret []string
		return ret
	}
	return o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceTokenSearchExpression) GetValuesOk() ([]string, bool) {
	if o == nil || IsNil(o.Values) {
		return nil, false
	}
	return o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *ServiceTokenSearchExpression) HasValues() bool {
	if o != nil && !IsNil(o.Values) {
		return true
	}

	return false
}

// SetValues gets a reference to the given []string and assigns it to the Values field.
func (o *ServiceTokenSearchExpression) SetValues(v []string) {
	o.Values = v
}

func (o ServiceTokenSearchExpression) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceTokenSearchExpression) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.And) {
		toSerialize["and"] = o.And
	}
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

func (o *ServiceTokenSearchExpression) UnmarshalJSON(data []byte) (err error) {
	varServiceTokenSearchExpression := _ServiceTokenSearchExpression{}

	err = json.Unmarshal(data, &varServiceTokenSearchExpression)

	if err != nil {
		return err
	}

	*o = ServiceTokenSearchExpression(varServiceTokenSearchExpression)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "and")
		delete(additionalProperties, "property")
		delete(additionalProperties, "operator")
		delete(additionalProperties, "values")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableServiceTokenSearchExpression struct {
	value *ServiceTokenSearchExpression
	isSet bool
}

func (v NullableServiceTokenSearchExpression) Get() *ServiceTokenSearchExpression {
	return v.value
}

func (v *NullableServiceTokenSearchExpression) Set(val *ServiceTokenSearchExpression) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceTokenSearchExpression) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceTokenSearchExpression) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceTokenSearchExpression(val *ServiceTokenSearchExpression) *NullableServiceTokenSearchExpression {
	return &NullableServiceTokenSearchExpression{value: val, isSet: true}
}

func (v NullableServiceTokenSearchExpression) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceTokenSearchExpression) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}