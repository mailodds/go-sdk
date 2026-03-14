# ClassifyContent200ResponseContentCheck

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Score** | Pointer to **float32** | Overall content quality score (0-100) | [optional] 
**Verdict** | Pointer to **string** | Overall verdict | [optional] 
**Categories** | Pointer to [**[]ClassifyContent200ResponseContentCheckCategoriesInner**](ClassifyContent200ResponseContentCheckCategoriesInner.md) |  | [optional] 
**Suggestions** | Pointer to **[]string** | Improvement suggestions | [optional] 

## Methods

### NewClassifyContent200ResponseContentCheck

`func NewClassifyContent200ResponseContentCheck() *ClassifyContent200ResponseContentCheck`

NewClassifyContent200ResponseContentCheck instantiates a new ClassifyContent200ResponseContentCheck object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClassifyContent200ResponseContentCheckWithDefaults

`func NewClassifyContent200ResponseContentCheckWithDefaults() *ClassifyContent200ResponseContentCheck`

NewClassifyContent200ResponseContentCheckWithDefaults instantiates a new ClassifyContent200ResponseContentCheck object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScore

`func (o *ClassifyContent200ResponseContentCheck) GetScore() float32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *ClassifyContent200ResponseContentCheck) GetScoreOk() (*float32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *ClassifyContent200ResponseContentCheck) SetScore(v float32)`

SetScore sets Score field to given value.

### HasScore

`func (o *ClassifyContent200ResponseContentCheck) HasScore() bool`

HasScore returns a boolean if a field has been set.

### GetVerdict

`func (o *ClassifyContent200ResponseContentCheck) GetVerdict() string`

GetVerdict returns the Verdict field if non-nil, zero value otherwise.

### GetVerdictOk

`func (o *ClassifyContent200ResponseContentCheck) GetVerdictOk() (*string, bool)`

GetVerdictOk returns a tuple with the Verdict field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerdict

`func (o *ClassifyContent200ResponseContentCheck) SetVerdict(v string)`

SetVerdict sets Verdict field to given value.

### HasVerdict

`func (o *ClassifyContent200ResponseContentCheck) HasVerdict() bool`

HasVerdict returns a boolean if a field has been set.

### GetCategories

`func (o *ClassifyContent200ResponseContentCheck) GetCategories() []ClassifyContent200ResponseContentCheckCategoriesInner`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *ClassifyContent200ResponseContentCheck) GetCategoriesOk() (*[]ClassifyContent200ResponseContentCheckCategoriesInner, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *ClassifyContent200ResponseContentCheck) SetCategories(v []ClassifyContent200ResponseContentCheckCategoriesInner)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *ClassifyContent200ResponseContentCheck) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### GetSuggestions

`func (o *ClassifyContent200ResponseContentCheck) GetSuggestions() []string`

GetSuggestions returns the Suggestions field if non-nil, zero value otherwise.

### GetSuggestionsOk

`func (o *ClassifyContent200ResponseContentCheck) GetSuggestionsOk() (*[]string, bool)`

GetSuggestionsOk returns a tuple with the Suggestions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuggestions

`func (o *ClassifyContent200ResponseContentCheck) SetSuggestions(v []string)`

SetSuggestions sets Suggestions field to given value.

### HasSuggestions

`func (o *ClassifyContent200ResponseContentCheck) HasSuggestions() bool`

HasSuggestions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


