# GetBounceRecords200ResponseRecordsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**BounceType** | Pointer to **string** |  | [optional] 
**SmtpCode** | Pointer to **int32** |  | [optional] 
**EnhancedStatus** | Pointer to **string** |  | [optional] 
**Diagnostic** | Pointer to **string** |  | [optional] 
**MxHost** | Pointer to **string** |  | [optional] 
**BouncedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewGetBounceRecords200ResponseRecordsInner

`func NewGetBounceRecords200ResponseRecordsInner() *GetBounceRecords200ResponseRecordsInner`

NewGetBounceRecords200ResponseRecordsInner instantiates a new GetBounceRecords200ResponseRecordsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetBounceRecords200ResponseRecordsInnerWithDefaults

`func NewGetBounceRecords200ResponseRecordsInnerWithDefaults() *GetBounceRecords200ResponseRecordsInner`

NewGetBounceRecords200ResponseRecordsInnerWithDefaults instantiates a new GetBounceRecords200ResponseRecordsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetBounceRecords200ResponseRecordsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetBounceRecords200ResponseRecordsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetBounceRecords200ResponseRecordsInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GetBounceRecords200ResponseRecordsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetEmail

`func (o *GetBounceRecords200ResponseRecordsInner) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *GetBounceRecords200ResponseRecordsInner) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *GetBounceRecords200ResponseRecordsInner) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *GetBounceRecords200ResponseRecordsInner) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetBounceType

`func (o *GetBounceRecords200ResponseRecordsInner) GetBounceType() string`

GetBounceType returns the BounceType field if non-nil, zero value otherwise.

### GetBounceTypeOk

`func (o *GetBounceRecords200ResponseRecordsInner) GetBounceTypeOk() (*string, bool)`

GetBounceTypeOk returns a tuple with the BounceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBounceType

`func (o *GetBounceRecords200ResponseRecordsInner) SetBounceType(v string)`

SetBounceType sets BounceType field to given value.

### HasBounceType

`func (o *GetBounceRecords200ResponseRecordsInner) HasBounceType() bool`

HasBounceType returns a boolean if a field has been set.

### GetSmtpCode

`func (o *GetBounceRecords200ResponseRecordsInner) GetSmtpCode() int32`

GetSmtpCode returns the SmtpCode field if non-nil, zero value otherwise.

### GetSmtpCodeOk

`func (o *GetBounceRecords200ResponseRecordsInner) GetSmtpCodeOk() (*int32, bool)`

GetSmtpCodeOk returns a tuple with the SmtpCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtpCode

`func (o *GetBounceRecords200ResponseRecordsInner) SetSmtpCode(v int32)`

SetSmtpCode sets SmtpCode field to given value.

### HasSmtpCode

`func (o *GetBounceRecords200ResponseRecordsInner) HasSmtpCode() bool`

HasSmtpCode returns a boolean if a field has been set.

### GetEnhancedStatus

`func (o *GetBounceRecords200ResponseRecordsInner) GetEnhancedStatus() string`

GetEnhancedStatus returns the EnhancedStatus field if non-nil, zero value otherwise.

### GetEnhancedStatusOk

`func (o *GetBounceRecords200ResponseRecordsInner) GetEnhancedStatusOk() (*string, bool)`

GetEnhancedStatusOk returns a tuple with the EnhancedStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnhancedStatus

`func (o *GetBounceRecords200ResponseRecordsInner) SetEnhancedStatus(v string)`

SetEnhancedStatus sets EnhancedStatus field to given value.

### HasEnhancedStatus

`func (o *GetBounceRecords200ResponseRecordsInner) HasEnhancedStatus() bool`

HasEnhancedStatus returns a boolean if a field has been set.

### GetDiagnostic

`func (o *GetBounceRecords200ResponseRecordsInner) GetDiagnostic() string`

GetDiagnostic returns the Diagnostic field if non-nil, zero value otherwise.

### GetDiagnosticOk

`func (o *GetBounceRecords200ResponseRecordsInner) GetDiagnosticOk() (*string, bool)`

GetDiagnosticOk returns a tuple with the Diagnostic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiagnostic

`func (o *GetBounceRecords200ResponseRecordsInner) SetDiagnostic(v string)`

SetDiagnostic sets Diagnostic field to given value.

### HasDiagnostic

`func (o *GetBounceRecords200ResponseRecordsInner) HasDiagnostic() bool`

HasDiagnostic returns a boolean if a field has been set.

### GetMxHost

`func (o *GetBounceRecords200ResponseRecordsInner) GetMxHost() string`

GetMxHost returns the MxHost field if non-nil, zero value otherwise.

### GetMxHostOk

`func (o *GetBounceRecords200ResponseRecordsInner) GetMxHostOk() (*string, bool)`

GetMxHostOk returns a tuple with the MxHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMxHost

`func (o *GetBounceRecords200ResponseRecordsInner) SetMxHost(v string)`

SetMxHost sets MxHost field to given value.

### HasMxHost

`func (o *GetBounceRecords200ResponseRecordsInner) HasMxHost() bool`

HasMxHost returns a boolean if a field has been set.

### GetBouncedAt

`func (o *GetBounceRecords200ResponseRecordsInner) GetBouncedAt() time.Time`

GetBouncedAt returns the BouncedAt field if non-nil, zero value otherwise.

### GetBouncedAtOk

`func (o *GetBounceRecords200ResponseRecordsInner) GetBouncedAtOk() (*time.Time, bool)`

GetBouncedAtOk returns a tuple with the BouncedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBouncedAt

`func (o *GetBounceRecords200ResponseRecordsInner) SetBouncedAt(v time.Time)`

SetBouncedAt sets BouncedAt field to given value.

### HasBouncedAt

`func (o *GetBounceRecords200ResponseRecordsInner) HasBouncedAt() bool`

HasBouncedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


