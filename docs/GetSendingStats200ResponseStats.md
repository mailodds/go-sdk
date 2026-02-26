# GetSendingStats200ResponseStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Period** | Pointer to **string** |  | [optional] 
**Sent** | Pointer to **int32** |  | [optional] 
**Delivered** | Pointer to **int32** |  | [optional] 
**Bounced** | Pointer to **int32** |  | [optional] 
**Deferred** | Pointer to **int32** |  | [optional] 
**Failed** | Pointer to **int32** |  | [optional] 
**OpenedTotal** | Pointer to **int32** |  | [optional] 
**OpenedUnique** | Pointer to **int32** |  | [optional] 
**ClickedTotal** | Pointer to **int32** |  | [optional] 
**ClickedUnique** | Pointer to **int32** |  | [optional] 
**Unsubscribed** | Pointer to **int32** |  | [optional] 
**DeliveryRate** | Pointer to **float32** |  | [optional] 
**OpenRate** | Pointer to **float32** |  | [optional] 
**ClickRate** | Pointer to **float32** |  | [optional] 
**BotOpens** | Pointer to **int32** | Opens from known bots/scanners | [optional] 
**HumanOpens** | Pointer to **int32** | Verified human opens | [optional] 
**BotOpenPct** | Pointer to **float32** | Percentage of opens from bots | [optional] 

## Methods

### NewGetSendingStats200ResponseStats

`func NewGetSendingStats200ResponseStats() *GetSendingStats200ResponseStats`

NewGetSendingStats200ResponseStats instantiates a new GetSendingStats200ResponseStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSendingStats200ResponseStatsWithDefaults

`func NewGetSendingStats200ResponseStatsWithDefaults() *GetSendingStats200ResponseStats`

NewGetSendingStats200ResponseStatsWithDefaults instantiates a new GetSendingStats200ResponseStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPeriod

`func (o *GetSendingStats200ResponseStats) GetPeriod() string`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *GetSendingStats200ResponseStats) GetPeriodOk() (*string, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *GetSendingStats200ResponseStats) SetPeriod(v string)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *GetSendingStats200ResponseStats) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### GetSent

`func (o *GetSendingStats200ResponseStats) GetSent() int32`

GetSent returns the Sent field if non-nil, zero value otherwise.

### GetSentOk

`func (o *GetSendingStats200ResponseStats) GetSentOk() (*int32, bool)`

GetSentOk returns a tuple with the Sent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSent

`func (o *GetSendingStats200ResponseStats) SetSent(v int32)`

SetSent sets Sent field to given value.

### HasSent

`func (o *GetSendingStats200ResponseStats) HasSent() bool`

HasSent returns a boolean if a field has been set.

### GetDelivered

`func (o *GetSendingStats200ResponseStats) GetDelivered() int32`

GetDelivered returns the Delivered field if non-nil, zero value otherwise.

### GetDeliveredOk

`func (o *GetSendingStats200ResponseStats) GetDeliveredOk() (*int32, bool)`

GetDeliveredOk returns a tuple with the Delivered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelivered

`func (o *GetSendingStats200ResponseStats) SetDelivered(v int32)`

SetDelivered sets Delivered field to given value.

### HasDelivered

`func (o *GetSendingStats200ResponseStats) HasDelivered() bool`

HasDelivered returns a boolean if a field has been set.

### GetBounced

`func (o *GetSendingStats200ResponseStats) GetBounced() int32`

GetBounced returns the Bounced field if non-nil, zero value otherwise.

### GetBouncedOk

`func (o *GetSendingStats200ResponseStats) GetBouncedOk() (*int32, bool)`

GetBouncedOk returns a tuple with the Bounced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBounced

`func (o *GetSendingStats200ResponseStats) SetBounced(v int32)`

SetBounced sets Bounced field to given value.

### HasBounced

`func (o *GetSendingStats200ResponseStats) HasBounced() bool`

HasBounced returns a boolean if a field has been set.

### GetDeferred

`func (o *GetSendingStats200ResponseStats) GetDeferred() int32`

GetDeferred returns the Deferred field if non-nil, zero value otherwise.

### GetDeferredOk

`func (o *GetSendingStats200ResponseStats) GetDeferredOk() (*int32, bool)`

GetDeferredOk returns a tuple with the Deferred field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeferred

