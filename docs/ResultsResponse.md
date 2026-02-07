# ResultsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SchemaVersion** | Pointer to **string** |  | [optional] 
**Results** | Pointer to [**[]ValidationResult**](ValidationResult.md) |  | [optional] 
**Pagination** | Pointer to [**Pagination**](Pagination.md) |  | [optional] 

## Methods

### NewResultsResponse

`func NewResultsResponse() *ResultsResponse`

NewResultsResponse instantiates a new ResultsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResultsResponseWithDefaults

`func NewResultsResponseWithDefaults() *ResultsResponse`

NewResultsResponseWithDefaults instantiates a new ResultsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemaVersion

`func (o *ResultsResponse) GetSchemaVersion() string`

GetSchemaVersion returns the SchemaVersion field if non-nil, zero value otherwise.

### GetSchemaVersionOk

`func (o *ResultsResponse) GetSchemaVersionOk() (*string, bool)`

GetSchemaVersionOk returns a tuple with the SchemaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaVersion

`func (o *ResultsResponse) SetSchemaVersion(v string)`

SetSchemaVersion sets SchemaVersion field to given value.

### HasSchemaVersion

`func (o *ResultsResponse) HasSchemaVersion() bool`

HasSchemaVersion returns a boolean if a field has been set.

### GetResults

`func (o *ResultsResponse) GetResults() []ValidationResult`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *ResultsResponse) GetResultsOk() (*[]ValidationResult, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *ResultsResponse) SetResults(v []ValidationResult)`

SetResults sets Results field to given value.

### HasResults

`func (o *ResultsResponse) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetPagination

`func (o *ResultsResponse) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ResultsResponse) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ResultsResponse) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *ResultsResponse) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


