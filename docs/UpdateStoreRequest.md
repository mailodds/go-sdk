# UpdateStoreRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StoreName** | Pointer to **string** | Display name for the store | [optional] 
**SyncIntervalSeconds** | Pointer to **int32** | Auto-sync interval in seconds (min 1800) | [optional] 
**Settings** | Pointer to **map[string]interface{}** | Platform-specific settings | [optional] 
**Credentials** | Pointer to **map[string]interface{}** | Updated store credentials (connection is tested before saving) | [optional] 

## Methods

### NewUpdateStoreRequest

`func NewUpdateStoreRequest() *UpdateStoreRequest`

NewUpdateStoreRequest instantiates a new UpdateStoreRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateStoreRequestWithDefaults

`func NewUpdateStoreRequestWithDefaults() *UpdateStoreRequest`

NewUpdateStoreRequestWithDefaults instantiates a new UpdateStoreRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStoreName

`func (o *UpdateStoreRequest) GetStoreName() string`

GetStoreName returns the StoreName field if non-nil, zero value otherwise.

### GetStoreNameOk

`func (o *UpdateStoreRequest) GetStoreNameOk() (*string, bool)`

GetStoreNameOk returns a tuple with the StoreName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreName

`func (o *UpdateStoreRequest) SetStoreName(v string)`

SetStoreName sets StoreName field to given value.

### HasStoreName

`func (o *UpdateStoreRequest) HasStoreName() bool`

HasStoreName returns a boolean if a field has been set.

### GetSyncIntervalSeconds

`func (o *UpdateStoreRequest) GetSyncIntervalSeconds() int32`

GetSyncIntervalSeconds returns the SyncIntervalSeconds field if non-nil, zero value otherwise.

### GetSyncIntervalSecondsOk

`func (o *UpdateStoreRequest) GetSyncIntervalSecondsOk() (*int32, bool)`

GetSyncIntervalSecondsOk returns a tuple with the SyncIntervalSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncIntervalSeconds

`func (o *UpdateStoreRequest) SetSyncIntervalSeconds(v int32)`

SetSyncIntervalSeconds sets SyncIntervalSeconds field to given value.

### HasSyncIntervalSeconds

`func (o *UpdateStoreRequest) HasSyncIntervalSeconds() bool`

HasSyncIntervalSeconds returns a boolean if a field has been set.

### GetSettings

`func (o *UpdateStoreRequest) GetSettings() map[string]interface{}`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *UpdateStoreRequest) GetSettingsOk() (*map[string]interface{}, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *UpdateStoreRequest) SetSettings(v map[string]interface{})`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *UpdateStoreRequest) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### GetCredentials

`func (o *UpdateStoreRequest) GetCredentials() map[string]interface{}`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *UpdateStoreRequest) GetCredentialsOk() (*map[string]interface{}, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *UpdateStoreRequest) SetCredentials(v map[string]interface{})`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *UpdateStoreRequest) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


