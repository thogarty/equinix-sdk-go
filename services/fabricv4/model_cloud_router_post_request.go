/*
Equinix Fabric API v4

Contact: api-support@equinix.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fabricv4

import (
	"encoding/json"
)

// checks if the CloudRouterPostRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CloudRouterPostRequest{}

// CloudRouterPostRequest Create Cloud Router
type CloudRouterPostRequest struct {
	Type *CloudRouterPostRequestType `json:"type,omitempty"`
	// Customer-provided Cloud Router name
	Name     *string                        `json:"name,omitempty"`
	Location *SimplifiedLocationWithoutIBX  `json:"location,omitempty"`
	Package  *CloudRouterPostRequestPackage `json:"package,omitempty"`
	Order    *Order                         `json:"order,omitempty"`
	Project  *Project                       `json:"project,omitempty"`
	Account  *SimplifiedAccount             `json:"account,omitempty"`
	// Preferences for notifications on connection configuration or status changes
	Notifications           []SimplifiedNotification `json:"notifications,omitempty"`
	MarketplaceSubscription *MarketplaceSubscription `json:"marketplaceSubscription,omitempty"`
	AdditionalProperties    map[string]interface{}
}

type _CloudRouterPostRequest CloudRouterPostRequest

// NewCloudRouterPostRequest instantiates a new CloudRouterPostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudRouterPostRequest() *CloudRouterPostRequest {
	this := CloudRouterPostRequest{}
	return &this
}

// NewCloudRouterPostRequestWithDefaults instantiates a new CloudRouterPostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudRouterPostRequestWithDefaults() *CloudRouterPostRequest {
	this := CloudRouterPostRequest{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *CloudRouterPostRequest) GetType() CloudRouterPostRequestType {
	if o == nil || IsNil(o.Type) {
		var ret CloudRouterPostRequestType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudRouterPostRequest) GetTypeOk() (*CloudRouterPostRequestType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CloudRouterPostRequest) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given CloudRouterPostRequestType and assigns it to the Type field.
func (o *CloudRouterPostRequest) SetType(v CloudRouterPostRequestType) {
	o.Type = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CloudRouterPostRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudRouterPostRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CloudRouterPostRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CloudRouterPostRequest) SetName(v string) {
	o.Name = &v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *CloudRouterPostRequest) GetLocation() SimplifiedLocationWithoutIBX {
	if o == nil || IsNil(o.Location) {
		var ret SimplifiedLocationWithoutIBX
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudRouterPostRequest) GetLocationOk() (*SimplifiedLocationWithoutIBX, bool) {
	if o == nil || IsNil(o.Location) {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *CloudRouterPostRequest) HasLocation() bool {
	if o != nil && !IsNil(o.Location) {
		return true
	}

	return false
}

// SetLocation gets a reference to the given SimplifiedLocationWithoutIBX and assigns it to the Location field.
func (o *CloudRouterPostRequest) SetLocation(v SimplifiedLocationWithoutIBX) {
	o.Location = &v
}

// GetPackage returns the Package field value if set, zero value otherwise.
func (o *CloudRouterPostRequest) GetPackage() CloudRouterPostRequestPackage {
	if o == nil || IsNil(o.Package) {
		var ret CloudRouterPostRequestPackage
		return ret
	}
	return *o.Package
}

// GetPackageOk returns a tuple with the Package field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudRouterPostRequest) GetPackageOk() (*CloudRouterPostRequestPackage, bool) {
	if o == nil || IsNil(o.Package) {
		return nil, false
	}
	return o.Package, true
}

// HasPackage returns a boolean if a field has been set.
func (o *CloudRouterPostRequest) HasPackage() bool {
	if o != nil && !IsNil(o.Package) {
		return true
	}

	return false
}

// SetPackage gets a reference to the given CloudRouterPostRequestPackage and assigns it to the Package field.
func (o *CloudRouterPostRequest) SetPackage(v CloudRouterPostRequestPackage) {
	o.Package = &v
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *CloudRouterPostRequest) GetOrder() Order {
	if o == nil || IsNil(o.Order) {
		var ret Order
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudRouterPostRequest) GetOrderOk() (*Order, bool) {
	if o == nil || IsNil(o.Order) {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *CloudRouterPostRequest) HasOrder() bool {
	if o != nil && !IsNil(o.Order) {
		return true
	}

	return false
}

// SetOrder gets a reference to the given Order and assigns it to the Order field.
func (o *CloudRouterPostRequest) SetOrder(v Order) {
	o.Order = &v
}

// GetProject returns the Project field value if set, zero value otherwise.
func (o *CloudRouterPostRequest) GetProject() Project {
	if o == nil || IsNil(o.Project) {
		var ret Project
		return ret
	}
	return *o.Project
}

// GetProjectOk returns a tuple with the Project field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudRouterPostRequest) GetProjectOk() (*Project, bool) {
	if o == nil || IsNil(o.Project) {
		return nil, false
	}
	return o.Project, true
}

// HasProject returns a boolean if a field has been set.
func (o *CloudRouterPostRequest) HasProject() bool {
	if o != nil && !IsNil(o.Project) {
		return true
	}

	return false
}

// SetProject gets a reference to the given Project and assigns it to the Project field.
func (o *CloudRouterPostRequest) SetProject(v Project) {
	o.Project = &v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *CloudRouterPostRequest) GetAccount() SimplifiedAccount {
	if o == nil || IsNil(o.Account) {
		var ret SimplifiedAccount
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudRouterPostRequest) GetAccountOk() (*SimplifiedAccount, bool) {
	if o == nil || IsNil(o.Account) {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *CloudRouterPostRequest) HasAccount() bool {
	if o != nil && !IsNil(o.Account) {
		return true
	}

	return false
}

// SetAccount gets a reference to the given SimplifiedAccount and assigns it to the Account field.
func (o *CloudRouterPostRequest) SetAccount(v SimplifiedAccount) {
	o.Account = &v
}

// GetNotifications returns the Notifications field value if set, zero value otherwise.
func (o *CloudRouterPostRequest) GetNotifications() []SimplifiedNotification {
	if o == nil || IsNil(o.Notifications) {
		var ret []SimplifiedNotification
		return ret
	}
	return o.Notifications
}

// GetNotificationsOk returns a tuple with the Notifications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudRouterPostRequest) GetNotificationsOk() ([]SimplifiedNotification, bool) {
	if o == nil || IsNil(o.Notifications) {
		return nil, false
	}
	return o.Notifications, true
}

// HasNotifications returns a boolean if a field has been set.
func (o *CloudRouterPostRequest) HasNotifications() bool {
	if o != nil && !IsNil(o.Notifications) {
		return true
	}

	return false
}

// SetNotifications gets a reference to the given []SimplifiedNotification and assigns it to the Notifications field.
func (o *CloudRouterPostRequest) SetNotifications(v []SimplifiedNotification) {
	o.Notifications = v
}

// GetMarketplaceSubscription returns the MarketplaceSubscription field value if set, zero value otherwise.
func (o *CloudRouterPostRequest) GetMarketplaceSubscription() MarketplaceSubscription {
	if o == nil || IsNil(o.MarketplaceSubscription) {
		var ret MarketplaceSubscription
		return ret
	}
	return *o.MarketplaceSubscription
}

// GetMarketplaceSubscriptionOk returns a tuple with the MarketplaceSubscription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudRouterPostRequest) GetMarketplaceSubscriptionOk() (*MarketplaceSubscription, bool) {
	if o == nil || IsNil(o.MarketplaceSubscription) {
		return nil, false
	}
	return o.MarketplaceSubscription, true
}

// HasMarketplaceSubscription returns a boolean if a field has been set.
func (o *CloudRouterPostRequest) HasMarketplaceSubscription() bool {
	if o != nil && !IsNil(o.MarketplaceSubscription) {
		return true
	}

	return false
}

// SetMarketplaceSubscription gets a reference to the given MarketplaceSubscription and assigns it to the MarketplaceSubscription field.
func (o *CloudRouterPostRequest) SetMarketplaceSubscription(v MarketplaceSubscription) {
	o.MarketplaceSubscription = &v
}

func (o CloudRouterPostRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CloudRouterPostRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Location) {
		toSerialize["location"] = o.Location
	}
	if !IsNil(o.Package) {
		toSerialize["package"] = o.Package
	}
	if !IsNil(o.Order) {
		toSerialize["order"] = o.Order
	}
	if !IsNil(o.Project) {
		toSerialize["project"] = o.Project
	}
	if !IsNil(o.Account) {
		toSerialize["account"] = o.Account
	}
	if !IsNil(o.Notifications) {
		toSerialize["notifications"] = o.Notifications
	}
	if !IsNil(o.MarketplaceSubscription) {
		toSerialize["marketplaceSubscription"] = o.MarketplaceSubscription
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CloudRouterPostRequest) UnmarshalJSON(data []byte) (err error) {
	varCloudRouterPostRequest := _CloudRouterPostRequest{}

	err = json.Unmarshal(data, &varCloudRouterPostRequest)

	if err != nil {
		return err
	}

	*o = CloudRouterPostRequest(varCloudRouterPostRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "name")
		delete(additionalProperties, "location")
		delete(additionalProperties, "package")
		delete(additionalProperties, "order")
		delete(additionalProperties, "project")
		delete(additionalProperties, "account")
		delete(additionalProperties, "notifications")
		delete(additionalProperties, "marketplaceSubscription")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCloudRouterPostRequest struct {
	value *CloudRouterPostRequest
	isSet bool
}

func (v NullableCloudRouterPostRequest) Get() *CloudRouterPostRequest {
	return v.value
}

func (v *NullableCloudRouterPostRequest) Set(val *CloudRouterPostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudRouterPostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudRouterPostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudRouterPostRequest(val *CloudRouterPostRequest) *NullableCloudRouterPostRequest {
	return &NullableCloudRouterPostRequest{value: val, isSet: true}
}

func (v NullableCloudRouterPostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudRouterPostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
