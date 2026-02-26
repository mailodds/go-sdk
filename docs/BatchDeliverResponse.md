# BatchDeliverResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SchemaVersion** | Pointer to **string** |  | [optional] 
**RequestId** | Pointer to **string** | Unique request identifier | [optional] 
**Total** | Pointer to **int32** | Total recipients submitted | [optional] 
**Accepted** | Pointer to **int32** | Number of recipients accepted for delivery | [optional] 
**Rejected** | Pointer to [**[]BatchDeliverResponseRejectedInner**](BatchDeliverResponseRejectedInner.md) | Recipients that were rejected (suppressed or failed validation) | [optional] 
**Status** | Pointer to **string** | Batch status | [optional] 
**Delivery** | Pointer to [**BatchDeliverResponseDelivery**](BatchDeliverResponseDelivery.md) |  | [optional] 

## Methods

### NewBatchDeliverResponse

`func NewBatchDeliverResponse() *BatchDeliverResponse`

NewBatchDeliverResponse instantiates a new BatchDeliverResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchDeliverResponseWithDefaults

`func NewBatchDeliverResponseWithDefaults() *BatchDeliverResponse`

NewBatchDeliverResponseWithDefaults instantiates a new BatchDeliverResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemaVersion

`func (o *BatchDeliverResponse) GetSchemaVersion() string`

GetSchemaVersion returns the SchemaVersion field if non-nil, zero value otherwise.

### GetSchemaVersionOk

`func (o *BatchDeliverResponse) GetSchemaVersionOk() (*string, bool)`

GetSchemaVersionOk returns a tuple with the SchemaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaVersion

`func (o *BatchDeliverResponse) SetSchemaVersion(v string)`

SetSchemaVersion sets SchemaVersion field to given value.

### HasSchemaVersion

`func (o *BatchDeliverResponse) HasSchemaVersion() bool`

HasSchemaVersion returns a boolean if a field has been set.

### GetRequestId

`func (o *BatchDeliverResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *BatchDeliverResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *BatchDeliverResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *BatchDeliverResponse) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetTotal

`func (o *BatchDeliverResponse) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *BatchDeliverResponse) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *BatchDeliverResponse) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *BatchDeliverResponse) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetAccepted

`func (o *BatchDeliverResponse) GetAccepted() int32`

GetAccepted returns the Accepted field if non-nil, zero value otherwise.

### GetAcceptedOk

`func (o *BatchDeliverResponse) GetAcceptedOk() (*int32, bool)`

GetAcceptedOk returns a tuple with the Accepted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccepted

`func (o *BatchDeliverResponse) SetAccepted(v int32)`

SetAccepted sets Accepted field to given value.

### HasAccepted

`func (o *BatchDeliverResponse) HasAccepted() bool`

HasAccepted returns a boolean if a field has been set.

### GetRejected

`func (o *BatchDeliverResponse) GetRejected() []BatchDeliverResponseRejectedInner`

GetRejected returns the Rejected field if non-nil, zero value otherwise.

### GetRejectedOk

`func (o *BatchDeliverResponse) GetRejectedOk() (*[]BatchDeliverResponseRejectedInner, bool)`

GetRejectedOk returns a tuple with the Rejected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejected

`func (o *BatchDeliverResponse) SetRejected(v []BatchDeliverResponseRejectedInner)`

SetRejected sets Rejected field to given value.

### HasRejected

`func (o *BatchDeliverResponse) HasRejected() bool`

HasRejected returns a boolean if a field has been set.

### GetStatus

`func (o *BatchDeliverResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BatchDeliverResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BatchDeliverResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BatchDeliverResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDelivery

`func (o *BatchDeliverResponse) GetDelivery() BatchDeliverResponseDelivery`

GetDelivery returns the Delivery field if non-nil, zero value otherwise.

### GetDeliveryOk

`func (o *BatchDeliverResponse) GetDeliveryOk() (*BatchDeliverResponseDelivery, bool)`

GetDeliveryOk returns a tuple with the Delivery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelivery

`func (o *BatchDeliverResponse) SetDelivery(v BatchDeliverResponseDelivery)`

SetDelivery sets Delivery field to given value.

### HasDelivery

`func (o *BatchDeliverResponse) HasDelivery() bool`

HasDelivery returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


