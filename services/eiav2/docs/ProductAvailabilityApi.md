# \ProductAvailabilityApi

All URIs are relative to *https://api.equinix.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetIbxs**](ProductAvailabilityApi.md#GetIbxs) | **Get** /internetAccess/v2/ibxs | Returns IBXs where EIA is available



## GetIbxs

> IbxPage GetIbxs(ctx).ServiceConnectionType(serviceConnectionType).Offset(offset).Limit(limit).ConnectionAsideAccessPointType(connectionAsideAccessPointType).AssetType(assetType).Execute()

Returns IBXs where EIA is available



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/equinix/equinix-sdk-go/services/eiav2"
)

func main() {
	serviceConnectionType := openapiclient.ConnectionType("IA_C") // ConnectionType | Service connection type (physical, virtual)
	offset := int32(0) // int32 | Pagination offset (optional) (default to 0)
	limit := int32(50) // int32 | Max number of returned results (optional) (default to 50)
	connectionAsideAccessPointType := openapiclient.getIbxs_connection_aside_accessPoint_type_parameter("COLO") // GetIbxsConnectionAsideAccessPointTypeParameter | Service connection access point type. When not provided, COLO type is used by default. (optional)
	assetType := openapiclient.getIbxs_asset_type_parameter("CABINET") // GetIbxsAssetTypeParameter | Specifies the type of an asset the user must have in the IBX (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductAvailabilityApi.GetIbxs(context.Background()).ServiceConnectionType(serviceConnectionType).Offset(offset).Limit(limit).ConnectionAsideAccessPointType(connectionAsideAccessPointType).AssetType(assetType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductAvailabilityApi.GetIbxs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIbxs`: IbxPage
	fmt.Fprintf(os.Stdout, "Response from `ProductAvailabilityApi.GetIbxs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetIbxsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **serviceConnectionType** | [**ConnectionType**](ConnectionType.md) | Service connection type (physical, virtual) | 
 **offset** | **int32** | Pagination offset | [default to 0]
 **limit** | **int32** | Max number of returned results | [default to 50]
 **connectionAsideAccessPointType** | [**GetIbxsConnectionAsideAccessPointTypeParameter**](GetIbxsConnectionAsideAccessPointTypeParameter.md) | Service connection access point type. When not provided, COLO type is used by default. | 
 **assetType** | [**GetIbxsAssetTypeParameter**](GetIbxsAssetTypeParameter.md) | Specifies the type of an asset the user must have in the IBX | 

### Return type

[**IbxPage**](IbxPage.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

