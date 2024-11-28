# Location

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metro** | **string** |  | 
**IbxCode** | **string** | IBX data center code | 

## Methods

### NewLocation

`func NewLocation(metro string, ibxCode string, ) *Location`

NewLocation instantiates a new Location object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocationWithDefaults

`func NewLocationWithDefaults() *Location`

NewLocationWithDefaults instantiates a new Location object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetro

`func (o *Location) GetMetro() string`

GetMetro returns the Metro field if non-nil, zero value otherwise.

### GetMetroOk

`func (o *Location) GetMetroOk() (*string, bool)`

GetMetroOk returns a tuple with the Metro field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetro

`func (o *Location) SetMetro(v string)`

SetMetro sets Metro field to given value.


### GetIbxCode

`func (o *Location) GetIbxCode() string`

GetIbxCode returns the IbxCode field if non-nil, zero value otherwise.

### GetIbxCodeOk

`func (o *Location) GetIbxCodeOk() (*string, bool)`

GetIbxCodeOk returns a tuple with the IbxCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIbxCode

`func (o *Location) SetIbxCode(v string)`

SetIbxCode sets IbxCode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


