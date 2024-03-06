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

// checks if the Connection type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Connection{}

// Connection Connection specification
type Connection struct {
	Type ConnectionType `json:"type"`
	// Connection URI
	Href *string `json:"href,omitempty"`
	// Equinix-assigned connection identifier
	Uuid *string `json:"uuid,omitempty"`
	// Customer-provided connection name
	Name string `json:"name"`
	// Customer-provided connection description
	Description *string              `json:"description,omitempty"`
	State       *ConnectionState     `json:"state,omitempty"`
	Change      *Change              `json:"change,omitempty"`
	Operation   *ConnectionOperation `json:"operation,omitempty"`
	Order       *Order               `json:"order,omitempty"`
	// Preferences for notifications on connection configuration or status changes
	Notifications []SimplifiedNotification `json:"notifications,omitempty"`
	Account       *SimplifiedAccount       `json:"account,omitempty"`
	ChangeLog     *Changelog               `json:"changeLog,omitempty"`
	// Connection bandwidth in Mbps
	Bandwidth  int32                 `json:"bandwidth"`
	GeoScope   *GeoScopeType         `json:"geoScope,omitempty"`
	Redundancy *ConnectionRedundancy `json:"redundancy,omitempty"`
	// Connection property derived from access point locations
	IsRemote  *bool                `json:"isRemote,omitempty"`
	Direction *ConnectionDirection `json:"direction,omitempty"`
	ASide     ConnectionSide       `json:"aSide"`
	ZSide     ConnectionSide       `json:"zSide"`
	// Connection additional information
	AdditionalInfo       []ConnectionSideAdditionalInfo `json:"additionalInfo,omitempty"`
	Project              *Project                       `json:"project,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Connection Connection

// NewConnection instantiates a new Connection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnection(type_ ConnectionType, name string, bandwidth int32, aSide ConnectionSide, zSide ConnectionSide) *Connection {
	this := Connection{}
	this.Type = type_
	this.Name = name
	this.Bandwidth = bandwidth
	this.ASide = aSide
	this.ZSide = zSide
	return &this
}

// NewConnectionWithDefaults instantiates a new Connection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectionWithDefaults() *Connection {
	this := Connection{}
	return &this
}

// GetType returns the Type field value
func (o *Connection) GetType() ConnectionType {
	if o == nil {
		var ret ConnectionType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Connection) GetTypeOk() (*ConnectionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Connection) SetType(v ConnectionType) {
	o.Type = v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *Connection) GetHref() string {
	if o == nil || IsNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Connection) GetHrefOk() (*string, bool) {
	if o == nil || IsNil(o.Href) {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *Connection) HasHref() bool {
	if o != nil && !IsNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *Connection) SetHref(v string) {
	o.Href = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *Connection) GetUuid() string {
	if o == nil || IsNil(o.Uuid) {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Connection) GetUuidOk() (*string, bool) {
	if o == nil || IsNil(o.Uuid) {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *Connection) HasUuid() bool {
	if o != nil && !IsNil(o.Uuid) {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *Connection) SetUuid(v string) {
	o.Uuid = &v
}

// GetName returns the Name field value
func (o *Connection) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Connection) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Connection) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Connection) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Connection) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Connection) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Connection) SetDescription(v string) {
	o.Description = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *Connection) GetState() ConnectionState {
	if o == nil || IsNil(o.State) {
		var ret ConnectionState
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Connection) GetStateOk() (*ConnectionState, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *Connection) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given ConnectionState and assigns it to the State field.
func (o *Connection) SetState(v ConnectionState) {
	o.State = &v
}

// GetChange returns the Change field value if set, zero value otherwise.
func (o *Connection) GetChange() Change {
	if o == nil || IsNil(o.Change) {
		var ret Change
		return ret
	}
	return *o.Change
}

// GetChangeOk returns a tuple with the Change field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Connection) GetChangeOk() (*Change, bool) {
	if o == nil || IsNil(o.Change) {
		return nil, false
	}
	return o.Change, true
}

// HasChange returns a boolean if a field has been set.
func (o *Connection) HasChange() bool {
	if o != nil && !IsNil(o.Change) {
		return true
	}

	return false
}

// SetChange gets a reference to the given Change and assigns it to the Change field.
func (o *Connection) SetChange(v Change) {
	o.Change = &v
}

// GetOperation returns the Operation field value if set, zero value otherwise.
func (o *Connection) GetOperation() ConnectionOperation {
	if o == nil || IsNil(o.Operation) {
		var ret ConnectionOperation
		return ret
	}
	return *o.Operation
}

// GetOperationOk returns a tuple with the Operation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Connection) GetOperationOk() (*ConnectionOperation, bool) {
	if o == nil || IsNil(o.Operation) {
		return nil, false
	}
	return o.Operation, true
}

// HasOperation returns a boolean if a field has been set.
func (o *Connection) HasOperation() bool {
	if o != nil && !IsNil(o.Operation) {
		return true
	}

	return false
}

// SetOperation gets a reference to the given ConnectionOperation and assigns it to the Operation field.
func (o *Connection) SetOperation(v ConnectionOperation) {
	o.Operation = &v
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *Connection) GetOrder() Order {
	if o == nil || IsNil(o.Order) {
		var ret Order
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Connection) GetOrderOk() (*Order, bool) {
	if o == nil || IsNil(o.Order) {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *Connection) HasOrder() bool {
	if o != nil && !IsNil(o.Order) {
		return true
	}

	return false
}

// SetOrder gets a reference to the given Order and assigns it to the Order field.
func (o *Connection) SetOrder(v Order) {
	o.Order = &v
}

// GetNotifications returns the Notifications field value if set, zero value otherwise.
func (o *Connection) GetNotifications() []SimplifiedNotification {
	if o == nil || IsNil(o.Notifications) {
		var ret []SimplifiedNotification
		return ret
	}
	return o.Notifications
}

// GetNotificationsOk returns a tuple with the Notifications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Connection) GetNotificationsOk() ([]SimplifiedNotification, bool) {
	if o == nil || IsNil(o.Notifications) {
		return nil, false
	}
	return o.Notifications, true
}

// HasNotifications returns a boolean if a field has been set.
func (o *Connection) HasNotifications() bool {
	if o != nil && !IsNil(o.Notifications) {
		return true
	}

	return false
}

// SetNotifications gets a reference to the given []SimplifiedNotification and assigns it to the Notifications field.
func (o *Connection) SetNotifications(v []SimplifiedNotification) {
	o.Notifications = v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *Connection) GetAccount() SimplifiedAccount {
	if o == nil || IsNil(o.Account) {
		var ret SimplifiedAccount
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Connection) GetAccountOk() (*SimplifiedAccount, bool) {
	if o == nil || IsNil(o.Account) {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *Connection) HasAccount() bool {
	if o != nil && !IsNil(o.Account) {
		return true
	}

	return false
}

// SetAccount gets a reference to the given SimplifiedAccount and assigns it to the Account field.
func (o *Connection) SetAccount(v SimplifiedAccount) {
	o.Account = &v
}

// GetChangeLog returns the ChangeLog field value if set, zero value otherwise.
func (o *Connection) GetChangeLog() Changelog {
	if o == nil || IsNil(o.ChangeLog) {
		var ret Changelog
		return ret
	}
	return *o.ChangeLog
}

// GetChangeLogOk returns a tuple with the ChangeLog field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Connection) GetChangeLogOk() (*Changelog, bool) {
	if o == nil || IsNil(o.ChangeLog) {
		return nil, false
	}
	return o.ChangeLog, true
}

// HasChangeLog returns a boolean if a field has been set.
func (o *Connection) HasChangeLog() bool {
	if o != nil && !IsNil(o.ChangeLog) {
		return true
	}

	return false
}

// SetChangeLog gets a reference to the given Changelog and assigns it to the ChangeLog field.
func (o *Connection) SetChangeLog(v Changelog) {
	o.ChangeLog = &v
}

// GetBandwidth returns the Bandwidth field value
func (o *Connection) GetBandwidth() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Bandwidth
}

// GetBandwidthOk returns a tuple with the Bandwidth field value
// and a boolean to check if the value has been set.
func (o *Connection) GetBandwidthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Bandwidth, true
}

// SetBandwidth sets field value
func (o *Connection) SetBandwidth(v int32) {
	o.Bandwidth = v
}

// GetGeoScope returns the GeoScope field value if set, zero value otherwise.
func (o *Connection) GetGeoScope() GeoScopeType {
	if o == nil || IsNil(o.GeoScope) {
		var ret GeoScopeType
		return ret
	}
	return *o.GeoScope
}

// GetGeoScopeOk returns a tuple with the GeoScope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Connection) GetGeoScopeOk() (*GeoScopeType, bool) {
	if o == nil || IsNil(o.GeoScope) {
		return nil, false
	}
	return o.GeoScope, true
}

// HasGeoScope returns a boolean if a field has been set.
func (o *Connection) HasGeoScope() bool {
	if o != nil && !IsNil(o.GeoScope) {
		return true
	}

	return false
}

// SetGeoScope gets a reference to the given GeoScopeType and assigns it to the GeoScope field.
func (o *Connection) SetGeoScope(v GeoScopeType) {
	o.GeoScope = &v
}

// GetRedundancy returns the Redundancy field value if set, zero value otherwise.
func (o *Connection) GetRedundancy() ConnectionRedundancy {
	if o == nil || IsNil(o.Redundancy) {
		var ret ConnectionRedundancy
		return ret
	}
	return *o.Redundancy
}

// GetRedundancyOk returns a tuple with the Redundancy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Connection) GetRedundancyOk() (*ConnectionRedundancy, bool) {
	if o == nil || IsNil(o.Redundancy) {
		return nil, false
	}
	return o.Redundancy, true
}

// HasRedundancy returns a boolean if a field has been set.
func (o *Connection) HasRedundancy() bool {
	if o != nil && !IsNil(o.Redundancy) {
		return true
	}

	return false
}

// SetRedundancy gets a reference to the given ConnectionRedundancy and assigns it to the Redundancy field.
func (o *Connection) SetRedundancy(v ConnectionRedundancy) {
	o.Redundancy = &v
}

// GetIsRemote returns the IsRemote field value if set, zero value otherwise.
func (o *Connection) GetIsRemote() bool {
	if o == nil || IsNil(o.IsRemote) {
		var ret bool
		return ret
	}
	return *o.IsRemote
}

// GetIsRemoteOk returns a tuple with the IsRemote field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Connection) GetIsRemoteOk() (*bool, bool) {
	if o == nil || IsNil(o.IsRemote) {
		return nil, false
	}
	return o.IsRemote, true
}

// HasIsRemote returns a boolean if a field has been set.
func (o *Connection) HasIsRemote() bool {
	if o != nil && !IsNil(o.IsRemote) {
		return true
	}

	return false
}

// SetIsRemote gets a reference to the given bool and assigns it to the IsRemote field.
func (o *Connection) SetIsRemote(v bool) {
	o.IsRemote = &v
}

// GetDirection returns the Direction field value if set, zero value otherwise.
func (o *Connection) GetDirection() ConnectionDirection {
	if o == nil || IsNil(o.Direction) {
		var ret ConnectionDirection
		return ret
	}
	return *o.Direction
}

// GetDirectionOk returns a tuple with the Direction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Connection) GetDirectionOk() (*ConnectionDirection, bool) {
	if o == nil || IsNil(o.Direction) {
		return nil, false
	}
	return o.Direction, true
}

// HasDirection returns a boolean if a field has been set.
func (o *Connection) HasDirection() bool {
	if o != nil && !IsNil(o.Direction) {
		return true
	}

	return false
}

// SetDirection gets a reference to the given ConnectionDirection and assigns it to the Direction field.
func (o *Connection) SetDirection(v ConnectionDirection) {
	o.Direction = &v
}

// GetASide returns the ASide field value
func (o *Connection) GetASide() ConnectionSide {
	if o == nil {
		var ret ConnectionSide
		return ret
	}

	return o.ASide
}

// GetASideOk returns a tuple with the ASide field value
// and a boolean to check if the value has been set.
func (o *Connection) GetASideOk() (*ConnectionSide, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ASide, true
}

// SetASide sets field value
func (o *Connection) SetASide(v ConnectionSide) {
	o.ASide = v
}

// GetZSide returns the ZSide field value
func (o *Connection) GetZSide() ConnectionSide {
	if o == nil {
		var ret ConnectionSide
		return ret
	}

	return o.ZSide
}

// GetZSideOk returns a tuple with the ZSide field value
// and a boolean to check if the value has been set.
func (o *Connection) GetZSideOk() (*ConnectionSide, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ZSide, true
}

// SetZSide sets field value
func (o *Connection) SetZSide(v ConnectionSide) {
	o.ZSide = v
}

// GetAdditionalInfo returns the AdditionalInfo field value if set, zero value otherwise.
func (o *Connection) GetAdditionalInfo() []ConnectionSideAdditionalInfo {
	if o == nil || IsNil(o.AdditionalInfo) {
		var ret []ConnectionSideAdditionalInfo
		return ret
	}
	return o.AdditionalInfo
}

// GetAdditionalInfoOk returns a tuple with the AdditionalInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Connection) GetAdditionalInfoOk() ([]ConnectionSideAdditionalInfo, bool) {
	if o == nil || IsNil(o.AdditionalInfo) {
		return nil, false
	}
	return o.AdditionalInfo, true
}

// HasAdditionalInfo returns a boolean if a field has been set.
func (o *Connection) HasAdditionalInfo() bool {
	if o != nil && !IsNil(o.AdditionalInfo) {
		return true
	}

	return false
}

// SetAdditionalInfo gets a reference to the given []ConnectionSideAdditionalInfo and assigns it to the AdditionalInfo field.
func (o *Connection) SetAdditionalInfo(v []ConnectionSideAdditionalInfo) {
	o.AdditionalInfo = v
}

// GetProject returns the Project field value if set, zero value otherwise.
func (o *Connection) GetProject() Project {
	if o == nil || IsNil(o.Project) {
		var ret Project
		return ret
	}
	return *o.Project
}

// GetProjectOk returns a tuple with the Project field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Connection) GetProjectOk() (*Project, bool) {
	if o == nil || IsNil(o.Project) {
		return nil, false
	}
	return o.Project, true
}

// HasProject returns a boolean if a field has been set.
func (o *Connection) HasProject() bool {
	if o != nil && !IsNil(o.Project) {
		return true
	}

	return false
}

// SetProject gets a reference to the given Project and assigns it to the Project field.
func (o *Connection) SetProject(v Project) {
	o.Project = &v
}

func (o Connection) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Connection) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	if !IsNil(o.Href) {
		toSerialize["href"] = o.Href
	}
	if !IsNil(o.Uuid) {
		toSerialize["uuid"] = o.Uuid
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.Change) {
		toSerialize["change"] = o.Change
	}
	if !IsNil(o.Operation) {
		toSerialize["operation"] = o.Operation
	}
	if !IsNil(o.Order) {
		toSerialize["order"] = o.Order
	}
	if !IsNil(o.Notifications) {
		toSerialize["notifications"] = o.Notifications
	}
	if !IsNil(o.Account) {
		toSerialize["account"] = o.Account
	}
	if !IsNil(o.ChangeLog) {
		toSerialize["changeLog"] = o.ChangeLog
	}
	toSerialize["bandwidth"] = o.Bandwidth
	if !IsNil(o.GeoScope) {
		toSerialize["geoScope"] = o.GeoScope
	}
	if !IsNil(o.Redundancy) {
		toSerialize["redundancy"] = o.Redundancy
	}
	if !IsNil(o.IsRemote) {
		toSerialize["isRemote"] = o.IsRemote
	}
	if !IsNil(o.Direction) {
		toSerialize["direction"] = o.Direction
	}
	toSerialize["aSide"] = o.ASide
	toSerialize["zSide"] = o.ZSide
	if !IsNil(o.AdditionalInfo) {
		toSerialize["additionalInfo"] = o.AdditionalInfo
	}
	if !IsNil(o.Project) {
		toSerialize["project"] = o.Project
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Connection) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"name",
		"bandwidth",
		"aSide",
		"zSide",
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

	varConnection := _Connection{}

	err = json.Unmarshal(data, &varConnection)

	if err != nil {
		return err
	}

	*o = Connection(varConnection)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "href")
		delete(additionalProperties, "uuid")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "state")
		delete(additionalProperties, "change")
		delete(additionalProperties, "operation")
		delete(additionalProperties, "order")
		delete(additionalProperties, "notifications")
		delete(additionalProperties, "account")
		delete(additionalProperties, "changeLog")
		delete(additionalProperties, "bandwidth")
		delete(additionalProperties, "geoScope")
		delete(additionalProperties, "redundancy")
		delete(additionalProperties, "isRemote")
		delete(additionalProperties, "direction")
		delete(additionalProperties, "aSide")
		delete(additionalProperties, "zSide")
		delete(additionalProperties, "additionalInfo")
		delete(additionalProperties, "project")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableConnection struct {
	value *Connection
	isSet bool
}

func (v NullableConnection) Get() *Connection {
	return v.value
}

func (v *NullableConnection) Set(val *Connection) {
	v.value = val
	v.isSet = true
}

func (v NullableConnection) IsSet() bool {
	return v.isSet
}

func (v *NullableConnection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnection(val *Connection) *NullableConnection {
	return &NullableConnection{value: val, isSet: true}
}

func (v NullableConnection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}