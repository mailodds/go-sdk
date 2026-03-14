# CreateBounceAnalysisRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DomainId** | **string** | Sending domain UUID to analyze bounces for | 
**Period** | Pointer to **string** | Time period to analyze | [optional] [default to "30d"]

## Methods

### NewCreateBounceAnalysisRequest

`func NewCreateBounceAnalysisRequest(domainId string, ) *CreateBounceAnalysisRequest`

NewCreateBounceAnalysisRequest instantiates a new CreateBounceAnalysisRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateBounceAnalysisRequestWithDefaults

`func NewCreateBounceAnalysisRequestWithDefaults() *CreateBounceAnalysisRequest`

NewCreateBounceAnalysisRequestWithDefaults instantiates a new CreateBounceAnalysisRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomainId

`func (o *CreateBounceAnalysisRequest) GetDomainId() string`

GetDomainId returns the DomainId field if non-nil, zero value otherwise.

### GetDomainIdOk

`func (o *CreateBounceAnalysisRequest) GetDomainIdOk() (*string, bool)`

GetDomainIdOk returns a tuple with the DomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainId

`func (o *CreateBounceAnalysisRequest) SetDomainId(v string)`

SetDomainId sets DomainId field to given value.


### GetPeriod

`func (o *CreateBounceAnalysisRequest) GetPeriod() string`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *CreateBounceAnalysisRequest) GetPeriodOk() (*string, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *CreateBounceAnalysisRequest) SetPeriod(v string)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *CreateBounceAnalysisRequest) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


