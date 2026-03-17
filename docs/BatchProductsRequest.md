# BatchProductsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Products** | [**[]BatchProductsRequestProductsInner**](BatchProductsRequestProductsInner.md) |  | 

## Methods

### NewBatchProductsRequest

`func NewBatchProductsRequest(products []BatchProductsRequestProductsInner, ) *BatchProductsRequest`

NewBatchProductsRequest instantiates a new BatchProductsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchProductsRequestWithDefaults

`func NewBatchProductsRequestWithDefaults() *BatchProductsRequest`

NewBatchProductsRequestWithDefaults instantiates a new BatchProductsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProducts

`func (o *BatchProductsRequest) GetProducts() []BatchProductsRequestProductsInner`

GetProducts returns the Products field if non-nil, zero value otherwise.

### GetProductsOk

`func (o *BatchProductsRequest) GetProductsOk() (*[]BatchProductsRequestProductsInner, bool)`

GetProductsOk returns a tuple with the Products field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducts

`func (o *BatchProductsRequest) SetProducts(v []BatchProductsRequestProductsInner)`

SetProducts sets Products field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


