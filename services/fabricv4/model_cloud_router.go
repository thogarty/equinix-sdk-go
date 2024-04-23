/*
Equinix Fabric API v4

Equinix Fabric is an advanced software-defined interconnection solution that enables you to directly, securely and dynamically connect to distributed infrastructure and digital ecosystems on platform Equinix via a single port, Customers can use Fabric to connect to: </br> 1. Cloud Service Providers - Clouds, network and other service providers.  </br> 2. Enterprises - Other Equinix customers, vendors and partners.  </br> 3. Myself - Another customer instance deployed at Equinix. </br>

Contact: api-support@equinix.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fabricv4

import (
	"encoding/json"
)

// checks if the CloudRouter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CloudRouter{}

// CloudRouter Fabric Cloud Router object
type CloudRouter struct {
	// Cloud Routers URI
	Href *string `json:"href,omitempty"`
	// Equinix-assigned access point identifier
	Uuid *string `json:"uuid,omitempty"`
	// Customer-provided Cloud Router name
	Name  *string                      `json:"name,omitempty"`
	State *CloudRouterAccessPointState `json:"state,omitempty"`
	// Equinix ASN
	EquinixAsn *int64 `json:"equinixAsn,omitempty"`
	// Access point used and maximum number of IPv4 BGP routes
	BgpIpv4RoutesCount *int32 `json:"bgpIpv4RoutesCount,omitempty"`
	// Access point used and maximum number of IPv6 BGP routes
	BgpIpv6RoutesCount *int32 `json:"bgpIpv6RoutesCount,omitempty"`
	// Number of connections associated with this Access point
	ConnectionsCount *int32 `json:"connectionsCount,omitempty"`
	// Number of distinct ipv4 routes
	DistinctIpv4PrefixesCount *int32 `json:"distinctIpv4PrefixesCount,omitempty"`
	// Number of distinct ipv6 routes
	DistinctIpv6PrefixesCount *int32                         `json:"distinctIpv6PrefixesCount,omitempty"`
	ChangeLog                 *Changelog                     `json:"changeLog,omitempty"`
	Change                    *CloudRouterChange             `json:"change,omitempty"`
	Type                      *CloudRouterPostRequestType    `json:"type,omitempty"`
	Location                  *SimplifiedLocationWithoutIBX  `json:"location,omitempty"`
	Package                   *CloudRouterPostRequestPackage `json:"package,omitempty"`
	Order                     *Order                         `json:"order,omitempty"`
	Project                   *Project                       `json:"project,omitempty"`
	Account                   *SimplifiedAccount             `json:"account,omitempty"`
	// Preferences for notifications on connection configuration or status changes
	Notifications        []SimplifiedNotification `json:"notifications,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CloudRouter CloudRouter

// NewCloudRouter instantiates a new CloudRouter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudRouter() *CloudRouter {
	this := CloudRouter{}
	return &this
}

// NewCloudRouterWithDefaults instantiates a new CloudRouter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudRouterWithDefaults() *CloudRouter {
	this := CloudRouter{}
	return &this
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *CloudRouter) GetHref() string {
	if o == nil || IsNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudRouter) GetHrefOk() (*string, bool) {
	if o == nil || IsNil(o.Href) {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *CloudRouter) HasHref() bool {
	if o != nil && !IsNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *CloudRouter) SetHref(v string) {
	o.Href = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *CloudRouter) GetUuid() string {
	if o == nil || IsNil(o.Uuid) {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudRouter) GetUuidOk() (*string, bool) {
	if o == nil || IsNil(o.Uuid) {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *CloudRouter) HasUuid() bool {
	if o != nil && !IsNil(o.Uuid) {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *CloudRouter) SetUuid(v string) {
	o.Uuid = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CloudRouter) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudRouter) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CloudRouter) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CloudRouter) SetName(v string) {
	o.Name = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *CloudRouter) GetState() CloudRouterAccessPointState {
	if o == nil || IsNil(o.State) {
		var ret CloudRouterAccessPointState
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudRouter) GetStateOk() (*CloudRouterAccessPointState, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *CloudRouter) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given CloudRouterAccessPointState and assigns it to the State field.
func (o *CloudRouter) SetState(v CloudRouterAccessPointState) {
	o.State = &v
}

// GetEquinixAsn returns the EquinixAsn field value if set, zero value otherwise.
func (o *CloudRouter) GetEquinixAsn() int64 {
	if o == nil || IsNil(o.EquinixAsn) {
		var ret int64
		return ret
	}
	return *o.EquinixAsn
}

// GetEquinixAsnOk returns a tuple with the EquinixAsn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudRouter) GetEquinixAsnOk() (*int64, bool) {
	if o == nil || IsNil(o.EquinixAsn) {
		return nil, false
	}
	return o.EquinixAsn, true
}

// HasEquinixAsn returns a boolean if a field has been set.
func (o *CloudRouter) HasEquinixAsn() bool {
	if o != nil && !IsNil(o.EquinixAsn) {
		return true
	}

	return false
}

// SetEquinixAsn gets a reference to the given int64 and assigns it to the EquinixAsn field.
func (o *CloudRouter) SetEquinixAsn(v int64) {
	o.EquinixAsn = &v
}

// GetBgpIpv4RoutesCount returns the BgpIpv4RoutesCount field value if set, zero value otherwise.
func (o *CloudRouter) GetBgpIpv4RoutesCount() int32 {
	if o == nil || IsNil(o.BgpIpv4RoutesCount) {
		var ret int32
		return ret
	}
	return *o.BgpIpv4RoutesCount
}

// GetBgpIpv4RoutesCountOk returns a tuple with the BgpIpv4RoutesCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudRouter) GetBgpIpv4RoutesCountOk() (*int32, bool) {
	if o == nil || IsNil(o.BgpIpv4RoutesCount) {
		return nil, false
	}
	return o.BgpIpv4RoutesCount, true
}

// HasBgpIpv4RoutesCount returns a boolean if a field has been set.
func (o *CloudRouter) HasBgpIpv4RoutesCount() bool {
	if o != nil && !IsNil(o.BgpIpv4RoutesCount) {
		return true
	}

	return false
}

// SetBgpIpv4RoutesCount gets a reference to the given int32 and assigns it to the BgpIpv4RoutesCount field.
func (o *CloudRouter) SetBgpIpv4RoutesCount(v int32) {
	o.BgpIpv4RoutesCount = &v
}

// GetBgpIpv6RoutesCount returns the BgpIpv6RoutesCount field value if set, zero value otherwise.
func (o *CloudRouter) GetBgpIpv6RoutesCount() int32 {
	if o == nil || IsNil(o.BgpIpv6RoutesCount) {
		var ret int32
		return ret
	}
	return *o.BgpIpv6RoutesCount
}

// GetBgpIpv6RoutesCountOk returns a tuple with the BgpIpv6RoutesCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudRouter) GetBgpIpv6RoutesCountOk() (*int32, bool) {
	if o == nil || IsNil(o.BgpIpv6RoutesCount) {
		return nil, false
	}
	return o.BgpIpv6RoutesCount, true
}

// HasBgpIpv6RoutesCount returns a boolean if a field has been set.
func (o *CloudRouter) HasBgpIpv6RoutesCount() bool {
	if o != nil && !IsNil(o.BgpIpv6RoutesCount) {
		return true
	}

	return false
}

// SetBgpIpv6RoutesCount gets a reference to the given int32 and assigns it to the BgpIpv6RoutesCount field.
func (o *CloudRouter) SetBgpIpv6RoutesCount(v int32) {
	o.BgpIpv6RoutesCount = &v
}

// GetConnectionsCount returns the ConnectionsCount field value if set, zero value otherwise.
func (o *CloudRouter) GetConnectionsCount() int32 {
	if o == nil || IsNil(o.ConnectionsCount) {
		var ret int32
		return ret
	}
	return *o.ConnectionsCount
}

// GetConnectionsCountOk returns a tuple with the ConnectionsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudRouter) GetConnectionsCountOk() (*int32, bool) {
	if o == nil || IsNil(o.ConnectionsCount) {
		return nil, false
	}
	return o.ConnectionsCount, true
}

// HasConnectionsCount returns a boolean if a field has been set.
func (o *CloudRouter) HasConnectionsCount() bool {
	if o != nil && !IsNil(o.ConnectionsCount) {
		return true
	}

	return false
}

// SetConnectionsCount gets a reference to the given int32 and assigns it to the ConnectionsCount field.
func (o *CloudRouter) SetConnectionsCount(v int32) {
	o.ConnectionsCount = &v
}

// GetDistinctIpv4PrefixesCount returns the DistinctIpv4PrefixesCount field value if set, zero value otherwise.
func (o *CloudRouter) GetDistinctIpv4PrefixesCount() int32 {
	if o == nil || IsNil(o.DistinctIpv4PrefixesCount) {
		var ret int32
		return ret
	}
	return *o.DistinctIpv4PrefixesCount
}

// GetDistinctIpv4PrefixesCountOk returns a tuple with the DistinctIpv4PrefixesCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudRouter) GetDistinctIpv4PrefixesCountOk() (*int32, bool) {
	if o == nil || IsNil(o.DistinctIpv4PrefixesCount) {
		return nil, false
	}
	return o.DistinctIpv4PrefixesCount, true
}

// HasDistinctIpv4PrefixesCount returns a boolean if a field has been set.
func (o *CloudRouter) HasDistinctIpv4PrefixesCount() bool {
	if o != nil && !IsNil(o.DistinctIpv4PrefixesCount) {
		return true
	}

	return false
}

// SetDistinctIpv4PrefixesCount gets a reference to the given int32 and assigns it to the DistinctIpv4PrefixesCount field.
func (o *CloudRouter) SetDistinctIpv4PrefixesCount(v int32) {
	o.DistinctIpv4PrefixesCount = &v
}

// GetDistinctIpv6PrefixesCount returns the DistinctIpv6PrefixesCount field value if set, zero value otherwise.
func (o *CloudRouter) GetDistinctIpv6PrefixesCount() int32 {
	if o == nil || IsNil(o.DistinctIpv6PrefixesCount) {
		var ret int32
		return ret
	}
	return *o.DistinctIpv6PrefixesCount
}

// GetDistinctIpv6PrefixesCountOk returns a tuple with the DistinctIpv6PrefixesCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudRouter) GetDistinctIpv6PrefixesCountOk() (*int32, bool) {
	if o == nil || IsNil(o.DistinctIpv6PrefixesCount) {
		return nil, false
	}
	return o.DistinctIpv6PrefixesCount, true
}

// HasDistinctIpv6PrefixesCount returns a boolean if a field has been set.
func (o *CloudRouter) HasDistinctIpv6PrefixesCount() bool {
	if o != nil && !IsNil(o.DistinctIpv6PrefixesCount) {
		return true
	}

	return false
}

// SetDistinctIpv6PrefixesCount gets a reference to the given int32 and assigns it to the DistinctIpv6PrefixesCount field.
func (o *CloudRouter) SetDistinctIpv6PrefixesCount(v int32) {
	o.DistinctIpv6PrefixesCount = &v
}

// GetChangeLog returns the ChangeLog field value if set, zero value otherwise.
func (o *CloudRouter) GetChangeLog() Changelog {
	if o == nil || IsNil(o.ChangeLog) {
		var ret Changelog
		return ret
	}
	return *o.ChangeLog
}

// GetChangeLogOk returns a tuple with the ChangeLog field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudRouter) GetChangeLogOk() (*Changelog, bool) {
	if o == nil || IsNil(o.ChangeLog) {
		return nil, false
	}
	return o.ChangeLog, true
}

// HasChangeLog returns a boolean if a field has been set.
func (o *CloudRouter) HasChangeLog() bool {
	if o != nil && !IsNil(o.ChangeLog) {
		return true
	}

	return false
}

// SetChangeLog gets a reference to the given Changelog and assigns it to the ChangeLog field.
func (o *CloudRouter) SetChangeLog(v Changelog) {
	o.ChangeLog = &v
}

// GetChange returns the Change field value if set, zero value otherwise.
func (o *CloudRouter) GetChange() CloudRouterChange {
	if o == nil || IsNil(o.Change) {
		var ret CloudRouterChange
		return ret
	}
	return *o.Change
}

// GetChangeOk returns a tuple with the Change field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudRouter) GetChangeOk() (*CloudRouterChange, bool) {
	if o == nil || IsNil(o.Change) {
		return nil, false
	}
	return o.Change, true
}

// HasChange returns a boolean if a field has been set.
func (o *CloudRouter) HasChange() bool {
	if o != nil && !IsNil(o.Change) {
		return true
	}

	return false
}

// SetChange gets a reference to the given CloudRouterChange and assigns it to the Change field.
func (o *CloudRouter) SetChange(v CloudRouterChange) {
	o.Change = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *CloudRouter) GetType() CloudRouterPostRequestType {
	if o == nil || IsNil(o.Type) {
		var ret CloudRouterPostRequestType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudRouter) GetTypeOk() (*CloudRouterPostRequestType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CloudRouter) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given CloudRouterPostRequestType and assigns it to the Type field.
func (o *CloudRouter) SetType(v CloudRouterPostRequestType) {
	o.Type = &v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *CloudRouter) GetLocation() SimplifiedLocationWithoutIBX {
	if o == nil || IsNil(o.Location) {
		var ret SimplifiedLocationWithoutIBX
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudRouter) GetLocationOk() (*SimplifiedLocationWithoutIBX, bool) {
	if o == nil || IsNil(o.Location) {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *CloudRouter) HasLocation() bool {
	if o != nil && !IsNil(o.Location) {
		return true
	}

	return false
}

// SetLocation gets a reference to the given SimplifiedLocationWithoutIBX and assigns it to the Location field.
func (o *CloudRouter) SetLocation(v SimplifiedLocationWithoutIBX) {
	o.Location = &v
}

// GetPackage returns the Package field value if set, zero value otherwise.
func (o *CloudRouter) GetPackage() CloudRouterPostRequestPackage {
	if o == nil || IsNil(o.Package) {
		var ret CloudRouterPostRequestPackage
		return ret
	}
	return *o.Package
}

// GetPackageOk returns a tuple with the Package field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudRouter) GetPackageOk() (*CloudRouterPostRequestPackage, bool) {
	if o == nil || IsNil(o.Package) {
		return nil, false
	}
	return o.Package, true
}

// HasPackage returns a boolean if a field has been set.
func (o *CloudRouter) HasPackage() bool {
	if o != nil && !IsNil(o.Package) {
		return true
	}

	return false
}

// SetPackage gets a reference to the given CloudRouterPostRequestPackage and assigns it to the Package field.
func (o *CloudRouter) SetPackage(v CloudRouterPostRequestPackage) {
	o.Package = &v
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *CloudRouter) GetOrder() Order {
	if o == nil || IsNil(o.Order) {
		var ret Order
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudRouter) GetOrderOk() (*Order, bool) {
	if o == nil || IsNil(o.Order) {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *CloudRouter) HasOrder() bool {
	if o != nil && !IsNil(o.Order) {
		return true
	}

	return false
}

// SetOrder gets a reference to the given Order and assigns it to the Order field.
func (o *CloudRouter) SetOrder(v Order) {
	o.Order = &v
}

// GetProject returns the Project field value if set, zero value otherwise.
func (o *CloudRouter) GetProject() Project {
	if o == nil || IsNil(o.Project) {
		var ret Project
		return ret
	}
	return *o.Project
}

// GetProjectOk returns a tuple with the Project field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudRouter) GetProjectOk() (*Project, bool) {
	if o == nil || IsNil(o.Project) {
		return nil, false
	}
	return o.Project, true
}

// HasProject returns a boolean if a field has been set.
func (o *CloudRouter) HasProject() bool {
	if o != nil && !IsNil(o.Project) {
		return true
	}

	return false
}

// SetProject gets a reference to the given Project and assigns it to the Project field.
func (o *CloudRouter) SetProject(v Project) {
	o.Project = &v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *CloudRouter) GetAccount() SimplifiedAccount {
	if o == nil || IsNil(o.Account) {
		var ret SimplifiedAccount
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudRouter) GetAccountOk() (*SimplifiedAccount, bool) {
	if o == nil || IsNil(o.Account) {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *CloudRouter) HasAccount() bool {
	if o != nil && !IsNil(o.Account) {
		return true
	}

	return false
}

// SetAccount gets a reference to the given SimplifiedAccount and assigns it to the Account field.
func (o *CloudRouter) SetAccount(v SimplifiedAccount) {
	o.Account = &v
}

// GetNotifications returns the Notifications field value if set, zero value otherwise.
func (o *CloudRouter) GetNotifications() []SimplifiedNotification {
	if o == nil || IsNil(o.Notifications) {
		var ret []SimplifiedNotification
		return ret
	}
	return o.Notifications
}

// GetNotificationsOk returns a tuple with the Notifications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudRouter) GetNotificationsOk() ([]SimplifiedNotification, bool) {
	if o == nil || IsNil(o.Notifications) {
		return nil, false
	}
	return o.Notifications, true
}

// HasNotifications returns a boolean if a field has been set.
func (o *CloudRouter) HasNotifications() bool {
	if o != nil && !IsNil(o.Notifications) {
		return true
	}

	return false
}

// SetNotifications gets a reference to the given []SimplifiedNotification and assigns it to the Notifications field.
func (o *CloudRouter) SetNotifications(v []SimplifiedNotification) {
	o.Notifications = v
}

func (o CloudRouter) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CloudRouter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Href) {
		toSerialize["href"] = o.Href
	}
	if !IsNil(o.Uuid) {
		toSerialize["uuid"] = o.Uuid
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.EquinixAsn) {
		toSerialize["equinixAsn"] = o.EquinixAsn
	}
	if !IsNil(o.BgpIpv4RoutesCount) {
		toSerialize["bgpIpv4RoutesCount"] = o.BgpIpv4RoutesCount
	}
	if !IsNil(o.BgpIpv6RoutesCount) {
		toSerialize["bgpIpv6RoutesCount"] = o.BgpIpv6RoutesCount
	}
	if !IsNil(o.ConnectionsCount) {
		toSerialize["connectionsCount"] = o.ConnectionsCount
	}
	if !IsNil(o.DistinctIpv4PrefixesCount) {
		toSerialize["distinctIpv4PrefixesCount"] = o.DistinctIpv4PrefixesCount
	}
	if !IsNil(o.DistinctIpv6PrefixesCount) {
		toSerialize["distinctIpv6PrefixesCount"] = o.DistinctIpv6PrefixesCount
	}
	if !IsNil(o.ChangeLog) {
		toSerialize["changeLog"] = o.ChangeLog
	}
	if !IsNil(o.Change) {
		toSerialize["change"] = o.Change
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
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

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CloudRouter) UnmarshalJSON(data []byte) (err error) {
	varCloudRouter := _CloudRouter{}

	err = json.Unmarshal(data, &varCloudRouter)

	if err != nil {
		return err
	}

	*o = CloudRouter(varCloudRouter)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "href")
		delete(additionalProperties, "uuid")
		delete(additionalProperties, "name")
		delete(additionalProperties, "state")
		delete(additionalProperties, "equinixAsn")
		delete(additionalProperties, "bgpIpv4RoutesCount")
		delete(additionalProperties, "bgpIpv6RoutesCount")
		delete(additionalProperties, "connectionsCount")
		delete(additionalProperties, "distinctIpv4PrefixesCount")
		delete(additionalProperties, "distinctIpv6PrefixesCount")
		delete(additionalProperties, "changeLog")
		delete(additionalProperties, "change")
		delete(additionalProperties, "type")
		delete(additionalProperties, "location")
		delete(additionalProperties, "package")
		delete(additionalProperties, "order")
		delete(additionalProperties, "project")
		delete(additionalProperties, "account")
		delete(additionalProperties, "notifications")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCloudRouter struct {
	value *CloudRouter
	isSet bool
}

func (v NullableCloudRouter) Get() *CloudRouter {
	return v.value
}

func (v *NullableCloudRouter) Set(val *CloudRouter) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudRouter) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudRouter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudRouter(val *CloudRouter) *NullableCloudRouter {
	return &NullableCloudRouter{value: val, isSet: true}
}

func (v NullableCloudRouter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudRouter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
