# \BlacklistMonitoringAPI

All URIs are relative to *https://api.mailodds.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddBlacklistMonitor**](BlacklistMonitoringAPI.md#AddBlacklistMonitor) | **Post** /v1/blacklist-monitors | Add blacklist monitor
[**DeleteBlacklistMonitor**](BlacklistMonitoringAPI.md#DeleteBlacklistMonitor) | **Delete** /v1/blacklist-monitors/{monitor_id} | Delete a blacklist monitor
[**GetBlacklistHistory**](BlacklistMonitoringAPI.md#GetBlacklistHistory) | **Get** /v1/blacklist-monitors/{monitor_id}/history | Get blacklist check history
[**ListBlacklistMonitors**](BlacklistMonitoringAPI.md#ListBlacklistMonitors) | **Get** /v1/blacklist-monitors | List blacklist monitors
[**RunBlacklistCheck**](BlacklistMonitoringAPI.md#RunBlacklistCheck) | **Post** /v1/blacklist-monitors/{monitor_id}/check | Run blacklist check



## AddBlacklistMonitor

> AddBlacklistMonitor201Response AddBlacklistMonitor(ctx).AddBlacklistMonitorRequest(addBlacklistMonitorRequest).Execute()

Add blacklist monitor



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
	addBlacklistMonitorRequest := *openapiclient.NewAddBlacklistMonitorRequest("192.0.2.1", "TargetType_example") // AddBlacklistMonitorRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlacklistMonitoringAPI.AddBlacklistMonitor(context.Background()).AddBlacklistMonitorRequest(addBlacklistMonitorRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlacklistMonitoringAPI.AddBlacklistMonitor``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddBlacklistMonitor`: AddBlacklistMonitor201Response
	fmt.Fprintf(os.Stdout, "Response from `BlacklistMonitoringAPI.AddBlacklistMonitor`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddBlacklistMonitorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addBlacklistMonitorRequest** | [**AddBlacklistMonitorRequest**](AddBlacklistMonitorRequest.md) |  | 

### Return type

[**AddBlacklistMonitor201Response**](AddBlacklistMonitor201Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteBlacklistMonitor

> DeletePolicyRule200Response DeleteBlacklistMonitor(ctx, monitorId).Execute()

Delete a blacklist monitor



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
	monitorId := "monitorId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlacklistMonitoringAPI.DeleteBlacklistMonitor(context.Background(), monitorId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlacklistMonitoringAPI.DeleteBlacklistMonitor``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteBlacklistMonitor`: DeletePolicyRule200Response
	fmt.Fprintf(os.Stdout, "Response from `BlacklistMonitoringAPI.DeleteBlacklistMonitor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**monitorId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBlacklistMonitorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeletePolicyRule200Response**](DeletePolicyRule200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBlacklistHistory

> GetBlacklistHistory200Response GetBlacklistHistory(ctx, monitorId).Page(page).PerPage(perPage).Execute()

Get blacklist check history



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
	monitorId := "monitorId_example" // string | 
	page := int32(56) // int32 |  (optional) (default to 1)
	perPage := int32(56) // int32 |  (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlacklistMonitoringAPI.GetBlacklistHistory(context.Background(), monitorId).Page(page).PerPage(perPage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlacklistMonitoringAPI.GetBlacklistHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBlacklistHistory`: GetBlacklistHistory200Response
	fmt.Fprintf(os.Stdout, "Response from `BlacklistMonitoringAPI.GetBlacklistHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**monitorId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBlacklistHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** |  | [default to 1]
 **perPage** | **int32** |  | [default to 20]

### Return type

[**GetBlacklistHistory200Response**](GetBlacklistHistory200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBlacklistMonitors

> ListBlacklistMonitors200Response ListBlacklistMonitors(ctx).Execute()

List blacklist monitors



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
	resp, r, err := apiClient.BlacklistMonitoringAPI.ListBlacklistMonitors(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlacklistMonitoringAPI.ListBlacklistMonitors``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListBlacklistMonitors`: ListBlacklistMonitors200Response
	fmt.Fprintf(os.Stdout, "Response from `BlacklistMonitoringAPI.ListBlacklistMonitors`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListBlacklistMonitorsRequest struct via the builder pattern


### Return type

[**ListBlacklistMonitors200Response**](ListBlacklistMonitors200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RunBlacklistCheck

> RunBlacklistCheck200Response RunBlacklistCheck(ctx, monitorId).Execute()

Run blacklist check



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
	monitorId := "monitorId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlacklistMonitoringAPI.RunBlacklistCheck(context.Background(), monitorId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlacklistMonitoringAPI.RunBlacklistCheck``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RunBlacklistCheck`: RunBlacklistCheck200Response
	fmt.Fprintf(os.Stdout, "Response from `BlacklistMonitoringAPI.RunBlacklistCheck`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**monitorId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRunBlacklistCheckRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RunBlacklistCheck200Response**](RunBlacklistCheck200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

