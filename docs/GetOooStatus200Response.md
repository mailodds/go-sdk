# GetOooStatus200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **string** |  | [optional] 
**IsOoo** | Pointer to **bool** |  | [optional] 
**DetectedAt** | Pointer to **NullableTime** |  | [optional] 
**ExpiresAt** | Pointer to **NullableTime** |  | [optional] 
**OooType** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewGetOooStatus200Response

`func NewGetOooStatus200Response() *GetOooStatus200Response`

NewGetOooStatus200Response instantiates a new GetOooStatus200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOooStatus200ResponseWithDefaults

`func NewGetOooStatus200ResponseWithDefaults() *GetOooStatus200Response`

NewGetOooStatus200ResponseWithDefaults instantiates a new GetOooStatus200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *GetOooStatus200Response) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *GetOooStatus200Response) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *GetOooStatus200Response) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *GetOooStatus200Response) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetIsOoo

`func (o *GetOooStatus200Response) GetIsOoo() bool`

GetIsOoo returns the IsOoo field if non-nil, zero value otherwise.

### GetIsOooOk

`func (o *GetOooStatus200Response) GetIsOooOk() (*bool, bool)`

GetIsOooOk returns a tuple with the IsOoo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOoo

`func (o *GetOooStatus200Response) SetIsOoo(v bool)`

SetIsOoo sets IsOoo field to given value.

### HasIsOoo

`func (o *GetOooStatus200Response) HasIsOoo() bool`

HasIsOoo returns a boolean if a field has been set.

### GetDetectedAt

`func (o *GetOooStatus200Response) GetDetectedAt() time.Time`

GetDetectedAt returns the DetectedAt field if non-nil, zero value otherwise.

### GetDetectedAtOk

`func (o *GetOooStatus200Response) GetDetectedAtOk() (*time.Time, bool)`

GetDetectedAtOk returns a tuple with the DetectedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetectedAt

`func (o *GetOooStatus200Response) SetDetectedAt(v time.Time)`

SetDetectedAt sets DetectedAt field to given value.

### HasDetectedAt

`func (o *GetOooStatus200Response) HasDetectedAt() bool`

HasDetectedAt returns a boolean if a field has been set.

### SetDetectedAtNil

`func (o *GetOooStatus200Response) SetDetectedAtNil(b bool)`

 SetDetectedAtNil sets the value for DetectedAt to be an explicit nil

### UnsetDetectedAt
`func (o *GetOooStatus200Response) UnsetDetectedAt()`

UnsetDetectedAt ensures that no value is present for DetectedAt, not even an explicit nil
### GetExpiresAt

`func (o *GetOooStatus200Response) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *GetOooStatus200Response) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *GetOooStatus200Response) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *GetOooStatus200Response) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### SetExpiresAtNil

`func (o *GetOooStatus200Response) SetExpiresAtNil(b bool)`

 SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil

### UnsetExpiresAt
`func (o *GetOooStatus200Response) UnsetExpiresAt()`

UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil
### GetOooType

`func (o *GetOooStatus200Response) GetOooType() string`

GetOooType returns the OooType field if non-nil, zero value otherwise.

### GetOooTypeOk

`func (o *GetOooStatus200Response) GetOooTypeOk() (*string, bool)`

GetOooTypeOk returns a tuple with the OooType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOooType

`func (o *GetOooStatus200Response) SetOooType(v string)`

SetOooType sets OooType field to given value.

### HasOooType

`func (o *GetOooStatus200Response) HasOooType() bool`

HasOooType returns a boolean if a field has been set.

### SetOooTypeNil

`func (o *GetOooStatus200Response) SetOooTypeNil(b bool)`

 SetOooTypeNil sets the value for OooType to be an explicit nil

### UnsetOooType
`func (o *GetOooStatus200Response) UnsetOooType()`

UnsetOooType ensures that no value is present for OooType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


