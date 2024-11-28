# AndQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Property** | **string** |  | 
**Operator** | [**AndQueryOperator**](AndQueryOperator.md) |  | 
**Values** | **[]string** |  | 

## Methods

### NewAndQuery

`func NewAndQuery(property string, operator AndQueryOperator, values []string, ) *AndQuery`

NewAndQuery instantiates a new AndQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAndQueryWithDefaults

`func NewAndQueryWithDefaults() *AndQuery`

NewAndQueryWithDefaults instantiates a new AndQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProperty

`func (o *AndQuery) GetProperty() string`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *AndQuery) GetPropertyOk() (*string, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *AndQuery) SetProperty(v string)`

SetProperty sets Property field to given value.


### GetOperator

`func (o *AndQuery) GetOperator() AndQueryOperator`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *AndQuery) GetOperatorOk() (*AndQueryOperator, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *AndQuery) SetOperator(v AndQueryOperator)`

SetOperator sets Operator field to given value.


### GetValues

`func (o *AndQuery) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *AndQuery) GetValuesOk() (*[]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *AndQuery) SetValues(v []string)`

SetValues sets Values field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


