# BatchDeliverRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**To** | **[]string** | List of recipient email addresses (max 100) | 
**From** | **string** |  | 
**Subject** | **string** |  | 
**Html** | Pointer to **string** |  | [optional] 
**Text** | Pointer to **string** |  | [optional] 
**DomainId** | **string** |  | 
**ReplyTo** | Pointer to **string** |  | [optional] 
**Headers** | Pointer to **map[string]interface{}** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**CampaignType** | Pointer to **string** |  | [optional] 
**StructuredData** | Pointer to [**BatchDeliverRequestStructuredData**](BatchDeliverRequestStructuredData.md) |  | [optional] 
**Options** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewBatchDeliverRequest

`func NewBatchDeliverRequest(to []string, from string, subject string, domainId string, ) *BatchDeliverRequest`

NewBatchDeliverRequest instantiates a new BatchDeliverRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchDeliverRequestWithDefaults

`func NewBatchDeliverRequestWithDefaults() *BatchDeliverRequest`

NewBatchDeliverRequestWithDefaults instantiates a new BatchDeliverRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTo

`func (o *BatchDeliverRequest) GetTo() []string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *BatchDeliverRequest) GetToOk() (*[]string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *BatchDeliverRequest) SetTo(v []string)`

SetTo sets To field to given value.


### GetFrom

`func (o *BatchDeliverRequest) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *BatchDeliverRequest) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *BatchDeliverRequest) SetFrom(v string)`

SetFrom sets From field to given value.


### GetSubject

`func (o *BatchDeliverRequest) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *BatchDeliverRequest) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *BatchDeliverRequest) SetSubject(v string)`

SetSubject sets Subject field to given value.


### GetHtml

`func (o *BatchDeliverRequest) GetHtml() string`

GetHtml returns the Html field if non-nil, zero value otherwise.

### GetHtmlOk

`func (o *BatchDeliverRequest) GetHtmlOk() (*string, bool)`

GetHtmlOk returns a tuple with the Html field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtml

`func (o *BatchDeliverRequest) SetHtml(v string)`

SetHtml sets Html field to given value.

### HasHtml

`func (o *BatchDeliverRequest) HasHtml() bool`

HasHtml returns a boolean if a field has been set.

### GetText

`func (o *BatchDeliverRequest) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *BatchDeliverRequest) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *BatchDeliverRequest) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *BatchDeliverRequest) HasText() bool`

HasText returns a boolean if a field has been set.

### GetDomainId

`func (o *BatchDeliverRequest) GetDomainId() string`

GetDomainId returns the DomainId field if non-nil, zero value otherwise.

### GetDomainIdOk

`func (o *BatchDeliverRequest) GetDomainIdOk() (*string, bool)`

GetDomainIdOk returns a tuple with the DomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainId

`func (o *BatchDeliverRequest) SetDomainId(v string)`

SetDomainId sets DomainId field to given value.


### GetReplyTo

`func (o *BatchDeliverRequest) GetReplyTo() string`

GetReplyTo returns the ReplyTo field if non-nil, zero value otherwise.

### GetReplyToOk

`func (o *BatchDeliverRequest) GetReplyToOk() (*string, bool)`

GetReplyToOk returns a tuple with the ReplyTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyTo

`func (o *BatchDeliverRequest) SetReplyTo(v string)`

SetReplyTo sets ReplyTo field to given value.

### HasReplyTo

`func (o *BatchDeliverRequest) HasReplyTo() bool`

HasReplyTo returns a boolean if a field has been set.

### GetHeaders

`func (o *BatchDeliverRequest) GetHeaders() map[string]interface{}`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *BatchDeliverRequest) GetHeadersOk() (*map[string]interface{}, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *BatchDeliverRequest) SetHeaders(v map[string]interface{})`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *BatchDeliverRequest) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetTags

`func (o *BatchDeliverRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *BatchDeliverRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *BatchDeliverRequest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *BatchDeliverRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCampaignType

`func (o *BatchDeliverRequest) GetCampaignType() string`

GetCampaignType returns the CampaignType field if non-nil, zero value otherwise.

### GetCampaignTypeOk

`func (o *BatchDeliverRequest) GetCampaignTypeOk() (*string, bool)`

GetCampaignTypeOk returns a tuple with the CampaignType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignType

`func (o *BatchDeliverRequest) SetCampaignType(v string)`

SetCampaignType sets CampaignType field to given value.

### HasCampaignType

`func (o *BatchDeliverRequest) HasCampaignType() bool`

HasCampaignType returns a boolean if a field has been set.

### GetStructuredData

`func (o *BatchDeliverRequest) GetStructuredData() BatchDeliverRequestStructuredData`

GetStructuredData returns the StructuredData field if non-nil, zero value otherwise.

### GetStructuredDataOk

`func (o *BatchDeliverRequest) GetStructuredDataOk() (*BatchDeliverRequestStructuredData, bool)`

GetStructuredDataOk returns a tuple with the StructuredData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStructuredData

`func (o *BatchDeliverRequest) SetStructuredData(v BatchDeliverRequestStructuredData)`

SetStructuredData sets StructuredData field to given value.

### HasStructuredData

`func (o *BatchDeliverRequest) HasStructuredData() bool`

HasStructuredData returns a boolean if a field has been set.

### GetOptions

`func (o *BatchDeliverRequest) GetOptions() map[string]interface{}`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *BatchDeliverRequest) GetOptionsOk() (*map[string]interface{}, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *BatchDeliverRequest) SetOptions(v map[string]interface{})`

SetOptions sets Options field to given value.

### HasOptions

`func (o *BatchDeliverRequest) HasOptions() bool`

HasOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


