# \SuppressionListsAPI

All URIs are relative to *https://api.mailodds.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddSuppression**](SuppressionListsAPI.md#AddSuppression) | **Post** /v1/suppression | Add suppression entries
[**CheckSuppression**](SuppressionListsAPI.md#CheckSuppression) | **Post** /v1/suppression/check | Check suppression status
[**GetSuppressionStats**](SuppressionListsAPI.md#GetSuppressionStats) | **Get** /v1/suppression/stats | Get suppression statistics
[**ListSuppression**](SuppressionListsAPI.md#ListSuppression) | **Get** /v1/suppression | List suppression entries
[**RemoveSuppression**](SuppressionListsAPI.md#RemoveSuppression) | **Delete** /v1/suppression | Remove suppression entries



## AddSuppression

> AddSuppressionResponse AddSuppression(ctx).AddSuppressionRequest(addSuppressionRequest).Execute()

Add suppression entries



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mailodds/go-sdk"
)

func main() {
	addSuppressionRequest := *openapiclient.NewAddSuppressionRequest([]openapiclient.AddSuppressionRequestEntriesInner{*openapiclient.NewAddSuppressionRequestEntriesInner("Type_example", "Value_example")}) // AddSuppressionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SuppressionListsAPI.AddSuppression(context.Background()).AddSuppressionRequest(addSuppressionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SuppressionListsAPI.AddSuppression``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddSuppression`: AddSuppressionResponse
	fmt.Fprintf(os.Stdout, "Response from `SuppressionListsAPI.AddSuppression`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddSuppressionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addSuppressionRequest** | [**AddSuppressionRequest**](AddSuppressionRequest.md) |  | 

### Return type

[**AddSuppressionResponse**](AddSuppressionResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CheckSuppression

> SuppressionCheckResponse CheckSuppression(ctx).CheckSuppressionRequest(checkSuppressionRequest).Execute()

Check suppression status



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mailodds/go-sdk"
)

func main() {
	checkSuppressionRequest := *openapiclient.NewCheckSuppressionRequest("Email_example") // CheckSuppressionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SuppressionListsAPI.CheckSuppression(context.Background()).CheckSuppressionRequest(checkSuppressionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SuppressionListsAPI.CheckSuppression``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CheckSuppression`: SuppressionCheckResponse
	fmt.Fprintf(os.Stdout, "Response from `SuppressionListsAPI.CheckSuppression`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckSuppressionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **checkSuppressionRequest** | [**CheckSuppressionRequest**](CheckSuppressionRequest.md) |  | 

### Return type

[**SuppressionCheckResponse**](SuppressionCheckResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSuppressionStats

> SuppressionStatsResponse GetSuppressionStats(ctx).Execute()

Get suppression statistics



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mailodds/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SuppressionListsAPI.GetSuppressionStats(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SuppressionListsAPI.GetSuppressionStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSuppressionStats`: SuppressionStatsResponse
	fmt.Fprintf(os.Stdout, "Response from `SuppressionListsAPI.GetSuppressionStats`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSuppressionStatsRequest struct via the builder pattern


### Return type

[**SuppressionStatsResponse**](SuppressionStatsResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSuppression

> SuppressionListResponse ListSuppression(ctx).Page(page).PerPage(perPage).Type_(type_).Search(search).Execute()

List suppression entries



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mailodds/go-sdk"
)

func main() {
	page := int32(56) // int32 |  (optional) (default to 1)
	perPage := int32(56) // int32 |  (optional) (default to 50)
	type_ := "type__example" // string |  (optional)
	search := "search_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SuppressionListsAPI.ListSuppression(context.Background()).Page(page).PerPage(perPage).Type_(type_).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SuppressionListsAPI.ListSuppression``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSuppression`: SuppressionListResponse
	fmt.Fprintf(os.Stdout, "Response from `SuppressionListsAPI.ListSuppression`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListSuppressionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | [default to 1]
 **perPage** | **int32** |  | [default to 50]
 **type_** | **string** |  | 
 **search** | **string** |  | 

### Return type

[**SuppressionListResponse**](SuppressionListResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveSuppression

> RemoveSuppression200Response RemoveSuppression(ctx).RemoveSuppressionRequest(removeSuppressionRequest).Execute()

Remove suppression entries



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mailodds/go-sdk"
)

func main() {
	removeSuppressionRequest := *openapiclient.NewRemoveSuppressionRequest([]string{"Entries_example"}) // RemoveSuppressionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SuppressionListsAPI.RemoveSuppression(context.Background()).RemoveSuppressionRequest(removeSuppressionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SuppressionListsAPI.RemoveSuppression``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveSuppression`: RemoveSuppression200Response
	fmt.Fprintf(os.Stdout, "Response from `SuppressionListsAPI.RemoveSuppression`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRemoveSuppressionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **removeSuppressionRequest** | [**RemoveSuppressionRequest**](RemoveSuppressionRequest.md) |  | 

### Return type

[**RemoveSuppression200Response**](RemoveSuppression200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

