# \EmailValidationAPI

All URIs are relative to *https://api.mailodds.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ValidateEmail**](EmailValidationAPI.md#ValidateEmail) | **Post** /v1/validate | Validate single email



## ValidateEmail

> ValidationResponse ValidateEmail(ctx).ValidateRequest(validateRequest).Execute()

Validate single email



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mailodds/go-sdk"
)

func main() {
	validateRequest := *openapiclient.NewValidateRequest("Email_example") // ValidateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmailValidationAPI.ValidateEmail(context.Background()).ValidateRequest(validateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailValidationAPI.ValidateEmail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ValidateEmail`: ValidationResponse
	fmt.Fprintf(os.Stdout, "Response from `EmailValidationAPI.ValidateEmail`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateEmailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **validateRequest** | [**ValidateRequest**](ValidateRequest.md) |  | 

### Return type

[**ValidationResponse**](ValidationResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

