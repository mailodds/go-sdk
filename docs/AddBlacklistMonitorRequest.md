# AddBlacklistMonitorRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Target** | **string** | IP address or domain to monitor | 
**TargetType** | **string** | Whether the target is an IP or domain | 

## Methods

### NewAddBlacklistMonitorRequest

`func NewAddBlacklistMonitorRequest(target string, targetType string, ) *AddBlacklistMonitorRequest`

NewAddBlacklistMonitorRequest instantiates a new AddBlacklistMonitorRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddBlacklistMonitorRequestWithDefaults

`func NewAddBlacklistMonitorRequestWithDefaults() *AddBlacklistMonitorRequest`

NewAddBlacklistMonitorRequestWithDefaults instantiates a new AddBlacklistMonitorRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTarget

`func (o *AddBlacklistMonitorRequest) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *AddBlacklistMonitorRequest) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *AddBlacklistMonitorRequest) SetTarget(v string)`

SetTarget sets Target field to given value.


### GetTargetType

`func (o *AddBlacklistMonitorRequest) GetTargetType() string`

GetTargetType returns the TargetType field if non-nil, zero value otherwise.

### GetTargetTypeOk

`func (o *AddBlacklistMonitorRequest) GetTargetTypeOk() (*string, bool)`

GetTargetTypeOk returns a tuple with the TargetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetType

`func (o *AddBlacklistMonitorRequest) SetTargetType(v string)`

SetTargetType sets TargetType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


