# CreatePolicyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**IsDefault** | Pointer to **bool** |  | [optional] [default to false]
**Rules** | Pointer to [**[]PolicyRule**](PolicyRule.md) |  | [optional] 

## Methods

### NewCreatePolicyRequest

`func NewCreatePolicyRequest(name string, ) *CreatePolicyRequest`

NewCreatePolicyRequest instantiates a new CreatePolicyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePolicyRequestWithDefaults

`func NewCreatePolicyRequestWithDefaults() *CreatePolicyRequest`

NewCreatePolicyRequestWithDefaults instantiates a new CreatePolicyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreatePolicyRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreatePolicyRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreatePolicyRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CreatePolicyRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreatePolicyRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreatePolicyRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreatePolicyRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIsDefault

`func (o *CreatePolicyRequest) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *CreatePolicyRequest) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *CreatePolicyRequest) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *CreatePolicyRequest) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetRules

`func (o *CreatePolicyRequest) GetRules() []PolicyRule`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *CreatePolicyRequest) GetRulesOk() (*[]PolicyRule, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *CreatePolicyRequest) SetRules(v []PolicyRule)`

SetRules sets Rules field to given value.

### HasRules

`func (o *CreatePolicyRequest) HasRules() bool`

HasRules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


