# RunServerTest201Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SchemaVersion** | Pointer to **string** |  | [optional] 
**RequestId** | Pointer to **string** |  | [optional] 
**Test** | Pointer to [**ServerTest**](ServerTest.md) |  | [optional] 

## Methods

### NewRunServerTest201Response

`func NewRunServerTest201Response() *RunServerTest201Response`

NewRunServerTest201Response instantiates a new RunServerTest201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRunServerTest201ResponseWithDefaults

`func NewRunServerTest201ResponseWithDefaults() *RunServerTest201Response`

NewRunServerTest201ResponseWithDefaults instantiates a new RunServerTest201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemaVersion

`func (o *RunServerTest201Response) GetSchemaVersion() string`

GetSchemaVersion returns the SchemaVersion field if non-nil, zero value otherwise.

### GetSchemaVersionOk

`func (o *RunServerTest201Response) GetSchemaVersionOk() (*string, bool)`

GetSchemaVersionOk returns a tuple with the SchemaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaVersion

`func (o *RunServerTest201Response) SetSchemaVersion(v string)`

SetSchemaVersion sets SchemaVersion field to given value.

### HasSchemaVersion

`func (o *RunServerTest201Response) HasSchemaVersion() bool`

HasSchemaVersion returns a boolean if a field has been set.

### GetRequestId

`func (o *RunServerTest201Response) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *RunServerTest201Response) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *RunServerTest201Response) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *RunServerTest201Response) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetTest

`func (o *RunServerTest201Response) GetTest() ServerTest`

GetTest returns the Test field if non-nil, zero value otherwise.

### GetTestOk

`func (o *RunServerTest201Response) GetTestOk() (*ServerTest, bool)`

GetTestOk returns a tuple with the Test field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTest

`func (o *RunServerTest201Response) SetTest(v ServerTest)`

SetTest sets Test field to given value.

### HasTest

`func (o *RunServerTest201Response) HasTest() bool`

HasTest returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


