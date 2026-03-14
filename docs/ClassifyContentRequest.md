# ClassifyContentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subject** | Pointer to **string** | Email subject line | [optional] 
**HtmlBody** | Pointer to **string** | HTML email body | [optional] 
**Content** | Pointer to **string** | Raw text content (alternative to subject+html_body) | [optional] 

## Methods

### NewClassifyContentRequest

`func NewClassifyContentRequest() *ClassifyContentRequest`

NewClassifyContentRequest instantiates a new ClassifyContentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClassifyContentRequestWithDefaults

`func NewClassifyContentRequestWithDefaults() *ClassifyContentRequest`

NewClassifyContentRequestWithDefaults instantiates a new ClassifyContentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubject

`func (o *ClassifyContentRequest) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *ClassifyContentRequest) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *ClassifyContentRequest) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *ClassifyContentRequest) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetHtmlBody

`func (o *ClassifyContentRequest) GetHtmlBody() string`

GetHtmlBody returns the HtmlBody field if non-nil, zero value otherwise.

### GetHtmlBodyOk

`func (o *ClassifyContentRequest) GetHtmlBodyOk() (*string, bool)`

GetHtmlBodyOk returns a tuple with the HtmlBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlBody

`func (o *ClassifyContentRequest) SetHtmlBody(v string)`

SetHtmlBody sets HtmlBody field to given value.

### HasHtmlBody

`func (o *ClassifyContentRequest) HasHtmlBody() bool`

HasHtmlBody returns a boolean if a field has been set.

### GetContent

`func (o *ClassifyContentRequest) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *ClassifyContentRequest) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *ClassifyContentRequest) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *ClassifyContentRequest) HasContent() bool`

HasContent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


