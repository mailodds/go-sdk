# SendingDomainIdentityScoreChecks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dkim** | Pointer to [**SendingDomainIdentityScoreChecksDkim**](SendingDomainIdentityScoreChecksDkim.md) |  | [optional] 
**Spf** | Pointer to [**SendingDomainIdentityScoreChecksDkim**](SendingDomainIdentityScoreChecksDkim.md) |  | [optional] 
**Dmarc** | Pointer to [**SendingDomainIdentityScoreChecksDmarc**](SendingDomainIdentityScoreChecksDmarc.md) |  | [optional] 
**Mx** | Pointer to [**SendingDomainIdentityScoreChecksDkim**](SendingDomainIdentityScoreChecksDkim.md) |  | [optional] 
**ReturnPath** | Pointer to [**SendingDomainIdentityScoreChecksDkim**](SendingDomainIdentityScoreChecksDkim.md) |  | [optional] 

## Methods

### NewSendingDomainIdentityScoreChecks

`func NewSendingDomainIdentityScoreChecks() *SendingDomainIdentityScoreChecks`

NewSendingDomainIdentityScoreChecks instantiates a new SendingDomainIdentityScoreChecks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendingDomainIdentityScoreChecksWithDefaults

`func NewSendingDomainIdentityScoreChecksWithDefaults() *SendingDomainIdentityScoreChecks`

NewSendingDomainIdentityScoreChecksWithDefaults instantiates a new SendingDomainIdentityScoreChecks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDkim

`func (o *SendingDomainIdentityScoreChecks) GetDkim() SendingDomainIdentityScoreChecksDkim`

GetDkim returns the Dkim field if non-nil, zero value otherwise.

### GetDkimOk

`func (o *SendingDomainIdentityScoreChecks) GetDkimOk() (*SendingDomainIdentityScoreChecksDkim, bool)`

GetDkimOk returns a tuple with the Dkim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDkim

`func (o *SendingDomainIdentityScoreChecks) SetDkim(v SendingDomainIdentityScoreChecksDkim)`

SetDkim sets Dkim field to given value.

### HasDkim

`func (o *SendingDomainIdentityScoreChecks) HasDkim() bool`

HasDkim returns a boolean if a field has been set.

### GetSpf

`func (o *SendingDomainIdentityScoreChecks) GetSpf() SendingDomainIdentityScoreChecksDkim`

GetSpf returns the Spf field if non-nil, zero value otherwise.

### GetSpfOk

`func (o *SendingDomainIdentityScoreChecks) GetSpfOk() (*SendingDomainIdentityScoreChecksDkim, bool)`

GetSpfOk returns a tuple with the Spf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpf

`func (o *SendingDomainIdentityScoreChecks) SetSpf(v SendingDomainIdentityScoreChecksDkim)`

SetSpf sets Spf field to given value.

### HasSpf

`func (o *SendingDomainIdentityScoreChecks) HasSpf() bool`

HasSpf returns a boolean if a field has been set.

### GetDmarc

`func (o *SendingDomainIdentityScoreChecks) GetDmarc() SendingDomainIdentityScoreChecksDmarc`

GetDmarc returns the Dmarc field if non-nil, zero value otherwise.

### GetDmarcOk

`func (o *SendingDomainIdentityScoreChecks) GetDmarcOk() (*SendingDomainIdentityScoreChecksDmarc, bool)`

GetDmarcOk returns a tuple with the Dmarc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDmarc

`func (o *SendingDomainIdentityScoreChecks) SetDmarc(v SendingDomainIdentityScoreChecksDmarc)`

SetDmarc sets Dmarc field to given value.

### HasDmarc

`func (o *SendingDomainIdentityScoreChecks) HasDmarc() bool`

HasDmarc returns a boolean if a field has been set.

### GetMx

`func (o *SendingDomainIdentityScoreChecks) GetMx() SendingDomainIdentityScoreChecksDkim`

GetMx returns the Mx field if non-nil, zero value otherwise.

### GetMxOk

`func (o *SendingDomainIdentityScoreChecks) GetMxOk() (*SendingDomainIdentityScoreChecksDkim, bool)`

GetMxOk returns a tuple with the Mx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMx

`func (o *SendingDomainIdentityScoreChecks) SetMx(v SendingDomainIdentityScoreChecksDkim)`

SetMx sets Mx field to given value.

### HasMx

`func (o *SendingDomainIdentityScoreChecks) HasMx() bool`

HasMx returns a boolean if a field has been set.

### GetReturnPath

`func (o *SendingDomainIdentityScoreChecks) GetReturnPath() SendingDomainIdentityScoreChecksDkim`

GetReturnPath returns the ReturnPath field if non-nil, zero value otherwise.

### GetReturnPathOk

`func (o *SendingDomainIdentityScoreChecks) GetReturnPathOk() (*SendingDomainIdentityScoreChecksDkim, bool)`

GetReturnPathOk returns a tuple with the ReturnPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnPath

`func (o *SendingDomainIdentityScoreChecks) SetReturnPath(v SendingDomainIdentityScoreChecksDkim)`

SetReturnPath sets ReturnPath field to given value.

### HasReturnPath

`func (o *SendingDomainIdentityScoreChecks) HasReturnPath() bool`

HasReturnPath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


