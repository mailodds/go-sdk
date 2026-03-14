# GetMessageEvents200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SchemaVersion** | Pointer to **string** |  | [optional] 
**RequestId** | Pointer to **string** |  | [optional] 
**MessageId** | Pointer to **string** |  | [optional] 
**Recipient** | Pointer to **string** |  | [optional] 
**Summary** | Pointer to [**GetMessageEvents200ResponseSummary**](GetMessageEvents200ResponseSummary.md) |  | [optional] 
**Clicks** | Pointer to [**[]GetMessageEvents200ResponseClicksInner**](GetMessageEvents200ResponseClicksInner.md) |  | [optional] 
**Events** | Pointer to [**[]GetMessageEvents200ResponseEventsInner**](GetMessageEvents200ResponseEventsInner.md) |  | [optional] 
**Total** | Pointer to **int32** |  | [optional] 

## Methods

### NewGetMessageEvents200Response

`func NewGetMessageEvents200Response() *GetMessageEvents200Response`

NewGetMessageEvents200Response instantiates a new GetMessageEvents200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMessageEvents200ResponseWithDefaults

`func NewGetMessageEvents200ResponseWithDefaults() *GetMessageEvents200Response`

NewGetMessageEvents200ResponseWithDefaults instantiates a new GetMessageEvents200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemaVersion

`func (o *GetMessageEvents200Response) GetSchemaVersion() string`

GetSchemaVersion returns the SchemaVersion field if non-nil, zero value otherwise.

### GetSchemaVersionOk

`func (o *GetMessageEvents200Response) GetSchemaVersionOk() (*string, bool)`

GetSchemaVersionOk returns a tuple with the SchemaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaVersion

`func (o *GetMessageEvents200Response) SetSchemaVersion(v string)`

SetSchemaVersion sets SchemaVersion field to given value.

### HasSchemaVersion

`func (o *GetMessageEvents200Response) HasSchemaVersion() bool`

HasSchemaVersion returns a boolean if a field has been set.

### GetRequestId

`func (o *GetMessageEvents200Response) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *GetMessageEvents200Response) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *GetMessageEvents200Response) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *GetMessageEvents200Response) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetMessageId

`func (o *GetMessageEvents200Response) GetMessageId() string`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *GetMessageEvents200Response) GetMessageIdOk() (*string, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *GetMessageEvents200Response) SetMessageId(v string)`

SetMessageId sets MessageId field to given value.

### HasMessageId

`func (o *GetMessageEvents200Response) HasMessageId() bool`

HasMessageId returns a boolean if a field has been set.

### GetRecipient

`func (o *GetMessageEvents200Response) GetRecipient() string`

GetRecipient returns the Recipient field if non-nil, zero value otherwise.

### GetRecipientOk

`func (o *GetMessageEvents200Response) GetRecipientOk() (*string, bool)`

GetRecipientOk returns a tuple with the Recipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipient

`func (o *GetMessageEvents200Response) SetRecipient(v string)`

SetRecipient sets Recipient field to given value.

### HasRecipient

`func (o *GetMessageEvents200Response) HasRecipient() bool`

HasRecipient returns a boolean if a field has been set.

### GetSummary

`func (o *GetMessageEvents200Response) GetSummary() GetMessageEvents200ResponseSummary`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *GetMessageEvents200Response) GetSummaryOk() (*GetMessageEvents200ResponseSummary, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *GetMessageEvents200Response) SetSummary(v GetMessageEvents200ResponseSummary)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *GetMessageEvents200Response) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetClicks

`func (o *GetMessageEvents200Response) GetClicks() []GetMessageEvents200ResponseClicksInner`

GetClicks returns the Clicks field if non-nil, zero value otherwise.

### GetClicksOk

`func (o *GetMessageEvents200Response) GetClicksOk() (*[]GetMessageEvents200ResponseClicksInner, bool)`

GetClicksOk returns a tuple with the Clicks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClicks

`func (o *GetMessageEvents200Response) SetClicks(v []GetMessageEvents200ResponseClicksInner)`

SetClicks sets Clicks field to given value.

### HasClicks

`func (o *GetMessageEvents200Response) HasClicks() bool`

HasClicks returns a boolean if a field has been set.

### GetEvents

`func (o *GetMessageEvents200Response) GetEvents() []GetMessageEvents200ResponseEventsInner`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *GetMessageEvents200Response) GetEventsOk() (*[]GetMessageEvents200ResponseEventsInner, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *GetMessageEvents200Response) SetEvents(v []GetMessageEvents200ResponseEventsInner)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *GetMessageEvents200Response) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetTotal

`func (o *GetMessageEvents200Response) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GetMessageEvents200Response) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GetMessageEvents200Response) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *GetMessageEvents200Response) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


