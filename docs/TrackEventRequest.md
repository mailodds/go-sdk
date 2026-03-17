# TrackEventRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventType** | **string** | Type of commerce event | 
**Email** | **string** | Email address associated with the event | 
**Properties** | Pointer to **map[string]interface{}** | Event-specific data (e.g., order_id, value, product_url) | [optional] 
**OccurredAt** | Pointer to **time.Time** | When the event occurred (defaults to now) | [optional] 
**IdempotencyKey** | Pointer to **string** | Unique key to prevent duplicate events (5 min TTL) | [optional] 

## Methods

### NewTrackEventRequest

`func NewTrackEventRequest(eventType string, email string, ) *TrackEventRequest`

NewTrackEventRequest instantiates a new TrackEventRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrackEventRequestWithDefaults

`func NewTrackEventRequestWithDefaults() *TrackEventRequest`

NewTrackEventRequestWithDefaults instantiates a new TrackEventRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventType

`func (o *TrackEventRequest) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *TrackEventRequest) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *TrackEventRequest) SetEventType(v string)`

SetEventType sets EventType field to given value.


### GetEmail

`func (o *TrackEventRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *TrackEventRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *TrackEventRequest) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetProperties

`func (o *TrackEventRequest) GetProperties() map[string]interface{}`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *TrackEventRequest) GetPropertiesOk() (*map[string]interface{}, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *TrackEventRequest) SetProperties(v map[string]interface{})`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *TrackEventRequest) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetOccurredAt

`func (o *TrackEventRequest) GetOccurredAt() time.Time`

GetOccurredAt returns the OccurredAt field if non-nil, zero value otherwise.

### GetOccurredAtOk

`func (o *TrackEventRequest) GetOccurredAtOk() (*time.Time, bool)`

GetOccurredAtOk returns a tuple with the OccurredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccurredAt

`func (o *TrackEventRequest) SetOccurredAt(v time.Time)`

SetOccurredAt sets OccurredAt field to given value.

### HasOccurredAt

`func (o *TrackEventRequest) HasOccurredAt() bool`

HasOccurredAt returns a boolean if a field has been set.

### GetIdempotencyKey

`func (o *TrackEventRequest) GetIdempotencyKey() string`

GetIdempotencyKey returns the IdempotencyKey field if non-nil, zero value otherwise.

### GetIdempotencyKeyOk

`func (o *TrackEventRequest) GetIdempotencyKeyOk() (*string, bool)`

GetIdempotencyKeyOk returns a tuple with the IdempotencyKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdempotencyKey

`func (o *TrackEventRequest) SetIdempotencyKey(v string)`

SetIdempotencyKey sets IdempotencyKey field to given value.

### HasIdempotencyKey

`func (o *TrackEventRequest) HasIdempotencyKey() bool`

HasIdempotencyKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


