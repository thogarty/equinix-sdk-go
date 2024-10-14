/*
Equinix Fabric API v4

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fabricv4

import (
	"encoding/json"
	"fmt"
)

// checks if the PrecisionTimePackageRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PrecisionTimePackageRequest{}

// PrecisionTimePackageRequest Precision Time Service Level Request
type PrecisionTimePackageRequest struct {
	Code                 PrecisionTimePackageRequestCode `json:"code"`
	AdditionalProperties map[string]interface{}
}

type _PrecisionTimePackageRequest PrecisionTimePackageRequest

// NewPrecisionTimePackageRequest instantiates a new PrecisionTimePackageRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrecisionTimePackageRequest(code PrecisionTimePackageRequestCode) *PrecisionTimePackageRequest {
	this := PrecisionTimePackageRequest{}
	this.Code = code
	return &this
}

// NewPrecisionTimePackageRequestWithDefaults instantiates a new PrecisionTimePackageRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrecisionTimePackageRequestWithDefaults() *PrecisionTimePackageRequest {
	this := PrecisionTimePackageRequest{}
	return &this
}

// GetCode returns the Code field value
func (o *PrecisionTimePackageRequest) GetCode() PrecisionTimePackageRequestCode {
	if o == nil {
		var ret PrecisionTimePackageRequestCode
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *PrecisionTimePackageRequest) GetCodeOk() (*PrecisionTimePackageRequestCode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *PrecisionTimePackageRequest) SetCode(v PrecisionTimePackageRequestCode) {
	o.Code = v
}

func (o PrecisionTimePackageRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PrecisionTimePackageRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PrecisionTimePackageRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"code",
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

	varPrecisionTimePackageRequest := _PrecisionTimePackageRequest{}

	err = json.Unmarshal(data, &varPrecisionTimePackageRequest)

	if err != nil {
		return err
	}

	*o = PrecisionTimePackageRequest(varPrecisionTimePackageRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "code")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePrecisionTimePackageRequest struct {
	value *PrecisionTimePackageRequest
	isSet bool
}

func (v NullablePrecisionTimePackageRequest) Get() *PrecisionTimePackageRequest {
	return v.value
}

func (v *NullablePrecisionTimePackageRequest) Set(val *PrecisionTimePackageRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePrecisionTimePackageRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePrecisionTimePackageRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrecisionTimePackageRequest(val *PrecisionTimePackageRequest) *NullablePrecisionTimePackageRequest {
	return &NullablePrecisionTimePackageRequest{value: val, isSet: true}
}

func (v NullablePrecisionTimePackageRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrecisionTimePackageRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
