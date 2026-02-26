# ValidateBatch200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SchemaVersion** | Pointer to **string** |  | [optional] 
**RequestId** | Pointer to **string** | Unique request identifier | [optional] 
**Total** | Pointer to **int32** |  | [optional] 
**Summary** | Pointer to [**ValidateBatch200ResponseSummary**](ValidateBatch200ResponseSummary.md) |  | [optional] 
**Results** | Pointer to [**[]ValidationResponse**](ValidationResponse.md) |  | [optional] 

## Methods

### NewValidateBatch200Response

`func NewValidateBatch200Response() *ValidateBatch200Response`

NewValidateBatch200Response instantiates a new ValidateBatch200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidateBatch200ResponseWithDefaults

`func NewValidateBatch200ResponseWithDefaults() *ValidateBatch200Response`

NewValidateBatch200ResponseWithDefaults instantiates a new ValidateBatch200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemaVersion

`func (o *ValidateBatch200Response) GetSchemaVersion() string`

GetSchemaVersion returns the SchemaVersion field if non-nil, zero value otherwise.

### GetSchemaVersionOk

`func (o *ValidateBatch200Response) GetSchemaVersionOk() (*string, bool)`

GetSchemaVersionOk returns a tuple with the SchemaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaVersion

`func (o *ValidateBatch200Response) SetSchemaVersion(v string)`

SetSchemaVersion sets SchemaVersion field to given value.

### HasSchemaVersion

`func (o *ValidateBatch200Response) HasSchemaVersion() bool`

HasSchemaVersion returns a boolean if a field has been set.

### GetRequestId

`func (o *ValidateBatch200Response) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *ValidateBatch200Response) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *ValidateBatch200Response) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *ValidateBatch200Response) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetTotal

`func (o *ValidateBatch200Response) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ValidateBatch200Response) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ValidateBatch200Response) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *ValidateBatch200Response) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetSummary

`func (o *ValidateBatch200Response) GetSummary() ValidateBatch200ResponseSummary`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *ValidateBatch200Response) GetSummaryOk() (*ValidateBatch200ResponseSummary, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *ValidateBatch200Response) SetSummary(v ValidateBatch200ResponseSummary)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *ValidateBatch200Response) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetResults

`func (o *ValidateBatch200Response) GetResults() []ValidationResponse`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *ValidateBatch200Response) GetResultsOk() (*[]ValidationResponse, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *ValidateBatch200Response) SetResults(v []ValidationResponse)`

SetResults sets Results field to given value.

### HasResults

`func (o *ValidateBatch200Response) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


