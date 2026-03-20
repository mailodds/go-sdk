# \SpamChecksAPI

All URIs are relative to *https://api.mailodds.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSpamCheck**](SpamChecksAPI.md#DeleteSpamCheck) | **Delete** /v1/spam-checks/{check_id} | Delete spam check
[**GetSpamCheck**](SpamChecksAPI.md#GetSpamCheck) | **Get** /v1/spam-checks/{check_id} | Get spam check
[**ListSpamChecks**](SpamChecksAPI.md#ListSpamChecks) | **Get** /v1/spam-checks | List spam checks
[**RunSpamCheck**](SpamChecksAPI.md#RunSpamCheck) | **Post** /v1/spam-checks | Run spam check



## DeleteSpamCheck

> DeletePolicyRule200Response DeleteSpamCheck(ctx, checkId).Execute()

Delete spam check



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
	checkId := "checkId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SpamChecksAPI.DeleteSpamCheck(context.Background(), checkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpamChecksAPI.DeleteSpamCheck``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteSpamCheck`: DeletePolicyRule200Response
	fmt.Fprintf(os.Stdout, "Response from `SpamChecksAPI.DeleteSpamCheck`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**checkId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSpamCheckRequest struct via the builder pattern


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


## GetSpamCheck

> RunSpamCheck201Response GetSpamCheck(ctx, checkId).Execute()

Get spam check



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
	checkId := "checkId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SpamChecksAPI.GetSpamCheck(context.Background(), checkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpamChecksAPI.GetSpamCheck``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSpamCheck`: RunSpamCheck201Response
	fmt.Fprintf(os.Stdout, "Response from `SpamChecksAPI.GetSpamCheck`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**checkId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSpamCheckRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RunSpamCheck201Response**](RunSpamCheck201Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSpamChecks

> ListSpamChecks200Response ListSpamChecks(ctx).Page(page).PerPage(perPage).Execute()

List spam checks



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
	page := int32(56) // int32 |  (optional) (default to 1)
	perPage := int32(56) // int32 |  (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SpamChecksAPI.ListSpamChecks(context.Background()).Page(page).PerPage(perPage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpamChecksAPI.ListSpamChecks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSpamChecks`: ListSpamChecks200Response
	fmt.Fprintf(os.Stdout, "Response from `SpamChecksAPI.ListSpamChecks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListSpamChecksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | [default to 1]
 **perPage** | **int32** |  | [default to 20]

### Return type

[**ListSpamChecks200Response**](ListSpamChecks200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RunSpamCheck

> RunSpamCheck201Response RunSpamCheck(ctx).RunSpamCheckRequest(runSpamCheckRequest).Execute()

Run spam check



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
	runSpamCheckRequest := *openapiclient.NewRunSpamCheckRequest("FromDomain_example") // RunSpamCheckRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SpamChecksAPI.RunSpamCheck(context.Background()).RunSpamCheckRequest(runSpamCheckRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpamChecksAPI.RunSpamCheck``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RunSpamCheck`: RunSpamCheck201Response
	fmt.Fprintf(os.Stdout, "Response from `SpamChecksAPI.RunSpamCheck`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRunSpamCheckRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **runSpamCheckRequest** | [**RunSpamCheckRequest**](RunSpamCheckRequest.md) |  | 

### Return type

[**RunSpamCheck201Response**](RunSpamCheck201Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

