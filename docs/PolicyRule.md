# PolicyRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Type** | **string** | Rule type determines how condition is evaluated | 
**Condition** | **map[string]interface{}** | Condition depends on rule type. status_override: {status}, domain_filter: {domain_mode, domains}, check_requirement: {check, required}, sub_status_override: {sub_status} | 
**Action** | [**PolicyRuleAction**](PolicyRuleAction.md) |  | 
**Sequence** | Pointer to **int32** |  | [optional] 
**IsEnabled** | Pointer to **bool** |  | [optional] 

## Methods

### NewPolicyRule

`func NewPolicyRule(type_ string, condition map[string]interface{}, action PolicyRuleAction, ) *PolicyRule`

NewPolicyRule instantiates a new PolicyRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyRuleWithDefaults

`func NewPolicyRuleWithDefaults() *PolicyRule`

NewPolicyRuleWithDefaults instantiates a new PolicyRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PolicyRule) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PolicyRule) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PolicyRule) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *PolicyRule) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *PolicyRule) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PolicyRule) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PolicyRule) SetType(v string)`

SetType sets Type field to given value.


### GetCondition

`func (o *PolicyRule) GetCondition() map[string]interface{}`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *PolicyRule) GetConditionOk() (*map[string]interface{}, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *PolicyRule) SetCondition(v map[string]interface{})`

SetCondition sets Condition field to given value.


### GetAction

`func (o *PolicyRule) GetAction() PolicyRuleAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *PolicyRule) GetActionOk() (*PolicyRuleAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *PolicyRule) SetAction(v PolicyRuleAction)`

SetAction sets Action field to given value.


### GetSequence

`func (o *PolicyRule) GetSequence() int32`

GetSequence returns the Sequence field if non-nil, zero value otherwise.

### GetSequenceOk

`func (o *PolicyRule) GetSequenceOk() (*int32, bool)`

GetSequenceOk returns a tuple with the Sequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequence

`func (o *PolicyRule) SetSequence(v int32)`

SetSequence sets Sequence field to given value.

### HasSequence

`func (o *PolicyRule) HasSequence() bool`

HasSequence returns a boolean if a field has been set.

### GetIsEnabled

`func (o *PolicyRule) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *PolicyRule) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *PolicyRule) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *PolicyRule) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


