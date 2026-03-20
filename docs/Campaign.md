# Campaign

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Campaign UUID | 
**AccountId** | Pointer to **int32** |  | [optional] 
**Name** | **string** | Campaign name | 
**Status** | **string** |  | 
**DomainId** | **string** | Sending domain UUID | 
**Subject** | Pointer to **string** |  | [optional] 
**FromAddress** | **string** | Sender email address | 
**ReplyTo** | Pointer to **NullableString** |  | [optional] 
**HtmlBody** | Pointer to **NullableString** |  | [optional] 
**TextBody** | Pointer to **NullableString** |  | [optional] 
**HtmlBodyDark** | Pointer to **NullableString** |  | [optional] 
**TextBodyDark** | Pointer to **NullableString** |  | [optional] 
**CampaignType** | Pointer to **NullableString** |  | [optional] 
**AutoDetectSchema** | Pointer to **bool** |  | [optional] 
**PromoAnnotations** | Pointer to **map[string]interface{}** |  | [optional] 
**ThrowawayPolicy** | Pointer to **string** |  | [optional] 
**ScheduledAt** | Pointer to **NullableTime** |  | [optional] 
**StartedAt** | Pointer to **NullableTime** |  | [optional] 
**CompletedAt** | Pointer to **NullableTime** |  | [optional] 
**RecipientCount** | Pointer to **int32** |  | [optional] 
**IsAbTest** | Pointer to **bool** |  | [optional] 
**WinningVariantId** | Pointer to **NullableString** |  | [optional] 
**AbTestConfig** | Pointer to **map[string]interface{}** |  | [optional] 
**ErrorMessage** | Pointer to **NullableString** |  | [optional] 
**Stats** | Pointer to [**CampaignStats**](CampaignStats.md) |  | [optional] 
**OpenRate** | Pointer to **float32** |  | [optional] 
**ClickRate** | Pointer to **float32** |  | [optional] 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewCampaign

`func NewCampaign(id string, name string, status string, domainId string, fromAddress string, createdAt time.Time, ) *Campaign`

NewCampaign instantiates a new Campaign object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignWithDefaults

`func NewCampaignWithDefaults() *Campaign`

NewCampaignWithDefaults instantiates a new Campaign object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Campaign) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Campaign) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Campaign) SetId(v string)`

SetId sets Id field to given value.


### GetAccountId

`func (o *Campaign) GetAccountId() int32`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *Campaign) GetAccountIdOk() (*int32, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *Campaign) SetAccountId(v int32)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *Campaign) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetName

`func (o *Campaign) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Campaign) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Campaign) SetName(v string)`

SetName sets Name field to given value.


### GetStatus

`func (o *Campaign) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Campaign) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Campaign) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetDomainId

`func (o *Campaign) GetDomainId() string`

GetDomainId returns the DomainId field if non-nil, zero value otherwise.

### GetDomainIdOk

`func (o *Campaign) GetDomainIdOk() (*string, bool)`

GetDomainIdOk returns a tuple with the DomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainId

`func (o *Campaign) SetDomainId(v string)`

SetDomainId sets DomainId field to given value.


### GetSubject

