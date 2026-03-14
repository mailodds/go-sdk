# CampaignStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sent** | Pointer to **int32** |  | [optional] 
**Delivered** | Pointer to **int32** |  | [optional] 
**Opened** | Pointer to **int32** |  | [optional] 
**Clicked** | Pointer to **int32** |  | [optional] 
**Bounced** | Pointer to **int32** |  | [optional] 
**Unsubscribed** | Pointer to **int32** |  | [optional] 
**Complained** | Pointer to **int32** |  | [optional] 
**DeliveryRate** | Pointer to **float32** |  | [optional] 
**OpenRate** | Pointer to **float32** |  | [optional] 
**ClickRate** | Pointer to **float32** |  | [optional] 

## Methods

### NewCampaignStats

`func NewCampaignStats() *CampaignStats`

NewCampaignStats instantiates a new CampaignStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignStatsWithDefaults

`func NewCampaignStatsWithDefaults() *CampaignStats`

NewCampaignStatsWithDefaults instantiates a new CampaignStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSent

`func (o *CampaignStats) GetSent() int32`

GetSent returns the Sent field if non-nil, zero value otherwise.

### GetSentOk

`func (o *CampaignStats) GetSentOk() (*int32, bool)`

GetSentOk returns a tuple with the Sent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSent

`func (o *CampaignStats) SetSent(v int32)`

SetSent sets Sent field to given value.

### HasSent

`func (o *CampaignStats) HasSent() bool`

HasSent returns a boolean if a field has been set.

### GetDelivered

`func (o *CampaignStats) GetDelivered() int32`

GetDelivered returns the Delivered field if non-nil, zero value otherwise.

### GetDeliveredOk

`func (o *CampaignStats) GetDeliveredOk() (*int32, bool)`

GetDeliveredOk returns a tuple with the Delivered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelivered

`func (o *CampaignStats) SetDelivered(v int32)`

SetDelivered sets Delivered field to given value.

### HasDelivered

`func (o *CampaignStats) HasDelivered() bool`

HasDelivered returns a boolean if a field has been set.

### GetOpened

`func (o *CampaignStats) GetOpened() int32`

GetOpened returns the Opened field if non-nil, zero value otherwise.

### GetOpenedOk

`func (o *CampaignStats) GetOpenedOk() (*int32, bool)`

GetOpenedOk returns a tuple with the Opened field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpened

`func (o *CampaignStats) SetOpened(v int32)`

SetOpened sets Opened field to given value.

### HasOpened

`func (o *CampaignStats) HasOpened() bool`

HasOpened returns a boolean if a field has been set.

### GetClicked

`func (o *CampaignStats) GetClicked() int32`

GetClicked returns the Clicked field if non-nil, zero value otherwise.

### GetClickedOk

`func (o *CampaignStats) GetClickedOk() (*int32, bool)`

GetClickedOk returns a tuple with the Clicked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClicked

`func (o *CampaignStats) SetClicked(v int32)`

SetClicked sets Clicked field to given value.

### HasClicked

`func (o *CampaignStats) HasClicked() bool`

HasClicked returns a boolean if a field has been set.

### GetBounced

`func (o *CampaignStats) GetBounced() int32`

GetBounced returns the Bounced field if non-nil, zero value otherwise.

### GetBouncedOk

`func (o *CampaignStats) GetBouncedOk() (*int32, bool)`

GetBouncedOk returns a tuple with the Bounced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBounced

`func (o *CampaignStats) SetBounced(v int32)`

SetBounced sets Bounced field to given value.

### HasBounced

`func (o *CampaignStats) HasBounced() bool`

HasBounced returns a boolean if a field has been set.

### GetUnsubscribed

`func (o *CampaignStats) GetUnsubscribed() int32`

GetUnsubscribed returns the Unsubscribed field if non-nil, zero value otherwise.

### GetUnsubscribedOk

`func (o *CampaignStats) GetUnsubscribedOk() (*int32, bool)`

GetUnsubscribedOk returns a tuple with the Unsubscribed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnsubscribed

`func (o *CampaignStats) SetUnsubscribed(v int32)`

SetUnsubscribed sets Unsubscribed field to given value.

### HasUnsubscribed

`func (o *CampaignStats) HasUnsubscribed() bool`

HasUnsubscribed returns a boolean if a field has been set.

### GetComplained

`func (o *CampaignStats) GetComplained() int32`

GetComplained returns the Complained field if non-nil, zero value otherwise.

### GetComplainedOk

`func (o *CampaignStats) GetComplainedOk() (*int32, bool)`

GetComplainedOk returns a tuple with the Complained field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplained

`func (o *CampaignStats) SetComplained(v int32)`

SetComplained sets Complained field to given value.

### HasComplained

`func (o *CampaignStats) HasComplained() bool`

HasComplained returns a boolean if a field has been set.

### GetDeliveryRate

`func (o *CampaignStats) GetDeliveryRate() float32`

GetDeliveryRate returns the DeliveryRate field if non-nil, zero value otherwise.

### GetDeliveryRateOk

`func (o *CampaignStats) GetDeliveryRateOk() (*float32, bool)`

GetDeliveryRateOk returns a tuple with the DeliveryRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryRate

`func (o *CampaignStats) SetDeliveryRate(v float32)`

SetDeliveryRate sets DeliveryRate field to given value.

### HasDeliveryRate

`func (o *CampaignStats) HasDeliveryRate() bool`

HasDeliveryRate returns a boolean if a field has been set.

### GetOpenRate

`func (o *CampaignStats) GetOpenRate() float32`

GetOpenRate returns the OpenRate field if non-nil, zero value otherwise.

### GetOpenRateOk

`func (o *CampaignStats) GetOpenRateOk() (*float32, bool)`

GetOpenRateOk returns a tuple with the OpenRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenRate

`func (o *CampaignStats) SetOpenRate(v float32)`

SetOpenRate sets OpenRate field to given value.

### HasOpenRate

`func (o *CampaignStats) HasOpenRate() bool`

HasOpenRate returns a boolean if a field has been set.

### GetClickRate

`func (o *CampaignStats) GetClickRate() float32`

GetClickRate returns the ClickRate field if non-nil, zero value otherwise.

### GetClickRateOk

`func (o *CampaignStats) GetClickRateOk() (*float32, bool)`

GetClickRateOk returns a tuple with the ClickRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClickRate

`func (o *CampaignStats) SetClickRate(v float32)`

SetClickRate sets ClickRate field to given value.

### HasClickRate

`func (o *CampaignStats) HasClickRate() bool`

HasClickRate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


