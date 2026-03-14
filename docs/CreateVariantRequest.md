# CreateVariantRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Variant name (e.g., \&quot;Variant A\&quot;) | 
**Subject** | **string** | Email subject line | 
**Html** | Pointer to **string** | HTML email body | [optional] 
**Text** | Pointer to **string** | Plain text email body | [optional] 
**Weight** | **int32** | Traffic weight percentage (all variant weights must sum to 100) | 

## Methods

### NewCreateVariantRequest

`func NewCreateVariantRequest(name string, subject string, weight int32, ) *CreateVariantRequest`

NewCreateVariantRequest instantiates a new CreateVariantRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateVariantRequestWithDefaults

`func NewCreateVariantRequestWithDefaults() *CreateVariantRequest`

NewCreateVariantRequestWithDefaults instantiates a new CreateVariantRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateVariantRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateVariantRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateVariantRequest) SetName(v string)`

SetName sets Name field to given value.


### GetSubject

`func (o *CreateVariantRequest) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *CreateVariantRequest) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *CreateVariantRequest) SetSubject(v string)`

SetSubject sets Subject field to given value.


### GetHtml

`func (o *CreateVariantRequest) GetHtml() string`

GetHtml returns the Html field if non-nil, zero value otherwise.

### GetHtmlOk

`func (o *CreateVariantRequest) GetHtmlOk() (*string, bool)`

GetHtmlOk returns a tuple with the Html field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtml

`func (o *CreateVariantRequest) SetHtml(v string)`

SetHtml sets Html field to given value.

### HasHtml

`func (o *CreateVariantRequest) HasHtml() bool`

HasHtml returns a boolean if a field has been set.

### GetText

`func (o *CreateVariantRequest) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *CreateVariantRequest) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *CreateVariantRequest) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *CreateVariantRequest) HasText() bool`

HasText returns a boolean if a field has been set.

### GetWeight

`func (o *CreateVariantRequest) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *CreateVariantRequest) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *CreateVariantRequest) SetWeight(v int32)`

SetWeight sets Weight field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


