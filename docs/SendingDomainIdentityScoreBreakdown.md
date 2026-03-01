# SendingDomainIdentityScoreBreakdown

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dkim** | Pointer to [**IdentityScoreCheck**](IdentityScoreCheck.md) |  | [optional] 
**Spf** | Pointer to [**IdentityScoreCheck**](IdentityScoreCheck.md) |  | [optional] 
**Dmarc** | Pointer to [**IdentityScoreCheck**](IdentityScoreCheck.md) |  | [optional] 
**Mx** | Pointer to [**IdentityScoreCheck**](IdentityScoreCheck.md) |  | [optional] 
**ReturnPath** | Pointer to [**IdentityScoreCheck**](IdentityScoreCheck.md) |  | [optional] 
**Bimi** | Pointer to [**IdentityScoreCheck**](IdentityScoreCheck.md) |  | [optional] 

## Methods

### NewSendingDomainIdentityScoreBreakdown

`func NewSendingDomainIdentityScoreBreakdown() *SendingDomainIdentityScoreBreakdown`

NewSendingDomainIdentityScoreBreakdown instantiates a new SendingDomainIdentityScoreBreakdown object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendingDomainIdentityScoreBreakdownWithDefaults

`func NewSendingDomainIdentityScoreBreakdownWithDefaults() *SendingDomainIdentityScoreBreakdown`

NewSendingDomainIdentityScoreBreakdownWithDefaults instantiates a new SendingDomainIdentityScoreBreakdown object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDkim

`func (o *SendingDomainIdentityScoreBreakdown) GetDkim() IdentityScoreCheck`

GetDkim returns the Dkim field if non-nil, zero value otherwise.

### GetDkimOk

`func (o *SendingDomainIdentityScoreBreakdown) GetDkimOk() (*IdentityScoreCheck, bool)`

GetDkimOk returns a tuple with the Dkim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDkim

`func (o *SendingDomainIdentityScoreBreakdown) SetDkim(v IdentityScoreCheck)`

SetDkim sets Dkim field to given value.

### HasDkim

`func (o *SendingDomainIdentityScoreBreakdown) HasDkim() bool`

HasDkim returns a boolean if a field has been set.

### GetSpf

`func (o *SendingDomainIdentityScoreBreakdown) GetSpf() IdentityScoreCheck`

GetSpf returns the Spf field if non-nil, zero value otherwise.

### GetSpfOk

`func (o *SendingDomainIdentityScoreBreakdown) GetSpfOk() (*IdentityScoreCheck, bool)`

GetSpfOk returns a tuple with the Spf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpf

`func (o *SendingDomainIdentityScoreBreakdown) SetSpf(v IdentityScoreCheck)`

SetSpf sets Spf field to given value.

### HasSpf

`func (o *SendingDomainIdentityScoreBreakdown) HasSpf() bool`

HasSpf returns a boolean if a field has been set.

### GetDmarc

`func (o *SendingDomainIdentityScoreBreakdown) GetDmarc() IdentityScoreCheck`

GetDmarc returns the Dmarc field if non-nil, zero value otherwise.

### GetDmarcOk

`func (o *SendingDomainIdentityScoreBreakdown) GetDmarcOk() (*IdentityScoreCheck, bool)`

GetDmarcOk returns a tuple with the Dmarc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDmarc

`func (o *SendingDomainIdentityScoreBreakdown) SetDmarc(v IdentityScoreCheck)`

SetDmarc sets Dmarc field to given value.

### HasDmarc

`func (o *SendingDomainIdentityScoreBreakdown) HasDmarc() bool`

HasDmarc returns a boolean if a field has been set.

### GetMx

`func (o *SendingDomainIdentityScoreBreakdown) GetMx() IdentityScoreCheck`

GetMx returns the Mx field if non-nil, zero value otherwise.

### GetMxOk

`func (o *SendingDomainIdentityScoreBreakdown) GetMxOk() (*IdentityScoreCheck, bool)`

GetMxOk returns a tuple with the Mx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMx

`func (o *SendingDomainIdentityScoreBreakdown) SetMx(v IdentityScoreCheck)`

SetMx sets Mx field to given value.

### HasMx

`func (o *SendingDomainIdentityScoreBreakdown) HasMx() bool`

HasMx returns a boolean if a field has been set.

### GetReturnPath

`func (o *SendingDomainIdentityScoreBreakdown) GetReturnPath() IdentityScoreCheck`

GetReturnPath returns the ReturnPath field if non-nil, zero value otherwise.

### GetReturnPathOk

`func (o *SendingDomainIdentityScoreBreakdown) GetReturnPathOk() (*IdentityScoreCheck, bool)`

GetReturnPathOk returns a tuple with the ReturnPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnPath

`func (o *SendingDomainIdentityScoreBreakdown) SetReturnPath(v IdentityScoreCheck)`

SetReturnPath sets ReturnPath field to given value.

### HasReturnPath

`func (o *SendingDomainIdentityScoreBreakdown) HasReturnPath() bool`

HasReturnPath returns a boolean if a field has been set.

### GetBimi

`func (o *SendingDomainIdentityScoreBreakdown) GetBimi() IdentityScoreCheck`

GetBimi returns the Bimi field if non-nil, zero value otherwise.

### GetBimiOk

`func (o *SendingDomainIdentityScoreBreakdown) GetBimiOk() (*IdentityScoreCheck, bool)`

GetBimiOk returns a tuple with the Bimi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBimi

`func (o *SendingDomainIdentityScoreBreakdown) SetBimi(v IdentityScoreCheck)`

SetBimi sets Bimi field to given value.

### HasBimi

`func (o *SendingDomainIdentityScoreBreakdown) HasBimi() bool`

HasBimi returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


