# DeliverRequestOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ValidateFirst** | Pointer to **bool** | Validate recipients before sending | [optional] 

## Methods

### NewDeliverRequestOptions

`func NewDeliverRequestOptions() *DeliverRequestOptions`

NewDeliverRequestOptions instantiates a new DeliverRequestOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeliverRequestOptionsWithDefaults

`func NewDeliverRequestOptionsWithDefaults() *DeliverRequestOptions`

NewDeliverRequestOptionsWithDefaults instantiates a new DeliverRequestOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValidateFirst

`func (o *DeliverRequestOptions) GetValidateFirst() bool`

GetValidateFirst returns the ValidateFirst field if non-nil, zero value otherwise.

### GetValidateFirstOk

`func (o *DeliverRequestOptions) GetValidateFirstOk() (*bool, bool)`

GetValidateFirstOk returns a tuple with the ValidateFirst field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidateFirst

`func (o *DeliverRequestOptions) SetValidateFirst(v bool)`

SetValidateFirst sets ValidateFirst field to given value.

### HasValidateFirst

`func (o *DeliverRequestOptions) HasValidateFirst() bool`

HasValidateFirst returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


