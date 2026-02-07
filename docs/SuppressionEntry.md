# SuppressionEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Value** | Pointer to **string** |  | [optional] 
**Reason** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewSuppressionEntry

`func NewSuppressionEntry() *SuppressionEntry`

NewSuppressionEntry instantiates a new SuppressionEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSuppressionEntryWithDefaults

`func NewSuppressionEntryWithDefaults() *SuppressionEntry`

NewSuppressionEntryWithDefaults instantiates a new SuppressionEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SuppressionEntry) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SuppressionEntry) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SuppressionEntry) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *SuppressionEntry) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *SuppressionEntry) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SuppressionEntry) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SuppressionEntry) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SuppressionEntry) HasType() bool`

HasType returns a boolean if a field has been set.

### GetValue

`func (o *SuppressionEntry) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *SuppressionEntry) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *SuppressionEntry) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *SuppressionEntry) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetReason

`func (o *SuppressionEntry) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *SuppressionEntry) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *SuppressionEntry) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *SuppressionEntry) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetCreatedAt

`func (o *SuppressionEntry) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SuppressionEntry) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SuppressionEntry) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *SuppressionEntry) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


