/*
Equinix Internet Access API

Equinix Internet Access provides direct access to the Internet with scalable bandwidth options in IBX data centers.

Contact: api-support@equinix.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eiav2

import (
	"encoding/json"
)

// checks if the ServiceOrderRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceOrderRequest{}

// ServiceOrderRequest struct for ServiceOrderRequest
type ServiceOrderRequest struct {
	Tags                 []string               `json:"tags,omitempty"`
	Contacts             []ServiceOrderContact  `json:"contacts,omitempty"`
	PurchaseOrder        *ServicePurchaseOrder  `json:"purchaseOrder,omitempty"`
	ReferenceNumber      *string                `json:"referenceNumber,omitempty"`
	Signature            *OrderSignatureRequest `json:"signature,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ServiceOrderRequest ServiceOrderRequest

// NewServiceOrderRequest instantiates a new ServiceOrderRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceOrderRequest() *ServiceOrderRequest {
	this := ServiceOrderRequest{}
	return &this
}

// NewServiceOrderRequestWithDefaults instantiates a new ServiceOrderRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceOrderRequestWithDefaults() *ServiceOrderRequest {
	this := ServiceOrderRequest{}
	return &this
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *ServiceOrderRequest) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceOrderRequest) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *ServiceOrderRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *ServiceOrderRequest) SetTags(v []string) {
	o.Tags = v
}

// GetContacts returns the Contacts field value if set, zero value otherwise.
func (o *ServiceOrderRequest) GetContacts() []ServiceOrderContact {
	if o == nil || IsNil(o.Contacts) {
		var ret []ServiceOrderContact
		return ret
	}
	return o.Contacts
}

// GetContactsOk returns a tuple with the Contacts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceOrderRequest) GetContactsOk() ([]ServiceOrderContact, bool) {
	if o == nil || IsNil(o.Contacts) {
		return nil, false
	}
	return o.Contacts, true
}

// HasContacts returns a boolean if a field has been set.
func (o *ServiceOrderRequest) HasContacts() bool {
	if o != nil && !IsNil(o.Contacts) {
		return true
	}

	return false
}

// SetContacts gets a reference to the given []ServiceOrderContact and assigns it to the Contacts field.
func (o *ServiceOrderRequest) SetContacts(v []ServiceOrderContact) {
	o.Contacts = v
}

// GetPurchaseOrder returns the PurchaseOrder field value if set, zero value otherwise.
func (o *ServiceOrderRequest) GetPurchaseOrder() ServicePurchaseOrder {
	if o == nil || IsNil(o.PurchaseOrder) {
		var ret ServicePurchaseOrder
		return ret
	}
	return *o.PurchaseOrder
}

// GetPurchaseOrderOk returns a tuple with the PurchaseOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceOrderRequest) GetPurchaseOrderOk() (*ServicePurchaseOrder, bool) {
	if o == nil || IsNil(o.PurchaseOrder) {
		return nil, false
	}
	return o.PurchaseOrder, true
}

// HasPurchaseOrder returns a boolean if a field has been set.
func (o *ServiceOrderRequest) HasPurchaseOrder() bool {
	if o != nil && !IsNil(o.PurchaseOrder) {
		return true
	}

	return false
}

// SetPurchaseOrder gets a reference to the given ServicePurchaseOrder and assigns it to the PurchaseOrder field.
func (o *ServiceOrderRequest) SetPurchaseOrder(v ServicePurchaseOrder) {
	o.PurchaseOrder = &v
}

// GetReferenceNumber returns the ReferenceNumber field value if set, zero value otherwise.
func (o *ServiceOrderRequest) GetReferenceNumber() string {
	if o == nil || IsNil(o.ReferenceNumber) {
		var ret string
		return ret
	}
	return *o.ReferenceNumber
}

// GetReferenceNumberOk returns a tuple with the ReferenceNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceOrderRequest) GetReferenceNumberOk() (*string, bool) {
	if o == nil || IsNil(o.ReferenceNumber) {
		return nil, false
	}
	return o.ReferenceNumber, true
}

// HasReferenceNumber returns a boolean if a field has been set.
func (o *ServiceOrderRequest) HasReferenceNumber() bool {
	if o != nil && !IsNil(o.ReferenceNumber) {
		return true
	}

	return false
}

// SetReferenceNumber gets a reference to the given string and assigns it to the ReferenceNumber field.
func (o *ServiceOrderRequest) SetReferenceNumber(v string) {
	o.ReferenceNumber = &v
}

// GetSignature returns the Signature field value if set, zero value otherwise.
func (o *ServiceOrderRequest) GetSignature() OrderSignatureRequest {
	if o == nil || IsNil(o.Signature) {
		var ret OrderSignatureRequest
		return ret
	}
	return *o.Signature
}

// GetSignatureOk returns a tuple with the Signature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceOrderRequest) GetSignatureOk() (*OrderSignatureRequest, bool) {
	if o == nil || IsNil(o.Signature) {
		return nil, false
	}
	return o.Signature, true
}

// HasSignature returns a boolean if a field has been set.
func (o *ServiceOrderRequest) HasSignature() bool {
	if o != nil && !IsNil(o.Signature) {
		return true
	}

	return false
}

// SetSignature gets a reference to the given OrderSignatureRequest and assigns it to the Signature field.
func (o *ServiceOrderRequest) SetSignature(v OrderSignatureRequest) {
	o.Signature = &v
}

func (o ServiceOrderRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceOrderRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.Contacts) {
		toSerialize["contacts"] = o.Contacts
	}
	if !IsNil(o.PurchaseOrder) {
		toSerialize["purchaseOrder"] = o.PurchaseOrder
	}
	if !IsNil(o.ReferenceNumber) {
		toSerialize["referenceNumber"] = o.ReferenceNumber
	}
	if !IsNil(o.Signature) {
		toSerialize["signature"] = o.Signature
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ServiceOrderRequest) UnmarshalJSON(data []byte) (err error) {
	varServiceOrderRequest := _ServiceOrderRequest{}

	err = json.Unmarshal(data, &varServiceOrderRequest)

	if err != nil {
		return err
	}

	*o = ServiceOrderRequest(varServiceOrderRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "tags")
		delete(additionalProperties, "contacts")
		delete(additionalProperties, "purchaseOrder")
		delete(additionalProperties, "referenceNumber")
		delete(additionalProperties, "signature")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableServiceOrderRequest struct {
	value *ServiceOrderRequest
	isSet bool
}

func (v NullableServiceOrderRequest) Get() *ServiceOrderRequest {
	return v.value
}

func (v *NullableServiceOrderRequest) Set(val *ServiceOrderRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceOrderRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceOrderRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceOrderRequest(val *ServiceOrderRequest) *NullableServiceOrderRequest {
	return &NullableServiceOrderRequest{value: val, isSet: true}
}

func (v NullableServiceOrderRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceOrderRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
