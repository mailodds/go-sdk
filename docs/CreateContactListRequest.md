# CreateContactListRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | List name | 
**SourceJobId** | Pointer to **string** | Single validation job ID to build list from | [optional] 
**SourceJobIds** | Pointer to **[]string** | Multiple validation job IDs to merge into one list | [optional] 
**Tags** | Pointer to **[]string** | Tags for categorization | [optional] 

## Methods

### NewCreateContactListRequest

`func NewCreateContactListRequest(name string, ) *CreateContactListRequest`

NewCreateContactListRequest instantiates a new CreateContactListRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateContactListRequestWithDefaults

`func NewCreateContactListRequestWithDefaults() *CreateContactListRequest`

NewCreateContactListRequestWithDefaults instantiates a new CreateContactListRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateContactListRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateContactListRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateContactListRequest) SetName(v string)`

SetName sets Name field to given value.


### GetSourceJobId

`func (o *CreateContactListRequest) GetSourceJobId() string`

GetSourceJobId returns the SourceJobId field if non-nil, zero value otherwise.

### GetSourceJobIdOk

`func (o *CreateContactListRequest) GetSourceJobIdOk() (*string, bool)`

GetSourceJobIdOk returns a tuple with the SourceJobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceJobId

`func (o *CreateContactListRequest) SetSourceJobId(v string)`

SetSourceJobId sets SourceJobId field to given value.

### HasSourceJobId

`func (o *CreateContactListRequest) HasSourceJobId() bool`

HasSourceJobId returns a boolean if a field has been set.

### GetSourceJobIds

`func (o *CreateContactListRequest) GetSourceJobIds() []string`

GetSourceJobIds returns the SourceJobIds field if non-nil, zero value otherwise.

### GetSourceJobIdsOk

`func (o *CreateContactListRequest) GetSourceJobIdsOk() (*[]string, bool)`

GetSourceJobIdsOk returns a tuple with the SourceJobIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceJobIds

`func (o *CreateContactListRequest) SetSourceJobIds(v []string)`

SetSourceJobIds sets SourceJobIds field to given value.

### HasSourceJobIds

`func (o *CreateContactListRequest) HasSourceJobIds() bool`

HasSourceJobIds returns a boolean if a field has been set.

### GetTags

`func (o *CreateContactListRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CreateContactListRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CreateContactListRequest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CreateContactListRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


