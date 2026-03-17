# StoreConnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Store connection UUID | [optional] 
**AccountId** | Pointer to **int32** |  | [optional] 
**Platform** | Pointer to **string** | E-commerce platform | [optional] 
**StoreName** | Pointer to **string** |  | [optional] 
**StoreUrl** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**AuthMethod** | Pointer to **string** |  | [optional] 
**ProductCount** | Pointer to **int32** | Number of active products | [optional] 
**LastSyncedAt** | Pointer to **NullableTime** |  | [optional] 
**LastError** | Pointer to **NullableString** | Last sync error message | [optional] 
**SyncIntervalSeconds** | Pointer to **int32** | Auto-sync interval in seconds | [optional] [default to 3600]
**Settings** | Pointer to **map[string]interface{}** | Platform-specific settings | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewStoreConnection

`func NewStoreConnection() *StoreConnection`

NewStoreConnection instantiates a new StoreConnection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoreConnectionWithDefaults

`func NewStoreConnectionWithDefaults() *StoreConnection`

NewStoreConnectionWithDefaults instantiates a new StoreConnection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *StoreConnection) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StoreConnection) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StoreConnection) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *StoreConnection) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAccountId

`func (o *StoreConnection) GetAccountId() int32`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *StoreConnection) GetAccountIdOk() (*int32, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *StoreConnection) SetAccountId(v int32)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *StoreConnection) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetPlatform

`func (o *StoreConnection) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *StoreConnection) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *StoreConnection) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *StoreConnection) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetStoreName

`func (o *StoreConnection) GetStoreName() string`

GetStoreName returns the StoreName field if non-nil, zero value otherwise.

### GetStoreNameOk

`func (o *StoreConnection) GetStoreNameOk() (*string, bool)`

GetStoreNameOk returns a tuple with the StoreName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreName

`func (o *StoreConnection) SetStoreName(v string)`

SetStoreName sets StoreName field to given value.

### HasStoreName

`func (o *StoreConnection) HasStoreName() bool`

HasStoreName returns a boolean if a field has been set.

### GetStoreUrl

`func (o *StoreConnection) GetStoreUrl() string`

GetStoreUrl returns the StoreUrl field if non-nil, zero value otherwise.

### GetStoreUrlOk

`func (o *StoreConnection) GetStoreUrlOk() (*string, bool)`

GetStoreUrlOk returns a tuple with the StoreUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreUrl

`func (o *StoreConnection) SetStoreUrl(v string)`

SetStoreUrl sets StoreUrl field to given value.

### HasStoreUrl

`func (o *StoreConnection) HasStoreUrl() bool`

HasStoreUrl returns a boolean if a field has been set.

### GetStatus

`func (o *StoreConnection) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *StoreConnection) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *StoreConnection) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *StoreConnection) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetAuthMethod

`func (o *StoreConnection) GetAuthMethod() string`

GetAuthMethod returns the AuthMethod field if non-nil, zero value otherwise.

### GetAuthMethodOk

`func (o *StoreConnection) GetAuthMethodOk() (*string, bool)`

GetAuthMethodOk returns a tuple with the AuthMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthMethod

`func (o *StoreConnection) SetAuthMethod(v string)`

SetAuthMethod sets AuthMethod field to given value.

### HasAuthMethod

`func (o *StoreConnection) HasAuthMethod() bool`

HasAuthMethod returns a boolean if a field has been set.

### GetProductCount

`func (o *StoreConnection) GetProductCount() int32`

GetProductCount returns the ProductCount field if non-nil, zero value otherwise.

### GetProductCountOk

`func (o *StoreConnection) GetProductCountOk() (*int32, bool)`

GetProductCountOk returns a tuple with the ProductCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductCount

`func (o *StoreConnection) SetProductCount(v int32)`

SetProductCount sets ProductCount field to given value.

### HasProductCount

`func (o *StoreConnection) HasProductCount() bool`

HasProductCount returns a boolean if a field has been set.

### GetLastSyncedAt

`func (o *StoreConnection) GetLastSyncedAt() time.Time`

GetLastSyncedAt returns the LastSyncedAt field if non-nil, zero value otherwise.

### GetLastSyncedAtOk

`func (o *StoreConnection) GetLastSyncedAtOk() (*time.Time, bool)`

GetLastSyncedAtOk returns a tuple with the LastSyncedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSyncedAt

`func (o *StoreConnection) SetLastSyncedAt(v time.Time)`

SetLastSyncedAt sets LastSyncedAt field to given value.

### HasLastSyncedAt

`func (o *StoreConnection) HasLastSyncedAt() bool`

HasLastSyncedAt returns a boolean if a field has been set.

### SetLastSyncedAtNil

`func (o *StoreConnection) SetLastSyncedAtNil(b bool)`

 SetLastSyncedAtNil sets the value for LastSyncedAt to be an explicit nil

### UnsetLastSyncedAt
`func (o *StoreConnection) UnsetLastSyncedAt()`

UnsetLastSyncedAt ensures that no value is present for LastSyncedAt, not even an explicit nil
### GetLastError

`func (o *StoreConnection) GetLastError() string`

GetLastError returns the LastError field if non-nil, zero value otherwise.

### GetLastErrorOk

`func (o *StoreConnection) GetLastErrorOk() (*string, bool)`

GetLastErrorOk returns a tuple with the LastError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastError

`func (o *StoreConnection) SetLastError(v string)`

SetLastError sets LastError field to given value.

### HasLastError

`func (o *StoreConnection) HasLastError() bool`

HasLastError returns a boolean if a field has been set.

### SetLastErrorNil

`func (o *StoreConnection) SetLastErrorNil(b bool)`

 SetLastErrorNil sets the value for LastError to be an explicit nil

### UnsetLastError
`func (o *StoreConnection) UnsetLastError()`

UnsetLastError ensures that no value is present for LastError, not even an explicit nil
### GetSyncIntervalSeconds

`func (o *StoreConnection) GetSyncIntervalSeconds() int32`

GetSyncIntervalSeconds returns the SyncIntervalSeconds field if non-nil, zero value otherwise.

### GetSyncIntervalSecondsOk

`func (o *StoreConnection) GetSyncIntervalSecondsOk() (*int32, bool)`

GetSyncIntervalSecondsOk returns a tuple with the SyncIntervalSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncIntervalSeconds

`func (o *StoreConnection) SetSyncIntervalSeconds(v int32)`

SetSyncIntervalSeconds sets SyncIntervalSeconds field to given value.

### HasSyncIntervalSeconds

`func (o *StoreConnection) HasSyncIntervalSeconds() bool`

HasSyncIntervalSeconds returns a boolean if a field has been set.

### GetSettings

`func (o *StoreConnection) GetSettings() map[string]interface{}`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *StoreConnection) GetSettingsOk() (*map[string]interface{}, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *StoreConnection) SetSettings(v map[string]interface{})`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *StoreConnection) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### GetCreatedAt

`func (o *StoreConnection) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *StoreConnection) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *StoreConnection) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *StoreConnection) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *StoreConnection) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *StoreConnection) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *StoreConnection) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *StoreConnection) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


