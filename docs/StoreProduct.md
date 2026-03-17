# StoreProduct

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Product UUID | [optional] 
**StoreId** | Pointer to **string** | Store connection UUID | [optional] 
**ExternalId** | Pointer to **string** | Product ID in the source store | [optional] 
**Sku** | Pointer to **NullableString** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**PriceCurrent** | Pointer to **NullableFloat32** | Current price | [optional] 
**PriceOriginal** | Pointer to **NullableFloat32** | Original price (before discount) | [optional] 
**Currency** | Pointer to **string** |  | [optional] [default to "USD"]
**SaleStart** | Pointer to **NullableTime** |  | [optional] 
**SaleEnd** | Pointer to **NullableTime** |  | [optional] 
**StockStatus** | Pointer to **string** |  | [optional] 
**StockQuantity** | Pointer to **NullableInt32** |  | [optional] 
**ImageUrl** | Pointer to **NullableString** |  | [optional] 
**AdditionalImages** | Pointer to **[]string** |  | [optional] 
**Categories** | Pointer to **[]string** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**ProductUrl** | Pointer to **string** |  | [optional] 
**Variants** | Pointer to **[]map[string]interface{}** |  | [optional] 
**IsActive** | Pointer to **bool** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewStoreProduct

`func NewStoreProduct() *StoreProduct`

NewStoreProduct instantiates a new StoreProduct object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoreProductWithDefaults

`func NewStoreProductWithDefaults() *StoreProduct`

NewStoreProductWithDefaults instantiates a new StoreProduct object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *StoreProduct) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StoreProduct) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StoreProduct) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *StoreProduct) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStoreId

`func (o *StoreProduct) GetStoreId() string`

GetStoreId returns the StoreId field if non-nil, zero value otherwise.

### GetStoreIdOk

`func (o *StoreProduct) GetStoreIdOk() (*string, bool)`

GetStoreIdOk returns a tuple with the StoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreId

`func (o *StoreProduct) SetStoreId(v string)`

SetStoreId sets StoreId field to given value.

### HasStoreId

`func (o *StoreProduct) HasStoreId() bool`

HasStoreId returns a boolean if a field has been set.

### GetExternalId

`func (o *StoreProduct) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *StoreProduct) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *StoreProduct) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *StoreProduct) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetSku

`func (o *StoreProduct) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *StoreProduct) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *StoreProduct) SetSku(v string)`

SetSku sets Sku field to given value.

### HasSku

`func (o *StoreProduct) HasSku() bool`

HasSku returns a boolean if a field has been set.

### SetSkuNil

`func (o *StoreProduct) SetSkuNil(b bool)`

 SetSkuNil sets the value for Sku to be an explicit nil

### UnsetSku
`func (o *StoreProduct) UnsetSku()`

UnsetSku ensures that no value is present for Sku, not even an explicit nil
### GetTitle

`func (o *StoreProduct) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *StoreProduct) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *StoreProduct) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *StoreProduct) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDescription

`func (o *StoreProduct) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *StoreProduct) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *StoreProduct) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *StoreProduct) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *StoreProduct) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *StoreProduct) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetPriceCurrent

`func (o *StoreProduct) GetPriceCurrent() float32`

GetPriceCurrent returns the PriceCurrent field if non-nil, zero value otherwise.

### GetPriceCurrentOk

`func (o *StoreProduct) GetPriceCurrentOk() (*float32, bool)`

GetPriceCurrentOk returns a tuple with the PriceCurrent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceCurrent

`func (o *StoreProduct) SetPriceCurrent(v float32)`

SetPriceCurrent sets PriceCurrent field to given value.

### HasPriceCurrent

`func (o *StoreProduct) HasPriceCurrent() bool`

HasPriceCurrent returns a boolean if a field has been set.

### SetPriceCurrentNil

`func (o *StoreProduct) SetPriceCurrentNil(b bool)`

 SetPriceCurrentNil sets the value for PriceCurrent to be an explicit nil

### UnsetPriceCurrent
`func (o *StoreProduct) UnsetPriceCurrent()`

UnsetPriceCurrent ensures that no value is present for PriceCurrent, not even an explicit nil
### GetPriceOriginal

`func (o *StoreProduct) GetPriceOriginal() float32`

GetPriceOriginal returns the PriceOriginal field if non-nil, zero value otherwise.

### GetPriceOriginalOk

`func (o *StoreProduct) GetPriceOriginalOk() (*float32, bool)`

GetPriceOriginalOk returns a tuple with the PriceOriginal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceOriginal

`func (o *StoreProduct) SetPriceOriginal(v float32)`

SetPriceOriginal sets PriceOriginal field to given value.

### HasPriceOriginal

`func (o *StoreProduct) HasPriceOriginal() bool`

HasPriceOriginal returns a boolean if a field has been set.

### SetPriceOriginalNil

`func (o *StoreProduct) SetPriceOriginalNil(b bool)`

 SetPriceOriginalNil sets the value for PriceOriginal to be an explicit nil

### UnsetPriceOriginal
`func (o *StoreProduct) UnsetPriceOriginal()`

UnsetPriceOriginal ensures that no value is present for PriceOriginal, not even an explicit nil
### GetCurrency

`func (o *StoreProduct) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *StoreProduct) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *StoreProduct) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *StoreProduct) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetSaleStart

