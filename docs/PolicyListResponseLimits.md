# PolicyListResponseLimits

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxPolicies** | Pointer to **int32** | -1 means unlimited | [optional] 
**MaxRulesPerPolicy** | Pointer to **int32** |  | [optional] 
**Plan** | Pointer to **string** |  | [optional] 

## Methods

### NewPolicyListResponseLimits

`func NewPolicyListResponseLimits() *PolicyListResponseLimits`

NewPolicyListResponseLimits instantiates a new PolicyListResponseLimits object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyListResponseLimitsWithDefaults

`func NewPolicyListResponseLimitsWithDefaults() *PolicyListResponseLimits`

NewPolicyListResponseLimitsWithDefaults instantiates a new PolicyListResponseLimits object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxPolicies

`func (o *PolicyListResponseLimits) GetMaxPolicies() int32`

GetMaxPolicies returns the MaxPolicies field if non-nil, zero value otherwise.

### GetMaxPoliciesOk

`func (o *PolicyListResponseLimits) GetMaxPoliciesOk() (*int32, bool)`

GetMaxPoliciesOk returns a tuple with the MaxPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPolicies

`func (o *PolicyListResponseLimits) SetMaxPolicies(v int32)`

SetMaxPolicies sets MaxPolicies field to given value.

### HasMaxPolicies

`func (o *PolicyListResponseLimits) HasMaxPolicies() bool`

HasMaxPolicies returns a boolean if a field has been set.

### GetMaxRulesPerPolicy

`func (o *PolicyListResponseLimits) GetMaxRulesPerPolicy() int32`

GetMaxRulesPerPolicy returns the MaxRulesPerPolicy field if non-nil, zero value otherwise.

### GetMaxRulesPerPolicyOk

`func (o *PolicyListResponseLimits) GetMaxRulesPerPolicyOk() (*int32, bool)`

GetMaxRulesPerPolicyOk returns a tuple with the MaxRulesPerPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRulesPerPolicy

`func (o *PolicyListResponseLimits) SetMaxRulesPerPolicy(v int32)`

SetMaxRulesPerPolicy sets MaxRulesPerPolicy field to given value.

### HasMaxRulesPerPolicy

`func (o *PolicyListResponseLimits) HasMaxRulesPerPolicy() bool`

HasMaxRulesPerPolicy returns a boolean if a field has been set.

### GetPlan

`func (o *PolicyListResponseLimits) GetPlan() string`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *PolicyListResponseLimits) GetPlanOk() (*string, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *PolicyListResponseLimits) SetPlan(v string)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *PolicyListResponseLimits) HasPlan() bool`

HasPlan returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


