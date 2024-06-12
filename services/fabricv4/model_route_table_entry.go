/*
Equinix Fabric API v4

Contact: api-support@equinix.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fabricv4

import (
	"encoding/json"
	"fmt"
)

// checks if the RouteTableEntry type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RouteTableEntry{}

// RouteTableEntry Route table entry object
type RouteTableEntry struct {
	Type                 RouteTableEntryType          `json:"type"`
	ProtocolType         *RouteTableEntryProtocolType `json:"protocolType,omitempty"`
	State                RouteTableEntryState         `json:"state"`
	Age                  *string                      `json:"age,omitempty"`
	Prefix               *string                      `json:"prefix,omitempty"`
	NextHop              *string                      `json:"nextHop,omitempty"`
	Metric               *int32                       `json:"metric,omitempty"`
	LocalPreference      *int32                       `json:"localPreference,omitempty"`
	AsPath               []int32                      `json:"asPath,omitempty"`
	Connection           *RouteTableEntryConnection   `json:"connection,omitempty"`
	ChangeLog            Changelog                    `json:"changeLog"`
	AdditionalProperties map[string]interface{}
}

type _RouteTableEntry RouteTableEntry

// NewRouteTableEntry instantiates a new RouteTableEntry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRouteTableEntry(type_ RouteTableEntryType, state RouteTableEntryState, changeLog Changelog) *RouteTableEntry {
	this := RouteTableEntry{}
	this.Type = type_
	this.State = state
	this.ChangeLog = changeLog
	return &this
}

// NewRouteTableEntryWithDefaults instantiates a new RouteTableEntry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRouteTableEntryWithDefaults() *RouteTableEntry {
	this := RouteTableEntry{}
	return &this
}

// GetType returns the Type field value
func (o *RouteTableEntry) GetType() RouteTableEntryType {
	if o == nil {
		var ret RouteTableEntryType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *RouteTableEntry) GetTypeOk() (*RouteTableEntryType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *RouteTableEntry) SetType(v RouteTableEntryType) {
	o.Type = v
}

// GetProtocolType returns the ProtocolType field value if set, zero value otherwise.
func (o *RouteTableEntry) GetProtocolType() RouteTableEntryProtocolType {
	if o == nil || IsNil(o.ProtocolType) {
		var ret RouteTableEntryProtocolType
		return ret
	}
	return *o.ProtocolType
}

// GetProtocolTypeOk returns a tuple with the ProtocolType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RouteTableEntry) GetProtocolTypeOk() (*RouteTableEntryProtocolType, bool) {
	if o == nil || IsNil(o.ProtocolType) {
		return nil, false
	}
	return o.ProtocolType, true
}

// HasProtocolType returns a boolean if a field has been set.
func (o *RouteTableEntry) HasProtocolType() bool {
	if o != nil && !IsNil(o.ProtocolType) {
		return true
	}

	return false
}

// SetProtocolType gets a reference to the given RouteTableEntryProtocolType and assigns it to the ProtocolType field.
func (o *RouteTableEntry) SetProtocolType(v RouteTableEntryProtocolType) {
	o.ProtocolType = &v
}

// GetState returns the State field value
func (o *RouteTableEntry) GetState() RouteTableEntryState {
	if o == nil {
		var ret RouteTableEntryState
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *RouteTableEntry) GetStateOk() (*RouteTableEntryState, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *RouteTableEntry) SetState(v RouteTableEntryState) {
	o.State = v
}

// GetAge returns the Age field value if set, zero value otherwise.
func (o *RouteTableEntry) GetAge() string {
	if o == nil || IsNil(o.Age) {
		var ret string
		return ret
	}
	return *o.Age
}

// GetAgeOk returns a tuple with the Age field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RouteTableEntry) GetAgeOk() (*string, bool) {
	if o == nil || IsNil(o.Age) {
		return nil, false
	}
	return o.Age, true
}

// HasAge returns a boolean if a field has been set.
func (o *RouteTableEntry) HasAge() bool {
	if o != nil && !IsNil(o.Age) {
		return true
	}

	return false
}

// SetAge gets a reference to the given string and assigns it to the Age field.
func (o *RouteTableEntry) SetAge(v string) {
	o.Age = &v
}

// GetPrefix returns the Prefix field value if set, zero value otherwise.
func (o *RouteTableEntry) GetPrefix() string {
	if o == nil || IsNil(o.Prefix) {
		var ret string
		return ret
	}
	return *o.Prefix
}

// GetPrefixOk returns a tuple with the Prefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RouteTableEntry) GetPrefixOk() (*string, bool) {
	if o == nil || IsNil(o.Prefix) {
		return nil, false
	}
	return o.Prefix, true
}

// HasPrefix returns a boolean if a field has been set.
func (o *RouteTableEntry) HasPrefix() bool {
	if o != nil && !IsNil(o.Prefix) {
		return true
	}

	return false
}

// SetPrefix gets a reference to the given string and assigns it to the Prefix field.
func (o *RouteTableEntry) SetPrefix(v string) {
	o.Prefix = &v
}

// GetNextHop returns the NextHop field value if set, zero value otherwise.
func (o *RouteTableEntry) GetNextHop() string {
	if o == nil || IsNil(o.NextHop) {
		var ret string
		return ret
	}
	return *o.NextHop
}

// GetNextHopOk returns a tuple with the NextHop field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RouteTableEntry) GetNextHopOk() (*string, bool) {
	if o == nil || IsNil(o.NextHop) {
		return nil, false
	}
	return o.NextHop, true
}

// HasNextHop returns a boolean if a field has been set.
func (o *RouteTableEntry) HasNextHop() bool {
	if o != nil && !IsNil(o.NextHop) {
		return true
	}

	return false
}

// SetNextHop gets a reference to the given string and assigns it to the NextHop field.
func (o *RouteTableEntry) SetNextHop(v string) {
	o.NextHop = &v
}

// GetMetric returns the Metric field value if set, zero value otherwise.
func (o *RouteTableEntry) GetMetric() int32 {
	if o == nil || IsNil(o.Metric) {
		var ret int32
		return ret
	}
	return *o.Metric
}

// GetMetricOk returns a tuple with the Metric field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RouteTableEntry) GetMetricOk() (*int32, bool) {
	if o == nil || IsNil(o.Metric) {
		return nil, false
	}
	return o.Metric, true
}

// HasMetric returns a boolean if a field has been set.
func (o *RouteTableEntry) HasMetric() bool {
	if o != nil && !IsNil(o.Metric) {
		return true
	}

	return false
}

// SetMetric gets a reference to the given int32 and assigns it to the Metric field.
func (o *RouteTableEntry) SetMetric(v int32) {
	o.Metric = &v
}

// GetLocalPreference returns the LocalPreference field value if set, zero value otherwise.
func (o *RouteTableEntry) GetLocalPreference() int32 {
	if o == nil || IsNil(o.LocalPreference) {
		var ret int32
		return ret
	}
	return *o.LocalPreference
}

// GetLocalPreferenceOk returns a tuple with the LocalPreference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RouteTableEntry) GetLocalPreferenceOk() (*int32, bool) {
	if o == nil || IsNil(o.LocalPreference) {
		return nil, false
	}
	return o.LocalPreference, true
}

// HasLocalPreference returns a boolean if a field has been set.
func (o *RouteTableEntry) HasLocalPreference() bool {
	if o != nil && !IsNil(o.LocalPreference) {
		return true
	}

	return false
}

// SetLocalPreference gets a reference to the given int32 and assigns it to the LocalPreference field.
func (o *RouteTableEntry) SetLocalPreference(v int32) {
	o.LocalPreference = &v
}

// GetAsPath returns the AsPath field value if set, zero value otherwise.
func (o *RouteTableEntry) GetAsPath() []int32 {
	if o == nil || IsNil(o.AsPath) {
		var ret []int32
		return ret
	}
	return o.AsPath
}

// GetAsPathOk returns a tuple with the AsPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RouteTableEntry) GetAsPathOk() ([]int32, bool) {
	if o == nil || IsNil(o.AsPath) {
		return nil, false
	}
	return o.AsPath, true
}

// HasAsPath returns a boolean if a field has been set.
func (o *RouteTableEntry) HasAsPath() bool {
	if o != nil && !IsNil(o.AsPath) {
		return true
	}

	return false
}

// SetAsPath gets a reference to the given []int32 and assigns it to the AsPath field.
func (o *RouteTableEntry) SetAsPath(v []int32) {
	o.AsPath = v
}

// GetConnection returns the Connection field value if set, zero value otherwise.
func (o *RouteTableEntry) GetConnection() RouteTableEntryConnection {
	if o == nil || IsNil(o.Connection) {
		var ret RouteTableEntryConnection
		return ret
	}
	return *o.Connection
}

// GetConnectionOk returns a tuple with the Connection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RouteTableEntry) GetConnectionOk() (*RouteTableEntryConnection, bool) {
	if o == nil || IsNil(o.Connection) {
		return nil, false
	}
	return o.Connection, true
}

// HasConnection returns a boolean if a field has been set.
func (o *RouteTableEntry) HasConnection() bool {
	if o != nil && !IsNil(o.Connection) {
		return true
	}

	return false
}

// SetConnection gets a reference to the given RouteTableEntryConnection and assigns it to the Connection field.
func (o *RouteTableEntry) SetConnection(v RouteTableEntryConnection) {
	o.Connection = &v
}

// GetChangeLog returns the ChangeLog field value
func (o *RouteTableEntry) GetChangeLog() Changelog {
	if o == nil {
		var ret Changelog
		return ret
	}

	return o.ChangeLog
}

// GetChangeLogOk returns a tuple with the ChangeLog field value
// and a boolean to check if the value has been set.
func (o *RouteTableEntry) GetChangeLogOk() (*Changelog, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChangeLog, true
}

// SetChangeLog sets field value
func (o *RouteTableEntry) SetChangeLog(v Changelog) {
	o.ChangeLog = v
}

func (o RouteTableEntry) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RouteTableEntry) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	if !IsNil(o.ProtocolType) {
		toSerialize["protocolType"] = o.ProtocolType
	}
	toSerialize["state"] = o.State
	if !IsNil(o.Age) {
		toSerialize["age"] = o.Age
	}
	if !IsNil(o.Prefix) {
		toSerialize["prefix"] = o.Prefix
	}
	if !IsNil(o.NextHop) {
		toSerialize["nextHop"] = o.NextHop
	}
	if !IsNil(o.Metric) {
		toSerialize["metric"] = o.Metric
	}
	if !IsNil(o.LocalPreference) {
		toSerialize["localPreference"] = o.LocalPreference
	}
	if !IsNil(o.AsPath) {
		toSerialize["asPath"] = o.AsPath
	}
	if !IsNil(o.Connection) {
		toSerialize["connection"] = o.Connection
	}
	toSerialize["changeLog"] = o.ChangeLog

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RouteTableEntry) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"state",
		"changeLog",
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

	varRouteTableEntry := _RouteTableEntry{}

	err = json.Unmarshal(data, &varRouteTableEntry)

	if err != nil {
		return err
	}

	*o = RouteTableEntry(varRouteTableEntry)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "protocolType")
		delete(additionalProperties, "state")
		delete(additionalProperties, "age")
		delete(additionalProperties, "prefix")
		delete(additionalProperties, "nextHop")
		delete(additionalProperties, "metric")
		delete(additionalProperties, "localPreference")
		delete(additionalProperties, "asPath")
		delete(additionalProperties, "connection")
		delete(additionalProperties, "changeLog")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRouteTableEntry struct {
	value *RouteTableEntry
	isSet bool
}

func (v NullableRouteTableEntry) Get() *RouteTableEntry {
	return v.value
}

func (v *NullableRouteTableEntry) Set(val *RouteTableEntry) {
	v.value = val
	v.isSet = true
}

func (v NullableRouteTableEntry) IsSet() bool {
	return v.isSet
}

func (v *NullableRouteTableEntry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRouteTableEntry(val *RouteTableEntry) *NullableRouteTableEntry {
	return &NullableRouteTableEntry{value: val, isSet: true}
}

func (v NullableRouteTableEntry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRouteTableEntry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
