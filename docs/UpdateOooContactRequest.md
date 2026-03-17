# UpdateOooContactRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsActive** | Pointer to **bool** | Set to false to clear OOO status | [optional] 
**OooType** | Pointer to **string** |  | [optional] [default to "out_of_office"]
**DurationDays** | Pointer to **int32** |  | [optional] [default to 7]

## Methods

### NewUpdateOooContactRequest

`func NewUpdateOooContactRequest() *UpdateOooContactRequest`

NewUpdateOooContactRequest instantiates a new UpdateOooContactRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateOooContactRequestWithDefaults

`func NewUpdateOooContactRequestWithDefaults() *UpdateOooContactRequest`

NewUpdateOooContactRequestWithDefaults instantiates a new UpdateOooContactRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsActive

`func (o *UpdateOooContactRequest) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *UpdateOooContactRequest) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *UpdateOooContactRequest) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *UpdateOooContactRequest) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetOooType

`func (o *UpdateOooContactRequest) GetOooType() string`

GetOooType returns the OooType field if non-nil, zero value otherwise.

### GetOooTypeOk

`func (o *UpdateOooContactRequest) GetOooTypeOk() (*string, bool)`

GetOooTypeOk returns a tuple with the OooType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOooType

`func (o *UpdateOooContactRequest) SetOooType(v string)`

SetOooType sets OooType field to given value.

### HasOooType

`func (o *UpdateOooContactRequest) HasOooType() bool`

HasOooType returns a boolean if a field has been set.

### GetDurationDays

`func (o *UpdateOooContactRequest) GetDurationDays() int32`

GetDurationDays returns the DurationDays field if non-nil, zero value otherwise.

### GetDurationDaysOk

`func (o *UpdateOooContactRequest) GetDurationDaysOk() (*int32, bool)`

GetDurationDaysOk returns a tuple with the DurationDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationDays

`func (o *UpdateOooContactRequest) SetDurationDays(v int32)`

SetDurationDays sets DurationDays field to given value.

### HasDurationDays

`func (o *UpdateOooContactRequest) HasDurationDays() bool`

HasDurationDays returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


