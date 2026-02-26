# SendingDomainDnsRecordsNs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | Record type (NS) | [optional] 
**Host** | Pointer to **string** | NS record host (mo.yourdomain.com) | [optional] 
**Targets** | Pointer to **[]string** | NS target servers | [optional] 
**Status** | Pointer to **string** | Verification status | [optional] 
**VerifiedAt** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewSendingDomainDnsRecordsNs

`func NewSendingDomainDnsRecordsNs() *SendingDomainDnsRecordsNs`

NewSendingDomainDnsRecordsNs instantiates a new SendingDomainDnsRecordsNs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendingDomainDnsRecordsNsWithDefaults

`func NewSendingDomainDnsRecordsNsWithDefaults() *SendingDomainDnsRecordsNs`

NewSendingDomainDnsRecordsNsWithDefaults instantiates a new SendingDomainDnsRecordsNs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SendingDomainDnsRecordsNs) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SendingDomainDnsRecordsNs) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SendingDomainDnsRecordsNs) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SendingDomainDnsRecordsNs) HasType() bool`

HasType returns a boolean if a field has been set.

### GetHost

`func (o *SendingDomainDnsRecordsNs) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *SendingDomainDnsRecordsNs) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *SendingDomainDnsRecordsNs) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *SendingDomainDnsRecordsNs) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetTargets

`func (o *SendingDomainDnsRecordsNs) GetTargets() []string`

GetTargets returns the Targets field if non-nil, zero value otherwise.

### GetTargetsOk

`func (o *SendingDomainDnsRecordsNs) GetTargetsOk() (*[]string, bool)`

GetTargetsOk returns a tuple with the Targets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargets

`func (o *SendingDomainDnsRecordsNs) SetTargets(v []string)`

SetTargets sets Targets field to given value.

### HasTargets

`func (o *SendingDomainDnsRecordsNs) HasTargets() bool`

HasTargets returns a boolean if a field has been set.

### GetStatus

`func (o *SendingDomainDnsRecordsNs) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SendingDomainDnsRecordsNs) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SendingDomainDnsRecordsNs) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SendingDomainDnsRecordsNs) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetVerifiedAt

`func (o *SendingDomainDnsRecordsNs) GetVerifiedAt() time.Time`

GetVerifiedAt returns the VerifiedAt field if non-nil, zero value otherwise.

### GetVerifiedAtOk

`func (o *SendingDomainDnsRecordsNs) GetVerifiedAtOk() (*time.Time, bool)`

GetVerifiedAtOk returns a tuple with the VerifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifiedAt

`func (o *SendingDomainDnsRecordsNs) SetVerifiedAt(v time.Time)`

SetVerifiedAt sets VerifiedAt field to given value.

### HasVerifiedAt

`func (o *SendingDomainDnsRecordsNs) HasVerifiedAt() bool`

HasVerifiedAt returns a boolean if a field has been set.

### SetVerifiedAtNil

`func (o *SendingDomainDnsRecordsNs) SetVerifiedAtNil(b bool)`

 SetVerifiedAtNil sets the value for VerifiedAt to be an explicit nil

### UnsetVerifiedAt
`func (o *SendingDomainDnsRecordsNs) UnsetVerifiedAt()`

UnsetVerifiedAt ensures that no value is present for VerifiedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


