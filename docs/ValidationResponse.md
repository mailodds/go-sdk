# ValidationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SchemaVersion** | **string** |  | 
**RequestId** | Pointer to **string** | Unique request identifier | [optional] 
**Email** | **string** |  | 
**Status** | **string** | Validation status | 
**Action** | **string** | Recommended action | 
**SubStatus** | Pointer to **string** | Detailed status reason. Omitted when none. | [optional] 
**Domain** | **string** |  | 
**MxFound** | **bool** | Whether MX records were found for the domain | 
**MxHost** | Pointer to **string** | Primary MX hostname. Omitted when MX not resolved. | [optional] 
**SmtpCheck** | Pointer to **bool** | Whether SMTP verification passed. Omitted when SMTP not checked. | [optional] 
**CatchAll** | Pointer to **bool** | Whether domain is catch-all. Omitted when SMTP not checked. | [optional] 
**Disposable** | **bool** | Whether domain is a known disposable email provider | 
**RoleAccount** | **bool** | Whether address is a role account (e.g., info@, admin@) | 
**FreeProvider** | **bool** | Whether domain is a known free email provider (e.g., gmail.com) | 
**Depth** | **string** | Validation depth used for this check | 
**ProcessedAt** | **time.Time** | ISO 8601 timestamp of validation | 
**SuggestedEmail** | Pointer to **string** | Typo correction suggestion. Omitted when no typo detected. | [optional] 
**RetryAfterMs** | Pointer to **int32** | Suggested retry delay in milliseconds. Present only for retry_later action. | [optional] 
**HasSpf** | Pointer to **bool** | Whether the domain has an SPF record. Omitted for standard depth. | [optional] 
**HasDmarc** | Pointer to **bool** | Whether the domain has a DMARC record. Omitted for standard depth. | [optional] 
**DmarcPolicy** | Pointer to **string** | The domain&#39;s DMARC policy. Omitted when no DMARC record found. | [optional] 
**DnsblListed** | Pointer to **bool** | Whether the domain&#39;s MX IP is on a DNS blocklist (Spamhaus ZEN). Omitted for standard depth. | [optional] 
**SuppressionMatch** | Pointer to [**ValidationResponseSuppressionMatch**](ValidationResponseSuppressionMatch.md) |  | [optional] 
**PolicyApplied** | Pointer to [**ValidationResponsePolicyApplied**](ValidationResponsePolicyApplied.md) |  | [optional] 

## Methods

### NewValidationResponse

`func NewValidationResponse(schemaVersion string, email string, status string, action string, domain string, mxFound bool, disposable bool, roleAccount bool, freeProvider bool, depth string, processedAt time.Time, ) *ValidationResponse`

NewValidationResponse instantiates a new ValidationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidationResponseWithDefaults

`func NewValidationResponseWithDefaults() *ValidationResponse`

NewValidationResponseWithDefaults instantiates a new ValidationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemaVersion

`func (o *ValidationResponse) GetSchemaVersion() string`

GetSchemaVersion returns the SchemaVersion field if non-nil, zero value otherwise.

### GetSchemaVersionOk

`func (o *ValidationResponse) GetSchemaVersionOk() (*string, bool)`

GetSchemaVersionOk returns a tuple with the SchemaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaVersion

`func (o *ValidationResponse) SetSchemaVersion(v string)`

SetSchemaVersion sets SchemaVersion field to given value.


### GetRequestId

`func (o *ValidationResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *ValidationResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *ValidationResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *ValidationResponse) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetEmail

`func (o *ValidationResponse) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ValidationResponse) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ValidationResponse) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetStatus

`func (o *ValidationResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ValidationResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ValidationResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetAction

`func (o *ValidationResponse) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *ValidationResponse) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *ValidationResponse) SetAction(v string)`

SetAction sets Action field to given value.


### GetSubStatus

`func (o *ValidationResponse) GetSubStatus() string`

GetSubStatus returns the SubStatus field if non-nil, zero value otherwise.

### GetSubStatusOk

`func (o *ValidationResponse) GetSubStatusOk() (*string, bool)`

GetSubStatusOk returns a tuple with the SubStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubStatus

`func (o *ValidationResponse) SetSubStatus(v string)`

SetSubStatus sets SubStatus field to given value.

### HasSubStatus

`func (o *ValidationResponse) HasSubStatus() bool`

HasSubStatus returns a boolean if a field has been set.

### GetDomain

`func (o *ValidationResponse) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *ValidationResponse) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *ValidationResponse) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetMxFound

`func (o *ValidationResponse) GetMxFound() bool`

GetMxFound returns the MxFound field if non-nil, zero value otherwise.

### GetMxFoundOk

`func (o *ValidationResponse) GetMxFoundOk() (*bool, bool)`

GetMxFoundOk returns a tuple with the MxFound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMxFound

`func (o *ValidationResponse) SetMxFound(v bool)`