`func (o *Campaign) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *Campaign) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *Campaign) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *Campaign) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetFromAddress

`func (o *Campaign) GetFromAddress() string`

GetFromAddress returns the FromAddress field if non-nil, zero value otherwise.

### GetFromAddressOk

`func (o *Campaign) GetFromAddressOk() (*string, bool)`

GetFromAddressOk returns a tuple with the FromAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromAddress

`func (o *Campaign) SetFromAddress(v string)`

SetFromAddress sets FromAddress field to given value.


### GetReplyTo

`func (o *Campaign) GetReplyTo() string`

GetReplyTo returns the ReplyTo field if non-nil, zero value otherwise.

### GetReplyToOk

`func (o *Campaign) GetReplyToOk() (*string, bool)`

GetReplyToOk returns a tuple with the ReplyTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyTo

`func (o *Campaign) SetReplyTo(v string)`

SetReplyTo sets ReplyTo field to given value.

### HasReplyTo

`func (o *Campaign) HasReplyTo() bool`

HasReplyTo returns a boolean if a field has been set.

### SetReplyToNil

`func (o *Campaign) SetReplyToNil(b bool)`

 SetReplyToNil sets the value for ReplyTo to be an explicit nil

### UnsetReplyTo
`func (o *Campaign) UnsetReplyTo()`

UnsetReplyTo ensures that no value is present for ReplyTo, not even an explicit nil
### GetHtmlBody

`func (o *Campaign) GetHtmlBody() string`

GetHtmlBody returns the HtmlBody field if non-nil, zero value otherwise.

### GetHtmlBodyOk

`func (o *Campaign) GetHtmlBodyOk() (*string, bool)`

GetHtmlBodyOk returns a tuple with the HtmlBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlBody

`func (o *Campaign) SetHtmlBody(v string)`

SetHtmlBody sets HtmlBody field to given value.

### HasHtmlBody

`func (o *Campaign) HasHtmlBody() bool`

HasHtmlBody returns a boolean if a field has been set.

### SetHtmlBodyNil

`func (o *Campaign) SetHtmlBodyNil(b bool)`

 SetHtmlBodyNil sets the value for HtmlBody to be an explicit nil

### UnsetHtmlBody
`func (o *Campaign) UnsetHtmlBody()`

UnsetHtmlBody ensures that no value is present for HtmlBody, not even an explicit nil
### GetTextBody

`func (o *Campaign) GetTextBody() string`

GetTextBody returns the TextBody field if non-nil, zero value otherwise.

### GetTextBodyOk

`func (o *Campaign) GetTextBodyOk() (*string, bool)`

GetTextBodyOk returns a tuple with the TextBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextBody

`func (o *Campaign) SetTextBody(v string)`

SetTextBody sets TextBody field to given value.

### HasTextBody

`func (o *Campaign) HasTextBody() bool`

HasTextBody returns a boolean if a field has been set.

### SetTextBodyNil

`func (o *Campaign) SetTextBodyNil(b bool)`

 SetTextBodyNil sets the value for TextBody to be an explicit nil

### UnsetTextBody
`func (o *Campaign) UnsetTextBody()`

UnsetTextBody ensures that no value is present for TextBody, not even an explicit nil
### GetHtmlBodyDark

`func (o *Campaign) GetHtmlBodyDark() string`

GetHtmlBodyDark returns the HtmlBodyDark field if non-nil, zero value otherwise.

### GetHtmlBodyDarkOk

`func (o *Campaign) GetHtmlBodyDarkOk() (*string, bool)`

GetHtmlBodyDarkOk returns a tuple with the HtmlBodyDark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlBodyDark

`func (o *Campaign) SetHtmlBodyDark(v string)`

SetHtmlBodyDark sets HtmlBodyDark field to given value.

### HasHtmlBodyDark

`func (o *Campaign) HasHtmlBodyDark() bool`

HasHtmlBodyDark returns a boolean if a field has been set.

### SetHtmlBodyDarkNil

`func (o *Campaign) SetHtmlBodyDarkNil(b bool)`

 SetHtmlBodyDarkNil sets the value for HtmlBodyDark to be an explicit nil

### UnsetHtmlBodyDark
`func (o *Campaign) UnsetHtmlBodyDark()`

UnsetHtmlBodyDark ensures that no value is present for HtmlBodyDark, not even an explicit nil
### GetTextBodyDark

`func (o *Campaign) GetTextBodyDark() string`

GetTextBodyDark returns the TextBodyDark field if non-nil, zero value otherwise.

### GetTextBodyDarkOk

`func (o *Campaign) GetTextBodyDarkOk() (*string, bool)`

GetTextBodyDarkOk returns a tuple with the TextBodyDark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextBodyDark

`func (o *Campaign) SetTextBodyDark(v string)`

SetTextBodyDark sets TextBodyDark field to given value.

### HasTextBodyDark

`func (o *Campaign) HasTextBodyDark() bool`

HasTextBodyDark returns a boolean if a field has been set.

### SetTextBodyDarkNil

`func (o *Campaign) SetTextBodyDarkNil(b bool)`

 SetTextBodyDarkNil sets the value for TextBodyDark to be an explicit nil

### UnsetTextBodyDark
`func (o *Campaign) UnsetTextBodyDark()`

UnsetTextBodyDark ensures that no value is present for TextBodyDark, not even an explicit nil
### GetCampaignType

`func (o *Campaign) GetCampaignType() string`

GetCampaignType returns the CampaignType field if non-nil, zero value otherwise.

### GetCampaignTypeOk

`func (o *Campaign) GetCampaignTypeOk() (*string, bool)`

GetCampaignTypeOk returns a tuple with the CampaignType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignType

`func (o *Campaign) SetCampaignType(v string)`

SetCampaignType sets CampaignType field to given value.

### HasCampaignType

`func (o *Campaign) HasCampaignType() bool`

HasCampaignType returns a boolean if a field has been set.

### SetCampaignTypeNil

`func (o *Campaign) SetCampaignTypeNil(b bool)`

 SetCampaignTypeNil sets the value for CampaignType to be an explicit nil

### UnsetCampaignType
`func (o *Campaign) UnsetCampaignType()`

UnsetCampaignType ensures that no value is present for CampaignType, not even an explicit nil
### GetAutoDetectSchema

`func (o *Campaign) GetAutoDetectSchema() bool`

GetAutoDetectSchema returns the AutoDetectSchema field if non-nil, zero value otherwise.

### GetAutoDetectSchemaOk

`func (o *Campaign) GetAutoDetectSchemaOk() (*bool, bool)`

GetAutoDetectSchemaOk returns a tuple with the AutoDetectSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoDetectSchema

`func (o *Campaign) SetAutoDetectSchema(v bool)`

SetAutoDetectSchema sets AutoDetectSchema field to given value.

### HasAutoDetectSchema

`func (o *Campaign) HasAutoDetectSchema() bool`

HasAutoDetectSchema returns a boolean if a field has been set.

### GetPromoAnnotations

`func (o *Campaign) GetPromoAnnotations() map[string]interface{}`

GetPromoAnnotations returns the PromoAnnotations field if non-nil, zero value otherwise.

### GetPromoAnnotationsOk

`func (o *Campaign) GetPromoAnnotationsOk() (*map[string]interface{}, bool)`

GetPromoAnnotationsOk returns a tuple with the PromoAnnotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromoAnnotations

`func (o *Campaign) SetPromoAnnotations(v map[string]interface{})`

SetPromoAnnotations sets PromoAnnotations field to given value.

### HasPromoAnnotations

`func (o *Campaign) HasPromoAnnotations() bool`

HasPromoAnnotations returns a boolean if a field has been set.

### SetPromoAnnotationsNil

`func (o *Campaign) SetPromoAnnotationsNil(b bool)`

 SetPromoAnnotationsNil sets the value for PromoAnnotations to be an explicit nil

### UnsetPromoAnnotations
`func (o *Campaign) UnsetPromoAnnotations()`

UnsetPromoAnnotations ensures that no value is present for PromoAnnotations, not even an explicit nil
### GetThrowawayPolicy

`func (o *Campaign) GetThrowawayPolicy() string`

GetThrowawayPolicy returns the ThrowawayPolicy field if non-nil, zero value otherwise.

### GetThrowawayPolicyOk

`func (o *Campaign) GetThrowawayPolicyOk() (*string, bool)`

GetThrowawayPolicyOk returns a tuple with the ThrowawayPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThrowawayPolicy

`func (o *Campaign) SetThrowawayPolicy(v string)`

SetThrowawayPolicy sets ThrowawayPolicy field to given value.

### HasThrowawayPolicy

`func (o *Campaign) HasThrowawayPolicy() bool`

HasThrowawayPolicy returns a boolean if a field has been set.

### GetScheduledAt

`func (o *Campaign) GetScheduledAt() time.Time`

GetScheduledAt returns the ScheduledAt field if non-nil, zero value otherwise.

### GetScheduledAtOk

`func (o *Campaign) GetScheduledAtOk() (*time.Time, bool)`

GetScheduledAtOk returns a tuple with the ScheduledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledAt

`func (o *Campaign) SetScheduledAt(v time.Time)`

SetScheduledAt sets ScheduledAt field to given value.

### HasScheduledAt

`func (o *Campaign) HasScheduledAt() bool`

HasScheduledAt returns a boolean if a field has been set.

### SetScheduledAtNil

`func (o *Campaign) SetScheduledAtNil(b bool)`

 SetScheduledAtNil sets the value for ScheduledAt to be an explicit nil

### UnsetScheduledAt
`func (o *Campaign) UnsetScheduledAt()`

UnsetScheduledAt ensures that no value is present for ScheduledAt, not even an explicit nil
### GetStartedAt

`func (o *Campaign) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *Campaign) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *Campaign) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *Campaign) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### SetStartedAtNil

