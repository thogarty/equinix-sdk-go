# RoutingProtocolReadModelStatic

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

### NewRoutingProtocolReadModelStatic

`func NewRoutingProtocolReadModelStatic(type_ string, name string, changeLog ChangeLog, ) *RoutingProtocolReadModelStatic`

NewRoutingProtocolReadModelStatic instantiates a new RoutingProtocolReadModelStatic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoutingProtocolReadModelStaticWithDefaults

`func NewRoutingProtocolReadModelStaticWithDefaults() *RoutingProtocolReadModelStatic`

NewRoutingProtocolReadModelStaticWithDefaults instantiates a new RoutingProtocolReadModelStatic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTags

`func (o *RoutingProtocolReadModelStatic) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *RoutingProtocolReadModelStatic) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *RoutingProtocolReadModelStatic) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *RoutingProtocolReadModelStatic) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetHref

`func (o *RoutingProtocolReadModelStatic) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *RoutingProtocolReadModelStatic) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *RoutingProtocolReadModelStatic) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *RoutingProtocolReadModelStatic) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetUuid

`func (o *RoutingProtocolReadModelStatic) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *RoutingProtocolReadModelStatic) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *RoutingProtocolReadModelStatic) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *RoutingProtocolReadModelStatic) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetType

`func (o *RoutingProtocolReadModelStatic) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RoutingProtocolReadModelStatic) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RoutingProtocolReadModelStatic) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *RoutingProtocolReadModelStatic) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RoutingProtocolReadModelStatic) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RoutingProtocolReadModelStatic) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *RoutingProtocolReadModelStatic) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RoutingProtocolReadModelStatic) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RoutingProtocolReadModelStatic) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RoutingProtocolReadModelStatic) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIpv4

`func (o *RoutingProtocolReadModelStatic) GetIpv4() RoutingProtocolIpv4`

GetIpv4 returns the Ipv4 field if non-nil, zero value otherwise.

### GetIpv4Ok

`func (o *RoutingProtocolReadModelStatic) GetIpv4Ok() (*RoutingProtocolIpv4, bool)`

GetIpv4Ok returns a tuple with the Ipv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4

`func (o *RoutingProtocolReadModelStatic) SetIpv4(v RoutingProtocolIpv4)`

SetIpv4 sets Ipv4 field to given value.

### HasIpv4

`func (o *RoutingProtocolReadModelStatic) HasIpv4() bool`

HasIpv4 returns a boolean if a field has been set.

### GetIpv6

`func (o *RoutingProtocolReadModelStatic) GetIpv6() RoutingProtocolIpv6`

GetIpv6 returns the Ipv6 field if non-nil, zero value otherwise.

### GetIpv6Ok

`func (o *RoutingProtocolReadModelStatic) GetIpv6Ok() (*RoutingProtocolIpv6, bool)`

GetIpv6Ok returns a tuple with the Ipv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6

`func (o *RoutingProtocolReadModelStatic) SetIpv6(v RoutingProtocolIpv6)`

SetIpv6 sets Ipv6 field to given value.

### HasIpv6

`func (o *RoutingProtocolReadModelStatic) HasIpv6() bool`

HasIpv6 returns a boolean if a field has been set.

### GetChangeLog

`func (o *RoutingProtocolReadModelStatic) GetChangeLog() ChangeLog`

GetChangeLog returns the ChangeLog field if non-nil, zero value otherwise.

### GetChangeLogOk

`func (o *RoutingProtocolReadModelStatic) GetChangeLogOk() (*ChangeLog, bool)`

GetChangeLogOk returns a tuple with the ChangeLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeLog

`func (o *RoutingProtocolReadModelStatic) SetChangeLog(v ChangeLog)`

SetChangeLog sets ChangeLog field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


