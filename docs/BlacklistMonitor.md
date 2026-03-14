# BlacklistMonitor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Monitor UUID | [optional] 
**Target** | Pointer to **string** | IP address or domain being monitored | [optional] 
**TargetType** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** | Current status (clean, listed) | [optional] 
**ListedCount** | Pointer to **int32** | Number of blacklists currently listing this target | [optional] 
**LastCheckedAt** | Pointer to **NullableTime** |  | [optional] 
**LatestCheck** | Pointer to [**NullableBlacklistMonitorLatestCheck**](BlacklistMonitorLatestCheck.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewBlacklistMonitor

`func NewBlacklistMonitor() *BlacklistMonitor`

NewBlacklistMonitor instantiates a new BlacklistMonitor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlacklistMonitorWithDefaults

`func NewBlacklistMonitorWithDefaults() *BlacklistMonitor`

NewBlacklistMonitorWithDefaults instantiates a new BlacklistMonitor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BlacklistMonitor) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BlacklistMonitor) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BlacklistMonitor) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BlacklistMonitor) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTarget

`func (o *BlacklistMonitor) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *BlacklistMonitor) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *BlacklistMonitor) SetTarget(v string)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *BlacklistMonitor) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetTargetType

`func (o *BlacklistMonitor) GetTargetType() string`

GetTargetType returns the TargetType field if non-nil, zero value otherwise.

### GetTargetTypeOk

`func (o *BlacklistMonitor) GetTargetTypeOk() (*string, bool)`

GetTargetTypeOk returns a tuple with the TargetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetType

`func (o *BlacklistMonitor) SetTargetType(v string)`

SetTargetType sets TargetType field to given value.

### HasTargetType

`func (o *BlacklistMonitor) HasTargetType() bool`

HasTargetType returns a boolean if a field has been set.

### GetStatus

`func (o *BlacklistMonitor) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BlacklistMonitor) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BlacklistMonitor) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BlacklistMonitor) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetListedCount

`func (o *BlacklistMonitor) GetListedCount() int32`

GetListedCount returns the ListedCount field if non-nil, zero value otherwise.

### GetListedCountOk

`func (o *BlacklistMonitor) GetListedCountOk() (*int32, bool)`

GetListedCountOk returns a tuple with the ListedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListedCount

`func (o *BlacklistMonitor) SetListedCount(v int32)`

SetListedCount sets ListedCount field to given value.

### HasListedCount

`func (o *BlacklistMonitor) HasListedCount() bool`

HasListedCount returns a boolean if a field has been set.

### GetLastCheckedAt

`func (o *BlacklistMonitor) GetLastCheckedAt() time.Time`

GetLastCheckedAt returns the LastCheckedAt field if non-nil, zero value otherwise.

### GetLastCheckedAtOk

`func (o *BlacklistMonitor) GetLastCheckedAtOk() (*time.Time, bool)`

GetLastCheckedAtOk returns a tuple with the LastCheckedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCheckedAt

`func (o *BlacklistMonitor) SetLastCheckedAt(v time.Time)`

SetLastCheckedAt sets LastCheckedAt field to given value.

### HasLastCheckedAt

`func (o *BlacklistMonitor) HasLastCheckedAt() bool`

HasLastCheckedAt returns a boolean if a field has been set.

### SetLastCheckedAtNil

`func (o *BlacklistMonitor) SetLastCheckedAtNil(b bool)`

 SetLastCheckedAtNil sets the value for LastCheckedAt to be an explicit nil

### UnsetLastCheckedAt
`func (o *BlacklistMonitor) UnsetLastCheckedAt()`

UnsetLastCheckedAt ensures that no value is present for LastCheckedAt, not even an explicit nil
### GetLatestCheck

`func (o *BlacklistMonitor) GetLatestCheck() BlacklistMonitorLatestCheck`

GetLatestCheck returns the LatestCheck field if non-nil, zero value otherwise.

### GetLatestCheckOk

`func (o *BlacklistMonitor) GetLatestCheckOk() (*BlacklistMonitorLatestCheck, bool)`

GetLatestCheckOk returns a tuple with the LatestCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestCheck

`func (o *BlacklistMonitor) SetLatestCheck(v BlacklistMonitorLatestCheck)`

SetLatestCheck sets LatestCheck field to given value.

### HasLatestCheck

`func (o *BlacklistMonitor) HasLatestCheck() bool`

HasLatestCheck returns a boolean if a field has been set.

### SetLatestCheckNil

`func (o *BlacklistMonitor) SetLatestCheckNil(b bool)`

 SetLatestCheckNil sets the value for LatestCheck to be an explicit nil

### UnsetLatestCheck
`func (o *BlacklistMonitor) UnsetLatestCheck()`

UnsetLatestCheck ensures that no value is present for LatestCheck, not even an explicit nil
### GetCreatedAt

`func (o *BlacklistMonitor) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BlacklistMonitor) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BlacklistMonitor) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *BlacklistMonitor) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


