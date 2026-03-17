# \CampaignAnalyticsAPI

All URIs are relative to *https://api.mailodds.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCampaignABResults**](CampaignAnalyticsAPI.md#GetCampaignABResults) | **Get** /v1/campaigns/{campaign_id}/ab-results | Get A/B test results
[**GetCampaignAttribution**](CampaignAnalyticsAPI.md#GetCampaignAttribution) | **Get** /v1/campaigns/{campaign_id}/conversions/attribution | Get campaign attribution
[**GetCampaignDeliveryConfidence**](CampaignAnalyticsAPI.md#GetCampaignDeliveryConfidence) | **Get** /v1/campaigns/{campaign_id}/delivery-confidence | Get pre-send delivery confidence
[**GetCampaignFunnel**](CampaignAnalyticsAPI.md#GetCampaignFunnel) | **Get** /v1/campaigns/{campaign_id}/funnel | Get campaign funnel
[**GetCampaignProviderIntelligence**](CampaignAnalyticsAPI.md#GetCampaignProviderIntelligence) | **Get** /v1/campaigns/{campaign_id}/provider-intelligence | Get provider intelligence



## GetCampaignABResults

> GetCampaignABResults200Response GetCampaignABResults(ctx, campaignId).Execute()

Get A/B test results



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	campaignId := "campaignId_example" // string | Campaign UUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CampaignAnalyticsAPI.GetCampaignABResults(context.Background(), campaignId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CampaignAnalyticsAPI.GetCampaignABResults``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCampaignABResults`: GetCampaignABResults200Response
	fmt.Fprintf(os.Stdout, "Response from `CampaignAnalyticsAPI.GetCampaignABResults`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **string** | Campaign UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCampaignABResultsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetCampaignABResults200Response**](GetCampaignABResults200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCampaignAttribution

> GetCampaignAttribution200Response GetCampaignAttribution(ctx, campaignId).Execute()

Get campaign attribution



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	campaignId := "campaignId_example" // string | Campaign UUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CampaignAnalyticsAPI.GetCampaignAttribution(context.Background(), campaignId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CampaignAnalyticsAPI.GetCampaignAttribution``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCampaignAttribution`: GetCampaignAttribution200Response
	fmt.Fprintf(os.Stdout, "Response from `CampaignAnalyticsAPI.GetCampaignAttribution`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **string** | Campaign UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCampaignAttributionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetCampaignAttribution200Response**](GetCampaignAttribution200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCampaignDeliveryConfidence

> GetCampaignDeliveryConfidence200Response GetCampaignDeliveryConfidence(ctx, campaignId).Execute()

Get pre-send delivery confidence



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	campaignId := "campaignId_example" // string | Campaign UUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CampaignAnalyticsAPI.GetCampaignDeliveryConfidence(context.Background(), campaignId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CampaignAnalyticsAPI.GetCampaignDeliveryConfidence``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCampaignDeliveryConfidence`: GetCampaignDeliveryConfidence200Response
	fmt.Fprintf(os.Stdout, "Response from `CampaignAnalyticsAPI.GetCampaignDeliveryConfidence`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **string** | Campaign UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCampaignDeliveryConfidenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetCampaignDeliveryConfidence200Response**](GetCampaignDeliveryConfidence200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCampaignFunnel

> GetCampaignFunnel200Response GetCampaignFunnel(ctx, campaignId).Execute()

Get campaign funnel



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	campaignId := "campaignId_example" // string | Campaign UUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CampaignAnalyticsAPI.GetCampaignFunnel(context.Background(), campaignId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CampaignAnalyticsAPI.GetCampaignFunnel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCampaignFunnel`: GetCampaignFunnel200Response
	fmt.Fprintf(os.Stdout, "Response from `CampaignAnalyticsAPI.GetCampaignFunnel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **string** | Campaign UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCampaignFunnelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetCampaignFunnel200Response**](GetCampaignFunnel200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCampaignProviderIntelligence

> GetCampaignProviderIntelligence200Response GetCampaignProviderIntelligence(ctx, campaignId).Execute()

Get provider intelligence



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	campaignId := "campaignId_example" // string | Campaign UUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CampaignAnalyticsAPI.GetCampaignProviderIntelligence(context.Background(), campaignId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CampaignAnalyticsAPI.GetCampaignProviderIntelligence``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCampaignProviderIntelligence`: GetCampaignProviderIntelligence200Response
	fmt.Fprintf(os.Stdout, "Response from `CampaignAnalyticsAPI.GetCampaignProviderIntelligence`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **string** | Campaign UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCampaignProviderIntelligenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetCampaignProviderIntelligence200Response**](GetCampaignProviderIntelligence200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

