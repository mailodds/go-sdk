# ListAlertRules200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SchemaVersion** | Pointer to **string** |  | [optional] 
**RequestId** | Pointer to **string** |  | [optional] 
**Rules** | Pointer to [**[]AlertRule**](AlertRule.md) |  | [optional] 

## Methods

### NewListAlertRules200Response

`func NewListAlertRules200Response() *ListAlertRules200Response`

NewListAlertRules200Response instantiates a new ListAlertRules200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListAlertRules200ResponseWithDefaults

`func NewListAlertRules200ResponseWithDefaults() *ListAlertRules200Response`

NewListAlertRules200ResponseWithDefaults instantiates a new ListAlertRules200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemaVersion

`func (o *ListAlertRules200Response) GetSchemaVersion() string`

GetSchemaVersion returns the SchemaVersion field if non-nil, zero value otherwise.

### GetSchemaVersionOk

`func (o *ListAlertRules200Response) GetSchemaVersionOk() (*string, bool)`

GetSchemaVersionOk returns a tuple with the SchemaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaVersion

`func (o *ListAlertRules200Response) SetSchemaVersion(v string)`

SetSchemaVersion sets SchemaVersion field to given value.

### HasSchemaVersion

`func (o *ListAlertRules200Response) HasSchemaVersion() bool`

HasSchemaVersion returns a boolean if a field has been set.

### GetRequestId

`func (o *ListAlertRules200Response) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *ListAlertRules200Response) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *ListAlertRules200Response) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *ListAlertRules200Response) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetRules

`func (o *ListAlertRules200Response) GetRules() []AlertRule`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *ListAlertRules200Response) GetRulesOk() (*[]AlertRule, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *ListAlertRules200Response) SetRules(v []AlertRule)`

SetRules sets Rules field to given value.

### HasRules

`func (o *ListAlertRules200Response) HasRules() bool`

HasRules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


