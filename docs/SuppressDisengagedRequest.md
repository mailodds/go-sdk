# SuppressDisengagedRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InactiveDays** | Pointer to **int32** | Days of inactivity threshold | [optional] [default to 180]
**MinSends** | Pointer to **int32** | Minimum sends to qualify | [optional] [default to 10]
**DryRun** | Pointer to **bool** | Preview without suppressing | [optional] [default to true]

## Methods

### NewSuppressDisengagedRequest

`func NewSuppressDisengagedRequest() *SuppressDisengagedRequest`

NewSuppressDisengagedRequest instantiates a new SuppressDisengagedRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSuppressDisengagedRequestWithDefaults

`func NewSuppressDisengagedRequestWithDefaults() *SuppressDisengagedRequest`

NewSuppressDisengagedRequestWithDefaults instantiates a new SuppressDisengagedRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInactiveDays

`func (o *SuppressDisengagedRequest) GetInactiveDays() int32`

GetInactiveDays returns the InactiveDays field if non-nil, zero value otherwise.

### GetInactiveDaysOk

`func (o *SuppressDisengagedRequest) GetInactiveDaysOk() (*int32, bool)`

GetInactiveDaysOk returns a tuple with the InactiveDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInactiveDays

`func (o *SuppressDisengagedRequest) SetInactiveDays(v int32)`

SetInactiveDays sets InactiveDays field to given value.

### HasInactiveDays

`func (o *SuppressDisengagedRequest) HasInactiveDays() bool`

HasInactiveDays returns a boolean if a field has been set.

### GetMinSends

`func (o *SuppressDisengagedRequest) GetMinSends() int32`

GetMinSends returns the MinSends field if non-nil, zero value otherwise.

### GetMinSendsOk

`func (o *SuppressDisengagedRequest) GetMinSendsOk() (*int32, bool)`

GetMinSendsOk returns a tuple with the MinSends field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinSends

`func (o *SuppressDisengagedRequest) SetMinSends(v int32)`

SetMinSends sets MinSends field to given value.

### HasMinSends

`func (o *SuppressDisengagedRequest) HasMinSends() bool`

HasMinSends returns a boolean if a field has been set.

### GetDryRun

`func (o *SuppressDisengagedRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *SuppressDisengagedRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *SuppressDisengagedRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *SuppressDisengagedRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


