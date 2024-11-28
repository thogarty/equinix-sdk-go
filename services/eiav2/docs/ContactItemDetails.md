# ContactItemDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**ContactItemDetailsType**](ContactItemDetailsType.md) |  | [optional] 
**Value** | Pointer to **string** |  | [optional] 
**Notes** | Pointer to **string** | Contact additional info  | [optional] 

## Methods

### NewContactItemDetails

`func NewContactItemDetails() *ContactItemDetails`

NewContactItemDetails instantiates a new ContactItemDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContactItemDetailsWithDefaults

`func NewContactItemDetailsWithDefaults() *ContactItemDetails`

NewContactItemDetailsWithDefaults instantiates a new ContactItemDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ContactItemDetails) GetType() ContactItemDetailsType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ContactItemDetails) GetTypeOk() (*ContactItemDetailsType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ContactItemDetails) SetType(v ContactItemDetailsType)`

SetType sets Type field to given value.

### HasType

`func (o *ContactItemDetails) HasType() bool`

HasType returns a boolean if a field has been set.

### GetValue

`func (o *ContactItemDetails) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ContactItemDetails) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ContactItemDetails) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *ContactItemDetails) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetNotes

`func (o *ContactItemDetails) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *ContactItemDetails) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *ContactItemDetails) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *ContactItemDetails) HasNotes() bool`

HasNotes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


