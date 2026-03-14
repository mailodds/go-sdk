# SpamCheck

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Spam check UUID | [optional] 
**FromDomain** | Pointer to **string** |  | [optional] 
**Score** | Pointer to **float32** | Overall spam score (0-10, lower is better) | [optional] 
**Verdict** | Pointer to **string** | Overall verdict | [optional] 
**Checks** | Pointer to [**SpamCheckChecks**](SpamCheckChecks.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewSpamCheck

`func NewSpamCheck() *SpamCheck`

NewSpamCheck instantiates a new SpamCheck object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpamCheckWithDefaults

`func NewSpamCheckWithDefaults() *SpamCheck`

NewSpamCheckWithDefaults instantiates a new SpamCheck object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SpamCheck) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SpamCheck) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SpamCheck) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SpamCheck) HasId() bool`

HasId returns a boolean if a field has been set.

### GetFromDomain

`func (o *SpamCheck) GetFromDomain() string`

GetFromDomain returns the FromDomain field if non-nil, zero value otherwise.

### GetFromDomainOk

`func (o *SpamCheck) GetFromDomainOk() (*string, bool)`

GetFromDomainOk returns a tuple with the FromDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromDomain

`func (o *SpamCheck) SetFromDomain(v string)`

SetFromDomain sets FromDomain field to given value.

### HasFromDomain

`func (o *SpamCheck) HasFromDomain() bool`

HasFromDomain returns a boolean if a field has been set.

### GetScore

`func (o *SpamCheck) GetScore() float32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *SpamCheck) GetScoreOk() (*float32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *SpamCheck) SetScore(v float32)`

SetScore sets Score field to given value.

### HasScore

`func (o *SpamCheck) HasScore() bool`

HasScore returns a boolean if a field has been set.

### GetVerdict

`func (o *SpamCheck) GetVerdict() string`

GetVerdict returns the Verdict field if non-nil, zero value otherwise.

### GetVerdictOk

`func (o *SpamCheck) GetVerdictOk() (*string, bool)`

GetVerdictOk returns a tuple with the Verdict field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerdict

`func (o *SpamCheck) SetVerdict(v string)`

SetVerdict sets Verdict field to given value.

### HasVerdict

`func (o *SpamCheck) HasVerdict() bool`

HasVerdict returns a boolean if a field has been set.

### GetChecks

`func (o *SpamCheck) GetChecks() SpamCheckChecks`

GetChecks returns the Checks field if non-nil, zero value otherwise.

### GetChecksOk

`func (o *SpamCheck) GetChecksOk() (*SpamCheckChecks, bool)`

GetChecksOk returns a tuple with the Checks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecks

`func (o *SpamCheck) SetChecks(v SpamCheckChecks)`

SetChecks sets Checks field to given value.

### HasChecks

`func (o *SpamCheck) HasChecks() bool`

HasChecks returns a boolean if a field has been set.

### GetCreatedAt

`func (o *SpamCheck) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SpamCheck) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SpamCheck) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *SpamCheck) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


