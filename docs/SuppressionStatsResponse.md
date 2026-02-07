# SuppressionStatsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SchemaVersion** | Pointer to **string** |  | [optional] 
**Total** | Pointer to **int32** |  | [optional] 
**ByType** | Pointer to [**SuppressionStatsResponseByType**](SuppressionStatsResponseByType.md) |  | [optional] 

## Methods

### NewSuppressionStatsResponse

`func NewSuppressionStatsResponse() *SuppressionStatsResponse`

NewSuppressionStatsResponse instantiates a new SuppressionStatsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSuppressionStatsResponseWithDefaults

`func NewSuppressionStatsResponseWithDefaults() *SuppressionStatsResponse`

NewSuppressionStatsResponseWithDefaults instantiates a new SuppressionStatsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemaVersion

`func (o *SuppressionStatsResponse) GetSchemaVersion() string`

GetSchemaVersion returns the SchemaVersion field if non-nil, zero value otherwise.

### GetSchemaVersionOk

`func (o *SuppressionStatsResponse) GetSchemaVersionOk() (*string, bool)`

GetSchemaVersionOk returns a tuple with the SchemaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaVersion

`func (o *SuppressionStatsResponse) SetSchemaVersion(v string)`

SetSchemaVersion sets SchemaVersion field to given value.

### HasSchemaVersion

`func (o *SuppressionStatsResponse) HasSchemaVersion() bool`

HasSchemaVersion returns a boolean if a field has been set.

### GetTotal

`func (o *SuppressionStatsResponse) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *SuppressionStatsResponse) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *SuppressionStatsResponse) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *SuppressionStatsResponse) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetByType

`func (o *SuppressionStatsResponse) GetByType() SuppressionStatsResponseByType`

GetByType returns the ByType field if non-nil, zero value otherwise.

### GetByTypeOk

`func (o *SuppressionStatsResponse) GetByTypeOk() (*SuppressionStatsResponseByType, bool)`

GetByTypeOk returns a tuple with the ByType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByType

`func (o *SuppressionStatsResponse) SetByType(v SuppressionStatsResponseByType)`

SetByType sets ByType field to given value.

### HasByType

`func (o *SuppressionStatsResponse) HasByType() bool`

HasByType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


