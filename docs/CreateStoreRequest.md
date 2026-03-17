# CreateStoreRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Platform** | **string** | E-commerce platform | 
**StoreName** | **string** | Display name for the store | 
**StoreUrl** | **string** | Store base URL | 
**AuthMethod** | **string** | Authentication method | 
**Settings** | Pointer to **map[string]interface{}** | Platform-specific settings (e.g., API keys, feed URL) | [optional] 

## Methods

### NewCreateStoreRequest

`func NewCreateStoreRequest(platform string, storeName string, storeUrl string, authMethod string, ) *CreateStoreRequest`

NewCreateStoreRequest instantiates a new CreateStoreRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateStoreRequestWithDefaults

`func NewCreateStoreRequestWithDefaults() *CreateStoreRequest`

NewCreateStoreRequestWithDefaults instantiates a new CreateStoreRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlatform

`func (o *CreateStoreRequest) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *CreateStoreRequest) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *CreateStoreRequest) SetPlatform(v string)`

SetPlatform sets Platform field to given value.


### GetStoreName

`func (o *CreateStoreRequest) GetStoreName() string`

GetStoreName returns the StoreName field if non-nil, zero value otherwise.

### GetStoreNameOk

`func (o *CreateStoreRequest) GetStoreNameOk() (*string, bool)`

GetStoreNameOk returns a tuple with the StoreName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreName

`func (o *CreateStoreRequest) SetStoreName(v string)`

SetStoreName sets StoreName field to given value.


### GetStoreUrl

`func (o *CreateStoreRequest) GetStoreUrl() string`

GetStoreUrl returns the StoreUrl field if non-nil, zero value otherwise.

### GetStoreUrlOk

`func (o *CreateStoreRequest) GetStoreUrlOk() (*string, bool)`

GetStoreUrlOk returns a tuple with the StoreUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreUrl

`func (o *CreateStoreRequest) SetStoreUrl(v string)`

SetStoreUrl sets StoreUrl field to given value.


### GetAuthMethod

`func (o *CreateStoreRequest) GetAuthMethod() string`

GetAuthMethod returns the AuthMethod field if non-nil, zero value otherwise.

### GetAuthMethodOk

`func (o *CreateStoreRequest) GetAuthMethodOk() (*string, bool)`

GetAuthMethodOk returns a tuple with the AuthMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthMethod

`func (o *CreateStoreRequest) SetAuthMethod(v string)`

SetAuthMethod sets AuthMethod field to given value.


### GetSettings

`func (o *CreateStoreRequest) GetSettings() map[string]interface{}`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *CreateStoreRequest) GetSettingsOk() (*map[string]interface{}, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *CreateStoreRequest) SetSettings(v map[string]interface{})`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *CreateStoreRequest) HasSettings() bool`

HasSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


