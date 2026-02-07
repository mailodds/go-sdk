# \ValidationPoliciesAPI

All URIs are relative to *https://api.mailodds.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddPolicyRule**](ValidationPoliciesAPI.md#AddPolicyRule) | **Post** /v1/policies/{policy_id}/rules | Add rule to policy
[**CreatePolicy**](ValidationPoliciesAPI.md#CreatePolicy) | **Post** /v1/policies | Create policy
[**CreatePolicyFromPreset**](ValidationPoliciesAPI.md#CreatePolicyFromPreset) | **Post** /v1/policies/from-preset | Create policy from preset
[**DeletePolicy**](ValidationPoliciesAPI.md#DeletePolicy) | **Delete** /v1/policies/{policy_id} | Delete policy
[**DeletePolicyRule**](ValidationPoliciesAPI.md#DeletePolicyRule) | **Delete** /v1/policies/{policy_id}/rules/{rule_id} | Delete rule
[**GetPolicy**](ValidationPoliciesAPI.md#GetPolicy) | **Get** /v1/policies/{policy_id} | Get policy
[**GetPolicyPresets**](ValidationPoliciesAPI.md#GetPolicyPresets) | **Get** /v1/policies/presets | Get policy presets
[**ListPolicies**](ValidationPoliciesAPI.md#ListPolicies) | **Get** /v1/policies | List policies
[**TestPolicy**](ValidationPoliciesAPI.md#TestPolicy) | **Post** /v1/policies/test | Test policy evaluation
[**UpdatePolicy**](ValidationPoliciesAPI.md#UpdatePolicy) | **Put** /v1/policies/{policy_id} | Update policy



## AddPolicyRule

> AddPolicyRule201Response AddPolicyRule(ctx, policyId).PolicyRule(policyRule).Execute()

Add rule to policy



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
	policyId := int32(56) // int32 | 
	policyRule := *openapiclient.NewPolicyRule("Type_example", map[string]interface{}(123), *openapiclient.NewPolicyRuleAction()) // PolicyRule | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ValidationPoliciesAPI.AddPolicyRule(context.Background(), policyId).PolicyRule(policyRule).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ValidationPoliciesAPI.AddPolicyRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddPolicyRule`: AddPolicyRule201Response
	fmt.Fprintf(os.Stdout, "Response from `ValidationPoliciesAPI.AddPolicyRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddPolicyRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **policyRule** | [**PolicyRule**](PolicyRule.md) |  | 

### Return type

[**AddPolicyRule201Response**](AddPolicyRule201Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePolicy

> PolicyResponse CreatePolicy(ctx).CreatePolicyRequest(createPolicyRequest).Execute()

Create policy



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
	createPolicyRequest := *openapiclient.NewCreatePolicyRequest("Name_example") // CreatePolicyRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ValidationPoliciesAPI.CreatePolicy(context.Background()).CreatePolicyRequest(createPolicyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ValidationPoliciesAPI.CreatePolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePolicy`: PolicyResponse
	fmt.Fprintf(os.Stdout, "Response from `ValidationPoliciesAPI.CreatePolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createPolicyRequest** | [**CreatePolicyRequest**](CreatePolicyRequest.md) |  | 

### Return type

[**PolicyResponse**](PolicyResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePolicyFromPreset

> PolicyResponse CreatePolicyFromPreset(ctx).CreatePolicyFromPresetRequest(createPolicyFromPresetRequest).Execute()

Create policy from preset



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
	createPolicyFromPresetRequest := *openapiclient.NewCreatePolicyFromPresetRequest("PresetId_example", "Name_example") // CreatePolicyFromPresetRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ValidationPoliciesAPI.CreatePolicyFromPreset(context.Background()).CreatePolicyFromPresetRequest(createPolicyFromPresetRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ValidationPoliciesAPI.CreatePolicyFromPreset``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePolicyFromPreset`: PolicyResponse
	fmt.Fprintf(os.Stdout, "Response from `ValidationPoliciesAPI.CreatePolicyFromPreset`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePolicyFromPresetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createPolicyFromPresetRequest** | [**CreatePolicyFromPresetRequest**](CreatePolicyFromPresetRequest.md) |  | 

### Return type

[**PolicyResponse**](PolicyResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePolicy

> DeletePolicy200Response DeletePolicy(ctx, policyId).Execute()

Delete policy



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
	policyId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ValidationPoliciesAPI.DeletePolicy(context.Background(), policyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ValidationPoliciesAPI.DeletePolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeletePolicy`: DeletePolicy200Response
	fmt.Fprintf(os.Stdout, "Response from `ValidationPoliciesAPI.DeletePolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeletePolicy200Response**](DeletePolicy200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePolicyRule

> DeletePolicyRule200Response DeletePolicyRule(ctx, policyId, ruleId).Execute()

Delete rule



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
	policyId := int32(56) // int32 | 
	ruleId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ValidationPoliciesAPI.DeletePolicyRule(context.Background(), policyId, ruleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ValidationPoliciesAPI.DeletePolicyRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeletePolicyRule`: DeletePolicyRule200Response
	fmt.Fprintf(os.Stdout, "Response from `ValidationPoliciesAPI.DeletePolicyRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **int32** |  | 
**ruleId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePolicyRuleRequest struct via the builder pattern


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


## GetPolicy

> PolicyResponse GetPolicy(ctx, policyId).Execute()

Get policy



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
	policyId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ValidationPoliciesAPI.GetPolicy(context.Background(), policyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ValidationPoliciesAPI.GetPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPolicy`: PolicyResponse
	fmt.Fprintf(os.Stdout, "Response from `ValidationPoliciesAPI.GetPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PolicyResponse**](PolicyResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPolicyPresets

> PolicyPresetsResponse GetPolicyPresets(ctx).Execute()

Get policy presets



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ValidationPoliciesAPI.GetPolicyPresets(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ValidationPoliciesAPI.GetPolicyPresets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPolicyPresets`: PolicyPresetsResponse
	fmt.Fprintf(os.Stdout, "Response from `ValidationPoliciesAPI.GetPolicyPresets`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPolicyPresetsRequest struct via the builder pattern


### Return type

[**PolicyPresetsResponse**](PolicyPresetsResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPolicies

> PolicyListResponse ListPolicies(ctx).IncludeRules(includeRules).Execute()

List policies



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
	includeRules := true // bool | Include full rules in response (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ValidationPoliciesAPI.ListPolicies(context.Background()).IncludeRules(includeRules).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ValidationPoliciesAPI.ListPolicies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListPolicies`: PolicyListResponse
	fmt.Fprintf(os.Stdout, "Response from `ValidationPoliciesAPI.ListPolicies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **includeRules** | **bool** | Include full rules in response | [default to false]

### Return type

[**PolicyListResponse**](PolicyListResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestPolicy

> PolicyTestResponse TestPolicy(ctx).TestPolicyRequest(testPolicyRequest).Execute()

Test policy evaluation



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
	testPolicyRequest := *openapiclient.NewTestPolicyRequest(int32(123), *openapiclient.NewTestPolicyRequestTestResult()) // TestPolicyRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ValidationPoliciesAPI.TestPolicy(context.Background()).TestPolicyRequest(testPolicyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ValidationPoliciesAPI.TestPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TestPolicy`: PolicyTestResponse
	fmt.Fprintf(os.Stdout, "Response from `ValidationPoliciesAPI.TestPolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTestPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **testPolicyRequest** | [**TestPolicyRequest**](TestPolicyRequest.md) |  | 

### Return type

[**PolicyTestResponse**](PolicyTestResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePolicy

> PolicyResponse UpdatePolicy(ctx, policyId).UpdatePolicyRequest(updatePolicyRequest).Execute()

Update policy



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
	policyId := int32(56) // int32 | 
	updatePolicyRequest := *openapiclient.NewUpdatePolicyRequest() // UpdatePolicyRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ValidationPoliciesAPI.UpdatePolicy(context.Background(), policyId).UpdatePolicyRequest(updatePolicyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ValidationPoliciesAPI.UpdatePolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdatePolicy`: PolicyResponse
	fmt.Fprintf(os.Stdout, "Response from `ValidationPoliciesAPI.UpdatePolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updatePolicyRequest** | [**UpdatePolicyRequest**](UpdatePolicyRequest.md) |  | 

### Return type

[**PolicyResponse**](PolicyResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

