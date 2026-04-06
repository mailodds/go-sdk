# ConnectDnsProviderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Provider** | **string** | DNS provider | 
**ApiToken** | **string** | API token with Zone &gt; DNS &gt; Edit permission | 

## Methods

### NewConnectDnsProviderRequest

`func NewConnectDnsProviderRequest(provider string, apiToken string, ) *ConnectDnsProviderRequest`

NewConnectDnsProviderRequest instantiates a new ConnectDnsProviderRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectDnsProviderRequestWithDefaults

`func NewConnectDnsProviderRequestWithDefaults() *ConnectDnsProviderRequest`

NewConnectDnsProviderRequestWithDefaults instantiates a new ConnectDnsProviderRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProvider

`func (o *ConnectDnsProviderRequest) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *ConnectDnsProviderRequest) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *ConnectDnsProviderRequest) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetApiToken

`func (o *ConnectDnsProviderRequest) GetApiToken() string`

GetApiToken returns the ApiToken field if non-nil, zero value otherwise.

### GetApiTokenOk

`func (o *ConnectDnsProviderRequest) GetApiTokenOk() (*string, bool)`

GetApiTokenOk returns a tuple with the ApiToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiToken

`func (o *ConnectDnsProviderRequest) SetApiToken(v string)`

SetApiToken sets ApiToken field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


