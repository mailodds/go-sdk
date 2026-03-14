# GetCampaignABResults200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SchemaVersion** | Pointer to **string** |  | [optional] 
**RequestId** | Pointer to **string** | Unique request identifier | [optional] 
**CampaignId** | Pointer to **string** |  | [optional] 
**Variants** | Pointer to [**[]GetCampaignABResults200ResponseVariantsInner**](GetCampaignABResults200ResponseVariantsInner.md) |  | [optional] 
**Winner** | Pointer to [**NullableGetCampaignABResults200ResponseWinner**](GetCampaignABResults200ResponseWinner.md) |  | [optional] 

## Methods

### NewGetCampaignABResults200Response

`func NewGetCampaignABResults200Response() *GetCampaignABResults200Response`

NewGetCampaignABResults200Response instantiates a new GetCampaignABResults200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCampaignABResults200ResponseWithDefaults

`func NewGetCampaignABResults200ResponseWithDefaults() *GetCampaignABResults200Response`

NewGetCampaignABResults200ResponseWithDefaults instantiates a new GetCampaignABResults200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemaVersion

`func (o *GetCampaignABResults200Response) GetSchemaVersion() string`

GetSchemaVersion returns the SchemaVersion field if non-nil, zero value otherwise.

### GetSchemaVersionOk

`func (o *GetCampaignABResults200Response) GetSchemaVersionOk() (*string, bool)`

GetSchemaVersionOk returns a tuple with the SchemaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaVersion

`func (o *GetCampaignABResults200Response) SetSchemaVersion(v string)`

SetSchemaVersion sets SchemaVersion field to given value.

### HasSchemaVersion

`func (o *GetCampaignABResults200Response) HasSchemaVersion() bool`

HasSchemaVersion returns a boolean if a field has been set.

### GetRequestId

`func (o *GetCampaignABResults200Response) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *GetCampaignABResults200Response) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *GetCampaignABResults200Response) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *GetCampaignABResults200Response) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetCampaignId

`func (o *GetCampaignABResults200Response) GetCampaignId() string`

GetCampaignId returns the CampaignId field if non-nil, zero value otherwise.

### GetCampaignIdOk

`func (o *GetCampaignABResults200Response) GetCampaignIdOk() (*string, bool)`

GetCampaignIdOk returns a tuple with the CampaignId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignId

`func (o *GetCampaignABResults200Response) SetCampaignId(v string)`

SetCampaignId sets CampaignId field to given value.

### HasCampaignId

`func (o *GetCampaignABResults200Response) HasCampaignId() bool`

HasCampaignId returns a boolean if a field has been set.

### GetVariants

`func (o *GetCampaignABResults200Response) GetVariants() []GetCampaignABResults200ResponseVariantsInner`

GetVariants returns the Variants field if non-nil, zero value otherwise.

### GetVariantsOk

`func (o *GetCampaignABResults200Response) GetVariantsOk() (*[]GetCampaignABResults200ResponseVariantsInner, bool)`

GetVariantsOk returns a tuple with the Variants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariants

`func (o *GetCampaignABResults200Response) SetVariants(v []GetCampaignABResults200ResponseVariantsInner)`

SetVariants sets Variants field to given value.

### HasVariants

`func (o *GetCampaignABResults200Response) HasVariants() bool`

HasVariants returns a boolean if a field has been set.

### GetWinner

`func (o *GetCampaignABResults200Response) GetWinner() GetCampaignABResults200ResponseWinner`

GetWinner returns the Winner field if non-nil, zero value otherwise.

### GetWinnerOk

`func (o *GetCampaignABResults200Response) GetWinnerOk() (*GetCampaignABResults200ResponseWinner, bool)`

GetWinnerOk returns a tuple with the Winner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWinner

`func (o *GetCampaignABResults200Response) SetWinner(v GetCampaignABResults200ResponseWinner)`

SetWinner sets Winner field to given value.

### HasWinner

`func (o *GetCampaignABResults200Response) HasWinner() bool`

HasWinner returns a boolean if a field has been set.

### SetWinnerNil

`func (o *GetCampaignABResults200Response) SetWinnerNil(b bool)`

 SetWinnerNil sets the value for Winner to be an explicit nil

### UnsetWinner
`func (o *GetCampaignABResults200Response) UnsetWinner()`

UnsetWinner ensures that no value is present for Winner, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


