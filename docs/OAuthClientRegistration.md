# OAuthClientRegistration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | **string** | Issued client identifier | 
**ClientName** | **string** | Human-readable client name | 
**RedirectUris** | **[]string** | Registered redirect URIs | 
**GrantTypes** | **[]string** | Allowed grant types | 
**ResponseTypes** | **[]string** | Allowed response types | 
**TokenEndpointAuthMethod** | **string** | Token endpoint auth method | 
**Scope** | Pointer to **string** | Allowed scope | [optional] 
**ClientIdIssuedAt** | **int32** | Unix timestamp of client registration | 
**ClientSecret** | Pointer to **string** | Client secret (only for confidential clients, shown once) | [optional] 
**ClientSecretExpiresAt** | Pointer to **int32** | Secret expiry (0 &#x3D; never) | [optional] 

## Methods

### NewOAuthClientRegistration

`func NewOAuthClientRegistration(clientId string, clientName string, redirectUris []string, grantTypes []string, responseTypes []string, tokenEndpointAuthMethod string, clientIdIssuedAt int32, ) *OAuthClientRegistration`

NewOAuthClientRegistration instantiates a new OAuthClientRegistration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOAuthClientRegistrationWithDefaults

`func NewOAuthClientRegistrationWithDefaults() *OAuthClientRegistration`

NewOAuthClientRegistrationWithDefaults instantiates a new OAuthClientRegistration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *OAuthClientRegistration) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *OAuthClientRegistration) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *OAuthClientRegistration) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetClientName

`func (o *OAuthClientRegistration) GetClientName() string`

GetClientName returns the ClientName field if non-nil, zero value otherwise.

### GetClientNameOk

`func (o *OAuthClientRegistration) GetClientNameOk() (*string, bool)`

GetClientNameOk returns a tuple with the ClientName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientName

`func (o *OAuthClientRegistration) SetClientName(v string)`

SetClientName sets ClientName field to given value.


### GetRedirectUris

`func (o *OAuthClientRegistration) GetRedirectUris() []string`

GetRedirectUris returns the RedirectUris field if non-nil, zero value otherwise.

### GetRedirectUrisOk

`func (o *OAuthClientRegistration) GetRedirectUrisOk() (*[]string, bool)`

GetRedirectUrisOk returns a tuple with the RedirectUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUris

`func (o *OAuthClientRegistration) SetRedirectUris(v []string)`

SetRedirectUris sets RedirectUris field to given value.


### GetGrantTypes

`func (o *OAuthClientRegistration) GetGrantTypes() []string`

GetGrantTypes returns the GrantTypes field if non-nil, zero value otherwise.

### GetGrantTypesOk

`func (o *OAuthClientRegistration) GetGrantTypesOk() (*[]string, bool)`

GetGrantTypesOk returns a tuple with the GrantTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantTypes

`func (o *OAuthClientRegistration) SetGrantTypes(v []string)`

SetGrantTypes sets GrantTypes field to given value.


### GetResponseTypes

`func (o *OAuthClientRegistration) GetResponseTypes() []string`

GetResponseTypes returns the ResponseTypes field if non-nil, zero value otherwise.

### GetResponseTypesOk

`func (o *OAuthClientRegistration) GetResponseTypesOk() (*[]string, bool)`

GetResponseTypesOk returns a tuple with the ResponseTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseTypes

`func (o *OAuthClientRegistration) SetResponseTypes(v []string)`

SetResponseTypes sets ResponseTypes field to given value.


### GetTokenEndpointAuthMethod

`func (o *OAuthClientRegistration) GetTokenEndpointAuthMethod() string`

GetTokenEndpointAuthMethod returns the TokenEndpointAuthMethod field if non-nil, zero value otherwise.

### GetTokenEndpointAuthMethodOk

`func (o *OAuthClientRegistration) GetTokenEndpointAuthMethodOk() (*string, bool)`

GetTokenEndpointAuthMethodOk returns a tuple with the TokenEndpointAuthMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenEndpointAuthMethod

`func (o *OAuthClientRegistration) SetTokenEndpointAuthMethod(v string)`

SetTokenEndpointAuthMethod sets TokenEndpointAuthMethod field to given value.


### GetScope

`func (o *OAuthClientRegistration) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *OAuthClientRegistration) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *OAuthClientRegistration) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *OAuthClientRegistration) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetClientIdIssuedAt

`func (o *OAuthClientRegistration) GetClientIdIssuedAt() int32`

GetClientIdIssuedAt returns the ClientIdIssuedAt field if non-nil, zero value otherwise.

### GetClientIdIssuedAtOk

`func (o *OAuthClientRegistration) GetClientIdIssuedAtOk() (*int32, bool)`

GetClientIdIssuedAtOk returns a tuple with the ClientIdIssuedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientIdIssuedAt

`func (o *OAuthClientRegistration) SetClientIdIssuedAt(v int32)`

SetClientIdIssuedAt sets ClientIdIssuedAt field to given value.


### GetClientSecret

`func (o *OAuthClientRegistration) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *OAuthClientRegistration) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *OAuthClientRegistration) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *OAuthClientRegistration) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### GetClientSecretExpiresAt

`func (o *OAuthClientRegistration) GetClientSecretExpiresAt() int32`

GetClientSecretExpiresAt returns the ClientSecretExpiresAt field if non-nil, zero value otherwise.

### GetClientSecretExpiresAtOk

`func (o *OAuthClientRegistration) GetClientSecretExpiresAtOk() (*int32, bool)`

GetClientSecretExpiresAtOk returns a tuple with the ClientSecretExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecretExpiresAt

`func (o *OAuthClientRegistration) SetClientSecretExpiresAt(v int32)`

SetClientSecretExpiresAt sets ClientSecretExpiresAt field to given value.

### HasClientSecretExpiresAt

`func (o *OAuthClientRegistration) HasClientSecretExpiresAt() bool`

HasClientSecretExpiresAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


