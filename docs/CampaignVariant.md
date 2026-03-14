# CampaignVariant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Variant UUID | 
**CampaignId** | **string** |  | 
**Name** | **string** | Variant name (e.g., \&quot;Variant A\&quot;) | 
**Subject** | **string** |  | 
**Html** | Pointer to **string** | HTML email body | [optional] 
**Text** | Pointer to **string** | Plain text email body | [optional] 
**Weight** | **int32** | Traffic weight percentage (all variant weights must sum to 100) | 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewCampaignVariant

`func NewCampaignVariant(id string, campaignId string, name string, subject string, weight int32, ) *CampaignVariant`

NewCampaignVariant instantiates a new CampaignVariant object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignVariantWithDefaults

`func NewCampaignVariantWithDefaults() *CampaignVariant`

NewCampaignVariantWithDefaults instantiates a new CampaignVariant object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CampaignVariant) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CampaignVariant) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CampaignVariant) SetId(v string)`

SetId sets Id field to given value.


### GetCampaignId

`func (o *CampaignVariant) GetCampaignId() string`

GetCampaignId returns the CampaignId field if non-nil, zero value otherwise.

### GetCampaignIdOk

`func (o *CampaignVariant) GetCampaignIdOk() (*string, bool)`

GetCampaignIdOk returns a tuple with the CampaignId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignId

`func (o *CampaignVariant) SetCampaignId(v string)`

SetCampaignId sets CampaignId field to given value.


### GetName

`func (o *CampaignVariant) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CampaignVariant) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CampaignVariant) SetName(v string)`

SetName sets Name field to given value.


### GetSubject

`func (o *CampaignVariant) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *CampaignVariant) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *CampaignVariant) SetSubject(v string)`

SetSubject sets Subject field to given value.


### GetHtml

`func (o *CampaignVariant) GetHtml() string`

GetHtml returns the Html field if non-nil, zero value otherwise.

### GetHtmlOk

`func (o *CampaignVariant) GetHtmlOk() (*string, bool)`

GetHtmlOk returns a tuple with the Html field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtml

`func (o *CampaignVariant) SetHtml(v string)`

SetHtml sets Html field to given value.

### HasHtml

`func (o *CampaignVariant) HasHtml() bool`

HasHtml returns a boolean if a field has been set.

### GetText

`func (o *CampaignVariant) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *CampaignVariant) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *CampaignVariant) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *CampaignVariant) HasText() bool`

HasText returns a boolean if a field has been set.

### GetWeight

`func (o *CampaignVariant) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *CampaignVariant) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *CampaignVariant) SetWeight(v int32)`

SetWeight sets Weight field to given value.


### GetCreatedAt

`func (o *CampaignVariant) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CampaignVariant) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CampaignVariant) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CampaignVariant) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


