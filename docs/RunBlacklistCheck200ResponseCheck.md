# RunBlacklistCheck200ResponseCheck

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**ListedOn** | Pointer to **[]string** | Blacklists where the target is listed | [optional] 
**CleanOn** | Pointer to **[]string** | Blacklists where the target is clean | [optional] 
**CheckedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewRunBlacklistCheck200ResponseCheck

`func NewRunBlacklistCheck200ResponseCheck() *RunBlacklistCheck200ResponseCheck`

NewRunBlacklistCheck200ResponseCheck instantiates a new RunBlacklistCheck200ResponseCheck object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRunBlacklistCheck200ResponseCheckWithDefaults

`func NewRunBlacklistCheck200ResponseCheckWithDefaults() *RunBlacklistCheck200ResponseCheck`

NewRunBlacklistCheck200ResponseCheckWithDefaults instantiates a new RunBlacklistCheck200ResponseCheck object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RunBlacklistCheck200ResponseCheck) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RunBlacklistCheck200ResponseCheck) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RunBlacklistCheck200ResponseCheck) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RunBlacklistCheck200ResponseCheck) HasId() bool`

HasId returns a boolean if a field has been set.

### GetListedOn

`func (o *RunBlacklistCheck200ResponseCheck) GetListedOn() []string`

GetListedOn returns the ListedOn field if non-nil, zero value otherwise.

### GetListedOnOk

`func (o *RunBlacklistCheck200ResponseCheck) GetListedOnOk() (*[]string, bool)`

GetListedOnOk returns a tuple with the ListedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListedOn

`func (o *RunBlacklistCheck200ResponseCheck) SetListedOn(v []string)`

SetListedOn sets ListedOn field to given value.

### HasListedOn

`func (o *RunBlacklistCheck200ResponseCheck) HasListedOn() bool`

HasListedOn returns a boolean if a field has been set.

### GetCleanOn

`func (o *RunBlacklistCheck200ResponseCheck) GetCleanOn() []string`

GetCleanOn returns the CleanOn field if non-nil, zero value otherwise.

### GetCleanOnOk

`func (o *RunBlacklistCheck200ResponseCheck) GetCleanOnOk() (*[]string, bool)`

GetCleanOnOk returns a tuple with the CleanOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCleanOn

`func (o *RunBlacklistCheck200ResponseCheck) SetCleanOn(v []string)`

SetCleanOn sets CleanOn field to given value.

### HasCleanOn

`func (o *RunBlacklistCheck200ResponseCheck) HasCleanOn() bool`

HasCleanOn returns a boolean if a field has been set.

### GetCheckedAt

`func (o *RunBlacklistCheck200ResponseCheck) GetCheckedAt() time.Time`

GetCheckedAt returns the CheckedAt field if non-nil, zero value otherwise.

### GetCheckedAtOk

`func (o *RunBlacklistCheck200ResponseCheck) GetCheckedAtOk() (*time.Time, bool)`

GetCheckedAtOk returns a tuple with the CheckedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckedAt

`func (o *RunBlacklistCheck200ResponseCheck) SetCheckedAt(v time.Time)`

SetCheckedAt sets CheckedAt field to given value.

### HasCheckedAt

`func (o *RunBlacklistCheck200ResponseCheck) HasCheckedAt() bool`

HasCheckedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


