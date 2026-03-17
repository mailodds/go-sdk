# \DMARCMonitoringAPI

All URIs are relative to *https://api.mailodds.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddDmarcDomain**](DMARCMonitoringAPI.md#AddDmarcDomain) | **Post** /v1/dmarc-domains | Add DMARC domain
[**DeleteDmarcDomain**](DMARCMonitoringAPI.md#DeleteDmarcDomain) | **Delete** /v1/dmarc-domains/{domain_id} | Delete a DMARC domain
[**GetDmarcDomain**](DMARCMonitoringAPI.md#GetDmarcDomain) | **Get** /v1/dmarc-domains/{domain_id} | Get DMARC domain
[**GetDmarcRecommendation**](DMARCMonitoringAPI.md#GetDmarcRecommendation) | **Get** /v1/dmarc-domains/{domain_id}/recommendation | Get DMARC policy recommendation
[**GetDmarcSources**](DMARCMonitoringAPI.md#GetDmarcSources) | **Get** /v1/dmarc-domains/{domain_id}/sources | Get DMARC sending sources
[**GetDmarcTrend**](DMARCMonitoringAPI.md#GetDmarcTrend) | **Get** /v1/dmarc-domains/{domain_id}/trend | Get DMARC trend
[**ListDmarcDomains**](DMARCMonitoringAPI.md#ListDmarcDomains) | **Get** /v1/dmarc-domains | List DMARC domains
[**VerifyDmarcDomain**](DMARCMonitoringAPI.md#VerifyDmarcDomain) | **Post** /v1/dmarc-domains/{domain_id}/verify | Verify DMARC DNS records



## AddDmarcDomain

> AddDmarcDomain201Response AddDmarcDomain(ctx).AddDmarcDomainRequest(addDmarcDomainRequest).Execute()

Add DMARC domain



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
	addDmarcDomainRequest := *openapiclient.NewAddDmarcDomainRequest("example.com") // AddDmarcDomainRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DMARCMonitoringAPI.AddDmarcDomain(context.Background()).AddDmarcDomainRequest(addDmarcDomainRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DMARCMonitoringAPI.AddDmarcDomain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddDmarcDomain`: AddDmarcDomain201Response
	fmt.Fprintf(os.Stdout, "Response from `DMARCMonitoringAPI.AddDmarcDomain`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddDmarcDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addDmarcDomainRequest** | [**AddDmarcDomainRequest**](AddDmarcDomainRequest.md) |  | 

### Return type

[**AddDmarcDomain201Response**](AddDmarcDomain201Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDmarcDomain

> DeletePolicyRule200Response DeleteDmarcDomain(ctx, domainId).Execute()

Delete a DMARC domain



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
	domainId := "domainId_example" // string | DMARC domain UUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DMARCMonitoringAPI.DeleteDmarcDomain(context.Background(), domainId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DMARCMonitoringAPI.DeleteDmarcDomain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteDmarcDomain`: DeletePolicyRule200Response
	fmt.Fprintf(os.Stdout, "Response from `DMARCMonitoringAPI.DeleteDmarcDomain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainId** | **string** | DMARC domain UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDmarcDomainRequest struct via the builder pattern


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


## GetDmarcDomain

> GetDmarcDomain200Response GetDmarcDomain(ctx, domainId).Days(days).Execute()

Get DMARC domain



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
	domainId := "domainId_example" // string | DMARC domain UUID
	days := int32(56) // int32 | Number of days for summary stats (optional) (default to 30)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DMARCMonitoringAPI.GetDmarcDomain(context.Background(), domainId).Days(days).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DMARCMonitoringAPI.GetDmarcDomain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDmarcDomain`: GetDmarcDomain200Response
	fmt.Fprintf(os.Stdout, "Response from `DMARCMonitoringAPI.GetDmarcDomain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainId** | **string** | DMARC domain UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDmarcDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **days** | **int32** | Number of days for summary stats | [default to 30]

### Return type

[**GetDmarcDomain200Response**](GetDmarcDomain200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDmarcRecommendation

> GetDmarcRecommendation200Response GetDmarcRecommendation(ctx, domainId).Execute()

Get DMARC policy recommendation



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
	domainId := "domainId_example" // string | DMARC domain UUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DMARCMonitoringAPI.GetDmarcRecommendation(context.Background(), domainId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DMARCMonitoringAPI.GetDmarcRecommendation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDmarcRecommendation`: GetDmarcRecommendation200Response
	fmt.Fprintf(os.Stdout, "Response from `DMARCMonitoringAPI.GetDmarcRecommendation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainId** | **string** | DMARC domain UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDmarcRecommendationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetDmarcRecommendation200Response**](GetDmarcRecommendation200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDmarcSources

> GetDmarcSources200Response GetDmarcSources(ctx, domainId).Days(days).Page(page).PerPage(perPage).Execute()

Get DMARC sending sources



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
	domainId := "domainId_example" // string | DMARC domain UUID
	days := int32(56) // int32 | Number of days to look back (optional) (default to 30)
	page := int32(56) // int32 |  (optional) (default to 1)
	perPage := int32(56) // int32 |  (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DMARCMonitoringAPI.GetDmarcSources(context.Background(), domainId).Days(days).Page(page).PerPage(perPage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DMARCMonitoringAPI.GetDmarcSources``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDmarcSources`: GetDmarcSources200Response
	fmt.Fprintf(os.Stdout, "Response from `DMARCMonitoringAPI.GetDmarcSources`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainId** | **string** | DMARC domain UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDmarcSourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **days** | **int32** | Number of days to look back | [default to 30]
 **page** | **int32** |  | [default to 1]
 **perPage** | **int32** |  | [default to 20]

### Return type

[**GetDmarcSources200Response**](GetDmarcSources200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDmarcTrend

> GetDmarcTrend200Response GetDmarcTrend(ctx, domainId).Days(days).Execute()

Get DMARC trend



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
	domainId := "domainId_example" // string | DMARC domain UUID
	days := int32(56) // int32 | Number of days of trend data (optional) (default to 30)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DMARCMonitoringAPI.GetDmarcTrend(context.Background(), domainId).Days(days).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DMARCMonitoringAPI.GetDmarcTrend``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDmarcTrend`: GetDmarcTrend200Response
	fmt.Fprintf(os.Stdout, "Response from `DMARCMonitoringAPI.GetDmarcTrend`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainId** | **string** | DMARC domain UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDmarcTrendRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **days** | **int32** | Number of days of trend data | [default to 30]

### Return type

[**GetDmarcTrend200Response**](GetDmarcTrend200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDmarcDomains

> ListDmarcDomains200Response ListDmarcDomains(ctx).Execute()

List DMARC domains



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
	resp, r, err := apiClient.DMARCMonitoringAPI.ListDmarcDomains(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DMARCMonitoringAPI.ListDmarcDomains``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDmarcDomains`: ListDmarcDomains200Response
	fmt.Fprintf(os.Stdout, "Response from `DMARCMonitoringAPI.ListDmarcDomains`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListDmarcDomainsRequest struct via the builder pattern


### Return type

[**ListDmarcDomains200Response**](ListDmarcDomains200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VerifyDmarcDomain

> AddDmarcDomain201Response VerifyDmarcDomain(ctx, domainId).Execute()

Verify DMARC DNS records



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
	domainId := "domainId_example" // string | DMARC domain UUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DMARCMonitoringAPI.VerifyDmarcDomain(context.Background(), domainId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DMARCMonitoringAPI.VerifyDmarcDomain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VerifyDmarcDomain`: AddDmarcDomain201Response
	fmt.Fprintf(os.Stdout, "Response from `DMARCMonitoringAPI.VerifyDmarcDomain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainId** | **string** | DMARC domain UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiVerifyDmarcDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AddDmarcDomain201Response**](AddDmarcDomain201Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

