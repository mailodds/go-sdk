# CreatePolicyFromPresetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PresetId** | **string** |  | 
**Name** | **string** |  | 
**IsDefault** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewCreatePolicyFromPresetRequest

`func NewCreatePolicyFromPresetRequest(presetId string, name string, ) *CreatePolicyFromPresetRequest`

NewCreatePolicyFromPresetRequest instantiates a new CreatePolicyFromPresetRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePolicyFromPresetRequestWithDefaults

`func NewCreatePolicyFromPresetRequestWithDefaults() *CreatePolicyFromPresetRequest`

NewCreatePolicyFromPresetRequestWithDefaults instantiates a new CreatePolicyFromPresetRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPresetId

`func (o *CreatePolicyFromPresetRequest) GetPresetId() string`

GetPresetId returns the PresetId field if non-nil, zero value otherwise.

### GetPresetIdOk

`func (o *CreatePolicyFromPresetRequest) GetPresetIdOk() (*string, bool)`

GetPresetIdOk returns a tuple with the PresetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresetId

`func (o *CreatePolicyFromPresetRequest) SetPresetId(v string)`

SetPresetId sets PresetId field to given value.


### GetName

`func (o *CreatePolicyFromPresetRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreatePolicyFromPresetRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreatePolicyFromPresetRequest) SetName(v string)`

SetName sets Name field to given value.


### GetIsDefault

`func (o *CreatePolicyFromPresetRequest) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *CreatePolicyFromPresetRequest) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *CreatePolicyFromPresetRequest) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *CreatePolicyFromPresetRequest) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