SetMxFound sets MxFound field to given value.


### GetMxHost

`func (o *ValidationResponse) GetMxHost() string`

GetMxHost returns the MxHost field if non-nil, zero value otherwise.

### GetMxHostOk

`func (o *ValidationResponse) GetMxHostOk() (*string, bool)`

GetMxHostOk returns a tuple with the MxHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMxHost

`func (o *ValidationResponse) SetMxHost(v string)`

SetMxHost sets MxHost field to given value.

### HasMxHost

`func (o *ValidationResponse) HasMxHost() bool`

HasMxHost returns a boolean if a field has been set.

### GetSmtpCheck

`func (o *ValidationResponse) GetSmtpCheck() bool`

GetSmtpCheck returns the SmtpCheck field if non-nil, zero value otherwise.

### GetSmtpCheckOk

`func (o *ValidationResponse) GetSmtpCheckOk() (*bool, bool)`

GetSmtpCheckOk returns a tuple with the SmtpCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtpCheck

`func (o *ValidationResponse) SetSmtpCheck(v bool)`

SetSmtpCheck sets SmtpCheck field to given value.

### HasSmtpCheck

`func (o *ValidationResponse) HasSmtpCheck() bool`

HasSmtpCheck returns a boolean if a field has been set.

### GetCatchAll

`func (o *ValidationResponse) GetCatchAll() bool`

GetCatchAll returns the CatchAll field if non-nil, zero value otherwise.

### GetCatchAllOk

`func (o *ValidationResponse) GetCatchAllOk() (*bool, bool)`

GetCatchAllOk returns a tuple with the CatchAll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatchAll

`func (o *ValidationResponse) SetCatchAll(v bool)`

SetCatchAll sets CatchAll field to given value.

### HasCatchAll

`func (o *ValidationResponse) HasCatchAll() bool`

HasCatchAll returns a boolean if a field has been set.

### GetDisposable

`func (o *ValidationResponse) GetDisposable() bool`

GetDisposable returns the Disposable field if non-nil, zero value otherwise.

### GetDisposableOk

`func (o *ValidationResponse) GetDisposableOk() (*bool, bool)`

GetDisposableOk returns a tuple with the Disposable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisposable

`func (o *ValidationResponse) SetDisposable(v bool)`

SetDisposable sets Disposable field to given value.


### GetRoleAccount

`func (o *ValidationResponse) GetRoleAccount() bool`

GetRoleAccount returns the RoleAccount field if non-nil, zero value otherwise.

### GetRoleAccountOk

`func (o *ValidationResponse) GetRoleAccountOk() (*bool, bool)`

GetRoleAccountOk returns a tuple with the RoleAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleAccount

`func (o *ValidationResponse) SetRoleAccount(v bool)`

SetRoleAccount sets RoleAccount field to given value.


### GetFreeProvider

`func (o *ValidationResponse) GetFreeProvider() bool`

GetFreeProvider returns the FreeProvider field if non-nil, zero value otherwise.

### GetFreeProviderOk

`func (o *ValidationResponse) GetFreeProviderOk() (*bool, bool)`

GetFreeProviderOk returns a tuple with the FreeProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeProvider

`func (o *ValidationResponse) SetFreeProvider(v bool)`

SetFreeProvider sets FreeProvider field to given value.


### GetDepth

`func (o *ValidationResponse) GetDepth() string`

GetDepth returns the Depth field if non-nil, zero value otherwise.

### GetDepthOk

`func (o *ValidationResponse) GetDepthOk() (*string, bool)`

GetDepthOk returns a tuple with the Depth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepth

`func (o *ValidationResponse) SetDepth(v string)`

SetDepth sets Depth field to given value.


### GetProcessedAt

`func (o *ValidationResponse) GetProcessedAt() time.Time`

GetProcessedAt returns the ProcessedAt field if non-nil, zero value otherwise.

### GetProcessedAtOk

`func (o *ValidationResponse) GetProcessedAtOk() (*time.Time, bool)`

GetProcessedAtOk returns a tuple with the ProcessedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessedAt

`func (o *ValidationResponse) SetProcessedAt(v time.Time)`

SetProcessedAt sets ProcessedAt field to given value.


### GetSuggestedEmail

`func (o *ValidationResponse) GetSuggestedEmail() string`

GetSuggestedEmail returns the SuggestedEmail field if non-nil, zero value otherwise.

### GetSuggestedEmailOk

`func (o *ValidationResponse) GetSuggestedEmailOk() (*string, bool)`

GetSuggestedEmailOk returns a tuple with the SuggestedEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuggestedEmail

`func (o *ValidationResponse) SetSuggestedEmail(v string)`

SetSuggestedEmail sets SuggestedEmail field to given value.

### HasSuggestedEmail

`func (o *ValidationResponse) HasSuggestedEmail() bool`

HasSuggestedEmail returns a boolean if a field has been set.

