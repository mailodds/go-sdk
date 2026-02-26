# SubscribeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** | Subscriber email address | 
**Name** | Pointer to **string** | Subscriber name | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Custom metadata key-value pairs | [optional] 
**PageUrl** | Pointer to **string** | URL of the page where the subscription form was submitted (for consent proof) | [optional] 
**FormId** | Pointer to **string** | Identifier of the form used to subscribe (for consent proof) | [optional] 

## Methods

### NewSubscribeRequest

`func NewSubscribeRequest(email string, ) *SubscribeRequest`

NewSubscribeRequest instantiates a new SubscribeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscribeRequestWithDefaults

`func NewSubscribeRequestWithDefaults() *SubscribeRequest`

NewSubscribeRequestWithDefaults instantiates a new SubscribeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *SubscribeRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *SubscribeRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *SubscribeRequest) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetName

`func (o *SubscribeRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SubscribeRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SubscribeRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SubscribeRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetMetadata

`func (o *SubscribeRequest) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SubscribeRequest) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SubscribeRequest) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *SubscribeRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetPageUrl

`func (o *SubscribeRequest) GetPageUrl() string`

GetPageUrl returns the PageUrl field if non-nil, zero value otherwise.

### GetPageUrlOk

`func (o *SubscribeRequest) GetPageUrlOk() (*string, bool)`

GetPageUrlOk returns a tuple with the PageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageUrl

`func (o *SubscribeRequest) SetPageUrl(v string)`

SetPageUrl sets PageUrl field to given value.

### HasPageUrl

`func (o *SubscribeRequest) HasPageUrl() bool`

HasPageUrl returns a boolean if a field has been set.

### GetFormId

`func (o *SubscribeRequest) GetFormId() string`

GetFormId returns the FormId field if non-nil, zero value otherwise.

### GetFormIdOk

`func (o *SubscribeRequest) GetFormIdOk() (*string, bool)`

GetFormIdOk returns a tuple with the FormId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormId

`func (o *SubscribeRequest) SetFormId(v string)`

SetFormId sets FormId field to given value.

### HasFormId

`func (o *SubscribeRequest) HasFormId() bool`

HasFormId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