`func (o *Campaign) SetStartedAtNil(b bool)`

 SetStartedAtNil sets the value for StartedAt to be an explicit nil

### UnsetStartedAt
`func (o *Campaign) UnsetStartedAt()`

UnsetStartedAt ensures that no value is present for StartedAt, not even an explicit nil
### GetCompletedAt

`func (o *Campaign) GetCompletedAt() time.Time`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *Campaign) GetCompletedAtOk() (*time.Time, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *Campaign) SetCompletedAt(v time.Time)`

SetCompletedAt sets CompletedAt field to given value.

### HasCompletedAt

`func (o *Campaign) HasCompletedAt() bool`

HasCompletedAt returns a boolean if a field has been set.

### SetCompletedAtNil

`func (o *Campaign) SetCompletedAtNil(b bool)`

 SetCompletedAtNil sets the value for CompletedAt to be an explicit nil

### UnsetCompletedAt
`func (o *Campaign) UnsetCompletedAt()`

UnsetCompletedAt ensures that no value is present for CompletedAt, not even an explicit nil
### GetRecipientCount

`func (o *Campaign) GetRecipientCount() int32`

GetRecipientCount returns the RecipientCount field if non-nil, zero value otherwise.

### GetRecipientCountOk

`func (o *Campaign) GetRecipientCountOk() (*int32, bool)`

GetRecipientCountOk returns a tuple with the RecipientCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientCount

`func (o *Campaign) SetRecipientCount(v int32)`

SetRecipientCount sets RecipientCount field to given value.

### HasRecipientCount

`func (o *Campaign) HasRecipientCount() bool`

HasRecipientCount returns a boolean if a field has been set.

### GetIsAbTest

`func (o *Campaign) GetIsAbTest() bool`

GetIsAbTest returns the IsAbTest field if non-nil, zero value otherwise.

### GetIsAbTestOk

`func (o *Campaign) GetIsAbTestOk() (*bool, bool)`

GetIsAbTestOk returns a tuple with the IsAbTest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAbTest

`func (o *Campaign) SetIsAbTest(v bool)`

SetIsAbTest sets IsAbTest field to given value.

### HasIsAbTest

`func (o *Campaign) HasIsAbTest() bool`

HasIsAbTest returns a boolean if a field has been set.

### GetWinningVariantId

`func (o *Campaign) GetWinningVariantId() string`

GetWinningVariantId returns the WinningVariantId field if non-nil, zero value otherwise.

### GetWinningVariantIdOk

`func (o *Campaign) GetWinningVariantIdOk() (*string, bool)`

GetWinningVariantIdOk returns a tuple with the WinningVariantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWinningVariantId

`func (o *Campaign) SetWinningVariantId(v string)`

SetWinningVariantId sets WinningVariantId field to given value.

### HasWinningVariantId

`func (o *Campaign) HasWinningVariantId() bool`

HasWinningVariantId returns a boolean if a field has been set.

### SetWinningVariantIdNil

`func (o *Campaign) SetWinningVariantIdNil(b bool)`

 SetWinningVariantIdNil sets the value for WinningVariantId to be an explicit nil

### UnsetWinningVariantId
`func (o *Campaign) UnsetWinningVariantId()`

UnsetWinningVariantId ensures that no value is present for WinningVariantId, not even an explicit nil
### GetAbTestConfig

`func (o *Campaign) GetAbTestConfig() map[string]interface{}`

GetAbTestConfig returns the AbTestConfig field if non-nil, zero value otherwise.

### GetAbTestConfigOk

`func (o *Campaign) GetAbTestConfigOk() (*map[string]interface{}, bool)`

GetAbTestConfigOk returns a tuple with the AbTestConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbTestConfig

`func (o *Campaign) SetAbTestConfig(v map[string]interface{})`

SetAbTestConfig sets AbTestConfig field to given value.

### HasAbTestConfig

`func (o *Campaign) HasAbTestConfig() bool`

HasAbTestConfig returns a boolean if a field has been set.

### SetAbTestConfigNil

`func (o *Campaign) SetAbTestConfigNil(b bool)`

 SetAbTestConfigNil sets the value for AbTestConfig to be an explicit nil

### UnsetAbTestConfig
`func (o *Campaign) UnsetAbTestConfig()`

UnsetAbTestConfig ensures that no value is present for AbTestConfig, not even an explicit nil
### GetErrorMessage

`func (o *Campaign) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *Campaign) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *Campaign) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *Campaign) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *Campaign) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *Campaign) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
### GetStats

