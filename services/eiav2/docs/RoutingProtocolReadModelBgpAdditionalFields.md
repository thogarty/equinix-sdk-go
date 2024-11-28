# RoutingProtocolReadModelBgpAdditionalFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerAsn** | **int64** | Customer Autonomous System Number  | 
**CustomerAsnRange** | Pointer to [**RoutingProtocolRequestBgpAllOfCustomerAsnRange**](RoutingProtocolRequestBgpAllOfCustomerAsnRange.md) |  | [optional] 
**EquinixAsn** | **int64** | Equinix Autonomous System Number  | 
**BgpAuthKey** | Pointer to **string** | BGP authentication key  | [optional] 
**ExportPolicy** | [**ExportPolicy**](ExportPolicy.md) |  | 

## Methods

### NewRoutingProtocolReadModelBgpAdditionalFields

`func NewRoutingProtocolReadModelBgpAdditionalFields(customerAsn int64, equinixAsn int64, exportPolicy ExportPolicy, ) *RoutingProtocolReadModelBgpAdditionalFields`

NewRoutingProtocolReadModelBgpAdditionalFields instantiates a new RoutingProtocolReadModelBgpAdditionalFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoutingProtocolReadModelBgpAdditionalFieldsWithDefaults

`func NewRoutingProtocolReadModelBgpAdditionalFieldsWithDefaults() *RoutingProtocolReadModelBgpAdditionalFields`

NewRoutingProtocolReadModelBgpAdditionalFieldsWithDefaults instantiates a new RoutingProtocolReadModelBgpAdditionalFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerAsn

`func (o *RoutingProtocolReadModelBgpAdditionalFields) GetCustomerAsn() int64`

GetCustomerAsn returns the CustomerAsn field if non-nil, zero value otherwise.

### GetCustomerAsnOk

`func (o *RoutingProtocolReadModelBgpAdditionalFields) GetCustomerAsnOk() (*int64, bool)`

GetCustomerAsnOk returns a tuple with the CustomerAsn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerAsn

`func (o *RoutingProtocolReadModelBgpAdditionalFields) SetCustomerAsn(v int64)`

SetCustomerAsn sets CustomerAsn field to given value.


### GetCustomerAsnRange

`func (o *RoutingProtocolReadModelBgpAdditionalFields) GetCustomerAsnRange() RoutingProtocolRequestBgpAllOfCustomerAsnRange`

GetCustomerAsnRange returns the CustomerAsnRange field if non-nil, zero value otherwise.

### GetCustomerAsnRangeOk

`func (o *RoutingProtocolReadModelBgpAdditionalFields) GetCustomerAsnRangeOk() (*RoutingProtocolRequestBgpAllOfCustomerAsnRange, bool)`

GetCustomerAsnRangeOk returns a tuple with the CustomerAsnRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerAsnRange

`func (o *RoutingProtocolReadModelBgpAdditionalFields) SetCustomerAsnRange(v RoutingProtocolRequestBgpAllOfCustomerAsnRange)`

SetCustomerAsnRange sets CustomerAsnRange field to given value.

### HasCustomerAsnRange

`func (o *RoutingProtocolReadModelBgpAdditionalFields) HasCustomerAsnRange() bool`

HasCustomerAsnRange returns a boolean if a field has been set.

### GetEquinixAsn

`func (o *RoutingProtocolReadModelBgpAdditionalFields) GetEquinixAsn() int64`

GetEquinixAsn returns the EquinixAsn field if non-nil, zero value otherwise.

### GetEquinixAsnOk

`func (o *RoutingProtocolReadModelBgpAdditionalFields) GetEquinixAsnOk() (*int64, bool)`

GetEquinixAsnOk returns a tuple with the EquinixAsn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquinixAsn

`func (o *RoutingProtocolReadModelBgpAdditionalFields) SetEquinixAsn(v int64)`

SetEquinixAsn sets EquinixAsn field to given value.


### GetBgpAuthKey

`func (o *RoutingProtocolReadModelBgpAdditionalFields) GetBgpAuthKey() string`

GetBgpAuthKey returns the BgpAuthKey field if non-nil, zero value otherwise.

### GetBgpAuthKeyOk

`func (o *RoutingProtocolReadModelBgpAdditionalFields) GetBgpAuthKeyOk() (*string, bool)`

GetBgpAuthKeyOk returns a tuple with the BgpAuthKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgpAuthKey

`func (o *RoutingProtocolReadModelBgpAdditionalFields) SetBgpAuthKey(v string)`

SetBgpAuthKey sets BgpAuthKey field to given value.

### HasBgpAuthKey

`func (o *RoutingProtocolReadModelBgpAdditionalFields) HasBgpAuthKey() bool`

HasBgpAuthKey returns a boolean if a field has been set.

### GetExportPolicy

`func (o *RoutingProtocolReadModelBgpAdditionalFields) GetExportPolicy() ExportPolicy`

GetExportPolicy returns the ExportPolicy field if non-nil, zero value otherwise.

### GetExportPolicyOk

`func (o *RoutingProtocolReadModelBgpAdditionalFields) GetExportPolicyOk() (*ExportPolicy, bool)`

GetExportPolicyOk returns a tuple with the ExportPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportPolicy

`func (o *RoutingProtocolReadModelBgpAdditionalFields) SetExportPolicy(v ExportPolicy)`

SetExportPolicy sets ExportPolicy field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


