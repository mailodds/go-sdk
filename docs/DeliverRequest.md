# DeliverRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**To** | [**[]DeliverRequestToInner**](DeliverRequestToInner.md) | List of recipient email addresses | 
**From** | **string** | Sender email address (must match sending domain) | 
**Subject** | **string** | Email subject line | 
**Html** | Pointer to **string** | HTML email body | [optional] 
**Text** | Pointer to **string** | Plain text email body | [optional] 
**DomainId** | **string** | Sending domain UUID | 
**ReplyTo** | Pointer to **string** | Reply-to address | [optional] 
**Headers** | Pointer to **map[string]interface{}** | Extra email headers | [optional] 
**Tags** | Pointer to **[]string** | Tags for categorization | [optional] 
**CampaignType** | Pointer to **string** | Campaign type for JSON-LD auto-generation | [optional] 
**StructuredData** | Pointer to [**DeliverRequestStructuredData**](DeliverRequestStructuredData.md) |  | [optional] 
**Options** | Pointer to [**DeliverRequestOptions**](DeliverRequestOptions.md) |  | [optional] 

## Methods

### NewDeliverRequest

`func NewDeliverRequest(to []DeliverRequestToInner, from string, subject string, domainId string, ) *DeliverRequest`

NewDeliverRequest instantiates a new DeliverRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeliverRequestWithDefaults

`func NewDeliverRequestWithDefaults() *DeliverRequest`

NewDeliverRequestWithDefaults instantiates a new DeliverRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTo

`func (o *DeliverRequest) GetTo() []DeliverRequestToInner`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *DeliverRequest) GetToOk() (*[]DeliverRequestToInner, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *DeliverRequest) SetTo(v []DeliverRequestToInner)`

SetTo sets To field to given value.


### GetFrom

`func (o *DeliverRequest) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *DeliverRequest) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *DeliverRequest) SetFrom(v string)`

SetFrom sets From field to given value.


### GetSubject

`func (o *DeliverRequest) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *DeliverRequest) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *DeliverRequest) SetSubject(v string)`

SetSubject sets Subject field to given value.


### GetHtml

`func (o *DeliverRequest) GetHtml() string`

GetHtml returns the Html field if non-nil, zero value otherwise.

### GetHtmlOk

`func (o *DeliverRequest) GetHtmlOk() (*string, bool)`

GetHtmlOk returns a tuple with the Html field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtml

`func (o *DeliverRequest) SetHtml(v string)`

SetHtml sets Html field to given value.

### HasHtml

`func (o *DeliverRequest) HasHtml() bool`

HasHtml returns a boolean if a field has been set.

### GetText

`func (o *DeliverRequest) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *DeliverRequest) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *DeliverRequest) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *DeliverRequest) HasText() bool`

HasText returns a boolean if a field has been set.

### GetDomainId

`func (o *DeliverRequest) GetDomainId() string`

GetDomainId returns the DomainId field if non-nil, zero value otherwise.

### GetDomainIdOk

`func (o *DeliverRequest) GetDomainIdOk() (*string, bool)`

GetDomainIdOk returns a tuple with the DomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainId

`func (o *DeliverRequest) SetDomainId(v string)`

SetDomainId sets DomainId field to given value.


### GetReplyTo

`func (o *DeliverRequest) GetReplyTo() string`

GetReplyTo returns the ReplyTo field if non-nil, zero value otherwise.

### GetReplyToOk

`func (o *DeliverRequest) GetReplyToOk() (*string, bool)`

GetReplyToOk returns a tuple with the ReplyTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyTo

`func (o *DeliverRequest) SetReplyTo(v string)`

SetReplyTo sets ReplyTo field to given value.

### HasReplyTo

`func (o *DeliverRequest) HasReplyTo() bool`

HasReplyTo returns a boolean if a field has been set.

### GetHeaders

`func (o *DeliverRequest) GetHeaders() map[string]interface{}`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *DeliverRequest) GetHeadersOk() (*map[string]interface{}, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *DeliverRequest) SetHeaders(v map[string]interface{})`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *DeliverRequest) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetTags

`func (o *DeliverRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DeliverRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DeliverRequest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *DeliverRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCampaignType

`func (o *DeliverRequest) GetCampaignType() string`

GetCampaignType returns the CampaignType field if non-nil, zero value otherwise.

### GetCampaignTypeOk

`func (o *DeliverRequest) GetCampaignTypeOk() (*string, bool)`

GetCampaignTypeOk returns a tuple with the CampaignType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignType

`func (o *DeliverRequest) SetCampaignType(v string)`

SetCampaignType sets CampaignType field to given value.

### HasCampaignType

`func (o *DeliverRequest) HasCampaignType() bool`

HasCampaignType returns a boolean if a field has been set.

### GetStructuredData

`func (o *DeliverRequest) GetStructuredData() DeliverRequestStructuredData`

GetStructuredData returns the StructuredData field if non-nil, zero value otherwise.

### GetStructuredDataOk

`func (o *DeliverRequest) GetStructuredDataOk() (*DeliverRequestStructuredData, bool)`

GetStructuredDataOk returns a tuple with the StructuredData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStructuredData

`func (o *DeliverRequest) SetStructuredData(v DeliverRequestStructuredData)`

SetStructuredData sets StructuredData field to given value.

### HasStructuredData

`func (o *DeliverRequest) HasStructuredData() bool`

HasStructuredData returns a boolean if a field has been set.

### GetOptions

`func (o *DeliverRequest) GetOptions() DeliverRequestOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *DeliverRequest) GetOptionsOk() (*DeliverRequestOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *DeliverRequest) SetOptions(v DeliverRequestOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *DeliverRequest) HasOptions() bool`

HasOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


