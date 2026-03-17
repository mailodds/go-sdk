# BulkUpdateProductsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProductIds** | **[]string** | Product IDs to update | 
**IsActive** | **bool** | Set product visibility | 

## Methods

### NewBulkUpdateProductsRequest

`func NewBulkUpdateProductsRequest(productIds []string, isActive bool, ) *BulkUpdateProductsRequest`

NewBulkUpdateProductsRequest instantiates a new BulkUpdateProductsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkUpdateProductsRequestWithDefaults

`func NewBulkUpdateProductsRequestWithDefaults() *BulkUpdateProductsRequest`

NewBulkUpdateProductsRequestWithDefaults instantiates a new BulkUpdateProductsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProductIds

`func (o *BulkUpdateProductsRequest) GetProductIds() []string`

GetProductIds returns the ProductIds field if non-nil, zero value otherwise.

### GetProductIdsOk

`func (o *BulkUpdateProductsRequest) GetProductIdsOk() (*[]string, bool)`

GetProductIdsOk returns a tuple with the ProductIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductIds

`func (o *BulkUpdateProductsRequest) SetProductIds(v []string)`

SetProductIds sets ProductIds field to given value.


### GetIsActive

`func (o *BulkUpdateProductsRequest) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *BulkUpdateProductsRequest) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *BulkUpdateProductsRequest) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


