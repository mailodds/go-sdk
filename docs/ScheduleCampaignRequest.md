# ScheduleCampaignRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SendAt** | **time.Time** | Scheduled send time (ISO 8601, must be in the future) | 

## Methods

### NewScheduleCampaignRequest

`func NewScheduleCampaignRequest(sendAt time.Time, ) *ScheduleCampaignRequest`

NewScheduleCampaignRequest instantiates a new ScheduleCampaignRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduleCampaignRequestWithDefaults

`func NewScheduleCampaignRequestWithDefaults() *ScheduleCampaignRequest`

NewScheduleCampaignRequestWithDefaults instantiates a new ScheduleCampaignRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSendAt

`func (o *ScheduleCampaignRequest) GetSendAt() time.Time`

GetSendAt returns the SendAt field if non-nil, zero value otherwise.

### GetSendAtOk

`func (o *ScheduleCampaignRequest) GetSendAtOk() (*time.Time, bool)`

GetSendAtOk returns a tuple with the SendAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendAt

`func (o *ScheduleCampaignRequest) SetSendAt(v time.Time)`

SetSendAt sets SendAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


