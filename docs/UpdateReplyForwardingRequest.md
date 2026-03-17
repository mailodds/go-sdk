# UpdateReplyForwardingRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ForwardRepliesTo** | Pointer to **NullableString** | Email to forward replies to, or null to disable | [optional] 

## Methods

### NewUpdateReplyForwardingRequest

`func NewUpdateReplyForwardingRequest() *UpdateReplyForwardingRequest`

NewUpdateReplyForwardingRequest instantiates a new UpdateReplyForwardingRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateReplyForwardingRequestWithDefaults

`func NewUpdateReplyForwardingRequestWithDefaults() *UpdateReplyForwardingRequest`

NewUpdateReplyForwardingRequestWithDefaults instantiates a new UpdateReplyForwardingRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetForwardRepliesTo

`func (o *UpdateReplyForwardingRequest) GetForwardRepliesTo() string`

GetForwardRepliesTo returns the ForwardRepliesTo field if non-nil, zero value otherwise.

### GetForwardRepliesToOk

`func (o *UpdateReplyForwardingRequest) GetForwardRepliesToOk() (*string, bool)`

GetForwardRepliesToOk returns a tuple with the ForwardRepliesTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardRepliesTo

`func (o *UpdateReplyForwardingRequest) SetForwardRepliesTo(v string)`

SetForwardRepliesTo sets ForwardRepliesTo field to given value.

### HasForwardRepliesTo

`func (o *UpdateReplyForwardingRequest) HasForwardRepliesTo() bool`

HasForwardRepliesTo returns a boolean if a field has been set.

### SetForwardRepliesToNil

`func (o *UpdateReplyForwardingRequest) SetForwardRepliesToNil(b bool)`

 SetForwardRepliesToNil sets the value for ForwardRepliesTo to be an explicit nil

### UnsetForwardRepliesTo
`func (o *UpdateReplyForwardingRequest) UnsetForwardRepliesTo()`

UnsetForwardRepliesTo ensures that no value is present for ForwardRepliesTo, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


