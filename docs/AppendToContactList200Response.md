# AppendToContactList200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SchemaVersion** | Pointer to **string** |  | [optional] 
**RequestId** | Pointer to **string** |  | [optional] 
**ContactList** | Pointer to [**ContactList**](ContactList.md) |  | [optional] 
**AddedCount** | Pointer to **int32** | Number of new emails added | [optional] 
**CandidateCount** | Pointer to **int32** | Total candidates from jobs | [optional] 
**DuplicateCount** | Pointer to **int32** | Duplicates skipped | [optional] 
**Breakdown** | Pointer to **map[string]interface{}** | Per-status breakdown of candidates | [optional] 

## Methods

### NewAppendToContactList200Response

`func NewAppendToContactList200Response() *AppendToContactList200Response`

NewAppendToContactList200Response instantiates a new AppendToContactList200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppendToContactList200ResponseWithDefaults

`func NewAppendToContactList200ResponseWithDefaults() *AppendToContactList200Response`

NewAppendToContactList200ResponseWithDefaults instantiates a new AppendToContactList200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemaVersion

`func (o *AppendToContactList200Response) GetSchemaVersion() string`

GetSchemaVersion returns the SchemaVersion field if non-nil, zero value otherwise.

### GetSchemaVersionOk

`func (o *AppendToContactList200Response) GetSchemaVersionOk() (*string, bool)`

GetSchemaVersionOk returns a tuple with the SchemaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaVersion

`func (o *AppendToContactList200Response) SetSchemaVersion(v string)`

SetSchemaVersion sets SchemaVersion field to given value.

### HasSchemaVersion

`func (o *AppendToContactList200Response) HasSchemaVersion() bool`

HasSchemaVersion returns a boolean if a field has been set.

### GetRequestId

`func (o *AppendToContactList200Response) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *AppendToContactList200Response) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *AppendToContactList200Response) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *AppendToContactList200Response) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetContactList

`func (o *AppendToContactList200Response) GetContactList() ContactList`

GetContactList returns the ContactList field if non-nil, zero value otherwise.

### GetContactListOk

`func (o *AppendToContactList200Response) GetContactListOk() (*ContactList, bool)`

GetContactListOk returns a tuple with the ContactList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactList

`func (o *AppendToContactList200Response) SetContactList(v ContactList)`

SetContactList sets ContactList field to given value.

### HasContactList

`func (o *AppendToContactList200Response) HasContactList() bool`

HasContactList returns a boolean if a field has been set.

### GetAddedCount

`func (o *AppendToContactList200Response) GetAddedCount() int32`

GetAddedCount returns the AddedCount field if non-nil, zero value otherwise.

### GetAddedCountOk

`func (o *AppendToContactList200Response) GetAddedCountOk() (*int32, bool)`

GetAddedCountOk returns a tuple with the AddedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddedCount

`func (o *AppendToContactList200Response) SetAddedCount(v int32)`

SetAddedCount sets AddedCount field to given value.

### HasAddedCount

`func (o *AppendToContactList200Response) HasAddedCount() bool`

HasAddedCount returns a boolean if a field has been set.

### GetCandidateCount

`func (o *AppendToContactList200Response) GetCandidateCount() int32`

GetCandidateCount returns the CandidateCount field if non-nil, zero value otherwise.

### GetCandidateCountOk

`func (o *AppendToContactList200Response) GetCandidateCountOk() (*int32, bool)`

GetCandidateCountOk returns a tuple with the CandidateCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCandidateCount

`func (o *AppendToContactList200Response) SetCandidateCount(v int32)`

SetCandidateCount sets CandidateCount field to given value.

### HasCandidateCount

`func (o *AppendToContactList200Response) HasCandidateCount() bool`

HasCandidateCount returns a boolean if a field has been set.

### GetDuplicateCount

`func (o *AppendToContactList200Response) GetDuplicateCount() int32`

GetDuplicateCount returns the DuplicateCount field if non-nil, zero value otherwise.

### GetDuplicateCountOk

`func (o *AppendToContactList200Response) GetDuplicateCountOk() (*int32, bool)`

GetDuplicateCountOk returns a tuple with the DuplicateCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuplicateCount

`func (o *AppendToContactList200Response) SetDuplicateCount(v int32)`

SetDuplicateCount sets DuplicateCount field to given value.

### HasDuplicateCount

`func (o *AppendToContactList200Response) HasDuplicateCount() bool`

HasDuplicateCount returns a boolean if a field has been set.

### GetBreakdown

`func (o *AppendToContactList200Response) GetBreakdown() map[string]interface{}`

GetBreakdown returns the Breakdown field if non-nil, zero value otherwise.

### GetBreakdownOk

`func (o *AppendToContactList200Response) GetBreakdownOk() (*map[string]interface{}, bool)`

GetBreakdownOk returns a tuple with the Breakdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreakdown

`func (o *AppendToContactList200Response) SetBreakdown(v map[string]interface{})`

SetBreakdown sets Breakdown field to given value.

### HasBreakdown

`func (o *AppendToContactList200Response) HasBreakdown() bool`

HasBreakdown returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


