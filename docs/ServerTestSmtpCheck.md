# ServerTestSmtpCheck

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Connectable** | Pointer to **bool** |  | [optional] 
**Banner** | Pointer to **string** |  | [optional] 
**Starttls** | Pointer to **bool** |  | [optional] 
**ResponseTimeMs** | Pointer to **int32** |  | [optional] 

## Methods

### NewServerTestSmtpCheck

`func NewServerTestSmtpCheck() *ServerTestSmtpCheck`

NewServerTestSmtpCheck instantiates a new ServerTestSmtpCheck object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerTestSmtpCheckWithDefaults

`func NewServerTestSmtpCheckWithDefaults() *ServerTestSmtpCheck`

NewServerTestSmtpCheckWithDefaults instantiates a new ServerTestSmtpCheck object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectable

`func (o *ServerTestSmtpCheck) GetConnectable() bool`

GetConnectable returns the Connectable field if non-nil, zero value otherwise.

### GetConnectableOk

`func (o *ServerTestSmtpCheck) GetConnectableOk() (*bool, bool)`

GetConnectableOk returns a tuple with the Connectable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectable

`func (o *ServerTestSmtpCheck) SetConnectable(v bool)`

SetConnectable sets Connectable field to given value.

### HasConnectable

`func (o *ServerTestSmtpCheck) HasConnectable() bool`

HasConnectable returns a boolean if a field has been set.

### GetBanner

`func (o *ServerTestSmtpCheck) GetBanner() string`

GetBanner returns the Banner field if non-nil, zero value otherwise.

### GetBannerOk

`func (o *ServerTestSmtpCheck) GetBannerOk() (*string, bool)`

GetBannerOk returns a tuple with the Banner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBanner

`func (o *ServerTestSmtpCheck) SetBanner(v string)`

SetBanner sets Banner field to given value.

### HasBanner

`func (o *ServerTestSmtpCheck) HasBanner() bool`

HasBanner returns a boolean if a field has been set.

### GetStarttls

`func (o *ServerTestSmtpCheck) GetStarttls() bool`

GetStarttls returns the Starttls field if non-nil, zero value otherwise.

### GetStarttlsOk

`func (o *ServerTestSmtpCheck) GetStarttlsOk() (*bool, bool)`

GetStarttlsOk returns a tuple with the Starttls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStarttls

`func (o *ServerTestSmtpCheck) SetStarttls(v bool)`

SetStarttls sets Starttls field to given value.

### HasStarttls

`func (o *ServerTestSmtpCheck) HasStarttls() bool`

HasStarttls returns a boolean if a field has been set.

### GetResponseTimeMs

`func (o *ServerTestSmtpCheck) GetResponseTimeMs() int32`

GetResponseTimeMs returns the ResponseTimeMs field if non-nil, zero value otherwise.

### GetResponseTimeMsOk

`func (o *ServerTestSmtpCheck) GetResponseTimeMsOk() (*int32, bool)`

GetResponseTimeMsOk returns a tuple with the ResponseTimeMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseTimeMs

`func (o *ServerTestSmtpCheck) SetResponseTimeMs(v int32)`

SetResponseTimeMs sets ResponseTimeMs field to given value.

### HasResponseTimeMs

`func (o *ServerTestSmtpCheck) HasResponseTimeMs() bool`

HasResponseTimeMs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


