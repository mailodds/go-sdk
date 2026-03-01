# Job

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** | Job name (from metadata or auto-generated) | 
**Status** | **string** |  | 
**TotalCount** | **int32** |  | 
**ProcessedCount** | **int32** |  | 
**Summary** | Pointer to [**JobSummary**](JobSummary.md) |  | [optional] 
**CreatedAt** | **time.Time** |  | 
**StartedAt** | Pointer to **time.Time** | When processing began. Omitted if not yet started. | [optional] 
**CompletedAt** | Pointer to **time.Time** | Omitted if not yet completed. | [optional] 
**ResultsExpireAt** | **time.Time** | When job results will be purged | 
**Metadata** | Pointer to **map[string]interface{}** | Custom metadata attached at creation | [optional] 
**ErrorMessage** | Pointer to **string** | Error details. Present only for failed jobs. | [optional] 
**RequestId** | Pointer to **string** | Request ID from the job creation request | [optional] 
**Artifacts** | Pointer to [**JobArtifacts**](JobArtifacts.md) |  | [optional] 

## Methods

### NewJob

`func NewJob(id string, name string, status string, totalCount int32, processedCount int32, createdAt time.Time, resultsExpireAt time.Time, ) *Job`

NewJob instantiates a new Job object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobWithDefaults

`func NewJobWithDefaults() *Job`

NewJobWithDefaults instantiates a new Job object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Job) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Job) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Job) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *Job) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Job) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Job) SetName(v string)`

SetName sets Name field to given value.


### GetStatus

`func (o *Job) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Job) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Job) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTotalCount

`func (o *Job) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *Job) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *Job) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.


### GetProcessedCount

`func (o *Job) GetProcessedCount() int32`

GetProcessedCount returns the ProcessedCount field if non-nil, zero value otherwise.

### GetProcessedCountOk

`func (o *Job) GetProcessedCountOk() (*int32, bool)`

GetProcessedCountOk returns a tuple with the ProcessedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessedCount

`func (o *Job) SetProcessedCount(v int32)`

SetProcessedCount sets ProcessedCount field to given value.


### GetSummary

`func (o *Job) GetSummary() JobSummary`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *Job) GetSummaryOk() (*JobSummary, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *Job) SetSummary(v JobSummary)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *Job) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Job) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Job) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Job) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetStartedAt

`func (o *Job) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *Job) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *Job) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *Job) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetCompletedAt

`func (o *Job) GetCompletedAt() time.Time`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *Job) GetCompletedAtOk() (*time.Time, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *Job) SetCompletedAt(v time.Time)`

SetCompletedAt sets CompletedAt field to given value.

### HasCompletedAt

`func (o *Job) HasCompletedAt() bool`

HasCompletedAt returns a boolean if a field has been set.

### GetResultsExpireAt

`func (o *Job) GetResultsExpireAt() time.Time`

GetResultsExpireAt returns the ResultsExpireAt field if non-nil, zero value otherwise.

### GetResultsExpireAtOk

`func (o *Job) GetResultsExpireAtOk() (*time.Time, bool)`

GetResultsExpireAtOk returns a tuple with the ResultsExpireAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultsExpireAt

`func (o *Job) SetResultsExpireAt(v time.Time)`

SetResultsExpireAt sets ResultsExpireAt field to given value.


### GetMetadata

`func (o *Job) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Job) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Job) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Job) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetErrorMessage

`func (o *Job) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *Job) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *Job) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *Job) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetRequestId

`func (o *Job) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *Job) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *Job) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *Job) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetArtifacts

`func (o *Job) GetArtifacts() JobArtifacts`

GetArtifacts returns the Artifacts field if non-nil, zero value otherwise.

### GetArtifactsOk

`func (o *Job) GetArtifactsOk() (*JobArtifacts, bool)`

GetArtifactsOk returns a tuple with the Artifacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifacts

`func (o *Job) SetArtifacts(v JobArtifacts)`

SetArtifacts sets Artifacts field to given value.

### HasArtifacts

`func (o *Job) HasArtifacts() bool`

HasArtifacts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


