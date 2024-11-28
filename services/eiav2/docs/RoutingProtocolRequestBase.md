# RoutingProtocolRequestBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tags** | Pointer to **[]string** |  | [optional] 
**Type** | **string** | Allowed values: - BGP - DIRECT - STATIC  | 
**Name** | Pointer to **string** | Name of the routing protocol instance.  | [optional] 
**Description** | Pointer to **string** | Description of the routing protocol instance  | [optional] 

## Methods

### NewRoutingProtocolRequestBase

`func NewRoutingProtocolRequestBase(type_ string, ) *RoutingProtocolRequestBase`

NewRoutingProtocolRequestBase instantiates a new RoutingProtocolRequestBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoutingProtocolRequestBaseWithDefaults

`func NewRoutingProtocolRequestBaseWithDefaults() *RoutingProtocolRequestBase`

NewRoutingProtocolRequestBaseWithDefaults instantiates a new RoutingProtocolRequestBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTags

`func (o *RoutingProtocolRequestBase) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *RoutingProtocolRequestBase) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *RoutingProtocolRequestBase) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *RoutingProtocolRequestBase) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetType

`func (o *RoutingProtocolRequestBase) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RoutingProtocolRequestBase) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RoutingProtocolRequestBase) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *RoutingProtocolRequestBase) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RoutingProtocolRequestBase) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RoutingProtocolRequestBase) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RoutingProtocolRequestBase) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *RoutingProtocolRequestBase) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RoutingProtocolRequestBase) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RoutingProtocolRequestBase) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RoutingProtocolRequestBase) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


