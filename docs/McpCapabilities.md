# McpCapabilities

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | Pointer to **string** |  | [optional] 
**ServerName** | Pointer to **string** |  | [optional] 
**ToolCount** | Pointer to **int32** | Total number of available tools | [optional] 
**Pillars** | Pointer to [**[]McpCapabilitiesPillarsInner**](McpCapabilitiesPillarsInner.md) |  | [optional] 
**SupportedTransports** | Pointer to **[]string** |  | [optional] 
**AuthRequired** | Pointer to **bool** |  | [optional] 

## Methods

### NewMcpCapabilities

`func NewMcpCapabilities() *McpCapabilities`

NewMcpCapabilities instantiates a new McpCapabilities object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMcpCapabilitiesWithDefaults

`func NewMcpCapabilitiesWithDefaults() *McpCapabilities`

NewMcpCapabilitiesWithDefaults instantiates a new McpCapabilities object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *McpCapabilities) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *McpCapabilities) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *McpCapabilities) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *McpCapabilities) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetServerName

`func (o *McpCapabilities) GetServerName() string`

GetServerName returns the ServerName field if non-nil, zero value otherwise.

### GetServerNameOk

`func (o *McpCapabilities) GetServerNameOk() (*string, bool)`

GetServerNameOk returns a tuple with the ServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerName

`func (o *McpCapabilities) SetServerName(v string)`

SetServerName sets ServerName field to given value.

### HasServerName

`func (o *McpCapabilities) HasServerName() bool`

HasServerName returns a boolean if a field has been set.

### GetToolCount

`func (o *McpCapabilities) GetToolCount() int32`

GetToolCount returns the ToolCount field if non-nil, zero value otherwise.

### GetToolCountOk

`func (o *McpCapabilities) GetToolCountOk() (*int32, bool)`

GetToolCountOk returns a tuple with the ToolCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolCount

`func (o *McpCapabilities) SetToolCount(v int32)`

SetToolCount sets ToolCount field to given value.

### HasToolCount

`func (o *McpCapabilities) HasToolCount() bool`

HasToolCount returns a boolean if a field has been set.

### GetPillars

`func (o *McpCapabilities) GetPillars() []McpCapabilitiesPillarsInner`

GetPillars returns the Pillars field if non-nil, zero value otherwise.

### GetPillarsOk

`func (o *McpCapabilities) GetPillarsOk() (*[]McpCapabilitiesPillarsInner, bool)`

GetPillarsOk returns a tuple with the Pillars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPillars

`func (o *McpCapabilities) SetPillars(v []McpCapabilitiesPillarsInner)`

SetPillars sets Pillars field to given value.

### HasPillars

`func (o *McpCapabilities) HasPillars() bool`

HasPillars returns a boolean if a field has been set.

### GetSupportedTransports

`func (o *McpCapabilities) GetSupportedTransports() []string`

GetSupportedTransports returns the SupportedTransports field if non-nil, zero value otherwise.

### GetSupportedTransportsOk

`func (o *McpCapabilities) GetSupportedTransportsOk() (*[]string, bool)`

GetSupportedTransportsOk returns a tuple with the SupportedTransports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedTransports

`func (o *McpCapabilities) SetSupportedTransports(v []string)`

SetSupportedTransports sets SupportedTransports field to given value.

### HasSupportedTransports

`func (o *McpCapabilities) HasSupportedTransports() bool`

HasSupportedTransports returns a boolean if a field has been set.

### GetAuthRequired

`func (o *McpCapabilities) GetAuthRequired() bool`

GetAuthRequired returns the AuthRequired field if non-nil, zero value otherwise.

### GetAuthRequiredOk

`func (o *McpCapabilities) GetAuthRequiredOk() (*bool, bool)`

GetAuthRequiredOk returns a tuple with the AuthRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthRequired

`func (o *McpCapabilities) SetAuthRequired(v bool)`

SetAuthRequired sets AuthRequired field to given value.

### HasAuthRequired

`func (o *McpCapabilities) HasAuthRequired() bool`

HasAuthRequired returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


