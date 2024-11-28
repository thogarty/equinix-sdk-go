# SearchResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pagination** | [**Pagination**](Pagination.md) |  | 
**Sort** | [**[]Sort**](Sort.md) |  | 
**Data** | [**[]ServiceReadModel**](ServiceReadModel.md) |  | 

## Methods

### NewSearchResponse

`func NewSearchResponse(pagination Pagination, sort []Sort, data []ServiceReadModel, ) *SearchResponse`

NewSearchResponse instantiates a new SearchResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchResponseWithDefaults

`func NewSearchResponseWithDefaults() *SearchResponse`

NewSearchResponseWithDefaults instantiates a new SearchResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPagination

`func (o *SearchResponse) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *SearchResponse) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *SearchResponse) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.


### GetSort

`func (o *SearchResponse) GetSort() []Sort`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *SearchResponse) GetSortOk() (*[]Sort, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *SearchResponse) SetSort(v []Sort)`

SetSort sets Sort field to given value.


### GetData

`func (o *SearchResponse) GetData() []ServiceReadModel`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SearchResponse) GetDataOk() (*[]ServiceReadModel, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SearchResponse) SetData(v []ServiceReadModel)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


