# CreateStore201Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SchemaVersion** | Pointer to **string** |  | [optional] 
**RequestId** | Pointer to **string** |  | [optional] 
**Store** | Pointer to [**StoreConnection**](StoreConnection.md) |  | [optional] 

## Methods

### NewCreateStore201Response

`func NewCreateStore201Response() *CreateStore201Response`

NewCreateStore201Response instantiates a new CreateStore201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateStore201ResponseWithDefaults

`func NewCreateStore201ResponseWithDefaults() *CreateStore201Response`

NewCreateStore201ResponseWithDefaults instantiates a new CreateStore201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemaVersion

`func (o *CreateStore201Response) GetSchemaVersion() string`

GetSchemaVersion returns the SchemaVersion field if non-nil, zero value otherwise.

### GetSchemaVersionOk

`func (o *CreateStore201Response) GetSchemaVersionOk() (*string, bool)`

GetSchemaVersionOk returns a tuple with the SchemaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaVersion

`func (o *CreateStore201Response) SetSchemaVersion(v string)`

SetSchemaVersion sets SchemaVersion field to given value.

### HasSchemaVersion

`func (o *CreateStore201Response) HasSchemaVersion() bool`

HasSchemaVersion returns a boolean if a field has been set.

### GetRequestId

`func (o *CreateStore201Response) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *CreateStore201Response) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *CreateStore201Response) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *CreateStore201Response) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetStore

`func (o *CreateStore201Response) GetStore() StoreConnection`

GetStore returns the Store field if non-nil, zero value otherwise.

### GetStoreOk

`func (o *CreateStore201Response) GetStoreOk() (*StoreConnection, bool)`

GetStoreOk returns a tuple with the Store field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStore

`func (o *CreateStore201Response) SetStore(v StoreConnection)`

SetStore sets Store field to given value.

### HasStore

`func (o *CreateStore201Response) HasStore() bool`

HasStore returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


