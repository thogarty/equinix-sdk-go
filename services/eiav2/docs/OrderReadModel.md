# OrderReadModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | **string** | Order URI | 
**Uuid** | **string** | Order identifier | 
**Type** | [**OrderReadModelType**](OrderReadModelType.md) |  | 
**Status** | [**OrderReadModelStatus**](OrderReadModelStatus.md) |  | 
**ChangeLog** | [**ChangeLog**](ChangeLog.md) |  | 
**Number** | **string** | Order number | 
**PurchaseOrder** | Pointer to [**PurchaseOrderReadModel**](PurchaseOrderReadModel.md) |  | [optional] 
**Contacts** | Pointer to [**[]ContactItem**](ContactItem.md) |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 

## Methods

### NewOrderReadModel

`func NewOrderReadModel(href string, uuid string, type_ OrderReadModelType, status OrderReadModelStatus, changeLog ChangeLog, number string, ) *OrderReadModel`

NewOrderReadModel instantiates a new OrderReadModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderReadModelWithDefaults

`func NewOrderReadModelWithDefaults() *OrderReadModel`

NewOrderReadModelWithDefaults instantiates a new OrderReadModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *OrderReadModel) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *OrderReadModel) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *OrderReadModel) SetHref(v string)`

SetHref sets Href field to given value.


### GetUuid

`func (o *OrderReadModel) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *OrderReadModel) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *OrderReadModel) SetUuid(v string)`

SetUuid sets Uuid field to given value.


### GetType

`func (o *OrderReadModel) GetType() OrderReadModelType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OrderReadModel) GetTypeOk() (*OrderReadModelType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OrderReadModel) SetType(v OrderReadModelType)`

SetType sets Type field to given value.


### GetStatus

`func (o *OrderReadModel) GetStatus() OrderReadModelStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OrderReadModel) GetStatusOk() (*OrderReadModelStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OrderReadModel) SetStatus(v OrderReadModelStatus)`

SetStatus sets Status field to given value.


### GetChangeLog

`func (o *OrderReadModel) GetChangeLog() ChangeLog`

GetChangeLog returns the ChangeLog field if non-nil, zero value otherwise.

### GetChangeLogOk

`func (o *OrderReadModel) GetChangeLogOk() (*ChangeLog, bool)`

GetChangeLogOk returns a tuple with the ChangeLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeLog

`func (o *OrderReadModel) SetChangeLog(v ChangeLog)`

SetChangeLog sets ChangeLog field to given value.


### GetNumber

`func (o *OrderReadModel) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *OrderReadModel) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *OrderReadModel) SetNumber(v string)`

SetNumber sets Number field to given value.


### GetPurchaseOrder

`func (o *OrderReadModel) GetPurchaseOrder() PurchaseOrderReadModel`

GetPurchaseOrder returns the PurchaseOrder field if non-nil, zero value otherwise.

### GetPurchaseOrderOk

`func (o *OrderReadModel) GetPurchaseOrderOk() (*PurchaseOrderReadModel, bool)`

GetPurchaseOrderOk returns a tuple with the PurchaseOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseOrder

`func (o *OrderReadModel) SetPurchaseOrder(v PurchaseOrderReadModel)`

SetPurchaseOrder sets PurchaseOrder field to given value.

### HasPurchaseOrder

`func (o *OrderReadModel) HasPurchaseOrder() bool`

HasPurchaseOrder returns a boolean if a field has been set.

### GetContacts

`func (o *OrderReadModel) GetContacts() []ContactItem`

GetContacts returns the Contacts field if non-nil, zero value otherwise.

### GetContactsOk

`func (o *OrderReadModel) GetContactsOk() (*[]ContactItem, bool)`

GetContactsOk returns a tuple with the Contacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContacts

`func (o *OrderReadModel) SetContacts(v []ContactItem)`

SetContacts sets Contacts field to given value.

### HasContacts

`func (o *OrderReadModel) HasContacts() bool`

HasContacts returns a boolean if a field has been set.

### GetTags

`func (o *OrderReadModel) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *OrderReadModel) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *OrderReadModel) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *OrderReadModel) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


