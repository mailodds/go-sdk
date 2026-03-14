# \ServerTestsAPI

All URIs are relative to *https://api.mailodds.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetServerTest**](ServerTestsAPI.md#GetServerTest) | **Get** /v1/server-tests/{test_id} | Get server test
[**ListServerTests**](ServerTestsAPI.md#ListServerTests) | **Get** /v1/server-tests | List server tests
[**RunServerTest**](ServerTestsAPI.md#RunServerTest) | **Post** /v1/server-tests | Run server test



## GetServerTest

> RunServerTest201Response GetServerTest(ctx, testId).Execute()

Get server test



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
	testId := "testId_example" // string | Server test UUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServerTestsAPI.GetServerTest(context.Background(), testId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerTestsAPI.GetServerTest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServerTest`: RunServerTest201Response
	fmt.Fprintf(os.Stdout, "Response from `ServerTestsAPI.GetServerTest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**testId** | **string** | Server test UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServerTestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RunServerTest201Response**](RunServerTest201Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListServerTests

> ListServerTests200Response ListServerTests(ctx).Page(page).PerPage(perPage).Execute()

List server tests



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
	page := int32(56) // int32 |  (optional) (default to 1)
	perPage := int32(56) // int32 |  (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServerTestsAPI.ListServerTests(context.Background()).Page(page).PerPage(perPage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerTestsAPI.ListServerTests``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListServerTests`: ListServerTests200Response
	fmt.Fprintf(os.Stdout, "Response from `ServerTestsAPI.ListServerTests`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListServerTestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | [default to 1]
 **perPage** | **int32** |  | [default to 20]

### Return type

[**ListServerTests200Response**](ListServerTests200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RunServerTest

> RunServerTest201Response RunServerTest(ctx).RunServerTestRequest(runServerTestRequest).Execute()

Run server test



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
	runServerTestRequest := *openapiclient.NewRunServerTestRequest("example.com") // RunServerTestRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServerTestsAPI.RunServerTest(context.Background()).RunServerTestRequest(runServerTestRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerTestsAPI.RunServerTest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RunServerTest`: RunServerTest201Response
	fmt.Fprintf(os.Stdout, "Response from `ServerTestsAPI.RunServerTest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRunServerTestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **runServerTestRequest** | [**RunServerTestRequest**](RunServerTestRequest.md) |  | 

### Return type

[**RunServerTest201Response**](RunServerTest201Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

