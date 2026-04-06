# \CampaignsAPI

All URIs are relative to *https://api.mailodds.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelCampaign**](CampaignsAPI.md#CancelCampaign) | **Post** /v1/campaigns/{campaign_id}/cancel | Cancel a campaign
[**CreateCampaign**](CampaignsAPI.md#CreateCampaign) | **Post** /v1/campaigns | Create a campaign
[**CreateCampaignVariant**](CampaignsAPI.md#CreateCampaignVariant) | **Post** /v1/campaigns/{campaign_id}/variants | Create A/B variant
[**DeleteCampaign**](CampaignsAPI.md#DeleteCampaign) | **Delete** /v1/campaigns/{campaign_id} | Delete a campaign
[**GetCampaign**](CampaignsAPI.md#GetCampaign) | **Get** /v1/campaigns/{campaign_id} | Get campaign with stats
[**ListCampaigns**](CampaignsAPI.md#ListCampaigns) | **Get** /v1/campaigns | List campaigns
[**ScheduleCampaign**](CampaignsAPI.md#ScheduleCampaign) | **Post** /v1/campaigns/{campaign_id}/schedule | Schedule a campaign
[**SendCampaign**](CampaignsAPI.md#SendCampaign) | **Post** /v1/campaigns/{campaign_id}/send | Send a campaign



## CancelCampaign

> CampaignResponse CancelCampaign(ctx, campaignId).Execute()

Cancel a campaign



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
	campaignId := "campaignId_example" // string | Campaign UUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CampaignsAPI.CancelCampaign(context.Background(), campaignId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CampaignsAPI.CancelCampaign``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CancelCampaign`: CampaignResponse
	fmt.Fprintf(os.Stdout, "Response from `CampaignsAPI.CancelCampaign`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **string** | Campaign UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelCampaignRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CampaignResponse**](CampaignResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCampaign

> CampaignResponse CreateCampaign(ctx).CreateCampaignRequest(createCampaignRequest).Execute()

Create a campaign



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
	createCampaignRequest := *openapiclient.NewCreateCampaignRequest("Name_example", "ListId_example", "DomainId_example", "FromEmail_example") // CreateCampaignRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CampaignsAPI.CreateCampaign(context.Background()).CreateCampaignRequest(createCampaignRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CampaignsAPI.CreateCampaign``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCampaign`: CampaignResponse
	fmt.Fprintf(os.Stdout, "Response from `CampaignsAPI.CreateCampaign`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCampaignRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createCampaignRequest** | [**CreateCampaignRequest**](CreateCampaignRequest.md) |  | 

### Return type

[**CampaignResponse**](CampaignResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCampaignVariant

> CreateCampaignVariant201Response CreateCampaignVariant(ctx, campaignId).CreateVariantRequest(createVariantRequest).Execute()

Create A/B variant



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
	campaignId := "campaignId_example" // string | Campaign UUID
	createVariantRequest := *openapiclient.NewCreateVariantRequest("Name_example", "Subject_example", int32(123)) // CreateVariantRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CampaignsAPI.CreateCampaignVariant(context.Background(), campaignId).CreateVariantRequest(createVariantRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CampaignsAPI.CreateCampaignVariant``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCampaignVariant`: CreateCampaignVariant201Response
	fmt.Fprintf(os.Stdout, "Response from `CampaignsAPI.CreateCampaignVariant`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **string** | Campaign UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCampaignVariantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createVariantRequest** | [**CreateVariantRequest**](CreateVariantRequest.md) |  | 

### Return type

[**CreateCampaignVariant201Response**](CreateCampaignVariant201Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCampaign

> DeletePolicyRule200Response DeleteCampaign(ctx, campaignId).Execute()

Delete a campaign



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
	campaignId := "campaignId_example" // string | Campaign UUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CampaignsAPI.DeleteCampaign(context.Background(), campaignId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CampaignsAPI.DeleteCampaign``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteCampaign`: DeletePolicyRule200Response
	fmt.Fprintf(os.Stdout, "Response from `CampaignsAPI.DeleteCampaign`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **string** | Campaign UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCampaignRequest struct via the builder pattern


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


## GetCampaign

> CampaignResponse GetCampaign(ctx, campaignId).Execute()

Get campaign with stats



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
	campaignId := "campaignId_example" // string | Campaign UUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CampaignsAPI.GetCampaign(context.Background(), campaignId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CampaignsAPI.GetCampaign``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCampaign`: CampaignResponse
	fmt.Fprintf(os.Stdout, "Response from `CampaignsAPI.GetCampaign`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **string** | Campaign UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCampaignRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CampaignResponse**](CampaignResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCampaigns

> ListCampaigns200Response ListCampaigns(ctx).Page(page).PerPage(perPage).Status(status).Execute()

List campaigns



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
	page := int32(56) // int32 | Page number (optional) (default to 1)
	perPage := int32(56) // int32 | Items per page (optional) (default to 25)
	status := "status_example" // string | Filter by campaign status (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CampaignsAPI.ListCampaigns(context.Background()).Page(page).PerPage(perPage).Status(status).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CampaignsAPI.ListCampaigns``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCampaigns`: ListCampaigns200Response
	fmt.Fprintf(os.Stdout, "Response from `CampaignsAPI.ListCampaigns`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListCampaignsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page number | [default to 1]
 **perPage** | **int32** | Items per page | [default to 25]
 **status** | **string** | Filter by campaign status | 

### Return type

[**ListCampaigns200Response**](ListCampaigns200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ScheduleCampaign

> CampaignResponse ScheduleCampaign(ctx, campaignId).ScheduleCampaignRequest(scheduleCampaignRequest).Execute()

Schedule a campaign



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
	campaignId := "campaignId_example" // string | Campaign UUID
	scheduleCampaignRequest := *openapiclient.NewScheduleCampaignRequest(time.Now()) // ScheduleCampaignRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CampaignsAPI.ScheduleCampaign(context.Background(), campaignId).ScheduleCampaignRequest(scheduleCampaignRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CampaignsAPI.ScheduleCampaign``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ScheduleCampaign`: CampaignResponse
	fmt.Fprintf(os.Stdout, "Response from `CampaignsAPI.ScheduleCampaign`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **string** | Campaign UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiScheduleCampaignRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **scheduleCampaignRequest** | [**ScheduleCampaignRequest**](ScheduleCampaignRequest.md) |  | 

### Return type

[**CampaignResponse**](CampaignResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendCampaign

> CampaignResponse SendCampaign(ctx, campaignId).Execute()

Send a campaign



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
	campaignId := "campaignId_example" // string | Campaign UUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CampaignsAPI.SendCampaign(context.Background(), campaignId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CampaignsAPI.SendCampaign``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SendCampaign`: CampaignResponse
	fmt.Fprintf(os.Stdout, "Response from `CampaignsAPI.SendCampaign`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **string** | Campaign UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSendCampaignRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CampaignResponse**](CampaignResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

