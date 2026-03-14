# AppendToContactListRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceJobIds** | **[]string** | Validation job IDs to append from | 
**IncludeCatchAll** | Pointer to **bool** | Include catch-all emails in addition to valid ones | [optional] [default to false]

## Methods

### NewAppendToContactListRequest

`func NewAppendToContactListRequest(sourceJobIds []string, ) *AppendToContactListRequest`

NewAppendToContactListRequest instantiates a new AppendToContactListRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppendToContactListRequestWithDefaults

`func NewAppendToContactListRequestWithDefaults() *AppendToContactListRequest`

NewAppendToContactListRequestWithDefaults instantiates a new AppendToContactListRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceJobIds

`func (o *AppendToContactListRequest) GetSourceJobIds() []string`

GetSourceJobIds returns the SourceJobIds field if non-nil, zero value otherwise.

### GetSourceJobIdsOk

`func (o *AppendToContactListRequest) GetSourceJobIdsOk() (*[]string, bool)`

GetSourceJobIdsOk returns a tuple with the SourceJobIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceJobIds

`func (o *AppendToContactListRequest) SetSourceJobIds(v []string)`

SetSourceJobIds sets SourceJobIds field to given value.


### GetIncludeCatchAll

`func (o *AppendToContactListRequest) GetIncludeCatchAll() bool`

GetIncludeCatchAll returns the IncludeCatchAll field if non-nil, zero value otherwise.

### GetIncludeCatchAllOk

`func (o *AppendToContactListRequest) GetIncludeCatchAllOk() (*bool, bool)`

GetIncludeCatchAllOk returns a tuple with the IncludeCatchAll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeCatchAll

`func (o *AppendToContactListRequest) SetIncludeCatchAll(v bool)`

SetIncludeCatchAll sets IncludeCatchAll field to given value.

### HasIncludeCatchAll

`func (o *AppendToContactListRequest) HasIncludeCatchAll() bool`

HasIncludeCatchAll returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


