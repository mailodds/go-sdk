# GetSenderHealth200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SchemaVersion** | Pointer to **string** |  | [optional] 
**RequestId** | Pointer to **string** | Unique request identifier | [optional] 
**Score** | Pointer to **int32** | Overall sender health score (0-100) | [optional] 
**Grade** | Pointer to **string** | Letter grade based on score | [optional] 
**Period** | Pointer to **string** |  | [optional] 
**Components** | Pointer to [**GetSenderHealth200ResponseComponents**](GetSenderHealth200ResponseComponents.md) |  | [optional] 
**Volume** | Pointer to [**GetSenderHealth200ResponseVolume**](GetSenderHealth200ResponseVolume.md) |  | [optional] 

## Methods

### NewGetSenderHealth200Response

`func NewGetSenderHealth200Response() *GetSenderHealth200Response`

NewGetSenderHealth200Response instantiates a new GetSenderHealth200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSenderHealth200ResponseWithDefaults

`func NewGetSenderHealth200ResponseWithDefaults() *GetSenderHealth200Response`

NewGetSenderHealth200ResponseWithDefaults instantiates a new GetSenderHealth200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemaVersion

`func (o *GetSenderHealth200Response) GetSchemaVersion() string`

GetSchemaVersion returns the SchemaVersion field if non-nil, zero value otherwise.

### GetSchemaVersionOk

`func (o *GetSenderHealth200Response) GetSchemaVersionOk() (*string, bool)`

GetSchemaVersionOk returns a tuple with the SchemaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaVersion

`func (o *GetSenderHealth200Response) SetSchemaVersion(v string)`

SetSchemaVersion sets SchemaVersion field to given value.

### HasSchemaVersion

`func (o *GetSenderHealth200Response) HasSchemaVersion() bool`

HasSchemaVersion returns a boolean if a field has been set.

### GetRequestId

`func (o *GetSenderHealth200Response) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *GetSenderHealth200Response) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *GetSenderHealth200Response) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *GetSenderHealth200Response) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetScore

`func (o *GetSenderHealth200Response) GetScore() int32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *GetSenderHealth200Response) GetScoreOk() (*int32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *GetSenderHealth200Response) SetScore(v int32)`

SetScore sets Score field to given value.

### HasScore

`func (o *GetSenderHealth200Response) HasScore() bool`

HasScore returns a boolean if a field has been set.

### GetGrade

`func (o *GetSenderHealth200Response) GetGrade() string`

GetGrade returns the Grade field if non-nil, zero value otherwise.

### GetGradeOk

`func (o *GetSenderHealth200Response) GetGradeOk() (*string, bool)`

GetGradeOk returns a tuple with the Grade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrade

`func (o *GetSenderHealth200Response) SetGrade(v string)`

SetGrade sets Grade field to given value.

### HasGrade

`func (o *GetSenderHealth200Response) HasGrade() bool`

HasGrade returns a boolean if a field has been set.

### GetPeriod

`func (o *GetSenderHealth200Response) GetPeriod() string`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *GetSenderHealth200Response) GetPeriodOk() (*string, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *GetSenderHealth200Response) SetPeriod(v string)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *GetSenderHealth200Response) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### GetComponents

`func (o *GetSenderHealth200Response) GetComponents() GetSenderHealth200ResponseComponents`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *GetSenderHealth200Response) GetComponentsOk() (*GetSenderHealth200ResponseComponents, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *GetSenderHealth200Response) SetComponents(v GetSenderHealth200ResponseComponents)`

SetComponents sets Components field to given value.

### HasComponents

`func (o *GetSenderHealth200Response) HasComponents() bool`

HasComponents returns a boolean if a field has been set.

### GetVolume

`func (o *GetSenderHealth200Response) GetVolume() GetSenderHealth200ResponseVolume`

GetVolume returns the Volume field if non-nil, zero value otherwise.

### GetVolumeOk

`func (o *GetSenderHealth200Response) GetVolumeOk() (*GetSenderHealth200ResponseVolume, bool)`

GetVolumeOk returns a tuple with the Volume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume

`func (o *GetSenderHealth200Response) SetVolume(v GetSenderHealth200ResponseVolume)`

SetVolume sets Volume field to given value.

### HasVolume

`func (o *GetSenderHealth200Response) HasVolume() bool`

HasVolume returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


