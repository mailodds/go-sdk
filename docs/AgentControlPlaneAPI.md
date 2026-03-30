# \AgentControlPlaneAPI

All URIs are relative to *https://api.mailodds.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMcpCapabilities**](AgentControlPlaneAPI.md#GetMcpCapabilities) | **Get** /v1/mcp/capabilities | Get MCP capabilities



## GetMcpCapabilities

> McpCapabilities GetMcpCapabilities(ctx).Execute()

Get MCP capabilities



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
	resp, r, err := apiClient.AgentControlPlaneAPI.GetMcpCapabilities(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentControlPlaneAPI.GetMcpCapabilities``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMcpCapabilities`: McpCapabilities
	fmt.Fprintf(os.Stdout, "Response from `AgentControlPlaneAPI.GetMcpCapabilities`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMcpCapabilitiesRequest struct via the builder pattern


### Return type

[**McpCapabilities**](McpCapabilities.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

