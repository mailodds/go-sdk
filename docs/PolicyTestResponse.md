# PolicyTestResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SchemaVersion** | Pointer to **string** |  | [optional] 
**Original** | Pointer to **map[string]interface{}** | Original validation result before policy | [optional] 
**Modified** | Pointer to **map[string]interface{}** | Result after policy applied | [optional] 
**MatchedRule** | Pointer to **map[string]interface{}** | The rule that matched, or null if none matched | [optional] 
**RulesEvaluated** | Pointer to **int32** |  | [optional] 

## Methods

### NewPolicyTestResponse

`func NewPolicyTestResponse() *PolicyTestResponse`

NewPolicyTestResponse instantiates a new PolicyTestResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyTestResponseWithDefaults

`func NewPolicyTestResponseWithDefaults() *PolicyTestResponse`

NewPolicyTestResponseWithDefaults instantiates a new PolicyTestResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemaVersion

`func (o *PolicyTestResponse) GetSchemaVersion() string`

GetSchemaVersion returns the SchemaVersion field if non-nil, zero value otherwise.

### GetSchemaVersionOk

`func (o *PolicyTestResponse) GetSchemaVersionOk() (*string, bool)`

GetSchemaVersionOk returns a tuple with the SchemaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaVersion

`func (o *PolicyTestResponse) SetSchemaVersion(v string)`

SetSchemaVersion sets SchemaVersion field to given value.

### HasSchemaVersion

`func (o *PolicyTestResponse) HasSchemaVersion() bool`

HasSchemaVersion returns a boolean if a field has been set.

### GetOriginal

`func (o *PolicyTestResponse) GetOriginal() map[string]interface{}`

GetOriginal returns the Original field if non-nil, zero value otherwise.

### GetOriginalOk

`func (o *PolicyTestResponse) GetOriginalOk() (*map[string]interface{}, bool)`

GetOriginalOk returns a tuple with the Original field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginal

`func (o *PolicyTestResponse) SetOriginal(v map[string]interface{})`

SetOriginal sets Original field to given value.

### HasOriginal

`func (o *PolicyTestResponse) HasOriginal() bool`

HasOriginal returns a boolean if a field has been set.

### GetModified

`func (o *PolicyTestResponse) GetModified() map[string]interface{}`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *PolicyTestResponse) GetModifiedOk() (*map[string]interface{}, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *PolicyTestResponse) SetModified(v map[string]interface{})`

SetModified sets Modified field to given value.

### HasModified

`func (o *PolicyTestResponse) HasModified() bool`

HasModified returns a boolean if a field has been set.

### GetMatchedRule

`func (o *PolicyTestResponse) GetMatchedRule() map[string]interface{}`

GetMatchedRule returns the MatchedRule field if non-nil, zero value otherwise.

### GetMatchedRuleOk

`func (o *PolicyTestResponse) GetMatchedRuleOk() (*map[string]interface{}, bool)`

GetMatchedRuleOk returns a tuple with the MatchedRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchedRule

`func (o *PolicyTestResponse) SetMatchedRule(v map[string]interface{})`

SetMatchedRule sets MatchedRule field to given value.

### HasMatchedRule

`func (o *PolicyTestResponse) HasMatchedRule() bool`

HasMatchedRule returns a boolean if a field has been set.

### SetMatchedRuleNil

`func (o *PolicyTestResponse) SetMatchedRuleNil(b bool)`

 SetMatchedRuleNil sets the value for MatchedRule to be an explicit nil

### UnsetMatchedRule
`func (o *PolicyTestResponse) UnsetMatchedRule()`

UnsetMatchedRule ensures that no value is present for MatchedRule, not even an explicit nil
### GetRulesEvaluated

`func (o *PolicyTestResponse) GetRulesEvaluated() int32`

GetRulesEvaluated returns the RulesEvaluated field if non-nil, zero value otherwise.

### GetRulesEvaluatedOk

`func (o *PolicyTestResponse) GetRulesEvaluatedOk() (*int32, bool)`

GetRulesEvaluatedOk returns a tuple with the RulesEvaluated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRulesEvaluated

`func (o *PolicyTestResponse) SetRulesEvaluated(v int32)`

SetRulesEvaluated sets RulesEvaluated field to given value.

### HasRulesEvaluated

`func (o *PolicyTestResponse) HasRulesEvaluated() bool`

HasRulesEvaluated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


