# RoutingProtocolReadModel

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
**CustomerAsn** | **int64** | Customer Autonomous System Number  | 
**CustomerAsnRange** | Pointer to [**RoutingProtocolRequestBgpAllOfCustomerAsnRange**](RoutingProtocolRequestBgpAllOfCustomerAsnRange.md) |  | [optional] 
**EquinixAsn** | **int64** | Equinix Autonomous System Number  | 
**BgpAuthKey** | Pointer to **string** | BGP authentication key  | [optional] 
**ExportPolicy** | [**ExportPolicy**](ExportPolicy.md) |  | 

## Methods

### NewRoutingProtocolReadModel

`func NewRoutingProtocolReadModel(type_ string, name string, changeLog ChangeLog, customerAsn int64, equinixAsn int64, exportPolicy ExportPolicy, ) *RoutingProtocolReadModel`

NewRoutingProtocolReadModel instantiates a new RoutingProtocolReadModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoutingProtocolReadModelWithDefaults

`func NewRoutingProtocolReadModelWithDefaults() *RoutingProtocolReadModel`

NewRoutingProtocolReadModelWithDefaults instantiates a new RoutingProtocolReadModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTags

`func (o *RoutingProtocolReadModel) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *RoutingProtocolReadModel) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *RoutingProtocolReadModel) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *RoutingProtocolReadModel) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetHref

`func (o *RoutingProtocolReadModel) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *RoutingProtocolReadModel) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *RoutingProtocolReadModel) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *RoutingProtocolReadModel) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetUuid

`func (o *RoutingProtocolReadModel) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *RoutingProtocolReadModel) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *RoutingProtocolReadModel) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *RoutingProtocolReadModel) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetType

`func (o *RoutingProtocolReadModel) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RoutingProtocolReadModel) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RoutingProtocolReadModel) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *RoutingProtocolReadModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RoutingProtocolReadModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RoutingProtocolReadModel) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *RoutingProtocolReadModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RoutingProtocolReadModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RoutingProtocolReadModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RoutingProtocolReadModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIpv4

`func (o *RoutingProtocolReadModel) GetIpv4() RoutingProtocolIpv4`

GetIpv4 returns the Ipv4 field if non-nil, zero value otherwise.

### GetIpv4Ok

`func (o *RoutingProtocolReadModel) GetIpv4Ok() (*RoutingProtocolIpv4, bool)`

GetIpv4Ok returns a tuple with the Ipv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4

`func (o *RoutingProtocolReadModel) SetIpv4(v RoutingProtocolIpv4)`

SetIpv4 sets Ipv4 field to given value.

### HasIpv4

`func (o *RoutingProtocolReadModel) HasIpv4() bool`

HasIpv4 returns a boolean if a field has been set.

### GetIpv6

`func (o *RoutingProtocolReadModel) GetIpv6() RoutingProtocolIpv6`

GetIpv6 returns the Ipv6 field if non-nil, zero value otherwise.

### GetIpv6Ok

`func (o *RoutingProtocolReadModel) GetIpv6Ok() (*RoutingProtocolIpv6, bool)`

GetIpv6Ok returns a tuple with the Ipv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6

`func (o *RoutingProtocolReadModel) SetIpv6(v RoutingProtocolIpv6)`

SetIpv6 sets Ipv6 field to given value.

### HasIpv6

`func (o *RoutingProtocolReadModel) HasIpv6() bool`

HasIpv6 returns a boolean if a field has been set.

### GetChangeLog

`func (o *RoutingProtocolReadModel) GetChangeLog() ChangeLog`

GetChangeLog returns the ChangeLog field if non-nil, zero value otherwise.

### GetChangeLogOk

`func (o *RoutingProtocolReadModel) GetChangeLogOk() (*ChangeLog, bool)`

GetChangeLogOk returns a tuple with the ChangeLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeLog

`func (o *RoutingProtocolReadModel) SetChangeLog(v ChangeLog)`

SetChangeLog sets ChangeLog field to given value.


### GetCustomerAsn

`func (o *RoutingProtocolReadModel) GetCustomerAsn() int64`

GetCustomerAsn returns the CustomerAsn field if non-nil, zero value otherwise.

### GetCustomerAsnOk

`func (o *RoutingProtocolReadModel) GetCustomerAsnOk() (*int64, bool)`

GetCustomerAsnOk returns a tuple with the CustomerAsn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerAsn

`func (o *RoutingProtocolReadModel) SetCustomerAsn(v int64)`

SetCustomerAsn sets CustomerAsn field to given value.


### GetCustomerAsnRange

`func (o *RoutingProtocolReadModel) GetCustomerAsnRange() RoutingProtocolRequestBgpAllOfCustomerAsnRange`

GetCustomerAsnRange returns the CustomerAsnRange field if non-nil, zero value otherwise.

### GetCustomerAsnRangeOk

`func (o *RoutingProtocolReadModel) GetCustomerAsnRangeOk() (*RoutingProtocolRequestBgpAllOfCustomerAsnRange, bool)`

GetCustomerAsnRangeOk returns a tuple with the CustomerAsnRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerAsnRange

`func (o *RoutingProtocolReadModel) SetCustomerAsnRange(v RoutingProtocolRequestBgpAllOfCustomerAsnRange)`

SetCustomerAsnRange sets CustomerAsnRange field to given value.

### HasCustomerAsnRange

`func (o *RoutingProtocolReadModel) HasCustomerAsnRange() bool`

HasCustomerAsnRange returns a boolean if a field has been set.

### GetEquinixAsn

`func (o *RoutingProtocolReadModel) GetEquinixAsn() int64`

GetEquinixAsn returns the EquinixAsn field if non-nil, zero value otherwise.

### GetEquinixAsnOk

`func (o *RoutingProtocolReadModel) GetEquinixAsnOk() (*int64, bool)`

GetEquinixAsnOk returns a tuple with the EquinixAsn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquinixAsn

`func (o *RoutingProtocolReadModel) SetEquinixAsn(v int64)`

SetEquinixAsn sets EquinixAsn field to given value.


### GetBgpAuthKey

`func (o *RoutingProtocolReadModel) GetBgpAuthKey() string`

GetBgpAuthKey returns the BgpAuthKey field if non-nil, zero value otherwise.

### GetBgpAuthKeyOk

`func (o *RoutingProtocolReadModel) GetBgpAuthKeyOk() (*string, bool)`

GetBgpAuthKeyOk returns a tuple with the BgpAuthKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgpAuthKey

`func (o *RoutingProtocolReadModel) SetBgpAuthKey(v string)`

SetBgpAuthKey sets BgpAuthKey field to given value.

### HasBgpAuthKey

`func (o *RoutingProtocolReadModel) HasBgpAuthKey() bool`

HasBgpAuthKey returns a boolean if a field has been set.

### GetExportPolicy

`func (o *RoutingProtocolReadModel) GetExportPolicy() ExportPolicy`

GetExportPolicy returns the ExportPolicy field if non-nil, zero value otherwise.

### GetExportPolicyOk

`func (o *RoutingProtocolReadModel) GetExportPolicyOk() (*ExportPolicy, bool)`

GetExportPolicyOk returns a tuple with the ExportPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportPolicy

`func (o *RoutingProtocolReadModel) SetExportPolicy(v ExportPolicy)`

SetExportPolicy sets ExportPolicy field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


