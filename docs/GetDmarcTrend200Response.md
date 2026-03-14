# GetDmarcTrend200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SchemaVersion** | Pointer to **string** |  | [optional] 
**RequestId** | Pointer to **string** |  | [optional] 
**Trend** | Pointer to [**[]GetDmarcTrend200ResponseTrendInner**](GetDmarcTrend200ResponseTrendInner.md) |  | [optional] 

## Methods

### NewGetDmarcTrend200Response

`func NewGetDmarcTrend200Response() *GetDmarcTrend200Response`

NewGetDmarcTrend200Response instantiates a new GetDmarcTrend200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDmarcTrend200ResponseWithDefaults

`func NewGetDmarcTrend200ResponseWithDefaults() *GetDmarcTrend200Response`

NewGetDmarcTrend200ResponseWithDefaults instantiates a new GetDmarcTrend200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemaVersion

`func (o *GetDmarcTrend200Response) GetSchemaVersion() string`

GetSchemaVersion returns the SchemaVersion field if non-nil, zero value otherwise.

### GetSchemaVersionOk

`func (o *GetDmarcTrend200Response) GetSchemaVersionOk() (*string, bool)`

GetSchemaVersionOk returns a tuple with the SchemaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaVersion

`func (o *GetDmarcTrend200Response) SetSchemaVersion(v string)`

SetSchemaVersion sets SchemaVersion field to given value.

### HasSchemaVersion

`func (o *GetDmarcTrend200Response) HasSchemaVersion() bool`

HasSchemaVersion returns a boolean if a field has been set.

### GetRequestId

`func (o *GetDmarcTrend200Response) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *GetDmarcTrend200Response) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *GetDmarcTrend200Response) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *GetDmarcTrend200Response) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetTrend

`func (o *GetDmarcTrend200Response) GetTrend() []GetDmarcTrend200ResponseTrendInner`

GetTrend returns the Trend field if non-nil, zero value otherwise.

### GetTrendOk

`func (o *GetDmarcTrend200Response) GetTrendOk() (*[]GetDmarcTrend200ResponseTrendInner, bool)`

GetTrendOk returns a tuple with the Trend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrend

`func (o *GetDmarcTrend200Response) SetTrend(v []GetDmarcTrend200ResponseTrendInner)`

SetTrend sets Trend field to given value.

### HasTrend

`func (o *GetDmarcTrend200Response) HasTrend() bool`

HasTrend returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


