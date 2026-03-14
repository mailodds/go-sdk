# \ContactListsAPI

All URIs are relative to *https://api.mailodds.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AppendToContactList**](ContactListsAPI.md#AppendToContactList) | **Post** /v1/contact-lists/{list_id}/append | Append to contact list
[**CreateContactList**](ContactListsAPI.md#CreateContactList) | **Post** /v1/contact-lists | Create contact list
[**GetInactiveContactsReport**](ContactListsAPI.md#GetInactiveContactsReport) | **Get** /v1/contacts/inactive-report | Get inactive contacts report
[**ListContactLists**](ContactListsAPI.md#ListContactLists) | **Get** /v1/contact-lists | List contact lists
[**QueryContactList**](ContactListsAPI.md#QueryContactList) | **Post** /v1/contact-lists/{list_id}/query | Query contact list



## AppendToContactList

> AppendToContactList200Response AppendToContactList(ctx, listId).AppendToContactListRequest(appendToContactListRequest).Execute()

Append to contact list



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
	listId := "listId_example" // string | Contact list UUID
	appendToContactListRequest := *openapiclient.NewAppendToContactListRequest([]string{"SourceJobIds_example"}) // AppendToContactListRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContactListsAPI.AppendToContactList(context.Background(), listId).AppendToContactListRequest(appendToContactListRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContactListsAPI.AppendToContactList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppendToContactList`: AppendToContactList200Response
	fmt.Fprintf(os.Stdout, "Response from `ContactListsAPI.AppendToContactList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listId** | **string** | Contact list UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppendToContactListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **appendToContactListRequest** | [**AppendToContactListRequest**](AppendToContactListRequest.md) |  | 

### Return type

[**AppendToContactList200Response**](AppendToContactList200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateContactList

> CreateContactList201Response CreateContactList(ctx).CreateContactListRequest(createContactListRequest).Execute()

Create contact list



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
	createContactListRequest := *openapiclient.NewCreateContactListRequest("Name_example") // CreateContactListRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContactListsAPI.CreateContactList(context.Background()).CreateContactListRequest(createContactListRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContactListsAPI.CreateContactList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateContactList`: CreateContactList201Response
	fmt.Fprintf(os.Stdout, "Response from `ContactListsAPI.CreateContactList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateContactListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createContactListRequest** | [**CreateContactListRequest**](CreateContactListRequest.md) |  | 

### Return type

[**CreateContactList201Response**](CreateContactList201Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInactiveContactsReport

> GetInactiveContactsReport200Response GetInactiveContactsReport(ctx).Days(days).Execute()

Get inactive contacts report



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
	days := int32(56) // int32 | Inactivity threshold in days (optional) (default to 90)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContactListsAPI.GetInactiveContactsReport(context.Background()).Days(days).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContactListsAPI.GetInactiveContactsReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInactiveContactsReport`: GetInactiveContactsReport200Response
	fmt.Fprintf(os.Stdout, "Response from `ContactListsAPI.GetInactiveContactsReport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetInactiveContactsReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **days** | **int32** | Inactivity threshold in days | [default to 90]

### Return type

[**GetInactiveContactsReport200Response**](GetInactiveContactsReport200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListContactLists

> ListContactLists200Response ListContactLists(ctx).Page(page).PerPage(perPage).Execute()

List contact lists



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
	resp, r, err := apiClient.ContactListsAPI.ListContactLists(context.Background()).Page(page).PerPage(perPage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContactListsAPI.ListContactLists``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListContactLists`: ListContactLists200Response
	fmt.Fprintf(os.Stdout, "Response from `ContactListsAPI.ListContactLists`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListContactListsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | [default to 1]
 **perPage** | **int32** |  | [default to 20]

### Return type

[**ListContactLists200Response**](ListContactLists200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QueryContactList

> QueryContactList200Response QueryContactList(ctx, listId).QueryContactListRequest(queryContactListRequest).Execute()

Query contact list



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
	listId := "listId_example" // string | Contact list UUID
	queryContactListRequest := *openapiclient.NewQueryContactListRequest() // QueryContactListRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContactListsAPI.QueryContactList(context.Background(), listId).QueryContactListRequest(queryContactListRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContactListsAPI.QueryContactList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QueryContactList`: QueryContactList200Response
	fmt.Fprintf(os.Stdout, "Response from `ContactListsAPI.QueryContactList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listId** | **string** | Contact list UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiQueryContactListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **queryContactListRequest** | [**QueryContactListRequest**](QueryContactListRequest.md) |  | 

### Return type

[**QueryContactList200Response**](QueryContactList200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

