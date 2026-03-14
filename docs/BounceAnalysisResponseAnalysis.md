# BounceAnalysisResponseAnalysis

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Analysis UUID | [optional] 
**DomainId** | Pointer to **string** |  | [optional] 
**Period** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**TotalBounces** | Pointer to **int32** |  | [optional] 
**HardBounces** | Pointer to **int32** |  | [optional] 
**SoftBounces** | Pointer to **int32** |  | [optional] 
**Categories** | Pointer to [**BounceAnalysisResponseAnalysisCategories**](BounceAnalysisResponseAnalysisCategories.md) |  | [optional] 
**TopDomains** | Pointer to [**[]BounceAnalysisResponseAnalysisTopDomainsInner**](BounceAnalysisResponseAnalysisTopDomainsInner.md) | Top bouncing recipient domains | [optional] 
**Recommendations** | Pointer to **[]string** | Actionable recommendations to reduce bounces | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewBounceAnalysisResponseAnalysis

`func NewBounceAnalysisResponseAnalysis() *BounceAnalysisResponseAnalysis`

NewBounceAnalysisResponseAnalysis instantiates a new BounceAnalysisResponseAnalysis object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBounceAnalysisResponseAnalysisWithDefaults

`func NewBounceAnalysisResponseAnalysisWithDefaults() *BounceAnalysisResponseAnalysis`

NewBounceAnalysisResponseAnalysisWithDefaults instantiates a new BounceAnalysisResponseAnalysis object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BounceAnalysisResponseAnalysis) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BounceAnalysisResponseAnalysis) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BounceAnalysisResponseAnalysis) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BounceAnalysisResponseAnalysis) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDomainId

`func (o *BounceAnalysisResponseAnalysis) GetDomainId() string`

GetDomainId returns the DomainId field if non-nil, zero value otherwise.

### GetDomainIdOk

`func (o *BounceAnalysisResponseAnalysis) GetDomainIdOk() (*string, bool)`

GetDomainIdOk returns a tuple with the DomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainId

`func (o *BounceAnalysisResponseAnalysis) SetDomainId(v string)`

SetDomainId sets DomainId field to given value.

### HasDomainId

`func (o *BounceAnalysisResponseAnalysis) HasDomainId() bool`

HasDomainId returns a boolean if a field has been set.

### GetPeriod

`func (o *BounceAnalysisResponseAnalysis) GetPeriod() string`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *BounceAnalysisResponseAnalysis) GetPeriodOk() (*string, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *BounceAnalysisResponseAnalysis) SetPeriod(v string)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *BounceAnalysisResponseAnalysis) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### GetStatus

`func (o *BounceAnalysisResponseAnalysis) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BounceAnalysisResponseAnalysis) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BounceAnalysisResponseAnalysis) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BounceAnalysisResponseAnalysis) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTotalBounces

`func (o *BounceAnalysisResponseAnalysis) GetTotalBounces() int32`

GetTotalBounces returns the TotalBounces field if non-nil, zero value otherwise.

### GetTotalBouncesOk

`func (o *BounceAnalysisResponseAnalysis) GetTotalBouncesOk() (*int32, bool)`

GetTotalBouncesOk returns a tuple with the TotalBounces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalBounces

`func (o *BounceAnalysisResponseAnalysis) SetTotalBounces(v int32)`

SetTotalBounces sets TotalBounces field to given value.

### HasTotalBounces

`func (o *BounceAnalysisResponseAnalysis) HasTotalBounces() bool`

HasTotalBounces returns a boolean if a field has been set.

### GetHardBounces

`func (o *BounceAnalysisResponseAnalysis) GetHardBounces() int32`

GetHardBounces returns the HardBounces field if non-nil, zero value otherwise.

### GetHardBouncesOk

`func (o *BounceAnalysisResponseAnalysis) GetHardBouncesOk() (*int32, bool)`

GetHardBouncesOk returns a tuple with the HardBounces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardBounces

`func (o *BounceAnalysisResponseAnalysis) SetHardBounces(v int32)`

SetHardBounces sets HardBounces field to given value.

### HasHardBounces

`func (o *BounceAnalysisResponseAnalysis) HasHardBounces() bool`

HasHardBounces returns a boolean if a field has been set.

### GetSoftBounces

`func (o *BounceAnalysisResponseAnalysis) GetSoftBounces() int32`

GetSoftBounces returns the SoftBounces field if non-nil, zero value otherwise.

### GetSoftBouncesOk

`func (o *BounceAnalysisResponseAnalysis) GetSoftBouncesOk() (*int32, bool)`

GetSoftBouncesOk returns a tuple with the SoftBounces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftBounces

`func (o *BounceAnalysisResponseAnalysis) SetSoftBounces(v int32)`

SetSoftBounces sets SoftBounces field to given value.

### HasSoftBounces

`func (o *BounceAnalysisResponseAnalysis) HasSoftBounces() bool`

HasSoftBounces returns a boolean if a field has been set.

### GetCategories

`func (o *BounceAnalysisResponseAnalysis) GetCategories() BounceAnalysisResponseAnalysisCategories`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *BounceAnalysisResponseAnalysis) GetCategoriesOk() (*BounceAnalysisResponseAnalysisCategories, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *BounceAnalysisResponseAnalysis) SetCategories(v BounceAnalysisResponseAnalysisCategories)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *BounceAnalysisResponseAnalysis) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### GetTopDomains

`func (o *BounceAnalysisResponseAnalysis) GetTopDomains() []BounceAnalysisResponseAnalysisTopDomainsInner`

GetTopDomains returns the TopDomains field if non-nil, zero value otherwise.

### GetTopDomainsOk

`func (o *BounceAnalysisResponseAnalysis) GetTopDomainsOk() (*[]BounceAnalysisResponseAnalysisTopDomainsInner, bool)`

GetTopDomainsOk returns a tuple with the TopDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopDomains

`func (o *BounceAnalysisResponseAnalysis) SetTopDomains(v []BounceAnalysisResponseAnalysisTopDomainsInner)`

SetTopDomains sets TopDomains field to given value.

### HasTopDomains

`func (o *BounceAnalysisResponseAnalysis) HasTopDomains() bool`

HasTopDomains returns a boolean if a field has been set.

### GetRecommendations

`func (o *BounceAnalysisResponseAnalysis) GetRecommendations() []string`

GetRecommendations returns the Recommendations field if non-nil, zero value otherwise.

### GetRecommendationsOk

`func (o *BounceAnalysisResponseAnalysis) GetRecommendationsOk() (*[]string, bool)`

GetRecommendationsOk returns a tuple with the Recommendations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommendations

`func (o *BounceAnalysisResponseAnalysis) SetRecommendations(v []string)`

SetRecommendations sets Recommendations field to given value.

### HasRecommendations

`func (o *BounceAnalysisResponseAnalysis) HasRecommendations() bool`

HasRecommendations returns a boolean if a field has been set.

### GetCreatedAt

`func (o *BounceAnalysisResponseAnalysis) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BounceAnalysisResponseAnalysis) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BounceAnalysisResponseAnalysis) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *BounceAnalysisResponseAnalysis) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


