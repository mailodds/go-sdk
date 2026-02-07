# CreateJobFromS3Request

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**S3Key** | **string** | S3 key from presigned upload | 
**Dedup** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewCreateJobFromS3Request

`func NewCreateJobFromS3Request(s3Key string, ) *CreateJobFromS3Request`

NewCreateJobFromS3Request instantiates a new CreateJobFromS3Request object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateJobFromS3RequestWithDefaults

`func NewCreateJobFromS3RequestWithDefaults() *CreateJobFromS3Request`

NewCreateJobFromS3RequestWithDefaults instantiates a new CreateJobFromS3Request object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetS3Key

`func (o *CreateJobFromS3Request) GetS3Key() string`

GetS3Key returns the S3Key field if non-nil, zero value otherwise.

### GetS3KeyOk

`func (o *CreateJobFromS3Request) GetS3KeyOk() (*string, bool)`

GetS3KeyOk returns a tuple with the S3Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3Key

`func (o *CreateJobFromS3Request) SetS3Key(v string)`

SetS3Key sets S3Key field to given value.


### GetDedup

`func (o *CreateJobFromS3Request) GetDedup() bool`

GetDedup returns the Dedup field if non-nil, zero value otherwise.

### GetDedupOk

`func (o *CreateJobFromS3Request) GetDedupOk() (*bool, bool)`

GetDedupOk returns a tuple with the Dedup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDedup

`func (o *CreateJobFromS3Request) SetDedup(v bool)`

SetDedup sets Dedup field to given value.

### HasDedup

`func (o *CreateJobFromS3Request) HasDedup() bool`

HasDedup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


