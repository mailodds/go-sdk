# GetDmarcDomain200ResponseDomain

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Domain UUID | [optional] 
**Domain** | Pointer to **string** | Domain name | [optional] 
**ReportingAddress** | Pointer to **string** | DMARC aggregate report receiving address | [optional] 
**DnsVerified** | Pointer to **bool** | Whether DNS record has been verified | [optional] 
**DnsVerifiedAt** | Pointer to **NullableTime** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Summary** | Pointer to [**GetDmarcDomain200ResponseDomainAllOfSummary**](GetDmarcDomain200ResponseDomainAllOfSummary.md) |  | [optional] 

## Methods

### NewGetDmarcDomain200ResponseDomain

`func NewGetDmarcDomain200ResponseDomain() *GetDmarcDomain200ResponseDomain`

NewGetDmarcDomain200ResponseDomain instantiates a new GetDmarcDomain200ResponseDomain object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDmarcDomain200ResponseDomainWithDefaults

`func NewGetDmarcDomain200ResponseDomainWithDefaults() *GetDmarcDomain200ResponseDomain`

NewGetDmarcDomain200ResponseDomainWithDefaults instantiates a new GetDmarcDomain200ResponseDomain object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetDmarcDomain200ResponseDomain) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetDmarcDomain200ResponseDomain) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetDmarcDomain200ResponseDomain) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GetDmarcDomain200ResponseDomain) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDomain

`func (o *GetDmarcDomain200ResponseDomain) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *GetDmarcDomain200ResponseDomain) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *GetDmarcDomain200ResponseDomain) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *GetDmarcDomain200ResponseDomain) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetReportingAddress

`func (o *GetDmarcDomain200ResponseDomain) GetReportingAddress() string`

GetReportingAddress returns the ReportingAddress field if non-nil, zero value otherwise.

### GetReportingAddressOk

`func (o *GetDmarcDomain200ResponseDomain) GetReportingAddressOk() (*string, bool)`

GetReportingAddressOk returns a tuple with the ReportingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportingAddress

`func (o *GetDmarcDomain200ResponseDomain) SetReportingAddress(v string)`

SetReportingAddress sets ReportingAddress field to given value.

### HasReportingAddress

`func (o *GetDmarcDomain200ResponseDomain) HasReportingAddress() bool`

HasReportingAddress returns a boolean if a field has been set.

### GetDnsVerified

`func (o *GetDmarcDomain200ResponseDomain) GetDnsVerified() bool`

GetDnsVerified returns the DnsVerified field if non-nil, zero value otherwise.

### GetDnsVerifiedOk

`func (o *GetDmarcDomain200ResponseDomain) GetDnsVerifiedOk() (*bool, bool)`

GetDnsVerifiedOk returns a tuple with the DnsVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsVerified

`func (o *GetDmarcDomain200ResponseDomain) SetDnsVerified(v bool)`

SetDnsVerified sets DnsVerified field to given value.

### HasDnsVerified

`func (o *GetDmarcDomain200ResponseDomain) HasDnsVerified() bool`

HasDnsVerified returns a boolean if a field has been set.

### GetDnsVerifiedAt

`func (o *GetDmarcDomain200ResponseDomain) GetDnsVerifiedAt() time.Time`

GetDnsVerifiedAt returns the DnsVerifiedAt field if non-nil, zero value otherwise.

### GetDnsVerifiedAtOk

`func (o *GetDmarcDomain200ResponseDomain) GetDnsVerifiedAtOk() (*time.Time, bool)`

GetDnsVerifiedAtOk returns a tuple with the DnsVerifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsVerifiedAt

`func (o *GetDmarcDomain200ResponseDomain) SetDnsVerifiedAt(v time.Time)`

SetDnsVerifiedAt sets DnsVerifiedAt field to given value.

### HasDnsVerifiedAt

`func (o *GetDmarcDomain200ResponseDomain) HasDnsVerifiedAt() bool`

HasDnsVerifiedAt returns a boolean if a field has been set.

### SetDnsVerifiedAtNil

`func (o *GetDmarcDomain200ResponseDomain) SetDnsVerifiedAtNil(b bool)`

 SetDnsVerifiedAtNil sets the value for DnsVerifiedAt to be an explicit nil

### UnsetDnsVerifiedAt
`func (o *GetDmarcDomain200ResponseDomain) UnsetDnsVerifiedAt()`

UnsetDnsVerifiedAt ensures that no value is present for DnsVerifiedAt, not even an explicit nil
### GetCreatedAt

`func (o *GetDmarcDomain200ResponseDomain) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GetDmarcDomain200ResponseDomain) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GetDmarcDomain200ResponseDomain) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GetDmarcDomain200ResponseDomain) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetSummary

`func (o *GetDmarcDomain200ResponseDomain) GetSummary() GetDmarcDomain200ResponseDomainAllOfSummary`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *GetDmarcDomain200ResponseDomain) GetSummaryOk() (*GetDmarcDomain200ResponseDomainAllOfSummary, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *GetDmarcDomain200ResponseDomain) SetSummary(v GetDmarcDomain200ResponseDomainAllOfSummary)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *GetDmarcDomain200ResponseDomain) HasSummary() bool`

HasSummary returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


