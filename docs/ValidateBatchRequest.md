# ValidateBatchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Emails** | **[]string** | List of emails to validate | 
**Depth** | Pointer to **string** |  | [optional] [default to "enhanced"]
**PolicyId** | Pointer to **int32** | Optional policy ID | [optional] 

## Methods

### NewValidateBatchRequest

`func NewValidateBatchRequest(emails []string, ) *ValidateBatchRequest`

NewValidateBatchRequest instantiates a new ValidateBatchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidateBatchRequestWithDefaults

`func NewValidateBatchRequestWithDefaults() *ValidateBatchRequest`

NewValidateBatchRequestWithDefaults instantiates a new ValidateBatchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmails

`func (o *ValidateBatchRequest) GetEmails() []string`

GetEmails returns the Emails field if non-nil, zero value otherwise.

### GetEmailsOk

`func (o *ValidateBatchRequest) GetEmailsOk() (*[]string, bool)`

GetEmailsOk returns a tuple with the Emails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmails

`func (o *ValidateBatchRequest) SetEmails(v []string)`

SetEmails sets Emails field to given value.


### GetDepth

`func (o *ValidateBatchRequest) GetDepth() string`

GetDepth returns the Depth field if non-nil, zero value otherwise.

### GetDepthOk

`func (o *ValidateBatchRequest) GetDepthOk() (*string, bool)`

GetDepthOk returns a tuple with the Depth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepth

`func (o *ValidateBatchRequest) SetDepth(v string)`

SetDepth sets Depth field to given value.

### HasDepth

`func (o *ValidateBatchRequest) HasDepth() bool`

HasDepth returns a boolean if a field has been set.

### GetPolicyId

`func (o *ValidateBatchRequest) GetPolicyId() int32`

GetPolicyId returns the PolicyId field if non-nil, zero value otherwise.

### GetPolicyIdOk

`func (o *ValidateBatchRequest) GetPolicyIdOk() (*int32, bool)`

GetPolicyIdOk returns a tuple with the PolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyId

`func (o *ValidateBatchRequest) SetPolicyId(v int32)`

SetPolicyId sets PolicyId field to given value.

### HasPolicyId

`func (o *ValidateBatchRequest) HasPolicyId() bool`

HasPolicyId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


