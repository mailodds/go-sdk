# GetDmarcRecommendation200ResponseRecommendation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentPolicy** | Pointer to **string** | Current DMARC policy (none, quarantine, reject) | [optional] 
**RecommendedPolicy** | Pointer to **string** | Recommended DMARC policy | [optional] 
**Confidence** | Pointer to **float32** | Confidence level (0-1) | [optional] 
**Reasons** | Pointer to **[]string** | Reasons for the recommendation | [optional] 
**ReadyToUpgrade** | Pointer to **bool** | Whether it is safe to upgrade | [optional] 

## Methods

### NewGetDmarcRecommendation200ResponseRecommendation

`func NewGetDmarcRecommendation200ResponseRecommendation() *GetDmarcRecommendation200ResponseRecommendation`

NewGetDmarcRecommendation200ResponseRecommendation instantiates a new GetDmarcRecommendation200ResponseRecommendation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDmarcRecommendation200ResponseRecommendationWithDefaults

`func NewGetDmarcRecommendation200ResponseRecommendationWithDefaults() *GetDmarcRecommendation200ResponseRecommendation`

NewGetDmarcRecommendation200ResponseRecommendationWithDefaults instantiates a new GetDmarcRecommendation200ResponseRecommendation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentPolicy

`func (o *GetDmarcRecommendation200ResponseRecommendation) GetCurrentPolicy() string`

GetCurrentPolicy returns the CurrentPolicy field if non-nil, zero value otherwise.

### GetCurrentPolicyOk

`func (o *GetDmarcRecommendation200ResponseRecommendation) GetCurrentPolicyOk() (*string, bool)`

GetCurrentPolicyOk returns a tuple with the CurrentPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPolicy

`func (o *GetDmarcRecommendation200ResponseRecommendation) SetCurrentPolicy(v string)`

SetCurrentPolicy sets CurrentPolicy field to given value.

### HasCurrentPolicy

`func (o *GetDmarcRecommendation200ResponseRecommendation) HasCurrentPolicy() bool`

HasCurrentPolicy returns a boolean if a field has been set.

### GetRecommendedPolicy

`func (o *GetDmarcRecommendation200ResponseRecommendation) GetRecommendedPolicy() string`

GetRecommendedPolicy returns the RecommendedPolicy field if non-nil, zero value otherwise.

### GetRecommendedPolicyOk

`func (o *GetDmarcRecommendation200ResponseRecommendation) GetRecommendedPolicyOk() (*string, bool)`

GetRecommendedPolicyOk returns a tuple with the RecommendedPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommendedPolicy

`func (o *GetDmarcRecommendation200ResponseRecommendation) SetRecommendedPolicy(v string)`

SetRecommendedPolicy sets RecommendedPolicy field to given value.

### HasRecommendedPolicy

`func (o *GetDmarcRecommendation200ResponseRecommendation) HasRecommendedPolicy() bool`

HasRecommendedPolicy returns a boolean if a field has been set.

### GetConfidence

`func (o *GetDmarcRecommendation200ResponseRecommendation) GetConfidence() float32`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *GetDmarcRecommendation200ResponseRecommendation) GetConfidenceOk() (*float32, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *GetDmarcRecommendation200ResponseRecommendation) SetConfidence(v float32)`

SetConfidence sets Confidence field to given value.

### HasConfidence

`func (o *GetDmarcRecommendation200ResponseRecommendation) HasConfidence() bool`

HasConfidence returns a boolean if a field has been set.

### GetReasons

`func (o *GetDmarcRecommendation200ResponseRecommendation) GetReasons() []string`

GetReasons returns the Reasons field if non-nil, zero value otherwise.

### GetReasonsOk

`func (o *GetDmarcRecommendation200ResponseRecommendation) GetReasonsOk() (*[]string, bool)`

GetReasonsOk returns a tuple with the Reasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasons

`func (o *GetDmarcRecommendation200ResponseRecommendation) SetReasons(v []string)`

SetReasons sets Reasons field to given value.

### HasReasons

`func (o *GetDmarcRecommendation200ResponseRecommendation) HasReasons() bool`

HasReasons returns a boolean if a field has been set.

### GetReadyToUpgrade

`func (o *GetDmarcRecommendation200ResponseRecommendation) GetReadyToUpgrade() bool`

GetReadyToUpgrade returns the ReadyToUpgrade field if non-nil, zero value otherwise.

### GetReadyToUpgradeOk

`func (o *GetDmarcRecommendation200ResponseRecommendation) GetReadyToUpgradeOk() (*bool, bool)`

GetReadyToUpgradeOk returns a tuple with the ReadyToUpgrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadyToUpgrade

`func (o *GetDmarcRecommendation200ResponseRecommendation) SetReadyToUpgrade(v bool)`

SetReadyToUpgrade sets ReadyToUpgrade field to given value.

### HasReadyToUpgrade

`func (o *GetDmarcRecommendation200ResponseRecommendation) HasReadyToUpgrade() bool`

HasReadyToUpgrade returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