`func (o *GetSendingStats200ResponseStats) SetDeferred(v int32)`

SetDeferred sets Deferred field to given value.

### HasDeferred

`func (o *GetSendingStats200ResponseStats) HasDeferred() bool`

HasDeferred returns a boolean if a field has been set.

### GetFailed

`func (o *GetSendingStats200ResponseStats) GetFailed() int32`

GetFailed returns the Failed field if non-nil, zero value otherwise.

### GetFailedOk

`func (o *GetSendingStats200ResponseStats) GetFailedOk() (*int32, bool)`

GetFailedOk returns a tuple with the Failed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailed

`func (o *GetSendingStats200ResponseStats) SetFailed(v int32)`

SetFailed sets Failed field to given value.

### HasFailed

`func (o *GetSendingStats200ResponseStats) HasFailed() bool`

HasFailed returns a boolean if a field has been set.

### GetOpenedTotal

`func (o *GetSendingStats200ResponseStats) GetOpenedTotal() int32`

GetOpenedTotal returns the OpenedTotal field if non-nil, zero value otherwise.

### GetOpenedTotalOk

`func (o *GetSendingStats200ResponseStats) GetOpenedTotalOk() (*int32, bool)`

GetOpenedTotalOk returns a tuple with the OpenedTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenedTotal

`func (o *GetSendingStats200ResponseStats) SetOpenedTotal(v int32)`

SetOpenedTotal sets OpenedTotal field to given value.

### HasOpenedTotal

`func (o *GetSendingStats200ResponseStats) HasOpenedTotal() bool`

HasOpenedTotal returns a boolean if a field has been set.

### GetOpenedUnique

`func (o *GetSendingStats200ResponseStats) GetOpenedUnique() int32`

GetOpenedUnique returns the OpenedUnique field if non-nil, zero value otherwise.

### GetOpenedUniqueOk

`func (o *GetSendingStats200ResponseStats) GetOpenedUniqueOk() (*int32, bool)`

GetOpenedUniqueOk returns a tuple with the OpenedUnique field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenedUnique

`func (o *GetSendingStats200ResponseStats) SetOpenedUnique(v int32)`

SetOpenedUnique sets OpenedUnique field to given value.

### HasOpenedUnique

`func (o *GetSendingStats200ResponseStats) HasOpenedUnique() bool`

HasOpenedUnique returns a boolean if a field has been set.

### GetClickedTotal

`func (o *GetSendingStats200ResponseStats) GetClickedTotal() int32`

GetClickedTotal returns the ClickedTotal field if non-nil, zero value otherwise.

### GetClickedTotalOk

`func (o *GetSendingStats200ResponseStats) GetClickedTotalOk() (*int32, bool)`

GetClickedTotalOk returns a tuple with the ClickedTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClickedTotal

`func (o *GetSendingStats200ResponseStats) SetClickedTotal(v int32)`

SetClickedTotal sets ClickedTotal field to given value.

### HasClickedTotal

`func (o *GetSendingStats200ResponseStats) HasClickedTotal() bool`

HasClickedTotal returns a boolean if a field has been set.

### GetClickedUnique

`func (o *GetSendingStats200ResponseStats) GetClickedUnique() int32`

GetClickedUnique returns the ClickedUnique field if non-nil, zero value otherwise.

### GetClickedUniqueOk

`func (o *GetSendingStats200ResponseStats) GetClickedUniqueOk() (*int32, bool)`

GetClickedUniqueOk returns a tuple with the ClickedUnique field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClickedUnique

`func (o *GetSendingStats200ResponseStats) SetClickedUnique(v int32)`

SetClickedUnique sets ClickedUnique field to given value.

### HasClickedUnique

`func (o *GetSendingStats200ResponseStats) HasClickedUnique() bool`

HasClickedUnique returns a boolean if a field has been set.

### GetUnsubscribed

`func (o *GetSendingStats200ResponseStats) GetUnsubscribed() int32`

GetUnsubscribed returns the Unsubscribed field if non-nil, zero value otherwise.

### GetUnsubscribedOk

`func (o *GetSendingStats200ResponseStats) GetUnsubscribedOk() (*int32, bool)`

GetUnsubscribedOk returns a tuple with the Unsubscribed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnsubscribed

`func (o *GetSendingStats200ResponseStats) SetUnsubscribed(v int32)`

SetUnsubscribed sets Unsubscribed field to given value.

