# GetProduct200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SchemaVersion** | Pointer to **string** |  | [optional] 
**RequestId** | Pointer to **string** |  | [optional] 
**Product** | Pointer to [**StoreProduct**](StoreProduct.md) |  | [optional] 

## Methods

### NewGetProduct200Response

`func NewGetProduct200Response() *GetProduct200Response`

NewGetProduct200Response instantiates a new GetProduct200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetProduct200ResponseWithDefaults

`func NewGetProduct200ResponseWithDefaults() *GetProduct200Response`

NewGetProduct200ResponseWithDefaults instantiates a new GetProduct200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemaVersion

`func (o *GetProduct200Response) GetSchemaVersion() string`

GetSchemaVersion returns the SchemaVersion field if non-nil, zero value otherwise.

### GetSchemaVersionOk

`func (o *GetProduct200Response) GetSchemaVersionOk() (*string, bool)`

GetSchemaVersionOk returns a tuple with the SchemaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaVersion

`func (o *GetProduct200Response) SetSchemaVersion(v string)`

SetSchemaVersion sets SchemaVersion field to given value.

### HasSchemaVersion

`func (o *GetProduct200Response) HasSchemaVersion() bool`

HasSchemaVersion returns a boolean if a field has been set.

### GetRequestId

`func (o *GetProduct200Response) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *GetProduct200Response) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *GetProduct200Response) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *GetProduct200Response) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetProduct

`func (o *GetProduct200Response) GetProduct() StoreProduct`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *GetProduct200Response) GetProductOk() (*StoreProduct, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *GetProduct200Response) SetProduct(v StoreProduct)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *GetProduct200Response) HasProduct() bool`

HasProduct returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


