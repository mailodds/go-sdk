# CreateListRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | List name (unique per account) | 
**Description** | Pointer to **string** | Optional list description | [optional] 
**ConfirmationRedirectUrl** | Pointer to **string** | URL to redirect subscribers after confirmation | [optional] 
**ConfirmationSubject** | Pointer to **string** | Custom confirmation email subject | [optional] 
**ConfirmationFromName** | Pointer to **string** | Custom sender name for confirmation emails | [optional] 

## Methods

### NewCreateListRequest

`func NewCreateListRequest(name string, ) *CreateListRequest`

NewCreateListRequest instantiates a new CreateListRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateListRequestWithDefaults

`func NewCreateListRequestWithDefaults() *CreateListRequest`

NewCreateListRequestWithDefaults instantiates a new CreateListRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateListRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateListRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateListRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CreateListRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateListRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateListRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateListRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetConfirmationRedirectUrl

`func (o *CreateListRequest) GetConfirmationRedirectUrl() string`

GetConfirmationRedirectUrl returns the ConfirmationRedirectUrl field if non-nil, zero value otherwise.

### GetConfirmationRedirectUrlOk

`func (o *CreateListRequest) GetConfirmationRedirectUrlOk() (*string, bool)`

GetConfirmationRedirectUrlOk returns a tuple with the ConfirmationRedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmationRedirectUrl

`func (o *CreateListRequest) SetConfirmationRedirectUrl(v string)`

SetConfirmationRedirectUrl sets ConfirmationRedirectUrl field to given value.

### HasConfirmationRedirectUrl

`func (o *CreateListRequest) HasConfirmationRedirectUrl() bool`

HasConfirmationRedirectUrl returns a boolean if a field has been set.

### GetConfirmationSubject

`func (o *CreateListRequest) GetConfirmationSubject() string`

GetConfirmationSubject returns the ConfirmationSubject field if non-nil, zero value otherwise.

### GetConfirmationSubjectOk

`func (o *CreateListRequest) GetConfirmationSubjectOk() (*string, bool)`

GetConfirmationSubjectOk returns a tuple with the ConfirmationSubject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmationSubject

`func (o *CreateListRequest) SetConfirmationSubject(v string)`

SetConfirmationSubject sets ConfirmationSubject field to given value.

### HasConfirmationSubject

`func (o *CreateListRequest) HasConfirmationSubject() bool`

HasConfirmationSubject returns a boolean if a field has been set.

### GetConfirmationFromName

`func (o *CreateListRequest) GetConfirmationFromName() string`

GetConfirmationFromName returns the ConfirmationFromName field if non-nil, zero value otherwise.

### GetConfirmationFromNameOk

`func (o *CreateListRequest) GetConfirmationFromNameOk() (*string, bool)`

GetConfirmationFromNameOk returns a tuple with the ConfirmationFromName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmationFromName

`func (o *CreateListRequest) SetConfirmationFromName(v string)`

SetConfirmationFromName sets ConfirmationFromName field to given value.

### HasConfirmationFromName

`func (o *CreateListRequest) HasConfirmationFromName() bool`

HasConfirmationFromName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


