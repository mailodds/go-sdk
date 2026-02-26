# SendingDomainIdentityScore

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OverallScore** | Pointer to **float32** | Composite score 0-100 | [optional] 
**Checks** | Pointer to [**SendingDomainIdentityScoreChecks**](SendingDomainIdentityScoreChecks.md) |  | [optional] 

## Methods

### NewSendingDomainIdentityScore

`func NewSendingDomainIdentityScore() *SendingDomainIdentityScore`

NewSendingDomainIdentityScore instantiates a new SendingDomainIdentityScore object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendingDomainIdentityScoreWithDefaults

`func NewSendingDomainIdentityScoreWithDefaults() *SendingDomainIdentityScore`

NewSendingDomainIdentityScoreWithDefaults instantiates a new SendingDomainIdentityScore object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOverallScore

`func (o *SendingDomainIdentityScore) GetOverallScore() float32`

GetOverallScore returns the OverallScore field if non-nil, zero value otherwise.

### GetOverallScoreOk

`func (o *SendingDomainIdentityScore) GetOverallScoreOk() (*float32, bool)`

GetOverallScoreOk returns a tuple with the OverallScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverallScore

`func (o *SendingDomainIdentityScore) SetOverallScore(v float32)`

SetOverallScore sets OverallScore field to given value.

### HasOverallScore

`func (o *SendingDomainIdentityScore) HasOverallScore() bool`

HasOverallScore returns a boolean if a field has been set.

### GetChecks

`func (o *SendingDomainIdentityScore) GetChecks() SendingDomainIdentityScoreChecks`

GetChecks returns the Checks field if non-nil, zero value otherwise.

### GetChecksOk

`func (o *SendingDomainIdentityScore) GetChecksOk() (*SendingDomainIdentityScoreChecks, bool)`

GetChecksOk returns a tuple with the Checks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecks

`func (o *SendingDomainIdentityScore) SetChecks(v SendingDomainIdentityScoreChecks)`

SetChecks sets Checks field to given value.

### HasChecks

`func (o *SendingDomainIdentityScore) HasChecks() bool`

HasChecks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


