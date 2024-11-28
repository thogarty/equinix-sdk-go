# ConnectionReadModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | **string** |  | 
**Href** | **string** | Connection URI | 
**Type** | Pointer to [**ConnectionType**](ConnectionType.md) |  | [optional] 
**ASide** | [**ASide**](ASide.md) |  | 

## Methods

### NewConnectionReadModel

`func NewConnectionReadModel(uuid string, href string, aSide ASide, ) *ConnectionReadModel`

NewConnectionReadModel instantiates a new ConnectionReadModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectionReadModelWithDefaults

`func NewConnectionReadModelWithDefaults() *ConnectionReadModel`

NewConnectionReadModelWithDefaults instantiates a new ConnectionReadModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *ConnectionReadModel) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *ConnectionReadModel) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *ConnectionReadModel) SetUuid(v string)`

SetUuid sets Uuid field to given value.


### GetHref

`func (o *ConnectionReadModel) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *ConnectionReadModel) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *ConnectionReadModel) SetHref(v string)`

SetHref sets Href field to given value.


### GetType

`func (o *ConnectionReadModel) GetType() ConnectionType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ConnectionReadModel) GetTypeOk() (*ConnectionType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ConnectionReadModel) SetType(v ConnectionType)`

SetType sets Type field to given value.

### HasType

`func (o *ConnectionReadModel) HasType() bool`

HasType returns a boolean if a field has been set.

### GetASide

`func (o *ConnectionReadModel) GetASide() ASide`

GetASide returns the ASide field if non-nil, zero value otherwise.

### GetASideOk

`func (o *ConnectionReadModel) GetASideOk() (*ASide, bool)`

GetASideOk returns a tuple with the ASide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetASide

`func (o *ConnectionReadModel) SetASide(v ASide)`

SetASide sets ASide field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


