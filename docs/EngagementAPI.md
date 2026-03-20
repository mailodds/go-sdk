# \EngagementAPI

All URIs are relative to *https://api.mailodds.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDisengagedContacts**](EngagementAPI.md#GetDisengagedContacts) | **Get** /v1/engagement/disengaged | List disengaged contacts
[**GetEngagementScore**](EngagementAPI.md#GetEngagementScore) | **Get** /v1/engagement/score/{email} | Get engagement score
[**GetEngagementSummary**](EngagementAPI.md#GetEngagementSummary) | **Get** /v1/engagement/summary | Get engagement summary
[**SuppressDisengaged**](EngagementAPI.md#SuppressDisengaged) | **Post** /v1/engagement/suppress-disengaged | Suppress disengaged contacts



## GetDisengagedContacts

> GetDisengagedContacts200Response GetDisengagedContacts(ctx).InactiveDays(inactiveDays).MinSends(minSends).DomainId(domainId).Page(page).PerPage(perPage).Execute()

List disengaged contacts



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
	inactiveDays := int32(56) // int32 | Days of inactivity (optional) (default to 90)
	minSends := int32(56) // int32 | Minimum emails sent to qualify (optional) (default to 5)
	domainId := "domainId_example" // string | Filter by sending domain ID (optional)
	page := int32(56) // int32 |  (optional) (default to 1)
	perPage := int32(56) // int32 |  (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EngagementAPI.GetDisengagedContacts(context.Background()).InactiveDays(inactiveDays).MinSends(minSends).DomainId(domainId).Page(page).PerPage(perPage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EngagementAPI.GetDisengagedContacts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDisengagedContacts`: GetDisengagedContacts200Response
	fmt.Fprintf(os.Stdout, "Response from `EngagementAPI.GetDisengagedContacts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDisengagedContactsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inactiveDays** | **int32** | Days of inactivity | [default to 90]
 **minSends** | **int32** | Minimum emails sent to qualify | [default to 5]
 **domainId** | **string** | Filter by sending domain ID | 
 **page** | **int32** |  | [default to 1]
 **perPage** | **int32** |  | [default to 100]

### Return type

[**GetDisengagedContacts200Response**](GetDisengagedContacts200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEngagementScore

> GetEngagementScore200Response GetEngagementScore(ctx, email).Execute()

Get engagement score



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
	email := "email_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EngagementAPI.GetEngagementScore(context.Background(), email).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EngagementAPI.GetEngagementScore``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEngagementScore`: GetEngagementScore200Response
	fmt.Fprintf(os.Stdout, "Response from `EngagementAPI.GetEngagementScore`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**email** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEngagementScoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetEngagementScore200Response**](GetEngagementScore200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEngagementSummary

> GetBounceStatsSummary200Response GetEngagementSummary(ctx).DomainId(domainId).Execute()

Get engagement summary



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EngagementAPI.GetEngagementSummary(context.Background()).DomainId(domainId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EngagementAPI.GetEngagementSummary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEngagementSummary`: GetBounceStatsSummary200Response
	fmt.Fprintf(os.Stdout, "Response from `EngagementAPI.GetEngagementSummary`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEngagementSummaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **domainId** | **string** | Filter by sending domain ID | 

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


## SuppressDisengaged

> SuppressDisengaged200Response SuppressDisengaged(ctx).SuppressDisengagedRequest(suppressDisengagedRequest).Execute()

Suppress disengaged contacts



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
	suppressDisengagedRequest := *openapiclient.NewSuppressDisengagedRequest() // SuppressDisengagedRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EngagementAPI.SuppressDisengaged(context.Background()).SuppressDisengagedRequest(suppressDisengagedRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EngagementAPI.SuppressDisengaged``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SuppressDisengaged`: SuppressDisengaged200Response
	fmt.Fprintf(os.Stdout, "Response from `EngagementAPI.SuppressDisengaged`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSuppressDisengagedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **suppressDisengagedRequest** | [**SuppressDisengagedRequest**](SuppressDisengagedRequest.md) |  | 

### Return type

[**SuppressDisengaged200Response**](SuppressDisengaged200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

