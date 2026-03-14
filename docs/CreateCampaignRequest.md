# CreateCampaignRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Campaign name | 
**ListId** | **string** | Target subscriber list UUID | 
**DomainId** | **string** | Sending domain UUID | 
**FromEmail** | **string** | Sender email address (must match the sending domain) | 
**FromName** | Pointer to **string** | Sender display name | [optional] 
**ReplyTo** | Pointer to **string** | Reply-to address | [optional] 
**Tags** | Pointer to **[]string** | Tags for categorization | [optional] 

## Methods

### NewCreateCampaignRequest

`func NewCreateCampaignRequest(name string, listId string, domainId string, fromEmail string, ) *CreateCampaignRequest`

NewCreateCampaignRequest instantiates a new CreateCampaignRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCampaignRequestWithDefaults

`func NewCreateCampaignRequestWithDefaults() *CreateCampaignRequest`

NewCreateCampaignRequestWithDefaults instantiates a new CreateCampaignRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateCampaignRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateCampaignRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateCampaignRequest) SetName(v string)`

SetName sets Name field to given value.


### GetListId

`func (o *CreateCampaignRequest) GetListId() string`

GetListId returns the ListId field if non-nil, zero value otherwise.

### GetListIdOk

`func (o *CreateCampaignRequest) GetListIdOk() (*string, bool)`

GetListIdOk returns a tuple with the ListId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListId

`func (o *CreateCampaignRequest) SetListId(v string)`

SetListId sets ListId field to given value.


### GetDomainId

`func (o *CreateCampaignRequest) GetDomainId() string`

GetDomainId returns the DomainId field if non-nil, zero value otherwise.

### GetDomainIdOk

`func (o *CreateCampaignRequest) GetDomainIdOk() (*string, bool)`

GetDomainIdOk returns a tuple with the DomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainId

`func (o *CreateCampaignRequest) SetDomainId(v string)`

SetDomainId sets DomainId field to given value.


### GetFromEmail

`func (o *CreateCampaignRequest) GetFromEmail() string`

GetFromEmail returns the FromEmail field if non-nil, zero value otherwise.

### GetFromEmailOk

`func (o *CreateCampaignRequest) GetFromEmailOk() (*string, bool)`

GetFromEmailOk returns a tuple with the FromEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromEmail

`func (o *CreateCampaignRequest) SetFromEmail(v string)`

SetFromEmail sets FromEmail field to given value.


### GetFromName

`func (o *CreateCampaignRequest) GetFromName() string`

GetFromName returns the FromName field if non-nil, zero value otherwise.

### GetFromNameOk

`func (o *CreateCampaignRequest) GetFromNameOk() (*string, bool)`

GetFromNameOk returns a tuple with the FromName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromName

`func (o *CreateCampaignRequest) SetFromName(v string)`

SetFromName sets FromName field to given value.

### HasFromName

`func (o *CreateCampaignRequest) HasFromName() bool`

HasFromName returns a boolean if a field has been set.

### GetReplyTo

`func (o *CreateCampaignRequest) GetReplyTo() string`

GetReplyTo returns the ReplyTo field if non-nil, zero value otherwise.

### GetReplyToOk

`func (o *CreateCampaignRequest) GetReplyToOk() (*string, bool)`

GetReplyToOk returns a tuple with the ReplyTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyTo

`func (o *CreateCampaignRequest) SetReplyTo(v string)`

SetReplyTo sets ReplyTo field to given value.

### HasReplyTo

`func (o *CreateCampaignRequest) HasReplyTo() bool`

HasReplyTo returns a boolean if a field has been set.

### GetTags

`func (o *CreateCampaignRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CreateCampaignRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CreateCampaignRequest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CreateCampaignRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


