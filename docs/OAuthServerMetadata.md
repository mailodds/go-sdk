# OAuthServerMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Issuer** | Pointer to **string** |  | [optional] 
**AuthorizationEndpoint** | Pointer to **string** |  | [optional] 
**TokenEndpoint** | Pointer to **string** |  | [optional] 
**RevocationEndpoint** | Pointer to **string** |  | [optional] 
**IntrospectionEndpoint** | Pointer to **string** |  | [optional] 
**JwksUri** | Pointer to **string** |  | [optional] 
**ResponseTypesSupported** | Pointer to **[]string** |  | [optional] 
**GrantTypesSupported** | Pointer to **[]string** |  | [optional] 
**TokenEndpointAuthMethodsSupported** | Pointer to **[]string** |  | [optional] 
**ScopesSupported** | Pointer to **[]string** |  | [optional] 
**CodeChallengeMethodsSupported** | Pointer to **[]string** |  | [optional] 
**RevocationEndpointAuthMethodsSupported** | Pointer to **[]string** |  | [optional] 
**IntrospectionEndpointAuthMethodsSupported** | Pointer to **[]string** |  | [optional] 

## Methods

### NewOAuthServerMetadata

`func NewOAuthServerMetadata() *OAuthServerMetadata`

NewOAuthServerMetadata instantiates a new OAuthServerMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOAuthServerMetadataWithDefaults

`func NewOAuthServerMetadataWithDefaults() *OAuthServerMetadata`

NewOAuthServerMetadataWithDefaults instantiates a new OAuthServerMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIssuer

