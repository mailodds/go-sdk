# AlertRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Metric** | Pointer to **string** | Monitored metric name | [optional] 
**Threshold** | Pointer to **float32** | Alert threshold value (0-1) | [optional] 
**Channel** | Pointer to **string** | Notification channel | [optional] 
**WindowMinutes** | Pointer to **int32** | Evaluation window in minutes (15, 60, 1440, or 2880) | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewAlertRule

`func NewAlertRule() *AlertRule`

NewAlertRule instantiates a new AlertRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertRuleWithDefaults

`func NewAlertRuleWithDefaults() *AlertRule`

NewAlertRuleWithDefaults instantiates a new AlertRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AlertRule) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AlertRule) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AlertRule) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AlertRule) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetric

`func (o *AlertRule) GetMetric() string`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *AlertRule) GetMetricOk() (*string, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *AlertRule) SetMetric(v string)`

SetMetric sets Metric field to given value.

### HasMetric

`func (o *AlertRule) HasMetric() bool`

HasMetric returns a boolean if a field has been set.

### GetThreshold

`func (o *AlertRule) GetThreshold() float32`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *AlertRule) GetThresholdOk() (*float32, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *AlertRule) SetThreshold(v float32)`

SetThreshold sets Threshold field to given value.

### HasThreshold

`func (o *AlertRule) HasThreshold() bool`

HasThreshold returns a boolean if a field has been set.

### GetChannel

`func (o *AlertRule) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *AlertRule) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *AlertRule) SetChannel(v string)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *AlertRule) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetWindowMinutes

`func (o *AlertRule) GetWindowMinutes() int32`

GetWindowMinutes returns the WindowMinutes field if non-nil, zero value otherwise.

### GetWindowMinutesOk

`func (o *AlertRule) GetWindowMinutesOk() (*int32, bool)`

GetWindowMinutesOk returns a tuple with the WindowMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowMinutes

`func (o *AlertRule) SetWindowMinutes(v int32)`

SetWindowMinutes sets WindowMinutes field to given value.

### HasWindowMinutes

`func (o *AlertRule) HasWindowMinutes() bool`

HasWindowMinutes returns a boolean if a field has been set.

### GetEnabled

`func (o *AlertRule) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AlertRule) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AlertRule) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *AlertRule) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AlertRule) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AlertRule) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AlertRule) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AlertRule) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *AlertRule) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AlertRule) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AlertRule) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AlertRule) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


