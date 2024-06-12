/*
Equinix Fabric API v4

Contact: api-support@equinix.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fabricv4

import (
	"encoding/json"
)

// checks if the Order type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Order{}

// Order struct for Order
type Order struct {
	// Purchase order number
	PurchaseOrderNumber *string `json:"purchaseOrderNumber,omitempty"`
	// Customer reference number
	CustomerReferenceNumber *string `json:"customerReferenceNumber,omitempty"`
	// Billing tier for connection bandwidth
	BillingTier *string `json:"billingTier,omitempty"`
	// Order Identification
	OrderId *string `json:"orderId,omitempty"`
	// Order Reference Number
	OrderNumber          *string `json:"orderNumber,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Order Order

// NewOrder instantiates a new Order object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrder() *Order {
	this := Order{}
	return &this
}

// NewOrderWithDefaults instantiates a new Order object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderWithDefaults() *Order {
	this := Order{}
	return &this
}

// GetPurchaseOrderNumber returns the PurchaseOrderNumber field value if set, zero value otherwise.
func (o *Order) GetPurchaseOrderNumber() string {
	if o == nil || IsNil(o.PurchaseOrderNumber) {
		var ret string
		return ret
	}
	return *o.PurchaseOrderNumber
}

// GetPurchaseOrderNumberOk returns a tuple with the PurchaseOrderNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Order) GetPurchaseOrderNumberOk() (*string, bool) {
	if o == nil || IsNil(o.PurchaseOrderNumber) {
		return nil, false
	}
	return o.PurchaseOrderNumber, true
}

// HasPurchaseOrderNumber returns a boolean if a field has been set.
func (o *Order) HasPurchaseOrderNumber() bool {
	if o != nil && !IsNil(o.PurchaseOrderNumber) {
		return true
	}

	return false
}

// SetPurchaseOrderNumber gets a reference to the given string and assigns it to the PurchaseOrderNumber field.
func (o *Order) SetPurchaseOrderNumber(v string) {
	o.PurchaseOrderNumber = &v
}

// GetCustomerReferenceNumber returns the CustomerReferenceNumber field value if set, zero value otherwise.
func (o *Order) GetCustomerReferenceNumber() string {
	if o == nil || IsNil(o.CustomerReferenceNumber) {
		var ret string
		return ret
	}
	return *o.CustomerReferenceNumber
}

// GetCustomerReferenceNumberOk returns a tuple with the CustomerReferenceNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Order) GetCustomerReferenceNumberOk() (*string, bool) {
	if o == nil || IsNil(o.CustomerReferenceNumber) {
		return nil, false
	}
	return o.CustomerReferenceNumber, true
}

// HasCustomerReferenceNumber returns a boolean if a field has been set.
func (o *Order) HasCustomerReferenceNumber() bool {
	if o != nil && !IsNil(o.CustomerReferenceNumber) {
		return true
	}

	return false
}

// SetCustomerReferenceNumber gets a reference to the given string and assigns it to the CustomerReferenceNumber field.
func (o *Order) SetCustomerReferenceNumber(v string) {
	o.CustomerReferenceNumber = &v
}

// GetBillingTier returns the BillingTier field value if set, zero value otherwise.
func (o *Order) GetBillingTier() string {
	if o == nil || IsNil(o.BillingTier) {
		var ret string
		return ret
	}
	return *o.BillingTier
}

// GetBillingTierOk returns a tuple with the BillingTier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Order) GetBillingTierOk() (*string, bool) {
	if o == nil || IsNil(o.BillingTier) {
		return nil, false
	}
	return o.BillingTier, true
}

// HasBillingTier returns a boolean if a field has been set.
func (o *Order) HasBillingTier() bool {
	if o != nil && !IsNil(o.BillingTier) {
		return true
	}

	return false
}

// SetBillingTier gets a reference to the given string and assigns it to the BillingTier field.
func (o *Order) SetBillingTier(v string) {
	o.BillingTier = &v
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *Order) GetOrderId() string {
	if o == nil || IsNil(o.OrderId) {
		var ret string
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Order) GetOrderIdOk() (*string, bool) {
	if o == nil || IsNil(o.OrderId) {
		return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *Order) HasOrderId() bool {
	if o != nil && !IsNil(o.OrderId) {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given string and assigns it to the OrderId field.
func (o *Order) SetOrderId(v string) {
	o.OrderId = &v
}

// GetOrderNumber returns the OrderNumber field value if set, zero value otherwise.
func (o *Order) GetOrderNumber() string {
	if o == nil || IsNil(o.OrderNumber) {
		var ret string
		return ret
	}
	return *o.OrderNumber
}

// GetOrderNumberOk returns a tuple with the OrderNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Order) GetOrderNumberOk() (*string, bool) {
	if o == nil || IsNil(o.OrderNumber) {
		return nil, false
	}
	return o.OrderNumber, true
}

// HasOrderNumber returns a boolean if a field has been set.
func (o *Order) HasOrderNumber() bool {
	if o != nil && !IsNil(o.OrderNumber) {
		return true
	}

	return false
}

// SetOrderNumber gets a reference to the given string and assigns it to the OrderNumber field.
func (o *Order) SetOrderNumber(v string) {
	o.OrderNumber = &v
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
	if !IsNil(o.PurchaseOrderNumber) {
		toSerialize["purchaseOrderNumber"] = o.PurchaseOrderNumber
	}
	if !IsNil(o.CustomerReferenceNumber) {
		toSerialize["customerReferenceNumber"] = o.CustomerReferenceNumber
	}
	if !IsNil(o.BillingTier) {
		toSerialize["billingTier"] = o.BillingTier
	}
	if !IsNil(o.OrderId) {
		toSerialize["orderId"] = o.OrderId
	}
	if !IsNil(o.OrderNumber) {
		toSerialize["orderNumber"] = o.OrderNumber
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Order) UnmarshalJSON(data []byte) (err error) {
	varOrder := _Order{}

	err = json.Unmarshal(data, &varOrder)

	if err != nil {
		return err
	}

	*o = Order(varOrder)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "purchaseOrderNumber")
		delete(additionalProperties, "customerReferenceNumber")
		delete(additionalProperties, "billingTier")
		delete(additionalProperties, "orderId")
		delete(additionalProperties, "orderNumber")
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
