# ContactItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | Pointer to **string** | Account URI | [optional] 
**Type** | [**ContactItemType**](ContactItemType.md) |  | 
**RegisteredUser** | Pointer to **string** | Identifies (e.g., contactId, userId, userKey) a registered user.  | [optional] 
**FirstName** | Pointer to **string** |  | [optional] 
**LastName** | Pointer to **string** |  | [optional] 
**Timezone** | Pointer to **string** |  | [optional] 
**Availability** | Pointer to [**ContactItemAvailability**](ContactItemAvailability.md) |  | [optional] 
**Notes** | Pointer to **string** | Describe any contact preferences  | [optional] 
**Details** | Pointer to [**[]ContactItemDetails**](ContactItemDetails.md) |  | [optional] 

## Methods

### NewContactItem

`func NewContactItem(type_ ContactItemType, ) *ContactItem`

NewContactItem instantiates a new ContactItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContactItemWithDefaults

`func NewContactItemWithDefaults() *ContactItem`

NewContactItemWithDefaults instantiates a new ContactItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *ContactItem) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *ContactItem) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *ContactItem) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *ContactItem) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetType

`func (o *ContactItem) GetType() ContactItemType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ContactItem) GetTypeOk() (*ContactItemType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ContactItem) SetType(v ContactItemType)`

SetType sets Type field to given value.


### GetRegisteredUser

`func (o *ContactItem) GetRegisteredUser() string`

GetRegisteredUser returns the RegisteredUser field if non-nil, zero value otherwise.

### GetRegisteredUserOk

`func (o *ContactItem) GetRegisteredUserOk() (*string, bool)`

GetRegisteredUserOk returns a tuple with the RegisteredUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredUser

`func (o *ContactItem) SetRegisteredUser(v string)`

SetRegisteredUser sets RegisteredUser field to given value.

### HasRegisteredUser

`func (o *ContactItem) HasRegisteredUser() bool`

HasRegisteredUser returns a boolean if a field has been set.

### GetFirstName

`func (o *ContactItem) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *ContactItem) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *ContactItem) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *ContactItem) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *ContactItem) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *ContactItem) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *ContactItem) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *ContactItem) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetTimezone

`func (o *ContactItem) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *ContactItem) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *ContactItem) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *ContactItem) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetAvailability

`func (o *ContactItem) GetAvailability() ContactItemAvailability`

GetAvailability returns the Availability field if non-nil, zero value otherwise.

### GetAvailabilityOk

`func (o *ContactItem) GetAvailabilityOk() (*ContactItemAvailability, bool)`

GetAvailabilityOk returns a tuple with the Availability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailability

`func (o *ContactItem) SetAvailability(v ContactItemAvailability)`

SetAvailability sets Availability field to given value.

### HasAvailability

`func (o *ContactItem) HasAvailability() bool`

HasAvailability returns a boolean if a field has been set.

### GetNotes

`func (o *ContactItem) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *ContactItem) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *ContactItem) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *ContactItem) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetDetails

`func (o *ContactItem) GetDetails() []ContactItemDetails`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *ContactItem) GetDetailsOk() (*[]ContactItemDetails, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *ContactItem) SetDetails(v []ContactItemDetails)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *ContactItem) HasDetails() bool`

HasDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


