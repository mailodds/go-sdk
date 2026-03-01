# ValidationResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** |  | 
**Status** | **string** |  | 
**SubStatus** | Pointer to **string** | Detailed reason. Omitted when none. | [optional] 
**Action** | **string** |  | 
**Domain** | **string** | Email domain | 
**MxHost** | Pointer to **string** | Primary MX hostname. Omitted when not resolved. | [optional] 
**Checks** | Pointer to **map[string]interface{}** | Detailed check results (JSONB). Omitted when not available. | [optional] 
**Suppression** | Pointer to [**ValidationResultSuppression**](ValidationResultSuppression.md) |  | [optional] 
**ProcessedAt** | **time.Time** |  | 

## Methods

### NewValidationResult

`func NewValidationResult(email string, status string, action string, domain string, processedAt time.Time, ) *ValidationResult`

NewValidationResult instantiates a new ValidationResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidationResultWithDefaults

`func NewValidationResultWithDefaults() *ValidationResult`

NewValidationResultWithDefaults instantiates a new ValidationResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *ValidationResult) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ValidationResult) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ValidationResult) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetStatus

`func (o *ValidationResult) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ValidationResult) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ValidationResult) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetSubStatus

`func (o *ValidationResult) GetSubStatus() string`

GetSubStatus returns the SubStatus field if non-nil, zero value otherwise.

### GetSubStatusOk

`func (o *ValidationResult) GetSubStatusOk() (*string, bool)`

GetSubStatusOk returns a tuple with the SubStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubStatus

`func (o *ValidationResult) SetSubStatus(v string)`

SetSubStatus sets SubStatus field to given value.

### HasSubStatus

`func (o *ValidationResult) HasSubStatus() bool`

HasSubStatus returns a boolean if a field has been set.

### GetAction

`func (o *ValidationResult) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *ValidationResult) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *ValidationResult) SetAction(v string)`

SetAction sets Action field to given value.


### GetDomain

`func (o *ValidationResult) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *ValidationResult) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *ValidationResult) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetMxHost

`func (o *ValidationResult) GetMxHost() string`

GetMxHost returns the MxHost field if non-nil, zero value otherwise.

### GetMxHostOk

`func (o *ValidationResult) GetMxHostOk() (*string, bool)`

GetMxHostOk returns a tuple with the MxHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMxHost

`func (o *ValidationResult) SetMxHost(v string)`

SetMxHost sets MxHost field to given value.

### HasMxHost

`func (o *ValidationResult) HasMxHost() bool`

HasMxHost returns a boolean if a field has been set.

### GetChecks

`func (o *ValidationResult) GetChecks() map[string]interface{}`

GetChecks returns the Checks field if non-nil, zero value otherwise.

### GetChecksOk

`func (o *ValidationResult) GetChecksOk() (*map[string]interface{}, bool)`

GetChecksOk returns a tuple with the Checks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecks

`func (o *ValidationResult) SetChecks(v map[string]interface{})`

SetChecks sets Checks field to given value.

### HasChecks

`func (o *ValidationResult) HasChecks() bool`

HasChecks returns a boolean if a field has been set.

### GetSuppression

`func (o *ValidationResult) GetSuppression() ValidationResultSuppression`

GetSuppression returns the Suppression field if non-nil, zero value otherwise.

### GetSuppressionOk

`func (o *ValidationResult) GetSuppressionOk() (*ValidationResultSuppression, bool)`

GetSuppressionOk returns a tuple with the Suppression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppression

`func (o *ValidationResult) SetSuppression(v ValidationResultSuppression)`

SetSuppression sets Suppression field to given value.

### HasSuppression

`func (o *ValidationResult) HasSuppression() bool`

HasSuppression returns a boolean if a field has been set.

### GetProcessedAt

`func (o *ValidationResult) GetProcessedAt() time.Time`

GetProcessedAt returns the ProcessedAt field if non-nil, zero value otherwise.

### GetProcessedAtOk

`func (o *ValidationResult) GetProcessedAtOk() (*time.Time, bool)`

GetProcessedAtOk returns a tuple with the ProcessedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessedAt

`func (o *ValidationResult) SetProcessedAt(v time.Time)`

SetProcessedAt sets ProcessedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