`func (o *StoreProduct) GetSaleStart() time.Time`

GetSaleStart returns the SaleStart field if non-nil, zero value otherwise.

### GetSaleStartOk

`func (o *StoreProduct) GetSaleStartOk() (*time.Time, bool)`

GetSaleStartOk returns a tuple with the SaleStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaleStart

`func (o *StoreProduct) SetSaleStart(v time.Time)`

SetSaleStart sets SaleStart field to given value.

### HasSaleStart

`func (o *StoreProduct) HasSaleStart() bool`

HasSaleStart returns a boolean if a field has been set.

### SetSaleStartNil

`func (o *StoreProduct) SetSaleStartNil(b bool)`

 SetSaleStartNil sets the value for SaleStart to be an explicit nil

### UnsetSaleStart
`func (o *StoreProduct) UnsetSaleStart()`

UnsetSaleStart ensures that no value is present for SaleStart, not even an explicit nil
### GetSaleEnd

`func (o *StoreProduct) GetSaleEnd() time.Time`

GetSaleEnd returns the SaleEnd field if non-nil, zero value otherwise.

### GetSaleEndOk

`func (o *StoreProduct) GetSaleEndOk() (*time.Time, bool)`

GetSaleEndOk returns a tuple with the SaleEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaleEnd

`func (o *StoreProduct) SetSaleEnd(v time.Time)`

SetSaleEnd sets SaleEnd field to given value.

### HasSaleEnd

`func (o *StoreProduct) HasSaleEnd() bool`

HasSaleEnd returns a boolean if a field has been set.

### SetSaleEndNil

`func (o *StoreProduct) SetSaleEndNil(b bool)`

 SetSaleEndNil sets the value for SaleEnd to be an explicit nil

### UnsetSaleEnd
`func (o *StoreProduct) UnsetSaleEnd()`

UnsetSaleEnd ensures that no value is present for SaleEnd, not even an explicit nil
### GetStockStatus

`func (o *StoreProduct) GetStockStatus() string`

GetStockStatus returns the StockStatus field if non-nil, zero value otherwise.

### GetStockStatusOk

`func (o *StoreProduct) GetStockStatusOk() (*string, bool)`

GetStockStatusOk returns a tuple with the StockStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockStatus

`func (o *StoreProduct) SetStockStatus(v string)`

SetStockStatus sets StockStatus field to given value.

### HasStockStatus

`func (o *StoreProduct) HasStockStatus() bool`

HasStockStatus returns a boolean if a field has been set.

### GetStockQuantity

`func (o *StoreProduct) GetStockQuantity() int32`

GetStockQuantity returns the StockQuantity field if non-nil, zero value otherwise.

### GetStockQuantityOk

`func (o *StoreProduct) GetStockQuantityOk() (*int32, bool)`

GetStockQuantityOk returns a tuple with the StockQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockQuantity

`func (o *StoreProduct) SetStockQuantity(v int32)`

SetStockQuantity sets StockQuantity field to given value.

### HasStockQuantity

`func (o *StoreProduct) HasStockQuantity() bool`

HasStockQuantity returns a boolean if a field has been set.

### SetStockQuantityNil

