# BatchDeliverResponseRejectedInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **string** |  | [optional] 
**Reason** | Pointer to **string** | Rejection reason (suppressed, validation_rejected) | [optional] 
**Status** | Pointer to **string** | Validation status if rejected by validation | [optional] 
**SubStatus** | Pointer to **string** | Validation sub-status if rejected by validation | [optional] 

## Methods

### NewBatchDeliverResponseRejectedInner

`func NewBatchDeliverResponseRejectedInner() *BatchDeliverResponseRejectedInner`

NewBatchDeliverResponseRejectedInner instantiates a new BatchDeliverResponseRejectedInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchDeliverResponseRejectedInnerWithDefaults

`func NewBatchDeliverResponseRejectedInnerWithDefaults() *BatchDeliverResponseRejectedInner`

NewBatchDeliverResponseRejectedInnerWithDefaults instantiates a new BatchDeliverResponseRejectedInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *BatchDeliverResponseRejectedInner) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *BatchDeliverResponseRejectedInner) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *BatchDeliverResponseRejectedInner) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *BatchDeliverResponseRejectedInner) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetReason

`func (o *BatchDeliverResponseRejectedInner) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *BatchDeliverResponseRejectedInner) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *BatchDeliverResponseRejectedInner) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *BatchDeliverResponseRejectedInner) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetStatus

`func (o *BatchDeliverResponseRejectedInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BatchDeliverResponseRejectedInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BatchDeliverResponseRejectedInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BatchDeliverResponseRejectedInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSubStatus

`func (o *BatchDeliverResponseRejectedInner) GetSubStatus() string`

GetSubStatus returns the SubStatus field if non-nil, zero value otherwise.

### GetSubStatusOk

`func (o *BatchDeliverResponseRejectedInner) GetSubStatusOk() (*string, bool)`

GetSubStatusOk returns a tuple with the SubStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubStatus

`func (o *BatchDeliverResponseRejectedInner) SetSubStatus(v string)`

SetSubStatus sets SubStatus field to given value.

### HasSubStatus

`func (o *BatchDeliverResponseRejectedInner) HasSubStatus() bool`

HasSubStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


