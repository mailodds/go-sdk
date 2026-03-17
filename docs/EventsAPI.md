# \EventsAPI

All URIs are relative to *https://api.mailodds.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TrackEvent**](EventsAPI.md#TrackEvent) | **Post** /v1/events/track | Track a commerce event



## TrackEvent

> TrackEventResponse TrackEvent(ctx).TrackEventRequest(trackEventRequest).Execute()

Track a commerce event



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
	trackEventRequest := *openapiclient.NewTrackEventRequest("EventType_example", "Email_example") // TrackEventRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EventsAPI.TrackEvent(context.Background()).TrackEventRequest(trackEventRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.TrackEvent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TrackEvent`: TrackEventResponse
	fmt.Fprintf(os.Stdout, "Response from `EventsAPI.TrackEvent`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTrackEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **trackEventRequest** | [**TrackEventRequest**](TrackEventRequest.md) |  | 

### Return type

[**TrackEventResponse**](TrackEventResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

