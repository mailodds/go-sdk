# ClassifyContent200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SchemaVersion** | Pointer to **string** |  | [optional] 
**RequestId** | Pointer to **string** |  | [optional] 
**ContentCheck** | Pointer to [**ClassifyContent200ResponseContentCheck**](ClassifyContent200ResponseContentCheck.md) |  | [optional] 

## Methods

### NewClassifyContent200Response

`func NewClassifyContent200Response() *ClassifyContent200Response`

NewClassifyContent200Response instantiates a new ClassifyContent200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClassifyContent200ResponseWithDefaults

`func NewClassifyContent200ResponseWithDefaults() *ClassifyContent200Response`

NewClassifyContent200ResponseWithDefaults instantiates a new ClassifyContent200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemaVersion

`func (o *ClassifyContent200Response) GetSchemaVersion() string`

GetSchemaVersion returns the SchemaVersion field if non-nil, zero value otherwise.

### GetSchemaVersionOk

`func (o *ClassifyContent200Response) GetSchemaVersionOk() (*string, bool)`

GetSchemaVersionOk returns a tuple with the SchemaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaVersion

`func (o *ClassifyContent200Response) SetSchemaVersion(v string)`

SetSchemaVersion sets SchemaVersion field to given value.

### HasSchemaVersion

`func (o *ClassifyContent200Response) HasSchemaVersion() bool`

HasSchemaVersion returns a boolean if a field has been set.

### GetRequestId

`func (o *ClassifyContent200Response) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *ClassifyContent200Response) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *ClassifyContent200Response) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *ClassifyContent200Response) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetContentCheck

`func (o *ClassifyContent200Response) GetContentCheck() ClassifyContent200ResponseContentCheck`

GetContentCheck returns the ContentCheck field if non-nil, zero value otherwise.

### GetContentCheckOk

`func (o *ClassifyContent200Response) GetContentCheckOk() (*ClassifyContent200ResponseContentCheck, bool)`

GetContentCheckOk returns a tuple with the ContentCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentCheck

`func (o *ClassifyContent200Response) SetContentCheck(v ClassifyContent200ResponseContentCheck)`

SetContentCheck sets ContentCheck field to given value.

### HasContentCheck

`func (o *ClassifyContent200Response) HasContentCheck() bool`

HasContentCheck returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


