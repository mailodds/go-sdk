# \DeliverabilityAdvisorAPI

All URIs are relative to *https://api.mailodds.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DismissDeliverabilityRecommendation**](DeliverabilityAdvisorAPI.md#DismissDeliverabilityRecommendation) | **Post** /v1/deliverability/recommendations/{recommendation_id}/dismiss | Dismiss a deliverability recommendation
[**GetDeliverabilityRecommendations**](DeliverabilityAdvisorAPI.md#GetDeliverabilityRecommendations) | **Get** /v1/deliverability/recommendations | Get deliverability recommendations



## DismissDeliverabilityRecommendation

> DismissDeliverabilityRecommendation(ctx, recommendationId).Execute()

Dismiss a deliverability recommendation



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
	recommendationId := "recommendationId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DeliverabilityAdvisorAPI.DismissDeliverabilityRecommendation(context.Background(), recommendationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeliverabilityAdvisorAPI.DismissDeliverabilityRecommendation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**recommendationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDismissDeliverabilityRecommendationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## GetDeliverabilityRecommendations

> GetDeliverabilityRecommendations(ctx).Execute()

Get deliverability recommendations



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DeliverabilityAdvisorAPI.GetDeliverabilityRecommendations(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeliverabilityAdvisorAPI.GetDeliverabilityRecommendations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeliverabilityRecommendationsRequest struct via the builder pattern


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

