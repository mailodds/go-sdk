# \InboundProcessingAPI

All URIs are relative to *https://api.mailodds.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CorrectInboundMessage**](InboundProcessingAPI.md#CorrectInboundMessage) | **Patch** /v1/inbound-messages/{message_id}/correction | Correct inbound message classification
[**GetBounceStats**](InboundProcessingAPI.md#GetBounceStats) | **Get** /v1/bounce-stats | Get bounce statistics
[**GetBounceStatsSummary**](InboundProcessingAPI.md#GetBounceStatsSummary) | **Get** /v1/bounce-stats/summary | Get bounce statistics summary
[**GetComplaintAssessment**](InboundProcessingAPI.md#GetComplaintAssessment) | **Get** /v1/complaint-assessment | Get complaint assessment
[**GetInboundMessage**](InboundProcessingAPI.md#GetInboundMessage) | **Get** /v1/inbound-messages/{message_id} | Get inbound message
[**ListInboundMessages**](InboundProcessingAPI.md#ListInboundMessages) | **Get** /v1/inbound-messages | List inbound messages



## CorrectInboundMessage

> GetInboundMessage200Response CorrectInboundMessage(ctx, messageId).CorrectInboundMessageRequest(correctInboundMessageRequest).Execute()

Correct inbound message classification



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
	messageId := "messageId_example" // string | 
	correctInboundMessageRequest := *openapiclient.NewCorrectInboundMessageRequest("Correction_example") // CorrectInboundMessageRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InboundProcessingAPI.CorrectInboundMessage(context.Background(), messageId).CorrectInboundMessageRequest(correctInboundMessageRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InboundProcessingAPI.CorrectInboundMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CorrectInboundMessage`: GetInboundMessage200Response
	fmt.Fprintf(os.Stdout, "Response from `InboundProcessingAPI.CorrectInboundMessage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**messageId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCorrectInboundMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **correctInboundMessageRequest** | [**CorrectInboundMessageRequest**](CorrectInboundMessageRequest.md) |  | 

### Return type

[**GetInboundMessage200Response**](GetInboundMessage200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBounceStats

> GetBounceStats200Response GetBounceStats(ctx).DomainId(domainId).Period(period).GroupBy(groupBy).Execute()

Get bounce statistics



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
	domainId := "domainId_example" // string | Filter by sending domain ID (optional)
	period := "period_example" // string | Time period (optional) (default to "7d")
	groupBy := "groupBy_example" // string | Grouping interval (optional) (default to "day")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InboundProcessingAPI.GetBounceStats(context.Background()).DomainId(domainId).Period(period).GroupBy(groupBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InboundProcessingAPI.GetBounceStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBounceStats`: GetBounceStats200Response
	fmt.Fprintf(os.Stdout, "Response from `InboundProcessingAPI.GetBounceStats`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBounceStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **domainId** | **string** | Filter by sending domain ID | 
 **period** | **string** | Time period | [default to &quot;7d&quot;]
 **groupBy** | **string** | Grouping interval | [default to &quot;day&quot;]

### Return type

[**GetBounceStats200Response**](GetBounceStats200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBounceStatsSummary

> GetBounceStatsSummary200Response GetBounceStatsSummary(ctx).DomainId(domainId).Period(period).Execute()

Get bounce statistics summary



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
	domainId := "domainId_example" // string | Filter by sending domain ID (optional)
	period := "period_example" // string | Time period (optional) (default to "30d")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InboundProcessingAPI.GetBounceStatsSummary(context.Background()).DomainId(domainId).Period(period).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InboundProcessingAPI.GetBounceStatsSummary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBounceStatsSummary`: GetBounceStatsSummary200Response
	fmt.Fprintf(os.Stdout, "Response from `InboundProcessingAPI.GetBounceStatsSummary`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBounceStatsSummaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **domainId** | **string** | Filter by sending domain ID | 
 **period** | **string** | Time period | [default to &quot;30d&quot;]

### Return type

[**GetBounceStatsSummary200Response**](GetBounceStatsSummary200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetComplaintAssessment

> GetComplaintAssessment200Response GetComplaintAssessment(ctx).DomainId(domainId).Period(period).Execute()

Get complaint assessment



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
	domainId := "domainId_example" // string | Filter by sending domain ID (optional)
	period := "period_example" // string | Time period (optional) (default to "30d")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InboundProcessingAPI.GetComplaintAssessment(context.Background()).DomainId(domainId).Period(period).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InboundProcessingAPI.GetComplaintAssessment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetComplaintAssessment`: GetComplaintAssessment200Response
	fmt.Fprintf(os.Stdout, "Response from `InboundProcessingAPI.GetComplaintAssessment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetComplaintAssessmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **domainId** | **string** | Filter by sending domain ID | 
 **period** | **string** | Time period | [default to &quot;30d&quot;]

### Return type

[**GetComplaintAssessment200Response**](GetComplaintAssessment200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInboundMessage

> GetInboundMessage200Response GetInboundMessage(ctx, messageId).Execute()

Get inbound message



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
	messageId := "messageId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InboundProcessingAPI.GetInboundMessage(context.Background(), messageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InboundProcessingAPI.GetInboundMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInboundMessage`: GetInboundMessage200Response
	fmt.Fprintf(os.Stdout, "Response from `InboundProcessingAPI.GetInboundMessage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**messageId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInboundMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetInboundMessage200Response**](GetInboundMessage200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListInboundMessages

> ListInboundMessages200Response ListInboundMessages(ctx).Category(category).DomainId(domainId).Since(since).Until(until).IsRead(isRead).Recipient(recipient).Search(search).Page(page).PerPage(perPage).Execute()

List inbound messages



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	category := "category_example" // string | Filter by category (optional)
	domainId := "domainId_example" // string | Filter by sending domain ID (optional)
	since := time.Now() // time.Time | Start date (ISO 8601) (optional)
	until := time.Now() // time.Time | End date (ISO 8601) (optional)
	isRead := true // bool | Filter by read status (optional)
	recipient := "recipient_example" // string | Filter by original recipient (optional)
	search := "search_example" // string | Search in subject and body (optional)
	page := int32(56) // int32 |  (optional) (default to 1)
	perPage := int32(56) // int32 |  (optional) (default to 50)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InboundProcessingAPI.ListInboundMessages(context.Background()).Category(category).DomainId(domainId).Since(since).Until(until).IsRead(isRead).Recipient(recipient).Search(search).Page(page).PerPage(perPage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InboundProcessingAPI.ListInboundMessages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListInboundMessages`: ListInboundMessages200Response
	fmt.Fprintf(os.Stdout, "Response from `InboundProcessingAPI.ListInboundMessages`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListInboundMessagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **category** | **string** | Filter by category | 
 **domainId** | **string** | Filter by sending domain ID | 
 **since** | **time.Time** | Start date (ISO 8601) | 
 **until** | **time.Time** | End date (ISO 8601) | 
 **isRead** | **bool** | Filter by read status | 
 **recipient** | **string** | Filter by original recipient | 
 **search** | **string** | Search in subject and body | 
 **page** | **int32** |  | [default to 1]
 **perPage** | **int32** |  | [default to 50]

### Return type

[**ListInboundMessages200Response**](ListInboundMessages200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

