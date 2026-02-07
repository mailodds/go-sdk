# TestPolicyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PolicyId** | **int32** |  | 
**TestResult** | [**TestPolicyRequestTestResult**](TestPolicyRequestTestResult.md) |  | 

## Methods

### NewTestPolicyRequest

`func NewTestPolicyRequest(policyId int32, testResult TestPolicyRequestTestResult, ) *TestPolicyRequest`

NewTestPolicyRequest instantiates a new TestPolicyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestPolicyRequestWithDefaults

`func NewTestPolicyRequestWithDefaults() *TestPolicyRequest`

NewTestPolicyRequestWithDefaults instantiates a new TestPolicyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPolicyId

`func (o *TestPolicyRequest) GetPolicyId() int32`

GetPolicyId returns the PolicyId field if non-nil, zero value otherwise.

### GetPolicyIdOk

`func (o *TestPolicyRequest) GetPolicyIdOk() (*int32, bool)`

GetPolicyIdOk returns a tuple with the PolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyId

`func (o *TestPolicyRequest) SetPolicyId(v int32)`

SetPolicyId sets PolicyId field to given value.


### GetTestResult

`func (o *TestPolicyRequest) GetTestResult() TestPolicyRequestTestResult`

GetTestResult returns the TestResult field if non-nil, zero value otherwise.

### GetTestResultOk

`func (o *TestPolicyRequest) GetTestResultOk() (*TestPolicyRequestTestResult, bool)`

GetTestResultOk returns a tuple with the TestResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestResult

`func (o *TestPolicyRequest) SetTestResult(v TestPolicyRequestTestResult)`

SetTestResult sets TestResult field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


