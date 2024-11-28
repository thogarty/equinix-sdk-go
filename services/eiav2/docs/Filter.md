# Filter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**And** | Pointer to [**[]AndQuery**](AndQuery.md) |  | [optional] 

## Methods

### NewFilter

`func NewFilter() *Filter`

NewFilter instantiates a new Filter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFilterWithDefaults

`func NewFilterWithDefaults() *Filter`

NewFilterWithDefaults instantiates a new Filter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnd

`func (o *Filter) GetAnd() []AndQuery`

GetAnd returns the And field if non-nil, zero value otherwise.

### GetAndOk

`func (o *Filter) GetAndOk() (*[]AndQuery, bool)`

GetAndOk returns a tuple with the And field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnd

`func (o *Filter) SetAnd(v []AndQuery)`

SetAnd sets And field to given value.

### HasAnd

`func (o *Filter) HasAnd() bool`

HasAnd returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


