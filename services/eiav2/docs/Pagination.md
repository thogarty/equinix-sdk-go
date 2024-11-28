# Pagination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Offset** | **int64** |  | 
**Limit** | **int64** |  | 
**Total** | **int64** |  | 
**Next** | Pointer to **string** | Next page link | [optional] 
**Previous** | Pointer to **string** | Previous page link | [optional] 

## Methods

### NewPagination

`func NewPagination(offset int64, limit int64, total int64, ) *Pagination`

NewPagination instantiates a new Pagination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginationWithDefaults

`func NewPaginationWithDefaults() *Pagination`

NewPaginationWithDefaults instantiates a new Pagination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOffset

`func (o *Pagination) GetOffset() int64`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *Pagination) GetOffsetOk() (*int64, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *Pagination) SetOffset(v int64)`

SetOffset sets Offset field to given value.


### GetLimit

`func (o *Pagination) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *Pagination) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *Pagination) SetLimit(v int64)`

SetLimit sets Limit field to given value.


### GetTotal

`func (o *Pagination) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *Pagination) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *Pagination) SetTotal(v int64)`

SetTotal sets Total field to given value.


### GetNext

`func (o *Pagination) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *Pagination) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *Pagination) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *Pagination) HasNext() bool`

HasNext returns a boolean if a field has been set.

### GetPrevious

`func (o *Pagination) GetPrevious() string`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *Pagination) GetPreviousOk() (*string, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *Pagination) SetPrevious(v string)`

SetPrevious sets Previous field to given value.

### HasPrevious

`func (o *Pagination) HasPrevious() bool`

HasPrevious returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


