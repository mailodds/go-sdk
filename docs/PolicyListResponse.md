# PolicyListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SchemaVersion** | Pointer to **string** |  | [optional] 
**Policies** | Pointer to [**[]Policy**](Policy.md) |  | [optional] 
**Limits** | Pointer to [**PolicyListResponseLimits**](PolicyListResponseLimits.md) |  | [optional] 

## Methods

### NewPolicyListResponse

`func NewPolicyListResponse() *PolicyListResponse`

NewPolicyListResponse instantiates a new PolicyListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyListResponseWithDefaults

`func NewPolicyListResponseWithDefaults() *PolicyListResponse`

NewPolicyListResponseWithDefaults instantiates a new PolicyListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemaVersion

`func (o *PolicyListResponse) GetSchemaVersion() string`

GetSchemaVersion returns the SchemaVersion field if non-nil, zero value otherwise.

### GetSchemaVersionOk

`func (o *PolicyListResponse) GetSchemaVersionOk() (*string, bool)`

GetSchemaVersionOk returns a tuple with the SchemaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaVersion

`func (o *PolicyListResponse) SetSchemaVersion(v string)`

SetSchemaVersion sets SchemaVersion field to given value.

### HasSchemaVersion

`func (o *PolicyListResponse) HasSchemaVersion() bool`

HasSchemaVersion returns a boolean if a field has been set.

### GetPolicies

`func (o *PolicyListResponse) GetPolicies() []Policy`

GetPolicies returns the Policies field if non-nil, zero value otherwise.

### GetPoliciesOk

`func (o *PolicyListResponse) GetPoliciesOk() (*[]Policy, bool)`

GetPoliciesOk returns a tuple with the Policies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicies

`func (o *PolicyListResponse) SetPolicies(v []Policy)`

SetPolicies sets Policies field to given value.

### HasPolicies

`func (o *PolicyListResponse) HasPolicies() bool`

HasPolicies returns a boolean if a field has been set.

### GetLimits

`func (o *PolicyListResponse) GetLimits() PolicyListResponseLimits`

GetLimits returns the Limits field if non-nil, zero value otherwise.

### GetLimitsOk

`func (o *PolicyListResponse) GetLimitsOk() (*PolicyListResponseLimits, bool)`

GetLimitsOk returns a tuple with the Limits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimits

`func (o *PolicyListResponse) SetLimits(v PolicyListResponseLimits)`

SetLimits sets Limits field to given value.

### HasLimits

`func (o *PolicyListResponse) HasLimits() bool`

HasLimits returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


