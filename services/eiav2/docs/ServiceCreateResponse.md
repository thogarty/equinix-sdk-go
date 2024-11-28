# ServiceCreateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Type** | [**ServiceType**](ServiceType.md) |  | 
**Connections** | [**[]Connection**](Connection.md) | Collection of service connections uuids | 
**RoutingProtocol** | [**RoutingProtocolRequest**](RoutingProtocolRequest.md) |  | 
**Order** | Pointer to [**ServiceOrderRequest**](ServiceOrderRequest.md) |  | [optional] 
**Uuid** | **string** | Service identifier | 
**State** | [**ServiceState**](ServiceState.md) |  | 

## Methods

### NewServiceCreateResponse

`func NewServiceCreateResponse(type_ ServiceType, connections []Connection, routingProtocol RoutingProtocolRequest, uuid string, state ServiceState, ) *ServiceCreateResponse`

NewServiceCreateResponse instantiates a new ServiceCreateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceCreateResponseWithDefaults

`func NewServiceCreateResponseWithDefaults() *ServiceCreateResponse`

NewServiceCreateResponseWithDefaults instantiates a new ServiceCreateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ServiceCreateResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServiceCreateResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServiceCreateResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ServiceCreateResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ServiceCreateResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ServiceCreateResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ServiceCreateResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ServiceCreateResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetType

`func (o *ServiceCreateResponse) GetType() ServiceType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ServiceCreateResponse) GetTypeOk() (*ServiceType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ServiceCreateResponse) SetType(v ServiceType)`

SetType sets Type field to given value.


### GetConnections

`func (o *ServiceCreateResponse) GetConnections() []Connection`

GetConnections returns the Connections field if non-nil, zero value otherwise.

### GetConnectionsOk

`func (o *ServiceCreateResponse) GetConnectionsOk() (*[]Connection, bool)`

GetConnectionsOk returns a tuple with the Connections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnections

`func (o *ServiceCreateResponse) SetConnections(v []Connection)`

SetConnections sets Connections field to given value.


### GetRoutingProtocol

`func (o *ServiceCreateResponse) GetRoutingProtocol() RoutingProtocolRequest`

GetRoutingProtocol returns the RoutingProtocol field if non-nil, zero value otherwise.

### GetRoutingProtocolOk

`func (o *ServiceCreateResponse) GetRoutingProtocolOk() (*RoutingProtocolRequest, bool)`

GetRoutingProtocolOk returns a tuple with the RoutingProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingProtocol

`func (o *ServiceCreateResponse) SetRoutingProtocol(v RoutingProtocolRequest)`

SetRoutingProtocol sets RoutingProtocol field to given value.


### GetOrder

`func (o *ServiceCreateResponse) GetOrder() ServiceOrderRequest`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *ServiceCreateResponse) GetOrderOk() (*ServiceOrderRequest, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *ServiceCreateResponse) SetOrder(v ServiceOrderRequest)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *ServiceCreateResponse) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetUuid

`func (o *ServiceCreateResponse) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *ServiceCreateResponse) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *ServiceCreateResponse) SetUuid(v string)`

SetUuid sets Uuid field to given value.


### GetState

`func (o *ServiceCreateResponse) GetState() ServiceState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ServiceCreateResponse) GetStateOk() (*ServiceState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ServiceCreateResponse) SetState(v ServiceState)`

SetState sets State field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


