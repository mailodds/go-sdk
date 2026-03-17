# SyncResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SchemaVersion** | Pointer to **string** |  | [optional] 
**RequestId** | Pointer to **string** |  | [optional] 
**Scheduled** | Pointer to **bool** |  | [optional] 
**StoreId** | Pointer to **string** |  | [optional] 
**Idempotent** | Pointer to **bool** | True if this was a duplicate request matched by Idempotency-Key | [optional] 
**ExistingJobId** | Pointer to **string** | ID of existing sync job if one was already running | [optional] 

## Methods

### NewSyncResponse

`func NewSyncResponse() *SyncResponse`

NewSyncResponse instantiates a new SyncResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyncResponseWithDefaults

`func NewSyncResponseWithDefaults() *SyncResponse`

NewSyncResponseWithDefaults instantiates a new SyncResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemaVersion

`func (o *SyncResponse) GetSchemaVersion() string`

GetSchemaVersion returns the SchemaVersion field if non-nil, zero value otherwise.

### GetSchemaVersionOk

`func (o *SyncResponse) GetSchemaVersionOk() (*string, bool)`

GetSchemaVersionOk returns a tuple with the SchemaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaVersion

`func (o *SyncResponse) SetSchemaVersion(v string)`

SetSchemaVersion sets SchemaVersion field to given value.

### HasSchemaVersion

`func (o *SyncResponse) HasSchemaVersion() bool`

HasSchemaVersion returns a boolean if a field has been set.

### GetRequestId

`func (o *SyncResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *SyncResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *SyncResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *SyncResponse) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetScheduled

`func (o *SyncResponse) GetScheduled() bool`

GetScheduled returns the Scheduled field if non-nil, zero value otherwise.

### GetScheduledOk

`func (o *SyncResponse) GetScheduledOk() (*bool, bool)`

GetScheduledOk returns a tuple with the Scheduled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduled

`func (o *SyncResponse) SetScheduled(v bool)`

SetScheduled sets Scheduled field to given value.

### HasScheduled

`func (o *SyncResponse) HasScheduled() bool`

HasScheduled returns a boolean if a field has been set.

### GetStoreId

`func (o *SyncResponse) GetStoreId() string`

GetStoreId returns the StoreId field if non-nil, zero value otherwise.

### GetStoreIdOk

`func (o *SyncResponse) GetStoreIdOk() (*string, bool)`

GetStoreIdOk returns a tuple with the StoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreId

`func (o *SyncResponse) SetStoreId(v string)`

SetStoreId sets StoreId field to given value.

### HasStoreId

`func (o *SyncResponse) HasStoreId() bool`

HasStoreId returns a boolean if a field has been set.

### GetIdempotent

`func (o *SyncResponse) GetIdempotent() bool`

GetIdempotent returns the Idempotent field if non-nil, zero value otherwise.

### GetIdempotentOk

`func (o *SyncResponse) GetIdempotentOk() (*bool, bool)`

GetIdempotentOk returns a tuple with the Idempotent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdempotent

`func (o *SyncResponse) SetIdempotent(v bool)`

SetIdempotent sets Idempotent field to given value.

### HasIdempotent

`func (o *SyncResponse) HasIdempotent() bool`

HasIdempotent returns a boolean if a field has been set.

### GetExistingJobId

`func (o *SyncResponse) GetExistingJobId() string`

GetExistingJobId returns the ExistingJobId field if non-nil, zero value otherwise.

### GetExistingJobIdOk

`func (o *SyncResponse) GetExistingJobIdOk() (*string, bool)`

GetExistingJobIdOk returns a tuple with the ExistingJobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExistingJobId

`func (o *SyncResponse) SetExistingJobId(v string)`

SetExistingJobId sets ExistingJobId field to given value.

### HasExistingJobId

`func (o *SyncResponse) HasExistingJobId() bool`

HasExistingJobId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


