# PresignedUploadResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SchemaVersion** | Pointer to **string** |  | [optional] 
**Upload** | Pointer to [**PresignedUploadResponseUpload**](PresignedUploadResponseUpload.md) |  | [optional] 

## Methods

### NewPresignedUploadResponse

`func NewPresignedUploadResponse() *PresignedUploadResponse`

NewPresignedUploadResponse instantiates a new PresignedUploadResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPresignedUploadResponseWithDefaults

`func NewPresignedUploadResponseWithDefaults() *PresignedUploadResponse`

NewPresignedUploadResponseWithDefaults instantiates a new PresignedUploadResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemaVersion

`func (o *PresignedUploadResponse) GetSchemaVersion() string`

GetSchemaVersion returns the SchemaVersion field if non-nil, zero value otherwise.

### GetSchemaVersionOk

`func (o *PresignedUploadResponse) GetSchemaVersionOk() (*string, bool)`

GetSchemaVersionOk returns a tuple with the SchemaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaVersion

`func (o *PresignedUploadResponse) SetSchemaVersion(v string)`

SetSchemaVersion sets SchemaVersion field to given value.

### HasSchemaVersion

`func (o *PresignedUploadResponse) HasSchemaVersion() bool`

HasSchemaVersion returns a boolean if a field has been set.

### GetUpload

`func (o *PresignedUploadResponse) GetUpload() PresignedUploadResponseUpload`

GetUpload returns the Upload field if non-nil, zero value otherwise.

### GetUploadOk

`func (o *PresignedUploadResponse) GetUploadOk() (*PresignedUploadResponseUpload, bool)`

GetUploadOk returns a tuple with the Upload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpload

`func (o *PresignedUploadResponse) SetUpload(v PresignedUploadResponseUpload)`

SetUpload sets Upload field to given value.

### HasUpload

`func (o *PresignedUploadResponse) HasUpload() bool`

HasUpload returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


