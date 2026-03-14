# GetDmarcSources200ResponseSourcesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceIp** | Pointer to **string** |  | [optional] 
**Hostname** | Pointer to **string** |  | [optional] 
**Org** | Pointer to **string** |  | [optional] 
**Count** | Pointer to **int32** |  | [optional] 
**DkimPass** | Pointer to **int32** |  | [optional] 
**DkimFail** | Pointer to **int32** |  | [optional] 
**SpfPass** | Pointer to **int32** |  | [optional] 
**SpfFail** | Pointer to **int32** |  | [optional] 
**Disposition** | Pointer to **string** |  | [optional] 

## Methods

### NewGetDmarcSources200ResponseSourcesInner

`func NewGetDmarcSources200ResponseSourcesInner() *GetDmarcSources200ResponseSourcesInner`

NewGetDmarcSources200ResponseSourcesInner instantiates a new GetDmarcSources200ResponseSourcesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDmarcSources200ResponseSourcesInnerWithDefaults

`func NewGetDmarcSources200ResponseSourcesInnerWithDefaults() *GetDmarcSources200ResponseSourcesInner`

NewGetDmarcSources200ResponseSourcesInnerWithDefaults instantiates a new GetDmarcSources200ResponseSourcesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceIp

`func (o *GetDmarcSources200ResponseSourcesInner) GetSourceIp() string`

GetSourceIp returns the SourceIp field if non-nil, zero value otherwise.

### GetSourceIpOk

`func (o *GetDmarcSources200ResponseSourcesInner) GetSourceIpOk() (*string, bool)`

GetSourceIpOk returns a tuple with the SourceIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceIp

`func (o *GetDmarcSources200ResponseSourcesInner) SetSourceIp(v string)`

SetSourceIp sets SourceIp field to given value.

### HasSourceIp

`func (o *GetDmarcSources200ResponseSourcesInner) HasSourceIp() bool`

HasSourceIp returns a boolean if a field has been set.

### GetHostname

`func (o *GetDmarcSources200ResponseSourcesInner) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *GetDmarcSources200ResponseSourcesInner) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *GetDmarcSources200ResponseSourcesInner) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *GetDmarcSources200ResponseSourcesInner) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetOrg

`func (o *GetDmarcSources200ResponseSourcesInner) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *GetDmarcSources200ResponseSourcesInner) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *GetDmarcSources200ResponseSourcesInner) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *GetDmarcSources200ResponseSourcesInner) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetCount

`func (o *GetDmarcSources200ResponseSourcesInner) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *GetDmarcSources200ResponseSourcesInner) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *GetDmarcSources200ResponseSourcesInner) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *GetDmarcSources200ResponseSourcesInner) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetDkimPass

`func (o *GetDmarcSources200ResponseSourcesInner) GetDkimPass() int32`

GetDkimPass returns the DkimPass field if non-nil, zero value otherwise.

### GetDkimPassOk

`func (o *GetDmarcSources200ResponseSourcesInner) GetDkimPassOk() (*int32, bool)`

GetDkimPassOk returns a tuple with the DkimPass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDkimPass

`func (o *GetDmarcSources200ResponseSourcesInner) SetDkimPass(v int32)`

SetDkimPass sets DkimPass field to given value.

### HasDkimPass

`func (o *GetDmarcSources200ResponseSourcesInner) HasDkimPass() bool`

HasDkimPass returns a boolean if a field has been set.

### GetDkimFail

`func (o *GetDmarcSources200ResponseSourcesInner) GetDkimFail() int32`

GetDkimFail returns the DkimFail field if non-nil, zero value otherwise.

### GetDkimFailOk

`func (o *GetDmarcSources200ResponseSourcesInner) GetDkimFailOk() (*int32, bool)`

GetDkimFailOk returns a tuple with the DkimFail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDkimFail

`func (o *GetDmarcSources200ResponseSourcesInner) SetDkimFail(v int32)`

SetDkimFail sets DkimFail field to given value.

### HasDkimFail

`func (o *GetDmarcSources200ResponseSourcesInner) HasDkimFail() bool`

HasDkimFail returns a boolean if a field has been set.

### GetSpfPass

`func (o *GetDmarcSources200ResponseSourcesInner) GetSpfPass() int32`

GetSpfPass returns the SpfPass field if non-nil, zero value otherwise.

### GetSpfPassOk

`func (o *GetDmarcSources200ResponseSourcesInner) GetSpfPassOk() (*int32, bool)`

GetSpfPassOk returns a tuple with the SpfPass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpfPass

`func (o *GetDmarcSources200ResponseSourcesInner) SetSpfPass(v int32)`

SetSpfPass sets SpfPass field to given value.

### HasSpfPass

`func (o *GetDmarcSources200ResponseSourcesInner) HasSpfPass() bool`

HasSpfPass returns a boolean if a field has been set.

### GetSpfFail

`func (o *GetDmarcSources200ResponseSourcesInner) GetSpfFail() int32`

GetSpfFail returns the SpfFail field if non-nil, zero value otherwise.

### GetSpfFailOk

`func (o *GetDmarcSources200ResponseSourcesInner) GetSpfFailOk() (*int32, bool)`

GetSpfFailOk returns a tuple with the SpfFail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpfFail

`func (o *GetDmarcSources200ResponseSourcesInner) SetSpfFail(v int32)`

SetSpfFail sets SpfFail field to given value.

### HasSpfFail

`func (o *GetDmarcSources200ResponseSourcesInner) HasSpfFail() bool`

HasSpfFail returns a boolean if a field has been set.

### GetDisposition

`func (o *GetDmarcSources200ResponseSourcesInner) GetDisposition() string`

GetDisposition returns the Disposition field if non-nil, zero value otherwise.

### GetDispositionOk

`func (o *GetDmarcSources200ResponseSourcesInner) GetDispositionOk() (*string, bool)`

GetDispositionOk returns a tuple with the Disposition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisposition

`func (o *GetDmarcSources200ResponseSourcesInner) SetDisposition(v string)`

SetDisposition sets Disposition field to given value.

### HasDisposition

`func (o *GetDmarcSources200ResponseSourcesInner) HasDisposition() bool`

HasDisposition returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


