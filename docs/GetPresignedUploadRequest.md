# GetPresignedUploadRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filename** | **string** | Original filename | 
**ContentType** | Pointer to **string** | File MIME type | [optional] 

## Methods

### NewGetPresignedUploadRequest

`func NewGetPresignedUploadRequest(filename string, ) *GetPresignedUploadRequest`

NewGetPresignedUploadRequest instantiates a new GetPresignedUploadRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPresignedUploadRequestWithDefaults

`func NewGetPresignedUploadRequestWithDefaults() *GetPresignedUploadRequest`

NewGetPresignedUploadRequestWithDefaults instantiates a new GetPresignedUploadRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilename

`func (o *GetPresignedUploadRequest) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *GetPresignedUploadRequest) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *GetPresignedUploadRequest) SetFilename(v string)`

SetFilename sets Filename field to given value.


### GetContentType

`func (o *GetPresignedUploadRequest) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *GetPresignedUploadRequest) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *GetPresignedUploadRequest) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *GetPresignedUploadRequest) HasContentType() bool`

HasContentType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


