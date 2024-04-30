/*
Equinix Internet Access API

Equinix Internet Access provides direct access to the Internet with scalable bandwidth options in IBX data centers.

API version: 2.0
Contact: api-support@equinix.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eiav2

import (
	"encoding/json"
	"fmt"
)

// checks if the Order type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Order{}

// Order struct for Order
type Order struct {
	Href                 string              `json:"href"`
	Uuid                 string              `json:"uuid"`
	Type                 *OrderType          `json:"type,omitempty"`
	Contacts             []Contact           `json:"contacts,omitempty"`
	Draft                *bool               `json:"draft,omitempty"`
	Links                []Link              `json:"links,omitempty"`
	PurchaseOrder        *OrderPurchaseOrder `json:"purchaseOrder,omitempty"`
	ReferenceNumber      *string             `json:"referenceNumber,omitempty"`
	Signature            *OrderSignature     `json:"signature,omitempty"`
	State                *OrderState         `json:"state,omitempty"`
	ChangeLog            *OrderChangeLog     `json:"changeLog,omitempty"`
	Tags                 []string            `json:"tags,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Order Order

// NewOrder instantiates a new Order object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrder(href string, uuid string) *Order {
	this := Order{}
	this.Href = href
	this.Uuid = uuid
	var draft bool = false
	this.Draft = &draft
	return &this
}

// NewOrderWithDefaults instantiates a new Order object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderWithDefaults() *Order {
	this := Order{}
	var draft bool = false
	this.Draft = &draft
	return &this
}

// GetHref returns the Href field value
func (o *Order) GetHref() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Href
}

// GetHrefOk returns a tuple with the Href field value
// and a boolean to check if the value has been set.
func (o *Order) GetHrefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Href, true
}

// SetHref sets field value
func (o *Order) SetHref(v string) {
	o.Href = v
}

// GetUuid returns the Uuid field value
func (o *Order) GetUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value
// and a boolean to check if the value has been set.
func (o *Order) GetUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uuid, true
}

// SetUuid sets field value
func (o *Order) SetUuid(v string) {
	o.Uuid = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Order) GetType() OrderType {
	if o == nil || IsNil(o.Type) {
		var ret OrderType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Order) GetTypeOk() (*OrderType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Order) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given OrderType and assigns it to the Type field.
func (o *Order) SetType(v OrderType) {
	o.Type = &v
}

// GetContacts returns the Contacts field value if set, zero value otherwise.
func (o *Order) GetContacts() []Contact {
	if o == nil || IsNil(o.Contacts) {
		var ret []Contact
		return ret
	}
	return o.Contacts
}

// GetContactsOk returns a tuple with the Contacts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Order) GetContactsOk() ([]Contact, bool) {
	if o == nil || IsNil(o.Contacts) {
		return nil, false
	}
	return o.Contacts, true
}

// HasContacts returns a boolean if a field has been set.
func (o *Order) HasContacts() bool {
	if o != nil && !IsNil(o.Contacts) {
		return true
	}

	return false
}

// SetContacts gets a reference to the given []Contact and assigns it to the Contacts field.
func (o *Order) SetContacts(v []Contact) {
	o.Contacts = v
}

// GetDraft returns the Draft field value if set, zero value otherwise.
func (o *Order) GetDraft() bool {
	if o == nil || IsNil(o.Draft) {
		var ret bool
		return ret
	}
	return *o.Draft
}

// GetDraftOk returns a tuple with the Draft field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Order) GetDraftOk() (*bool, bool) {
	if o == nil || IsNil(o.Draft) {
		return nil, false
	}
	return o.Draft, true
}

// HasDraft returns a boolean if a field has been set.
func (o *Order) HasDraft() bool {
	if o != nil && !IsNil(o.Draft) {
		return true
	}

	return false
}

// SetDraft gets a reference to the given bool and assigns it to the Draft field.
func (o *Order) SetDraft(v bool) {
	o.Draft = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *Order) GetLinks() []Link {
	if o == nil || IsNil(o.Links) {
		var ret []Link
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Order) GetLinksOk() ([]Link, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *Order) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []Link and assigns it to the Links field.
func (o *Order) SetLinks(v []Link) {
	o.Links = v
}

// GetPurchaseOrder returns the PurchaseOrder field value if set, zero value otherwise.
func (o *Order) GetPurchaseOrder() OrderPurchaseOrder {
	if o == nil || IsNil(o.PurchaseOrder) {
		var ret OrderPurchaseOrder
		return ret
	}
	return *o.PurchaseOrder
}

// GetPurchaseOrderOk returns a tuple with the PurchaseOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Order) GetPurchaseOrderOk() (*OrderPurchaseOrder, bool) {
	if o == nil || IsNil(o.PurchaseOrder) {
		return nil, false
	}
	return o.PurchaseOrder, true
}

// HasPurchaseOrder returns a boolean if a field has been set.
func (o *Order) HasPurchaseOrder() bool {
	if o != nil && !IsNil(o.PurchaseOrder) {
		return true
	}

	return false
}

// SetPurchaseOrder gets a reference to the given OrderPurchaseOrder and assigns it to the PurchaseOrder field.
func (o *Order) SetPurchaseOrder(v OrderPurchaseOrder) {
	o.PurchaseOrder = &v
}

// GetReferenceNumber returns the ReferenceNumber field value if set, zero value otherwise.
func (o *Order) GetReferenceNumber() string {
	if o == nil || IsNil(o.ReferenceNumber) {
		var ret string
		return ret
	}
	return *o.ReferenceNumber
}

// GetReferenceNumberOk returns a tuple with the ReferenceNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Order) GetReferenceNumberOk() (*string, bool) {
	if o == nil || IsNil(o.ReferenceNumber) {
		return nil, false
	}
	return o.ReferenceNumber, true
}

// HasReferenceNumber returns a boolean if a field has been set.
func (o *Order) HasReferenceNumber() bool {
	if o != nil && !IsNil(o.ReferenceNumber) {
		return true
	}

	return false
}

// SetReferenceNumber gets a reference to the given string and assigns it to the ReferenceNumber field.
func (o *Order) SetReferenceNumber(v string) {
	o.ReferenceNumber = &v
}

// GetSignature returns the Signature field value if set, zero value otherwise.
func (o *Order) GetSignature() OrderSignature {
	if o == nil || IsNil(o.Signature) {
		var ret OrderSignature
		return ret
	}
	return *o.Signature
}

// GetSignatureOk returns a tuple with the Signature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Order) GetSignatureOk() (*OrderSignature, bool) {
	if o == nil || IsNil(o.Signature) {
		return nil, false
	}
	return o.Signature, true
}

// HasSignature returns a boolean if a field has been set.
func (o *Order) HasSignature() bool {
	if o != nil && !IsNil(o.Signature) {
		return true
	}

	return false
}

// SetSignature gets a reference to the given OrderSignature and assigns it to the Signature field.
func (o *Order) SetSignature(v OrderSignature) {
	o.Signature = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *Order) GetState() OrderState {
	if o == nil || IsNil(o.State) {
		var ret OrderState
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Order) GetStateOk() (*OrderState, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *Order) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given OrderState and assigns it to the State field.
func (o *Order) SetState(v OrderState) {
	o.State = &v
}

// GetChangeLog returns the ChangeLog field value if set, zero value otherwise.
func (o *Order) GetChangeLog() OrderChangeLog {
	if o == nil || IsNil(o.ChangeLog) {
		var ret OrderChangeLog
		return ret
	}
	return *o.ChangeLog
}

// GetChangeLogOk returns a tuple with the ChangeLog field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Order) GetChangeLogOk() (*OrderChangeLog, bool) {
	if o == nil || IsNil(o.ChangeLog) {
		return nil, false
	}
	return o.ChangeLog, true
}

// HasChangeLog returns a boolean if a field has been set.
func (o *Order) HasChangeLog() bool {
	if o != nil && !IsNil(o.ChangeLog) {
		return true
	}

	return false
}

// SetChangeLog gets a reference to the given OrderChangeLog and assigns it to the ChangeLog field.
func (o *Order) SetChangeLog(v OrderChangeLog) {
	o.ChangeLog = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *Order) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Order) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *Order) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *Order) SetTags(v []string) {
	o.Tags = v
}

func (o Order) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Order) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["href"] = o.Href
	toSerialize["uuid"] = o.Uuid
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Contacts) {
		toSerialize["contacts"] = o.Contacts
	}
	if !IsNil(o.Draft) {
		toSerialize["draft"] = o.Draft
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
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
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.ChangeLog) {
		toSerialize["changeLog"] = o.ChangeLog
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Order) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"href",
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

	varOrder := _Order{}

	err = json.Unmarshal(data, &varOrder)

	if err != nil {
		return err
	}

	*o = Order(varOrder)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "href")
		delete(additionalProperties, "uuid")
		delete(additionalProperties, "type")
		delete(additionalProperties, "contacts")
		delete(additionalProperties, "draft")
		delete(additionalProperties, "links")
		delete(additionalProperties, "purchaseOrder")
		delete(additionalProperties, "referenceNumber")
		delete(additionalProperties, "signature")
		delete(additionalProperties, "state")
		delete(additionalProperties, "changeLog")
		delete(additionalProperties, "tags")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOrder struct {
	value *Order
	isSet bool
}

func (v NullableOrder) Get() *Order {
	return v.value
}

func (v *NullableOrder) Set(val *Order) {
	v.value = val
	v.isSet = true
}

func (v NullableOrder) IsSet() bool {
	return v.isSet
}

func (v *NullableOrder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrder(val *Order) *NullableOrder {
	return &NullableOrder{value: val, isSet: true}
}

func (v NullableOrder) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
