# UpdatePixelSettingsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PixelSubscribeListId** | **NullableInt32** | Contact list ID for pixel subscriptions, or null to disable | 

## Methods

### NewUpdatePixelSettingsRequest

`func NewUpdatePixelSettingsRequest(pixelSubscribeListId NullableInt32, ) *UpdatePixelSettingsRequest`

NewUpdatePixelSettingsRequest instantiates a new UpdatePixelSettingsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdatePixelSettingsRequestWithDefaults

`func NewUpdatePixelSettingsRequestWithDefaults() *UpdatePixelSettingsRequest`

NewUpdatePixelSettingsRequestWithDefaults instantiates a new UpdatePixelSettingsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPixelSubscribeListId

`func (o *UpdatePixelSettingsRequest) GetPixelSubscribeListId() int32`

GetPixelSubscribeListId returns the PixelSubscribeListId field if non-nil, zero value otherwise.

### GetPixelSubscribeListIdOk

`func (o *UpdatePixelSettingsRequest) GetPixelSubscribeListIdOk() (*int32, bool)`

GetPixelSubscribeListIdOk returns a tuple with the PixelSubscribeListId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPixelSubscribeListId

`func (o *UpdatePixelSettingsRequest) SetPixelSubscribeListId(v int32)`

SetPixelSubscribeListId sets PixelSubscribeListId field to given value.


### SetPixelSubscribeListIdNil

`func (o *UpdatePixelSettingsRequest) SetPixelSubscribeListIdNil(b bool)`

 SetPixelSubscribeListIdNil sets the value for PixelSubscribeListId to be an explicit nil

### UnsetPixelSubscribeListId
`func (o *UpdatePixelSettingsRequest) UnsetPixelSubscribeListId()`

UnsetPixelSubscribeListId ensures that no value is present for PixelSubscribeListId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


