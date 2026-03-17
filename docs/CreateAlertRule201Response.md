# CreateAlertRule201Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SchemaVersion** | Pointer to **string** |  | [optional] 
**RequestId** | Pointer to **string** |  | [optional] 
**Rule** | Pointer to [**AlertRule**](AlertRule.md) |  | [optional] 

## Methods

### NewCreateAlertRule201Response

`func NewCreateAlertRule201Response() *CreateAlertRule201Response`

NewCreateAlertRule201Response instantiates a new CreateAlertRule201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAlertRule201ResponseWithDefaults

`func NewCreateAlertRule201ResponseWithDefaults() *CreateAlertRule201Response`

NewCreateAlertRule201ResponseWithDefaults instantiates a new CreateAlertRule201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemaVersion

`func (o *CreateAlertRule201Response) GetSchemaVersion() string`

GetSchemaVersion returns the SchemaVersion field if non-nil, zero value otherwise.

### GetSchemaVersionOk

`func (o *CreateAlertRule201Response) GetSchemaVersionOk() (*string, bool)`

GetSchemaVersionOk returns a tuple with the SchemaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaVersion

`func (o *CreateAlertRule201Response) SetSchemaVersion(v string)`

SetSchemaVersion sets SchemaVersion field to given value.

### HasSchemaVersion

`func (o *CreateAlertRule201Response) HasSchemaVersion() bool`

HasSchemaVersion returns a boolean if a field has been set.

### GetRequestId

`func (o *CreateAlertRule201Response) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *CreateAlertRule201Response) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *CreateAlertRule201Response) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *CreateAlertRule201Response) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetRule

`func (o *CreateAlertRule201Response) GetRule() AlertRule`

GetRule returns the Rule field if non-nil, zero value otherwise.

### GetRuleOk

`func (o *CreateAlertRule201Response) GetRuleOk() (*AlertRule, bool)`

GetRuleOk returns a tuple with the Rule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRule

`func (o *CreateAlertRule201Response) SetRule(v AlertRule)`

SetRule sets Rule field to given value.

### HasRule

`func (o *CreateAlertRule201Response) HasRule() bool`

HasRule returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


