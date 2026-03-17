# \ProductsAPI

All URIs are relative to *https://api.mailodds.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BatchProducts**](ProductsAPI.md#BatchProducts) | **Post** /v1/stores/{store_id}/products/batch | Batch push products
[**BulkUpdateProducts**](ProductsAPI.md#BulkUpdateProducts) | **Patch** /v1/store-products/bulk | Bulk update products
[**GetProduct**](ProductsAPI.md#GetProduct) | **Get** /v1/store-products/{product_id} | Get a product
[**QueryProducts**](ProductsAPI.md#QueryProducts) | **Get** /v1/store-products | Query products



## BatchProducts

> BatchProductsResponse BatchProducts(ctx, storeId).BatchProductsRequest(batchProductsRequest).Execute()

Batch push products



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	storeId := "storeId_example" // string | Store connection UUID
	batchProductsRequest := *openapiclient.NewBatchProductsRequest([]openapiclient.BatchProductsRequestProductsInner{*openapiclient.NewBatchProductsRequestProductsInner("ExternalId_example", "Title_example", "ProductUrl_example")}) // BatchProductsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductsAPI.BatchProducts(context.Background(), storeId).BatchProductsRequest(batchProductsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductsAPI.BatchProducts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BatchProducts`: BatchProductsResponse
	fmt.Fprintf(os.Stdout, "Response from `ProductsAPI.BatchProducts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeId** | **string** | Store connection UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiBatchProductsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **batchProductsRequest** | [**BatchProductsRequest**](BatchProductsRequest.md) |  | 

### Return type

[**BatchProductsResponse**](BatchProductsResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BulkUpdateProducts

> BulkUpdateProducts200Response BulkUpdateProducts(ctx).BulkUpdateProductsRequest(bulkUpdateProductsRequest).Execute()

Bulk update products



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	bulkUpdateProductsRequest := *openapiclient.NewBulkUpdateProductsRequest([]string{"ProductIds_example"}, false) // BulkUpdateProductsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductsAPI.BulkUpdateProducts(context.Background()).BulkUpdateProductsRequest(bulkUpdateProductsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductsAPI.BulkUpdateProducts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BulkUpdateProducts`: BulkUpdateProducts200Response
	fmt.Fprintf(os.Stdout, "Response from `ProductsAPI.BulkUpdateProducts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkUpdateProductsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bulkUpdateProductsRequest** | [**BulkUpdateProductsRequest**](BulkUpdateProductsRequest.md) |  | 

### Return type

[**BulkUpdateProducts200Response**](BulkUpdateProducts200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProduct

> GetProduct200Response GetProduct(ctx, productId).Execute()

Get a product



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	productId := "productId_example" // string | Product UUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductsAPI.GetProduct(context.Background(), productId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductsAPI.GetProduct``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProduct`: GetProduct200Response
	fmt.Fprintf(os.Stdout, "Response from `ProductsAPI.GetProduct`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**productId** | **string** | Product UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProductRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetProduct200Response**](GetProduct200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QueryProducts

> QueryProducts200Response QueryProducts(ctx).StoreId(storeId).Category(category).StockStatus(stockStatus).OnSale(onSale).Search(search).Facets(facets).GroupBySku(groupBySku).Page(page).PerPage(perPage).Execute()

Query products



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	storeId := "storeId_example" // string | Filter by store connection UUID (optional)
	category := "category_example" // string | Filter by category name (optional)
	stockStatus := "stockStatus_example" // string | Filter by stock status (optional)
	onSale := true // bool | Filter to products currently on sale (optional)
	search := "search_example" // string | Search by title or SKU (optional)
	facets := true // bool | Include facet aggregations (categories, price ranges, stores) (optional) (default to false)
	groupBySku := true // bool | Merge products with same SKU across stores into unified entries (optional) (default to false)
	page := int32(56) // int32 |  (optional) (default to 1)
	perPage := int32(56) // int32 |  (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductsAPI.QueryProducts(context.Background()).StoreId(storeId).Category(category).StockStatus(stockStatus).OnSale(onSale).Search(search).Facets(facets).GroupBySku(groupBySku).Page(page).PerPage(perPage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductsAPI.QueryProducts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QueryProducts`: QueryProducts200Response
	fmt.Fprintf(os.Stdout, "Response from `ProductsAPI.QueryProducts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQueryProductsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **storeId** | **string** | Filter by store connection UUID | 
 **category** | **string** | Filter by category name | 
 **stockStatus** | **string** | Filter by stock status | 
 **onSale** | **bool** | Filter to products currently on sale | 
 **search** | **string** | Search by title or SKU | 
 **facets** | **bool** | Include facet aggregations (categories, price ranges, stores) | [default to false]
 **groupBySku** | **bool** | Merge products with same SKU across stores into unified entries | [default to false]
 **page** | **int32** |  | [default to 1]
 **perPage** | **int32** |  | [default to 20]

### Return type

[**QueryProducts200Response**](QueryProducts200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

