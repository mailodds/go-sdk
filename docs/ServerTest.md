# ServerTest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Test UUID | [optional] 
**Domain** | Pointer to **string** | Tested domain | [optional] 
**Status** | Pointer to **string** | Test status | [optional] 
**MxRecords** | Pointer to [**[]ServerTestMxRecordsInner**](ServerTestMxRecordsInner.md) |  | [optional] 
**SmtpCheck** | Pointer to [**ServerTestSmtpCheck**](ServerTestSmtpCheck.md) |  | [optional] 
**DnsChecks** | Pointer to [**ServerTestDnsChecks**](ServerTestDnsChecks.md) |  | [optional] 
**Score** | Pointer to **int32** | Overall server configuration score (0-100) | [optional] 
**Recommendations** | Pointer to **[]string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewServerTest

`func NewServerTest() *ServerTest`

NewServerTest instantiates a new ServerTest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerTestWithDefaults

`func NewServerTestWithDefaults() *ServerTest`

NewServerTestWithDefaults instantiates a new ServerTest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ServerTest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServerTest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServerTest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ServerTest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDomain

`func (o *ServerTest) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *ServerTest) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *ServerTest) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *ServerTest) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetStatus

`func (o *ServerTest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ServerTest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ServerTest) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ServerTest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMxRecords

`func (o *ServerTest) GetMxRecords() []ServerTestMxRecordsInner`

GetMxRecords returns the MxRecords field if non-nil, zero value otherwise.

### GetMxRecordsOk

`func (o *ServerTest) GetMxRecordsOk() (*[]ServerTestMxRecordsInner, bool)`

GetMxRecordsOk returns a tuple with the MxRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMxRecords

`func (o *ServerTest) SetMxRecords(v []ServerTestMxRecordsInner)`

SetMxRecords sets MxRecords field to given value.

### HasMxRecords

`func (o *ServerTest) HasMxRecords() bool`

HasMxRecords returns a boolean if a field has been set.

### GetSmtpCheck

`func (o *ServerTest) GetSmtpCheck() ServerTestSmtpCheck`

GetSmtpCheck returns the SmtpCheck field if non-nil, zero value otherwise.

### GetSmtpCheckOk

`func (o *ServerTest) GetSmtpCheckOk() (*ServerTestSmtpCheck, bool)`

GetSmtpCheckOk returns a tuple with the SmtpCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtpCheck

`func (o *ServerTest) SetSmtpCheck(v ServerTestSmtpCheck)`

SetSmtpCheck sets SmtpCheck field to given value.

### HasSmtpCheck

`func (o *ServerTest) HasSmtpCheck() bool`

HasSmtpCheck returns a boolean if a field has been set.

### GetDnsChecks

`func (o *ServerTest) GetDnsChecks() ServerTestDnsChecks`

GetDnsChecks returns the DnsChecks field if non-nil, zero value otherwise.

### GetDnsChecksOk

`func (o *ServerTest) GetDnsChecksOk() (*ServerTestDnsChecks, bool)`

GetDnsChecksOk returns a tuple with the DnsChecks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsChecks

`func (o *ServerTest) SetDnsChecks(v ServerTestDnsChecks)`

SetDnsChecks sets DnsChecks field to given value.

### HasDnsChecks

`func (o *ServerTest) HasDnsChecks() bool`

HasDnsChecks returns a boolean if a field has been set.

### GetScore

`func (o *ServerTest) GetScore() int32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *ServerTest) GetScoreOk() (*int32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *ServerTest) SetScore(v int32)`

SetScore sets Score field to given value.

### HasScore

`func (o *ServerTest) HasScore() bool`

HasScore returns a boolean if a field has been set.

### GetRecommendations

`func (o *ServerTest) GetRecommendations() []string`

GetRecommendations returns the Recommendations field if non-nil, zero value otherwise.

### GetRecommendationsOk

`func (o *ServerTest) GetRecommendationsOk() (*[]string, bool)`

GetRecommendationsOk returns a tuple with the Recommendations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommendations

`func (o *ServerTest) SetRecommendations(v []string)`

SetRecommendations sets Recommendations field to given value.

### HasRecommendations

`func (o *ServerTest) HasRecommendations() bool`

HasRecommendations returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ServerTest) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ServerTest) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ServerTest) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ServerTest) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


