# GetCampaignDeliveryConfidence200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SchemaVersion** | Pointer to **string** |  | [optional] 
**RequestId** | Pointer to **string** | Unique request identifier | [optional] 
**CampaignId** | Pointer to **string** |  | [optional] 
**ConfidenceScore** | Pointer to **int32** | Predicted delivery confidence (0-100) | [optional] 
**Factors** | Pointer to [**GetCampaignDeliveryConfidence200ResponseFactors**](GetCampaignDeliveryConfidence200ResponseFactors.md) |  | [optional] 
**Recommendations** | Pointer to **[]string** | Actionable recommendations to improve delivery confidence | [optional] 

## Methods

### NewGetCampaignDeliveryConfidence200Response

`func NewGetCampaignDeliveryConfidence200Response() *GetCampaignDeliveryConfidence200Response`

NewGetCampaignDeliveryConfidence200Response instantiates a new GetCampaignDeliveryConfidence200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCampaignDeliveryConfidence200ResponseWithDefaults

`func NewGetCampaignDeliveryConfidence200ResponseWithDefaults() *GetCampaignDeliveryConfidence200Response`

NewGetCampaignDeliveryConfidence200ResponseWithDefaults instantiates a new GetCampaignDeliveryConfidence200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemaVersion

`func (o *GetCampaignDeliveryConfidence200Response) GetSchemaVersion() string`

GetSchemaVersion returns the SchemaVersion field if non-nil, zero value otherwise.

### GetSchemaVersionOk

`func (o *GetCampaignDeliveryConfidence200Response) GetSchemaVersionOk() (*string, bool)`

GetSchemaVersionOk returns a tuple with the SchemaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaVersion

`func (o *GetCampaignDeliveryConfidence200Response) SetSchemaVersion(v string)`

SetSchemaVersion sets SchemaVersion field to given value.

### HasSchemaVersion

`func (o *GetCampaignDeliveryConfidence200Response) HasSchemaVersion() bool`

HasSchemaVersion returns a boolean if a field has been set.

### GetRequestId

`func (o *GetCampaignDeliveryConfidence200Response) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *GetCampaignDeliveryConfidence200Response) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *GetCampaignDeliveryConfidence200Response) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *GetCampaignDeliveryConfidence200Response) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetCampaignId

`func (o *GetCampaignDeliveryConfidence200Response) GetCampaignId() string`

GetCampaignId returns the CampaignId field if non-nil, zero value otherwise.

### GetCampaignIdOk

`func (o *GetCampaignDeliveryConfidence200Response) GetCampaignIdOk() (*string, bool)`

GetCampaignIdOk returns a tuple with the CampaignId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignId

`func (o *GetCampaignDeliveryConfidence200Response) SetCampaignId(v string)`

SetCampaignId sets CampaignId field to given value.

### HasCampaignId

`func (o *GetCampaignDeliveryConfidence200Response) HasCampaignId() bool`

HasCampaignId returns a boolean if a field has been set.

### GetConfidenceScore

`func (o *GetCampaignDeliveryConfidence200Response) GetConfidenceScore() int32`

GetConfidenceScore returns the ConfidenceScore field if non-nil, zero value otherwise.

### GetConfidenceScoreOk

`func (o *GetCampaignDeliveryConfidence200Response) GetConfidenceScoreOk() (*int32, bool)`

GetConfidenceScoreOk returns a tuple with the ConfidenceScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidenceScore

`func (o *GetCampaignDeliveryConfidence200Response) SetConfidenceScore(v int32)`

SetConfidenceScore sets ConfidenceScore field to given value.

### HasConfidenceScore

`func (o *GetCampaignDeliveryConfidence200Response) HasConfidenceScore() bool`

HasConfidenceScore returns a boolean if a field has been set.

### GetFactors

`func (o *GetCampaignDeliveryConfidence200Response) GetFactors() GetCampaignDeliveryConfidence200ResponseFactors`

GetFactors returns the Factors field if non-nil, zero value otherwise.

### GetFactorsOk

`func (o *GetCampaignDeliveryConfidence200Response) GetFactorsOk() (*GetCampaignDeliveryConfidence200ResponseFactors, bool)`

GetFactorsOk returns a tuple with the Factors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFactors

`func (o *GetCampaignDeliveryConfidence200Response) SetFactors(v GetCampaignDeliveryConfidence200ResponseFactors)`

SetFactors sets Factors field to given value.

### HasFactors

`func (o *GetCampaignDeliveryConfidence200Response) HasFactors() bool`

HasFactors returns a boolean if a field has been set.

### GetRecommendations

`func (o *GetCampaignDeliveryConfidence200Response) GetRecommendations() []string`

GetRecommendations returns the Recommendations field if non-nil, zero value otherwise.

### GetRecommendationsOk

`func (o *GetCampaignDeliveryConfidence200Response) GetRecommendationsOk() (*[]string, bool)`

GetRecommendationsOk returns a tuple with the Recommendations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommendations

`func (o *GetCampaignDeliveryConfidence200Response) SetRecommendations(v []string)`

SetRecommendations sets Recommendations field to given value.

### HasRecommendations

`func (o *GetCampaignDeliveryConfidence200Response) HasRecommendations() bool`

HasRecommendations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


