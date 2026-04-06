# \StorefrontDomainsAPI

All URIs are relative to *https://api.mailodds.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateStorefrontDomain**](StorefrontDomainsAPI.md#CreateStorefrontDomain) | **Post** /v1/storefront-domains | Add a custom storefront domain
[**DeleteStorefrontDomain**](StorefrontDomainsAPI.md#DeleteStorefrontDomain) | **Delete** /v1/storefront-domains/{domain_id} | Delete a storefront domain
[**GetStorefrontDomain**](StorefrontDomainsAPI.md#GetStorefrontDomain) | **Get** /v1/storefront-domains/{domain_id} | Get storefront domain details
[**ListStorefrontDomains**](StorefrontDomainsAPI.md#ListStorefrontDomains) | **Get** /v1/storefront-domains | List storefront domains
[**VerifyStorefrontDomain**](StorefrontDomainsAPI.md#VerifyStorefrontDomain) | **Post** /v1/storefront-domains/{domain_id}/verify | Verify storefront domain DNS



## CreateStorefrontDomain

> CreateStorefrontDomain(ctx).CreateStorefrontDomainRequest(createStorefrontDomainRequest).Execute()

Add a custom storefront domain



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
	createStorefrontDomainRequest := *openapiclient.NewCreateStorefrontDomainRequest("shop.merchant.com", "uuid") // CreateStorefrontDomainRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.StorefrontDomainsAPI.CreateStorefrontDomain(context.Background()).CreateStorefrontDomainRequest(createStorefrontDomainRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorefrontDomainsAPI.CreateStorefrontDomain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateStorefrontDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createStorefrontDomainRequest** | [**CreateStorefrontDomainRequest**](CreateStorefrontDomainRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteStorefrontDomain

> DeleteStorefrontDomain(ctx, domainId).Execute()

Delete a storefront domain



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
	r, err := apiClient.StorefrontDomainsAPI.DeleteStorefrontDomain(context.Background(), domainId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorefrontDomainsAPI.DeleteStorefrontDomain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteStorefrontDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStorefrontDomain

> GetStorefrontDomain(ctx, domainId).Execute()

Get storefront domain details



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
	r, err := apiClient.StorefrontDomainsAPI.GetStorefrontDomain(context.Background(), domainId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorefrontDomainsAPI.GetStorefrontDomain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStorefrontDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListStorefrontDomains

> ListStorefrontDomains(ctx).Execute()

List storefront domains



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
	r, err := apiClient.StorefrontDomainsAPI.ListStorefrontDomains(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorefrontDomainsAPI.ListStorefrontDomains``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListStorefrontDomainsRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VerifyStorefrontDomain

> VerifyStorefrontDomain(ctx, domainId).Execute()

Verify storefront domain DNS



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
	r, err := apiClient.StorefrontDomainsAPI.VerifyStorefrontDomain(context.Background(), domainId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorefrontDomainsAPI.VerifyStorefrontDomain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiVerifyStorefrontDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

