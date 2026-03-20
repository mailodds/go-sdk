# \OAuth20API

All URIs are relative to *https://api.mailodds.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateToken**](OAuth20API.md#CreateToken) | **Post** /oauth/token | Create token
[**GetJwks**](OAuth20API.md#GetJwks) | **Get** /.well-known/jwks.json | Get JSON Web Key Set
[**IntrospectToken**](OAuth20API.md#IntrospectToken) | **Post** /oauth/introspect | Introspect token
[**OauthServerMetadata**](OAuth20API.md#OauthServerMetadata) | **Get** /.well-known/oauth-authorization-server | OAuth server metadata
[**RevokeToken**](OAuth20API.md#RevokeToken) | **Post** /oauth/revoke | Revoke token



## CreateToken

> CreateToken200Response CreateToken(ctx).GrantType(grantType).Code(code).RedirectUri(redirectUri).ClientId(clientId).ClientSecret(clientSecret).RefreshToken(refreshToken).Scope(scope).CodeVerifier(codeVerifier).Execute()

Create token



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
	grantType := "grantType_example" // string | 
	code := "code_example" // string | Authorization code (for authorization_code grant) (optional)
	redirectUri := "redirectUri_example" // string | Must match the original redirect_uri (optional)
	clientId := "clientId_example" // string |  (optional)
	clientSecret := "clientSecret_example" // string |  (optional)
	refreshToken := "refreshToken_example" // string | Refresh token (for refresh_token grant) (optional)
	scope := "scope_example" // string | Space-separated scopes (optional)
	codeVerifier := "codeVerifier_example" // string | PKCE code verifier (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OAuth20API.CreateToken(context.Background()).GrantType(grantType).Code(code).RedirectUri(redirectUri).ClientId(clientId).ClientSecret(clientSecret).RefreshToken(refreshToken).Scope(scope).CodeVerifier(codeVerifier).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OAuth20API.CreateToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateToken`: CreateToken200Response
	fmt.Fprintf(os.Stdout, "Response from `OAuth20API.CreateToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **grantType** | **string** |  | 
 **code** | **string** | Authorization code (for authorization_code grant) | 
 **redirectUri** | **string** | Must match the original redirect_uri | 
 **clientId** | **string** |  | 
 **clientSecret** | **string** |  | 
 **refreshToken** | **string** | Refresh token (for refresh_token grant) | 
 **scope** | **string** | Space-separated scopes | 
 **codeVerifier** | **string** | PKCE code verifier | 

### Return type

[**CreateToken200Response**](CreateToken200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetJwks

> JwksResponse GetJwks(ctx).Execute()

Get JSON Web Key Set



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
	resp, r, err := apiClient.OAuth20API.GetJwks(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OAuth20API.GetJwks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetJwks`: JwksResponse
	fmt.Fprintf(os.Stdout, "Response from `OAuth20API.GetJwks`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetJwksRequest struct via the builder pattern


### Return type

[**JwksResponse**](JwksResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IntrospectToken

> IntrospectToken200Response IntrospectToken(ctx).Token(token).TokenTypeHint(tokenTypeHint).ClientId(clientId).ClientSecret(clientSecret).Execute()

Introspect token



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
	token := "token_example" // string | Token to introspect
	tokenTypeHint := "tokenTypeHint_example" // string |  (optional)
	clientId := "clientId_example" // string |  (optional)
	clientSecret := "clientSecret_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OAuth20API.IntrospectToken(context.Background()).Token(token).TokenTypeHint(tokenTypeHint).ClientId(clientId).ClientSecret(clientSecret).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OAuth20API.IntrospectToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IntrospectToken`: IntrospectToken200Response
	fmt.Fprintf(os.Stdout, "Response from `OAuth20API.IntrospectToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIntrospectTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string** | Token to introspect | 
 **tokenTypeHint** | **string** |  | 
 **clientId** | **string** |  | 
 **clientSecret** | **string** |  | 

### Return type

[**IntrospectToken200Response**](IntrospectToken200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OauthServerMetadata

> OAuthServerMetadata OauthServerMetadata(ctx).Execute()

OAuth server metadata



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
	resp, r, err := apiClient.OAuth20API.OauthServerMetadata(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OAuth20API.OauthServerMetadata``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OauthServerMetadata`: OAuthServerMetadata
	fmt.Fprintf(os.Stdout, "Response from `OAuth20API.OauthServerMetadata`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiOauthServerMetadataRequest struct via the builder pattern


### Return type

[**OAuthServerMetadata**](OAuthServerMetadata.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RevokeToken

> RevokeToken(ctx).Token(token).TokenTypeHint(tokenTypeHint).ClientId(clientId).ClientSecret(clientSecret).Execute()

Revoke token



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
	token := "token_example" // string | Token to revoke
	tokenTypeHint := "tokenTypeHint_example" // string |  (optional)
	clientId := "clientId_example" // string |  (optional)
	clientSecret := "clientSecret_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OAuth20API.RevokeToken(context.Background()).Token(token).TokenTypeHint(tokenTypeHint).ClientId(clientId).ClientSecret(clientSecret).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OAuth20API.RevokeToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRevokeTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string** | Token to revoke | 
 **tokenTypeHint** | **string** |  | 
 **clientId** | **string** |  | 
 **clientSecret** | **string** |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

