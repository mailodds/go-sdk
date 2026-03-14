# GetCampaignABResults200ResponseWinner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VariantId** | Pointer to **string** |  | [optional] 
**Metric** | Pointer to **string** | Metric used to determine winner (open_rate or click_rate) | [optional] 
**Confidence** | Pointer to **float32** | Statistical confidence level (0-1) | [optional] 

## Methods

### NewGetCampaignABResults200ResponseWinner

`func NewGetCampaignABResults200ResponseWinner() *GetCampaignABResults200ResponseWinner`

NewGetCampaignABResults200ResponseWinner instantiates a new GetCampaignABResults200ResponseWinner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCampaignABResults200ResponseWinnerWithDefaults

`func NewGetCampaignABResults200ResponseWinnerWithDefaults() *GetCampaignABResults200ResponseWinner`

NewGetCampaignABResults200ResponseWinnerWithDefaults instantiates a new GetCampaignABResults200ResponseWinner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVariantId

`func (o *GetCampaignABResults200ResponseWinner) GetVariantId() string`

GetVariantId returns the VariantId field if non-nil, zero value otherwise.

### GetVariantIdOk

`func (o *GetCampaignABResults200ResponseWinner) GetVariantIdOk() (*string, bool)`

GetVariantIdOk returns a tuple with the VariantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariantId

`func (o *GetCampaignABResults200ResponseWinner) SetVariantId(v string)`

SetVariantId sets VariantId field to given value.

### HasVariantId

`func (o *GetCampaignABResults200ResponseWinner) HasVariantId() bool`

HasVariantId returns a boolean if a field has been set.

### GetMetric

`func (o *GetCampaignABResults200ResponseWinner) GetMetric() string`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *GetCampaignABResults200ResponseWinner) GetMetricOk() (*string, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *GetCampaignABResults200ResponseWinner) SetMetric(v string)`

SetMetric sets Metric field to given value.

### HasMetric

`func (o *GetCampaignABResults200ResponseWinner) HasMetric() bool`

HasMetric returns a boolean if a field has been set.

### GetConfidence

`func (o *GetCampaignABResults200ResponseWinner) GetConfidence() float32`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *GetCampaignABResults200ResponseWinner) GetConfidenceOk() (*float32, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *GetCampaignABResults200ResponseWinner) SetConfidence(v float32)`

SetConfidence sets Confidence field to given value.

### HasConfidence

`func (o *GetCampaignABResults200ResponseWinner) HasConfidence() bool`

HasConfidence returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


