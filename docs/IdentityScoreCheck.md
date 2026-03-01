# IdentityScoreCheck

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** | Check status (e.g. verified, pending, missing) | 
**Points** | **int32** | Points earned for this check | 
**MaxPoints** | **int32** | Maximum points available for this check | 
**VerifiedAt** | Pointer to **NullableTime** | When the check was last verified | [optional] 

## Methods

### NewIdentityScoreCheck

`func NewIdentityScoreCheck(status string, points int32, maxPoints int32, ) *IdentityScoreCheck`

NewIdentityScoreCheck instantiates a new IdentityScoreCheck object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityScoreCheckWithDefaults

`func NewIdentityScoreCheckWithDefaults() *IdentityScoreCheck`

NewIdentityScoreCheckWithDefaults instantiates a new IdentityScoreCheck object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *IdentityScoreCheck) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *IdentityScoreCheck) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *IdentityScoreCheck) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetPoints

`func (o *IdentityScoreCheck) GetPoints() int32`

GetPoints returns the Points field if non-nil, zero value otherwise.

### GetPointsOk

`func (o *IdentityScoreCheck) GetPointsOk() (*int32, bool)`

GetPointsOk returns a tuple with the Points field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoints

`func (o *IdentityScoreCheck) SetPoints(v int32)`

SetPoints sets Points field to given value.


### GetMaxPoints

`func (o *IdentityScoreCheck) GetMaxPoints() int32`

GetMaxPoints returns the MaxPoints field if non-nil, zero value otherwise.

### GetMaxPointsOk

`func (o *IdentityScoreCheck) GetMaxPointsOk() (*int32, bool)`

GetMaxPointsOk returns a tuple with the MaxPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPoints

`func (o *IdentityScoreCheck) SetMaxPoints(v int32)`

SetMaxPoints sets MaxPoints field to given value.


### GetVerifiedAt

`func (o *IdentityScoreCheck) GetVerifiedAt() time.Time`

GetVerifiedAt returns the VerifiedAt field if non-nil, zero value otherwise.

### GetVerifiedAtOk

`func (o *IdentityScoreCheck) GetVerifiedAtOk() (*time.Time, bool)`

GetVerifiedAtOk returns a tuple with the VerifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifiedAt

`func (o *IdentityScoreCheck) SetVerifiedAt(v time.Time)`

SetVerifiedAt sets VerifiedAt field to given value.

### HasVerifiedAt

`func (o *IdentityScoreCheck) HasVerifiedAt() bool`

HasVerifiedAt returns a boolean if a field has been set.

### SetVerifiedAtNil

`func (o *IdentityScoreCheck) SetVerifiedAtNil(b bool)`

 SetVerifiedAtNil sets the value for VerifiedAt to be an explicit nil

### UnsetVerifiedAt
`func (o *IdentityScoreCheck) UnsetVerifiedAt()`

UnsetVerifiedAt ensures that no value is present for VerifiedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