`func (o *Campaign) GetStats() CampaignStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *Campaign) GetStatsOk() (*CampaignStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *Campaign) SetStats(v CampaignStats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *Campaign) HasStats() bool`

HasStats returns a boolean if a field has been set.

### GetOpenRate

`func (o *Campaign) GetOpenRate() float32`

GetOpenRate returns the OpenRate field if non-nil, zero value otherwise.

### GetOpenRateOk

`func (o *Campaign) GetOpenRateOk() (*float32, bool)`

GetOpenRateOk returns a tuple with the OpenRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenRate

`func (o *Campaign) SetOpenRate(v float32)`

SetOpenRate sets OpenRate field to given value.

### HasOpenRate

`func (o *Campaign) HasOpenRate() bool`

HasOpenRate returns a boolean if a field has been set.

### GetClickRate

`func (o *Campaign) GetClickRate() float32`

GetClickRate returns the ClickRate field if non-nil, zero value otherwise.

### GetClickRateOk

`func (o *Campaign) GetClickRateOk() (*float32, bool)`

GetClickRateOk returns a tuple with the ClickRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClickRate

`func (o *Campaign) SetClickRate(v float32)`

SetClickRate sets ClickRate field to given value.

### HasClickRate

`func (o *Campaign) HasClickRate() bool`

HasClickRate returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Campaign) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Campaign) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Campaign) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *Campaign) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Campaign) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Campaign) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Campaign) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


