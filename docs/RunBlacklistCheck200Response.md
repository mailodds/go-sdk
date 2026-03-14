# RunBlacklistCheck200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SchemaVersion** | Pointer to **string** |  | [optional] 
**RequestId** | Pointer to **string** |  | [optional] 
**Monitor** | Pointer to [**BlacklistMonitor**](BlacklistMonitor.md) |  | [optional] 
**Check** | Pointer to [**RunBlacklistCheck200ResponseCheck**](RunBlacklistCheck200ResponseCheck.md) |  | [optional] 

## Methods

### NewRunBlacklistCheck200Response

`func NewRunBlacklistCheck200Response() *RunBlacklistCheck200Response`

NewRunBlacklistCheck200Response instantiates a new RunBlacklistCheck200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRunBlacklistCheck200ResponseWithDefaults

`func NewRunBlacklistCheck200ResponseWithDefaults() *RunBlacklistCheck200Response`

NewRunBlacklistCheck200ResponseWithDefaults instantiates a new RunBlacklistCheck200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemaVersion

`func (o *RunBlacklistCheck200Response) GetSchemaVersion() string`

GetSchemaVersion returns the SchemaVersion field if non-nil, zero value otherwise.

### GetSchemaVersionOk

`func (o *RunBlacklistCheck200Response) GetSchemaVersionOk() (*string, bool)`

GetSchemaVersionOk returns a tuple with the SchemaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaVersion

`func (o *RunBlacklistCheck200Response) SetSchemaVersion(v string)`

SetSchemaVersion sets SchemaVersion field to given value.

### HasSchemaVersion

`func (o *RunBlacklistCheck200Response) HasSchemaVersion() bool`

HasSchemaVersion returns a boolean if a field has been set.

### GetRequestId

`func (o *RunBlacklistCheck200Response) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *RunBlacklistCheck200Response) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *RunBlacklistCheck200Response) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *RunBlacklistCheck200Response) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetMonitor

`func (o *RunBlacklistCheck200Response) GetMonitor() BlacklistMonitor`

GetMonitor returns the Monitor field if non-nil, zero value otherwise.

### GetMonitorOk

`func (o *RunBlacklistCheck200Response) GetMonitorOk() (*BlacklistMonitor, bool)`

GetMonitorOk returns a tuple with the Monitor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitor

`func (o *RunBlacklistCheck200Response) SetMonitor(v BlacklistMonitor)`

SetMonitor sets Monitor field to given value.

### HasMonitor

`func (o *RunBlacklistCheck200Response) HasMonitor() bool`

HasMonitor returns a boolean if a field has been set.

### GetCheck

`func (o *RunBlacklistCheck200Response) GetCheck() RunBlacklistCheck200ResponseCheck`

GetCheck returns the Check field if non-nil, zero value otherwise.

### GetCheckOk

`func (o *RunBlacklistCheck200Response) GetCheckOk() (*RunBlacklistCheck200ResponseCheck, bool)`

GetCheckOk returns a tuple with the Check field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheck

`func (o *RunBlacklistCheck200Response) SetCheck(v RunBlacklistCheck200ResponseCheck)`

SetCheck sets Check field to given value.

### HasCheck

`func (o *RunBlacklistCheck200Response) HasCheck() bool`

HasCheck returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


