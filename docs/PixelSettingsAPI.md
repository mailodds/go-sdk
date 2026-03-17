# \PixelSettingsAPI

All URIs are relative to *https://api.mailodds.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPixelSettings**](PixelSettingsAPI.md#GetPixelSettings) | **Get** /v1/pixel-settings | Get pixel settings
[**UpdatePixelSettings**](PixelSettingsAPI.md#UpdatePixelSettings) | **Patch** /v1/pixel-settings | Update pixel settings



## GetPixelSettings

> GetPixelSettings200Response GetPixelSettings(ctx).Execute()

Get pixel settings



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PixelSettingsAPI.GetPixelSettings(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PixelSettingsAPI.GetPixelSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPixelSettings`: GetPixelSettings200Response
	fmt.Fprintf(os.Stdout, "Response from `PixelSettingsAPI.GetPixelSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPixelSettingsRequest struct via the builder pattern


### Return type

[**GetPixelSettings200Response**](GetPixelSettings200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePixelSettings

> GetPixelSettings200Response UpdatePixelSettings(ctx).UpdatePixelSettingsRequest(updatePixelSettingsRequest).Execute()

Update pixel settings



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
	updatePixelSettingsRequest := *openapiclient.NewUpdatePixelSettingsRequest(NullableInt32(123)) // UpdatePixelSettingsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PixelSettingsAPI.UpdatePixelSettings(context.Background()).UpdatePixelSettingsRequest(updatePixelSettingsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PixelSettingsAPI.UpdatePixelSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdatePixelSettings`: GetPixelSettings200Response
	fmt.Fprintf(os.Stdout, "Response from `PixelSettingsAPI.UpdatePixelSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePixelSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updatePixelSettingsRequest** | [**UpdatePixelSettingsRequest**](UpdatePixelSettingsRequest.md) |  | 

### Return type

[**GetPixelSettings200Response**](GetPixelSettings200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

