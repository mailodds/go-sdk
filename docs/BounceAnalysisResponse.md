# BounceAnalysisResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SchemaVersion** | Pointer to **string** |  | [optional] 
**RequestId** | Pointer to **string** | Unique request identifier | [optional] 
**Analysis** | Pointer to [**BounceAnalysisResponseAnalysis**](BounceAnalysisResponseAnalysis.md) |  | [optional] 

## Methods

### NewBounceAnalysisResponse

`func NewBounceAnalysisResponse() *BounceAnalysisResponse`

NewBounceAnalysisResponse instantiates a new BounceAnalysisResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBounceAnalysisResponseWithDefaults

`func NewBounceAnalysisResponseWithDefaults() *BounceAnalysisResponse`

NewBounceAnalysisResponseWithDefaults instantiates a new BounceAnalysisResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemaVersion

`func (o *BounceAnalysisResponse) GetSchemaVersion() string`

GetSchemaVersion returns the SchemaVersion field if non-nil, zero value otherwise.

### GetSchemaVersionOk

`func (o *BounceAnalysisResponse) GetSchemaVersionOk() (*string, bool)`

GetSchemaVersionOk returns a tuple with the SchemaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaVersion

`func (o *BounceAnalysisResponse) SetSchemaVersion(v string)`

SetSchemaVersion sets SchemaVersion field to given value.

### HasSchemaVersion

`func (o *BounceAnalysisResponse) HasSchemaVersion() bool`

HasSchemaVersion returns a boolean if a field has been set.

### GetRequestId

`func (o *BounceAnalysisResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *BounceAnalysisResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *BounceAnalysisResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *BounceAnalysisResponse) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetAnalysis

`func (o *BounceAnalysisResponse) GetAnalysis() BounceAnalysisResponseAnalysis`

GetAnalysis returns the Analysis field if non-nil, zero value otherwise.

### GetAnalysisOk

`func (o *BounceAnalysisResponse) GetAnalysisOk() (*BounceAnalysisResponseAnalysis, bool)`

GetAnalysisOk returns a tuple with the Analysis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalysis

`func (o *BounceAnalysisResponse) SetAnalysis(v BounceAnalysisResponseAnalysis)`

SetAnalysis sets Analysis field to given value.

### HasAnalysis

`func (o *BounceAnalysisResponse) HasAnalysis() bool`

HasAnalysis returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


