# ServiceCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tags** | Pointer to **[]string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Type** | [**ServiceType**](ServiceType.md) |  | 
**Connections** | **[]string** | Collection of service connections uuids | 
**RoutingProtocol** | [**RoutingProtocolRequest**](RoutingProtocolRequest.md) |  | 
**Order** | Pointer to [**ServiceOrderRequest**](ServiceOrderRequest.md) |  | [optional] 

## Methods

### NewServiceCreateRequest

`func NewServiceCreateRequest(type_ ServiceType, connections []string, routingProtocol RoutingProtocolRequest, ) *ServiceCreateRequest`

NewServiceCreateRequest instantiates a new ServiceCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceCreateRequestWithDefaults

`func NewServiceCreateRequestWithDefaults() *ServiceCreateRequest`

NewServiceCreateRequestWithDefaults instantiates a new ServiceCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTags

`func (o *ServiceCreateRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ServiceCreateRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ServiceCreateRequest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ServiceCreateRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetName

`func (o *ServiceCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServiceCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServiceCreateRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ServiceCreateRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ServiceCreateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ServiceCreateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ServiceCreateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ServiceCreateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetType

`func (o *ServiceCreateRequest) GetType() ServiceType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ServiceCreateRequest) GetTypeOk() (*ServiceType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ServiceCreateRequest) SetType(v ServiceType)`

SetType sets Type field to given value.


### GetConnections

`func (o *ServiceCreateRequest) GetConnections() []string`

GetConnections returns the Connections field if non-nil, zero value otherwise.

### GetConnectionsOk

`func (o *ServiceCreateRequest) GetConnectionsOk() (*[]string, bool)`

GetConnectionsOk returns a tuple with the Connections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnections

`func (o *ServiceCreateRequest) SetConnections(v []string)`

SetConnections sets Connections field to given value.


### GetRoutingProtocol

`func (o *ServiceCreateRequest) GetRoutingProtocol() RoutingProtocolRequest`

GetRoutingProtocol returns the RoutingProtocol field if non-nil, zero value otherwise.

### GetRoutingProtocolOk

`func (o *ServiceCreateRequest) GetRoutingProtocolOk() (*RoutingProtocolRequest, bool)`

GetRoutingProtocolOk returns a tuple with the RoutingProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingProtocol

`func (o *ServiceCreateRequest) SetRoutingProtocol(v RoutingProtocolRequest)`

SetRoutingProtocol sets RoutingProtocol field to given value.


### GetOrder

`func (o *ServiceCreateRequest) GetOrder() ServiceOrderRequest`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *ServiceCreateRequest) GetOrderOk() (*ServiceOrderRequest, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *ServiceCreateRequest) SetOrder(v ServiceOrderRequest)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *ServiceCreateRequest) HasOrder() bool`

HasOrder returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


