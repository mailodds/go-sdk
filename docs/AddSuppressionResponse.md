# AddSuppressionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SchemaVersion** | Pointer to **string** |  | [optional] 
**RequestId** | Pointer to **string** | Unique request identifier | [optional] 
**Added** | Pointer to **int32** | Number of entries successfully added | [optional] 
**Duplicates** | Pointer to **int32** | Number of duplicate entries skipped | [optional] 
**Invalid** | Pointer to **int32** | Number of invalid entries rejected | [optional] 
**Total** | Pointer to **int32** | Total entries in the request | [optional] 

## Methods

### NewAddSuppressionResponse

`func NewAddSuppressionResponse() *AddSuppressionResponse`

NewAddSuppressionResponse instantiates a new AddSuppressionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddSuppressionResponseWithDefaults

`func NewAddSuppressionResponseWithDefaults() *AddSuppressionResponse`

NewAddSuppressionResponseWithDefaults instantiates a new AddSuppressionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemaVersion

`func (o *AddSuppressionResponse) GetSchemaVersion() string`

GetSchemaVersion returns the SchemaVersion field if non-nil, zero value otherwise.

### GetSchemaVersionOk

`func (o *AddSuppressionResponse) GetSchemaVersionOk() (*string, bool)`

GetSchemaVersionOk returns a tuple with the SchemaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaVersion

`func (o *AddSuppressionResponse) SetSchemaVersion(v string)`

SetSchemaVersion sets SchemaVersion field to given value.

### HasSchemaVersion

`func (o *AddSuppressionResponse) HasSchemaVersion() bool`

HasSchemaVersion returns a boolean if a field has been set.

### GetRequestId

`func (o *AddSuppressionResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *AddSuppressionResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *AddSuppressionResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *AddSuppressionResponse) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetAdded

`func (o *AddSuppressionResponse) GetAdded() int32`

GetAdded returns the Added field if non-nil, zero value otherwise.

### GetAddedOk

`func (o *AddSuppressionResponse) GetAddedOk() (*int32, bool)`

GetAddedOk returns a tuple with the Added field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdded

`func (o *AddSuppressionResponse) SetAdded(v int32)`

SetAdded sets Added field to given value.

### HasAdded

`func (o *AddSuppressionResponse) HasAdded() bool`

HasAdded returns a boolean if a field has been set.

### GetDuplicates

`func (o *AddSuppressionResponse) GetDuplicates() int32`

GetDuplicates returns the Duplicates field if non-nil, zero value otherwise.

### GetDuplicatesOk

`func (o *AddSuppressionResponse) GetDuplicatesOk() (*int32, bool)`

GetDuplicatesOk returns a tuple with the Duplicates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuplicates

`func (o *AddSuppressionResponse) SetDuplicates(v int32)`

SetDuplicates sets Duplicates field to given value.

### HasDuplicates

`func (o *AddSuppressionResponse) HasDuplicates() bool`

HasDuplicates returns a boolean if a field has been set.

### GetInvalid

`func (o *AddSuppressionResponse) GetInvalid() int32`

GetInvalid returns the Invalid field if non-nil, zero value otherwise.

### GetInvalidOk

`func (o *AddSuppressionResponse) GetInvalidOk() (*int32, bool)`

GetInvalidOk returns a tuple with the Invalid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalid

`func (o *AddSuppressionResponse) SetInvalid(v int32)`

SetInvalid sets Invalid field to given value.

### HasInvalid

`func (o *AddSuppressionResponse) HasInvalid() bool`

HasInvalid returns a boolean if a field has been set.

### GetTotal

`func (o *AddSuppressionResponse) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *AddSuppressionResponse) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *AddSuppressionResponse) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *AddSuppressionResponse) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


