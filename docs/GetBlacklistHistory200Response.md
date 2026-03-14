# GetBlacklistHistory200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SchemaVersion** | Pointer to **string** |  | [optional] 
**RequestId** | Pointer to **string** |  | [optional] 
**Checks** | Pointer to [**[]GetBlacklistHistory200ResponseChecksInner**](GetBlacklistHistory200ResponseChecksInner.md) |  | [optional] 
**Pagination** | Pointer to [**Pagination**](Pagination.md) |  | [optional] 

## Methods

### NewGetBlacklistHistory200Response

`func NewGetBlacklistHistory200Response() *GetBlacklistHistory200Response`

NewGetBlacklistHistory200Response instantiates a new GetBlacklistHistory200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetBlacklistHistory200ResponseWithDefaults

`func NewGetBlacklistHistory200ResponseWithDefaults() *GetBlacklistHistory200Response`

NewGetBlacklistHistory200ResponseWithDefaults instantiates a new GetBlacklistHistory200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemaVersion

`func (o *GetBlacklistHistory200Response) GetSchemaVersion() string`

GetSchemaVersion returns the SchemaVersion field if non-nil, zero value otherwise.

### GetSchemaVersionOk

`func (o *GetBlacklistHistory200Response) GetSchemaVersionOk() (*string, bool)`

GetSchemaVersionOk returns a tuple with the SchemaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaVersion

`func (o *GetBlacklistHistory200Response) SetSchemaVersion(v string)`

SetSchemaVersion sets SchemaVersion field to given value.

### HasSchemaVersion

`func (o *GetBlacklistHistory200Response) HasSchemaVersion() bool`

HasSchemaVersion returns a boolean if a field has been set.

### GetRequestId

`func (o *GetBlacklistHistory200Response) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *GetBlacklistHistory200Response) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *GetBlacklistHistory200Response) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *GetBlacklistHistory200Response) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetChecks

`func (o *GetBlacklistHistory200Response) GetChecks() []GetBlacklistHistory200ResponseChecksInner`

GetChecks returns the Checks field if non-nil, zero value otherwise.

### GetChecksOk

`func (o *GetBlacklistHistory200Response) GetChecksOk() (*[]GetBlacklistHistory200ResponseChecksInner, bool)`

GetChecksOk returns a tuple with the Checks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecks

`func (o *GetBlacklistHistory200Response) SetChecks(v []GetBlacklistHistory200ResponseChecksInner)`

SetChecks sets Checks field to given value.

### HasChecks

`func (o *GetBlacklistHistory200Response) HasChecks() bool`

HasChecks returns a boolean if a field has been set.

### GetPagination

`func (o *GetBlacklistHistory200Response) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *GetBlacklistHistory200Response) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *GetBlacklistHistory200Response) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *GetBlacklistHistory200Response) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


