# ServerTestDnsChecks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Spf** | Pointer to **bool** |  | [optional] 
**Dkim** | Pointer to **bool** |  | [optional] 
**Dmarc** | Pointer to **bool** |  | [optional] 
**DmarcPolicy** | Pointer to **string** |  | [optional] 

## Methods

### NewServerTestDnsChecks

`func NewServerTestDnsChecks() *ServerTestDnsChecks`

NewServerTestDnsChecks instantiates a new ServerTestDnsChecks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerTestDnsChecksWithDefaults

`func NewServerTestDnsChecksWithDefaults() *ServerTestDnsChecks`

NewServerTestDnsChecksWithDefaults instantiates a new ServerTestDnsChecks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSpf

`func (o *ServerTestDnsChecks) GetSpf() bool`

GetSpf returns the Spf field if non-nil, zero value otherwise.

### GetSpfOk

`func (o *ServerTestDnsChecks) GetSpfOk() (*bool, bool)`

GetSpfOk returns a tuple with the Spf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpf

`func (o *ServerTestDnsChecks) SetSpf(v bool)`

SetSpf sets Spf field to given value.

### HasSpf

`func (o *ServerTestDnsChecks) HasSpf() bool`

HasSpf returns a boolean if a field has been set.

### GetDkim

`func (o *ServerTestDnsChecks) GetDkim() bool`

GetDkim returns the Dkim field if non-nil, zero value otherwise.

### GetDkimOk

`func (o *ServerTestDnsChecks) GetDkimOk() (*bool, bool)`

GetDkimOk returns a tuple with the Dkim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDkim

`func (o *ServerTestDnsChecks) SetDkim(v bool)`

SetDkim sets Dkim field to given value.

### HasDkim

`func (o *ServerTestDnsChecks) HasDkim() bool`

HasDkim returns a boolean if a field has been set.

### GetDmarc

`func (o *ServerTestDnsChecks) GetDmarc() bool`

GetDmarc returns the Dmarc field if non-nil, zero value otherwise.

### GetDmarcOk

`func (o *ServerTestDnsChecks) GetDmarcOk() (*bool, bool)`

GetDmarcOk returns a tuple with the Dmarc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDmarc

`func (o *ServerTestDnsChecks) SetDmarc(v bool)`

SetDmarc sets Dmarc field to given value.

### HasDmarc

`func (o *ServerTestDnsChecks) HasDmarc() bool`

HasDmarc returns a boolean if a field has been set.

### GetDmarcPolicy

`func (o *ServerTestDnsChecks) GetDmarcPolicy() string`

GetDmarcPolicy returns the DmarcPolicy field if non-nil, zero value otherwise.

### GetDmarcPolicyOk

`func (o *ServerTestDnsChecks) GetDmarcPolicyOk() (*string, bool)`

GetDmarcPolicyOk returns a tuple with the DmarcPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDmarcPolicy

`func (o *ServerTestDnsChecks) SetDmarcPolicy(v string)`

SetDmarcPolicy sets DmarcPolicy field to given value.

### HasDmarcPolicy

`func (o *ServerTestDnsChecks) HasDmarcPolicy() bool`

HasDmarcPolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


