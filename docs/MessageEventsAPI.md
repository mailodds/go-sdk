# \MessageEventsAPI

All URIs are relative to *https://api.mailodds.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMessageEvents**](MessageEventsAPI.md#GetMessageEvents) | **Get** /v1/message-events | Get message events



## GetMessageEvents

> GetMessageEvents200Response GetMessageEvents(ctx).MessageId(messageId).Execute()

Get message events



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
	messageId := "messageId_example" // string | UUID of the sent message

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MessageEventsAPI.GetMessageEvents(context.Background()).MessageId(messageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MessageEventsAPI.GetMessageEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMessageEvents`: GetMessageEvents200Response
	fmt.Fprintf(os.Stdout, "Response from `MessageEventsAPI.GetMessageEvents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMessageEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **messageId** | **string** | UUID of the sent message | 

### Return type

[**GetMessageEvents200Response**](GetMessageEvents200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

