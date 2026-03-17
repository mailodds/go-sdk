# \ContentClassificationAPI

All URIs are relative to *https://api.mailodds.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ClassifyContent**](ContentClassificationAPI.md#ClassifyContent) | **Post** /v1/content-check | Classify email content



## ClassifyContent

> ClassifyContent200Response ClassifyContent(ctx).ClassifyContentRequest(classifyContentRequest).Execute()

Classify email content



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
	classifyContentRequest := *openapiclient.NewClassifyContentRequest() // ClassifyContentRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentClassificationAPI.ClassifyContent(context.Background()).ClassifyContentRequest(classifyContentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentClassificationAPI.ClassifyContent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ClassifyContent`: ClassifyContent200Response
	fmt.Fprintf(os.Stdout, "Response from `ContentClassificationAPI.ClassifyContent`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiClassifyContentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **classifyContentRequest** | [**ClassifyContentRequest**](ClassifyContentRequest.md) |  | 

### Return type

[**ClassifyContent200Response**](ClassifyContent200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

