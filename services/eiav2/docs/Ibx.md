# Ibx

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | **string** |  | 
**CountryCode** | Pointer to **string** |  | [optional] 
**CountryName** | Pointer to **string** |  | [optional] 
**Region** | Pointer to [**IbxRegion**](IbxRegion.md) |  | [optional] 
**MetroCode** | Pointer to **string** |  | [optional] 
**IbxCode** | Pointer to **string** | IBX data center code | [optional] 
**GeoCoordinates** | Pointer to [**GeoCoordinates**](GeoCoordinates.md) |  | [optional] 

## Methods

### NewIbx

`func NewIbx(href string, ) *Ibx`

NewIbx instantiates a new Ibx object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIbxWithDefaults

`func NewIbxWithDefaults() *Ibx`

NewIbxWithDefaults instantiates a new Ibx object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *Ibx) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *Ibx) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *Ibx) SetHref(v string)`

SetHref sets Href field to given value.


### GetCountryCode

`func (o *Ibx) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *Ibx) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *Ibx) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *Ibx) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetCountryName

`func (o *Ibx) GetCountryName() string`

GetCountryName returns the CountryName field if non-nil, zero value otherwise.

### GetCountryNameOk

`func (o *Ibx) GetCountryNameOk() (*string, bool)`

GetCountryNameOk returns a tuple with the CountryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryName

`func (o *Ibx) SetCountryName(v string)`

SetCountryName sets CountryName field to given value.

### HasCountryName

`func (o *Ibx) HasCountryName() bool`

HasCountryName returns a boolean if a field has been set.

### GetRegion

`func (o *Ibx) GetRegion() IbxRegion`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *Ibx) GetRegionOk() (*IbxRegion, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *Ibx) SetRegion(v IbxRegion)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *Ibx) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetMetroCode

`func (o *Ibx) GetMetroCode() string`

GetMetroCode returns the MetroCode field if non-nil, zero value otherwise.

### GetMetroCodeOk

`func (o *Ibx) GetMetroCodeOk() (*string, bool)`

GetMetroCodeOk returns a tuple with the MetroCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetroCode

`func (o *Ibx) SetMetroCode(v string)`

SetMetroCode sets MetroCode field to given value.

### HasMetroCode

`func (o *Ibx) HasMetroCode() bool`

HasMetroCode returns a boolean if a field has been set.

### GetIbxCode

`func (o *Ibx) GetIbxCode() string`

GetIbxCode returns the IbxCode field if non-nil, zero value otherwise.

### GetIbxCodeOk

`func (o *Ibx) GetIbxCodeOk() (*string, bool)`

GetIbxCodeOk returns a tuple with the IbxCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIbxCode

`func (o *Ibx) SetIbxCode(v string)`

SetIbxCode sets IbxCode field to given value.

### HasIbxCode

`func (o *Ibx) HasIbxCode() bool`

HasIbxCode returns a boolean if a field has been set.

### GetGeoCoordinates

`func (o *Ibx) GetGeoCoordinates() GeoCoordinates`

GetGeoCoordinates returns the GeoCoordinates field if non-nil, zero value otherwise.

### GetGeoCoordinatesOk

`func (o *Ibx) GetGeoCoordinatesOk() (*GeoCoordinates, bool)`

GetGeoCoordinatesOk returns a tuple with the GeoCoordinates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeoCoordinates

`func (o *Ibx) SetGeoCoordinates(v GeoCoordinates)`

SetGeoCoordinates sets GeoCoordinates field to given value.

### HasGeoCoordinates

`func (o *Ibx) HasGeoCoordinates() bool`

HasGeoCoordinates returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


