# \ReputationAPI

All URIs are relative to *https://api.mailodds.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetReputation**](ReputationAPI.md#GetReputation) | **Get** /v1/reputation | Get account reputation
[**GetReputationTimeline**](ReputationAPI.md#GetReputationTimeline) | **Get** /v1/reputation/timeline | Get reputation timeline



## GetReputation

> GetReputation200Response GetReputation(ctx).Period(period).Execute()

Get account reputation



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
	period := "period_example" // string | Evaluation period (optional) (default to "7d")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReputationAPI.GetReputation(context.Background()).Period(period).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReputationAPI.GetReputation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetReputation`: GetReputation200Response
	fmt.Fprintf(os.Stdout, "Response from `ReputationAPI.GetReputation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetReputationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **period** | **string** | Evaluation period | [default to &quot;7d&quot;]

### Return type

[**GetReputation200Response**](GetReputation200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReputationTimeline

> GetReputationTimeline200Response GetReputationTimeline(ctx).Period(period).Execute()

Get reputation timeline



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
	period := "period_example" // string | Timeline period (optional) (default to "30d")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReputationAPI.GetReputationTimeline(context.Background()).Period(period).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReputationAPI.GetReputationTimeline``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetReputationTimeline`: GetReputationTimeline200Response
	fmt.Fprintf(os.Stdout, "Response from `ReputationAPI.GetReputationTimeline`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetReputationTimelineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **period** | **string** | Timeline period | [default to &quot;30d&quot;]

### Return type

[**GetReputationTimeline200Response**](GetReputationTimeline200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

