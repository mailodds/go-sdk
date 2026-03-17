# BatchProductsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SchemaVersion** | Pointer to **string** |  | [optional] 
**RequestId** | Pointer to **string** |  | [optional] 
**Created** | Pointer to **int32** | Products created | [optional] 
**Updated** | Pointer to **int32** | Products updated | [optional] 
**Errored** | Pointer to **int32** | Products that failed | [optional] 
**Errors** | Pointer to [**[]BatchProductsResponseErrorsInner**](BatchProductsResponseErrorsInner.md) | Error details (max 20) | [optional] 

## Methods

### NewBatchProductsResponse

`func NewBatchProductsResponse() *BatchProductsResponse`

NewBatchProductsResponse instantiates a new BatchProductsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchProductsResponseWithDefaults

`func NewBatchProductsResponseWithDefaults() *BatchProductsResponse`

NewBatchProductsResponseWithDefaults instantiates a new BatchProductsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemaVersion

`func (o *BatchProductsResponse) GetSchemaVersion() string`

GetSchemaVersion returns the SchemaVersion field if non-nil, zero value otherwise.

### GetSchemaVersionOk

`func (o *BatchProductsResponse) GetSchemaVersionOk() (*string, bool)`

GetSchemaVersionOk returns a tuple with the SchemaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaVersion

`func (o *BatchProductsResponse) SetSchemaVersion(v string)`

SetSchemaVersion sets SchemaVersion field to given value.

### HasSchemaVersion

`func (o *BatchProductsResponse) HasSchemaVersion() bool`

HasSchemaVersion returns a boolean if a field has been set.

### GetRequestId

`func (o *BatchProductsResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *BatchProductsResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *BatchProductsResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *BatchProductsResponse) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetCreated

`func (o *BatchProductsResponse) GetCreated() int32`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *BatchProductsResponse) GetCreatedOk() (*int32, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *BatchProductsResponse) SetCreated(v int32)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *BatchProductsResponse) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetUpdated

`func (o *BatchProductsResponse) GetUpdated() int32`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *BatchProductsResponse) GetUpdatedOk() (*int32, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *BatchProductsResponse) SetUpdated(v int32)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *BatchProductsResponse) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### GetErrored

`func (o *BatchProductsResponse) GetErrored() int32`

GetErrored returns the Errored field if non-nil, zero value otherwise.

### GetErroredOk

`func (o *BatchProductsResponse) GetErroredOk() (*int32, bool)`

GetErroredOk returns a tuple with the Errored field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrored

`func (o *BatchProductsResponse) SetErrored(v int32)`

SetErrored sets Errored field to given value.

### HasErrored

`func (o *BatchProductsResponse) HasErrored() bool`

HasErrored returns a boolean if a field has been set.

### GetErrors

`func (o *BatchProductsResponse) GetErrors() []BatchProductsResponseErrorsInner`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *BatchProductsResponse) GetErrorsOk() (*[]BatchProductsResponseErrorsInner, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *BatchProductsResponse) SetErrors(v []BatchProductsResponseErrorsInner)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *BatchProductsResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


