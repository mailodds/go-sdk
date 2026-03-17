# \OutOfOfficeAPI

All URIs are relative to *https://api.mailodds.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BatchCheckOoo**](OutOfOfficeAPI.md#BatchCheckOoo) | **Post** /v1/out-of-office/batch-check | Batch check OOO status
[**DeleteOooContact**](OutOfOfficeAPI.md#DeleteOooContact) | **Delete** /v1/out-of-office/{email} | Delete OOO contact
[**GetOooStatus**](OutOfOfficeAPI.md#GetOooStatus) | **Get** /v1/out-of-office/{email}/status | Get OOO status for email
[**ListOooContacts**](OutOfOfficeAPI.md#ListOooContacts) | **Get** /v1/out-of-office | List out-of-office contacts
[**UpdateOooContact**](OutOfOfficeAPI.md#UpdateOooContact) | **Patch** /v1/out-of-office/{email} | Update OOO contact



## BatchCheckOoo

> BatchCheckOoo200Response BatchCheckOoo(ctx).BatchCheckOooRequest(batchCheckOooRequest).Execute()

Batch check OOO status



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
	batchCheckOooRequest := *openapiclient.NewBatchCheckOooRequest([]string{"Emails_example"}) // BatchCheckOooRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OutOfOfficeAPI.BatchCheckOoo(context.Background()).BatchCheckOooRequest(batchCheckOooRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OutOfOfficeAPI.BatchCheckOoo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BatchCheckOoo`: BatchCheckOoo200Response
	fmt.Fprintf(os.Stdout, "Response from `OutOfOfficeAPI.BatchCheckOoo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBatchCheckOooRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batchCheckOooRequest** | [**BatchCheckOooRequest**](BatchCheckOooRequest.md) |  | 

### Return type

[**BatchCheckOoo200Response**](BatchCheckOoo200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOooContact

> DeleteOooContact200Response DeleteOooContact(ctx, email).Execute()

Delete OOO contact



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
	email := "email_example" // string | Email address

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OutOfOfficeAPI.DeleteOooContact(context.Background(), email).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OutOfOfficeAPI.DeleteOooContact``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteOooContact`: DeleteOooContact200Response
	fmt.Fprintf(os.Stdout, "Response from `OutOfOfficeAPI.DeleteOooContact`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**email** | **string** | Email address | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOooContactRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeleteOooContact200Response**](DeleteOooContact200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOooStatus

> GetOooStatus200Response GetOooStatus(ctx, email).Execute()

Get OOO status for email



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
	email := "email_example" // string | Email address to check

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OutOfOfficeAPI.GetOooStatus(context.Background(), email).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OutOfOfficeAPI.GetOooStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOooStatus`: GetOooStatus200Response
	fmt.Fprintf(os.Stdout, "Response from `OutOfOfficeAPI.GetOooStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**email** | **string** | Email address to check | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOooStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetOooStatus200Response**](GetOooStatus200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOooContacts

> ListOooContacts200Response ListOooContacts(ctx).ActiveOnly(activeOnly).Page(page).PerPage(perPage).Execute()

List out-of-office contacts



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
	activeOnly := true // bool | Only return currently active OOO contacts (optional) (default to true)
	page := int32(56) // int32 |  (optional) (default to 1)
	perPage := int32(56) // int32 |  (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OutOfOfficeAPI.ListOooContacts(context.Background()).ActiveOnly(activeOnly).Page(page).PerPage(perPage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OutOfOfficeAPI.ListOooContacts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOooContacts`: ListOooContacts200Response
	fmt.Fprintf(os.Stdout, "Response from `OutOfOfficeAPI.ListOooContacts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListOooContactsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **activeOnly** | **bool** | Only return currently active OOO contacts | [default to true]
 **page** | **int32** |  | [default to 1]
 **perPage** | **int32** |  | [default to 100]

### Return type

[**ListOooContacts200Response**](ListOooContacts200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOooContact

> map[string]interface{} UpdateOooContact(ctx, email).UpdateOooContactRequest(updateOooContactRequest).Execute()

Update OOO contact



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
	email := "email_example" // string | Email address
	updateOooContactRequest := *openapiclient.NewUpdateOooContactRequest() // UpdateOooContactRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OutOfOfficeAPI.UpdateOooContact(context.Background(), email).UpdateOooContactRequest(updateOooContactRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OutOfOfficeAPI.UpdateOooContact``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOooContact`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `OutOfOfficeAPI.UpdateOooContact`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**email** | **string** | Email address | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOooContactRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateOooContactRequest** | [**UpdateOooContactRequest**](UpdateOooContactRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

