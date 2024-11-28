# ServiceReadModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tags** | Pointer to **[]string** |  | [optional] 
**Href** | **string** | Internet access URI | 
**Uuid** | **string** | Service identifier | 
**Type** | [**ServiceType**](ServiceType.md) |  | 
**Name** | **string** | Name of the service This name should only contain spaces, the characters \&quot;_\&quot;, \&quot;-\&quot;, letters or numbers. The name cannot start or end with a \&quot;-\&quot;. This name can have only maximum of 24 characters  | 
**Description** | Pointer to **string** | Description of the service | [optional] 
**Bandwidth** | **int64** | Bandwidth of the service in Mbps | 
**Billing** | [**ServiceBilling**](ServiceBilling.md) |  | 
**State** | [**ServiceState**](ServiceState.md) |  | 
**Change** | [**Change**](Change.md) |  | 
**ChangeLog** | [**ChangeLog**](ChangeLog.md) |  | 
**Connections** | [**[]ConnectionReadModel**](ConnectionReadModel.md) |  | 
**RoutingProtocol** | [**RoutingProtocolReadModel**](RoutingProtocolReadModel.md) |  | 
**Locations** | [**[]Location**](Location.md) |  | 
**Account** | [**Account**](Account.md) |  | 
**Project** | [**ProjectReadModel**](ProjectReadModel.md) |  | 
**Order** | [**OrderReadModel**](OrderReadModel.md) |  | 

## Methods

### NewServiceReadModel

`func NewServiceReadModel(href string, uuid string, type_ ServiceType, name string, bandwidth int64, billing ServiceBilling, state ServiceState, change Change, changeLog ChangeLog, connections []ConnectionReadModel, routingProtocol RoutingProtocolReadModel, locations []Location, account Account, project ProjectReadModel, order OrderReadModel, ) *ServiceReadModel`

NewServiceReadModel instantiates a new ServiceReadModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceReadModelWithDefaults

`func NewServiceReadModelWithDefaults() *ServiceReadModel`

NewServiceReadModelWithDefaults instantiates a new ServiceReadModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTags

`func (o *ServiceReadModel) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ServiceReadModel) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ServiceReadModel) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ServiceReadModel) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetHref

`func (o *ServiceReadModel) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *ServiceReadModel) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *ServiceReadModel) SetHref(v string)`

SetHref sets Href field to given value.


### GetUuid

`func (o *ServiceReadModel) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *ServiceReadModel) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *ServiceReadModel) SetUuid(v string)`

SetUuid sets Uuid field to given value.


### GetType

`func (o *ServiceReadModel) GetType() ServiceType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ServiceReadModel) GetTypeOk() (*ServiceType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ServiceReadModel) SetType(v ServiceType)`

SetType sets Type field to given value.


### GetName

`func (o *ServiceReadModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServiceReadModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServiceReadModel) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ServiceReadModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ServiceReadModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ServiceReadModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ServiceReadModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetBandwidth

`func (o *ServiceReadModel) GetBandwidth() int64`

GetBandwidth returns the Bandwidth field if non-nil, zero value otherwise.

### GetBandwidthOk

`func (o *ServiceReadModel) GetBandwidthOk() (*int64, bool)`

GetBandwidthOk returns a tuple with the Bandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidth

`func (o *ServiceReadModel) SetBandwidth(v int64)`

SetBandwidth sets Bandwidth field to given value.


### GetBilling

`func (o *ServiceReadModel) GetBilling() ServiceBilling`

GetBilling returns the Billing field if non-nil, zero value otherwise.

### GetBillingOk

`func (o *ServiceReadModel) GetBillingOk() (*ServiceBilling, bool)`

GetBillingOk returns a tuple with the Billing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBilling

`func (o *ServiceReadModel) SetBilling(v ServiceBilling)`

SetBilling sets Billing field to given value.


### GetState

`func (o *ServiceReadModel) GetState() ServiceState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ServiceReadModel) GetStateOk() (*ServiceState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ServiceReadModel) SetState(v ServiceState)`

SetState sets State field to given value.


### GetChange

`func (o *ServiceReadModel) GetChange() Change`

GetChange returns the Change field if non-nil, zero value otherwise.

### GetChangeOk

`func (o *ServiceReadModel) GetChangeOk() (*Change, bool)`

GetChangeOk returns a tuple with the Change field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChange

`func (o *ServiceReadModel) SetChange(v Change)`

SetChange sets Change field to given value.


### GetChangeLog

`func (o *ServiceReadModel) GetChangeLog() ChangeLog`

GetChangeLog returns the ChangeLog field if non-nil, zero value otherwise.

### GetChangeLogOk

`func (o *ServiceReadModel) GetChangeLogOk() (*ChangeLog, bool)`

GetChangeLogOk returns a tuple with the ChangeLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeLog

`func (o *ServiceReadModel) SetChangeLog(v ChangeLog)`

SetChangeLog sets ChangeLog field to given value.


### GetConnections

`func (o *ServiceReadModel) GetConnections() []ConnectionReadModel`

GetConnections returns the Connections field if non-nil, zero value otherwise.

### GetConnectionsOk

`func (o *ServiceReadModel) GetConnectionsOk() (*[]ConnectionReadModel, bool)`

GetConnectionsOk returns a tuple with the Connections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnections

`func (o *ServiceReadModel) SetConnections(v []ConnectionReadModel)`

SetConnections sets Connections field to given value.


### GetRoutingProtocol

`func (o *ServiceReadModel) GetRoutingProtocol() RoutingProtocolReadModel`

GetRoutingProtocol returns the RoutingProtocol field if non-nil, zero value otherwise.

### GetRoutingProtocolOk

`func (o *ServiceReadModel) GetRoutingProtocolOk() (*RoutingProtocolReadModel, bool)`

GetRoutingProtocolOk returns a tuple with the RoutingProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingProtocol

`func (o *ServiceReadModel) SetRoutingProtocol(v RoutingProtocolReadModel)`

SetRoutingProtocol sets RoutingProtocol field to given value.


### GetLocations

`func (o *ServiceReadModel) GetLocations() []Location`

GetLocations returns the Locations field if non-nil, zero value otherwise.

### GetLocationsOk

`func (o *ServiceReadModel) GetLocationsOk() (*[]Location, bool)`

GetLocationsOk returns a tuple with the Locations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocations

`func (o *ServiceReadModel) SetLocations(v []Location)`

SetLocations sets Locations field to given value.


### GetAccount

`func (o *ServiceReadModel) GetAccount() Account`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *ServiceReadModel) GetAccountOk() (*Account, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *ServiceReadModel) SetAccount(v Account)`

SetAccount sets Account field to given value.


### GetProject

`func (o *ServiceReadModel) GetProject() ProjectReadModel`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *ServiceReadModel) GetProjectOk() (*ProjectReadModel, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *ServiceReadModel) SetProject(v ProjectReadModel)`

SetProject sets Project field to given value.


### GetOrder

`func (o *ServiceReadModel) GetOrder() OrderReadModel`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *ServiceReadModel) GetOrderOk() (*OrderReadModel, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *ServiceReadModel) SetOrder(v OrderReadModel)`

SetOrder sets Order field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


