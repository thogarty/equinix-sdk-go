# Account

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | **string** | Account URI | 
**AccountNumber** | **string** | Account identifier | 
**AccountName** | Pointer to **string** | Name of account | [optional] 
**OrgId** | Pointer to **string** | Customer organization identifier | [optional] 
**OrganizationName** | Pointer to **string** | Customer organization name | [optional] 

## Methods

### NewAccount

`func NewAccount(href string, accountNumber string, ) *Account`

NewAccount instantiates a new Account object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountWithDefaults

`func NewAccountWithDefaults() *Account`

NewAccountWithDefaults instantiates a new Account object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *Account) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *Account) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *Account) SetHref(v string)`

SetHref sets Href field to given value.


### GetAccountNumber

`func (o *Account) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *Account) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *Account) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.


### GetAccountName

`func (o *Account) GetAccountName() string`

GetAccountName returns the AccountName field if non-nil, zero value otherwise.

### GetAccountNameOk

`func (o *Account) GetAccountNameOk() (*string, bool)`

GetAccountNameOk returns a tuple with the AccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountName

`func (o *Account) SetAccountName(v string)`

SetAccountName sets AccountName field to given value.

### HasAccountName

`func (o *Account) HasAccountName() bool`

HasAccountName returns a boolean if a field has been set.

### GetOrgId

`func (o *Account) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *Account) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *Account) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *Account) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetOrganizationName

`func (o *Account) GetOrganizationName() string`

GetOrganizationName returns the OrganizationName field if non-nil, zero value otherwise.

### GetOrganizationNameOk

`func (o *Account) GetOrganizationNameOk() (*string, bool)`

GetOrganizationNameOk returns a tuple with the OrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationName

`func (o *Account) SetOrganizationName(v string)`

SetOrganizationName sets OrganizationName field to given value.

### HasOrganizationName

`func (o *Account) HasOrganizationName() bool`

HasOrganizationName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


