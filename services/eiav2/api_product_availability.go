/*
Equinix Internet Access API

Equinix Internet Access provides direct access to the Internet with scalable bandwidth options in IBX data centers.

Contact: api-support@equinix.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eiav2

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)

// ProductAvailabilityApiService ProductAvailabilityApi service
type ProductAvailabilityApiService service

type ApiGetIbxsRequest struct {
	ctx                            context.Context
	ApiService                     *ProductAvailabilityApiService
	serviceConnectionType          *ConnectionType
	offset                         *int32
	limit                          *int32
	connectionAsideAccessPointType *GetIbxsConnectionAsideAccessPointTypeParameter
	assetType                      *GetIbxsAssetTypeParameter
}

// Service connection type (physical, virtual)
func (r ApiGetIbxsRequest) ServiceConnectionType(serviceConnectionType ConnectionType) ApiGetIbxsRequest {
	r.serviceConnectionType = &serviceConnectionType
	return r
}

// Pagination offset
func (r ApiGetIbxsRequest) Offset(offset int32) ApiGetIbxsRequest {
	r.offset = &offset
	return r
}

// Max number of returned results
func (r ApiGetIbxsRequest) Limit(limit int32) ApiGetIbxsRequest {
	r.limit = &limit
	return r
}

// Service connection access point type. When not provided, COLO type is used by default.
func (r ApiGetIbxsRequest) ConnectionAsideAccessPointType(connectionAsideAccessPointType GetIbxsConnectionAsideAccessPointTypeParameter) ApiGetIbxsRequest {
	r.connectionAsideAccessPointType = &connectionAsideAccessPointType
	return r
}

// Specifies the type of an asset the user must have in the IBX
func (r ApiGetIbxsRequest) AssetType(assetType GetIbxsAssetTypeParameter) ApiGetIbxsRequest {
	r.assetType = &assetType
	return r
}

func (r ApiGetIbxsRequest) Execute() (*IbxPage, *http.Response, error) {
	return r.ApiService.GetIbxsExecute(r)
}

/*
GetIbxs Returns IBXs where EIA is available

Get IBXs <font color="red"><sup color="red">Beta</sup></font>

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetIbxsRequest
*/
func (a *ProductAvailabilityApiService) GetIbxs(ctx context.Context) ApiGetIbxsRequest {
	return ApiGetIbxsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return IbxPage
func (a *ProductAvailabilityApiService) GetIbxsExecute(r ApiGetIbxsRequest) (*IbxPage, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *IbxPage
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProductAvailabilityApiService.GetIbxs")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/internetAccess/v2/ibxs"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.serviceConnectionType == nil {
		return localVarReturnValue, nil, reportError("serviceConnectionType is required and must be specified")
	}

	if r.offset != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "offset", r.offset, "")
	} else {
		var defaultValue int32 = 0
		r.offset = &defaultValue
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	} else {
		var defaultValue int32 = 50
		r.limit = &defaultValue
	}
	parameterAddToHeaderOrQuery(localVarQueryParams, "service.connection.type", r.serviceConnectionType, "")
	if r.connectionAsideAccessPointType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "connection.aside.accessPoint.type", r.connectionAsideAccessPointType, "")
	}
	if r.assetType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "asset.type", r.assetType, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v []Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v []Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v []Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v []Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v []Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
