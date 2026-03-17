# BatchProductsRequestProductsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalId** | **string** |  | 
**Title** | **string** |  | 
**ProductUrl** | **string** |  | 
**Sku** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**PriceCurrent** | Pointer to **float32** |  | [optional] 
**PriceOriginal** | Pointer to **float32** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**StockStatus** | Pointer to **string** |  | [optional] 
**StockQuantity** | Pointer to **int32** |  | [optional] 
**ImageUrl** | Pointer to **string** |  | [optional] 
**AdditionalImages** | Pointer to **[]string** |  | [optional] 
**Categories** | Pointer to **[]string** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**Variants** | Pointer to **[]map[string]interface{}** |  | [optional] 

## Methods

### NewBatchProductsRequestProductsInner

`func NewBatchProductsRequestProductsInner(externalId string, title string, productUrl string, ) *BatchProductsRequestProductsInner`

NewBatchProductsRequestProductsInner instantiates a new BatchProductsRequestProductsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchProductsRequestProductsInnerWithDefaults

`func NewBatchProductsRequestProductsInnerWithDefaults() *BatchProductsRequestProductsInner`

NewBatchProductsRequestProductsInnerWithDefaults instantiates a new BatchProductsRequestProductsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalId

`func (o *BatchProductsRequestProductsInner) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *BatchProductsRequestProductsInner) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *BatchProductsRequestProductsInner) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.


### GetTitle

`func (o *BatchProductsRequestProductsInner) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *BatchProductsRequestProductsInner) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *BatchProductsRequestProductsInner) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetProductUrl

`func (o *BatchProductsRequestProductsInner) GetProductUrl() string`

GetProductUrl returns the ProductUrl field if non-nil, zero value otherwise.

### GetProductUrlOk

`func (o *BatchProductsRequestProductsInner) GetProductUrlOk() (*string, bool)`

GetProductUrlOk returns a tuple with the ProductUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductUrl

`func (o *BatchProductsRequestProductsInner) SetProductUrl(v string)`

SetProductUrl sets ProductUrl field to given value.


### GetSku

`func (o *BatchProductsRequestProductsInner) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *BatchProductsRequestProductsInner) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *BatchProductsRequestProductsInner) SetSku(v string)`

SetSku sets Sku field to given value.

### HasSku

`func (o *BatchProductsRequestProductsInner) HasSku() bool`

HasSku returns a boolean if a field has been set.

### GetDescription

`func (o *BatchProductsRequestProductsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BatchProductsRequestProductsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BatchProductsRequestProductsInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BatchProductsRequestProductsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPriceCurrent

`func (o *BatchProductsRequestProductsInner) GetPriceCurrent() float32`

GetPriceCurrent returns the PriceCurrent field if non-nil, zero value otherwise.

### GetPriceCurrentOk

`func (o *BatchProductsRequestProductsInner) GetPriceCurrentOk() (*float32, bool)`

GetPriceCurrentOk returns a tuple with the PriceCurrent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceCurrent

`func (o *BatchProductsRequestProductsInner) SetPriceCurrent(v float32)`

SetPriceCurrent sets PriceCurrent field to given value.

### HasPriceCurrent

`func (o *BatchProductsRequestProductsInner) HasPriceCurrent() bool`

HasPriceCurrent returns a boolean if a field has been set.

### GetPriceOriginal

`func (o *BatchProductsRequestProductsInner) GetPriceOriginal() float32`

GetPriceOriginal returns the PriceOriginal field if non-nil, zero value otherwise.

### GetPriceOriginalOk

`func (o *BatchProductsRequestProductsInner) GetPriceOriginalOk() (*float32, bool)`

GetPriceOriginalOk returns a tuple with the PriceOriginal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceOriginal

`func (o *BatchProductsRequestProductsInner) SetPriceOriginal(v float32)`

SetPriceOriginal sets PriceOriginal field to given value.

### HasPriceOriginal

`func (o *BatchProductsRequestProductsInner) HasPriceOriginal() bool`

HasPriceOriginal returns a boolean if a field has been set.

### GetCurrency

`func (o *BatchProductsRequestProductsInner) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *BatchProductsRequestProductsInner) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *BatchProductsRequestProductsInner) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *BatchProductsRequestProductsInner) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetStockStatus

`func (o *BatchProductsRequestProductsInner) GetStockStatus() string`

GetStockStatus returns the StockStatus field if non-nil, zero value otherwise.

### GetStockStatusOk

`func (o *BatchProductsRequestProductsInner) GetStockStatusOk() (*string, bool)`

GetStockStatusOk returns a tuple with the StockStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockStatus

`func (o *BatchProductsRequestProductsInner) SetStockStatus(v string)`

SetStockStatus sets StockStatus field to given value.

### HasStockStatus

`func (o *BatchProductsRequestProductsInner) HasStockStatus() bool`

HasStockStatus returns a boolean if a field has been set.

### GetStockQuantity

`func (o *BatchProductsRequestProductsInner) GetStockQuantity() int32`

GetStockQuantity returns the StockQuantity field if non-nil, zero value otherwise.

### GetStockQuantityOk

`func (o *BatchProductsRequestProductsInner) GetStockQuantityOk() (*int32, bool)`

GetStockQuantityOk returns a tuple with the StockQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockQuantity

`func (o *BatchProductsRequestProductsInner) SetStockQuantity(v int32)`

SetStockQuantity sets StockQuantity field to given value.

### HasStockQuantity

`func (o *BatchProductsRequestProductsInner) HasStockQuantity() bool`

HasStockQuantity returns a boolean if a field has been set.

### GetImageUrl

`func (o *BatchProductsRequestProductsInner) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *BatchProductsRequestProductsInner) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *BatchProductsRequestProductsInner) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *BatchProductsRequestProductsInner) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.

### GetAdditionalImages

`func (o *BatchProductsRequestProductsInner) GetAdditionalImages() []string`

GetAdditionalImages returns the AdditionalImages field if non-nil, zero value otherwise.

### GetAdditionalImagesOk

`func (o *BatchProductsRequestProductsInner) GetAdditionalImagesOk() (*[]string, bool)`

GetAdditionalImagesOk returns a tuple with the AdditionalImages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalImages

`func (o *BatchProductsRequestProductsInner) SetAdditionalImages(v []string)`

SetAdditionalImages sets AdditionalImages field to given value.

### HasAdditionalImages

`func (o *BatchProductsRequestProductsInner) HasAdditionalImages() bool`

HasAdditionalImages returns a boolean if a field has been set.

### GetCategories

`func (o *BatchProductsRequestProductsInner) GetCategories() []string`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *BatchProductsRequestProductsInner) GetCategoriesOk() (*[]string, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *BatchProductsRequestProductsInner) SetCategories(v []string)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *BatchProductsRequestProductsInner) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### GetTags

`func (o *BatchProductsRequestProductsInner) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *BatchProductsRequestProductsInner) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *BatchProductsRequestProductsInner) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *BatchProductsRequestProductsInner) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVariants

`func (o *BatchProductsRequestProductsInner) GetVariants() []map[string]interface{}`

GetVariants returns the Variants field if non-nil, zero value otherwise.

### GetVariantsOk

`func (o *BatchProductsRequestProductsInner) GetVariantsOk() (*[]map[string]interface{}, bool)`

GetVariantsOk returns a tuple with the Variants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariants

`func (o *BatchProductsRequestProductsInner) SetVariants(v []map[string]interface{})`

SetVariants sets Variants field to given value.

### HasVariants

`func (o *BatchProductsRequestProductsInner) HasVariants() bool`

HasVariants returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


