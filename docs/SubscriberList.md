# SubscriberList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | List UUID | [optional] 
**AccountId** | Pointer to **int32** | Account ID | [optional] 
**Name** | Pointer to **string** | List name | [optional] 
**Description** | Pointer to **NullableString** | List description | [optional] 
**ConfirmationRedirectUrl** | Pointer to **NullableString** | Redirect URL after confirmation | [optional] 
**ConfirmationSubject** | Pointer to **NullableString** | Confirmation email subject | [optional] 
**ConfirmationFromName** | Pointer to **NullableString** | Confirmation email sender name | [optional] 
**SubscriberCount** | Pointer to **int32** | Total subscriber count | [optional] 
**ConfirmedCount** | Pointer to **int32** | Confirmed subscriber count | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewSubscriberList

`func NewSubscriberList() *SubscriberList`

NewSubscriberList instantiates a new SubscriberList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriberListWithDefaults

`func NewSubscriberListWithDefaults() *SubscriberList`

NewSubscriberListWithDefaults instantiates a new SubscriberList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SubscriberList) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SubscriberList) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SubscriberList) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SubscriberList) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAccountId

`func (o *SubscriberList) GetAccountId() int32`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *SubscriberList) GetAccountIdOk() (*int32, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *SubscriberList) SetAccountId(v int32)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *SubscriberList) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetName

`func (o *SubscriberList) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SubscriberList) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SubscriberList) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SubscriberList) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *SubscriberList) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SubscriberList) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SubscriberList) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SubscriberList) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *SubscriberList) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *SubscriberList) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetConfirmationRedirectUrl

`func (o *SubscriberList) GetConfirmationRedirectUrl() string`

GetConfirmationRedirectUrl returns the ConfirmationRedirectUrl field if non-nil, zero value otherwise.

### GetConfirmationRedirectUrlOk

`func (o *SubscriberList) GetConfirmationRedirectUrlOk() (*string, bool)`

GetConfirmationRedirectUrlOk returns a tuple with the ConfirmationRedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmationRedirectUrl

`func (o *SubscriberList) SetConfirmationRedirectUrl(v string)`

SetConfirmationRedirectUrl sets ConfirmationRedirectUrl field to given value.

### HasConfirmationRedirectUrl

`func (o *SubscriberList) HasConfirmationRedirectUrl() bool`

HasConfirmationRedirectUrl returns a boolean if a field has been set.

### SetConfirmationRedirectUrlNil

`func (o *SubscriberList) SetConfirmationRedirectUrlNil(b bool)`

 SetConfirmationRedirectUrlNil sets the value for ConfirmationRedirectUrl to be an explicit nil

### UnsetConfirmationRedirectUrl
`func (o *SubscriberList) UnsetConfirmationRedirectUrl()`

UnsetConfirmationRedirectUrl ensures that no value is present for ConfirmationRedirectUrl, not even an explicit nil
### GetConfirmationSubject

`func (o *SubscriberList) GetConfirmationSubject() string`

GetConfirmationSubject returns the ConfirmationSubject field if non-nil, zero value otherwise.

### GetConfirmationSubjectOk

`func (o *SubscriberList) GetConfirmationSubjectOk() (*string, bool)`

GetConfirmationSubjectOk returns a tuple with the ConfirmationSubject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmationSubject

`func (o *SubscriberList) SetConfirmationSubject(v string)`

SetConfirmationSubject sets ConfirmationSubject field to given value.

### HasConfirmationSubject

`func (o *SubscriberList) HasConfirmationSubject() bool`

HasConfirmationSubject returns a boolean if a field has been set.

### SetConfirmationSubjectNil

`func (o *SubscriberList) SetConfirmationSubjectNil(b bool)`

 SetConfirmationSubjectNil sets the value for ConfirmationSubject to be an explicit nil

### UnsetConfirmationSubject
`func (o *SubscriberList) UnsetConfirmationSubject()`

UnsetConfirmationSubject ensures that no value is present for ConfirmationSubject, not even an explicit nil
### GetConfirmationFromName

`func (o *SubscriberList) GetConfirmationFromName() string`

GetConfirmationFromName returns the ConfirmationFromName field if non-nil, zero value otherwise.

### GetConfirmationFromNameOk

`func (o *SubscriberList) GetConfirmationFromNameOk() (*string, bool)`

GetConfirmationFromNameOk returns a tuple with the ConfirmationFromName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmationFromName

`func (o *SubscriberList) SetConfirmationFromName(v string)`

SetConfirmationFromName sets ConfirmationFromName field to given value.

### HasConfirmationFromName

`func (o *SubscriberList) HasConfirmationFromName() bool`

HasConfirmationFromName returns a boolean if a field has been set.

### SetConfirmationFromNameNil

`func (o *SubscriberList) SetConfirmationFromNameNil(b bool)`

 SetConfirmationFromNameNil sets the value for ConfirmationFromName to be an explicit nil

### UnsetConfirmationFromName
`func (o *SubscriberList) UnsetConfirmationFromName()`

UnsetConfirmationFromName ensures that no value is present for ConfirmationFromName, not even an explicit nil
### GetSubscriberCount

`func (o *SubscriberList) GetSubscriberCount() int32`

GetSubscriberCount returns the SubscriberCount field if non-nil, zero value otherwise.

### GetSubscriberCountOk

`func (o *SubscriberList) GetSubscriberCountOk() (*int32, bool)`

GetSubscriberCountOk returns a tuple with the SubscriberCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriberCount

`func (o *SubscriberList) SetSubscriberCount(v int32)`

SetSubscriberCount sets SubscriberCount field to given value.

### HasSubscriberCount

`func (o *SubscriberList) HasSubscriberCount() bool`

HasSubscriberCount returns a boolean if a field has been set.

### GetConfirmedCount

`func (o *SubscriberList) GetConfirmedCount() int32`

GetConfirmedCount returns the ConfirmedCount field if non-nil, zero value otherwise.

### GetConfirmedCountOk

`func (o *SubscriberList) GetConfirmedCountOk() (*int32, bool)`

GetConfirmedCountOk returns a tuple with the ConfirmedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmedCount

`func (o *SubscriberList) SetConfirmedCount(v int32)`

SetConfirmedCount sets ConfirmedCount field to given value.

### HasConfirmedCount

`func (o *SubscriberList) HasConfirmedCount() bool`

HasConfirmedCount returns a boolean if a field has been set.

### GetCreatedAt

`func (o *SubscriberList) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SubscriberList) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SubscriberList) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *SubscriberList) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *SubscriberList) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *SubscriberList) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *SubscriberList) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *SubscriberList) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


