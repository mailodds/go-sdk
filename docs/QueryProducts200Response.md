# QueryProducts200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SchemaVersion** | Pointer to **string** |  | [optional] 
**RequestId** | Pointer to **string** |  | [optional] 
**Products** | Pointer to [**[]StoreProduct**](StoreProduct.md) |  | [optional] 
**Total** | Pointer to **int32** |  | [optional] 
**Page** | Pointer to **int32** |  | [optional] 
**PerPage** | Pointer to **int32** |  | [optional] 
**Pages** | Pointer to **int32** |  | [optional] 
**Facets** | Pointer to [**ProductFacets**](ProductFacets.md) |  | [optional] 

## Methods

### NewQueryProducts200Response

`func NewQueryProducts200Response() *QueryProducts200Response`

NewQueryProducts200Response instantiates a new QueryProducts200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryProducts200ResponseWithDefaults

`func NewQueryProducts200ResponseWithDefaults() *QueryProducts200Response`

NewQueryProducts200ResponseWithDefaults instantiates a new QueryProducts200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemaVersion

`func (o *QueryProducts200Response) GetSchemaVersion() string`

GetSchemaVersion returns the SchemaVersion field if non-nil, zero value otherwise.

### GetSchemaVersionOk

`func (o *QueryProducts200Response) GetSchemaVersionOk() (*string, bool)`

GetSchemaVersionOk returns a tuple with the SchemaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaVersion

`func (o *QueryProducts200Response) SetSchemaVersion(v string)`

SetSchemaVersion sets SchemaVersion field to given value.

### HasSchemaVersion

`func (o *QueryProducts200Response) HasSchemaVersion() bool`

HasSchemaVersion returns a boolean if a field has been set.

### GetRequestId

`func (o *QueryProducts200Response) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *QueryProducts200Response) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *QueryProducts200Response) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *QueryProducts200Response) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetProducts

`func (o *QueryProducts200Response) GetProducts() []StoreProduct`

GetProducts returns the Products field if non-nil, zero value otherwise.

### GetProductsOk

`func (o *QueryProducts200Response) GetProductsOk() (*[]StoreProduct, bool)`

GetProductsOk returns a tuple with the Products field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducts

`func (o *QueryProducts200Response) SetProducts(v []StoreProduct)`

SetProducts sets Products field to given value.

### HasProducts

`func (o *QueryProducts200Response) HasProducts() bool`

HasProducts returns a boolean if a field has been set.

### GetTotal

`func (o *QueryProducts200Response) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *QueryProducts200Response) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *QueryProducts200Response) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *QueryProducts200Response) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetPage

`func (o *QueryProducts200Response) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *QueryProducts200Response) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *QueryProducts200Response) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *QueryProducts200Response) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetPerPage

`func (o *QueryProducts200Response) GetPerPage() int32`

GetPerPage returns the PerPage field if non-nil, zero value otherwise.

### GetPerPageOk

`func (o *QueryProducts200Response) GetPerPageOk() (*int32, bool)`

GetPerPageOk returns a tuple with the PerPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerPage

`func (o *QueryProducts200Response) SetPerPage(v int32)`

SetPerPage sets PerPage field to given value.

### HasPerPage

`func (o *QueryProducts200Response) HasPerPage() bool`

HasPerPage returns a boolean if a field has been set.

### GetPages

`func (o *QueryProducts200Response) GetPages() int32`

GetPages returns the Pages field if non-nil, zero value otherwise.

### GetPagesOk

`func (o *QueryProducts200Response) GetPagesOk() (*int32, bool)`

GetPagesOk returns a tuple with the Pages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPages

`func (o *QueryProducts200Response) SetPages(v int32)`

SetPages sets Pages field to given value.

### HasPages

`func (o *QueryProducts200Response) HasPages() bool`

HasPages returns a boolean if a field has been set.

### GetFacets

`func (o *QueryProducts200Response) GetFacets() ProductFacets`

GetFacets returns the Facets field if non-nil, zero value otherwise.

### GetFacetsOk

`func (o *QueryProducts200Response) GetFacetsOk() (*ProductFacets, bool)`

GetFacetsOk returns a tuple with the Facets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacets

`func (o *QueryProducts200Response) SetFacets(v ProductFacets)`

SetFacets sets Facets field to given value.

### HasFacets

`func (o *QueryProducts200Response) HasFacets() bool`

HasFacets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


