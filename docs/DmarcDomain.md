# DmarcDomain

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Domain UUID | [optional] 
**Domain** | Pointer to **string** | Domain name | [optional] 
**ReportingAddress** | Pointer to **string** | DMARC aggregate report receiving address | [optional] 
**DnsVerified** | Pointer to **bool** | Whether DNS record has been verified | [optional] 
**DnsVerifiedAt** | Pointer to **NullableTime** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewDmarcDomain

`func NewDmarcDomain() *DmarcDomain`

NewDmarcDomain instantiates a new DmarcDomain object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDmarcDomainWithDefaults

`func NewDmarcDomainWithDefaults() *DmarcDomain`

NewDmarcDomainWithDefaults instantiates a new DmarcDomain object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DmarcDomain) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DmarcDomain) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DmarcDomain) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DmarcDomain) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDomain

`func (o *DmarcDomain) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *DmarcDomain) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *DmarcDomain) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *DmarcDomain) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetReportingAddress

`func (o *DmarcDomain) GetReportingAddress() string`

GetReportingAddress returns the ReportingAddress field if non-nil, zero value otherwise.

### GetReportingAddressOk

`func (o *DmarcDomain) GetReportingAddressOk() (*string, bool)`

GetReportingAddressOk returns a tuple with the ReportingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportingAddress

`func (o *DmarcDomain) SetReportingAddress(v string)`

SetReportingAddress sets ReportingAddress field to given value.

### HasReportingAddress

`func (o *DmarcDomain) HasReportingAddress() bool`

HasReportingAddress returns a boolean if a field has been set.

### GetDnsVerified

`func (o *DmarcDomain) GetDnsVerified() bool`

GetDnsVerified returns the DnsVerified field if non-nil, zero value otherwise.

### GetDnsVerifiedOk

`func (o *DmarcDomain) GetDnsVerifiedOk() (*bool, bool)`

GetDnsVerifiedOk returns a tuple with the DnsVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsVerified

`func (o *DmarcDomain) SetDnsVerified(v bool)`

SetDnsVerified sets DnsVerified field to given value.

### HasDnsVerified

`func (o *DmarcDomain) HasDnsVerified() bool`

HasDnsVerified returns a boolean if a field has been set.

### GetDnsVerifiedAt

`func (o *DmarcDomain) GetDnsVerifiedAt() time.Time`

GetDnsVerifiedAt returns the DnsVerifiedAt field if non-nil, zero value otherwise.

### GetDnsVerifiedAtOk

`func (o *DmarcDomain) GetDnsVerifiedAtOk() (*time.Time, bool)`

GetDnsVerifiedAtOk returns a tuple with the DnsVerifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsVerifiedAt

`func (o *DmarcDomain) SetDnsVerifiedAt(v time.Time)`

SetDnsVerifiedAt sets DnsVerifiedAt field to given value.

### HasDnsVerifiedAt

`func (o *DmarcDomain) HasDnsVerifiedAt() bool`

HasDnsVerifiedAt returns a boolean if a field has been set.

### SetDnsVerifiedAtNil

`func (o *DmarcDomain) SetDnsVerifiedAtNil(b bool)`

 SetDnsVerifiedAtNil sets the value for DnsVerifiedAt to be an explicit nil

### UnsetDnsVerifiedAt
`func (o *DmarcDomain) UnsetDnsVerifiedAt()`

UnsetDnsVerifiedAt ensures that no value is present for DnsVerifiedAt, not even an explicit nil
### GetCreatedAt

`func (o *DmarcDomain) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DmarcDomain) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DmarcDomain) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DmarcDomain) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


