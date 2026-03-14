# CampaignResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SchemaVersion** | Pointer to **string** |  | [optional] 
**RequestId** | Pointer to **string** | Unique request identifier | [optional] 
**Campaign** | Pointer to [**Campaign**](Campaign.md) |  | [optional] 
**Variants** | Pointer to [**[]CampaignVariant**](CampaignVariant.md) | Campaign variants. Present when fetching a single campaign. | [optional] 

## Methods

### NewCampaignResponse

`func NewCampaignResponse() *CampaignResponse`

NewCampaignResponse instantiates a new CampaignResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignResponseWithDefaults

`func NewCampaignResponseWithDefaults() *CampaignResponse`

NewCampaignResponseWithDefaults instantiates a new CampaignResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemaVersion

`func (o *CampaignResponse) GetSchemaVersion() string`

GetSchemaVersion returns the SchemaVersion field if non-nil, zero value otherwise.

### GetSchemaVersionOk

`func (o *CampaignResponse) GetSchemaVersionOk() (*string, bool)`

GetSchemaVersionOk returns a tuple with the SchemaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaVersion

`func (o *CampaignResponse) SetSchemaVersion(v string)`

SetSchemaVersion sets SchemaVersion field to given value.

### HasSchemaVersion

`func (o *CampaignResponse) HasSchemaVersion() bool`

HasSchemaVersion returns a boolean if a field has been set.

### GetRequestId

`func (o *CampaignResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *CampaignResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *CampaignResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *CampaignResponse) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetCampaign

`func (o *CampaignResponse) GetCampaign() Campaign`

GetCampaign returns the Campaign field if non-nil, zero value otherwise.

### GetCampaignOk

`func (o *CampaignResponse) GetCampaignOk() (*Campaign, bool)`

GetCampaignOk returns a tuple with the Campaign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaign

`func (o *CampaignResponse) SetCampaign(v Campaign)`

SetCampaign sets Campaign field to given value.

### HasCampaign

`func (o *CampaignResponse) HasCampaign() bool`

HasCampaign returns a boolean if a field has been set.

### GetVariants

`func (o *CampaignResponse) GetVariants() []CampaignVariant`

GetVariants returns the Variants field if non-nil, zero value otherwise.

### GetVariantsOk

`func (o *CampaignResponse) GetVariantsOk() (*[]CampaignVariant, bool)`

GetVariantsOk returns a tuple with the Variants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariants

`func (o *CampaignResponse) SetVariants(v []CampaignVariant)`

SetVariants sets Variants field to given value.

### HasVariants

`func (o *CampaignResponse) HasVariants() bool`

HasVariants returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


