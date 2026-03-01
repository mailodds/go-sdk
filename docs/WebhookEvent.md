# WebhookEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Event** | **string** | Event type | 
**Timestamp** | **time.Time** | When the event occurred | 
**Job** | Pointer to [**Job**](Job.md) |  | [optional] 
**MessageId** | Pointer to **string** | Message ID (present for message.* and delivery events) | [optional] 
**AccountId** | Pointer to **int32** | Account ID (present for delivery events) | [optional] 
**DomainId** | Pointer to **string** | Sending domain UUID (present for delivery events) | [optional] 
**To** | Pointer to **string** | Recipient email (present for delivery events) | [optional] 
**Status** | Pointer to **string** | Delivery status (present for delivery events) | [optional] 
**SmtpCode** | Pointer to **int32** | SMTP response code (present for bounced/deferred/failed) | [optional] 
**SmtpResponse** | Pointer to **string** | SMTP diagnostic string (present for bounced/deferred/failed) | [optional] 
**MxHost** | Pointer to **string** | MX host that handled delivery | [optional] 
**BounceType** | Pointer to **NullableString** | Bounce classification (present for message.bounced) | [optional] 
**EnhancedStatusCode** | Pointer to **string** | Enhanced SMTP status code (e.g., 5.1.1) | [optional] 
**Attempts** | Pointer to **int32** | Number of delivery attempts | [optional] 
**Isp** | Pointer to **string** | Receiving ISP name | [optional] 
**IsMpp** | Pointer to **bool** | Whether the open was from Apple Mail Privacy Protection | [optional] 
**IpAddress** | Pointer to **string** | Client IP (present for message.opened/clicked) | [optional] 
**UserAgent** | Pointer to **string** | Client user agent (present for message.opened/clicked) | [optional] 
**IsBot** | Pointer to **bool** | Whether the event was triggered by a bot (present for message.opened/clicked) | [optional] 
**LinkUrl** | Pointer to **string** | Clicked URL (present for message.clicked) | [optional] 
**LinkIndex** | Pointer to **int32** | Position of clicked link in message (present for message.clicked) | [optional] 

## Methods

### NewWebhookEvent

`func NewWebhookEvent(event string, timestamp time.Time, ) *WebhookEvent`

NewWebhookEvent instantiates a new WebhookEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookEventWithDefaults

`func NewWebhookEventWithDefaults() *WebhookEvent`

NewWebhookEventWithDefaults instantiates a new WebhookEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvent

`func (o *WebhookEvent) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *WebhookEvent) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *WebhookEvent) SetEvent(v string)`

SetEvent sets Event field to given value.


### GetTimestamp

`func (o *WebhookEvent) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *WebhookEvent) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *WebhookEvent) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.


### GetJob

`func (o *WebhookEvent) GetJob() Job`

GetJob returns the Job field if non-nil, zero value otherwise.

### GetJobOk

`func (o *WebhookEvent) GetJobOk() (*Job, bool)`

GetJobOk returns a tuple with the Job field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJob

`func (o *WebhookEvent) SetJob(v Job)`

SetJob sets Job field to given value.

### HasJob

`func (o *WebhookEvent) HasJob() bool`

HasJob returns a boolean if a field has been set.

### GetMessageId

`func (o *WebhookEvent) GetMessageId() string`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *WebhookEvent) GetMessageIdOk() (*string, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *WebhookEvent) SetMessageId(v string)`

SetMessageId sets MessageId field to given value.

### HasMessageId

`func (o *WebhookEvent) HasMessageId() bool`

HasMessageId returns a boolean if a field has been set.

### GetAccountId

`func (o *WebhookEvent) GetAccountId() int32`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *WebhookEvent) GetAccountIdOk() (*int32, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *WebhookEvent) SetAccountId(v int32)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *WebhookEvent) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetDomainId

`func (o *WebhookEvent) GetDomainId() string`

GetDomainId returns the DomainId field if non-nil, zero value otherwise.

### GetDomainIdOk

`func (o *WebhookEvent) GetDomainIdOk() (*string, bool)`

GetDomainIdOk returns a tuple with the DomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainId

`func (o *WebhookEvent) SetDomainId(v string)`

SetDomainId sets DomainId field to given value.

### HasDomainId

`func (o *WebhookEvent) HasDomainId() bool`

HasDomainId returns a boolean if a field has been set.

### GetTo

`func (o *WebhookEvent) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *WebhookEvent) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *WebhookEvent) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *WebhookEvent) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetStatus

`func (o *WebhookEvent) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WebhookEvent) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WebhookEvent) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *WebhookEvent) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSmtpCode

`func (o *WebhookEvent) GetSmtpCode() int32`

GetSmtpCode returns the SmtpCode field if non-nil, zero value otherwise.

### GetSmtpCodeOk

`func (o *WebhookEvent) GetSmtpCodeOk() (*int32, bool)`

GetSmtpCodeOk returns a tuple with the SmtpCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtpCode

`func (o *WebhookEvent) SetSmtpCode(v int32)`

SetSmtpCode sets SmtpCode field to given value.

### HasSmtpCode

`func (o *WebhookEvent) HasSmtpCode() bool`

HasSmtpCode returns a boolean if a field has been set.

### GetSmtpResponse

`func (o *WebhookEvent) GetSmtpResponse() string`

GetSmtpResponse returns the SmtpResponse field if non-nil, zero value otherwise.

### GetSmtpResponseOk

`func (o *WebhookEvent) GetSmtpResponseOk() (*string, bool)`

GetSmtpResponseOk returns a tuple with the SmtpResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtpResponse

`func (o *WebhookEvent) SetSmtpResponse(v string)`

SetSmtpResponse sets SmtpResponse field to given value.

### HasSmtpResponse

`func (o *WebhookEvent) HasSmtpResponse() bool`

HasSmtpResponse returns a boolean if a field has been set.

### GetMxHost

`func (o *WebhookEvent) GetMxHost() string`

GetMxHost returns the MxHost field if non-nil, zero value otherwise.

### GetMxHostOk

`func (o *WebhookEvent) GetMxHostOk() (*string, bool)`

GetMxHostOk returns a tuple with the MxHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMxHost

`func (o *WebhookEvent) SetMxHost(v string)`

SetMxHost sets MxHost field to given value.

### HasMxHost

`func (o *WebhookEvent) HasMxHost() bool`

HasMxHost returns a boolean if a field has been set.

### GetBounceType

`func (o *WebhookEvent) GetBounceType() string`

GetBounceType returns the BounceType field if non-nil, zero value otherwise.

### GetBounceTypeOk

`func (o *WebhookEvent) GetBounceTypeOk() (*string, bool)`

GetBounceTypeOk returns a tuple with the BounceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBounceType

`func (o *WebhookEvent) SetBounceType(v string)`

SetBounceType sets BounceType field to given value.

### HasBounceType

`func (o *WebhookEvent) HasBounceType() bool`

HasBounceType returns a boolean if a field has been set.

### SetBounceTypeNil

`func (o *WebhookEvent) SetBounceTypeNil(b bool)`

 SetBounceTypeNil sets the value for BounceType to be an explicit nil

### UnsetBounceType
`func (o *WebhookEvent) UnsetBounceType()`

UnsetBounceType ensures that no value is present for BounceType, not even an explicit nil
### GetEnhancedStatusCode

`func (o *WebhookEvent) GetEnhancedStatusCode() string`

GetEnhancedStatusCode returns the EnhancedStatusCode field if non-nil, zero value otherwise.

### GetEnhancedStatusCodeOk

`func (o *WebhookEvent) GetEnhancedStatusCodeOk() (*string, bool)`

GetEnhancedStatusCodeOk returns a tuple with the EnhancedStatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnhancedStatusCode

`func (o *WebhookEvent) SetEnhancedStatusCode(v string)`

SetEnhancedStatusCode sets EnhancedStatusCode field to given value.

### HasEnhancedStatusCode

`func (o *WebhookEvent) HasEnhancedStatusCode() bool`

HasEnhancedStatusCode returns a boolean if a field has been set.

### GetAttempts

`func (o *WebhookEvent) GetAttempts() int32`

