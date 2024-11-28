# \EIAServiceApi

All URIs are relative to *https://api.equinix.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateEquinixInternetAccessv2**](EIAServiceApi.md#CreateEquinixInternetAccessv2) | **Post** /internetAccess/v2/services | Creates Equinix Internet Access Service
[**DeleteEquinixInternetAccess**](EIAServiceApi.md#DeleteEquinixInternetAccess) | **Delete** /internetAccess/v2/services/{serviceId} | Deletes Equinix Internet Access Service
[**GetEquinixInternetAccessServiceDetails**](EIAServiceApi.md#GetEquinixInternetAccessServiceDetails) | **Get** /internetAccess/v2/services/{serviceId} | Get Equinix Internet Access Service details
[**GetEquinixInternetAccessServices**](EIAServiceApi.md#GetEquinixInternetAccessServices) | **Post** /internetAccess/v2/services/search | Get Equinix Internet Access Services



## CreateEquinixInternetAccessv2

> ServiceCreateResponse CreateEquinixInternetAccessv2(ctx).ServiceCreateRequest(serviceCreateRequest).Execute()

Creates Equinix Internet Access Service



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
	serviceCreateRequest := *openapiclient.NewServiceCreateRequest(openapiclient.ServiceType("SINGLE"), []string{"9b8c5042-b553-4d5e-a2ac-c73bf6d4fd81"}, openapiclient.RoutingProtocolRequest{RoutingProtocolRequestBgp: openapiclient.NewRoutingProtocolRequestBgp("Type_example", openapiclient.RoutingProtocolRequestBgp_allOf_exportPolicy("FULL"))}) // ServiceCreateRequest | Options for creating Equinix Internet Access Service product 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EIAServiceApi.CreateEquinixInternetAccessv2(context.Background()).ServiceCreateRequest(serviceCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EIAServiceApi.CreateEquinixInternetAccessv2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateEquinixInternetAccessv2`: ServiceCreateResponse
	fmt.Fprintf(os.Stdout, "Response from `EIAServiceApi.CreateEquinixInternetAccessv2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateEquinixInternetAccessv2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **serviceCreateRequest** | [**ServiceCreateRequest**](ServiceCreateRequest.md) | Options for creating Equinix Internet Access Service product  | 

### Return type

[**ServiceCreateResponse**](ServiceCreateResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteEquinixInternetAccess

> DeleteEquinixInternetAccess(ctx, serviceId).DryRun(dryRun).Execute()

Deletes Equinix Internet Access Service



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
	serviceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Service identifier
	dryRun := true // bool | Setting this parameter to true will perform only request validation without actually deleting the service (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EIAServiceApi.DeleteEquinixInternetAccess(context.Background(), serviceId).DryRun(dryRun).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EIAServiceApi.DeleteEquinixInternetAccess``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | Service identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEquinixInternetAccessRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dryRun** | **bool** | Setting this parameter to true will perform only request validation without actually deleting the service | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEquinixInternetAccessServiceDetails

> ServiceReadModel GetEquinixInternetAccessServiceDetails(ctx, serviceId).Execute()

Get Equinix Internet Access Service details



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
	serviceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Service identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EIAServiceApi.GetEquinixInternetAccessServiceDetails(context.Background(), serviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EIAServiceApi.GetEquinixInternetAccessServiceDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEquinixInternetAccessServiceDetails`: ServiceReadModel
	fmt.Fprintf(os.Stdout, "Response from `EIAServiceApi.GetEquinixInternetAccessServiceDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** | Service identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEquinixInternetAccessServiceDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ServiceReadModel**](ServiceReadModel.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEquinixInternetAccessServices

> SearchResponse GetEquinixInternetAccessServices(ctx).Offset(offset).Limit(limit).Sort(sort).SearchRequest(searchRequest).Execute()

Get Equinix Internet Access Services



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
	offset := float32(20) // float32 | Search result offset (optional)
	limit := float32(35) // float32 | Search result limit (optional)
	sort := "/name,-/type" // string | Search result sorting (optional)
	searchRequest := *openapiclient.NewSearchRequest() // SearchRequest | Search filters (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EIAServiceApi.GetEquinixInternetAccessServices(context.Background()).Offset(offset).Limit(limit).Sort(sort).SearchRequest(searchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EIAServiceApi.GetEquinixInternetAccessServices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEquinixInternetAccessServices`: SearchResponse
	fmt.Fprintf(os.Stdout, "Response from `EIAServiceApi.GetEquinixInternetAccessServices`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEquinixInternetAccessServicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **float32** | Search result offset | 
 **limit** | **float32** | Search result limit | 
 **sort** | **string** | Search result sorting | 
 **searchRequest** | [**SearchRequest**](SearchRequest.md) | Search filters | 

### Return type

[**SearchResponse**](SearchResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

