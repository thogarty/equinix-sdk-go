# IbxPage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pagination** | [**Pagination**](Pagination.md) |  | 
**Data** | [**[]Ibx**](Ibx.md) |  | 

## Methods

### NewIbxPage

`func NewIbxPage(pagination Pagination, data []Ibx, ) *IbxPage`

NewIbxPage instantiates a new IbxPage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIbxPageWithDefaults

`func NewIbxPageWithDefaults() *IbxPage`

NewIbxPageWithDefaults instantiates a new IbxPage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPagination

`func (o *IbxPage) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *IbxPage) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *IbxPage) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.


### GetData

`func (o *IbxPage) GetData() []Ibx`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *IbxPage) GetDataOk() (*[]Ibx, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *IbxPage) SetData(v []Ibx)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


