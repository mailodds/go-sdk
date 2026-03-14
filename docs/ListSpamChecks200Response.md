# ListSpamChecks200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SchemaVersion** | Pointer to **string** |  | [optional] 
**RequestId** | Pointer to **string** |  | [optional] 
**SpamChecks** | Pointer to [**[]SpamCheck**](SpamCheck.md) |  | [optional] 
**Pagination** | Pointer to [**Pagination**](Pagination.md) |  | [optional] 

## Methods

### NewListSpamChecks200Response

`func NewListSpamChecks200Response() *ListSpamChecks200Response`

NewListSpamChecks200Response instantiates a new ListSpamChecks200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListSpamChecks200ResponseWithDefaults

`func NewListSpamChecks200ResponseWithDefaults() *ListSpamChecks200Response`

NewListSpamChecks200ResponseWithDefaults instantiates a new ListSpamChecks200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemaVersion

`func (o *ListSpamChecks200Response) GetSchemaVersion() string`

GetSchemaVersion returns the SchemaVersion field if non-nil, zero value otherwise.

### GetSchemaVersionOk

`func (o *ListSpamChecks200Response) GetSchemaVersionOk() (*string, bool)`

GetSchemaVersionOk returns a tuple with the SchemaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaVersion

`func (o *ListSpamChecks200Response) SetSchemaVersion(v string)`

SetSchemaVersion sets SchemaVersion field to given value.

### HasSchemaVersion

`func (o *ListSpamChecks200Response) HasSchemaVersion() bool`

HasSchemaVersion returns a boolean if a field has been set.

### GetRequestId

`func (o *ListSpamChecks200Response) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *ListSpamChecks200Response) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *ListSpamChecks200Response) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *ListSpamChecks200Response) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetSpamChecks

`func (o *ListSpamChecks200Response) GetSpamChecks() []SpamCheck`

GetSpamChecks returns the SpamChecks field if non-nil, zero value otherwise.

### GetSpamChecksOk

`func (o *ListSpamChecks200Response) GetSpamChecksOk() (*[]SpamCheck, bool)`

GetSpamChecksOk returns a tuple with the SpamChecks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpamChecks

`func (o *ListSpamChecks200Response) SetSpamChecks(v []SpamCheck)`

SetSpamChecks sets SpamChecks field to given value.

### HasSpamChecks

`func (o *ListSpamChecks200Response) HasSpamChecks() bool`

HasSpamChecks returns a boolean if a field has been set.

### GetPagination

`func (o *ListSpamChecks200Response) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ListSpamChecks200Response) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ListSpamChecks200Response) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *ListSpamChecks200Response) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


