# RoutingProtocolReadModelBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tags** | Pointer to **[]string** |  | [optional] 
**Href** | Pointer to **string** | Routing protocol URI | [optional] 
**Uuid** | Pointer to **string** | Routing protocol identifier | [optional] 
**Type** | **string** | Allowed values: - BGP - DIRECT - STATIC  | 
**Name** | **string** | Name of the routing protocol instance. | 
**Description** | Pointer to **string** | Description of the routing protocol instance | [optional] 
**Ipv4** | Pointer to [**RoutingProtocolIpv4**](RoutingProtocolIpv4.md) |  | [optional] 
**Ipv6** | Pointer to [**RoutingProtocolIpv6**](RoutingProtocolIpv6.md) |  | [optional] 
**ChangeLog** | [**ChangeLog**](ChangeLog.md) |  | 

## Methods

### NewRoutingProtocolReadModelBase

`func NewRoutingProtocolReadModelBase(type_ string, name string, changeLog ChangeLog, ) *RoutingProtocolReadModelBase`

NewRoutingProtocolReadModelBase instantiates a new RoutingProtocolReadModelBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoutingProtocolReadModelBaseWithDefaults

`func NewRoutingProtocolReadModelBaseWithDefaults() *RoutingProtocolReadModelBase`

NewRoutingProtocolReadModelBaseWithDefaults instantiates a new RoutingProtocolReadModelBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTags

`func (o *RoutingProtocolReadModelBase) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *RoutingProtocolReadModelBase) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *RoutingProtocolReadModelBase) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *RoutingProtocolReadModelBase) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetHref

`func (o *RoutingProtocolReadModelBase) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *RoutingProtocolReadModelBase) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *RoutingProtocolReadModelBase) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *RoutingProtocolReadModelBase) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetUuid

`func (o *RoutingProtocolReadModelBase) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *RoutingProtocolReadModelBase) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *RoutingProtocolReadModelBase) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *RoutingProtocolReadModelBase) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetType

`func (o *RoutingProtocolReadModelBase) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RoutingProtocolReadModelBase) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RoutingProtocolReadModelBase) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *RoutingProtocolReadModelBase) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RoutingProtocolReadModelBase) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RoutingProtocolReadModelBase) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *RoutingProtocolReadModelBase) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RoutingProtocolReadModelBase) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RoutingProtocolReadModelBase) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RoutingProtocolReadModelBase) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIpv4

`func (o *RoutingProtocolReadModelBase) GetIpv4() RoutingProtocolIpv4`

GetIpv4 returns the Ipv4 field if non-nil, zero value otherwise.

### GetIpv4Ok

`func (o *RoutingProtocolReadModelBase) GetIpv4Ok() (*RoutingProtocolIpv4, bool)`

GetIpv4Ok returns a tuple with the Ipv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4

`func (o *RoutingProtocolReadModelBase) SetIpv4(v RoutingProtocolIpv4)`

SetIpv4 sets Ipv4 field to given value.

### HasIpv4

`func (o *RoutingProtocolReadModelBase) HasIpv4() bool`

HasIpv4 returns a boolean if a field has been set.

### GetIpv6

`func (o *RoutingProtocolReadModelBase) GetIpv6() RoutingProtocolIpv6`

GetIpv6 returns the Ipv6 field if non-nil, zero value otherwise.

### GetIpv6Ok

`func (o *RoutingProtocolReadModelBase) GetIpv6Ok() (*RoutingProtocolIpv6, bool)`

GetIpv6Ok returns a tuple with the Ipv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6

`func (o *RoutingProtocolReadModelBase) SetIpv6(v RoutingProtocolIpv6)`

SetIpv6 sets Ipv6 field to given value.

### HasIpv6

`func (o *RoutingProtocolReadModelBase) HasIpv6() bool`

HasIpv6 returns a boolean if a field has been set.

### GetChangeLog

`func (o *RoutingProtocolReadModelBase) GetChangeLog() ChangeLog`

GetChangeLog returns the ChangeLog field if non-nil, zero value otherwise.

### GetChangeLogOk

`func (o *RoutingProtocolReadModelBase) GetChangeLogOk() (*ChangeLog, bool)`

GetChangeLogOk returns a tuple with the ChangeLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeLog

`func (o *RoutingProtocolReadModelBase) SetChangeLog(v ChangeLog)`

SetChangeLog sets ChangeLog field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


