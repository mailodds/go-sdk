# DeliverResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SchemaVersion** | Pointer to **string** |  | [optional] 
**RequestId** | Pointer to **string** | Unique request identifier | [optional] 
**MessageId** | Pointer to **string** | Unique message identifier | [optional] 
**Status** | Pointer to **string** | Delivery status | [optional] 
**Delivery** | Pointer to [**DeliverResponseDelivery**](DeliverResponseDelivery.md) |  | [optional] 
**Validation** | Pointer to **map[string]interface{}** | Pre-send validation results (when validate_first is true) | [optional] 
**ContentScan** | Pointer to **map[string]interface{}** | Content scan results | [optional] 

## Methods

### NewDeliverResponse

`func NewDeliverResponse() *DeliverResponse`

NewDeliverResponse instantiates a new DeliverResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeliverResponseWithDefaults

`func NewDeliverResponseWithDefaults() *DeliverResponse`

NewDeliverResponseWithDefaults instantiates a new DeliverResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemaVersion

`func (o *DeliverResponse) GetSchemaVersion() string`

GetSchemaVersion returns the SchemaVersion field if non-nil, zero value otherwise.

### GetSchemaVersionOk

`func (o *DeliverResponse) GetSchemaVersionOk() (*string, bool)`

GetSchemaVersionOk returns a tuple with the SchemaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaVersion

`func (o *DeliverResponse) SetSchemaVersion(v string)`

SetSchemaVersion sets SchemaVersion field to given value.

### HasSchemaVersion

`func (o *DeliverResponse) HasSchemaVersion() bool`

HasSchemaVersion returns a boolean if a field has been set.

### GetRequestId

`func (o *DeliverResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *DeliverResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *DeliverResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *DeliverResponse) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetMessageId

`func (o *DeliverResponse) GetMessageId() string`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *DeliverResponse) GetMessageIdOk() (*string, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *DeliverResponse) SetMessageId(v string)`

SetMessageId sets MessageId field to given value.

### HasMessageId

`func (o *DeliverResponse) HasMessageId() bool`

HasMessageId returns a boolean if a field has been set.

### GetStatus

`func (o *DeliverResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DeliverResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DeliverResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DeliverResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDelivery

`func (o *DeliverResponse) GetDelivery() DeliverResponseDelivery`

GetDelivery returns the Delivery field if non-nil, zero value otherwise.

### GetDeliveryOk

`func (o *DeliverResponse) GetDeliveryOk() (*DeliverResponseDelivery, bool)`

GetDeliveryOk returns a tuple with the Delivery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelivery

`func (o *DeliverResponse) SetDelivery(v DeliverResponseDelivery)`

SetDelivery sets Delivery field to given value.

### HasDelivery

`func (o *DeliverResponse) HasDelivery() bool`

HasDelivery returns a boolean if a field has been set.

### GetValidation

`func (o *DeliverResponse) GetValidation() map[string]interface{}`

GetValidation returns the Validation field if non-nil, zero value otherwise.

### GetValidationOk

`func (o *DeliverResponse) GetValidationOk() (*map[string]interface{}, bool)`

GetValidationOk returns a tuple with the Validation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidation

`func (o *DeliverResponse) SetValidation(v map[string]interface{})`

SetValidation sets Validation field to given value.

### HasValidation

`func (o *DeliverResponse) HasValidation() bool`

HasValidation returns a boolean if a field has been set.

### GetContentScan

`func (o *DeliverResponse) GetContentScan() map[string]interface{}`

GetContentScan returns the ContentScan field if non-nil, zero value otherwise.

### GetContentScanOk

`func (o *DeliverResponse) GetContentScanOk() (*map[string]interface{}, bool)`

GetContentScanOk returns a tuple with the ContentScan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentScan

`func (o *DeliverResponse) SetContentScan(v map[string]interface{})`

SetContentScan sets ContentScan field to given value.

### HasContentScan

`func (o *DeliverResponse) HasContentScan() bool`

HasContentScan returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


