# TrackEventResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SchemaVersion** | Pointer to **string** |  | [optional] 
**RequestId** | Pointer to **string** |  | [optional] 
**EventId** | Pointer to **int32** | Event ID | [optional] 
**Created** | Pointer to **bool** | True if newly created, false if idempotent duplicate | [optional] 
**Idempotent** | Pointer to **bool** | Present and true when returning an existing event | [optional] 

## Methods

### NewTrackEventResponse

`func NewTrackEventResponse() *TrackEventResponse`

NewTrackEventResponse instantiates a new TrackEventResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrackEventResponseWithDefaults

`func NewTrackEventResponseWithDefaults() *TrackEventResponse`

NewTrackEventResponseWithDefaults instantiates a new TrackEventResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemaVersion

`func (o *TrackEventResponse) GetSchemaVersion() string`

GetSchemaVersion returns the SchemaVersion field if non-nil, zero value otherwise.

### GetSchemaVersionOk

`func (o *TrackEventResponse) GetSchemaVersionOk() (*string, bool)`

GetSchemaVersionOk returns a tuple with the SchemaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaVersion

`func (o *TrackEventResponse) SetSchemaVersion(v string)`

SetSchemaVersion sets SchemaVersion field to given value.

### HasSchemaVersion

`func (o *TrackEventResponse) HasSchemaVersion() bool`

HasSchemaVersion returns a boolean if a field has been set.

### GetRequestId

`func (o *TrackEventResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *TrackEventResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *TrackEventResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *TrackEventResponse) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetEventId

`func (o *TrackEventResponse) GetEventId() int32`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *TrackEventResponse) GetEventIdOk() (*int32, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *TrackEventResponse) SetEventId(v int32)`

SetEventId sets EventId field to given value.

### HasEventId

`func (o *TrackEventResponse) HasEventId() bool`

HasEventId returns a boolean if a field has been set.

### GetCreated

`func (o *TrackEventResponse) GetCreated() bool`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *TrackEventResponse) GetCreatedOk() (*bool, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *TrackEventResponse) SetCreated(v bool)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *TrackEventResponse) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetIdempotent

`func (o *TrackEventResponse) GetIdempotent() bool`

GetIdempotent returns the Idempotent field if non-nil, zero value otherwise.

### GetIdempotentOk

`func (o *TrackEventResponse) GetIdempotentOk() (*bool, bool)`

GetIdempotentOk returns a tuple with the Idempotent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdempotent

`func (o *TrackEventResponse) SetIdempotent(v bool)`

SetIdempotent sets Idempotent field to given value.

### HasIdempotent

`func (o *TrackEventResponse) HasIdempotent() bool`

HasIdempotent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


