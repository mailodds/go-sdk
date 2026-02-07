# TelemetrySummaryRates

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Deliverable** | Pointer to **float32** | Rate of valid + catch_all emails | [optional] 
**Reject** | Pointer to **float32** | Rate of invalid + do_not_mail emails | [optional] 
**Unknown** | Pointer to **float32** | Rate of unknown status | [optional] 
**Suppressed** | Pointer to **float32** | Rate of suppressed emails | [optional] 

## Methods

### NewTelemetrySummaryRates

`func NewTelemetrySummaryRates() *TelemetrySummaryRates`

NewTelemetrySummaryRates instantiates a new TelemetrySummaryRates object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTelemetrySummaryRatesWithDefaults

`func NewTelemetrySummaryRatesWithDefaults() *TelemetrySummaryRates`

NewTelemetrySummaryRatesWithDefaults instantiates a new TelemetrySummaryRates object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeliverable

`func (o *TelemetrySummaryRates) GetDeliverable() float32`

GetDeliverable returns the Deliverable field if non-nil, zero value otherwise.

### GetDeliverableOk

`func (o *TelemetrySummaryRates) GetDeliverableOk() (*float32, bool)`

GetDeliverableOk returns a tuple with the Deliverable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliverable

`func (o *TelemetrySummaryRates) SetDeliverable(v float32)`

SetDeliverable sets Deliverable field to given value.

### HasDeliverable

`func (o *TelemetrySummaryRates) HasDeliverable() bool`

HasDeliverable returns a boolean if a field has been set.

### GetReject

`func (o *TelemetrySummaryRates) GetReject() float32`

GetReject returns the Reject field if non-nil, zero value otherwise.

### GetRejectOk

`func (o *TelemetrySummaryRates) GetRejectOk() (*float32, bool)`

GetRejectOk returns a tuple with the Reject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReject

`func (o *TelemetrySummaryRates) SetReject(v float32)`

SetReject sets Reject field to given value.

### HasReject

`func (o *TelemetrySummaryRates) HasReject() bool`

HasReject returns a boolean if a field has been set.

### GetUnknown

`func (o *TelemetrySummaryRates) GetUnknown() float32`

GetUnknown returns the Unknown field if non-nil, zero value otherwise.

### GetUnknownOk

`func (o *TelemetrySummaryRates) GetUnknownOk() (*float32, bool)`

GetUnknownOk returns a tuple with the Unknown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnknown

`func (o *TelemetrySummaryRates) SetUnknown(v float32)`

SetUnknown sets Unknown field to given value.

### HasUnknown

`func (o *TelemetrySummaryRates) HasUnknown() bool`

HasUnknown returns a boolean if a field has been set.

### GetSuppressed

`func (o *TelemetrySummaryRates) GetSuppressed() float32`

GetSuppressed returns the Suppressed field if non-nil, zero value otherwise.

### GetSuppressedOk

`func (o *TelemetrySummaryRates) GetSuppressedOk() (*float32, bool)`

GetSuppressedOk returns a tuple with the Suppressed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppressed

`func (o *TelemetrySummaryRates) SetSuppressed(v float32)`

SetSuppressed sets Suppressed field to given value.

### HasSuppressed

`func (o *TelemetrySummaryRates) HasSuppressed() bool`

HasSuppressed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


