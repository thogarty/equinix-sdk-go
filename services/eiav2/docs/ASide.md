# ASide

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Service** | [**ASideService**](ASideService.md) |  | 

## Methods

### NewASide

`func NewASide(type_ string, service ASideService, ) *ASide`

NewASide instantiates a new ASide object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewASideWithDefaults

`func NewASideWithDefaults() *ASide`

NewASideWithDefaults instantiates a new ASide object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ASide) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ASide) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ASide) SetType(v string)`

SetType sets Type field to given value.


### GetService

`func (o *ASide) GetService() ASideService`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *ASide) GetServiceOk() (*ASideService, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *ASide) SetService(v ASideService)`

SetService sets Service field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


