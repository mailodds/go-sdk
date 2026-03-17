# CreateBounceAnalysisRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Text** | **string** | Bounce log text to analyze. Identifies patterns, categorizes bounce types, and provides remediation recommendations. | 
**Name** | Pointer to **string** | Optional name for this bounce analysis | [optional] 

## Methods

### NewCreateBounceAnalysisRequest

`func NewCreateBounceAnalysisRequest(text string, ) *CreateBounceAnalysisRequest`

NewCreateBounceAnalysisRequest instantiates a new CreateBounceAnalysisRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateBounceAnalysisRequestWithDefaults

`func NewCreateBounceAnalysisRequestWithDefaults() *CreateBounceAnalysisRequest`

NewCreateBounceAnalysisRequestWithDefaults instantiates a new CreateBounceAnalysisRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetText

`func (o *CreateBounceAnalysisRequest) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *CreateBounceAnalysisRequest) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *CreateBounceAnalysisRequest) SetText(v string)`

SetText sets Text field to given value.


### GetName

`func (o *CreateBounceAnalysisRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateBounceAnalysisRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateBounceAnalysisRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateBounceAnalysisRequest) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


