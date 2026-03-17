# ImportContactList200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SchemaVersion** | Pointer to **string** |  | [optional] 
**RequestId** | Pointer to **string** |  | [optional] 
**Imported** | Pointer to **int32** |  | [optional] 
**Skipped** | Pointer to **int32** |  | [optional] 
**Duplicates** | Pointer to **int32** |  | [optional] 
**Errors** | Pointer to **int32** |  | [optional] 
**Total** | Pointer to **int32** |  | [optional] 
**ContactList** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewImportContactList200Response

`func NewImportContactList200Response() *ImportContactList200Response`

NewImportContactList200Response instantiates a new ImportContactList200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImportContactList200ResponseWithDefaults

`func NewImportContactList200ResponseWithDefaults() *ImportContactList200Response`

NewImportContactList200ResponseWithDefaults instantiates a new ImportContactList200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemaVersion

`func (o *ImportContactList200Response) GetSchemaVersion() string`

GetSchemaVersion returns the SchemaVersion field if non-nil, zero value otherwise.

### GetSchemaVersionOk

`func (o *ImportContactList200Response) GetSchemaVersionOk() (*string, bool)`

GetSchemaVersionOk returns a tuple with the SchemaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaVersion

`func (o *ImportContactList200Response) SetSchemaVersion(v string)`

SetSchemaVersion sets SchemaVersion field to given value.

### HasSchemaVersion

`func (o *ImportContactList200Response) HasSchemaVersion() bool`

HasSchemaVersion returns a boolean if a field has been set.

### GetRequestId

`func (o *ImportContactList200Response) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *ImportContactList200Response) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *ImportContactList200Response) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *ImportContactList200Response) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetImported

`func (o *ImportContactList200Response) GetImported() int32`

GetImported returns the Imported field if non-nil, zero value otherwise.

### GetImportedOk

`func (o *ImportContactList200Response) GetImportedOk() (*int32, bool)`

GetImportedOk returns a tuple with the Imported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImported

`func (o *ImportContactList200Response) SetImported(v int32)`

SetImported sets Imported field to given value.

### HasImported

`func (o *ImportContactList200Response) HasImported() bool`

HasImported returns a boolean if a field has been set.

### GetSkipped

`func (o *ImportContactList200Response) GetSkipped() int32`

GetSkipped returns the Skipped field if non-nil, zero value otherwise.

### GetSkippedOk

`func (o *ImportContactList200Response) GetSkippedOk() (*int32, bool)`

GetSkippedOk returns a tuple with the Skipped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipped

`func (o *ImportContactList200Response) SetSkipped(v int32)`

SetSkipped sets Skipped field to given value.

### HasSkipped

`func (o *ImportContactList200Response) HasSkipped() bool`

HasSkipped returns a boolean if a field has been set.

### GetDuplicates

`func (o *ImportContactList200Response) GetDuplicates() int32`

GetDuplicates returns the Duplicates field if non-nil, zero value otherwise.

### GetDuplicatesOk

`func (o *ImportContactList200Response) GetDuplicatesOk() (*int32, bool)`

GetDuplicatesOk returns a tuple with the Duplicates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuplicates

`func (o *ImportContactList200Response) SetDuplicates(v int32)`

SetDuplicates sets Duplicates field to given value.

### HasDuplicates

`func (o *ImportContactList200Response) HasDuplicates() bool`

HasDuplicates returns a boolean if a field has been set.

### GetErrors

`func (o *ImportContactList200Response) GetErrors() int32`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *ImportContactList200Response) GetErrorsOk() (*int32, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *ImportContactList200Response) SetErrors(v int32)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *ImportContactList200Response) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetTotal

`func (o *ImportContactList200Response) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ImportContactList200Response) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ImportContactList200Response) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *ImportContactList200Response) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetContactList

`func (o *ImportContactList200Response) GetContactList() map[string]interface{}`

GetContactList returns the ContactList field if non-nil, zero value otherwise.

### GetContactListOk

`func (o *ImportContactList200Response) GetContactListOk() (*map[string]interface{}, bool)`

GetContactListOk returns a tuple with the ContactList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactList

`func (o *ImportContactList200Response) SetContactList(v map[string]interface{})`

SetContactList sets ContactList field to given value.

### HasContactList

`func (o *ImportContactList200Response) HasContactList() bool`

HasContactList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


