# ContactList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Contact list UUID | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**EmailCount** | Pointer to **int32** | Number of emails in the list | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**SourceJobIds** | Pointer to **[]string** | Validation job IDs this list was built from | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewContactList

`func NewContactList() *ContactList`

NewContactList instantiates a new ContactList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContactListWithDefaults

`func NewContactListWithDefaults() *ContactList`

NewContactListWithDefaults instantiates a new ContactList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ContactList) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ContactList) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ContactList) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ContactList) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ContactList) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ContactList) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ContactList) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ContactList) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEmailCount

`func (o *ContactList) GetEmailCount() int32`

GetEmailCount returns the EmailCount field if non-nil, zero value otherwise.

### GetEmailCountOk

`func (o *ContactList) GetEmailCountOk() (*int32, bool)`

GetEmailCountOk returns a tuple with the EmailCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailCount

`func (o *ContactList) SetEmailCount(v int32)`

SetEmailCount sets EmailCount field to given value.

### HasEmailCount

`func (o *ContactList) HasEmailCount() bool`

HasEmailCount returns a boolean if a field has been set.

### GetTags

`func (o *ContactList) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ContactList) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ContactList) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ContactList) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetSourceJobIds

`func (o *ContactList) GetSourceJobIds() []string`

GetSourceJobIds returns the SourceJobIds field if non-nil, zero value otherwise.

### GetSourceJobIdsOk

`func (o *ContactList) GetSourceJobIdsOk() (*[]string, bool)`

GetSourceJobIdsOk returns a tuple with the SourceJobIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceJobIds

`func (o *ContactList) SetSourceJobIds(v []string)`

SetSourceJobIds sets SourceJobIds field to given value.

### HasSourceJobIds

`func (o *ContactList) HasSourceJobIds() bool`

HasSourceJobIds returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ContactList) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ContactList) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ContactList) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ContactList) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ContactList) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ContactList) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ContactList) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ContactList) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