### GetRetryAfterMs

`func (o *ValidationResponse) GetRetryAfterMs() int32`

GetRetryAfterMs returns the RetryAfterMs field if non-nil, zero value otherwise.

### GetRetryAfterMsOk

`func (o *ValidationResponse) GetRetryAfterMsOk() (*int32, bool)`

GetRetryAfterMsOk returns a tuple with the RetryAfterMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryAfterMs

`func (o *ValidationResponse) SetRetryAfterMs(v int32)`

SetRetryAfterMs sets RetryAfterMs field to given value.

### HasRetryAfterMs

`func (o *ValidationResponse) HasRetryAfterMs() bool`

HasRetryAfterMs returns a boolean if a field has been set.

### GetHasSpf

`func (o *ValidationResponse) GetHasSpf() bool`

GetHasSpf returns the HasSpf field if non-nil, zero value otherwise.

### GetHasSpfOk

`func (o *ValidationResponse) GetHasSpfOk() (*bool, bool)`

GetHasSpfOk returns a tuple with the HasSpf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasSpf

`func (o *ValidationResponse) SetHasSpf(v bool)`

SetHasSpf sets HasSpf field to given value.

### HasHasSpf

`func (o *ValidationResponse) HasHasSpf() bool`

HasHasSpf returns a boolean if a field has been set.

### GetHasDmarc

`func (o *ValidationResponse) GetHasDmarc() bool`

GetHasDmarc returns the HasDmarc field if non-nil, zero value otherwise.

### GetHasDmarcOk

`func (o *ValidationResponse) GetHasDmarcOk() (*bool, bool)`

GetHasDmarcOk returns a tuple with the HasDmarc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasDmarc

`func (o *ValidationResponse) SetHasDmarc(v bool)`

SetHasDmarc sets HasDmarc field to given value.

### HasHasDmarc

`func (o *ValidationResponse) HasHasDmarc() bool`

HasHasDmarc returns a boolean if a field has been set.

### GetDmarcPolicy

`func (o *ValidationResponse) GetDmarcPolicy() string`

GetDmarcPolicy returns the DmarcPolicy field if non-nil, zero value otherwise.

### GetDmarcPolicyOk

`func (o *ValidationResponse) GetDmarcPolicyOk() (*string, bool)`

GetDmarcPolicyOk returns a tuple with the DmarcPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDmarcPolicy

`func (o *ValidationResponse) SetDmarcPolicy(v string)`

SetDmarcPolicy sets DmarcPolicy field to given value.

### HasDmarcPolicy

`func (o *ValidationResponse) HasDmarcPolicy() bool`

HasDmarcPolicy returns a boolean if a field has been set.

### GetDnsblListed

`func (o *ValidationResponse) GetDnsblListed() bool`

GetDnsblListed returns the DnsblListed field if non-nil, zero value otherwise.

### GetDnsblListedOk

`func (o *ValidationResponse) GetDnsblListedOk() (*bool, bool)`

GetDnsblListedOk returns a tuple with the DnsblListed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsblListed

`func (o *ValidationResponse) SetDnsblListed(v bool)`

SetDnsblListed sets DnsblListed field to given value.

### HasDnsblListed

`func (o *ValidationResponse) HasDnsblListed() bool`

HasDnsblListed returns a boolean if a field has been set.

### GetSuppressionMatch

`func (o *ValidationResponse) GetSuppressionMatch() ValidationResponseSuppressionMatch`

GetSuppressionMatch returns the SuppressionMatch field if non-nil, zero value otherwise.

### GetSuppressionMatchOk

`func (o *ValidationResponse) GetSuppressionMatchOk() (*ValidationResponseSuppressionMatch, bool)`

GetSuppressionMatchOk returns a tuple with the SuppressionMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppressionMatch

`func (o *ValidationResponse) SetSuppressionMatch(v ValidationResponseSuppressionMatch)`

SetSuppressionMatch sets SuppressionMatch field to given value.

### HasSuppressionMatch

`func (o *ValidationResponse) HasSuppressionMatch() bool`

HasSuppressionMatch returns a boolean if a field has been set.

### GetPolicyApplied

`func (o *ValidationResponse) GetPolicyApplied() ValidationResponsePolicyApplied`

GetPolicyApplied returns the PolicyApplied field if non-nil, zero value otherwise.

### GetPolicyAppliedOk

`func (o *ValidationResponse) GetPolicyAppliedOk() (*ValidationResponsePolicyApplied, bool)`

GetPolicyAppliedOk returns a tuple with the PolicyApplied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyApplied

`func (o *ValidationResponse) SetPolicyApplied(v ValidationResponsePolicyApplied)`

SetPolicyApplied sets PolicyApplied field to given value.

### HasPolicyApplied

`func (o *ValidationResponse) HasPolicyApplied() bool`

HasPolicyApplied returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