GetAttempts returns the Attempts field if non-nil, zero value otherwise.

### GetAttemptsOk

`func (o *WebhookEvent) GetAttemptsOk() (*int32, bool)`

GetAttemptsOk returns a tuple with the Attempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttempts

`func (o *WebhookEvent) SetAttempts(v int32)`

SetAttempts sets Attempts field to given value.

### HasAttempts

`func (o *WebhookEvent) HasAttempts() bool`

HasAttempts returns a boolean if a field has been set.

### GetIsp

`func (o *WebhookEvent) GetIsp() string`

GetIsp returns the Isp field if non-nil, zero value otherwise.

### GetIspOk

`func (o *WebhookEvent) GetIspOk() (*string, bool)`

GetIspOk returns a tuple with the Isp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsp

`func (o *WebhookEvent) SetIsp(v string)`

SetIsp sets Isp field to given value.

### HasIsp

`func (o *WebhookEvent) HasIsp() bool`

HasIsp returns a boolean if a field has been set.

### GetIsMpp

`func (o *WebhookEvent) GetIsMpp() bool`

GetIsMpp returns the IsMpp field if non-nil, zero value otherwise.

### GetIsMppOk

`func (o *WebhookEvent) GetIsMppOk() (*bool, bool)`

GetIsMppOk returns a tuple with the IsMpp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMpp

`func (o *WebhookEvent) SetIsMpp(v bool)`

SetIsMpp sets IsMpp field to given value.

### HasIsMpp

`func (o *WebhookEvent) HasIsMpp() bool`

HasIsMpp returns a boolean if a field has been set.

### GetIpAddress

`func (o *WebhookEvent) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *WebhookEvent) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *WebhookEvent) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *WebhookEvent) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetUserAgent

`func (o *WebhookEvent) GetUserAgent() string`

GetUserAgent returns the UserAgent field if non-nil, zero value otherwise.

### GetUserAgentOk

`func (o *WebhookEvent) GetUserAgentOk() (*string, bool)`

GetUserAgentOk returns a tuple with the UserAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAgent

`func (o *WebhookEvent) SetUserAgent(v string)`

SetUserAgent sets UserAgent field to given value.

### HasUserAgent

`func (o *WebhookEvent) HasUserAgent() bool`

HasUserAgent returns a boolean if a field has been set.

### GetIsBot

`func (o *WebhookEvent) GetIsBot() bool`

GetIsBot returns the IsBot field if non-nil, zero value otherwise.

### GetIsBotOk

`func (o *WebhookEvent) GetIsBotOk() (*bool, bool)`

GetIsBotOk returns a tuple with the IsBot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBot

`func (o *WebhookEvent) SetIsBot(v bool)`

SetIsBot sets IsBot field to given value.

### HasIsBot

`func (o *WebhookEvent) HasIsBot() bool`

HasIsBot returns a boolean if a field has been set.

### GetLinkUrl

`func (o *WebhookEvent) GetLinkUrl() string`

GetLinkUrl returns the LinkUrl field if non-nil, zero value otherwise.

### GetLinkUrlOk

`func (o *WebhookEvent) GetLinkUrlOk() (*string, bool)`

GetLinkUrlOk returns a tuple with the LinkUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkUrl

`func (o *WebhookEvent) SetLinkUrl(v string)`

SetLinkUrl sets LinkUrl field to given value.

### HasLinkUrl

`func (o *WebhookEvent) HasLinkUrl() bool`

HasLinkUrl returns a boolean if a field has been set.

### GetLinkIndex

`func (o *WebhookEvent) GetLinkIndex() int32`

GetLinkIndex returns the LinkIndex field if non-nil, zero value otherwise.

### GetLinkIndexOk

`func (o *WebhookEvent) GetLinkIndexOk() (*int32, bool)`

GetLinkIndexOk returns a tuple with the LinkIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkIndex

`func (o *WebhookEvent) SetLinkIndex(v int32)`

SetLinkIndex sets LinkIndex field to given value.

### HasLinkIndex

`func (o *WebhookEvent) HasLinkIndex() bool`

HasLinkIndex returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


