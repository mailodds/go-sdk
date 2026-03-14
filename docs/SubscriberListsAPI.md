# \SubscriberListsAPI

All URIs are relative to *https://api.mailodds.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConfirmSubscription**](SubscriberListsAPI.md#ConfirmSubscription) | **Get** /v1/confirm/{token} | Confirm subscription
[**CreateList**](SubscriberListsAPI.md#CreateList) | **Post** /v1/lists | Create a subscriber list
[**DeleteList**](SubscriberListsAPI.md#DeleteList) | **Delete** /v1/lists/{list_id} | Delete a subscriber list
[**GetList**](SubscriberListsAPI.md#GetList) | **Get** /v1/lists/{list_id} | Get a subscriber list
[**GetLists**](SubscriberListsAPI.md#GetLists) | **Get** /v1/lists | List subscriber lists
[**GetSubscribers**](SubscriberListsAPI.md#GetSubscribers) | **Get** /v1/lists/{list_id}/subscribers | List subscribers
[**Subscribe**](SubscriberListsAPI.md#Subscribe) | **Post** /v1/subscribe/{list_id} | Subscribe to a list
[**UnsubscribeSubscriber**](SubscriberListsAPI.md#UnsubscribeSubscriber) | **Delete** /v1/lists/{list_id}/subscribers/{subscriber_id} | Unsubscribe a subscriber



## ConfirmSubscription

> ConfirmSubscription200Response ConfirmSubscription(ctx, token).Execute()

Confirm subscription



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
	token := "token_example" // string | Confirmation token from email

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriberListsAPI.ConfirmSubscription(context.Background(), token).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriberListsAPI.ConfirmSubscription``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConfirmSubscription`: ConfirmSubscription200Response
	fmt.Fprintf(os.Stdout, "Response from `SubscriberListsAPI.ConfirmSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** | Confirmation token from email | 

### Other Parameters

Other parameters are passed through a pointer to a apiConfirmSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ConfirmSubscription200Response**](ConfirmSubscription200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateList

> CreateList201Response CreateList(ctx).CreateListRequest(createListRequest).Execute()

Create a subscriber list



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
	createListRequest := *openapiclient.NewCreateListRequest("Name_example") // CreateListRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriberListsAPI.CreateList(context.Background()).CreateListRequest(createListRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriberListsAPI.CreateList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateList`: CreateList201Response
	fmt.Fprintf(os.Stdout, "Response from `SubscriberListsAPI.CreateList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createListRequest** | [**CreateListRequest**](CreateListRequest.md) |  | 

### Return type

[**CreateList201Response**](CreateList201Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteList

> DeletePolicyRule200Response DeleteList(ctx, listId).Execute()

Delete a subscriber list



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
	listId := "listId_example" // string | List UUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriberListsAPI.DeleteList(context.Background(), listId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriberListsAPI.DeleteList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteList`: DeletePolicyRule200Response
	fmt.Fprintf(os.Stdout, "Response from `SubscriberListsAPI.DeleteList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listId** | **string** | List UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteListRequest struct via the builder pattern


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


## GetList

> CreateList201Response GetList(ctx, listId).Execute()

Get a subscriber list



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
	listId := "listId_example" // string | List UUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriberListsAPI.GetList(context.Background(), listId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriberListsAPI.GetList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetList`: CreateList201Response
	fmt.Fprintf(os.Stdout, "Response from `SubscriberListsAPI.GetList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listId** | **string** | List UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreateList201Response**](CreateList201Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLists

> GetLists200Response GetLists(ctx).Page(page).PerPage(perPage).Execute()

List subscriber lists



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
	page := int32(56) // int32 | Page number (optional) (default to 1)
	perPage := int32(56) // int32 | Items per page (optional) (default to 25)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriberListsAPI.GetLists(context.Background()).Page(page).PerPage(perPage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriberListsAPI.GetLists``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLists`: GetLists200Response
	fmt.Fprintf(os.Stdout, "Response from `SubscriberListsAPI.GetLists`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetListsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page number | [default to 1]
 **perPage** | **int32** | Items per page | [default to 25]

### Return type

[**GetLists200Response**](GetLists200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSubscribers

> GetSubscribers200Response GetSubscribers(ctx, listId).Page(page).PerPage(perPage).Status(status).Execute()

List subscribers



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
	listId := "listId_example" // string | List UUID
	page := int32(56) // int32 | Page number (optional) (default to 1)
	perPage := int32(56) // int32 | Items per page (optional) (default to 50)
	status := "status_example" // string | Filter by subscriber status (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriberListsAPI.GetSubscribers(context.Background(), listId).Page(page).PerPage(perPage).Status(status).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriberListsAPI.GetSubscribers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSubscribers`: GetSubscribers200Response
	fmt.Fprintf(os.Stdout, "Response from `SubscriberListsAPI.GetSubscribers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listId** | **string** | List UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSubscribersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Page number | [default to 1]
 **perPage** | **int32** | Items per page | [default to 50]
 **status** | **string** | Filter by subscriber status | 

### Return type

[**GetSubscribers200Response**](GetSubscribers200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Subscribe

> UnsubscribeSubscriber200Response Subscribe(ctx, listId).SubscribeRequest(subscribeRequest).Execute()

Subscribe to a list



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
	listId := "listId_example" // string | List UUID
	subscribeRequest := *openapiclient.NewSubscribeRequest("Email_example") // SubscribeRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriberListsAPI.Subscribe(context.Background(), listId).SubscribeRequest(subscribeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriberListsAPI.Subscribe``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Subscribe`: UnsubscribeSubscriber200Response
	fmt.Fprintf(os.Stdout, "Response from `SubscriberListsAPI.Subscribe`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listId** | **string** | List UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubscribeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **subscribeRequest** | [**SubscribeRequest**](SubscribeRequest.md) |  | 

### Return type

[**UnsubscribeSubscriber200Response**](UnsubscribeSubscriber200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnsubscribeSubscriber

> UnsubscribeSubscriber200Response UnsubscribeSubscriber(ctx, listId, subscriberId).Execute()

Unsubscribe a subscriber



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
	listId := "listId_example" // string | List UUID
	subscriberId := "subscriberId_example" // string | Subscriber UUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriberListsAPI.UnsubscribeSubscriber(context.Background(), listId, subscriberId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriberListsAPI.UnsubscribeSubscriber``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UnsubscribeSubscriber`: UnsubscribeSubscriber200Response
	fmt.Fprintf(os.Stdout, "Response from `SubscriberListsAPI.UnsubscribeSubscriber`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listId** | **string** | List UUID | 
**subscriberId** | **string** | Subscriber UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnsubscribeSubscriberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**UnsubscribeSubscriber200Response**](UnsubscribeSubscriber200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

