# Subscriber

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Subscriber UUID | [optional] 
**ListId** | Pointer to **string** | List UUID | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**ConsentSourceIp** | Pointer to **NullableString** | IP address of subscription | [optional] 
**ConsentPageUrl** | Pointer to **NullableString** | Page URL where form was submitted | [optional] 
**ConsentFormId** | Pointer to **NullableString** | Form identifier | [optional] 
**ConsentTimestamp** | Pointer to **NullableTime** |  | [optional] 
**ConfirmedAt** | Pointer to **NullableTime** |  | [optional] 
**UnsubscribedAt** | Pointer to **NullableTime** |  | [optional] 
**ValidationResult** | Pointer to **map[string]interface{}** | Email validation result | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Custom metadata | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewSubscriber

`func NewSubscriber() *Subscriber`

NewSubscriber instantiates a new Subscriber object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriberWithDefaults

`func NewSubscriberWithDefaults() *Subscriber`

NewSubscriberWithDefaults instantiates a new Subscriber object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Subscriber) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Subscriber) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Subscriber) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Subscriber) HasId() bool`

HasId returns a boolean if a field has been set.

### GetListId

`func (o *Subscriber) GetListId() string`

GetListId returns the ListId field if non-nil, zero value otherwise.

### GetListIdOk

`func (o *Subscriber) GetListIdOk() (*string, bool)`

GetListIdOk returns a tuple with the ListId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListId

`func (o *Subscriber) SetListId(v string)`

SetListId sets ListId field to given value.

### HasListId

`func (o *Subscriber) HasListId() bool`

HasListId returns a boolean if a field has been set.

### GetEmail

`func (o *Subscriber) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *Subscriber) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *Subscriber) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *Subscriber) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetName

`func (o *Subscriber) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Subscriber) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Subscriber) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Subscriber) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *Subscriber) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *Subscriber) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetStatus

`func (o *Subscriber) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Subscriber) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Subscriber) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Subscriber) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetConsentSourceIp

`func (o *Subscriber) GetConsentSourceIp() string`

GetConsentSourceIp returns the ConsentSourceIp field if non-nil, zero value otherwise.

### GetConsentSourceIpOk

`func (o *Subscriber) GetConsentSourceIpOk() (*string, bool)`

GetConsentSourceIpOk returns a tuple with the ConsentSourceIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsentSourceIp

`func (o *Subscriber) SetConsentSourceIp(v string)`

SetConsentSourceIp sets ConsentSourceIp field to given value.

### HasConsentSourceIp

`func (o *Subscriber) HasConsentSourceIp() bool`

HasConsentSourceIp returns a boolean if a field has been set.

### SetConsentSourceIpNil

`func (o *Subscriber) SetConsentSourceIpNil(b bool)`

 SetConsentSourceIpNil sets the value for ConsentSourceIp to be an explicit nil

### UnsetConsentSourceIp
`func (o *Subscriber) UnsetConsentSourceIp()`

UnsetConsentSourceIp ensures that no value is present for ConsentSourceIp, not even an explicit nil
### GetConsentPageUrl

`func (o *Subscriber) GetConsentPageUrl() string`

GetConsentPageUrl returns the ConsentPageUrl field if non-nil, zero value otherwise.

### GetConsentPageUrlOk

`func (o *Subscriber) GetConsentPageUrlOk() (*string, bool)`

GetConsentPageUrlOk returns a tuple with the ConsentPageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsentPageUrl

`func (o *Subscriber) SetConsentPageUrl(v string)`

SetConsentPageUrl sets ConsentPageUrl field to given value.

### HasConsentPageUrl

`func (o *Subscriber) HasConsentPageUrl() bool`

HasConsentPageUrl returns a boolean if a field has been set.

### SetConsentPageUrlNil

`func (o *Subscriber) SetConsentPageUrlNil(b bool)`

 SetConsentPageUrlNil sets the value for ConsentPageUrl to be an explicit nil

### UnsetConsentPageUrl
`func (o *Subscriber) UnsetConsentPageUrl()`

UnsetConsentPageUrl ensures that no value is present for ConsentPageUrl, not even an explicit nil
### GetConsentFormId

`func (o *Subscriber) GetConsentFormId() string`

GetConsentFormId returns the ConsentFormId field if non-nil, zero value otherwise.

### GetConsentFormIdOk

`func (o *Subscriber) GetConsentFormIdOk() (*string, bool)`

GetConsentFormIdOk returns a tuple with the ConsentFormId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsentFormId

`func (o *Subscriber) SetConsentFormId(v string)`

SetConsentFormId sets ConsentFormId field to given value.

### HasConsentFormId

`func (o *Subscriber) HasConsentFormId() bool`

HasConsentFormId returns a boolean if a field has been set.

### SetConsentFormIdNil

