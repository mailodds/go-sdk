# CreateAlertRuleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metric** | **string** | Metric to monitor (e.g., bounce_rate, complaint_rate) | 
**Threshold** | **float32** | Threshold value to trigger alert | 
**Channel** | **string** | Notification channel (e.g., webhook) | 
**WindowMinutes** | Pointer to **int32** | Evaluation window in minutes | [optional] [default to 60]
**Enabled** | Pointer to **bool** |  | [optional] [default to true]

## Methods

### NewCreateAlertRuleRequest

`func NewCreateAlertRuleRequest(metric string, threshold float32, channel string, ) *CreateAlertRuleRequest`

NewCreateAlertRuleRequest instantiates a new CreateAlertRuleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAlertRuleRequestWithDefaults

`func NewCreateAlertRuleRequestWithDefaults() *CreateAlertRuleRequest`

NewCreateAlertRuleRequestWithDefaults instantiates a new CreateAlertRuleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetric

`func (o *CreateAlertRuleRequest) GetMetric() string`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *CreateAlertRuleRequest) GetMetricOk() (*string, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *CreateAlertRuleRequest) SetMetric(v string)`

SetMetric sets Metric field to given value.


### GetThreshold

`func (o *CreateAlertRuleRequest) GetThreshold() float32`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *CreateAlertRuleRequest) GetThresholdOk() (*float32, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *CreateAlertRuleRequest) SetThreshold(v float32)`

SetThreshold sets Threshold field to given value.


### GetChannel

`func (o *CreateAlertRuleRequest) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *CreateAlertRuleRequest) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *CreateAlertRuleRequest) SetChannel(v string)`

SetChannel sets Channel field to given value.


### GetWindowMinutes

`func (o *CreateAlertRuleRequest) GetWindowMinutes() int32`

GetWindowMinutes returns the WindowMinutes field if non-nil, zero value otherwise.

### GetWindowMinutesOk

`func (o *CreateAlertRuleRequest) GetWindowMinutesOk() (*int32, bool)`

GetWindowMinutesOk returns a tuple with the WindowMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowMinutes

`func (o *CreateAlertRuleRequest) SetWindowMinutes(v int32)`

SetWindowMinutes sets WindowMinutes field to given value.

### HasWindowMinutes

`func (o *CreateAlertRuleRequest) HasWindowMinutes() bool`

HasWindowMinutes returns a boolean if a field has been set.

### GetEnabled

`func (o *CreateAlertRuleRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CreateAlertRuleRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CreateAlertRuleRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CreateAlertRuleRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


