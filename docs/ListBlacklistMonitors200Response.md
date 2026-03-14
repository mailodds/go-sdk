# ListBlacklistMonitors200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SchemaVersion** | Pointer to **string** |  | [optional] 
**RequestId** | Pointer to **string** |  | [optional] 
**Monitors** | Pointer to [**[]BlacklistMonitor**](BlacklistMonitor.md) |  | [optional] 

## Methods

### NewListBlacklistMonitors200Response

`func NewListBlacklistMonitors200Response() *ListBlacklistMonitors200Response`

NewListBlacklistMonitors200Response instantiates a new ListBlacklistMonitors200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListBlacklistMonitors200ResponseWithDefaults

`func NewListBlacklistMonitors200ResponseWithDefaults() *ListBlacklistMonitors200Response`

NewListBlacklistMonitors200ResponseWithDefaults instantiates a new ListBlacklistMonitors200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemaVersion

`func (o *ListBlacklistMonitors200Response) GetSchemaVersion() string`

GetSchemaVersion returns the SchemaVersion field if non-nil, zero value otherwise.

### GetSchemaVersionOk

`func (o *ListBlacklistMonitors200Response) GetSchemaVersionOk() (*string, bool)`

GetSchemaVersionOk returns a tuple with the SchemaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaVersion

`func (o *ListBlacklistMonitors200Response) SetSchemaVersion(v string)`

SetSchemaVersion sets SchemaVersion field to given value.

### HasSchemaVersion

`func (o *ListBlacklistMonitors200Response) HasSchemaVersion() bool`

HasSchemaVersion returns a boolean if a field has been set.

### GetRequestId

`func (o *ListBlacklistMonitors200Response) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *ListBlacklistMonitors200Response) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *ListBlacklistMonitors200Response) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *ListBlacklistMonitors200Response) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetMonitors

`func (o *ListBlacklistMonitors200Response) GetMonitors() []BlacklistMonitor`

GetMonitors returns the Monitors field if non-nil, zero value otherwise.

### GetMonitorsOk

`func (o *ListBlacklistMonitors200Response) GetMonitorsOk() (*[]BlacklistMonitor, bool)`

GetMonitorsOk returns a tuple with the Monitors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitors

`func (o *ListBlacklistMonitors200Response) SetMonitors(v []BlacklistMonitor)`

SetMonitors sets Monitors field to given value.

### HasMonitors

`func (o *ListBlacklistMonitors200Response) HasMonitors() bool`

HasMonitors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


