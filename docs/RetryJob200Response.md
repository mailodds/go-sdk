# RetryJob200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SchemaVersion** | Pointer to **string** |  | [optional] 
**RequestId** | Pointer to **string** |  | [optional] 
**Job** | Pointer to **map[string]interface{}** |  | [optional] 
**PendingCount** | Pointer to **int32** | Number of emails re-queued | [optional] 

## Methods

### NewRetryJob200Response

`func NewRetryJob200Response() *RetryJob200Response`

NewRetryJob200Response instantiates a new RetryJob200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRetryJob200ResponseWithDefaults

`func NewRetryJob200ResponseWithDefaults() *RetryJob200Response`

NewRetryJob200ResponseWithDefaults instantiates a new RetryJob200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemaVersion

`func (o *RetryJob200Response) GetSchemaVersion() string`

GetSchemaVersion returns the SchemaVersion field if non-nil, zero value otherwise.

### GetSchemaVersionOk

`func (o *RetryJob200Response) GetSchemaVersionOk() (*string, bool)`

GetSchemaVersionOk returns a tuple with the SchemaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaVersion

`func (o *RetryJob200Response) SetSchemaVersion(v string)`

SetSchemaVersion sets SchemaVersion field to given value.

### HasSchemaVersion

`func (o *RetryJob200Response) HasSchemaVersion() bool`

HasSchemaVersion returns a boolean if a field has been set.

### GetRequestId

`func (o *RetryJob200Response) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *RetryJob200Response) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *RetryJob200Response) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *RetryJob200Response) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetJob

`func (o *RetryJob200Response) GetJob() map[string]interface{}`

GetJob returns the Job field if non-nil, zero value otherwise.

### GetJobOk

`func (o *RetryJob200Response) GetJobOk() (*map[string]interface{}, bool)`

GetJobOk returns a tuple with the Job field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJob

`func (o *RetryJob200Response) SetJob(v map[string]interface{})`

SetJob sets Job field to given value.

### HasJob

`func (o *RetryJob200Response) HasJob() bool`

HasJob returns a boolean if a field has been set.

### GetPendingCount

`func (o *RetryJob200Response) GetPendingCount() int32`

GetPendingCount returns the PendingCount field if non-nil, zero value otherwise.

### GetPendingCountOk

`func (o *RetryJob200Response) GetPendingCountOk() (*int32, bool)`

GetPendingCountOk returns a tuple with the PendingCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingCount

`func (o *RetryJob200Response) SetPendingCount(v int32)`

SetPendingCount sets PendingCount field to given value.

### HasPendingCount

`func (o *RetryJob200Response) HasPendingCount() bool`

HasPendingCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


