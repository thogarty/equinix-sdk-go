# RoutingProtocolRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tags** | Pointer to **[]string** |  | [optional] 
**Type** | **string** | Allowed values: - BGP - DIRECT - STATIC  | 
**Name** | Pointer to **string** | Name of the routing protocol instance.  | [optional] 
**Description** | Pointer to **string** | Description of the routing protocol instance  | [optional] 
**Ipv4** | Pointer to [**RoutingProtocolIpv4Request**](RoutingProtocolIpv4Request.md) |  | [optional] 
**Ipv6** | Pointer to [**RoutingProtocolIpv6Request**](RoutingProtocolIpv6Request.md) |  | [optional] 
**CustomerAsnRange** | Pointer to [**RoutingProtocolRequestBgpAllOfCustomerAsnRange**](RoutingProtocolRequestBgpAllOfCustomerAsnRange.md) |  | [optional] 
**CustomerAsn** | Pointer to **int64** | Customer Autonomous System Number  | [optional] 
**BgpAuthKey** | Pointer to **string** | BGP authentication key  | [optional] 
**ExportPolicy** | [**RoutingProtocolRequestBgpAllOfExportPolicy**](RoutingProtocolRequestBgpAllOfExportPolicy.md) |  | 

## Methods

### NewRoutingProtocolRequest

`func NewRoutingProtocolRequest(type_ string, exportPolicy RoutingProtocolRequestBgpAllOfExportPolicy, ) *RoutingProtocolRequest`

NewRoutingProtocolRequest instantiates a new RoutingProtocolRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoutingProtocolRequestWithDefaults

`func NewRoutingProtocolRequestWithDefaults() *RoutingProtocolRequest`

NewRoutingProtocolRequestWithDefaults instantiates a new RoutingProtocolRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTags

`func (o *RoutingProtocolRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *RoutingProtocolRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *RoutingProtocolRequest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *RoutingProtocolRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetType

`func (o *RoutingProtocolRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RoutingProtocolRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RoutingProtocolRequest) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *RoutingProtocolRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RoutingProtocolRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RoutingProtocolRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RoutingProtocolRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *RoutingProtocolRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RoutingProtocolRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RoutingProtocolRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RoutingProtocolRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIpv4

`func (o *RoutingProtocolRequest) GetIpv4() RoutingProtocolIpv4Request`

GetIpv4 returns the Ipv4 field if non-nil, zero value otherwise.

### GetIpv4Ok

`func (o *RoutingProtocolRequest) GetIpv4Ok() (*RoutingProtocolIpv4Request, bool)`

GetIpv4Ok returns a tuple with the Ipv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4

`func (o *RoutingProtocolRequest) SetIpv4(v RoutingProtocolIpv4Request)`

SetIpv4 sets Ipv4 field to given value.

### HasIpv4

`func (o *RoutingProtocolRequest) HasIpv4() bool`

HasIpv4 returns a boolean if a field has been set.

### GetIpv6

`func (o *RoutingProtocolRequest) GetIpv6() RoutingProtocolIpv6Request`

GetIpv6 returns the Ipv6 field if non-nil, zero value otherwise.

### GetIpv6Ok

`func (o *RoutingProtocolRequest) GetIpv6Ok() (*RoutingProtocolIpv6Request, bool)`

GetIpv6Ok returns a tuple with the Ipv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6

`func (o *RoutingProtocolRequest) SetIpv6(v RoutingProtocolIpv6Request)`

SetIpv6 sets Ipv6 field to given value.

### HasIpv6

`func (o *RoutingProtocolRequest) HasIpv6() bool`

HasIpv6 returns a boolean if a field has been set.

### GetCustomerAsnRange

`func (o *RoutingProtocolRequest) GetCustomerAsnRange() RoutingProtocolRequestBgpAllOfCustomerAsnRange`

GetCustomerAsnRange returns the CustomerAsnRange field if non-nil, zero value otherwise.

### GetCustomerAsnRangeOk

`func (o *RoutingProtocolRequest) GetCustomerAsnRangeOk() (*RoutingProtocolRequestBgpAllOfCustomerAsnRange, bool)`

GetCustomerAsnRangeOk returns a tuple with the CustomerAsnRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerAsnRange

`func (o *RoutingProtocolRequest) SetCustomerAsnRange(v RoutingProtocolRequestBgpAllOfCustomerAsnRange)`

SetCustomerAsnRange sets CustomerAsnRange field to given value.

### HasCustomerAsnRange

`func (o *RoutingProtocolRequest) HasCustomerAsnRange() bool`

HasCustomerAsnRange returns a boolean if a field has been set.

### GetCustomerAsn

`func (o *RoutingProtocolRequest) GetCustomerAsn() int64`

GetCustomerAsn returns the CustomerAsn field if non-nil, zero value otherwise.

### GetCustomerAsnOk

`func (o *RoutingProtocolRequest) GetCustomerAsnOk() (*int64, bool)`

GetCustomerAsnOk returns a tuple with the CustomerAsn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerAsn

`func (o *RoutingProtocolRequest) SetCustomerAsn(v int64)`

SetCustomerAsn sets CustomerAsn field to given value.

### HasCustomerAsn

`func (o *RoutingProtocolRequest) HasCustomerAsn() bool`

HasCustomerAsn returns a boolean if a field has been set.

### GetBgpAuthKey

`func (o *RoutingProtocolRequest) GetBgpAuthKey() string`

GetBgpAuthKey returns the BgpAuthKey field if non-nil, zero value otherwise.

### GetBgpAuthKeyOk

`func (o *RoutingProtocolRequest) GetBgpAuthKeyOk() (*string, bool)`

GetBgpAuthKeyOk returns a tuple with the BgpAuthKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgpAuthKey

`func (o *RoutingProtocolRequest) SetBgpAuthKey(v string)`

SetBgpAuthKey sets BgpAuthKey field to given value.

### HasBgpAuthKey

`func (o *RoutingProtocolRequest) HasBgpAuthKey() bool`

HasBgpAuthKey returns a boolean if a field has been set.

### GetExportPolicy

`func (o *RoutingProtocolRequest) GetExportPolicy() RoutingProtocolRequestBgpAllOfExportPolicy`

GetExportPolicy returns the ExportPolicy field if non-nil, zero value otherwise.

### GetExportPolicyOk

`func (o *RoutingProtocolRequest) GetExportPolicyOk() (*RoutingProtocolRequestBgpAllOfExportPolicy, bool)`

GetExportPolicyOk returns a tuple with the ExportPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportPolicy

`func (o *RoutingProtocolRequest) SetExportPolicy(v RoutingProtocolRequestBgpAllOfExportPolicy)`

SetExportPolicy sets ExportPolicy field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


