# GetSenderHealthTrend200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SchemaVersion** | Pointer to **string** |  | [optional] 
**RequestId** | Pointer to **string** | Unique request identifier | [optional] 
**Period** | Pointer to **string** |  | [optional] 
**DataPoints** | Pointer to [**[]GetSenderHealthTrend200ResponseDataPointsInner**](GetSenderHealthTrend200ResponseDataPointsInner.md) |  | [optional] 

## Methods

### NewGetSenderHealthTrend200Response

`func NewGetSenderHealthTrend200Response() *GetSenderHealthTrend200Response`

NewGetSenderHealthTrend200Response instantiates a new GetSenderHealthTrend200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSenderHealthTrend200ResponseWithDefaults

`func NewGetSenderHealthTrend200ResponseWithDefaults() *GetSenderHealthTrend200Response`

NewGetSenderHealthTrend200ResponseWithDefaults instantiates a new GetSenderHealthTrend200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemaVersion

`func (o *GetSenderHealthTrend200Response) GetSchemaVersion() string`

GetSchemaVersion returns the SchemaVersion field if non-nil, zero value otherwise.

### GetSchemaVersionOk

`func (o *GetSenderHealthTrend200Response) GetSchemaVersionOk() (*string, bool)`

GetSchemaVersionOk returns a tuple with the SchemaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaVersion

`func (o *GetSenderHealthTrend200Response) SetSchemaVersion(v string)`

SetSchemaVersion sets SchemaVersion field to given value.

### HasSchemaVersion

`func (o *GetSenderHealthTrend200Response) HasSchemaVersion() bool`

HasSchemaVersion returns a boolean if a field has been set.

### GetRequestId

`func (o *GetSenderHealthTrend200Response) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *GetSenderHealthTrend200Response) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *GetSenderHealthTrend200Response) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *GetSenderHealthTrend200Response) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetPeriod

`func (o *GetSenderHealthTrend200Response) GetPeriod() string`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *GetSenderHealthTrend200Response) GetPeriodOk() (*string, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *GetSenderHealthTrend200Response) SetPeriod(v string)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *GetSenderHealthTrend200Response) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### GetDataPoints

`func (o *GetSenderHealthTrend200Response) GetDataPoints() []GetSenderHealthTrend200ResponseDataPointsInner`

GetDataPoints returns the DataPoints field if non-nil, zero value otherwise.

### GetDataPointsOk

`func (o *GetSenderHealthTrend200Response) GetDataPointsOk() (*[]GetSenderHealthTrend200ResponseDataPointsInner, bool)`

GetDataPointsOk returns a tuple with the DataPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataPoints

`func (o *GetSenderHealthTrend200Response) SetDataPoints(v []GetSenderHealthTrend200ResponseDataPointsInner)`

SetDataPoints sets DataPoints field to given value.

### HasDataPoints

`func (o *GetSenderHealthTrend200Response) HasDataPoints() bool`

HasDataPoints returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


