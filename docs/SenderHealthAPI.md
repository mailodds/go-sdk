# \SenderHealthAPI

All URIs are relative to *https://api.mailodds.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSenderHealth**](SenderHealthAPI.md#GetSenderHealth) | **Get** /v1/sender-health | Get sender health score
[**GetSenderHealthTrend**](SenderHealthAPI.md#GetSenderHealthTrend) | **Get** /v1/sender-health/trend | Get sender health trend



## GetSenderHealth

> GetSenderHealth200Response GetSenderHealth(ctx).Period(period).Execute()

Get sender health score



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/mailodds"
)

func main() {
	period := "period_example" // string | Time period for health calculation (optional) (default to "30d")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SenderHealthAPI.GetSenderHealth(context.Background()).Period(period).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SenderHealthAPI.GetSenderHealth``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSenderHealth`: GetSenderHealth200Response
	fmt.Fprintf(os.Stdout, "Response from `SenderHealthAPI.GetSenderHealth`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSenderHealthRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **period** | **string** | Time period for health calculation | [default to &quot;30d&quot;]

### Return type

[**GetSenderHealth200Response**](GetSenderHealth200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSenderHealthTrend

> GetSenderHealthTrend200Response GetSenderHealthTrend(ctx).Period(period).Execute()

Get sender health trend



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/mailodds"
)

func main() {
	period := "period_example" // string | Time period for trend data (optional) (default to "30d")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SenderHealthAPI.GetSenderHealthTrend(context.Background()).Period(period).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SenderHealthAPI.GetSenderHealthTrend``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSenderHealthTrend`: GetSenderHealthTrend200Response
	fmt.Fprintf(os.Stdout, "Response from `SenderHealthAPI.GetSenderHealthTrend`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSenderHealthTrendRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **period** | **string** | Time period for trend data | [default to &quot;30d&quot;]

### Return type

[**GetSenderHealthTrend200Response**](GetSenderHealthTrend200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

