# SendingDomain

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Domain UUID | [optional] 
**Domain** | Pointer to **string** | Domain name | [optional] 
**DomainType** | Pointer to **string** | Domain type (ns_delegated) | [optional] 
**Status** | Pointer to **string** | Domain verification status | [optional] 
**DkimSelector** | Pointer to **string** | DKIM selector (e.g. mo1) | [optional] 
**DnsRecords** | Pointer to [**SendingDomainDnsRecords**](SendingDomainDnsRecords.md) |  | [optional] 
**BimiSvgUrl** | Pointer to **NullableString** | BIMI SVG logo URL | [optional] 
**BimiVmcUrl** | Pointer to **NullableString** | BIMI VMC certificate URL | [optional] 
**BimiEnabled** | Pointer to **bool** | Whether BIMI is enabled | [optional] 
**ForwardRepliesTo** | Pointer to **NullableString** | Reply forwarding address | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewSendingDomain

`func NewSendingDomain() *SendingDomain`

NewSendingDomain instantiates a new SendingDomain object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendingDomainWithDefaults

`func NewSendingDomainWithDefaults() *SendingDomain`

NewSendingDomainWithDefaults instantiates a new SendingDomain object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SendingDomain) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SendingDomain) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SendingDomain) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SendingDomain) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDomain

`func (o *SendingDomain) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *SendingDomain) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *SendingDomain) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *SendingDomain) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetDomainType

`func (o *SendingDomain) GetDomainType() string`

GetDomainType returns the DomainType field if non-nil, zero value otherwise.

### GetDomainTypeOk

`func (o *SendingDomain) GetDomainTypeOk() (*string, bool)`

GetDomainTypeOk returns a tuple with the DomainType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainType

`func (o *SendingDomain) SetDomainType(v string)`

SetDomainType sets DomainType field to given value.

### HasDomainType

`func (o *SendingDomain) HasDomainType() bool`

HasDomainType returns a boolean if a field has been set.

### GetStatus

`func (o *SendingDomain) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SendingDomain) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SendingDomain) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SendingDomain) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDkimSelector

`func (o *SendingDomain) GetDkimSelector() string`

GetDkimSelector returns the DkimSelector field if non-nil, zero value otherwise.

### GetDkimSelectorOk

`func (o *SendingDomain) GetDkimSelectorOk() (*string, bool)`

GetDkimSelectorOk returns a tuple with the DkimSelector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDkimSelector

`func (o *SendingDomain) SetDkimSelector(v string)`

SetDkimSelector sets DkimSelector field to given value.

### HasDkimSelector

`func (o *SendingDomain) HasDkimSelector() bool`

HasDkimSelector returns a boolean if a field has been set.

### GetDnsRecords

`func (o *SendingDomain) GetDnsRecords() SendingDomainDnsRecords`

GetDnsRecords returns the DnsRecords field if non-nil, zero value otherwise.

### GetDnsRecordsOk

`func (o *SendingDomain) GetDnsRecordsOk() (*SendingDomainDnsRecords, bool)`

GetDnsRecordsOk returns a tuple with the DnsRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsRecords

`func (o *SendingDomain) SetDnsRecords(v SendingDomainDnsRecords)`

SetDnsRecords sets DnsRecords field to given value.

### HasDnsRecords

`func (o *SendingDomain) HasDnsRecords() bool`

HasDnsRecords returns a boolean if a field has been set.

### GetBimiSvgUrl

`func (o *SendingDomain) GetBimiSvgUrl() string`

GetBimiSvgUrl returns the BimiSvgUrl field if non-nil, zero value otherwise.

### GetBimiSvgUrlOk

`func (o *SendingDomain) GetBimiSvgUrlOk() (*string, bool)`

GetBimiSvgUrlOk returns a tuple with the BimiSvgUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBimiSvgUrl

`func (o *SendingDomain) SetBimiSvgUrl(v string)`

SetBimiSvgUrl sets BimiSvgUrl field to given value.

### HasBimiSvgUrl

`func (o *SendingDomain) HasBimiSvgUrl() bool`

HasBimiSvgUrl returns a boolean if a field has been set.

### SetBimiSvgUrlNil

`func (o *SendingDomain) SetBimiSvgUrlNil(b bool)`

 SetBimiSvgUrlNil sets the value for BimiSvgUrl to be an explicit nil

### UnsetBimiSvgUrl
`func (o *SendingDomain) UnsetBimiSvgUrl()`

UnsetBimiSvgUrl ensures that no value is present for BimiSvgUrl, not even an explicit nil
### GetBimiVmcUrl

`func (o *SendingDomain) GetBimiVmcUrl() string`

GetBimiVmcUrl returns the BimiVmcUrl field if non-nil, zero value otherwise.

### GetBimiVmcUrlOk

`func (o *SendingDomain) GetBimiVmcUrlOk() (*string, bool)`

GetBimiVmcUrlOk returns a tuple with the BimiVmcUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBimiVmcUrl

`func (o *SendingDomain) SetBimiVmcUrl(v string)`

SetBimiVmcUrl sets BimiVmcUrl field to given value.

### HasBimiVmcUrl

`func (o *SendingDomain) HasBimiVmcUrl() bool`

HasBimiVmcUrl returns a boolean if a field has been set.

### SetBimiVmcUrlNil

`func (o *SendingDomain) SetBimiVmcUrlNil(b bool)`

 SetBimiVmcUrlNil sets the value for BimiVmcUrl to be an explicit nil

### UnsetBimiVmcUrl
`func (o *SendingDomain) UnsetBimiVmcUrl()`

UnsetBimiVmcUrl ensures that no value is present for BimiVmcUrl, not even an explicit nil
### GetBimiEnabled

`func (o *SendingDomain) GetBimiEnabled() bool`

GetBimiEnabled returns the BimiEnabled field if non-nil, zero value otherwise.

### GetBimiEnabledOk

`func (o *SendingDomain) GetBimiEnabledOk() (*bool, bool)`

GetBimiEnabledOk returns a tuple with the BimiEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBimiEnabled

`func (o *SendingDomain) SetBimiEnabled(v bool)`

SetBimiEnabled sets BimiEnabled field to given value.

### HasBimiEnabled

`func (o *SendingDomain) HasBimiEnabled() bool`

HasBimiEnabled returns a boolean if a field has been set.

### GetForwardRepliesTo

`func (o *SendingDomain) GetForwardRepliesTo() string`

GetForwardRepliesTo returns the ForwardRepliesTo field if non-nil, zero value otherwise.

### GetForwardRepliesToOk

`func (o *SendingDomain) GetForwardRepliesToOk() (*string, bool)`

GetForwardRepliesToOk returns a tuple with the ForwardRepliesTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardRepliesTo

`func (o *SendingDomain) SetForwardRepliesTo(v string)`

SetForwardRepliesTo sets ForwardRepliesTo field to given value.

### HasForwardRepliesTo

`func (o *SendingDomain) HasForwardRepliesTo() bool`

HasForwardRepliesTo returns a boolean if a field has been set.

### SetForwardRepliesToNil

`func (o *SendingDomain) SetForwardRepliesToNil(b bool)`

 SetForwardRepliesToNil sets the value for ForwardRepliesTo to be an explicit nil

### UnsetForwardRepliesTo
`func (o *SendingDomain) UnsetForwardRepliesTo()`

UnsetForwardRepliesTo ensures that no value is present for ForwardRepliesTo, not even an explicit nil
### GetCreatedAt

`func (o *SendingDomain) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SendingDomain) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SendingDomain) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *SendingDomain) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *SendingDomain) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *SendingDomain) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *SendingDomain) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *SendingDomain) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


