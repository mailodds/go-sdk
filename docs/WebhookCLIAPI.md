# \WebhookCLIAPI

All URIs are relative to *https://api.mailodds.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateWebhookCliSession**](WebhookCLIAPI.md#CreateWebhookCliSession) | **Post** /v1/webhook-cli/sessions | Create CLI forwarding session
[**DeleteWebhookCliSession**](WebhookCLIAPI.md#DeleteWebhookCliSession) | **Delete** /v1/webhook-cli/sessions/{session_id} | Close CLI session
[**ListWebhookDeliveries**](WebhookCLIAPI.md#ListWebhookDeliveries) | **Get** /v1/webhook-cli/deliveries | List recent webhook deliveries
[**ReplayWebhookDelivery**](WebhookCLIAPI.md#ReplayWebhookDelivery) | **Post** /v1/webhook-cli/deliveries/{delivery_id}/replay | Replay webhook delivery



## CreateWebhookCliSession

> CreateWebhookCliSession201Response CreateWebhookCliSession(ctx).CreateWebhookCliSessionRequest(createWebhookCliSessionRequest).Execute()

Create CLI forwarding session



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
	createWebhookCliSessionRequest := *openapiclient.NewCreateWebhookCliSessionRequest() // CreateWebhookCliSessionRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhookCLIAPI.CreateWebhookCliSession(context.Background()).CreateWebhookCliSessionRequest(createWebhookCliSessionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookCLIAPI.CreateWebhookCliSession``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateWebhookCliSession`: CreateWebhookCliSession201Response
	fmt.Fprintf(os.Stdout, "Response from `WebhookCLIAPI.CreateWebhookCliSession`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateWebhookCliSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createWebhookCliSessionRequest** | [**CreateWebhookCliSessionRequest**](CreateWebhookCliSessionRequest.md) |  | 

### Return type

[**CreateWebhookCliSession201Response**](CreateWebhookCliSession201Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteWebhookCliSession

> DeleteWebhookCliSession200Response DeleteWebhookCliSession(ctx, sessionId).Execute()

Close CLI session



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
	sessionId := "sessionId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhookCLIAPI.DeleteWebhookCliSession(context.Background(), sessionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookCLIAPI.DeleteWebhookCliSession``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteWebhookCliSession`: DeleteWebhookCliSession200Response
	fmt.Fprintf(os.Stdout, "Response from `WebhookCLIAPI.DeleteWebhookCliSession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWebhookCliSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeleteWebhookCliSession200Response**](DeleteWebhookCliSession200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWebhookDeliveries

> ListWebhookDeliveries200Response ListWebhookDeliveries(ctx).Limit(limit).Execute()

List recent webhook deliveries



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
	limit := int32(56) // int32 | Maximum deliveries to return (optional) (default to 50)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhookCLIAPI.ListWebhookDeliveries(context.Background()).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookCLIAPI.ListWebhookDeliveries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListWebhookDeliveries`: ListWebhookDeliveries200Response
	fmt.Fprintf(os.Stdout, "Response from `WebhookCLIAPI.ListWebhookDeliveries`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListWebhookDeliveriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Maximum deliveries to return | [default to 50]

### Return type

[**ListWebhookDeliveries200Response**](ListWebhookDeliveries200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplayWebhookDelivery

> ReplayWebhookDelivery200Response ReplayWebhookDelivery(ctx, deliveryId).Execute()

Replay webhook delivery



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
	deliveryId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhookCLIAPI.ReplayWebhookDelivery(context.Background(), deliveryId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookCLIAPI.ReplayWebhookDelivery``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReplayWebhookDelivery`: ReplayWebhookDelivery200Response
	fmt.Fprintf(os.Stdout, "Response from `WebhookCLIAPI.ReplayWebhookDelivery`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deliveryId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplayWebhookDeliveryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ReplayWebhookDelivery200Response**](ReplayWebhookDelivery200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