`func (o *Subscriber) SetConsentFormIdNil(b bool)`

 SetConsentFormIdNil sets the value for ConsentFormId to be an explicit nil

### UnsetConsentFormId
`func (o *Subscriber) UnsetConsentFormId()`

UnsetConsentFormId ensures that no value is present for ConsentFormId, not even an explicit nil
### GetConsentTimestamp

`func (o *Subscriber) GetConsentTimestamp() time.Time`

GetConsentTimestamp returns the ConsentTimestamp field if non-nil, zero value otherwise.

### GetConsentTimestampOk

`func (o *Subscriber) GetConsentTimestampOk() (*time.Time, bool)`

GetConsentTimestampOk returns a tuple with the ConsentTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsentTimestamp

`func (o *Subscriber) SetConsentTimestamp(v time.Time)`

SetConsentTimestamp sets ConsentTimestamp field to given value.

### HasConsentTimestamp

`func (o *Subscriber) HasConsentTimestamp() bool`

HasConsentTimestamp returns a boolean if a field has been set.

### SetConsentTimestampNil

`func (o *Subscriber) SetConsentTimestampNil(b bool)`

 SetConsentTimestampNil sets the value for ConsentTimestamp to be an explicit nil

### UnsetConsentTimestamp
`func (o *Subscriber) UnsetConsentTimestamp()`

UnsetConsentTimestamp ensures that no value is present for ConsentTimestamp, not even an explicit nil
### GetConfirmedAt

`func (o *Subscriber) GetConfirmedAt() time.Time`

GetConfirmedAt returns the ConfirmedAt field if non-nil, zero value otherwise.

### GetConfirmedAtOk

`func (o *Subscriber) GetConfirmedAtOk() (*time.Time, bool)`

GetConfirmedAtOk returns a tuple with the ConfirmedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmedAt

`func (o *Subscriber) SetConfirmedAt(v time.Time)`

SetConfirmedAt sets ConfirmedAt field to given value.

### HasConfirmedAt

`func (o *Subscriber) HasConfirmedAt() bool`

HasConfirmedAt returns a boolean if a field has been set.

### SetConfirmedAtNil

`func (o *Subscriber) SetConfirmedAtNil(b bool)`

 SetConfirmedAtNil sets the value for ConfirmedAt to be an explicit nil

### UnsetConfirmedAt
`func (o *Subscriber) UnsetConfirmedAt()`

UnsetConfirmedAt ensures that no value is present for ConfirmedAt, not even an explicit nil
### GetUnsubscribedAt

`func (o *Subscriber) GetUnsubscribedAt() time.Time`

GetUnsubscribedAt returns the UnsubscribedAt field if non-nil, zero value otherwise.

### GetUnsubscribedAtOk

`func (o *Subscriber) GetUnsubscribedAtOk() (*time.Time, bool)`

GetUnsubscribedAtOk returns a tuple with the UnsubscribedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnsubscribedAt

`func (o *Subscriber) SetUnsubscribedAt(v time.Time)`

SetUnsubscribedAt sets UnsubscribedAt field to given value.

### HasUnsubscribedAt

`func (o *Subscriber) HasUnsubscribedAt() bool`

HasUnsubscribedAt returns a boolean if a field has been set.

### SetUnsubscribedAtNil

`func (o *Subscriber) SetUnsubscribedAtNil(b bool)`

 SetUnsubscribedAtNil sets the value for UnsubscribedAt to be an explicit nil

### UnsetUnsubscribedAt
`func (o *Subscriber) UnsetUnsubscribedAt()`

UnsetUnsubscribedAt ensures that no value is present for UnsubscribedAt, not even an explicit nil
### GetValidationResult

`func (o *Subscriber) GetValidationResult() map[string]interface{}`

GetValidationResult returns the ValidationResult field if non-nil, zero value otherwise.

### GetValidationResultOk

`func (o *Subscriber) GetValidationResultOk() (*map[string]interface{}, bool)`

GetValidationResultOk returns a tuple with the ValidationResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationResult

`func (o *Subscriber) SetValidationResult(v map[string]interface{})`

SetValidationResult sets ValidationResult field to given value.

### HasValidationResult

`func (o *Subscriber) HasValidationResult() bool`

HasValidationResult returns a boolean if a field has been set.

### SetValidationResultNil

`func (o *Subscriber) SetValidationResultNil(b bool)`

 SetValidationResultNil sets the value for ValidationResult to be an explicit nil

### UnsetValidationResult
`func (o *Subscriber) UnsetValidationResult()`

UnsetValidationResult ensures that no value is present for ValidationResult, not even an explicit nil
### GetMetadata

`func (o *Subscriber) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Subscriber) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Subscriber) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Subscriber) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Subscriber) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Subscriber) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Subscriber) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Subscriber) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


