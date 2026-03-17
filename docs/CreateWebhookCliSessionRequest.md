# CreateWebhookCliSessionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ForwardUrl** | Pointer to **string** | Local URL where webhooks will be forwarded | [optional] 

## Methods

### NewCreateWebhookCliSessionRequest

`func NewCreateWebhookCliSessionRequest() *CreateWebhookCliSessionRequest`

NewCreateWebhookCliSessionRequest instantiates a new CreateWebhookCliSessionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateWebhookCliSessionRequestWithDefaults

`func NewCreateWebhookCliSessionRequestWithDefaults() *CreateWebhookCliSessionRequest`

NewCreateWebhookCliSessionRequestWithDefaults instantiates a new CreateWebhookCliSessionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetForwardUrl

`func (o *CreateWebhookCliSessionRequest) GetForwardUrl() string`

GetForwardUrl returns the ForwardUrl field if non-nil, zero value otherwise.

### GetForwardUrlOk

`func (o *CreateWebhookCliSessionRequest) GetForwardUrlOk() (*string, bool)`

GetForwardUrlOk returns a tuple with the ForwardUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardUrl

`func (o *CreateWebhookCliSessionRequest) SetForwardUrl(v string)`

SetForwardUrl sets ForwardUrl field to given value.

### HasForwardUrl

`func (o *CreateWebhookCliSessionRequest) HasForwardUrl() bool`

HasForwardUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


