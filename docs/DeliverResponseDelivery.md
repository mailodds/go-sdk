# DeliverResponseDelivery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pool** | Pointer to **string** | IP pool used | [optional] 
**Lane** | Pointer to **string** | Delivery lane | [optional] 
**WarmupLimited** | Pointer to **bool** | Whether warmup throttling was applied | [optional] 

## Methods

### NewDeliverResponseDelivery

`func NewDeliverResponseDelivery() *DeliverResponseDelivery`

NewDeliverResponseDelivery instantiates a new DeliverResponseDelivery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeliverResponseDeliveryWithDefaults

`func NewDeliverResponseDeliveryWithDefaults() *DeliverResponseDelivery`

NewDeliverResponseDeliveryWithDefaults instantiates a new DeliverResponseDelivery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPool

`func (o *DeliverResponseDelivery) GetPool() string`

GetPool returns the Pool field if non-nil, zero value otherwise.

### GetPoolOk

`func (o *DeliverResponseDelivery) GetPoolOk() (*string, bool)`

GetPoolOk returns a tuple with the Pool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool

`func (o *DeliverResponseDelivery) SetPool(v string)`

SetPool sets Pool field to given value.

### HasPool

`func (o *DeliverResponseDelivery) HasPool() bool`

HasPool returns a boolean if a field has been set.

### GetLane

`func (o *DeliverResponseDelivery) GetLane() string`

GetLane returns the Lane field if non-nil, zero value otherwise.

### GetLaneOk

`func (o *DeliverResponseDelivery) GetLaneOk() (*string, bool)`

GetLaneOk returns a tuple with the Lane field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLane

`func (o *DeliverResponseDelivery) SetLane(v string)`

SetLane sets Lane field to given value.

### HasLane

`func (o *DeliverResponseDelivery) HasLane() bool`

HasLane returns a boolean if a field has been set.

### GetWarmupLimited

`func (o *DeliverResponseDelivery) GetWarmupLimited() bool`

GetWarmupLimited returns the WarmupLimited field if non-nil, zero value otherwise.

### GetWarmupLimitedOk

`func (o *DeliverResponseDelivery) GetWarmupLimitedOk() (*bool, bool)`

GetWarmupLimitedOk returns a tuple with the WarmupLimited field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarmupLimited

`func (o *DeliverResponseDelivery) SetWarmupLimited(v bool)`

SetWarmupLimited sets WarmupLimited field to given value.

### HasWarmupLimited

`func (o *DeliverResponseDelivery) HasWarmupLimited() bool`

HasWarmupLimited returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


