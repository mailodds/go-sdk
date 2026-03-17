# ClassifyContent200ResponseContentCheck

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** | Overall content status | [optional] 
**Flag** | Pointer to **bool** | Whether the content is flagged | [optional] 
**Reason** | Pointer to **string** | Human-readable reason for the status | [optional] 
**Priority** | Pointer to **int32** | Priority level (1&#x3D;lowest, 5&#x3D;highest) | [optional] 
**Suggestions** | Pointer to **[]string** | Improvement suggestions | [optional] 
**DurationMs** | Pointer to **int32** | Classification duration in milliseconds | [optional] 

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

### GetStatus

`func (o *ClassifyContent200ResponseContentCheck) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ClassifyContent200ResponseContentCheck) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ClassifyContent200ResponseContentCheck) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ClassifyContent200ResponseContentCheck) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetFlag

`func (o *ClassifyContent200ResponseContentCheck) GetFlag() bool`

GetFlag returns the Flag field if non-nil, zero value otherwise.

### GetFlagOk

`func (o *ClassifyContent200ResponseContentCheck) GetFlagOk() (*bool, bool)`

GetFlagOk returns a tuple with the Flag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlag

`func (o *ClassifyContent200ResponseContentCheck) SetFlag(v bool)`

SetFlag sets Flag field to given value.

### HasFlag

`func (o *ClassifyContent200ResponseContentCheck) HasFlag() bool`

HasFlag returns a boolean if a field has been set.

### GetReason

`func (o *ClassifyContent200ResponseContentCheck) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *ClassifyContent200ResponseContentCheck) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *ClassifyContent200ResponseContentCheck) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *ClassifyContent200ResponseContentCheck) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetPriority

`func (o *ClassifyContent200ResponseContentCheck) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *ClassifyContent200ResponseContentCheck) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *ClassifyContent200ResponseContentCheck) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *ClassifyContent200ResponseContentCheck) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

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

### GetDurationMs

`func (o *ClassifyContent200ResponseContentCheck) GetDurationMs() int32`

GetDurationMs returns the DurationMs field if non-nil, zero value otherwise.

### GetDurationMsOk

`func (o *ClassifyContent200ResponseContentCheck) GetDurationMsOk() (*int32, bool)`

GetDurationMsOk returns a tuple with the DurationMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationMs

`func (o *ClassifyContent200ResponseContentCheck) SetDurationMs(v int32)`

SetDurationMs sets DurationMs field to given value.

### HasDurationMs

`func (o *ClassifyContent200ResponseContentCheck) HasDurationMs() bool`

HasDurationMs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


