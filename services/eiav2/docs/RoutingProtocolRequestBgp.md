# RoutingProtocolRequestBgp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tags** | Pointer to **[]string** |  | [optional] 
**Type** | **string** | Allowed values: - BGP - DIRECT - STATIC  | 
**Name** | Pointer to **string** | Name of the routing protocol instance.  | [optional] 
**Description** | Pointer to **string** | Description of the routing protocol instance  | [optional] 
**CustomerAsnRange** | Pointer to [**RoutingProtocolRequestBgpAllOfCustomerAsnRange**](RoutingProtocolRequestBgpAllOfCustomerAsnRange.md) |  | [optional] 
**CustomerAsn** | Pointer to **int64** | Customer Autonomous System Number  | [optional] 
**BgpAuthKey** | Pointer to **string** | BGP authentication key  | [optional] 
**ExportPolicy** | [**RoutingProtocolRequestBgpAllOfExportPolicy**](RoutingProtocolRequestBgpAllOfExportPolicy.md) |  | 
**Ipv4** | Pointer to [**RoutingProtocolIpv4Request**](RoutingProtocolIpv4Request.md) |  | [optional] 
**Ipv6** | Pointer to [**RoutingProtocolIpv6Request**](RoutingProtocolIpv6Request.md) |  | [optional] 

## Methods

### NewRoutingProtocolRequestBgp

`func NewRoutingProtocolRequestBgp(type_ string, exportPolicy RoutingProtocolRequestBgpAllOfExportPolicy, ) *RoutingProtocolRequestBgp`

NewRoutingProtocolRequestBgp instantiates a new RoutingProtocolRequestBgp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoutingProtocolRequestBgpWithDefaults

`func NewRoutingProtocolRequestBgpWithDefaults() *RoutingProtocolRequestBgp`

NewRoutingProtocolRequestBgpWithDefaults instantiates a new RoutingProtocolRequestBgp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTags

`func (o *RoutingProtocolRequestBgp) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *RoutingProtocolRequestBgp) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *RoutingProtocolRequestBgp) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *RoutingProtocolRequestBgp) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetType

`func (o *RoutingProtocolRequestBgp) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RoutingProtocolRequestBgp) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RoutingProtocolRequestBgp) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *RoutingProtocolRequestBgp) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RoutingProtocolRequestBgp) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RoutingProtocolRequestBgp) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RoutingProtocolRequestBgp) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *RoutingProtocolRequestBgp) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RoutingProtocolRequestBgp) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RoutingProtocolRequestBgp) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RoutingProtocolRequestBgp) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCustomerAsnRange

`func (o *RoutingProtocolRequestBgp) GetCustomerAsnRange() RoutingProtocolRequestBgpAllOfCustomerAsnRange`

GetCustomerAsnRange returns the CustomerAsnRange field if non-nil, zero value otherwise.

### GetCustomerAsnRangeOk

`func (o *RoutingProtocolRequestBgp) GetCustomerAsnRangeOk() (*RoutingProtocolRequestBgpAllOfCustomerAsnRange, bool)`

GetCustomerAsnRangeOk returns a tuple with the CustomerAsnRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerAsnRange

`func (o *RoutingProtocolRequestBgp) SetCustomerAsnRange(v RoutingProtocolRequestBgpAllOfCustomerAsnRange)`

SetCustomerAsnRange sets CustomerAsnRange field to given value.

### HasCustomerAsnRange

`func (o *RoutingProtocolRequestBgp) HasCustomerAsnRange() bool`

HasCustomerAsnRange returns a boolean if a field has been set.

### GetCustomerAsn

`func (o *RoutingProtocolRequestBgp) GetCustomerAsn() int64`

GetCustomerAsn returns the CustomerAsn field if non-nil, zero value otherwise.

### GetCustomerAsnOk

`func (o *RoutingProtocolRequestBgp) GetCustomerAsnOk() (*int64, bool)`

GetCustomerAsnOk returns a tuple with the CustomerAsn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerAsn

`func (o *RoutingProtocolRequestBgp) SetCustomerAsn(v int64)`

SetCustomerAsn sets CustomerAsn field to given value.

### HasCustomerAsn

`func (o *RoutingProtocolRequestBgp) HasCustomerAsn() bool`

HasCustomerAsn returns a boolean if a field has been set.

### GetBgpAuthKey

`func (o *RoutingProtocolRequestBgp) GetBgpAuthKey() string`

GetBgpAuthKey returns the BgpAuthKey field if non-nil, zero value otherwise.

### GetBgpAuthKeyOk

`func (o *RoutingProtocolRequestBgp) GetBgpAuthKeyOk() (*string, bool)`

GetBgpAuthKeyOk returns a tuple with the BgpAuthKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgpAuthKey

`func (o *RoutingProtocolRequestBgp) SetBgpAuthKey(v string)`

SetBgpAuthKey sets BgpAuthKey field to given value.

### HasBgpAuthKey

`func (o *RoutingProtocolRequestBgp) HasBgpAuthKey() bool`

HasBgpAuthKey returns a boolean if a field has been set.

### GetExportPolicy

`func (o *RoutingProtocolRequestBgp) GetExportPolicy() RoutingProtocolRequestBgpAllOfExportPolicy`

GetExportPolicy returns the ExportPolicy field if non-nil, zero value otherwise.

### GetExportPolicyOk

`func (o *RoutingProtocolRequestBgp) GetExportPolicyOk() (*RoutingProtocolRequestBgpAllOfExportPolicy, bool)`

GetExportPolicyOk returns a tuple with the ExportPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportPolicy

`func (o *RoutingProtocolRequestBgp) SetExportPolicy(v RoutingProtocolRequestBgpAllOfExportPolicy)`

SetExportPolicy sets ExportPolicy field to given value.


### GetIpv4

`func (o *RoutingProtocolRequestBgp) GetIpv4() RoutingProtocolIpv4Request`

GetIpv4 returns the Ipv4 field if non-nil, zero value otherwise.

### GetIpv4Ok

`func (o *RoutingProtocolRequestBgp) GetIpv4Ok() (*RoutingProtocolIpv4Request, bool)`

GetIpv4Ok returns a tuple with the Ipv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4

`func (o *RoutingProtocolRequestBgp) SetIpv4(v RoutingProtocolIpv4Request)`

SetIpv4 sets Ipv4 field to given value.

### HasIpv4

`func (o *RoutingProtocolRequestBgp) HasIpv4() bool`

HasIpv4 returns a boolean if a field has been set.

### GetIpv6

`func (o *RoutingProtocolRequestBgp) GetIpv6() RoutingProtocolIpv6Request`

GetIpv6 returns the Ipv6 field if non-nil, zero value otherwise.

### GetIpv6Ok

`func (o *RoutingProtocolRequestBgp) GetIpv6Ok() (*RoutingProtocolIpv6Request, bool)`

GetIpv6Ok returns a tuple with the Ipv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6

`func (o *RoutingProtocolRequestBgp) SetIpv6(v RoutingProtocolIpv6Request)`

SetIpv6 sets Ipv6 field to given value.

### HasIpv6

`func (o *RoutingProtocolRequestBgp) HasIpv6() bool`

HasIpv6 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


