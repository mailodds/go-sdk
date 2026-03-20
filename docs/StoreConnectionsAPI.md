# \StoreConnectionsAPI

All URIs are relative to *https://api.mailodds.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateStore**](StoreConnectionsAPI.md#CreateStore) | **Post** /v1/stores | Create a store connection
[**DisconnectStore**](StoreConnectionsAPI.md#DisconnectStore) | **Delete** /v1/stores/{store_id} | Disconnect a store
[**GetStore**](StoreConnectionsAPI.md#GetStore) | **Get** /v1/stores/{store_id} | Get a store connection
[**GetSyncJobErrors**](StoreConnectionsAPI.md#GetSyncJobErrors) | **Get** /v1/stores/{store_id}/sync-jobs/{job_id}/errors | Get sync job errors
[**ListStores**](StoreConnectionsAPI.md#ListStores) | **Get** /v1/stores | List store connections
[**ListSyncJobs**](StoreConnectionsAPI.md#ListSyncJobs) | **Get** /v1/stores/{store_id}/sync-jobs | List sync jobs
[**TriggerSync**](StoreConnectionsAPI.md#TriggerSync) | **Post** /v1/stores/{store_id}/sync | Trigger product sync
[**UpdateStore**](StoreConnectionsAPI.md#UpdateStore) | **Put** /v1/stores/{store_id} | Update a store connection



## CreateStore

> CreateStore201Response CreateStore(ctx).CreateStoreRequest(createStoreRequest).Execute()

Create a store connection



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
	createStoreRequest := *openapiclient.NewCreateStoreRequest("Platform_example", "StoreName_example", "StoreUrl_example", "AuthMethod_example") // CreateStoreRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StoreConnectionsAPI.CreateStore(context.Background()).CreateStoreRequest(createStoreRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StoreConnectionsAPI.CreateStore``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateStore`: CreateStore201Response
	fmt.Fprintf(os.Stdout, "Response from `StoreConnectionsAPI.CreateStore`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateStoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createStoreRequest** | [**CreateStoreRequest**](CreateStoreRequest.md) |  | 

### Return type

[**CreateStore201Response**](CreateStore201Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DisconnectStore

> DisconnectStore200Response DisconnectStore(ctx, storeId).Execute()

Disconnect a store



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
	storeId := "storeId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StoreConnectionsAPI.DisconnectStore(context.Background(), storeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StoreConnectionsAPI.DisconnectStore``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DisconnectStore`: DisconnectStore200Response
	fmt.Fprintf(os.Stdout, "Response from `StoreConnectionsAPI.DisconnectStore`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDisconnectStoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DisconnectStore200Response**](DisconnectStore200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStore

> CreateStore201Response GetStore(ctx, storeId).Execute()

Get a store connection



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
	storeId := "storeId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StoreConnectionsAPI.GetStore(context.Background(), storeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StoreConnectionsAPI.GetStore``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStore`: CreateStore201Response
	fmt.Fprintf(os.Stdout, "Response from `StoreConnectionsAPI.GetStore`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreateStore201Response**](CreateStore201Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSyncJobErrors

> GetSyncJobErrors200Response GetSyncJobErrors(ctx, storeId, jobId).Page(page).PerPage(perPage).Execute()

Get sync job errors



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
	storeId := "storeId_example" // string | 
	jobId := "jobId_example" // string | 
	page := int32(56) // int32 |  (optional) (default to 1)
	perPage := int32(56) // int32 |  (optional) (default to 50)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StoreConnectionsAPI.GetSyncJobErrors(context.Background(), storeId, jobId).Page(page).PerPage(perPage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StoreConnectionsAPI.GetSyncJobErrors``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSyncJobErrors`: GetSyncJobErrors200Response
	fmt.Fprintf(os.Stdout, "Response from `StoreConnectionsAPI.GetSyncJobErrors`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeId** | **string** |  | 
**jobId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSyncJobErrorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** |  | [default to 1]
 **perPage** | **int32** |  | [default to 50]

### Return type

[**GetSyncJobErrors200Response**](GetSyncJobErrors200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListStores

> ListStores200Response ListStores(ctx).Status(status).Execute()

List store connections



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
	status := "status_example" // string | Filter by connection status (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StoreConnectionsAPI.ListStores(context.Background()).Status(status).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StoreConnectionsAPI.ListStores``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListStores`: ListStores200Response
	fmt.Fprintf(os.Stdout, "Response from `StoreConnectionsAPI.ListStores`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListStoresRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **status** | **string** | Filter by connection status | 

### Return type

[**ListStores200Response**](ListStores200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSyncJobs

> ListSyncJobs200Response ListSyncJobs(ctx, storeId).Page(page).PerPage(perPage).Execute()

List sync jobs



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
	storeId := "storeId_example" // string | 
	page := int32(56) // int32 |  (optional) (default to 1)
	perPage := int32(56) // int32 |  (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StoreConnectionsAPI.ListSyncJobs(context.Background(), storeId).Page(page).PerPage(perPage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StoreConnectionsAPI.ListSyncJobs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSyncJobs`: ListSyncJobs200Response
	fmt.Fprintf(os.Stdout, "Response from `StoreConnectionsAPI.ListSyncJobs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSyncJobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** |  | [default to 1]
 **perPage** | **int32** |  | [default to 20]

### Return type

[**ListSyncJobs200Response**](ListSyncJobs200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TriggerSync

> SyncResponse TriggerSync(ctx, storeId).IdempotencyKey(idempotencyKey).Execute()

Trigger product sync



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
	storeId := "storeId_example" // string | 
	idempotencyKey := "idempotencyKey_example" // string | Idempotency key to prevent duplicate syncs (5 min TTL) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StoreConnectionsAPI.TriggerSync(context.Background(), storeId).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StoreConnectionsAPI.TriggerSync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TriggerSync`: SyncResponse
	fmt.Fprintf(os.Stdout, "Response from `StoreConnectionsAPI.TriggerSync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTriggerSyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **idempotencyKey** | **string** | Idempotency key to prevent duplicate syncs (5 min TTL) | 

### Return type

[**SyncResponse**](SyncResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateStore

> CreateStore201Response UpdateStore(ctx, storeId).UpdateStoreRequest(updateStoreRequest).Execute()

Update a store connection



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
	storeId := "storeId_example" // string | 
	updateStoreRequest := *openapiclient.NewUpdateStoreRequest() // UpdateStoreRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StoreConnectionsAPI.UpdateStore(context.Background(), storeId).UpdateStoreRequest(updateStoreRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StoreConnectionsAPI.UpdateStore``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateStore`: CreateStore201Response
	fmt.Fprintf(os.Stdout, "Response from `StoreConnectionsAPI.UpdateStore`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateStoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateStoreRequest** | [**UpdateStoreRequest**](UpdateStoreRequest.md) |  | 

### Return type

[**CreateStore201Response**](CreateStore201Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

