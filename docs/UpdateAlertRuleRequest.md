# UpdateAlertRuleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metric** | Pointer to **string** |  | [optional] 
**Threshold** | Pointer to **float32** |  | [optional] 
**Channel** | Pointer to **string** |  | [optional] 
**WindowMinutes** | Pointer to **int32** | Evaluation window in minutes (15, 60, 1440, or 2880) | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 

## Methods

### NewUpdateAlertRuleRequest

`func NewUpdateAlertRuleRequest() *UpdateAlertRuleRequest`

NewUpdateAlertRuleRequest instantiates a new UpdateAlertRuleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAlertRuleRequestWithDefaults

`func NewUpdateAlertRuleRequestWithDefaults() *UpdateAlertRuleRequest`

NewUpdateAlertRuleRequestWithDefaults instantiates a new UpdateAlertRuleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetric

`func (o *UpdateAlertRuleRequest) GetMetric() string`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *UpdateAlertRuleRequest) GetMetricOk() (*string, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *UpdateAlertRuleRequest) SetMetric(v string)`

SetMetric sets Metric field to given value.

### HasMetric

`func (o *UpdateAlertRuleRequest) HasMetric() bool`

HasMetric returns a boolean if a field has been set.

### GetThreshold

`func (o *UpdateAlertRuleRequest) GetThreshold() float32`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *UpdateAlertRuleRequest) GetThresholdOk() (*float32, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *UpdateAlertRuleRequest) SetThreshold(v float32)`

SetThreshold sets Threshold field to given value.

### HasThreshold

`func (o *UpdateAlertRuleRequest) HasThreshold() bool`

HasThreshold returns a boolean if a field has been set.

### GetChannel

`func (o *UpdateAlertRuleRequest) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *UpdateAlertRuleRequest) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *UpdateAlertRuleRequest) SetChannel(v string)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *UpdateAlertRuleRequest) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetWindowMinutes

`func (o *UpdateAlertRuleRequest) GetWindowMinutes() int32`

GetWindowMinutes returns the WindowMinutes field if non-nil, zero value otherwise.

### GetWindowMinutesOk

`func (o *UpdateAlertRuleRequest) GetWindowMinutesOk() (*int32, bool)`

GetWindowMinutesOk returns a tuple with the WindowMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowMinutes

`func (o *UpdateAlertRuleRequest) SetWindowMinutes(v int32)`

SetWindowMinutes sets WindowMinutes field to given value.

### HasWindowMinutes

`func (o *UpdateAlertRuleRequest) HasWindowMinutes() bool`

HasWindowMinutes returns a boolean if a field has been set.

### GetEnabled

`func (o *UpdateAlertRuleRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UpdateAlertRuleRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UpdateAlertRuleRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *UpdateAlertRuleRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


