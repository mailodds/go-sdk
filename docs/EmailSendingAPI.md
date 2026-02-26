# \EmailSendingAPI

All URIs are relative to *https://api.mailodds.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeliverBatch**](EmailSendingAPI.md#DeliverBatch) | **Post** /v1/deliver/batch | Send to multiple recipients (max 100)
[**DeliverEmail**](EmailSendingAPI.md#DeliverEmail) | **Post** /v1/deliver | Send a single email



## DeliverBatch

> BatchDeliverResponse DeliverBatch(ctx).BatchDeliverRequest(batchDeliverRequest).Execute()

Send to multiple recipients (max 100)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mailodds/go-sdk/mailodds"
)

func main() {
	batchDeliverRequest := *openapiclient.NewBatchDeliverRequest([]string{"To_example"}, "From_example", "Subject_example", "DomainId_example") // BatchDeliverRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmailSendingAPI.DeliverBatch(context.Background()).BatchDeliverRequest(batchDeliverRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailSendingAPI.DeliverBatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeliverBatch`: BatchDeliverResponse
	fmt.Fprintf(os.Stdout, "Response from `EmailSendingAPI.DeliverBatch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeliverBatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batchDeliverRequest** | [**BatchDeliverRequest**](BatchDeliverRequest.md) |  | 

### Return type

[**BatchDeliverResponse**](BatchDeliverResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeliverEmail

> DeliverResponse DeliverEmail(ctx).DeliverRequest(deliverRequest).Execute()

Send a single email



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mailodds/go-sdk/mailodds"
)

func main() {
	deliverRequest := *openapiclient.NewDeliverRequest([]openapiclient.DeliverRequestToInner{*openapiclient.NewDeliverRequestToInner("Email_example")}, "From_example", "Subject_example", "DomainId_example") // DeliverRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmailSendingAPI.DeliverEmail(context.Background()).DeliverRequest(deliverRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailSendingAPI.DeliverEmail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeliverEmail`: DeliverResponse
	fmt.Fprintf(os.Stdout, "Response from `EmailSendingAPI.DeliverEmail`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeliverEmailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deliverRequest** | [**DeliverRequest**](DeliverRequest.md) |  | 

### Return type

[**DeliverResponse**](DeliverResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

