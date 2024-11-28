# RoutingProtocolRequestDirect

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tags** | Pointer to **[]string** |  | [optional] 
**Type** | **string** | Allowed values: - BGP - DIRECT - STATIC  | 
**Name** | Pointer to **string** | Name of the routing protocol instance.  | [optional] 
**Description** | Pointer to **string** | Description of the routing protocol instance  | [optional] 
**Ipv4** | Pointer to [**DirectRoutingProtocolIpv4Request**](DirectRoutingProtocolIpv4Request.md) |  | [optional] 
**Ipv6** | Pointer to [**DirectRoutingProtocolIpv6Request**](DirectRoutingProtocolIpv6Request.md) |  | [optional] 

## Methods

### NewRoutingProtocolRequestDirect

`func NewRoutingProtocolRequestDirect(type_ string, ) *RoutingProtocolRequestDirect`

NewRoutingProtocolRequestDirect instantiates a new RoutingProtocolRequestDirect object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoutingProtocolRequestDirectWithDefaults

`func NewRoutingProtocolRequestDirectWithDefaults() *RoutingProtocolRequestDirect`

NewRoutingProtocolRequestDirectWithDefaults instantiates a new RoutingProtocolRequestDirect object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTags

`func (o *RoutingProtocolRequestDirect) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *RoutingProtocolRequestDirect) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *RoutingProtocolRequestDirect) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *RoutingProtocolRequestDirect) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetType

`func (o *RoutingProtocolRequestDirect) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RoutingProtocolRequestDirect) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RoutingProtocolRequestDirect) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *RoutingProtocolRequestDirect) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RoutingProtocolRequestDirect) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RoutingProtocolRequestDirect) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RoutingProtocolRequestDirect) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *RoutingProtocolRequestDirect) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RoutingProtocolRequestDirect) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RoutingProtocolRequestDirect) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RoutingProtocolRequestDirect) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIpv4

`func (o *RoutingProtocolRequestDirect) GetIpv4() DirectRoutingProtocolIpv4Request`

GetIpv4 returns the Ipv4 field if non-nil, zero value otherwise.

### GetIpv4Ok

`func (o *RoutingProtocolRequestDirect) GetIpv4Ok() (*DirectRoutingProtocolIpv4Request, bool)`

GetIpv4Ok returns a tuple with the Ipv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4

`func (o *RoutingProtocolRequestDirect) SetIpv4(v DirectRoutingProtocolIpv4Request)`

SetIpv4 sets Ipv4 field to given value.

### HasIpv4

`func (o *RoutingProtocolRequestDirect) HasIpv4() bool`

HasIpv4 returns a boolean if a field has been set.

### GetIpv6

`func (o *RoutingProtocolRequestDirect) GetIpv6() DirectRoutingProtocolIpv6Request`

GetIpv6 returns the Ipv6 field if non-nil, zero value otherwise.

### GetIpv6Ok

`func (o *RoutingProtocolRequestDirect) GetIpv6Ok() (*DirectRoutingProtocolIpv6Request, bool)`

GetIpv6Ok returns a tuple with the Ipv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6

`func (o *RoutingProtocolRequestDirect) SetIpv6(v DirectRoutingProtocolIpv6Request)`

SetIpv6 sets Ipv6 field to given value.

### HasIpv6

`func (o *RoutingProtocolRequestDirect) HasIpv6() bool`

HasIpv6 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


