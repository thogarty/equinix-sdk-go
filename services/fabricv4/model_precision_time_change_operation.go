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

// checks if the PrecisionTimeChangeOperation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PrecisionTimeChangeOperation{}

// PrecisionTimeChangeOperation Fabric Precision Timing change operation data
type PrecisionTimeChangeOperation struct {
	Op   PrecisionTimeChangeOperationOp   `json:"op"`
	Path PrecisionTimeChangeOperationPath `json:"path"`
	// new value for updated parameter
	Value                interface{} `json:"value"`
	AdditionalProperties map[string]interface{}
}

type _PrecisionTimeChangeOperation PrecisionTimeChangeOperation

// NewPrecisionTimeChangeOperation instantiates a new PrecisionTimeChangeOperation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrecisionTimeChangeOperation(op PrecisionTimeChangeOperationOp, path PrecisionTimeChangeOperationPath, value interface{}) *PrecisionTimeChangeOperation {
	this := PrecisionTimeChangeOperation{}
	this.Op = op
	this.Path = path
	this.Value = value
	return &this
}

// NewPrecisionTimeChangeOperationWithDefaults instantiates a new PrecisionTimeChangeOperation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrecisionTimeChangeOperationWithDefaults() *PrecisionTimeChangeOperation {
	this := PrecisionTimeChangeOperation{}
	return &this
}

// GetOp returns the Op field value
func (o *PrecisionTimeChangeOperation) GetOp() PrecisionTimeChangeOperationOp {
	if o == nil {
		var ret PrecisionTimeChangeOperationOp
		return ret
	}

	return o.Op
}

// GetOpOk returns a tuple with the Op field value
// and a boolean to check if the value has been set.
func (o *PrecisionTimeChangeOperation) GetOpOk() (*PrecisionTimeChangeOperationOp, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Op, true
}

// SetOp sets field value
func (o *PrecisionTimeChangeOperation) SetOp(v PrecisionTimeChangeOperationOp) {
	o.Op = v
}

// GetPath returns the Path field value
func (o *PrecisionTimeChangeOperation) GetPath() PrecisionTimeChangeOperationPath {
	if o == nil {
		var ret PrecisionTimeChangeOperationPath
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *PrecisionTimeChangeOperation) GetPathOk() (*PrecisionTimeChangeOperationPath, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *PrecisionTimeChangeOperation) SetPath(v PrecisionTimeChangeOperationPath) {
	o.Path = v
}

// GetValue returns the Value field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *PrecisionTimeChangeOperation) GetValue() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PrecisionTimeChangeOperation) GetValueOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *PrecisionTimeChangeOperation) SetValue(v interface{}) {
	o.Value = v
}

func (o PrecisionTimeChangeOperation) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PrecisionTimeChangeOperation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["op"] = o.Op
	toSerialize["path"] = o.Path
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PrecisionTimeChangeOperation) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"op",
		"path",
		"value",
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

	varPrecisionTimeChangeOperation := _PrecisionTimeChangeOperation{}

	err = json.Unmarshal(data, &varPrecisionTimeChangeOperation)

	if err != nil {
		return err
	}

	*o = PrecisionTimeChangeOperation(varPrecisionTimeChangeOperation)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "op")
		delete(additionalProperties, "path")
		delete(additionalProperties, "value")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePrecisionTimeChangeOperation struct {
	value *PrecisionTimeChangeOperation
	isSet bool
}

func (v NullablePrecisionTimeChangeOperation) Get() *PrecisionTimeChangeOperation {
	return v.value
}

func (v *NullablePrecisionTimeChangeOperation) Set(val *PrecisionTimeChangeOperation) {
	v.value = val
	v.isSet = true
}

func (v NullablePrecisionTimeChangeOperation) IsSet() bool {
	return v.isSet
}

func (v *NullablePrecisionTimeChangeOperation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrecisionTimeChangeOperation(val *PrecisionTimeChangeOperation) *NullablePrecisionTimeChangeOperation {
	return &NullablePrecisionTimeChangeOperation{value: val, isSet: true}
}

func (v NullablePrecisionTimeChangeOperation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrecisionTimeChangeOperation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}