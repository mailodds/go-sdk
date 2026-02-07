# ValidationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SchemaVersion** | Pointer to **string** |  | [optional] 
**Email** | **string** |  | 
**Status** | **string** | Validation status | 
**SubStatus** | Pointer to **string** | Detailed status reason | [optional] 
**Action** | **string** | Recommended action | 
**Domain** | Pointer to **string** |  | [optional] 
**MxFound** | Pointer to **bool** |  | [optional] 
**SmtpCheck** | Pointer to **bool** |  | [optional] 
**Disposable** | Pointer to **bool** |  | [optional] 
**RoleAccount** | Pointer to **bool** |  | [optional] 
**FreeProvider** | Pointer to **bool** |  | [optional] 
**SuppressionMatch** | Pointer to [**ValidationResponseSuppressionMatch**](ValidationResponseSuppressionMatch.md) |  | [optional] 

## Methods

### NewValidationResponse

`func NewValidationResponse(email string, status string, action string, ) *ValidationResponse`

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

### HasSchemaVersion

`func (o *ValidationResponse) HasSchemaVersion() bool`

HasSchemaVersion returns a boolean if a field has been set.

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

### HasDomain

`func (o *ValidationResponse) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

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

### HasMxFound

`func (o *ValidationResponse) HasMxFound() bool`

HasMxFound returns a boolean if a field has been set.

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

### HasDisposable

`func (o *ValidationResponse) HasDisposable() bool`

HasDisposable returns a boolean if a field has been set.

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

### HasRoleAccount

`func (o *ValidationResponse) HasRoleAccount() bool`

HasRoleAccount returns a boolean if a field has been set.

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

### HasFreeProvider

`func (o *ValidationResponse) HasFreeProvider() bool`

HasFreeProvider returns a boolean if a field has been set.

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


