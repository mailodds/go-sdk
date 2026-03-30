# \ISPFBLGuidesAPI

All URIs are relative to *https://api.mailodds.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetIspFblGuide**](ISPFBLGuidesAPI.md#GetIspFblGuide) | **Get** /v1/isp-fbl/guides/{isp_id} | Get ISP FBL guide
[**ListIspFblGuides**](ISPFBLGuidesAPI.md#ListIspFblGuides) | **Get** /v1/isp-fbl/guides | List ISP FBL guides



## GetIspFblGuide

> GetIspFblGuide(ctx, ispId).Execute()

Get ISP FBL guide



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
	ispId := "ispId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ISPFBLGuidesAPI.GetIspFblGuide(context.Background(), ispId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ISPFBLGuidesAPI.GetIspFblGuide``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ispId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIspFblGuideRequest struct via the builder pattern


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


## ListIspFblGuides

> ListIspFblGuides(ctx).Execute()

List ISP FBL guides



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
	r, err := apiClient.ISPFBLGuidesAPI.ListIspFblGuides(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ISPFBLGuidesAPI.ListIspFblGuides``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListIspFblGuidesRequest struct via the builder pattern


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