`func (o *OAuthServerMetadata) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *OAuthServerMetadata) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *OAuthServerMetadata) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *OAuthServerMetadata) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### GetAuthorizationEndpoint

`func (o *OAuthServerMetadata) GetAuthorizationEndpoint() string`

GetAuthorizationEndpoint returns the AuthorizationEndpoint field if non-nil, zero value otherwise.

### GetAuthorizationEndpointOk

`func (o *OAuthServerMetadata) GetAuthorizationEndpointOk() (*string, bool)`

GetAuthorizationEndpointOk returns a tuple with the AuthorizationEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationEndpoint

`func (o *OAuthServerMetadata) SetAuthorizationEndpoint(v string)`

SetAuthorizationEndpoint sets AuthorizationEndpoint field to given value.

### HasAuthorizationEndpoint

`func (o *OAuthServerMetadata) HasAuthorizationEndpoint() bool`

HasAuthorizationEndpoint returns a boolean if a field has been set.

### GetTokenEndpoint

`func (o *OAuthServerMetadata) GetTokenEndpoint() string`

GetTokenEndpoint returns the TokenEndpoint field if non-nil, zero value otherwise.

### GetTokenEndpointOk

`func (o *OAuthServerMetadata) GetTokenEndpointOk() (*string, bool)`

GetTokenEndpointOk returns a tuple with the TokenEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenEndpoint

`func (o *OAuthServerMetadata) SetTokenEndpoint(v string)`

SetTokenEndpoint sets TokenEndpoint field to given value.

### HasTokenEndpoint

`func (o *OAuthServerMetadata) HasTokenEndpoint() bool`

HasTokenEndpoint returns a boolean if a field has been set.

### GetRevocationEndpoint

`func (o *OAuthServerMetadata) GetRevocationEndpoint() string`

GetRevocationEndpoint returns the RevocationEndpoint field if non-nil, zero value otherwise.

### GetRevocationEndpointOk

`func (o *OAuthServerMetadata) GetRevocationEndpointOk() (*string, bool)`

GetRevocationEndpointOk returns a tuple with the RevocationEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevocationEndpoint

`func (o *OAuthServerMetadata) SetRevocationEndpoint(v string)`

SetRevocationEndpoint sets RevocationEndpoint field to given value.

### HasRevocationEndpoint

`func (o *OAuthServerMetadata) HasRevocationEndpoint() bool`

HasRevocationEndpoint returns a boolean if a field has been set.

### GetIntrospectionEndpoint

`func (o *OAuthServerMetadata) GetIntrospectionEndpoint() string`

GetIntrospectionEndpoint returns the IntrospectionEndpoint field if non-nil, zero value otherwise.

### GetIntrospectionEndpointOk

`func (o *OAuthServerMetadata) GetIntrospectionEndpointOk() (*string, bool)`

GetIntrospectionEndpointOk returns a tuple with the IntrospectionEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntrospectionEndpoint

`func (o *OAuthServerMetadata) SetIntrospectionEndpoint(v string)`

SetIntrospectionEndpoint sets IntrospectionEndpoint field to given value.

### HasIntrospectionEndpoint

`func (o *OAuthServerMetadata) HasIntrospectionEndpoint() bool`

HasIntrospectionEndpoint returns a boolean if a field has been set.

### GetJwksUri

`func (o *OAuthServerMetadata) GetJwksUri() string`

GetJwksUri returns the JwksUri field if non-nil, zero value otherwise.

### GetJwksUriOk

`func (o *OAuthServerMetadata) GetJwksUriOk() (*string, bool)`

GetJwksUriOk returns a tuple with the JwksUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwksUri

`func (o *OAuthServerMetadata) SetJwksUri(v string)`

SetJwksUri sets JwksUri field to given value.

### HasJwksUri

`func (o *OAuthServerMetadata) HasJwksUri() bool`

HasJwksUri returns a boolean if a field has been set.

### GetResponseTypesSupported

`func (o *OAuthServerMetadata) GetResponseTypesSupported() []string`

GetResponseTypesSupported returns the ResponseTypesSupported field if non-nil, zero value otherwise.

### GetResponseTypesSupportedOk

`func (o *OAuthServerMetadata) GetResponseTypesSupportedOk() (*[]string, bool)`

GetResponseTypesSupportedOk returns a tuple with the ResponseTypesSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseTypesSupported

`func (o *OAuthServerMetadata) SetResponseTypesSupported(v []string)`

SetResponseTypesSupported sets ResponseTypesSupported field to given value.

### HasResponseTypesSupported

`func (o *OAuthServerMetadata) HasResponseTypesSupported() bool`

HasResponseTypesSupported returns a boolean if a field has been set.

### GetGrantTypesSupported

`func (o *OAuthServerMetadata) GetGrantTypesSupported() []string`

GetGrantTypesSupported returns the GrantTypesSupported field if non-nil, zero value otherwise.

### GetGrantTypesSupportedOk

`func (o *OAuthServerMetadata) GetGrantTypesSupportedOk() (*[]string, bool)`

GetGrantTypesSupportedOk returns a tuple with the GrantTypesSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantTypesSupported

`func (o *OAuthServerMetadata) SetGrantTypesSupported(v []string)`

SetGrantTypesSupported sets GrantTypesSupported field to given value.

### HasGrantTypesSupported

`func (o *OAuthServerMetadata) HasGrantTypesSupported() bool`

HasGrantTypesSupported returns a boolean if a field has been set.

### GetTokenEndpointAuthMethodsSupported

`func (o *OAuthServerMetadata) GetTokenEndpointAuthMethodsSupported() []string`

GetTokenEndpointAuthMethodsSupported returns the TokenEndpointAuthMethodsSupported field if non-nil, zero value otherwise.

### GetTokenEndpointAuthMethodsSupportedOk

`func (o *OAuthServerMetadata) GetTokenEndpointAuthMethodsSupportedOk() (*[]string, bool)`

GetTokenEndpointAuthMethodsSupportedOk returns a tuple with the TokenEndpointAuthMethodsSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenEndpointAuthMethodsSupported

`func (o *OAuthServerMetadata) SetTokenEndpointAuthMethodsSupported(v []string)`

SetTokenEndpointAuthMethodsSupported sets TokenEndpointAuthMethodsSupported field to given value.

### HasTokenEndpointAuthMethodsSupported

`func (o *OAuthServerMetadata) HasTokenEndpointAuthMethodsSupported() bool`

HasTokenEndpointAuthMethodsSupported returns a boolean if a field has been set.

### GetScopesSupported

`func (o *OAuthServerMetadata) GetScopesSupported() []string`

GetScopesSupported returns the ScopesSupported field if non-nil, zero value otherwise.

### GetScopesSupportedOk

`func (o *OAuthServerMetadata) GetScopesSupportedOk() (*[]string, bool)`

GetScopesSupportedOk returns a tuple with the ScopesSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopesSupported

`func (o *OAuthServerMetadata) SetScopesSupported(v []string)`

SetScopesSupported sets ScopesSupported field to given value.

### HasScopesSupported

`func (o *OAuthServerMetadata) HasScopesSupported() bool`

HasScopesSupported returns a boolean if a field has been set.

### GetCodeChallengeMethodsSupported

`func (o *OAuthServerMetadata) GetCodeChallengeMethodsSupported() []string`

GetCodeChallengeMethodsSupported returns the CodeChallengeMethodsSupported field if non-nil, zero value otherwise.

### GetCodeChallengeMethodsSupportedOk

`func (o *OAuthServerMetadata) GetCodeChallengeMethodsSupportedOk() (*[]string, bool)`

GetCodeChallengeMethodsSupportedOk returns a tuple with the CodeChallengeMethodsSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeChallengeMethodsSupported

`func (o *OAuthServerMetadata) SetCodeChallengeMethodsSupported(v []string)`

SetCodeChallengeMethodsSupported sets CodeChallengeMethodsSupported field to given value.

### HasCodeChallengeMethodsSupported

`func (o *OAuthServerMetadata) HasCodeChallengeMethodsSupported() bool`

HasCodeChallengeMethodsSupported returns a boolean if a field has been set.

### GetRevocationEndpointAuthMethodsSupported

`func (o *OAuthServerMetadata) GetRevocationEndpointAuthMethodsSupported() []string`

GetRevocationEndpointAuthMethodsSupported returns the RevocationEndpointAuthMethodsSupported field if non-nil, zero value otherwise.

### GetRevocationEndpointAuthMethodsSupportedOk

`func (o *OAuthServerMetadata) GetRevocationEndpointAuthMethodsSupportedOk() (*[]string, bool)`

GetRevocationEndpointAuthMethodsSupportedOk returns a tuple with the RevocationEndpointAuthMethodsSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevocationEndpointAuthMethodsSupported

`func (o *OAuthServerMetadata) SetRevocationEndpointAuthMethodsSupported(v []string)`

SetRevocationEndpointAuthMethodsSupported sets RevocationEndpointAuthMethodsSupported field to given value.

### HasRevocationEndpointAuthMethodsSupported

`func (o *OAuthServerMetadata) HasRevocationEndpointAuthMethodsSupported() bool`

HasRevocationEndpointAuthMethodsSupported returns a boolean if a field has been set.

### GetIntrospectionEndpointAuthMethodsSupported

`func (o *OAuthServerMetadata) GetIntrospectionEndpointAuthMethodsSupported() []string`

GetIntrospectionEndpointAuthMethodsSupported returns the IntrospectionEndpointAuthMethodsSupported field if non-nil, zero value otherwise.

### GetIntrospectionEndpointAuthMethodsSupportedOk

`func (o *OAuthServerMetadata) GetIntrospectionEndpointAuthMethodsSupportedOk() (*[]string, bool)`

GetIntrospectionEndpointAuthMethodsSupportedOk returns a tuple with the IntrospectionEndpointAuthMethodsSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntrospectionEndpointAuthMethodsSupported

`func (o *OAuthServerMetadata) SetIntrospectionEndpointAuthMethodsSupported(v []string)`

SetIntrospectionEndpointAuthMethodsSupported sets IntrospectionEndpointAuthMethodsSupported field to given value.

### HasIntrospectionEndpointAuthMethodsSupported

`func (o *OAuthServerMetadata) HasIntrospectionEndpointAuthMethodsSupported() bool`

HasIntrospectionEndpointAuthMethodsSupported returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


