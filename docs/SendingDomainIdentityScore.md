# SendingDomainIdentityScore

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Score** | **int32** | Total points earned across all checks | 
**MaxScore** | **int32** | Maximum possible score (100) | 
**Percentage** | **int32** | Score as percentage (same as score since max is 100) | 
**Breakdown** | [**SendingDomainIdentityScoreBreakdown**](SendingDomainIdentityScoreBreakdown.md) |  | 
**Grade** | **string** | Letter grade (A+, A, B, C, D, F) | 

## Methods

### NewSendingDomainIdentityScore

`func NewSendingDomainIdentityScore(score int32, maxScore int32, percentage int32, breakdown SendingDomainIdentityScoreBreakdown, grade string, ) *SendingDomainIdentityScore`

NewSendingDomainIdentityScore instantiates a new SendingDomainIdentityScore object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendingDomainIdentityScoreWithDefaults

`func NewSendingDomainIdentityScoreWithDefaults() *SendingDomainIdentityScore`

NewSendingDomainIdentityScoreWithDefaults instantiates a new SendingDomainIdentityScore object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScore

`func (o *SendingDomainIdentityScore) GetScore() int32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *SendingDomainIdentityScore) GetScoreOk() (*int32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *SendingDomainIdentityScore) SetScore(v int32)`

SetScore sets Score field to given value.


### GetMaxScore

`func (o *SendingDomainIdentityScore) GetMaxScore() int32`

GetMaxScore returns the MaxScore field if non-nil, zero value otherwise.

### GetMaxScoreOk

`func (o *SendingDomainIdentityScore) GetMaxScoreOk() (*int32, bool)`

GetMaxScoreOk returns a tuple with the MaxScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxScore

`func (o *SendingDomainIdentityScore) SetMaxScore(v int32)`

SetMaxScore sets MaxScore field to given value.


### GetPercentage

`func (o *SendingDomainIdentityScore) GetPercentage() int32`

GetPercentage returns the Percentage field if non-nil, zero value otherwise.

### GetPercentageOk

`func (o *SendingDomainIdentityScore) GetPercentageOk() (*int32, bool)`

GetPercentageOk returns a tuple with the Percentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentage

`func (o *SendingDomainIdentityScore) SetPercentage(v int32)`

SetPercentage sets Percentage field to given value.


### GetBreakdown

`func (o *SendingDomainIdentityScore) GetBreakdown() SendingDomainIdentityScoreBreakdown`

GetBreakdown returns the Breakdown field if non-nil, zero value otherwise.

### GetBreakdownOk

`func (o *SendingDomainIdentityScore) GetBreakdownOk() (*SendingDomainIdentityScoreBreakdown, bool)`

GetBreakdownOk returns a tuple with the Breakdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreakdown

`func (o *SendingDomainIdentityScore) SetBreakdown(v SendingDomainIdentityScoreBreakdown)`

SetBreakdown sets Breakdown field to given value.


### GetGrade

`func (o *SendingDomainIdentityScore) GetGrade() string`

GetGrade returns the Grade field if non-nil, zero value otherwise.

### GetGradeOk

`func (o *SendingDomainIdentityScore) GetGradeOk() (*string, bool)`

GetGradeOk returns a tuple with the Grade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrade

`func (o *SendingDomainIdentityScore) SetGrade(v string)`

SetGrade sets Grade field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


