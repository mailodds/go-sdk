# ProductFacets

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Categories** | Pointer to [**[]ProductFacetsCategoriesInner**](ProductFacetsCategoriesInner.md) |  | [optional] 
**PriceRanges** | Pointer to [**[]ProductFacetsPriceRangesInner**](ProductFacetsPriceRangesInner.md) |  | [optional] 
**Stores** | Pointer to [**[]ProductFacetsStoresInner**](ProductFacetsStoresInner.md) |  | [optional] 

## Methods

### NewProductFacets

`func NewProductFacets() *ProductFacets`

NewProductFacets instantiates a new ProductFacets object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductFacetsWithDefaults

`func NewProductFacetsWithDefaults() *ProductFacets`

NewProductFacetsWithDefaults instantiates a new ProductFacets object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategories

`func (o *ProductFacets) GetCategories() []ProductFacetsCategoriesInner`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *ProductFacets) GetCategoriesOk() (*[]ProductFacetsCategoriesInner, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *ProductFacets) SetCategories(v []ProductFacetsCategoriesInner)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *ProductFacets) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### GetPriceRanges

`func (o *ProductFacets) GetPriceRanges() []ProductFacetsPriceRangesInner`

GetPriceRanges returns the PriceRanges field if non-nil, zero value otherwise.

### GetPriceRangesOk

`func (o *ProductFacets) GetPriceRangesOk() (*[]ProductFacetsPriceRangesInner, bool)`

GetPriceRangesOk returns a tuple with the PriceRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceRanges

`func (o *ProductFacets) SetPriceRanges(v []ProductFacetsPriceRangesInner)`

SetPriceRanges sets PriceRanges field to given value.

### HasPriceRanges

`func (o *ProductFacets) HasPriceRanges() bool`

HasPriceRanges returns a boolean if a field has been set.

### GetStores

`func (o *ProductFacets) GetStores() []ProductFacetsStoresInner`

GetStores returns the Stores field if non-nil, zero value otherwise.

### GetStoresOk

`func (o *ProductFacets) GetStoresOk() (*[]ProductFacetsStoresInner, bool)`

GetStoresOk returns a tuple with the Stores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStores

`func (o *ProductFacets) SetStores(v []ProductFacetsStoresInner)`

SetStores sets Stores field to given value.

### HasStores

`func (o *ProductFacets) HasStores() bool`

HasStores returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


