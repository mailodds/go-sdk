# \SendingDomainsAPI

All URIs are relative to *https://api.mailodds.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSendingDomain**](SendingDomainsAPI.md#CreateSendingDomain) | **Post** /v1/sending-domains | Add a sending domain
[**DeleteSendingDomain**](SendingDomainsAPI.md#DeleteSendingDomain) | **Delete** /v1/sending-domains/{domain_id} | Delete a sending domain
[**GetReplyForwarding**](SendingDomainsAPI.md#GetReplyForwarding) | **Get** /v1/sending-domains/{domain_id}/reply-forwarding | Get reply forwarding config
[**GetSendingDomain**](SendingDomainsAPI.md#GetSendingDomain) | **Get** /v1/sending-domains/{domain_id} | Get a sending domain
[**GetSendingDomainIdentityScore**](SendingDomainsAPI.md#GetSendingDomainIdentityScore) | **Get** /v1/sending-domains/{domain_id}/identity-score | Get domain identity score
[**GetSendingStats**](SendingDomainsAPI.md#GetSendingStats) | **Get** /v1/sending-stats | Get sending statistics
[**ListSendingDomains**](SendingDomainsAPI.md#ListSendingDomains) | **Get** /v1/sending-domains | List sending domains
[**UpdateReplyForwarding**](SendingDomainsAPI.md#UpdateReplyForwarding) | **Patch** /v1/sending-domains/{domain_id}/reply-forwarding | Update reply forwarding config
[**VerifySendingDomain**](SendingDomainsAPI.md#VerifySendingDomain) | **Post** /v1/sending-domains/{domain_id}/verify | Verify domain DNS records



## CreateSendingDomain

> CreateSendingDomain201Response CreateSendingDomain(ctx).CreateSendingDomainRequest(createSendingDomainRequest).Execute()

Add a sending domain



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
	createSendingDomainRequest := *openapiclient.NewCreateSendingDomainRequest("example.com") // CreateSendingDomainRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SendingDomainsAPI.CreateSendingDomain(context.Background()).CreateSendingDomainRequest(createSendingDomainRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SendingDomainsAPI.CreateSendingDomain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSendingDomain`: CreateSendingDomain201Response
	fmt.Fprintf(os.Stdout, "Response from `SendingDomainsAPI.CreateSendingDomain`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSendingDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createSendingDomainRequest** | [**CreateSendingDomainRequest**](CreateSendingDomainRequest.md) |  | 

### Return type

[**CreateSendingDomain201Response**](CreateSendingDomain201Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSendingDomain

> DeletePolicyRule200Response DeleteSendingDomain(ctx, domainId).Execute()

Delete a sending domain



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
	domainId := "domainId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SendingDomainsAPI.DeleteSendingDomain(context.Background(), domainId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SendingDomainsAPI.DeleteSendingDomain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteSendingDomain`: DeletePolicyRule200Response
	fmt.Fprintf(os.Stdout, "Response from `SendingDomainsAPI.DeleteSendingDomain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSendingDomainRequest struct via the builder pattern


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


## GetReplyForwarding

> GetReplyForwarding200Response GetReplyForwarding(ctx, domainId).Execute()

Get reply forwarding config



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
	domainId := "domainId_example" // string | Sending domain ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SendingDomainsAPI.GetReplyForwarding(context.Background(), domainId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SendingDomainsAPI.GetReplyForwarding``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetReplyForwarding`: GetReplyForwarding200Response
	fmt.Fprintf(os.Stdout, "Response from `SendingDomainsAPI.GetReplyForwarding`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainId** | **string** | Sending domain ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReplyForwardingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetReplyForwarding200Response**](GetReplyForwarding200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSendingDomain

> CreateSendingDomain201Response GetSendingDomain(ctx, domainId).Execute()

Get a sending domain



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
	domainId := "domainId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SendingDomainsAPI.GetSendingDomain(context.Background(), domainId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SendingDomainsAPI.GetSendingDomain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSendingDomain`: CreateSendingDomain201Response
	fmt.Fprintf(os.Stdout, "Response from `SendingDomainsAPI.GetSendingDomain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSendingDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreateSendingDomain201Response**](CreateSendingDomain201Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSendingDomainIdentityScore

> GetSendingDomainIdentityScore200Response GetSendingDomainIdentityScore(ctx, domainId).Execute()

Get domain identity score



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
	domainId := "domainId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SendingDomainsAPI.GetSendingDomainIdentityScore(context.Background(), domainId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SendingDomainsAPI.GetSendingDomainIdentityScore``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSendingDomainIdentityScore`: GetSendingDomainIdentityScore200Response
	fmt.Fprintf(os.Stdout, "Response from `SendingDomainsAPI.GetSendingDomainIdentityScore`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSendingDomainIdentityScoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetSendingDomainIdentityScore200Response**](GetSendingDomainIdentityScore200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSendingStats

> GetSendingStats200Response GetSendingStats(ctx).Period(period).DomainId(domainId).Execute()

Get sending statistics



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
	period := "period_example" // string | Time period (optional) (default to "7d")
	domainId := "domainId_example" // string | Filter by domain (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SendingDomainsAPI.GetSendingStats(context.Background()).Period(period).DomainId(domainId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SendingDomainsAPI.GetSendingStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSendingStats`: GetSendingStats200Response
	fmt.Fprintf(os.Stdout, "Response from `SendingDomainsAPI.GetSendingStats`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSendingStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **period** | **string** | Time period | [default to &quot;7d&quot;]
 **domainId** | **string** | Filter by domain | 

### Return type

[**GetSendingStats200Response**](GetSendingStats200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSendingDomains

> ListSendingDomains200Response ListSendingDomains(ctx).Execute()

List sending domains



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
	resp, r, err := apiClient.SendingDomainsAPI.ListSendingDomains(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SendingDomainsAPI.ListSendingDomains``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSendingDomains`: ListSendingDomains200Response
	fmt.Fprintf(os.Stdout, "Response from `SendingDomainsAPI.ListSendingDomains`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListSendingDomainsRequest struct via the builder pattern


### Return type

[**ListSendingDomains200Response**](ListSendingDomains200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateReplyForwarding

> GetReplyForwarding200Response UpdateReplyForwarding(ctx, domainId).UpdateReplyForwardingRequest(updateReplyForwardingRequest).Execute()

Update reply forwarding config



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
	domainId := "domainId_example" // string | Sending domain ID
	updateReplyForwardingRequest := *openapiclient.NewUpdateReplyForwardingRequest() // UpdateReplyForwardingRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SendingDomainsAPI.UpdateReplyForwarding(context.Background(), domainId).UpdateReplyForwardingRequest(updateReplyForwardingRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SendingDomainsAPI.UpdateReplyForwarding``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateReplyForwarding`: GetReplyForwarding200Response
	fmt.Fprintf(os.Stdout, "Response from `SendingDomainsAPI.UpdateReplyForwarding`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainId** | **string** | Sending domain ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateReplyForwardingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateReplyForwardingRequest** | [**UpdateReplyForwardingRequest**](UpdateReplyForwardingRequest.md) |  | 

### Return type

[**GetReplyForwarding200Response**](GetReplyForwarding200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VerifySendingDomain

> CreateSendingDomain201Response VerifySendingDomain(ctx, domainId).Execute()

Verify domain DNS records



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
	domainId := "domainId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SendingDomainsAPI.VerifySendingDomain(context.Background(), domainId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SendingDomainsAPI.VerifySendingDomain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VerifySendingDomain`: CreateSendingDomain201Response
	fmt.Fprintf(os.Stdout, "Response from `SendingDomainsAPI.VerifySendingDomain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiVerifySendingDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreateSendingDomain201Response**](CreateSendingDomain201Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

