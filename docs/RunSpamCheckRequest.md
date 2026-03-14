# RunSpamCheckRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FromDomain** | **string** | Sending domain to check | 
**Links** | Pointer to **[]string** | URLs included in the email | [optional] 
**SubjectPreview** | Pointer to **string** | Email subject line to analyze | [optional] 
**ClientScores** | Pointer to **map[string]interface{}** | Client-side spam scores to include in analysis | [optional] 

## Methods

### NewRunSpamCheckRequest

`func NewRunSpamCheckRequest(fromDomain string, ) *RunSpamCheckRequest`

NewRunSpamCheckRequest instantiates a new RunSpamCheckRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRunSpamCheckRequestWithDefaults

`func NewRunSpamCheckRequestWithDefaults() *RunSpamCheckRequest`

NewRunSpamCheckRequestWithDefaults instantiates a new RunSpamCheckRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFromDomain

`func (o *RunSpamCheckRequest) GetFromDomain() string`

GetFromDomain returns the FromDomain field if non-nil, zero value otherwise.

### GetFromDomainOk

`func (o *RunSpamCheckRequest) GetFromDomainOk() (*string, bool)`

GetFromDomainOk returns a tuple with the FromDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromDomain

`func (o *RunSpamCheckRequest) SetFromDomain(v string)`

SetFromDomain sets FromDomain field to given value.


### GetLinks

`func (o *RunSpamCheckRequest) GetLinks() []string`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *RunSpamCheckRequest) GetLinksOk() (*[]string, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *RunSpamCheckRequest) SetLinks(v []string)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *RunSpamCheckRequest) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetSubjectPreview

`func (o *RunSpamCheckRequest) GetSubjectPreview() string`

GetSubjectPreview returns the SubjectPreview field if non-nil, zero value otherwise.

### GetSubjectPreviewOk

`func (o *RunSpamCheckRequest) GetSubjectPreviewOk() (*string, bool)`

GetSubjectPreviewOk returns a tuple with the SubjectPreview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectPreview

`func (o *RunSpamCheckRequest) SetSubjectPreview(v string)`

SetSubjectPreview sets SubjectPreview field to given value.

### HasSubjectPreview

`func (o *RunSpamCheckRequest) HasSubjectPreview() bool`

HasSubjectPreview returns a boolean if a field has been set.

### GetClientScores

`func (o *RunSpamCheckRequest) GetClientScores() map[string]interface{}`

GetClientScores returns the ClientScores field if non-nil, zero value otherwise.

### GetClientScoresOk

`func (o *RunSpamCheckRequest) GetClientScoresOk() (*map[string]interface{}, bool)`

GetClientScoresOk returns a tuple with the ClientScores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientScores

`func (o *RunSpamCheckRequest) SetClientScores(v map[string]interface{})`

SetClientScores sets ClientScores field to given value.

### HasClientScores

`func (o *RunSpamCheckRequest) HasClientScores() bool`

HasClientScores returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


