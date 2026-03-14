# CreateContactList201Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SchemaVersion** | Pointer to **string** |  | [optional] 
**RequestId** | Pointer to **string** |  | [optional] 
**ContactList** | Pointer to [**ContactList**](ContactList.md) |  | [optional] 

## Methods

### NewCreateContactList201Response

`func NewCreateContactList201Response() *CreateContactList201Response`

NewCreateContactList201Response instantiates a new CreateContactList201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateContactList201ResponseWithDefaults

`func NewCreateContactList201ResponseWithDefaults() *CreateContactList201Response`

NewCreateContactList201ResponseWithDefaults instantiates a new CreateContactList201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemaVersion

`func (o *CreateContactList201Response) GetSchemaVersion() string`

GetSchemaVersion returns the SchemaVersion field if non-nil, zero value otherwise.

### GetSchemaVersionOk

`func (o *CreateContactList201Response) GetSchemaVersionOk() (*string, bool)`

GetSchemaVersionOk returns a tuple with the SchemaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaVersion

`func (o *CreateContactList201Response) SetSchemaVersion(v string)`

SetSchemaVersion sets SchemaVersion field to given value.

### HasSchemaVersion

`func (o *CreateContactList201Response) HasSchemaVersion() bool`

HasSchemaVersion returns a boolean if a field has been set.

### GetRequestId

`func (o *CreateContactList201Response) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *CreateContactList201Response) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *CreateContactList201Response) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *CreateContactList201Response) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetContactList

`func (o *CreateContactList201Response) GetContactList() ContactList`

GetContactList returns the ContactList field if non-nil, zero value otherwise.

### GetContactListOk

`func (o *CreateContactList201Response) GetContactListOk() (*ContactList, bool)`

GetContactListOk returns a tuple with the ContactList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactList

`func (o *CreateContactList201Response) SetContactList(v ContactList)`

SetContactList sets ContactList field to given value.

### HasContactList

`func (o *CreateContactList201Response) HasContactList() bool`

HasContactList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