### HasUnsubscribed

`func (o *GetSendingStats200ResponseStats) HasUnsubscribed() bool`

HasUnsubscribed returns a boolean if a field has been set.

### GetDeliveryRate

`func (o *GetSendingStats200ResponseStats) GetDeliveryRate() float32`

GetDeliveryRate returns the DeliveryRate field if non-nil, zero value otherwise.

### GetDeliveryRateOk

`func (o *GetSendingStats200ResponseStats) GetDeliveryRateOk() (*float32, bool)`

GetDeliveryRateOk returns a tuple with the DeliveryRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryRate

`func (o *GetSendingStats200ResponseStats) SetDeliveryRate(v float32)`

SetDeliveryRate sets DeliveryRate field to given value.

### HasDeliveryRate

`func (o *GetSendingStats200ResponseStats) HasDeliveryRate() bool`

HasDeliveryRate returns a boolean if a field has been set.

### GetOpenRate

`func (o *GetSendingStats200ResponseStats) GetOpenRate() float32`

GetOpenRate returns the OpenRate field if non-nil, zero value otherwise.

### GetOpenRateOk

`func (o *GetSendingStats200ResponseStats) GetOpenRateOk() (*float32, bool)`

GetOpenRateOk returns a tuple with the OpenRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenRate

`func (o *GetSendingStats200ResponseStats) SetOpenRate(v float32)`

SetOpenRate sets OpenRate field to given value.

### HasOpenRate

`func (o *GetSendingStats200ResponseStats) HasOpenRate() bool`

HasOpenRate returns a boolean if a field has been set.

### GetClickRate

`func (o *GetSendingStats200ResponseStats) GetClickRate() float32`

GetClickRate returns the ClickRate field if non-nil, zero value otherwise.

### GetClickRateOk

`func (o *GetSendingStats200ResponseStats) GetClickRateOk() (*float32, bool)`

GetClickRateOk returns a tuple with the ClickRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClickRate

`func (o *GetSendingStats200ResponseStats) SetClickRate(v float32)`

SetClickRate sets ClickRate field to given value.

### HasClickRate

`func (o *GetSendingStats200ResponseStats) HasClickRate() bool`

HasClickRate returns a boolean if a field has been set.

### GetBotOpens

`func (o *GetSendingStats200ResponseStats) GetBotOpens() int32`

GetBotOpens returns the BotOpens field if non-nil, zero value otherwise.

### GetBotOpensOk

`func (o *GetSendingStats200ResponseStats) GetBotOpensOk() (*int32, bool)`

GetBotOpensOk returns a tuple with the BotOpens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBotOpens

`func (o *GetSendingStats200ResponseStats) SetBotOpens(v int32)`

SetBotOpens sets BotOpens field to given value.

### HasBotOpens

`func (o *GetSendingStats200ResponseStats) HasBotOpens() bool`

HasBotOpens returns a boolean if a field has been set.

### GetHumanOpens

`func (o *GetSendingStats200ResponseStats) GetHumanOpens() int32`

GetHumanOpens returns the HumanOpens field if non-nil, zero value otherwise.

### GetHumanOpensOk

`func (o *GetSendingStats200ResponseStats) GetHumanOpensOk() (*int32, bool)`

GetHumanOpensOk returns a tuple with the HumanOpens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHumanOpens

`func (o *GetSendingStats200ResponseStats) SetHumanOpens(v int32)`

SetHumanOpens sets HumanOpens field to given value.

### HasHumanOpens

`func (o *GetSendingStats200ResponseStats) HasHumanOpens() bool`

HasHumanOpens returns a boolean if a field has been set.

### GetBotOpenPct

`func (o *GetSendingStats200ResponseStats) GetBotOpenPct() float32`

GetBotOpenPct returns the BotOpenPct field if non-nil, zero value otherwise.

### GetBotOpenPctOk

`func (o *GetSendingStats200ResponseStats) GetBotOpenPctOk() (*float32, bool)`

GetBotOpenPctOk returns a tuple with the BotOpenPct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBotOpenPct

`func (o *GetSendingStats200ResponseStats) SetBotOpenPct(v float32)`

SetBotOpenPct sets BotOpenPct field to given value.

### HasBotOpenPct

`func (o *GetSendingStats200ResponseStats) HasBotOpenPct() bool`

HasBotOpenPct returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


