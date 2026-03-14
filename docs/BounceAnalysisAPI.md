# \BounceAnalysisAPI

All URIs are relative to *https://api.mailodds.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBounceAnalysis**](BounceAnalysisAPI.md#CreateBounceAnalysis) | **Post** /v1/bounce-analyses | Analyze bounce logs
[**CrossReferenceBounces**](BounceAnalysisAPI.md#CrossReferenceBounces) | **Get** /v1/bounce-analyses/{analysis_id}/cross-reference | Cross-reference bounces with validation logs
[**GetBounceAnalysis**](BounceAnalysisAPI.md#GetBounceAnalysis) | **Get** /v1/bounce-analyses/{analysis_id} | Get bounce analysis
[**GetBounceRecords**](BounceAnalysisAPI.md#GetBounceRecords) | **Get** /v1/bounce-analyses/{analysis_id}/records | Get bounce records



## CreateBounceAnalysis

> BounceAnalysisResponse CreateBounceAnalysis(ctx).CreateBounceAnalysisRequest(createBounceAnalysisRequest).Execute()

Analyze bounce logs



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mailodds/go-sdk/mailodds"
)

func main() {
	createBounceAnalysisRequest := *openapiclient.NewCreateBounceAnalysisRequest("DomainId_example") // CreateBounceAnalysisRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BounceAnalysisAPI.CreateBounceAnalysis(context.Background()).CreateBounceAnalysisRequest(createBounceAnalysisRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BounceAnalysisAPI.CreateBounceAnalysis``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateBounceAnalysis`: BounceAnalysisResponse
	fmt.Fprintf(os.Stdout, "Response from `BounceAnalysisAPI.CreateBounceAnalysis`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateBounceAnalysisRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createBounceAnalysisRequest** | [**CreateBounceAnalysisRequest**](CreateBounceAnalysisRequest.md) |  | 

### Return type

[**BounceAnalysisResponse**](BounceAnalysisResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CrossReferenceBounces

> CrossReferenceBounces200Response CrossReferenceBounces(ctx, analysisId).Execute()

Cross-reference bounces with validation logs



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mailodds/go-sdk/mailodds"
)

func main() {
	analysisId := "analysisId_example" // string | Bounce analysis UUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BounceAnalysisAPI.CrossReferenceBounces(context.Background(), analysisId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BounceAnalysisAPI.CrossReferenceBounces``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CrossReferenceBounces`: CrossReferenceBounces200Response
	fmt.Fprintf(os.Stdout, "Response from `BounceAnalysisAPI.CrossReferenceBounces`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**analysisId** | **string** | Bounce analysis UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCrossReferenceBouncesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CrossReferenceBounces200Response**](CrossReferenceBounces200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBounceAnalysis

> BounceAnalysisResponse GetBounceAnalysis(ctx, analysisId).Execute()

Get bounce analysis



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mailodds/go-sdk/mailodds"
)

func main() {
	analysisId := "analysisId_example" // string | Bounce analysis UUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BounceAnalysisAPI.GetBounceAnalysis(context.Background(), analysisId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BounceAnalysisAPI.GetBounceAnalysis``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBounceAnalysis`: BounceAnalysisResponse
	fmt.Fprintf(os.Stdout, "Response from `BounceAnalysisAPI.GetBounceAnalysis`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**analysisId** | **string** | Bounce analysis UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBounceAnalysisRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BounceAnalysisResponse**](BounceAnalysisResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBounceRecords

> GetBounceRecords200Response GetBounceRecords(ctx, analysisId).Page(page).PerPage(perPage).Type_(type_).Execute()

Get bounce records



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mailodds/go-sdk/mailodds"
)

func main() {
	analysisId := "analysisId_example" // string | Bounce analysis UUID
	page := int32(56) // int32 | Page number (optional) (default to 1)
	perPage := int32(56) // int32 | Items per page (optional) (default to 50)
	type_ := "type__example" // string | Filter by bounce type (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BounceAnalysisAPI.GetBounceRecords(context.Background(), analysisId).Page(page).PerPage(perPage).Type_(type_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BounceAnalysisAPI.GetBounceRecords``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBounceRecords`: GetBounceRecords200Response
	fmt.Fprintf(os.Stdout, "Response from `BounceAnalysisAPI.GetBounceRecords`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**analysisId** | **string** | Bounce analysis UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBounceRecordsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Page number | [default to 1]
 **perPage** | **int32** | Items per page | [default to 50]
 **type_** | **string** | Filter by bounce type | 

### Return type

[**GetBounceRecords200Response**](GetBounceRecords200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

