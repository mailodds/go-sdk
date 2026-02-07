# CreateJobRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Emails** | **[]string** | List of emails to validate | 
**Dedup** | Pointer to **bool** | Remove duplicate emails | [optional] [default to false]
**Metadata** | Pointer to **map[string]interface{}** | Custom metadata for the job | [optional] 
**WebhookUrl** | Pointer to **string** | URL for completion webhook | [optional] 
**IdempotencyKey** | Pointer to **string** | Unique key for idempotent requests | [optional] 

## Methods

### NewCreateJobRequest

`func NewCreateJobRequest(emails []string, ) *CreateJobRequest`

NewCreateJobRequest instantiates a new CreateJobRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateJobRequestWithDefaults

`func NewCreateJobRequestWithDefaults() *CreateJobRequest`

NewCreateJobRequestWithDefaults instantiates a new CreateJobRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmails

`func (o *CreateJobRequest) GetEmails() []string`

GetEmails returns the Emails field if non-nil, zero value otherwise.

### GetEmailsOk

`func (o *CreateJobRequest) GetEmailsOk() (*[]string, bool)`

GetEmailsOk returns a tuple with the Emails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmails

`func (o *CreateJobRequest) SetEmails(v []string)`

SetEmails sets Emails field to given value.


### GetDedup

`func (o *CreateJobRequest) GetDedup() bool`

GetDedup returns the Dedup field if non-nil, zero value otherwise.

### GetDedupOk

`func (o *CreateJobRequest) GetDedupOk() (*bool, bool)`

GetDedupOk returns a tuple with the Dedup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDedup

`func (o *CreateJobRequest) SetDedup(v bool)`

SetDedup sets Dedup field to given value.

### HasDedup

`func (o *CreateJobRequest) HasDedup() bool`

HasDedup returns a boolean if a field has been set.

### GetMetadata

`func (o *CreateJobRequest) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CreateJobRequest) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CreateJobRequest) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CreateJobRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetWebhookUrl

`func (o *CreateJobRequest) GetWebhookUrl() string`

GetWebhookUrl returns the WebhookUrl field if non-nil, zero value otherwise.

### GetWebhookUrlOk

`func (o *CreateJobRequest) GetWebhookUrlOk() (*string, bool)`

GetWebhookUrlOk returns a tuple with the WebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookUrl

`func (o *CreateJobRequest) SetWebhookUrl(v string)`

SetWebhookUrl sets WebhookUrl field to given value.

### HasWebhookUrl

`func (o *CreateJobRequest) HasWebhookUrl() bool`

HasWebhookUrl returns a boolean if a field has been set.

### GetIdempotencyKey

`func (o *CreateJobRequest) GetIdempotencyKey() string`

GetIdempotencyKey returns the IdempotencyKey field if non-nil, zero value otherwise.

### GetIdempotencyKeyOk

`func (o *CreateJobRequest) GetIdempotencyKeyOk() (*string, bool)`

GetIdempotencyKeyOk returns a tuple with the IdempotencyKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdempotencyKey

`func (o *CreateJobRequest) SetIdempotencyKey(v string)`

SetIdempotencyKey sets IdempotencyKey field to given value.

### HasIdempotencyKey

`func (o *CreateJobRequest) HasIdempotencyKey() bool`

HasIdempotencyKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


