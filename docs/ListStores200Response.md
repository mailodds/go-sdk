# ListStores200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SchemaVersion** | Pointer to **string** |  | [optional] 
**RequestId** | Pointer to **string** |  | [optional] 
**Stores** | Pointer to [**[]StoreConnection**](StoreConnection.md) |  | [optional] 

## Methods

### NewListStores200Response

`func NewListStores200Response() *ListStores200Response`

NewListStores200Response instantiates a new ListStores200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListStores200ResponseWithDefaults

`func NewListStores200ResponseWithDefaults() *ListStores200Response`

NewListStores200ResponseWithDefaults instantiates a new ListStores200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemaVersion

`func (o *ListStores200Response) GetSchemaVersion() string`

GetSchemaVersion returns the SchemaVersion field if non-nil, zero value otherwise.

### GetSchemaVersionOk

`func (o *ListStores200Response) GetSchemaVersionOk() (*string, bool)`

GetSchemaVersionOk returns a tuple with the SchemaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaVersion

`func (o *ListStores200Response) SetSchemaVersion(v string)`

SetSchemaVersion sets SchemaVersion field to given value.

### HasSchemaVersion

`func (o *ListStores200Response) HasSchemaVersion() bool`

HasSchemaVersion returns a boolean if a field has been set.

### GetRequestId

`func (o *ListStores200Response) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *ListStores200Response) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *ListStores200Response) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *ListStores200Response) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetStores

`func (o *ListStores200Response) GetStores() []StoreConnection`

GetStores returns the Stores field if non-nil, zero value otherwise.

### GetStoresOk

`func (o *ListStores200Response) GetStoresOk() (*[]StoreConnection, bool)`

GetStoresOk returns a tuple with the Stores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStores

`func (o *ListStores200Response) SetStores(v []StoreConnection)`

SetStores sets Stores field to given value.

### HasStores

`func (o *ListStores200Response) HasStores() bool`

HasStores returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


