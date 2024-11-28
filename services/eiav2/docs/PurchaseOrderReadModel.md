# PurchaseOrderReadModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | **string** | Purchase Order URI | 
**Type** | [**PurchaseOrderReadModelType**](PurchaseOrderReadModelType.md) |  | 
**Number** | **string** | Purchase Order number | 

## Methods

### NewPurchaseOrderReadModel

`func NewPurchaseOrderReadModel(href string, type_ PurchaseOrderReadModelType, number string, ) *PurchaseOrderReadModel`

NewPurchaseOrderReadModel instantiates a new PurchaseOrderReadModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPurchaseOrderReadModelWithDefaults

`func NewPurchaseOrderReadModelWithDefaults() *PurchaseOrderReadModel`

NewPurchaseOrderReadModelWithDefaults instantiates a new PurchaseOrderReadModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *PurchaseOrderReadModel) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *PurchaseOrderReadModel) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *PurchaseOrderReadModel) SetHref(v string)`

SetHref sets Href field to given value.


### GetType

`func (o *PurchaseOrderReadModel) GetType() PurchaseOrderReadModelType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PurchaseOrderReadModel) GetTypeOk() (*PurchaseOrderReadModelType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PurchaseOrderReadModel) SetType(v PurchaseOrderReadModelType)`

SetType sets Type field to given value.


### GetNumber

`func (o *PurchaseOrderReadModel) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *PurchaseOrderReadModel) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *PurchaseOrderReadModel) SetNumber(v string)`

SetNumber sets Number field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

