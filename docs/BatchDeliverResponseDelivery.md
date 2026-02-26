# BatchDeliverResponseDelivery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pool** | Pointer to **string** | IP pool used | [optional] 
**Lane** | Pointer to **string** | Delivery lane (green or yellow) | [optional] 
**QueuedAt** | Pointer to **time.Time** | Timestamp when batch was queued | [optional] 

## Methods

### NewBatchDeliverResponseDelivery

`func NewBatchDeliverResponseDelivery() *BatchDeliverResponseDelivery`

NewBatchDeliverResponseDelivery instantiates a new BatchDeliverResponseDelivery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchDeliverResponseDeliveryWithDefaults

`func NewBatchDeliverResponseDeliveryWithDefaults() *BatchDeliverResponseDelivery`

NewBatchDeliverResponseDeliveryWithDefaults instantiates a new BatchDeliverResponseDelivery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPool

`func (o *BatchDeliverResponseDelivery) GetPool() string`

GetPool returns the Pool field if non-nil, zero value otherwise.

### GetPoolOk

`func (o *BatchDeliverResponseDelivery) GetPoolOk() (*string, bool)`

GetPoolOk returns a tuple with the Pool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool

`func (o *BatchDeliverResponseDelivery) SetPool(v string)`

SetPool sets Pool field to given value.

### HasPool

`func (o *BatchDeliverResponseDelivery) HasPool() bool`

HasPool returns a boolean if a field has been set.

### GetLane

`func (o *BatchDeliverResponseDelivery) GetLane() string`

GetLane returns the Lane field if non-nil, zero value otherwise.

### GetLaneOk

`func (o *BatchDeliverResponseDelivery) GetLaneOk() (*string, bool)`

GetLaneOk returns a tuple with the Lane field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLane

`func (o *BatchDeliverResponseDelivery) SetLane(v string)`

SetLane sets Lane field to given value.

### HasLane

`func (o *BatchDeliverResponseDelivery) HasLane() bool`

HasLane returns a boolean if a field has been set.

### GetQueuedAt

`func (o *BatchDeliverResponseDelivery) GetQueuedAt() time.Time`

GetQueuedAt returns the QueuedAt field if non-nil, zero value otherwise.

### GetQueuedAtOk

`func (o *BatchDeliverResponseDelivery) GetQueuedAtOk() (*time.Time, bool)`

GetQueuedAtOk returns a tuple with the QueuedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueuedAt

`func (o *BatchDeliverResponseDelivery) SetQueuedAt(v time.Time)`

SetQueuedAt sets QueuedAt field to given value.

### HasQueuedAt

`func (o *BatchDeliverResponseDelivery) HasQueuedAt() bool`

HasQueuedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