`func (o *StoreProduct) SetStockQuantityNil(b bool)`

 SetStockQuantityNil sets the value for StockQuantity to be an explicit nil

### UnsetStockQuantity
`func (o *StoreProduct) UnsetStockQuantity()`

UnsetStockQuantity ensures that no value is present for StockQuantity, not even an explicit nil
### GetImageUrl

`func (o *StoreProduct) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *StoreProduct) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *StoreProduct) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *StoreProduct) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.

### SetImageUrlNil

`func (o *StoreProduct) SetImageUrlNil(b bool)`

 SetImageUrlNil sets the value for ImageUrl to be an explicit nil

### UnsetImageUrl
`func (o *StoreProduct) UnsetImageUrl()`

UnsetImageUrl ensures that no value is present for ImageUrl, not even an explicit nil
### GetAdditionalImages

`func (o *StoreProduct) GetAdditionalImages() []string`

GetAdditionalImages returns the AdditionalImages field if non-nil, zero value otherwise.

### GetAdditionalImagesOk

`func (o *StoreProduct) GetAdditionalImagesOk() (*[]string, bool)`

GetAdditionalImagesOk returns a tuple with the AdditionalImages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalImages

`func (o *StoreProduct) SetAdditionalImages(v []string)`

SetAdditionalImages sets AdditionalImages field to given value.

### HasAdditionalImages

`func (o *StoreProduct) HasAdditionalImages() bool`

HasAdditionalImages returns a boolean if a field has been set.

### SetAdditionalImagesNil

`func (o *StoreProduct) SetAdditionalImagesNil(b bool)`

 SetAdditionalImagesNil sets the value for AdditionalImages to be an explicit nil

### UnsetAdditionalImages
`func (o *StoreProduct) UnsetAdditionalImages()`

UnsetAdditionalImages ensures that no value is present for AdditionalImages, not even an explicit nil
### GetCategories

`func (o *StoreProduct) GetCategories() []string`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *StoreProduct) GetCategoriesOk() (*[]string, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *StoreProduct) SetCategories(v []string)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *StoreProduct) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### SetCategoriesNil

`func (o *StoreProduct) SetCategoriesNil(b bool)`

 SetCategoriesNil sets the value for Categories to be an explicit nil

### UnsetCategories
`func (o *StoreProduct) UnsetCategories()`

UnsetCategories ensures that no value is present for Categories, not even an explicit nil
### GetTags

`func (o *StoreProduct) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *StoreProduct) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *StoreProduct) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *StoreProduct) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *StoreProduct) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *StoreProduct) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetProductUrl

`func (o *StoreProduct) GetProductUrl() string`

GetProductUrl returns the ProductUrl field if non-nil, zero value otherwise.

### GetProductUrlOk

`func (o *StoreProduct) GetProductUrlOk() (*string, bool)`

GetProductUrlOk returns a tuple with the ProductUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductUrl

`func (o *StoreProduct) SetProductUrl(v string)`

SetProductUrl sets ProductUrl field to given value.

### HasProductUrl

`func (o *StoreProduct) HasProductUrl() bool`

HasProductUrl returns a boolean if a field has been set.

### GetVariants

`func (o *StoreProduct) GetVariants() []map[string]interface{}`

GetVariants returns the Variants field if non-nil, zero value otherwise.

### GetVariantsOk

`func (o *StoreProduct) GetVariantsOk() (*[]map[string]interface{}, bool)`

GetVariantsOk returns a tuple with the Variants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariants

`func (o *StoreProduct) SetVariants(v []map[string]interface{})`

SetVariants sets Variants field to given value.

### HasVariants

`func (o *StoreProduct) HasVariants() bool`

HasVariants returns a boolean if a field has been set.

### SetVariantsNil

`func (o *StoreProduct) SetVariantsNil(b bool)`

 SetVariantsNil sets the value for Variants to be an explicit nil

### UnsetVariants
`func (o *StoreProduct) UnsetVariants()`

UnsetVariants ensures that no value is present for Variants, not even an explicit nil
### GetIsActive

`func (o *StoreProduct) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *StoreProduct) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *StoreProduct) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *StoreProduct) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetCreatedAt

`func (o *StoreProduct) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *StoreProduct) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *StoreProduct) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *StoreProduct) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *StoreProduct) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *StoreProduct) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *StoreProduct) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *StoreProduct) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


