# \DomainInsightsAPI

All URIs are relative to *https://api.mailodds.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDomainHookEffectiveness**](DomainInsightsAPI.md#GetDomainHookEffectiveness) | **Get** /v1/sending-domains/{domain_id}/insights/hook-effectiveness | Get hook effectiveness metrics
[**GetDomainInsightsFunnel**](DomainInsightsAPI.md#GetDomainInsightsFunnel) | **Get** /v1/sending-domains/{domain_id}/insights/funnel | Get domain engagement funnel
[**GetDomainInsightsTrends**](DomainInsightsAPI.md#GetDomainInsightsTrends) | **Get** /v1/sending-domains/{domain_id}/insights/trends | Get domain engagement trends



## GetDomainHookEffectiveness

> GetDomainHookEffectiveness200Response GetDomainHookEffectiveness(ctx, domainId).Days(days).Execute()

Get hook effectiveness metrics



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
	domainId := "domainId_example" // string | Sending domain ID
	days := int32(56) // int32 | Lookback period in days (optional) (default to 30)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DomainInsightsAPI.GetDomainHookEffectiveness(context.Background(), domainId).Days(days).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainInsightsAPI.GetDomainHookEffectiveness``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDomainHookEffectiveness`: GetDomainHookEffectiveness200Response
	fmt.Fprintf(os.Stdout, "Response from `DomainInsightsAPI.GetDomainHookEffectiveness`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainId** | **string** | Sending domain ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDomainHookEffectivenessRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **days** | **int32** | Lookback period in days | [default to 30]

### Return type

[**GetDomainHookEffectiveness200Response**](GetDomainHookEffectiveness200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDomainInsightsFunnel

> GetDomainInsightsFunnel200Response GetDomainInsightsFunnel(ctx, domainId).Days(days).Execute()

Get domain engagement funnel



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
	domainId := "domainId_example" // string | Sending domain ID
	days := int32(56) // int32 | Lookback period in days (optional) (default to 30)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DomainInsightsAPI.GetDomainInsightsFunnel(context.Background(), domainId).Days(days).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainInsightsAPI.GetDomainInsightsFunnel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDomainInsightsFunnel`: GetDomainInsightsFunnel200Response
	fmt.Fprintf(os.Stdout, "Response from `DomainInsightsAPI.GetDomainInsightsFunnel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainId** | **string** | Sending domain ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDomainInsightsFunnelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **days** | **int32** | Lookback period in days | [default to 30]

### Return type

[**GetDomainInsightsFunnel200Response**](GetDomainInsightsFunnel200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDomainInsightsTrends

> GetDomainInsightsTrends200Response GetDomainInsightsTrends(ctx, domainId).Days(days).Execute()

Get domain engagement trends



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
	domainId := "domainId_example" // string | Sending domain ID
	days := int32(56) // int32 | Lookback period in days (optional) (default to 30)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DomainInsightsAPI.GetDomainInsightsTrends(context.Background(), domainId).Days(days).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainInsightsAPI.GetDomainInsightsTrends``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDomainInsightsTrends`: GetDomainInsightsTrends200Response
	fmt.Fprintf(os.Stdout, "Response from `DomainInsightsAPI.GetDomainInsightsTrends`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainId** | **string** | Sending domain ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDomainInsightsTrendsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **days** | **int32** | Lookback period in days | [default to 30]

### Return type

[**GetDomainInsightsTrends200Response**](GetDomainInsightsTrends200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

